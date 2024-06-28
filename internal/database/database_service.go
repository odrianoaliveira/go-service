package database

import (
	"context"
	"github.com/odrianoaliveira/go-service/models"
)

func (c *Client) GetAllServices(ctx context.Context) ([]models.Service, error) {
	var services []models.Service

	c.DB.WithContext(ctx).Find(&services)

	return services, nil
}
