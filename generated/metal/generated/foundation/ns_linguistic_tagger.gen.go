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
	SetDominantLanguage(value IString)
	String() IString
	SetString(value IString)
	TagSchemes() LinguisticTagScheme /* not a class type */
	SetTagSchemes(value LinguisticTagScheme /* not a class type */)
	// methods:
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



// Returns the dominant language of the string set for the linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagger/dominantlanguage
func (l_ LinguisticTagger) DominantLanguage() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("dominantLanguage"))
	return rv
}


// Returns the dominant language of the string set for the linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagger/dominantlanguage
func (l_ LinguisticTagger) SetDominantLanguage(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDominantLanguage:"), value)
}


// The string being analyzed by the linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagger/string
func (l_ LinguisticTagger) String() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("string"))
	return rv
}


// The string being analyzed by the linguistic tagger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagger/string
func (l_ LinguisticTagger) SetString(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setString:"), value)
}


// Returns the tag schemes configured for this linguistic tagger. For possible values, see
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagger/tagschemes
func (l_ LinguisticTagger) TagSchemes() LinguisticTagScheme /* not a class type */ {
	rv := objc.Send[LinguisticTagScheme](l_.ID, objc.Sel("tagSchemes"))
	return rv
}


// Returns the tag schemes configured for this linguistic tagger. For possible values, see
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslinguistictagger/tagschemes
func (l_ LinguisticTagger) SetTagSchemes(value LinguisticTagScheme /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTagSchemes:"), value)
}



