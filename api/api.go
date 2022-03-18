package api

import (
	"github.com/VerTox/zipcodes/domain"
	"github.com/gorilla/mux"
)

type Api struct {
	Router  *mux.Router
	Context *domain.Context
}
