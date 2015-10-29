package input

import (
	"testing"
	"github.com/stromausfall/ssmpg/utils"
)

func textCreateConfigExpectException(expectedPanic, argument string, t *testing.T) {
	utils.ExpectException(
		func() {
			CreateConfigFile(argument)
		},
		expectedPanic,
		t)
}

func TestCreateConfigFile(t *testing.T) {
	textCreateConfigExpectException("create config - incorrect argument", "", t)
}
