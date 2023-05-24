package config

var profile = ""

func GetActiveProfile() string {
	if profile != "" {
		return profile
	}
	profile = GetString("profile")
	if profile == "" {
		profile = "test"
	}
	return profile
}

func IsProd() bool {
	return GetActiveProfile() == "production"
}

func IsDev() bool {
	return GetActiveProfile() == "development"
}

func IsTest() bool {
	return GetActiveProfile() == "test"
}
