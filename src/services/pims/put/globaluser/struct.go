package globaluser

import (
	"github.com/OrangeMoney/pims/src/services/pims/pims"

	validator "gopkg.in/asaskevich/govalidator.v9"
)

const MID_LAYER_PIMS = "pimsMiddleRequest"

//Request
type pimsRequest struct {
	FirstName string `json:"firstName" valid:"required~first name  cannot be empty"`
	LastName  string `json:"lastName"`
	DOB       string `json:"dob"`
}

type ResponseData struct {
	Success          bool   `json:"success"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage,omitempty"`
	SKU              string `json:"sku"`
}

type reqResponse struct {
	middleLayerPims pims.RequestUser
	Responses       ResponseData
	PimsInterface   pims.PimsInterface
}

/*
return mongoObject
From consumer can get mongodata
in case of success
*/
func (this pimsRequest) Audit() error {
	_, err := validator.ValidateStruct(this)
	if err != nil {
		return err
	}
	//other validation can be done here if any
	return nil
}
