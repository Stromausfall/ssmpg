package app

// tests the examples from : http://daringfireball.net/projects/markdown/syntax

import (
	"testing"
	"fmt"
	"strings"
)

// testing : http://daringfireball.net/projects/markdown/syntax#html
func TestBlockHtml(t *testing.T) {
	input := ""
	input += "This is a regular paragraph.\n"
	input += "\n"
	input += "<table>\n"
	input += "    <tr>\n"
	input += "        <td>*Foo*</td>\n"
	input += "    </tr>\n"
	input += "</table>\n"
	input += "\n"
	input += "This is another regular paragraph.\n"
	
	expected := ""
	expected += "This is a regular paragraph.\n"
	expected += "\n"
	expected += "<table>\n"
	expected += "    <tr>\n"
	expected += "        <td>*Foo*</td>\n"
	expected += "    </tr>\n"
	expected += "</table>\n"
	expected += "\n"
	expected += "This is another regular paragraph.</br>\n"
	
	result := Process(input)
	
	if strings.Compare(result, expected) != 0 {
		// they are not the same !
		fmt.Println("expected : '" + expected + "'")
		fmt.Println("but was  : '" + result + "'")
		t.Fail()
	}
}
