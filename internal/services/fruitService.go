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

func (f *FruitService) InsertFruit(ctx context.Context, name string, colour string) (*repository.Fruit, error) {

	params := repository.InsertFruitParams{
		Name:   name,
		Colour: colour,
	}

	fruit, err := f.queries.InsertFruit(ctx, params)

	if err != nil {
		return nil, err
	}
	return &fruit, nil
}

func (f *FruitService) DeleteFruit(ctx context.Context, id int) error {

	err := f.queries.DeleteFruit(ctx, id)

	if err != nil {
		return err
	}
	return nil
}
