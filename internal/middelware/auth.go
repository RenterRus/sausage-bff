package middelware

import (
	"context"
	"net/http"

	"github.com/RenterRus/sausage-bff/internal/usecase/auth"
)

// !!!
func AuthMiddelware(ctx context.Context, next http.Handler, auth auth.AuthCases) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		access, err := r.Cookie("access_token")
		if err != nil && err != http.ErrNoCookie {
			http.Error(w, "Ошибка сервера", http.StatusBadRequest)
			return
		}

		refresh, err := r.Cookie("refresh_token")
		if err != nil && err != http.ErrNoCookie {
			http.Error(w, "Ошибка сервера", http.StatusBadRequest)
			return
		}

		uuid, err := auth.ValidateToken(ctx, &access.Value)

		_ = refresh
		_ = uuid
		r.AddCookie(&http.Cookie{})

		next.ServeHTTP(w, r)
	})
}
