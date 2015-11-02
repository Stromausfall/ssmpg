title: ArchChinese to Anki
date: 2014-12-09 23:48
categories:
- Python

---

[http://www.archchinese.com/](http://www.archchinese.com/) offers hands down the best and most comprehensive vocabulary resource for learning Chinese I've found so far. Printing their marvellous flashcards works great, but the lack of a systematic approach when using flashcards manually (when to repeat and how often to repeat certain vocablury) - is a negative factor.

So I stumbled upon [http://ankisrs.net/](http://ankisrs.net/) and the excellent port of Anki for Android 
[AnkiDroid](https://play.google.com/store/apps/details?id=com.ichi2.anki&hl=en) - but losing all flashcards manually created in ArchChinese is not an option. Of course you can manully create a deck and copy all data from archchinese.

But archchinese can export vocabulary lists to excel and using a [python script (has to be renamed to .py)](http://www.matthias-auer.net/Projects/ArchChineseToAnki/ArchChineseToAnki.txt) I created it can be converted to the .csv format which can be imported by anki - making the progress of migration quick and painful.

To be able to use the script you have to have the python xlrd paket install (to instal the python package use : <code>pip install xlrd</code>)

<code>python ArchChineseToAnki.py inputFileFromArchChinese.xls</code>

the created outputfile will be named : <code>inputFileFromArchChinese.xls.csv</code> and can directly be imported into an anki deck !

