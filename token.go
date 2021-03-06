package acl

// Token ...
type Token interface {
	PermIsPrivileged() bool
	PermPolicies() []string
}
