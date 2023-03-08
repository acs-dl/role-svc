package service

import (
	"github.com/go-chi/chi"
	auth "gitlab.com/distributed_lab/acs/auth/middlewares"
	"gitlab.com/distributed_lab/acs/role-svc/internal/data"
	"gitlab.com/distributed_lab/acs/role-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	secret := s.cfg.JwtParams().Secret

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxLinksParams(s.cfg.Links()),
		),
	)
	r.Route("/integrations/role-svc", func(r chi.Router) {
		r.Route("/requests/users", func(r chi.Router) {
			r.With(auth.Jwt(secret, data.ModuleName, []string{"write"}...)).
				Post("/add", handlers.AddUsers)
		})
	})

	return r
}
