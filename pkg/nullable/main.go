package nullable

import "database/sql"

func NullStringToPtr(ns sql.NullString) *string {

	if ns.Valid {
		return &ns.String
	}

	return nil
}

func PtrToNullString(s *string) sql.NullString {

	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}

	return sql.NullString{Valid: false}
}
