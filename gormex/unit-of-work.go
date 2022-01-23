package gormex

import "gorm.io/gorm"

type unitOfWork struct {
	tx *gorm.DB
}

func (u unitOfWork) Commit() error {
	u.tx.Commit()
	return u.tx.Error
}
