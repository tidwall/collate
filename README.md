Collate
======
<a href="https://travis-ci.org/tidwall/collate"><img src="https://img.shields.io/travis/tidwall/collate.svg?style=flat-square" alt="Build Status"></a>
<a href="https://godoc.org/github.com/tidwall/collate"><img src="https://img.shields.io/badge/api-reference-blue.svg?style=flat-square" alt="GoDoc"></a>

Collate is a simple collation library for comparing strings in various languages for Go. 
It's designed for the [BuntDB](https://github.com/tidwall/buntdb) project, and 
is simliar to the 
[collation](https://msdn.microsoft.com/en-us/library/ms174596.aspx) that is 
found in traditional database systems

The idea is that you call a function with a collation name and it generates 
a `Less(a, b string) bool` function that can be used for sorting in B-Tree 
style databases.

Install
-------
```
go get -u github.com/tidwall/collate
```

Example
-------
```go
// create a case-insensitive collation for spanish.
less := collate.Index("SPANISH_CI")
println(less("Hola", "hola"))

// Output
// false
```

Options
-------

### Case Sensitivity

Add `_CI` to the collation name to specify case-insensitive comparing.
Add `_CS` for case-sensitive compares, this is the default.

```go
collate.Index("SPANISH_CI") // Case-insensitive collation for spanish
collate.Index("SPANISH_CS") // Case-sensitive collation for spanish
```

### Loose Compares

Add `_LOOSE` to ignores diacritics, case and weight.

### Numeric Compares

Add `_NUM` to specifies that numbers should sort numerically ("2" < "12")

Supported Languages
-------------------

```
Afrikaans
Albanian
AmericanEnglish
Amharic
Arabic
Armenian
Azerbaijani
Bengali
BrazilianPortuguese
BritishEnglish
Bulgarian
Burmese
CanadianFrench
Catalan
Chinese
Croatian
Czech
Danish
Dutch
English
Estonian
EuropeanPortuguese
EuropeanSpanish
Filipino
Finnish
French
Georgian
German
Greek
Gujarati
Hebrew
Hindi
Hungarian
Icelandic
Indonesian
Italian
Japanese
Kannada
Kazakh
Khmer
Kirghiz
Korean
Lao
LatinAmericanSpanish
Latvian
Lithuanian
Macedonian
Malay
Malayalam
Marathi
ModernStandardArabic
Mongolian
Nepali
Norwegian
Persian
Polish
Portuguese
Punjabi
Romanian
Russian
Serbian
SerbianLatin
SimplifiedChinese
Sinhala
Slovak
Slovenian
Spanish
Swahili
Swedish
Tamil
Telugu
Thai
TraditionalChinese
Turkish
Ukrainian
Urdu
Uzbek
Vietnamese
Zulu
```





## Contact
Josh Baker [@tidwall](http://twitter.com/tidwall)

## License

GJSON source code is available under the MIT [License](/LICENSE).


