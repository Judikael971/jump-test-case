package connectors

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	Connector *gorm.DB
)

func Set(dialector gorm.Dialector, opts ...gorm.Option) error {
	db, err := gorm.Open(dialector, opts...)
	Connector = db
	if err != nil {
		return fmt.Errorf("gorm.Open err: %w", err)
	}
	return nil
}
