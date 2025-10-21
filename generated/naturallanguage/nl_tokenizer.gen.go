// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Tokenizer] class.
var (
	TokenizerClass     _TokenizerClass
	TokenizerClassOnce sync.Once
)

func getTokenizerClass() _TokenizerClass {
	TokenizerClassOnce.Do(func() {
		TokenizerClass = _TokenizerClass{objc.GetClass("NLTokenizer")}
	})
	return TokenizerClass
}

type _TokenizerClass struct {
	class objc.Class
}

// An interface definition for the [Tokenizer] class.
type ITokenizer interface {
	objectivec.IObject
	EnumerateTokensInRangeUsingBlock(range_ foundation.IRange, block unsafe.Pointer)
	SetLanguage(language ILanguage)
	TokenRangeAtIndex(characterIndex uint) foundation.Range
	TokenRangeForRange(range_ foundation.IRange) foundation.Range
	TokensForRange(range_ foundation.IRange) []foundation.Value
}

// A tokenizer that segments natural language text into semantic units.
//
// creates individual units from natural language text. Define the desired unit (word, sentence, paragraph, or document as declared in the ) for tokenization, and then assign a string to tokenize. The method provides the ranges of the tokens in the string based on the tokenization unit. For more information, see .
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer
type Tokenizer struct {
	objectivec.Object
}

// TokenizerFrom constructs a [Tokenizer] from an unsafe.Pointer.
//
// A tokenizer that segments natural language text into semantic units.
func TokenizerFrom(ptr unsafe.Pointer) Tokenizer {
	return Tokenizer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TokenizerClass) Alloc() Tokenizer {
	rv := objc.Send[Tokenizer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TokenizerClass) New() Tokenizer {
	rv := objc.Send[Tokenizer](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Tokenizer) Init() Tokenizer {
	rv := objc.Send[Tokenizer](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Tokenizer) Autorelease() Tokenizer {
	rv := objc.Send[Tokenizer](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTokenizer creates a new Tokenizer instance.
func NewTokenizer() Tokenizer {
	return getTokenizerClass().New()
}




// Creates a tokenizer with the specified unit.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/init(unit:)
func NewTokenizerWithUnit(unit ITokenUnit) Tokenizer {
	instance := getTokenizerClass().Alloc()
	rv := objc.Send[Tokenizer](instance.ID, objc.Sel("initWithUnit:"), unit)
	rv.Autorelease()
	return rv
}


// Enumerates over a given range of the string and calls the specified block for each token.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/enumerateTokensInRange:usingBlock:
func (t_ Tokenizer) EnumerateTokensInRangeUsingBlock(range_ foundation.IRange, block unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("enumerateTokensInRange:usingBlock:"), range_, block)
}

// Sets the language of the text to be tokenized.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/setLanguage(_:)
func (t_ Tokenizer) SetLanguage(language ILanguage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLanguage:"), language)
}

// Finds the range of the token at the given index.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/tokenRangeAtIndex:
func (t_ Tokenizer) TokenRangeAtIndex(characterIndex uint) foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("tokenRangeAtIndex:"), characterIndex)
	return rv
}

// Finds the entire range of all tokens contained completely or partially within the specified range.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/tokenRangeForRange:
func (t_ Tokenizer) TokenRangeForRange(range_ foundation.IRange) foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("tokenRangeForRange:"), range_)
	return rv
}

// Tokenizes the string within the provided range.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/tokensForRange:
func (t_ Tokenizer) TokensForRange(range_ foundation.IRange) []foundation.Value {
	rv := objc.Send[[]foundation.Value](t_.ID, objc.Sel("tokensForRange:"), range_)
	return rv
}

// The text to be tokenized.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/string
func (t_ Tokenizer) String() appkit.string {
	rv := objc.Send[appkit.string](t_.ID, objc.Sel("string"))
	return rv
}


// SetString sets the value of the string property.
// The text to be tokenized.

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/string
func (t_ Tokenizer) SetString(value appkit.string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}

// The linguistic unit that this tokenizer uses.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/unit
func (t_ Tokenizer) Unit() TokenUnit {
	rv := objc.Send[TokenUnit](t_.ID, objc.Sel("unit"))
	return rv
}


