package db

import (
    "fmt"
    "os"
    "time"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

func NewGormDBFromEnv() (*gorm.DB, error) {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASSWORD")
    name := os.Getenv("DB_NAME")
    ssl := os.Getenv("DB_SSLMODE")
    if ssl == "" {
        ssl = "disable"
    }

    dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, name, pass, ssl)
    cfg := &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    }
    db, err := gorm.Open(postgres.Open(dsn), cfg)
    if err != nil {
        return nil, err
    }
    // set conn pool
    sqlDB, err := db.DB()
    if err == nil {
        sqlDB.SetMaxOpenConns(25)
        sqlDB.SetMaxIdleConns(5)
        sqlDB.SetConnMaxLifetime(time.Hour)
    }
    return db, nil
}
