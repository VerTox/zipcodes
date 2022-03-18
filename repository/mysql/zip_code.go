package mysql

import (
	"errors"
	"github.com/VerTox/zipcodes/domain/model"
	"github.com/jinzhu/gorm"
)

type ZipCode struct {
	Id   int    `gorm:"primary_key"`
	Zip  string `gorm:"zip"`
	City string `gorm:"city"`
}

func (ZipCode) TableName() string {
	return "zip_codes"
}

func (zp *ZipCode) Model() *model.ZipCode {
	return &model.ZipCode{
		Zip:  zp.Zip,
		City: zp.City,
	}
}

type ZipCodeCast model.ZipCode

func (zp ZipCodeCast) Repository() *ZipCode {
	return &ZipCode{
		Zip:  zp.Zip,
		City: zp.City,
	}
}

type ZipCodeFilter struct {
	Value *string
}

func (f *ZipCodeFilter) filter(q *gorm.DB) *gorm.DB {
	if f.Value != nil {
		q = q.Debug().Where("zip LIKE ?", *f.Value+"%").Or("city LIKE ?", *f.Value+"%")
	}

	return q
}

func (f *ZipCodeFilter) WithValue(value *string) model.ZipCodeFilter {
	f.Value = value
	return f
}

type ZipCodeRepository struct {
	Connection *gorm.DB
}

func NewZipCodeRepository(c *gorm.DB) *ZipCodeRepository {
	return &ZipCodeRepository{
		Connection: c,
	}
}

func (r *ZipCodeRepository) GetList(f model.ZipCodeFilter, limit int) ([]*model.ZipCode, error) {
	filter, ok := f.(*ZipCodeFilter)

	if !ok {
		return nil, errors.New("invalid filter")
	}

	fq := filter.filter(r.Connection)

	if limit > 0 {
		fq = fq.Limit(limit)
	}

	var zcs []*ZipCode

	err := fq.Find(&zcs).Error

	if err != nil {
		return nil, err
	}

	ts := make([]*model.ZipCode, len(zcs))

	for i, tr := range zcs {
		ts[i] = tr.Model()
	}

	return ts, nil
}

func (r *ZipCodeRepository) NewFilter() model.ZipCodeFilter {
	return &ZipCodeFilter{}
}
