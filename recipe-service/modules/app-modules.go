package modules

/*
 * This will be used for wiring of the project modules based on interface
 *  implementations we want to provide
 * If there is any change in implementation we will just need to update here the new implementation
 */
import (
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/config"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/core"
	db "github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/dbhandler"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/logger"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/service"
)

/*
 * Returns the interface implementation for Mongo handler
 */
func NewMongoHandler(config *config.Config) *db.MongoHandler {
	logger.Info("New mongo store created")
	return &db.MongoHandler{
		Session: db.StartConnection(&config.DBConfig),
		Config:  config,
	}
}

/*
 * Returns the interface implementation for RecipeRepo
 */
func NewRecipeDBRepo(handler *db.MongoHandler) core.RecipeDBRepo {
	logger.Info("New mongo repo created")
	return &db.MongoDBRepo{MongoHandler: handler}
}

/*
 * Returns the interface implementation for RecipeService
 */
func NewRecipeService(repo core.RecipeDBRepo) core.RecipeService {
	logger.Info("Create RecipeServiceImpl created ")
	return &service.RecSerImpl{repo}
}
