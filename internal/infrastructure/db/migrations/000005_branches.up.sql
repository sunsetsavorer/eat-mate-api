CREATE TABLE branches (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	brand_id UUID NOT NULL REFERENCES brands(id) ON DELETE CASCADE ON UPDATE CASCADE,
	address VARCHAR(255) NULL,
	contact_phone VARCHAR(20) NULL,
	coordinates JSON NULL
);