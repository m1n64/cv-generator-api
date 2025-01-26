package enums

type GenerateType string

const (
	TypeInformation  GenerateType = "information"
	TypeCertificates GenerateType = "certificates"
	TypeEducations   GenerateType = "educations"
	TypeExperiences  GenerateType = "experiences"
)

var GenerateTypeToKeyMap = map[GenerateType]string{
	TypeInformation:  "info",
	TypeCertificates: "certs",
	TypeEducations:   "edu",
	TypeExperiences:  "exp",
}

func IsValidGenerateType(t string) bool {
	_, exists := GenerateTypeToKeyMap[GenerateType(t)]
	return exists
}

func GetKeyByGenerateType(t string) (string, bool) {
	key, exists := GenerateTypeToKeyMap[GenerateType(t)]
	return key, exists
}
