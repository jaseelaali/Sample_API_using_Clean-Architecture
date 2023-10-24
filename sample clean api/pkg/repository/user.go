package repository

import (
	"context"

	interfaces "github.com/jaseelaali/Sample_API_using_Clean-Architecture/pkg/repository/interface"
	"gorm.io/gorm"
)

type db struct {
	db *gorm.DB
}

func NewRepository(dB *gorm.DB) interfaces.UserRepository {
	return &db{dB}
}
func (d *db) Wish(c context.Context) (string, error) {
	var name string
	err := d.db.Raw("select * from users").Scan(&name)
	if err != nil {
		return name, err.Error
	}
	return name, nil

}
