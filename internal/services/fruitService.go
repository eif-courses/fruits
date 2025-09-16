package services

import (
	"context"

	"github.com/eif-courses/fruits/internal/repository"
)

type FruitService struct {
	queries *repository.Queries
}

func NewFruitService(queries *repository.Queries) *FruitService {
	return &FruitService{queries: queries}
}

func (f *FruitService) GetFruits(ctx context.Context) ([]repository.Fruit, error) {
	return f.queries.ListFruits(ctx)
}
