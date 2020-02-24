package service

const collection = "software"

type Software struct {
	SoftwareId    string 	     `bson:"softwareId" json:"softwareId"`
	SoftwareName  string        `bson:"softwareName" json:"softwareName"`
	LnaguageUsed  string        `bson:"languageUsed" json:"languageUsed"`
	CreateDate   string       `bson:"createDate" json:"createDate"`
	Owner  string      `bson:"owner" json:"owner"`
}

// Softwares is an array of Software structs
type Softwares []Software
