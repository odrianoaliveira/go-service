package database

import (
	"context"
	"fmt"
	"github.com/odrianoaliveira/go-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type DatabaseClient interface {
	Ready() bool
	GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error)
}

type Client struct {
	DB *gorm.DB
}

func NewClient() (DatabaseClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		"localhost", "postgres", "postgres", "postgres", "5432")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom.",
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})

	if err != nil {
		return nil, err
	}

	client := &Client{
		DB: db,
	}
	return client, nil
}

func (c *Client) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 'ready'").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready != "ready" {
		return false
	}
	return true
}
