// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NLTokenizer] class.
var nLTokenizerClass = _NLTokenizerClass{objc.GetClass("NLTokenizer")}

type _NLTokenizerClass struct {
	class objc.Class
}

// An interface definition for the [NLTokenizer] class.
type INLTokenizer interface {
	objectivec.IObject
	EnumerateTokensInRangeUsingBlock(range_ unsafe.Pointer, block unsafe.Pointer)
	SetLanguage(language unsafe.Pointer)
	TokenRangeAtIndex(characterIndex uint) unsafe.Pointer
	TokenRangeForRange(range_ unsafe.Pointer) unsafe.Pointer
	TokensForRange(range_ unsafe.Pointer) unsafe.Pointer
}

// A tokenizer that segments natural language text into semantic units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer

type NLTokenizer struct {
	objectivec.Object
}

// NLTokenizerFrom constructs a [NLTokenizer] from an unsafe.Pointer.
//
// A tokenizer that segments natural language text into semantic units.
func NLTokenizerFrom(ptr unsafe.Pointer) NLTokenizer {
	return NLTokenizer{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NLTokenizerClass) Alloc() NLTokenizer {
	rv := objc.Send[NLTokenizer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NLTokenizerClass) New() NLTokenizer {
	rv := objc.Send[NLTokenizer](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NLTokenizer) Init() NLTokenizer {
	rv := objc.Send[NLTokenizer](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NLTokenizer) Autorelease() NLTokenizer {
	rv := objc.Send[NLTokenizer](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLTokenizer creates a new NLTokenizer instance.
func NewNLTokenizer() NLTokenizer {
	return nLTokenizerClass.New()
}


// Creates a tokenizer with the specified unit. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/init(unit:)
func NewNLTokenizerWithUnit(unit unsafe.Pointer) NLTokenizer {
	instance := nLTokenizerClass.Alloc()
	rv := objc.Send[NLTokenizer](instance.ID, objc.Sel("initWithUnit:"), unit)
	rv.Autorelease()
	return rv
}


// Enumerates over a given range of the string and calls the specified block for each token. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/enumerateTokensInRange:usingBlock:
func (n_ NLTokenizer) EnumerateTokensInRangeUsingBlock(range_ unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("enumerateTokensInRange:usingBlock:"), range_, block)
}
// Sets the language of the text to be tokenized. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/setLanguage(_:)
func (n_ NLTokenizer) SetLanguage(language unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLanguage:"), language)
}
// Finds the range of the token at the given index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/tokenRangeAtIndex:
func (n_ NLTokenizer) TokenRangeAtIndex(characterIndex uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("tokenRangeAtIndex:"), characterIndex)
	return rv
}
// Finds the entire range of all tokens contained completely or partially within the specified range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/tokenRangeForRange:
func (n_ NLTokenizer) TokenRangeForRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("tokenRangeForRange:"), range_)
	return rv
}
// Tokenizes the string within the provided range. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/tokensForRange:
func (n_ NLTokenizer) TokensForRange(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("tokensForRange:"), range_)
	return rv
}

