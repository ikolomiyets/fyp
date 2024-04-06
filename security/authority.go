package security

type AuthorityKey struct {
}

type Authority struct {
	UserID string
	Scopes []string
	Roles  []string
}
