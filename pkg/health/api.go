package health

import (
	"errors"
)

type Status uint8

const (
	HEALTHY      Status = 0
	INITIALIZING Status = 1
	UNHEALTHY    Status = 2
	UNKNOWN      Status = 3
)

type CheckResult struct {
	Status  Status
	Message string
}

type Result struct {
	Status     Status
	Components map[string]CheckResult
}

type Check func() CheckResult

var (
	CheckForComponentAlreadyRegistered = errors.New("a check for the requested component is already registered")
	checkerInstance                    *checker
)

func init() {
	checkerInstance = &checker{
		componentChecks: map[string]Check{},
	}
}

func CheckerInstance() Checker {
	return checkerInstance
}
