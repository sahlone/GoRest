package dbhandler

import (
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/logger"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MongoDBRepo struct {
	*MongoHandler
}

/*
 * Get  recipe
 */
func (accessor *MongoDBRepo) Get(id core.ID) (model.Recipe, error) {
	logger.Info("Get Recipe: ID : %s", id)
	collection, session := accessor.RecipeAccess()
	defer session.Close()
	recipe := model.Recipe{}
	err := collection.Find(bson.M{core.Identity(): id}).One(&recipe)
	return recipe, err
}

/*
 * Create create recipe
 */
func (accessor *MongoDBRepo) Create(recipe model.Recipe) error {
	logger.Info("Create Recipe with data : %+v", recipe)
	collection, session := accessor.RecipeAccess()
	defer session.Close()
	return collection.Insert(recipe)
}

/*
 * Update recipe
 */
func (accessor *MongoDBRepo) Update(recipe model.Recipe) error {
	logger.Info("Update Recipe  with ID : %+v", recipe)
	collection, session := accessor.RecipeAccess()
	defer session.Close()
	query := bson.M{core.Identity(): recipe.ID}
	change := bson.M{"$set": bson.M{core.Name(): recipe.Name, core.Prep(): recipe.Prep,
		core.Difficultylevel(): recipe.Difficulty, core.IsVeg(): recipe.IsVeg,
		core.IsArchive(): recipe.Archive, core.Modified(): time.Now()}}
	return collection.Update(query, change)
}

/*
 * Delete recipe
 */
func (accessor *MongoDBRepo) Delete(id core.ID) error {
	logger.Info("Delete Recipe with ID : %s", id)
	collection, session := accessor.RecipeAccess()
	defer session.Close()
	query := bson.M{core.Identity(): id}
	change := bson.M{"$set": bson.M{core.IsArchive(): true, core.Modified(): time.Now()}}
	return collection.Update(query, change)
}

/*
 * Rate recipe
 */
func (accessor *MongoDBRepo) Rate(id core.ID, rate int) error {
	logger.Info("Rate Recipe with ID : %v and rate : %v", id, rate)
	collection, session := accessor.RecipeRateAccess()
	defer session.Close()
	return collection.Insert(&model.RecipeRate{RecipeID: string(id), Rate: rate, Modified: time.Now()})
}

/*
 * List recipe
 * This service will return paginated results. Max item limit is 100
 *  which will be controlled by service
 * Listing is done based on the last record retrieved.
 * This approach will make effective use of default index on _id and nature of ObjectId.
 */
func (accessor *MongoDBRepo) List(lastRecord string, limit int) ([]model.Recipe, error) {
	var recipes []model.Recipe
	collection, session := accessor.RecipeAccess()
	defer session.Close()
	var query bson.M = nil
	var err error = nil
	if len(lastRecord) > 0 {
		query = bson.M{core.Identity(): bson.M{"$gt": lastRecord}}
		err = collection.Find(query).Sort(core.Identity()).Limit(limit).All(&recipes)
	} else {
		err = collection.Find(nil).Sort(core.Identity()).Limit(limit).All(&recipes)
	}

	return recipes, err
}

/*
 * List recipe
 * This service will return paginated results. Max item limit is 100
 *  which will be controlled by service
 */
func (accessor *MongoDBRepo) Search(start, limit int, search model.SearchCriteria) ([]model.Recipe, error) {
	var recipes []model.Recipe
	collection, session := accessor.RecipeAccess()
	defer session.Close()
	query := bson.M{}

	if search.SearchByArchive != "" {
		query[core.IsArchive()] = true
	}
	if search.SearchByDifficulty != 0 {
		query[core.Difficultylevel()] = search.SearchByDifficulty
	}
	if search.SearchByID != "" {
		query[core.Identity()] = search.SearchByID
	}
	if search.SearchByIsVeg != "" {
		query[core.IsVeg()] = search.SearchByIsVeg
	}
	if search.SearchByName != "" {
		query[core.Name()] = search.SearchByName
	}
	if &search.SearchByPrep != nil {
		isValid := false
		innerQuery := bson.M{}
		if search.SearchByPrep.Time != 0 {
			innerQuery[core.PrepTime()] = search.SearchByPrep.Time
			isValid = true
		}
		if search.SearchByPrep.Total != 0 {
			innerQuery[core.PrepTotal()] = search.SearchByPrep.Total
			isValid = true
		}
		if isValid {
			query[core.Prep()] = innerQuery
		}
	}
	err := collection.Find(query).Sort(core.Identity()).Skip(start).Limit(limit).All(&recipes)
	return recipes, err
}
