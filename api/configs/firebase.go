package configs

import (
	"context"

	firebase "firebase.google.com/go/v4"
)

var app *firebase.App

func GetFirebaseApp() *firebase.App {
	if app == nil {
		var err error
		app, err = firebase.NewApp(context.Background(), nil)
		if err != nil {
			panic("failed to create firebase app")
		}
	}
	return app
}
