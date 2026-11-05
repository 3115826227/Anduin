package database

import (
 "fmt"

 "gorm.io/driver/mysql"
 "gorm.io/gorm"
)

func Open(dsn string) (*gorm.DB, error) {
 if dsn == "" {
  return nil, fmt.Errorf("database DSN is empty")
 }
 db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
 if err != nil {
  return nil, fmt.Errorf("connect mysql: %w", err)
 }
 return db, nil
}
