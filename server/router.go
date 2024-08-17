package server

import (
	"github.com/bchaillou003/marvel-family-backend/server/comics"
	"github.com/bchaillou003/marvel-family-backend/server/rivals"
	"github.com/gin-gonic/gin"
)

func (api *API) InitRouter() {
	api.Router = gin.Default()
	api.Router.Use(api.ContextMiddleware)

	comics.SetupRoutes(api.Router)
	rivals.SetupRoutes(api.Router)
}
