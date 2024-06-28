package database

import (
	"context"
	"github.com/odrianoaliveira/go-service/models"
)

func (c *Client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor

	c.DB.WithContext(ctx).Find(&vendors)

	return vendors, nil
}
