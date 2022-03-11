package graph

import "github.com/jtyler258/graphql-blogpost/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	recipes []*model.Recipe
}
