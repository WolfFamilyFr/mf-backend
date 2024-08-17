package server

import (
	"log"

	"github.com/bchaillou003/marvel-family-backend/application"
	"github.com/gin-gonic/gin"
)

type API struct {
	Router *gin.Engine
	App    *application.App
}

func NewAPI() *API {
	var api API

	// Initialize app
	api.App = application.NewApplication()

	api.InitRouter()

	return &api
}

func (api *API) Run() {
	// Listen serves HTTP requests from the given addr.
	log.Fatal(api.Router.Run(":3000"))
}

func (api *API) ContextMiddleware(c *gin.Context) {
	ctx := application.ToContext(c.Request.Context(), api.App)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
