CREATE TABLE tenants (
    id UUID PRIMARY KEY,
    slug TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_tenants_slug ON tenants(slug);

INSERT INTO tenants (id, slug, name)
VALUES
  (gen_random_uuid(), 'tenant1', 'Tenant One'),
  (gen_random_uuid(), 'tenant2', 'Tenant Two');
