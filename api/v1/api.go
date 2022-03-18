package v1

import (
	"github.com/VerTox/zipcodes/api"
	"github.com/VerTox/zipcodes/domain"
	"github.com/gorilla/mux"
)

const RoutePrefix = "/api/v1"

type ApiV1 struct {
	*api.Api
}

func NewV1(r *mux.Router, c *domain.Context) {
	a := &ApiV1{
		Api: &api.Api{
			Router:  r.PathPrefix(RoutePrefix).Subrouter(),
			Context: c,
		},
	}

	a.bootRoutes()
}

func (a *ApiV1) bootRoutes() {
	r := a.Router

	r.HandleFunc("/zip_find/{query}", a.RouteWrapper(a.GetZipCodeList)).Methods("GET")
}
