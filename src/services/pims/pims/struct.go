package pims

type RequestUser struct {
	FirstName string `bson:"firstname"`
	LastName  string `bson:"lastname"`
	DOB       string `bson:"dob"`
	PimsId    string `bson:"pimsid"`
}
