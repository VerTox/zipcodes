package zip_code_list

import (
	"errors"
	"github.com/VerTox/zipcodes/domain/model"
)

type Request struct {
	Filter *model.ZipCodeQuery
	Limit  int
}

type Response struct {
	ZipCodes []*model.ZipCode
	Total    int
}

type Repositories struct {
	ZipCode model.ZipCodeRepository
}

func Run(r *Repositories, request *Request) (*Response, error) {
	if r == nil || request == nil {
		return nil, errors.New("invalid case initialization")
	}

	f := request.Filter.Fill(r.ZipCode.NewFilter())

	zcl, err := r.ZipCode.GetList(f, request.Limit)

	if err != nil {
		return nil, err
	}

	return &Response{ZipCodes: zcl, Total: len(zcl)}, nil
}
