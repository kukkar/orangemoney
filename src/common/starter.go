package common

import (
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/core/service"
	"github.com/OrangeMoney/pims/src/common/appconstant"
	appconfig "github.com/OrangeMoney/pims/src/config"
	"github.com/OrangeMoney/pims/src/services/hello"
	"github.com/OrangeMoney/pims/src/services/pims/put/globaluser"
)

//main is the entry point of the florest web service

func StartServer() {
	fmt.Println("APPLICATION BEGIN")
	webserver := new(service.Webserver)
	Register()
	webserver.PreStart(func() {}, func() {})
	webserver.Start()
}

func Register() {
	registerConfig()
	registerErrors()
	registerAllApis()
}

func registerAllApis() {
	service.RegisterAPI(new(hello.HelloAPI))
	service.RegisterAPI(new(globaluser.CreateStockAPI))
}

func registerConfig() {
	service.RegisterConfig(new(appconfig.OMAppConfig))
}

func registerErrors() {
	service.RegisterHTTPErrors(appconstant.APPErrorCodeToHTTPCodeMap)
}
