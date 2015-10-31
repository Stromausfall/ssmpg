package input

import (
	"github.com/stromausfall/ssmpg/utils"
	"testing"
)

const topBar = "foo"
const bottomBar = "(c) foo"
const validConfigFileContent = "---\ntopbar: " + topBar + "\nbottombar: " + bottomBar + "\n"

func testCreateConfigExpectException(expectedPanic, argument string, t *testing.T) {
	defer utils.TestExpectException(expectedPanic, t)

	CreateConfigData(argument)
}

func TestCreateConfigData(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.json", validConfigFileContent)
	created := CreateConfigData(path)

	if created.Topbar != topBar {
		t.Error("expected value : " + topBar + " but was : " + created.Topbar)
	}
	if created.Bottombar != bottomBar {
		t.Error("expected value : " + bottomBar + " but was : " + created.Bottombar)
	}
}

func TestCreateConfigDataWrongData(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.json", validConfigFileContent+"asdf asdf")

	testCreateConfigExpectException("error while unmarshaling configFile", path, t)
}

func TestCreateConfigDataWrongFile(t *testing.T) {
	path := utils.CreateTestFileReturnPath("foo.json", validConfigFileContent+"asdf asdf")

	testCreateConfigExpectException("error while opening configFile", path+"234", t)
}
