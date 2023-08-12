package transport

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pivetta21/planning-go/internal/configs"
	login "github.com/Pivetta21/planning-go/internal/feature/auth/login"
	logout "github.com/Pivetta21/planning-go/internal/feature/auth/logout"
	refresh "github.com/Pivetta21/planning-go/internal/feature/auth/refresh"
	register "github.com/Pivetta21/planning-go/internal/feature/auth/register"
	profilefind "github.com/Pivetta21/planning-go/internal/feature/profile/find"
	profileupdate "github.com/Pivetta21/planning-go/internal/feature/profile/update"
	sessiondel "github.com/Pivetta21/planning-go/internal/feature/session/delete"
	sessionlist "github.com/Pivetta21/planning-go/internal/feature/session/list"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartHttpServer() error {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.AllowContentType("application/json"))

	// Routes
	mux.Route("/auth", authRoutes())
	mux.Route("/session", sessionRoutes())
	mux.Route("/profile", profileRoutes())

	log.Println("server running on port:", configs.APIConfig.Port)
	return http.ListenAndServe(fmt.Sprintf("localhost:%d", configs.APIConfig.Port), mux)
}

func authRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/sign-up", register.HandleRegister)
		r.Post("/sign-in", login.HandleLogin)
		r.Post("/refresh", refresh.HandleRefresh)
		r.With(AuthMiddleware).Delete("/logout", logout.HandleLogout)
	}
}

func sessionRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Get("/", sessionlist.HandleList)
		r.Delete("/{identifier}", sessiondel.HandleDelete)
	}
}

func profileRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Get("/", profilefind.HandleFindProfile)
		r.Patch("/", profileupdate.HandleUpdateProfile)
	}
}
