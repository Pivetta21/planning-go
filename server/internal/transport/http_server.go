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
	profiledel "github.com/Pivetta21/planning-go/internal/feature/profile/delete"
	profilefind "github.com/Pivetta21/planning-go/internal/feature/profile/find"
	profileupdate "github.com/Pivetta21/planning-go/internal/feature/profile/update"
	roomwsconn "github.com/Pivetta21/planning-go/internal/feature/room/conn"
	sessiondel "github.com/Pivetta21/planning-go/internal/feature/session/delete"
	sessionlist "github.com/Pivetta21/planning-go/internal/feature/session/list"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func StartHttpServer() {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.AllowContentType("application/json"))

	// CORS
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	// Routes
	mux.Route("/auth", authRoutes())
	mux.Route("/session", sessionRoutes())
	mux.Route("/profile", profileRoutes())
	mux.Route("/room", roomRoutes())

	log.Println("http server running on port:", configs.APIConfig.Port)

	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", configs.APIConfig.Port), mux)
	log.Fatalln(err)
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
		r.Delete("/", profiledel.HandleDeleteProfile)
	}
}

func roomRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/conn", roomwsconn.HandleRoomConnection)
	}
}
