package application

import (
	"context"
	"log"

	"github.com/bchaillou003/marvel-family-backend/database"
	"github.com/bchaillou003/marvel-family-backend/sdk"
)

type App struct {
	Database  database.IDB
	ClientSDK sdk.ISDK
}

var appKey *App

func NewApplication() *App {
	var app App

	app.Database = database.NewDatabase()
	app.ClientSDK = sdk.NewClientSDK()

	return &app
}

func (app *App) Close() error {
	if db, ok := app.Database.(database.Database); ok {
		return db.Close()
	}
	return nil
}

func ToContext(ctx context.Context, app *App) context.Context {
	return context.WithValue(ctx, appKey, app)
}

func FromContext(ctx context.Context) *App {
	appFromContext := ctx.Value(appKey)
	app, ok := appFromContext.(*App)
	if !ok {
		log.Fatal("App type assertion failed")
	}

	return app
}

// DatabaseFromContext returns Database
func DatabaseFromContext(ctx context.Context) database.IDB {
	return FromContext(ctx).Database
}

// ClientSDKFromContext returns ClientSDK
func ClientSDKFromContext(ctx context.Context) sdk.ISDK {
	return FromContext(ctx).ClientSDK
}
