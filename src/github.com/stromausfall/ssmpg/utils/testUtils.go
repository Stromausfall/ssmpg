package utils

import (
	"fmt"
    "reflect"
	"testing"
)
    
func ExpectException(testFunction func(), expectedPanic string, t *testing.T) {
    defer func() {
	    if r := recover(); r != nil {
	        if r == expectedPanic {
	        	// as expected !
	        	return
	        }
	
			switch r.(type) {
				case string:
					t.Error("incorrect panic message : " + fmt.Sprintf("%v", r))
	            default:
					t.Error("incorrect panic type : " + fmt.Sprintf("%v", reflect.TypeOf(r)))
	        }
	
			return
	    }
	
	    // either no panic or not the correct one !
	    t.Error("expected panic did not happen")
	}()
	
	testFunction()
}
