package v1

import (
	"errors"
	"fmt"
	"github.com/VerTox/logger"
	"github.com/VerTox/zipcodes/api"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

func (a *ApiV1) RouteWrapper(callable func(c *api.Context, w http.ResponseWriter, r *http.Request) (int, interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		context := r.Context().Value("context")

		if context == nil {
			JsonError(w, http.StatusInternalServerError, errors.New("invalid request context"))
		} else {
			c := context.(*api.Context)

			c.Vars = mux.Vars(r)

			bodyFailed := false

			if r.Header.Get("Content-Type") == "application/json" {
				if r.Method == http.MethodPost || r.Method == http.MethodPatch || r.Method == http.MethodPut {
					body, err := ioutil.ReadAll(r.Body)

					if err != nil {
						bodyFailed = true
						JsonError(w, http.StatusBadRequest, errors.New("cannot read request"))
					} else {
						c.Body = body
					}
				}
			}

			c.With = strings.Split(r.URL.Query().Get("with"), ",")

			if !bodyFailed {
				status, response, err := callable(c, w, r)

				l := logger.New(c.TraceId)

				if err != nil {
					if a.Context.Connection.IsErrNotFound(err) && (r.Method == http.MethodGet || r.Method == http.MethodDelete || r.Method == http.MethodPatch) {
						status = http.StatusNotFound
					}

					l.Info(fmt.Sprintf("response [%d] %s", status, err.Error()), nil)

					JsonError(w, status, err)
				} else {
					if status == http.StatusCreated {
						JsonSuccessCreated(w, response)
					} else {
						JsonSuccess(w, response)
					}
				}
			}
		}
	}
}
