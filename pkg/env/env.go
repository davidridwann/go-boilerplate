package env

import "os"

type Mode string

const (
	Development Mode = "dev"
	Testing     Mode = "testing"
	Staging     Mode = "staging"
	Production  Mode = "production"
)

// Current get the current environment
func Current() Mode {
	return fromString(os.Getenv("APP_ENV"))
}

// fromString convert string literal into constant with Mode type.
func fromString(input string) Mode {
	switch input {
	case string(Staging):
		return Staging
	case string(Production):
		return Production
	case string(Testing):
		return Testing
	}

	return Development
}

// IsDevelopment used to check whether current env is development.
func IsDevelopment() bool {
	return Current() == Development
}

// IsTesting used to check whether current env is testing.
func IsTesting() bool {
	return Current() == Testing
}
