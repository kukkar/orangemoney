package config

import (
	"errors"
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/config"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
	"github.com/BoutiqaatREPO/getsetgo/src/components/mongodb"
)

type OMAppConfig struct {
	Mongo *mongodb.MDBConfig `json:"Mongo,omitempty"`
}

func GetGlobalConfig() *config.AppConfig {
	return config.GlobalAppConfig
}

func GetBQAppConfig() (*OMAppConfig, error) {
	c := config.GlobalAppConfig.ApplicationConfig

	appConfig, ok := c.(*OMAppConfig)

	if !ok {
		msg := fmt.Sprintf("Pims APP Config Not correct %+v", c)
		logger.Error(msg)
		return nil, errors.New(msg)
	}
	return appConfig, nil
}
func EnvUpdateMap() map[string]string {
	m := make(map[string]string)

	m["ApplicationConfig.Mongo.URL"] = "ATOM_MONGO_URL"
	m["ApplicationConfig.Mongo.DbName"] = "ATOM_MONGO_DB"
	m["ApplicationConfig.Cache.RedisCluster.Addrs"] = "ATOM_REDIS_CLUSTER"
	m["ApplicationConfig.Raven.RedisCluster.Addrs"] = "RAVEN_REDIS_CLUSTER"
	m["ApplicationConfig.Stock.MongoDbName"] = "STOCK_MONGO_DB"

	m["ApplicationConfig.OshrinkService.IPPort"] = "ATOM_OSHRINK_SERVICE"
	m["ApplicationConfig.OshrinkService.PubIPPort"] = "ATOM_OSHRINK_SERVICE_PUBLIC"
	m["ApplicationConfig.HydrazineService.IPPort"] = "ATOM_HYDRAZINE_SERVICE"
	m["ApplicationConfig.DiscoveryService.IPPort"] = "ATOM_DISCOVERY_SERVICE"
	m["ApplicationConfig.DiscoveryService.Sync"] = "ATOM_DISCOVERY_SYNC"

	m["MonitorConfig.NewRelicConfig.APIKey"] = "NEWRELIC_KEY"
	m["MonitorConfig.NewRelicConfig.AppId"] = "NEWRELIC_APP_NAME"

	m["ApplicationConfig.NemoService.IPPort"] = "ATOM_NEMO_SERVICE"

	m["ApplicationConfig.MagentoMig.IPPort"] = "ATOM_MAGENTOMIG_IPPORT"
	m["ApplicationConfig.MagentoMig.Creds.ConsumerKey"] = "ATOM_MAGENTOMIG_CREDS_CONSUMERKEY"
	m["ApplicationConfig.MagentoMig.Creds.ConsumerSecret"] = "ATOM_MAGENTOMIG_CREDS_CONSUMERSECRET"
	m["ApplicationConfig.MagentoMig.Creds.AccessToken"] = "ATOM_MAGENTOMIG_CREDS_ACCESSTOKEN"
	m["ApplicationConfig.MagentoMig.Creds.TokenSecret"] = "ATOM_MAGENTOMIG_CREDS_TOKENSECRET"
	m["ApplicationConfig.MagentoMig.Trigger"] = "ATOM_MAGENTOMIG_TRIGGER"
	m["AllowedHosts"] = "ATOM_ALLOWED_HOSTS"

	m["ApplicationConfig.ImageUrl"] = "ATOM_OSHRINK_IMAGE_URL"

	return m
}
