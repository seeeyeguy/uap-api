package main

import (
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
)

func (s *Server) requireAuth(next http.Handler) http.Handler {
	clerk.SetKey(s.cfg.ClerkSecretKey)
	return clerkhttp.RequireHeaderAuthorization()(next)
}

func (s *Server) getUserID(r *http.Request) string {
	claims, ok := clerk.SessionClaimsFromContext(r.Context())
	if !ok {
		return ""
	}
	return claims.Subject
}
