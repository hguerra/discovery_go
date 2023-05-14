package config

func GetActiveProfile() string {
	return GetString("profile")
}

func IsDev() bool {
	return GetActiveProfile() == "development"
}

func IsProd() bool {
	return GetActiveProfile() == "production"
}

func IsTest() bool {
	return GetActiveProfile() == "test"
}
