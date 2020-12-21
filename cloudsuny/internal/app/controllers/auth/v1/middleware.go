package auth

import (
	"io"
	"net/http"
)

func (i *Implementation) Middleware(next http.Handler) http.Handler {
	withAuth := map[string]struct{}{
		"/api/v1/badge/create":  {},
		"/api/v1/badge/update":  {},
		"/api/v1/user/create":   {},
		"/api/v1/user/delete":   {},
		"/api/v1/image/upload":  {},
		"/api/v1/auth/sign_out": {},
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем сессию на всех запросах, если сессия протухла - сразу удаляем куку, т.к. в базе ее больше нет.
		session, err := i.service.Authenticate(r.Context(), i.readHTTPCookie(r))
		if err != nil {
			http.SetCookie(w, i.cookie("", true))
		}

		// Если пользователь авторизован, устанавливаем данный хедер для UI, что бы там можно было это понять.
		if session != nil && err == nil {
			w.Header().Set("X-Admin", "true")
		}

		// Некоторые запросы (/api/v1/user/list/with_badges, etc...) не нуждаются
		// в авторизации, по этому данный этап не обязателен.
		if _, ok := withAuth[r.URL.Path]; ok && err != nil {
			i.logger.
				With("client-ip", r.RemoteAddr).
				With("client-user-agent", r.UserAgent()).
				With("req-url", r.URL.String()).
				Warnf(r.Context(), "can't authenticate request: %v", err)

			w.WriteHeader(http.StatusForbidden)
			_, _ = io.WriteString(w, http.StatusText(http.StatusForbidden))

			return
		}

		next.ServeHTTP(w, r)
	})
}
