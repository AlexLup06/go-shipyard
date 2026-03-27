-- +migrate Up
CREATE SCHEMA IF NOT EXISTS __APP_SLUG_DB__;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION __APP_SLUG_DB__.set_updated_at()
RETURNS trigger AS $func$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$func$ LANGUAGE plpgsql;
-- +migrate StatementEnd

CREATE TABLE IF NOT EXISTS public.__APP_SLUG_DB___schema_version (
  version INTEGER NOT NULL PRIMARY KEY
);

INSERT INTO public.__APP_SLUG_DB___schema_version (version)
VALUES (1)
ON CONFLICT (version) DO NOTHING;


-- +migrate Down
DROP FUNCTION IF EXISTS __APP_SLUG_DB__.set_updated_at();
DROP TABLE IF EXISTS __APP_SLUG_DB__.schema_version;
DROP SCHEMA IF EXISTS __APP_SLUG_DB__ CASCADE;
