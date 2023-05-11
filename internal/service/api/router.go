package api

import (
	auth "github.com/acs-dl/auth-svc/middlewares"
	"github.com/acs-dl/role-svc/internal/data"
	handlers2 "github.com/acs-dl/role-svc/internal/service/api/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	secret := s.cfg.JwtParams().Secret

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers2.CtxLog(s.log),
			handlers2.CtxLinksParams(s.cfg.Links()),
		),
	)
	r.Route("/integrations/role-svc", func(r chi.Router) {
		r.Get("/user_roles", handlers2.GetUserRolesMap) // comes from orchestrator
		r.Get("/roles", handlers2.GetRolesMap)          // comes from orchestrator

		r.Route("/requests/users", func(r chi.Router) {
			r.With(auth.Jwt(secret, data.ModuleName, []string{"write"}...)).
				Post("/add", handlers2.AddUsers)
		})
	})

	return r
}
