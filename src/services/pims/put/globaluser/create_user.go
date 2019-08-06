package globaluser

import (
	"fmt"

	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
)

type createStockAPINode struct {
	id string
}

func (a *createStockAPINode) SetID(id string) {
	a.id = id
}

func (a createStockAPINode) GetID() (string, error) {
	return a.id, nil
}

func (a createStockAPINode) Name() string {
	return "createStockAPINode"
}

func (a createStockAPINode) Execute(io orchestrator.WorkFlowData) (orchestrator.WorkFlowData, error) {
	logger.Info("Executing Create Stock Node")

	//stockDetail
	pimsUserDetail, errGet := io.IOData.Get(MID_LAYER_PIMS)
	if errGet != nil {
		return io, &constants.AppError{
			Code:             constants.ResourceErrorCode,
			Message:          "Could not GET request data",
			DeveloperMessage: fmt.Sprintf("Error: %s", errGet.Error()),
		}
	}

	data, ok := pimsUserDetail.(reqResponse)
	if !ok {
		return io, &constants.AppError{
			Code:             constants.IncorrectDataErrorCode,
			Message:          "Unable to read data from IOData",
			DeveloperMessage: "Get Data from IOData Failed",
		}
	}

	intf := data.PimsInterface
	err := intf.UpdateUser(data.middleLayerPims)
	if err != nil {
		io.IOData.Set(constants.Result, ResponseData{
			Success: false,
		})
	}
	io.IOData.Set(constants.Result, ResponseData{
		Success: true,
	})
	return io, nil
}
