package globaluser

type CreateStockHealthCheck struct {
}

func (a CreateStockHealthCheck) GetName() string {
	return "CreateStockHealthCheck"
}

func (a CreateStockHealthCheck) GetHealth() map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
	}
}
