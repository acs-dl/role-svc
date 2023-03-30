package api

import (
	"github.com/go-chi/chi"
	auth "gitlab.com/distributed_lab/acs/auth/middlewares"
	"gitlab.com/distributed_lab/acs/role-svc/internal/data"
	handlers2 "gitlab.com/distributed_lab/acs/role-svc/internal/service/api/handlers"
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
