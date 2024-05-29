package model

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"math/big"
	"time"
)

type BaseData struct {
	// ID
	ID        int64          `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"softDelete:flag" json:"deleted_at"`
}

// BigFloat wraps big.Float to implement sql.Scanner
type BigFloat struct {
	*big.Float
}

// Scan implements the sql.Scanner interface for BigFloat
func (bf *BigFloat) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		f, _, err := big.ParseFloat(string(v), 10, 256, big.ToNearestEven)
		if err != nil {
			return err
		}
		bf.Float = f
	case string:
		f, _, err := big.ParseFloat(v, 10, 256, big.ToNearestEven)
		if err != nil {
			return err
		}
		bf.Float = f
	case float64:
		bf.Float = big.NewFloat(v)
	case nil:
		bf.Float = nil
	default:
		return fmt.Errorf("cannot scan type %T into BigFloat", value)
	}
	return nil
}
func (bf *BigFloat) Value() (driver.Value, error) {
	if bf.Float == nil {
		return nil, nil
	}
	return bf.Text('f', -1), nil
}
