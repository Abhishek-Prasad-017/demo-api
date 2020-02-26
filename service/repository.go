package service

const collection1 = "software"
const collection2 = "service-details"

type Software struct {
	SoftwareId    string 	    `bson:"softwareId" json:"softwareId"`
	SoftwareName  string        `bson:"softwareName" json:"softwareName"`
	LanguageUsed  string        `bson:"languageUsed" json:"languageUsed"`
	CreateDate   string       	`bson:"createDate" json:"createDate"`
	Owner  string      			`bson:"owner" json:"owner"`
	AdminGroupName string       `bson:"adminGroupName" json:"adminGroupName"`
	DeveloperGroupName string   `bson:"developerGroupName" json:"developerGroupName"`
	QaGroupName string        	`bson:"qaGroupName" json:"qaGroupName"`
	ReleaseGroupName string  	`bson:"releaseGroupName" json:"releaseGroupName"`
	AsgName    string 	    `bson:"asgName" json:"asgName"`
	ServiceName    string 	    `bson:"serviceName" json:"serviceName"`
	Designation    []string 	    `bson:"designation" json:"designation"`
	People    []string 	    `bson:"people" json:"people"`
}

// Softwares is an array of Software structs
type Softwares []Software

