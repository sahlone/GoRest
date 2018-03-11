package core

import (
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
)

type ID string

/*
 * Database interface should not accept any pointers so data can be non-mutable
 * and help in making the concurrent access error free s its side effect(DB) producing layer
 */
type RecipeDBRepo interface {
	List(lastRecord string, limit int) ([]model.Recipe, error)
	Create(recipe model.Recipe) error
	Get(id ID) (model.Recipe, error)
	Update(recipe model.Recipe) error
	Delete(id ID) error
	Rate(id ID, rate int) error
	Search(start, limit int, search model.SearchCriteria) ([]model.Recipe, error)
}
type Column string
type Table struct {
	Columns []Column
}

var RECIPE_TABLE = Table{Columns: []Column{
	"_id",
	"name",
	"preparation",
	"difficulty",
	"vegetarian",
	"archive",
	"time",
	"total",
	"modified",
}}

func Identity() string {
	return string(RECIPE_TABLE.Columns[0])
}
func Name() string {
	return string(RECIPE_TABLE.Columns[1])
}
func Prep() string {
	return string(RECIPE_TABLE.Columns[2])
}
func Difficultylevel() string {
	return string(RECIPE_TABLE.Columns[3])
}
func IsVeg() string {
	return string(RECIPE_TABLE.Columns[4])
}
func IsArchive() string {
	return string(RECIPE_TABLE.Columns[5])
}
func PrepTime() string {
	return string(RECIPE_TABLE.Columns[6])
}
func PrepTotal() string {
	return string(RECIPE_TABLE.Columns[7])
}
func Modified() string {
	return string(RECIPE_TABLE.Columns[8])
}
