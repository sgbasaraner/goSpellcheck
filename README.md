goSpellcheck
===================

goSpellcheck is a multi-purpose command line spellchecker that is written purely in Go. 

----------


Features
-------------
**Ignores words with numbers**
Just like Microsoft Word, goSpellcheck ignores words with numbers to keep spellchecking output clean.

**Takes only alphabetic characters and apostrophe**
To ensure compatibility with dictionaries that store words with possessives like

>car's
people's
wall's

goSpellcheck includes apostrophe when spellchecking. All other characters are excluded in order not to mislead the user when a word is followed by a punctuation mark.

**Ignores too long words**
[Wikipedia](https://en.wikipedia.org/wiki/Longest_word_in_English"Longest word in English") claims that the longest word in English dictionaries is Pneumonoultramicroscopicsilicovolcanoconiosis and consists of 45 characters. Therefore goSpellcheck ignores words longer than that, but you can customize it anytime by changing the longestWord constant in the beginning of the main.go file.
```go
const longestWord int = 45
```
 **Lean Mode**
Lean mode allows you to disable all other features mentioned here, and search in a stricter fashion.

----------

Usage
-------------------
```bash
$ ./spellcheck text dictionary [lean]
```
[lean] at the end is optional.
```bash
$ ./spellcheck the_iliad dictionary_of_mythology lean
```
will result in spellcheck running in lean mode. In contrast
```bash
$ ./spellcheck the_iliad dictionary_of_mythology
```
will run in regular mode.
