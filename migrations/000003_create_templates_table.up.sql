CREATE TABLE templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_templates_slug ON templates(slug);

ALTER TABLE templates
ADD CONSTRAINT templates_slug_format_check
CHECK (slug ~ '^[a-z0-9-]+$');

ALTER TABLE templates
ADD CONSTRAINT templates_name_not_empty_check
CHECK (length(trim(name)) > 0);

INSERT INTO templates (slug, name, description)
VALUES
  ('modern', 'Modern', 'Strak en modern design met grote afbeeldingen'),
  ('classic', 'Classic', 'Tijdloos klassiek design voor de traditionele dealer');
