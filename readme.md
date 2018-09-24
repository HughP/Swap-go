# SWAP

## Operation

Command line function to translate a series of Unicode strings from one arrangement to another. It operates on a file editing in place (one argument). It has a look up table with a before and after syntax `before:after`i.e. `q:1` would replace the `q` with`1`. a third optional argument can change the delimiter.

## Tests

### ASCII [a-z] test
Status: passing

q:a
a:q

### Unicode combining characters test
Nature: This passes if UAX \#29 is implemented correctly. If NFD is used in comparisons, then it will fail and `ö:o` will also fail.
Status: failing
ʋ̈:a
a:ʋ̈

### Unicode Punctuation - numeric test
Status: passing

!:1
1:!

### Unicode composed character interchange test
status: passing
í:i
i:í

### Unicode decomposition test
nature: will pass with NFC will fail with NFD
status: passing

o:ö
ö:o

### Unicode letter and symbol test
status: passing
˗:+
+:˗

Repttitive
k:T
ɔ:k
T:ɔ
### unique character test with UAX 29
Status: fail
ʋ:4
4:ʋ

### Hash test, will hash also convert
Status: fail
8:\#
\#:8

### forward slash test
Status: passing
/:§
§:/

### back slash test
Status: fail - because not working via hash tag
\\:;
;:{
}:\

## Fails
as currently written this program fails:

1. because it can not replace Unicode NFD encoded characters following the Unicode text segmentation algorithm. http://unicode.org/reports/tr29/
2. because one can not put the path to the translator or the corpus such as `swap/translator.txt`.
3. because it can not convert hash tags
4. because it can not get more content after a hash tag
5. backslash test fails.


What Hugh doesn't understand is why the program cant be written so that any string before the delimiter is matched in the text (corpus) and replaced with any string after the delimiter? why is the abstraction need to go through a rune by rune comparison?
