package input

import (
	"github.com/stromausfall/ssmpg/utils/test"
	"testing"
)

func testCreateConfigExpectException(expectedPanic, argument string, t *testing.T) {
	defer test.TestExpectException(expectedPanic, t)

	CreateConfigData(argument)
}

func TestCreateConfigData(t *testing.T) {
	path := test.CreateTestFileReturnPath("foo.json", test.ConfigDataContent)
	created := CreateConfigData(path)

	if created.Topbar != test.ConfigDataTopBar {
		t.Error("expected value : " + test.ConfigDataTopBar + " but was : " + created.Topbar)
	}
	if created.Bottombar != test.ConfigDataBottomBar {
		t.Error("expected value : " + test.ConfigDataBottomBar + " but was : " + created.Bottombar)
	}
	if created.TitleType != test.ConfigDataTitleAndCatgoryType {
		t.Error("expected value : " + test.ConfigDataTitleAndCatgoryType + " but was : " + created.TitleType)
	}
	if created.IndexName != test.ConfigDataIndexName {
		t.Error("expected value : " + test.ConfigDataIndexName + " but was : " + created.IndexName)
	}
}

func TestCreateConfigDataWrongData(t *testing.T) {
	path := test.CreateTestFileReturnPath("foo.json", test.ConfigDataContent+"asdf asdf")

	testCreateConfigExpectException("error while unmarshaling configFile", path, t)
}

func TestCreateConfigDataWrongFile(t *testing.T) {
	path := test.CreateTestFileReturnPath("foo.json", test.ConfigDataContent+"asdf asdf")

	testCreateConfigExpectException("error while opening configFile", path+"234", t)
}
