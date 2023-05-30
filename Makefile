go_bin := "$$(go env GOPATH)/bin"
gen_authopt_path := "$(go_bin)/protoc-gen-authoption"
gen_zitadel_path := "$(go_bin)/protoc-gen-zitadel"

.PHONY: compile
compile: core_build console_build compile_pipeline

.PHONY: compile_pipeline
compile_pipeline:
	cp -r console/dist/console internal/api/ui/console/static/
	go build -o zitadel-$$(go env GOOS)-$$(go env GOARCH) -ldflags="-s -w"

.PHONY: core_dependencies
core_dependencies:
	go mod download

.PHONY: core_static
core_static:
	go install github.com/rakyll/statik@v0.1.7
	go generate internal/api/ui/login/statik/generate.go
	go generate internal/api/ui/login/static/resources/generate.go
	go generate internal/notification/statik/generate.go
	go generate internal/statik/generate.go

.PHONY: core_assets
core_assets:
	mkdir -p docs/apis/assets
	go run internal/api/assets/generator/asset_generator.go -directory=internal/api/assets/generator/ -assets=docs/apis/assets/assets.md

.PHONY: core_api_generator
core_api_generator:
ifeq (,$(wildcard $(gen_authopt_path)))
	go install internal/protoc/protoc-gen-authoption/main.go \
    && mv $$(go env GOPATH)/bin/main $(gen_authopt_path)
endif
ifeq (,$(wildcard $(gen_zitadel_path)))
	go install internal/protoc/protoc-gen-zitadel/main.go \
    && mv $$(go env GOPATH)/bin/main $(gen_zitadel_path)
endif

.PHONY: core_grpc_dependencies
core_grpc_dependencies:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30 
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3 
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2 
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2 
	go install github.com/envoyproxy/protoc-gen-validate@v0.10.1

.PHONY: core_api
core_api: core_api_generator core_grpc_dependencies
	buf generate
	mkdir -p pkg/grpc
	cp -r .artifacts/grpc/github.com/zitadel/zitadel/pkg/grpc/* pkg/grpc/
	mkdir -p openapi/v2/zitadel
	cp -r .artifacts/grpc/zitadel/ openapi/v2/zitadel

.PHONY: core_build
core_build: core_dependencies core_api core_static core_assets

.PHONY: console_dependencies
console_dependencies:
	cd console && \
	yarn install --frozen-lockfile

.PHONY: console_client
console_client:
	cd console && \
	yarn generate

.PHONY: console_build
console_build: console_dependencies console_client
	cd console && \
	yarn build

.PHONY: clean
clean:
	$(RM) .artifacts/grpc
	$(RM) $(gen_authopt_path)
	$(RM) $(gen_zitadel_path)

.PHONY: core_unit_test
core_unit_test:
	go test -race -v -coverprofile=profile.cov ./...

.PHONY: core_integration_test
core_integration_test:
	go run main.go init --config internal/integration/config/zitadel.yaml --config internal/integration/config/${INTEGRATION_DB_FLAVOR}.yaml
	go run main.go setup --masterkeyFromEnv --config internal/integration/config/zitadel.yaml --config internal/integration/config/${INTEGRATION_DB_FLAVOR}.yaml
	go test -tags=integration -race -parallel 1 -v -coverprofile=profile.cov -coverpkg=./... ./internal/integration ./internal/api/grpc/...

.PHONY: console_lint
console_lint:
	cd console && \
	yarn lint

# .PHONE: core_lint
# core_lint:
# 	golangci-lint run \
# 		--timeout 10m \
# 		--config ./.golangci.yaml \
# 		--out-format=github-actions \
# 		--concurrency=$$(getconf _NPROCESSORS_ONLN)