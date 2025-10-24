// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NLTokenizer */


/* debug [class_header]: Header for NLTokenizer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Tokenizer */
// An interface definition for the [Tokenizer] class.
type ITokenizer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Tokenizer */
	// properties:
	String() objc.IObject /* cross-framework: NSString */
	SetString(value objc.IObject /* cross-framework: NSString */)
	Unit() TokenUnit
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Tokenizer */
	// methods:
	EnumerateTokensInRangeUsingBlock(range_ corefoundation.Range, block bool)
	SetLanguage(language Language /* typedef */)
	TokenRangeAtIndex(characterIndex uint) corefoundation.Range
	TokenRangeForRange(range_ corefoundation.Range) corefoundation.Range
	TokensForRange(range_ corefoundation.Range) []foundation.Value
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Tokenizer */
// Alloc allocates a new instance without initialization.
func (tc _TokenizerClass) Alloc() Tokenizer {
	rv := objc.Send[Tokenizer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Tokenizer */
// A tokenizer that segments natural language text into semantic units.
//
// creates individual units from natural language text. Define the desired unit (word, sentence, paragraph, or document as declared in the ) for tokenization, and then assign a string to tokenize. The method provides the ranges of the tokens in the string based on the tokenization unit. For more information, see .


// A tokenizer that segments natural language text into semantic units.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Tokenizer */

// Creates a tokenizer with the specified unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/init(unit:)
func NewTokenizerWithUnit(unit TokenUnit) Tokenizer {
	instance := getTokenizerClass().Alloc()
	rv := objc.Send[Tokenizer](instance.ID, objc.Sel("initWithUnit:"), unit)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTokenizerWithUnit */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Tokenizer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Tokenizer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Tokenizer */

// Enumerates over a given range of the string and calls the specified block for each token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/enumerateTokensInRange:usingBlock:
func (t_ Tokenizer) EnumerateTokensInRangeUsingBlock(range_ corefoundation.Range, block bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("enumerateTokensInRange:usingBlock:"), range_, block)
}/* debug [instance_methods/method]: EnumerateTokensInRangeUsingBlock */


// Sets the language of the text to be tokenized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/setLanguage(_:)
func (t_ Tokenizer) SetLanguage(language Language /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLanguage:"), language)
}/* debug [instance_methods/method]: SetLanguage */


// Finds the range of the token at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/tokenRangeAtIndex:
func (t_ Tokenizer) TokenRangeAtIndex(characterIndex uint) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("tokenRangeAtIndex:"), characterIndex)
	return rv
}/* debug [instance_methods/method]: TokenRangeAtIndex */


// Finds the entire range of all tokens contained completely or partially within the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/tokenRangeForRange:
func (t_ Tokenizer) TokenRangeForRange(range_ corefoundation.Range) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("tokenRangeForRange:"), range_)
	return rv
}/* debug [instance_methods/method]: TokenRangeForRange */


// Tokenizes the string within the provided range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/tokensForRange:
func (t_ Tokenizer) TokensForRange(range_ corefoundation.Range) []foundation.Value {
	rv := objc.Send[[]foundation.Value](t_.ID, objc.Sel("tokensForRange:"), range_)
	return rv
}/* debug [instance_methods/method]: TokensForRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Tokenizer */

// The text to be tokenized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/string
func (t_ Tokenizer) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */


// The text to be tokenized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/string
func (t_ Tokenizer) SetString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setString:"), value)
}/* debug [instance_properties/setter]: string */


// The linguistic unit that this tokenizer uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLTokenizer/unit
func (t_ Tokenizer) Unit() TokenUnit {
	rv := objc.Send[TokenUnit](t_.ID, objc.Sel("unit"))
	return rv
}/* debug [instance_properties/getter]: unit */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NLTokenizer */


