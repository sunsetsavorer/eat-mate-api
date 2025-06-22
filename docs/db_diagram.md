Table users {
  id bigint [pk]
  name string [not null]
  photo_url string [null]
}

Table places {
  id uuid [pk]
  name string [not null, unique]
  icon_path string [null]
}

Table place_branches {
  id uuid [pk]
  place_id uuid [ref: - places.id]
  address string [null]
  contact_phone string [null]
  coordinates json [default: '{}']
}

Table groups {
  id uuid [pk]
  name string [not null]
  owner_id bigint [ref: - users.id]
  is_public bool [default: true]
  selection_mode selection_mode [not null]
  place_branch_id uuid [ref: - place_branches.id]
}

Table group_members {
  group_id uuid [ref: - groups.id]
  user_id bigint [ref: - users.id]
}

Table votes {
  group_id uuid [ref: - groups.id]
  user_id bigint [ref: - users.id]
  place_branch_id uuid [ref: - place_branches.id]
}

Table group_place_options {
  group_id uuid [ref: - groups.id]
  place_branch_id uuid [ref: - place_branches.id]
}

Enum selection_mode {
  defined
  voting
  random
}