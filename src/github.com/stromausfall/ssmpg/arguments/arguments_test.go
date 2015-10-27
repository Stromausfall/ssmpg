package arguments

import (
	"testing"
)

func TestArgumentNull(t *testing.T) {
    defer func() {
        if r := recover(); r != nil {
        	if r == "argument error - no args" {
        		// as expected !
        		return
        	}
        }
        
        // either no panic or not the correct one !
        t.Fail()
    }()
    
    BasicValidationOfConsoleArguments(nil)
}

func argumentCountTest(argumentCount int, t *testing.T) {
    defer func() {
        if r := recover(); r != nil {
        	if r == "argument error - incorrect argument count" {
        		// as expected !
        		return
        	}
        }
        
        // either no panic or not the correct one !
        t.Fail()
    }()
    
    args := []string{}
    for i := 0; i < argumentCount; i++ {
    	args = append(args, "dummyArg")
    }
    
    BasicValidationOfConsoleArguments(args)
}

func TestIncorrectArgumentCount(t *testing.T) {
	argumentCountTest(0, t)
	argumentCountTest(1, t)
	argumentCountTest(2, t)
	argumentCountTest(4, t)
}
