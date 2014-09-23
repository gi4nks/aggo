aggo
====

I track all my daily notes in **markdown** files divided per day. Often I need to search what I did or some specific keywords in the archive and I use existing tools. 
But I wanted to develo something from scratch in **go**, to learn the language, to do it as I want.

This is an experiment and currently the code is really simple.

Things done (but still to improve):
- create Index structure 
- create File structure
- create a Logger structure
- create Scan of files word by word

Things to do:
- manage stop words and so on to purge archive and have a better index class
- add serialization/deserialization in bson of index
- add serialization/deserialization in json of index
- implement query language
- implement a reverse index
- create statistics
