CREATE TABLE brands (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(255) UNIQUE NOT NULL,
	icon_path TEXT NULL
);