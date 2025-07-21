CREATE TYPE group_mode AS ENUM ('defined', 'voting', 'random');

CREATE TABLE groups (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(50) NOT NULL,
	owner_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
	is_public BOOLEAN NOT NULL DEFAULT true,
	selection_mode group_mode NOT NULL,
	place_branch_id UUID NULL REFERENCES place_branches(id) ON DELETE CASCADE ON UPDATE CASCADE
);