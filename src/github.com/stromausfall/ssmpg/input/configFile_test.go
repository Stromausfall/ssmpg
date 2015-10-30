package input

import (
	"testing"
	"github.com/stromausfall/ssmpg/utils"
)

const topBar = "foo"
const bottomBar = "(c) foo"
const validConfigFileContent = "{\n\"topBar\":\"" + topBar + "\",\n\"bottomBar\":\"" + bottomBar+ "\"\n}"

func textCreateConfigExpectException(expectedPanic, argument string, t *testing.T) {
	utils.ExpectException(
		func() {
			CreateConfigFile(argument)
		},
		expectedPanic,
		t)
}

func TestCreateConfigFile(t *testing.T) {
	path := utils.CreateTestFile("foo.json", validConfigFileContent)
	created := CreateConfigFile(path)
	
	if created.TopBar != topBar {
		t.Error("expected value : " + topBar + " but was : " + created.TopBar)
	}
	if created.BottomBar != bottomBar {
		t.Error("expected value : " + bottomBar + " but was : " + created.BottomBar)
	}
}
