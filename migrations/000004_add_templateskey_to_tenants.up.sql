ALTER TABLE tenants
ADD COLUMN template_id UUID REFERENCES templates(id);

UPDATE tenants
SET template_id = (SELECT id FROM templates WHERE slug = 'modern');

ALTER TABLE tenants
ALTER COLUMN template_id SET NOT NULL;

CREATE INDEX idx_tenants_template_id ON tenants(template_id);
