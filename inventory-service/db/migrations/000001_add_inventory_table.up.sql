CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS inventory(
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
product_id VARCHAR(24) NOT NULL,
properties JSONB NOT NULL,
available_stock INT DEFAULT 0,
reserved_stock INT DEFAULT 0
);