package utils

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

func GetValidator() *validator.Validate {
	return validate
}
