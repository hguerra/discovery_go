package user

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/domain/user"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/domain/user/usecase"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/logging"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/web/req"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/web/res"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/web/validate"
)

func RegisterUserRoutes() http.Handler {
	logger := logging.GetLogger()
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		users, err := usecase.FindUsers()
		if err != nil {
			res.BadRequest(w, "abc", "error to find users")
			return
		}
		res.OK(w, users)
	})

	r.Get("/pages", func(w http.ResponseWriter, r *http.Request) {
		page := req.NewPage(r)
		sort := req.NewSort(r)
		logger.Infof("Request page %d (%d)", page.Page, page.Size)
		users, err := usecase.FindUsers()
		if err != nil {
			res.BadRequest(w, "abc", "error to find users")
			return
		}

		newPage := 1
		newSize := 10
		newTotal := 100
		res.PageOf(w, res.M{"pageRequest": page, "sortRequest": sort, "randomData": users}, newPage, newSize, newTotal)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var u user.User
		if err := req.BindJSON(w, r, &u); err != nil {
			res.BadRequest(w, "cde", err.Error())
			return
		}
		errs := validate.Validate(u)
		logger.Infof("Request user with name %s, errors %v", u.FirstName, errs)
		if len(errs) > 0 {
			res.UnprocessableEntity(w, "123", "Invalid body", errs...)
			return
		}
		res.Created(w, u)
	})

	r.Post("/fields", func(w http.ResponseWriter, r *http.Request) {
		o, err := req.BindJSONObject(w, r)
		if err != nil {
			res.InternalServerError(w, "efg", err.Error())
			return
		}
		name := o.Get("firstName").String()
		logger.Infof("Request user with name %s", name)
		res.Created(w, res.M{"name": name})
	})

	return r
}
