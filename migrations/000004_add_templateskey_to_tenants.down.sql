DROP INDEX IF EXISTS idx_tenants_template_id;
ALTER TABLE tenants DROP COLUMN IF EXISTS template_id;
