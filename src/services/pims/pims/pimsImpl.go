package pims

import (
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
	"github.com/OrangeMoney/pims/src/common/factory/mongof"
	"gopkg.in/mgo.v2/bson"
)

type pimsService struct {
	//mongo client
	MongoSession *mongof.MDB
}

func (this pimsService) UpdateUser(re RequestUser) error {

	query := map[string]interface{}{}
	query["pimsid"] = re.PimsId
	//Update Document over sku  New Document
	//db.Update(appconstant.PRODUCT_COLLECTION, query, bson.M{"$set": this}); errUpdate != nil
	if errUpdate := (*this.MongoSession).Upsert("userdata", query, bson.M{"$set": re}); errUpdate != nil {
		msg := fmt.Errorf("Update Product Document Failed - %v sku :[%s]", errUpdate, re.PimsId)
		logger.Error(msg)
		return msg
	}
	//Update SubProduct Details
	return nil
}

func (this pimsService) GetUser() {

}
