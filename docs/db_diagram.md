Table users {
  id bigint [pk]
  name string [not null]
  photo_url string [null]
}

Table brands {
  id uuid [pk]
  name string [not null, unique]
  icon_path string [null]
}

Table branches {
  id uuid [pk]
  brand_id uuid [ref: - brands.id]
  address string [null]
  contact_phone string [null]
  coordinates json [null]
}

Table groups {
  id uuid [pk]
  name string [not null]
  is_public bool [default: true]
  is_active bool [default: true]
  selection_mode selection_mode [not null]
  branch_id uuid [null, ref: - branches.id]
}

Table group_members {
  group_id uuid [ref: - groups.id]
  user_id bigint [ref: - users.id]
  role group_member_role [not null]
}

Table votes {
  group_id uuid [ref: - groups.id]
  user_id bigint [ref: - users.id]
  branch_id uuid [null, ref: - branches.id]
}

Table group_branch_options {
  group_id uuid [ref: - groups.id]
  branch_id uuid [ref: - branches.id]
}

Enum selection_mode {
  defined
  voting
  random
}

Enum group_member_role {
  owner
  participant
}