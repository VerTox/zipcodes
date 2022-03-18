package model

type ZipCode struct {
	Zip  string `json:"zip"`
	City string `json:"city"`
}

type ZipCodeRepository interface {
	GetList(f ZipCodeFilter, limit int) ([]*ZipCode, error)
	NewFilter() ZipCodeFilter
}

type ZipCodeQuery struct {
	Value *string `json:"value"`
}

type ZipCodeFilter interface {
	WithValue(value *string) ZipCodeFilter
}

func (q *ZipCodeQuery) Fill(f ZipCodeFilter) ZipCodeFilter {
	f.WithValue(q.Value)
	return f
}
