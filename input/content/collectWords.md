title: Collecting words from text files
date: 2015-04-07 12:15
categories:
- Go

---

The following GO program ([collectWords.go](http://matthias-auer.net/Projects/CollectWords/collectWords.go)) takes one argument (a path to a folder containing files) and then collects all the words in the files. Then this collection is written to an output text file inside that directory.

The program works in the following way :

 - first the directory is scanned for all files
 - then for each of the files a go routine is started to run concurrently - which opens and reads all the content of that file
 - when chunk of a certain size has been read, a new worker is spawned that collects all the words from that chunk
 - finally another method is started that collects these collections and merges them (this process ends once all other concurrent workers have finished)
 - finally the collection is written to a file
