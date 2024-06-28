package database

import (
	"context"
	"github.com/odrianoaliveira/go-service/models"
)

func (c *Client) GetAllProducts(ctx context.Context, vendorID string) ([]models.Product, error) {
	var products []models.Product

	c.DB.WithContext(ctx).
		Where(models.Product{VendorID: vendorID}).
		Find(&products)

	return products, nil
}
