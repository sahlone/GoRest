package core

import (
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
)

/*
 * Service interface should not accept any pointers so data can be non-mutable
 * and help in making the pure functions where side effect will be addressed by db layer
 */
type RecipeService interface {
	GetRecipe(id ID) (model.Recipe, error)
	UpdateRecipe(recipe model.Recipe) error
	PatchRecipe(id ID, recipe model.RecipePatch) error
	DeleteRecipe(id ID) error
	CreateRecipe(recipe model.Recipe) error
	GetRecipes(lastRecord string, limit int) ([]model.Recipe, error)
	SearchRecipes(start, limit int, search model.SearchCriteria) ([]model.Recipe, error)
	RateRecipe(id ID, rating int) error
}
