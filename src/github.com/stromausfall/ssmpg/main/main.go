package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/stromausfall/ssmpg/conversion"
)

func main() {
	
	//args := os.Args[1:]
	//x1, x2, x3 := BasicValidationOfConsoleArguments(args)
	
	dat, _ := ioutil.ReadFile("C:\\Users\\mail_000\\OneDrive\\GO\\ssmpg\\base.html")
    data := string(dat)
    topBarString := conversion.Convert("# My Homepage #")
    bodyString := conversion.Convert("foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* foo *foo* ")
    indexString := conversion.Convert("[This link](http://example.net/)")
    bottomBarString := conversion.Convert("{C} bla")
    
    data = strings.Replace(data, "XXX-TOPBAR-XXX", topBarString, 1)
    data = strings.Replace(data, "XXX-BODY-XXX", bodyString, 1)
    data = strings.Replace(data, "XXX-INDEX-XXX", indexString, 1)
    data = strings.Replace(data, "XXX-BOTTOMBAR-XXX", bottomBarString, 1)
    
    fmt.Print(data)
}
