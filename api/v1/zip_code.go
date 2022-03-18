package v1

import (
	"github.com/VerTox/zipcodes/api"
	"github.com/VerTox/zipcodes/domain/cases/zip_code_list"
	"github.com/VerTox/zipcodes/domain/model"
	"github.com/davecgh/go-spew/spew"
	"net/http"
)

func (a *ApiV1) GetZipCodeList(c *api.Context, w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	spew.Dump("GetZipCodeList entered")
	filter := &model.ZipCodeQuery{}

	q := c.Vars["query"]

	filter.Value = &q

	resp, err := zip_code_list.Run(
		&zip_code_list.Repositories{ZipCode: a.Context.Connection.ZipCode()},
		&zip_code_list.Request{
			Filter: filter,
			Limit:  10,
		})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, &ListResponse{Items: resp.ZipCodes, Total: resp.Total}, nil

}
