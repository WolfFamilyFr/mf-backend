package mocks

import (
	"github.com/bchaillou003/marvel-family-backend/application"
	dbMocks "github.com/bchaillou003/marvel-family-backend/database/mocks"
	sdkMocks "github.com/bchaillou003/marvel-family-backend/sdk/mocks"
)

func NewMockedApplication() (*application.App, *dbMocks.DatabaseMock, *sdkMocks.ClientSDKMock) {
	d := &dbMocks.DatabaseMock{}
	c := &sdkMocks.ClientSDKMock{}

	return &application.App{
		Database:  d,
		ClientSDK: c,
	}, d, c
}
