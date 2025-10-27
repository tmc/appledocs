// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [LinguisticTagger] class.
var (
	LinguisticTaggerClass     _LinguisticTaggerClass
	LinguisticTaggerClassOnce sync.Once
)

func getLinguisticTaggerClass() _LinguisticTaggerClass {
	LinguisticTaggerClassOnce.Do(func() {
		LinguisticTaggerClass = _LinguisticTaggerClass{objc.GetClass("NSLinguisticTagger")}
	})
	return LinguisticTaggerClass
}

type _LinguisticTaggerClass struct {
	class objc.Class
}





// An interface definition for the [LinguisticTagger] class.
type ILinguisticTagger interface {
	objectivec.IObject
	

	// properties:
	DominantLanguage() IString
	String() IString
	SetString(value IString)
	TagSchemes() []string


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (lc _LinguisticTaggerClass) Alloc() LinguisticTagger {
	rv := objc.Send[LinguisticTagger](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LinguisticTaggerClass) New() LinguisticTagger {
	rv := objc.Send[LinguisticTagger](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LinguisticTagger) Init() LinguisticTagger {
	rv := objc.Send[LinguisticTagger](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LinguisticTagger) Autorelease() LinguisticTagger {
	rv := objc.Send[LinguisticTagger](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLinguisticTagger creates a new LinguisticTagger instance.
func NewLinguisticTagger() LinguisticTagger {
	return getLinguisticTaggerClass().New()
}





// Analyze natural language text to tag part of speech and lexical class, identify names, perform lemmatization, and determine the language and script.
//
// provides a uniform interface to a variety of natural language processing functionality with support for many different languages and scripts. You can use this class to segment natural language text into paragraphs, sentences, or words, and tag information about those segments, such as part of speech, lexical class, lemma, script, and language. When you create a linguistic tagger, you specify what kind of information you’re interested in by passing one or more values. Set the property to the natural language text you want to analyze, and the linguistic tagger processes it according to the specified tag schemes. You can then enumerate over the tags in a specified range, using the methods described in Enumerating Linguistic Tags, to get the information requested for a given scheme and unit.


// Analyze natural language text to tag part of speech and lexical class, identify names, perform lemmatization, and determine the language and script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger
type LinguisticTagger struct {
	objectivec.Object
}

// LinguisticTaggerFrom constructs a [LinguisticTagger] from an unsafe.Pointer.
//
// Analyze natural language text to tag part of speech and lexical class, identify names, perform lemmatization, and determine the language and script.
func LinguisticTaggerFrom(ptr unsafe.Pointer) LinguisticTagger {
	return LinguisticTagger{objectivec.Object{objc.ID(ptr)}}
}






// Creates a linguistic tagger instance using the specified tag schemes and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/init(tagSchemes:options:)
func NewLinguisticTaggerWithTagSchemesOptions(tagSchemes []string, opts uint) LinguisticTagger {
	instance := getLinguisticTaggerClass().Alloc()
	rv := objc.Send[LinguisticTagger](instance.ID, objc.Sel("initWithTagSchemes:options:"), tagSchemes, opts)
	rv.Autorelease()
	return rv
}







// Returns the tag schemes available for a particular unit and language on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/availableTagSchemes(for:language:)
func (lc _LinguisticTaggerClass) AvailableTagSchemesForUnitLanguage(unit LinguisticTaggerUnit, language IString) []string {
	rv := objc.Send[[]string](objc.ID(lc.class), objc.Sel("availableTagSchemesForUnit:language:"), unit, language)
	return rv
}


// Returns the tag schemes available for a particular language on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/availableTagSchemes(forLanguage:)
func (lc _LinguisticTaggerClass) AvailableTagSchemesForLanguage(language IString) []string {
	rv := objc.Send[[]string](objc.ID(lc.class), objc.Sel("availableTagSchemesForLanguage:"), language)
	return rv
}


// Returns the dominant language for the specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/dominantLanguage(for:)
func (lc _LinguisticTaggerClass) DominantLanguageForString(string_ IString) IString {
	rv := objc.Send[String](objc.ID(lc.class), objc.Sel("dominantLanguageForString:"), string_)
	return rv
}


// Enumerates over a given string and calls the specified block for each tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/enumerateTags(for:range:unit:scheme:options:orthography:using:)
func (lc _LinguisticTaggerClass) EnumerateTagsForStringRangeUnitSchemeOptionsOrthographyUsingBlock(string_ IString, range_ Range, unit LinguisticTaggerUnit, scheme LinguisticTagScheme, options LinguisticTaggerOptions, orthography IOrthography, block unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("enumerateTagsForString:range:unit:scheme:options:orthography:usingBlock:"), string_, range_, unit, scheme, options, orthography, block)
}


// Returns a tag for a single scheme, for a given linguistic unit, at the specified character position in a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/tag(for:at:unit:scheme:orthography:tokenRange:)
func (lc _LinguisticTaggerClass) TagForStringAtIndexUnitSchemeOrthographyTokenRange(string_ IString, charIndex uint, unit LinguisticTaggerUnit, scheme LinguisticTagScheme, orthography IOrthography, tokenRange RangePointer) LinguisticTag {
	rv := objc.Send[LinguisticTag](objc.ID(lc.class), objc.Sel("tagForString:atIndex:unit:scheme:orthography:tokenRange:"), string_, charIndex, unit, scheme, orthography, tokenRange)
	return rv
}


// Returns an array of linguistic tags and token ranges for a given string and linguistic unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/tags(for:range:unit:scheme:options:orthography:tokenRanges:)
func (lc _LinguisticTaggerClass) TagsForStringRangeUnitSchemeOptionsOrthographyTokenRanges(string_ IString, range_ Range, unit LinguisticTaggerUnit, scheme LinguisticTagScheme, options LinguisticTaggerOptions, orthography IOrthography, tokenRanges []Value) []string {
	rv := objc.Send[[]string](objc.ID(lc.class), objc.Sel("tagsForString:range:unit:scheme:options:orthography:tokenRanges:"), string_, range_, unit, scheme, options, orthography, tokenRanges)
	return rv
}

















// Returns the dominant language of the string set for the linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/dominantLanguage
func (l_ LinguisticTagger) DominantLanguage() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("dominantLanguage"))
	return rv
}


// The string being analyzed by the linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/string
func (l_ LinguisticTagger) String() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("string"))
	return rv
}


// The string being analyzed by the linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/string
func (l_ LinguisticTagger) SetString(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setString:"), value)
}


// Returns the tag schemes configured for this linguistic tagger. For possible values, see .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/tagSchemes
func (l_ LinguisticTagger) TagSchemes() []string {
	rv := objc.Send[[]string](l_.ID, objc.Sel("tagSchemes"))
	return rv
}







