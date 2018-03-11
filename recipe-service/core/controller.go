package core

import (
	"net/http"
)

type RequestContext func(http.ResponseWriter, *http.Request)

type Controller interface {
	Apply(service RecipeService) RequestContext
}
