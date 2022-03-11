package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/jtyler258/graphql-blogpost/graph/generated"
	"github.com/jtyler258/graphql-blogpost/graph/model"
)

func (r *mutationResolver) CreateRecipe(ctx context.Context, input model.CreateRecipeInput) (*model.Recipe, error) {
	recipe := &model.Recipe {
		ID: uuid.New().String(),
		Title: input.Title,
		Author: input.Author,
		Ingredients: input.Ingredients,
		Steps: input.Steps,
	}

	r.recipes = append(r.recipes, recipe)
	return recipe, nil
}

func (r *queryResolver) Recipes(ctx context.Context) ([]*model.Recipe, error) {
	return r.recipes, nil
}

func (r *queryResolver) Recipe(ctx context.Context, id string) (*model.Recipe, error) {
	for _, v := range r.recipes {
		if v.ID == id {
			return v, nil
		}
	}

	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
