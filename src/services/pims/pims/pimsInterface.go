package pims

// PimsInterface Reprent pims actvity
type PimsInterface interface {
	UpdateUser(RequestUser) error
	GetUser()
}
