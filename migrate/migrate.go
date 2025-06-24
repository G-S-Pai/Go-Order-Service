package main

import (
	"fmt"
	"log"

	"github.com/g-s-pai/go-order-service/initializers"
	"github.com/g-s-pai/go-order-service/models"

	spannergorm "github.com/googleapis/go-gorm-spanner"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	tables := []interface{}{&models.Order{}}

	m := initializers.DB.Migrator()
	migrator := m.(spannergorm.SpannerMigrator)
	err := migrator.AutoMigrate(tables...)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("👍 Migration complete")
}
