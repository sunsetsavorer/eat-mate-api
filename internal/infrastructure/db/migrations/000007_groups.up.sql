CREATE TABLE groups (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(50) NOT NULL,
	is_public BOOLEAN NOT NULL DEFAULT true,
	is_active BOOLEAN NOT NULL DEFAULT true,
	selection_mode group_mode NOT NULL,
	branch_id UUID NULL REFERENCES branches(id) ON DELETE SET NULL ON UPDATE CASCADE
);