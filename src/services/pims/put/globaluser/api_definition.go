package globaluser

import (
	"github.com/BoutiqaatREPO/getsetgo/src/common/constants"
	"github.com/BoutiqaatREPO/getsetgo/src/common/logger"
	"github.com/BoutiqaatREPO/getsetgo/src/common/ratelimiter"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/orchestrator"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/utils/healthcheck"
	"github.com/BoutiqaatREPO/getsetgo/src/core/common/versionmanager"
)

type CreateStockAPI struct {
}

func (a *CreateStockAPI) GetVersion() versionmanager.Version {

	return versionmanager.Version{
		Resource: "USER",
		Version:  "V1",
		Action:   "PUT",
		BucketID: constants.OrchestratorBucketDefaultValue,
		Path:     "",
	}
}

func (a *CreateStockAPI) GetOrchestrator() orchestrator.Orchestrator {
	logger.Info("Create Stock Pipeline creation begin")

	createOrchestrator := new(orchestrator.Orchestrator)
	createWorkflow := new(orchestrator.WorkFlowDefinition)
	createWorkflow.Create()

	validationNode := new(RequestValidatorNode)
	validationNode.SetID("Node 2")

	createStockNode := new(createStockAPINode)
	createStockNode.SetID("Node 3")

	//createWorkflow.AddExecutionNode(authNode)
	createWorkflow.AddExecutionNode(validationNode)
	createWorkflow.AddExecutionNode(createStockNode)

	createWorkflow.AddConnection(validationNode, createStockNode)

	createWorkflow.SetStartNode(validationNode)

	createOrchestrator.Create(createWorkflow)
	return *createOrchestrator
}

func (a *CreateStockAPI) GetHealthCheck() healthcheck.HCInterface {
	return new(CreateStockHealthCheck)
}

func (a *CreateStockAPI) Init() {

}

func (a *CreateStockAPI) GetRateLimiter() ratelimiter.RateLimiter {
	return nil
}
