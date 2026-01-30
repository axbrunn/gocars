-- Add CHECK constraint for slug format
ALTER TABLE tenants
ADD CONSTRAINT tenants_slug_format_check
CHECK (slug ~ '^[a-z0-9-]+$');

-- Add CHECK constraint for non-empty name
ALTER TABLE tenants
ADD CONSTRAINT tenants_name_not_empty_check
CHECK (length(trim(name)) > 0);

-- Optional but recommended index (if not already present via UNIQUE)
CREATE INDEX IF NOT EXISTS idx_tenants_slug ON tenants(slug);
