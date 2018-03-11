package service

import (
	. "github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/logger"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
)

type RecSerImpl struct {
	RecipeDBRepo
}

func (ser *RecSerImpl) GetRecipe(id ID) (model.Recipe, error) {
	recipe, err := ser.Get(id)
	if err != nil && err.Error() == "not found" {
		err = model.ErrNotFound
	}
	return recipe, err
}

func (ser RecSerImpl) CreateRecipe(recipe model.Recipe) error {
	return ser.Create(recipe)
}

func (ser RecSerImpl) RateRecipe(id ID, rating int) error {
	_, err := ser.GetRecipe(id)
	if (err != nil) {
		return err
	}
	return ser.Rate(id, rating)
}

/*
 * This is used for PUT Operation on recipe wher we will just replace the
 * old object with new data.
 * The rest client is supposed to provide data for all field , otherwise will be overwritten by defaults
 */
func (ser RecSerImpl) UpdateRecipe(recipe model.Recipe) error {
	err := ser.Update(recipe)
	if err != nil && err.Error() == "not found" {
		err = model.ErrNotFound
	}
	return err
}

/*
 * This is used for PATCH Operation on recipe where we will just replace the
 * provided fields with new data
 * The rest client is supposed to provide data for fields that need to be updated
 */
func (ser RecSerImpl) PatchRecipe(id ID, patch model.RecipePatch) error {
	recipe, err := ser.GetRecipe(id)
	if err != nil {
		logger.Info("GET recipe for patch with ID : %v Err :%v", id, err)
		return err
	}
	if patch.Name != "" {
		recipe.Name = patch.Name
	}
	if &patch.Prep != nil {
		// error check for these should be job of controller
		//time,_:= strconv.ParseInt(patch.Prep.Time,10,64)
		//total,_:= strconv.ParseInt(patch.Prep.Total,10,64)
		if patch.Prep.Total != 0 {
			recipe.Prep.Total = patch.Prep.Total
		}
		if patch.Prep.Time != 0 {
			recipe.Prep.Time = patch.Prep.Time
		}

		recipe.Prep = model.Preparation{patch.Prep.Time, patch.Prep.Total}
	}
	if patch.IsVeg != "" {
		if !(patch.IsVeg == "true" || patch.IsVeg == "false") {
			return model.ErrPatchIsVeg
		}
		if patch.IsVeg == "true" {
			recipe.IsVeg = true
		} else {
			recipe.IsVeg = false
		}

	}
	if patch.Difficulty != 0 {
		recipe.Difficulty = patch.Difficulty
	}
	return ser.Update(recipe)
}

/*
 * Implements soft delete. No actual record will be deleted
 */
func (ser RecSerImpl) DeleteRecipe(id ID) error {
	err := ser.Delete(id)
	if err != nil && err.Error() == "not found" {
		err = model.ErrNotFound
	}
	return err
}

/*
 * Implements pagination based GET for recipes
 */
func (ser RecSerImpl) GetRecipes(lastRecord string, limit int) ([]model.Recipe, error) {
	if limit > 100 || limit < 1 {
		return nil, model.ErrorGetLimit
	}
	return ser.List(lastRecord, limit)
}

/*
 * Implements searching based on Criteria passed
 */
func (ser RecSerImpl) SearchRecipes(start, limit int, search model.SearchCriteria) ([]model.Recipe, error) {
	if limit == 0 {
		limit = 10
	}
	if limit > 100 || limit < 1 {
		return nil, model.ErrorGetLimit
	}
	if start < 0 {
		return nil, model.ErrorGetSkip
	}

	return ser.Search(start, limit, search)
}
