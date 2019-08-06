package globaluser

import (
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
	"github.com/OrangeMoney/pims/src/services/pims/pims"
)

type RequestValidatorNode struct {
	id string
}

func (a *RequestValidatorNode) SetID(id string) {
	a.id = id
}

func (a RequestValidatorNode) GetID() (string, error) {
	return a.id, nil
}

func (a RequestValidatorNode) Name() string {
	return "RequestValidatorNode"
}

func (a RequestValidatorNode) Execute(io orchestrator.WorkFlowData) (orchestrator.WorkFlowData, error) {
	logger.Info("Stock creation validator node")

	//Business Logic
	appHttpReq, err := io.IOData.GetRequest()
	if err != nil {
		return io, &constants.AppError{
			Code:             constants.ParamsInValidErrorCode,
			Message:          "Could not read request body",
			DeveloperMessage: fmt.Sprintf("Error: %s", err.Error()),
		}
	}

	useRequest := pimsRequest{}
	err = appHttpReq.LoadBody(&useRequest)
	if err != nil {
		return io, &constants.AppError{
			Code:             constants.ParamsInValidErrorCode,
			Message:          "Request Unmarshalling failed",
			DeveloperMessage: fmt.Sprintf("Error: %s", err.Error()),
		}
	}
	pimIntf, err := pims.GetPimsIntf("")
	if err != nil {
		return io, &constants.AppError{
			Code:             constants.ParamsInValidErrorCode,
			Message:          "Get Pims Interface return with error",
			DeveloperMessage: fmt.Sprintf("Error: %s", err.Error()),
		}
	}
	reqResps := reqResponse{
		middleLayerPims: pims.RequestUser{
			PimsId:    "101",
			LastName:  useRequest.LastName,
			FirstName: useRequest.FirstName,
			DOB:       useRequest.DOB,
		},
		PimsInterface: pimIntf,
	}

	io.IOData.Set(MID_LAYER_PIMS, reqResps)
	return io, nil
}
