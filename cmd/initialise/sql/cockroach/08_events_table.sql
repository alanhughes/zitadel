SET experimental_enable_hash_sharded_indexes = on;

CREATE TABLE IF NOT EXISTS eventstore.events (
	id UUID DEFAULT gen_random_uuid()
	, event_type TEXT NOT NULL
	, aggregate_type TEXT NOT NULL
	, aggregate_id TEXT NOT NULL
	, aggregate_version TEXT NOT NULL
	, event_sequence BIGINT NOT NULL
	, previous_aggregate_sequence BIGINT
	, previous_aggregate_type_sequence INT8
	, creation_date TIMESTAMPTZ NOT NULL DEFAULT now()
	, created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
	, event_data JSONB
	, editor_user TEXT NOT NULL 
	, editor_service TEXT
	, resource_owner TEXT NOT NULL
	, instance_id TEXT NOT NULL
	, global_sequence SERIAL

	, PRIMARY KEY (instance_id, aggregate_type, aggregate_id, event_sequence DESC)
);