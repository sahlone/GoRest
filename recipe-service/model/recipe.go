package model

import (
	"time"
)

type Difficulty int

const (
	Easy Difficulty = iota + 1
	Normal
	Hard
)

type Creds struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type IDToken struct {
	Token string `json:"token"`
}

type Preparation struct {
	Time  int `json:"time"`
	Total int `json:"total"`
}

type Recipe struct {
	ID         string      `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string      `json:"name" bson:"name,omitempty"`
	Prep       Preparation `json:"preparation" bson:"preparation,omitempty"`
	Difficulty Difficulty  `json:"difficulty" bson:"difficulty,omitempty" validate:"difficulty,min=1,max=3"`
	IsVeg      bool        `json:"vegetarian" bson:"vegetarian,omitempty"`
	Archive    bool        `json:"archive" bson:"archive,omitempty"`
	Modified   time.Time   `json:"modified" bson:"modified,omitempty"`
}

type RecipeRate struct {
	// ID
	ID       string    `json:"_id,omitempty" bson:"_id,omitempty"`
	RecipeID string    `json:"recipe_id,omitempty" bson:"recipe_id,omitempty"`
	Rate     int       `json:"rate,omitempty" bson:"rate,omitempty"`
	Modified time.Time `json:"modified,omitempty" bson:"modified,omitempty"`
}

/*
 * The Struct represents the fields that can be available for PATCH
 * Booleans/integers are replaced by strings as the defaults in boolean can't help in determining
 * the PATCH update perfectly
 */
type RecipePatch struct {
	Name       string      `json:"name" bson:"name,omitempty"`
	Prep       Preparation `json:"preparation" bson:"prep,omitempty"`
	Difficulty Difficulty  `json:"difficulty" bson:"difficulty,omitempty" validate:"difficulty,min=1,max=3"`
	IsVeg      string      `json:"vegetarian" bson:"vegetarian,omitempty"`
}
