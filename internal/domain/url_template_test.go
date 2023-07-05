//go:build unit

package domain

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	caos_errs "github.com/zitadel/zitadel/internal/errors"
)

func Test_renderURLTemplate(t *testing.T) {
	type args struct {
		tmpl string
		data any
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr error
	}{
		{
			name: "parse error",
			args: args{
				tmpl: "{{",
			},
			wantErr: caos_errs.ThrowInvalidArgument(nil, "DOMAIN-oGh5e", "Errors.User.InvalidURLTemplate"),
		},
		{
			name: "execution error",
			args: args{
				tmpl: "{{.Some}}",
				data: struct{ Foo int }{Foo: 1},
			},
			wantErr: caos_errs.ThrowInvalidArgument(nil, "DOMAIN-ieYa7", "Errors.User.InvalidURLTemplate"),
		},
		{
			name: "success",
			args: args{
				tmpl: "{{.Foo}}",
				data: struct{ Foo int }{Foo: 1},
			},
			wantW: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := renderURLTemplate(w, tt.args.tmpl, tt.args.data)
			require.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.wantW, w.String())
		})
	}
}
