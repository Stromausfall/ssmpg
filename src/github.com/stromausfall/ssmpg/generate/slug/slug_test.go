package slug

import (
	"fmt"
	"github.com/stromausfall/ssmpg/utils/test"
	"testing"
)

func testGenerateSlug(preSlugify, postSlugify string, t *testing.T) {
	actualSlug := GenerateSlug(preSlugify)

	if actualSlug != postSlugify {
		fmt.Println("to slugify : " + preSlugify)
		fmt.Println("expected slug : " + postSlugify)
		fmt.Println("but was : " + actualSlug)
		t.Error("problem with CreateSlug")
	}
}

func TestGenerateSlugEmptyArg(t *testing.T) {
	defer test.TestExpectException("no argument passed to convert", t)

	GenerateSlug("")
}

func TestGenerateSlugWithoutSomethingToSlugify(t *testing.T) {
	testGenerateSlug("foo", "foo", t)
}

func TestGenerateSlugSimpleSlug(t *testing.T) {
	testGenerateSlug("foo bar", "foo-bar", t)
}

func TestGenerateSlugMoreWhitespaces(t *testing.T) {
	whitespaces := " "

	for i := 0; i < 25; i++ {
		testGenerateSlug("foo"+whitespaces+"bar", "foo-bar", t)
		whitespaces += " "
	}
}

func TestGenerateSlugWithPunctuation(t *testing.T) {
	testGenerateSlug("foo.bar", "foo-bar", t)
}

func TestGenerateSlugWithOtherLanguages(t *testing.T) {
	testGenerateSlug("你.好", "你-好", t)
	testGenerateSlug("你#好", "你-好", t)
	testGenerateSlug("你????好", "你-好", t)
}

func TestGenerateSlugWithNumbers(t *testing.T) {
	testGenerateSlug("foo3.bar", "foo3-bar", t)
}

func TestGenerateSlugLowerCase(t *testing.T) {
	testGenerateSlug("Foo3.Bar - 你好", "foo3-bar-你好", t)
}

func TestGenerateSlugRemoveTrailingAndLeadingHyphens(t *testing.T) {
	testGenerateSlug("Foo3.Bar - 你好-", "foo3-bar-你好", t)
	testGenerateSlug("-Foo3.Bar - 你好", "foo3-bar-你好", t)
	testGenerateSlug("-Foo3.Bar - 你好-", "foo3-bar-你好", t)
}
