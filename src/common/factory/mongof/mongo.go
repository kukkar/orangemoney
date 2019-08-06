package mongof

import (
	"errors"
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/collections/maps/concurrentmap/concurrenthashmap"
	"github.com/BoutiqaatREPO/getsetgo/src/components/mongodb"
	appconfig "github.com/OrangeMoney/pims/src/config"
)

var ErrNotFound = errors.New("not found")

var mongoMap = concurrenthashmap.New()
var stockMongoMap = concurrenthashmap.New()

const DEFAULT_KEY = "default"

type MDB struct {
	mongodb.MDBInterface
}

func GetPool(key string) (*MDB, error) {
	if val, ok := mongoMap.Get(key); !ok {
		adapter, err := InitPool(key)
		if err != nil {
			return nil, fmt.Errorf("Could not get Mongo Adapter for key:%s", key)
		}
		mongoMap.Put(key, adapter)
		return adapter, nil
	} else {
		return val.(*MDB), nil
	}
}

func InitPool(key string) (*MDB, error) {
	config, _ := appconfig.GetBQAppConfig()
	mongoDriver := new(mongodb.MongoDriver)

	merr := mongoDriver.Init(config.Mongo)
	if merr != nil {
		return nil, merr
	}
	var i mongodb.MDBInterface = mongoDriver
	n := MDB{i}

	return &n, nil
}
