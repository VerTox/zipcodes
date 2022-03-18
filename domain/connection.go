package domain

import "github.com/VerTox/zipcodes/domain/model"

type Connection interface {
	ZipCode() model.ZipCodeRepository
	IsErrNotFound(err error) bool
}
