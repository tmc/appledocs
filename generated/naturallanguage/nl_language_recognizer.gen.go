// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LanguageRecognizer] class.
var (
	LanguageRecognizerClass     _LanguageRecognizerClass
	LanguageRecognizerClassOnce sync.Once
)

func getLanguageRecognizerClass() _LanguageRecognizerClass {
	LanguageRecognizerClassOnce.Do(func() {
		LanguageRecognizerClass = _LanguageRecognizerClass{objc.GetClass("NLLanguageRecognizer")}
	})
	return LanguageRecognizerClass
}

type _LanguageRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [LanguageRecognizer] class.
type ILanguageRecognizer interface {
	objectivec.IObject
	// properties:
	DominantLanguage() objc.IObject /* cross-framework: Language */
	LanguageConstraints() []string
	SetLanguageConstraints(value []string)
	LanguageHints() foundation.IDictionary
	SetLanguageHints(value foundation.IDictionary)
	// methods:
	LanguageHypothesesWithMaximum(maxHypotheses uint) foundation.IDictionary
	ProcessString(string_ objc.IObject /* cross-framework: NSString */)
	Reset()
}

// The language of a body of text.
//
// An object automatically detects the language of a piece of text. It performs language identification by: Identifying the dominant script of a piece of text. Some languages have a unique script (like Greek), but others share the same script (like English, French, and German, which all share the Latin script). Identifying the language itself. The identification obtained from an object can be either a single most likely language, access through , or a set of language candidates with probabilities, using . You can reset the recognizer to its initial state, to be reused for new analysis. Use the convenience method, , to get the most likely language without creating an .


// The language of a body of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer
type LanguageRecognizer struct {
	objectivec.Object
}

// LanguageRecognizerFrom constructs a [LanguageRecognizer] from an unsafe.Pointer.
//
// The language of a body of text.
func LanguageRecognizerFrom(ptr unsafe.Pointer) LanguageRecognizer {
	return LanguageRecognizer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LanguageRecognizerClass) Alloc() LanguageRecognizer {
	rv := objc.Send[LanguageRecognizer](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LanguageRecognizerClass) New() LanguageRecognizer {
	rv := objc.Send[LanguageRecognizer](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LanguageRecognizer) Init() LanguageRecognizer {
	rv := objc.Send[LanguageRecognizer](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LanguageRecognizer) Autorelease() LanguageRecognizer {
	rv := objc.Send[LanguageRecognizer](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLanguageRecognizer creates a new LanguageRecognizer instance.
func NewLanguageRecognizer() LanguageRecognizer {
	return getLanguageRecognizerClass().New()
}




// Finds the most likely language of a piece of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/dominantLanguage(for:)
func (lc _LanguageRecognizerClass) DominantLanguageForString(string_ objc.IObject /* cross-framework: NSString */) objc.IObject /* cross-framework: Language */ {
	rv := objc.Send[Language](objc.ID(lc.class), objc.Sel("dominantLanguageForString:"), string_)
	return rv
}


// Generates the probabilities of possible languages for the processed text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/languageHypothesesWithMaximum:
func (l_ LanguageRecognizer) LanguageHypothesesWithMaximum(maxHypotheses uint) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](l_.ID, objc.Sel("languageHypothesesWithMaximum:"), maxHypotheses)
	return rv
}


// Analyzes the piece of text to determine its dominant language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/processString(_:)
func (l_ LanguageRecognizer) ProcessString(string_ objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("processString:"), string_)
}


// Resets the recognizer to its initial state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/reset()
func (l_ LanguageRecognizer) Reset() {
	objc.Send[objc.ID](l_.ID, objc.Sel("reset"))
}


// The most likely language for the processed text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/dominantLanguage
func (l_ LanguageRecognizer) DominantLanguage() objc.IObject /* cross-framework: Language */ {
	rv := objc.Send[Language](l_.ID, objc.Sel("dominantLanguage"))
	return rv
}


// Limits the set of possible languages that the recognizer will return.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/languageConstraints
func (l_ LanguageRecognizer) LanguageConstraints() []string {
	rv := objc.Send[[]string](l_.ID, objc.Sel("languageConstraints"))
	return rv
}


// Limits the set of possible languages that the recognizer will return.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/languageConstraints
func (l_ LanguageRecognizer) SetLanguageConstraints(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](l_.ID, objc.Sel("setLanguageConstraints:"), nsArray)
}


// A dictionary that maps languages to their probabilities in the language identification process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/languageHints-3gy00
func (l_ LanguageRecognizer) LanguageHints() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](l_.ID, objc.Sel("languageHints"))
	return rv
}


// A dictionary that maps languages to their probabilities in the language identification process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/languageHints-3gy00
func (l_ LanguageRecognizer) SetLanguageHints(value foundation.IDictionary) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLanguageHints:"), value)
}


