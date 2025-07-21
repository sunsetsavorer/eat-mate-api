CREATE TABLE IF NOT EXISTS place_branches (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	place_id UUID NOT NULL REFERENCES places(id) ON DELETE CASCADE ON UPDATE CASCADE,
	address VARCHAR(255) NULL,
	contact_phone VARCHAR(20) NULL,
	coordinates JSON NULL
);