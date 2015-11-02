package test

import (
)

const ContentDataTitle = "Old Pages"
const ContentDataDate = "2015-03-24 22:28"
const ContentDataContent = "Old unsupported pages :"
var ContentDataCategories = []string{"X1", "X2"}
var ContentData = "title: " + ContentDataTitle + "\ndate: " + ContentDataDate + "\ncategories:\n- " + ContentDataCategories[0] + "\n- " + ContentDataCategories[1] + "\n---\n" + ContentDataContent

const ConfigDataTopBar = "foo"
const ConfigDataBottomBar = "(c) foo"
const ConfigDataTitleAndCatgoryType = "h2"
const ConfigDataIndexName = "Index"
const ConfigDataContent = "---\ntopbar: " + ConfigDataTopBar + "\nbottombar: " + ConfigDataBottomBar + "\ntitletype: " + ConfigDataTitleAndCatgoryType + "\nindexname: " + ConfigDataIndexName + "\n"

const BaseHtmlContent = "XXX-TITLE-XXX ||| XXX-TOPBAR-XXX ||| XXX-BODY-XXX ||| XXX-INDEX-XXX ||| XXX-BOTTOMBAR-XXX"
