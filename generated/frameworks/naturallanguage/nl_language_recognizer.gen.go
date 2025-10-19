// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NLLanguageRecognizer] class.
var nLLanguageRecognizerClass = _NLLanguageRecognizerClass{objc.GetClass("NLLanguageRecognizer")}

type _NLLanguageRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [NLLanguageRecognizer] class.
type INLLanguageRecognizer interface {
	objectivec.IObject
	LanguageHypothesesWithMaximum(maxHypotheses uint) unsafe.Pointer
	ProcessString(string string)
	Reset()
}

// The language of a body of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer

type NLLanguageRecognizer struct {
	objectivec.Object
}

// NLLanguageRecognizerFrom constructs a [NLLanguageRecognizer] from an unsafe.Pointer.
//
// The language of a body of text.
func NLLanguageRecognizerFrom(ptr unsafe.Pointer) NLLanguageRecognizer {
	return NLLanguageRecognizer{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NLLanguageRecognizerClass) Alloc() NLLanguageRecognizer {
	rv := objc.Send[NLLanguageRecognizer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NLLanguageRecognizerClass) New() NLLanguageRecognizer {
	rv := objc.Send[NLLanguageRecognizer](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NLLanguageRecognizer) Init() NLLanguageRecognizer {
	rv := objc.Send[NLLanguageRecognizer](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NLLanguageRecognizer) Autorelease() NLLanguageRecognizer {
	rv := objc.Send[NLLanguageRecognizer](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLLanguageRecognizer creates a new NLLanguageRecognizer instance.
func NewNLLanguageRecognizer() NLLanguageRecognizer {
	return nLLanguageRecognizerClass.New()
}




// Finds the most likely language of a piece of text. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/dominantLanguage(for:)
func (nc _NLLanguageRecognizerClass) DominantLanguageForString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("dominantLanguageForString:"), string)
	return rv
}
// Generates the probabilities of possible languages for the processed text. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/languageHypothesesWithMaximum:
func (n_ NLLanguageRecognizer) LanguageHypothesesWithMaximum(maxHypotheses uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("languageHypothesesWithMaximum:"), maxHypotheses)
	return rv
}
// Analyzes the piece of text to determine its dominant language. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/processString(_:)
func (n_ NLLanguageRecognizer) ProcessString(string string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("processString:"), string)
}
// Resets the recognizer to its initial state. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLLanguageRecognizer/reset()
func (n_ NLLanguageRecognizer) Reset() {
	objc.Send[objc.ID](n_.ID, objc.Sel("reset"))
}

