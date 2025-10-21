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
	EnumerateTagsInRangeSchemeOptionsUsingBlock(range_ Range, tagScheme unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer)
}

// Analyze natural language text to tag part of speech and lexical class, identify names, perform lemmatization, and determine the language and script.
//
// provides a uniform interface to a variety of natural language processing functionality with support for many different languages and scripts. You can use this class to segment natural language text into paragraphs, sentences, or words, and tag information about those segments, such as part of speech, lexical class, lemma, script, and language. When you create a linguistic tagger, you specify what kind of information you’re interested in by passing one or more values. Set the property to the natural language text you want to analyze, and the linguistic tagger processes it according to the specified tag schemes. You can then enumerate over the tags in a specified range, using the methods described in Enumerating Linguistic Tags, to get the information requested for a given scheme and unit.
//
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

// Alloc allocates a new instance without initialization.
func (lc _LinguisticTaggerClass) Alloc() LinguisticTagger {
	rv := objc.Send[LinguisticTagger](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Enumerates over a given range of the string and calls the specified block for each tag.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/enumerateTags(in:scheme:options:using:)
func (l_ LinguisticTagger) EnumerateTagsInRangeSchemeOptionsUsingBlock(range_ Range, tagScheme unsafe.Pointer, opts unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("enumerateTagsInRange:scheme:options:usingBlock:"), range_, tagScheme, opts, block)
}

// Returns the dominant language of the string set for the linguistic tagger.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagger/dominantlanguage
func (l_ LinguisticTagger) DominantLanguage() string {
	rv := objc.Send[string](l_.ID, objc.Sel("dominantLanguage"))
	return rv
}


// SetDominantLanguage sets the value of the dominantLanguage property.
// Returns the dominant language of the string set for the linguistic tagger.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagger/dominantlanguage
func (l_ LinguisticTagger) SetDominantLanguage(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDominantLanguage:"), objc.String(value))
}

// The string being analyzed by the linguistic tagger.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagger/string
func (l_ LinguisticTagger) String() string {
	rv := objc.Send[string](l_.ID, objc.Sel("string"))
	return rv
}


// SetString sets the value of the string property.
// The string being analyzed by the linguistic tagger.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagger/string
func (l_ LinguisticTagger) SetString(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setString:"), objc.String(value))
}

// Returns the tag schemes configured for this linguistic tagger. For possible values, see
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagger/tagschemes
func (l_ LinguisticTagger) TagSchemes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("tagSchemes"))
	return rv
}


// SetTagSchemes sets the value of the tagSchemes property.
// Returns the tag schemes configured for this linguistic tagger. For possible values, see

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagger/tagschemes
func (l_ LinguisticTagger) SetTagSchemes(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTagSchemes:"), value)
}



