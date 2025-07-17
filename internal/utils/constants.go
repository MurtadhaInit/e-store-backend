package utils

type scopes struct {
	ScopeAuth string
}

const (
	ScopeAuth = "authentication"
)

var Scopes = scopes{
	ScopeAuth: ScopeAuth,
}
