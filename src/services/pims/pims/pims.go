package pims

import (
	"github.com/OrangeMoney/pims/src/common/appconstant"
	"github.com/OrangeMoney/pims/src/common/factory/mongof"
)

//
//medium to talk to stock using stock intf outer world can access all its exposed functions
//

func GetPimsIntf(key string) (PimsInterface, error) {

	var mongoKey string
	var pimsServiceImpl pimsService
	if key != "" {
		mongoKey = key
	} else {
		mongoKey = appconstant.DEFAULT_KEY
	}

	db, errG := mongof.GetPool(mongoKey)
	if errG != nil {
		return nil, errG
	}

	pimsServiceImpl.MongoSession = db

	return &pimsServiceImpl, nil
}
