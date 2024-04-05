package middleware

import (
	"context"
	"log"
	"net/http"
	"ws/internal/config"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := config.Store.Get(r, "ws-session")
		if err != nil {
			log.Printf("session error: %v", err)
			http.Redirect(w, r, `/auth/login?error=token_not_valid`, http.StatusSeeOther)
			return
		}

		id := session.Values["userID"]
		if id == nil {
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}
		log.Println("user logged in:", id)

		ctx := context.WithValue(r.Context(), "userID", id)
		rWithCtx := r.WithContext(ctx)

		next.ServeHTTP(w, rWithCtx)
	})
}
