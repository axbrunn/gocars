ALTER TABLE tenants
DROP CONSTRAINT IF EXISTS tenants_slug_format_check;

ALTER TABLE tenants
DROP CONSTRAINT IF EXISTS tenants_name_not_empty_check;

DROP INDEX IF EXISTS idx_tenants_slug;
