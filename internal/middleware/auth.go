package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/user"
)

// AuthUser represents the authenticated user information
type AuthUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// AuthContextKey is the key used to store the authenticated user in the request context
type AuthContextKey string

const AuthUserContextKey AuthContextKey = "auth_user"

// AuthMiddleware handles Clerk authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip auth for health check endpoint
		if r.URL.Path == "/health" {
			next.ServeHTTP(w, r)
			return
		}

		clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))
		claims, ok := clerk.SessionClaimsFromContext(r.Context())
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"access": "unauthorized"}`))
			return
		}

		user, err := user.Get(r.Context(), claims.Subject)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Create auth user object
		authUser := AuthUser{
			ID:        user.ID,
			Email:     user.EmailAddresses[0].EmailAddress,
			FirstName: *user.FirstName,
			LastName:  *user.LastName,
		}

		// Add user to context
		ctx := context.WithValue(r.Context(), AuthUserContextKey, authUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetAuthUser retrieves the authenticated user from the request context
func GetAuthUser(r *http.Request) (*AuthUser, error) {
	user, ok := r.Context().Value(AuthUserContextKey).(AuthUser)
	if !ok {
		return nil, nil
	}
	return &user, nil
}

// RequireAuth is a helper function to check if a user is authenticated
func RequireAuth(w http.ResponseWriter, r *http.Request) (*AuthUser, bool) {
	user, err := GetAuthUser(r)
	if err != nil || user == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Unauthorized - Authentication required",
		})
		return nil, false
	}
	return user, true
}
