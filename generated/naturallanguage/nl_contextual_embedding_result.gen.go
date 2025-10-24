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

/* debug [class.gen.go]: Generating class NLContextualEmbeddingResult */


/* debug [class_header]: Header for NLContextualEmbeddingResult */
// The class instance for the [ContextualEmbeddingResult] class.
var (
	ContextualEmbeddingResultClass     _ContextualEmbeddingResultClass
	ContextualEmbeddingResultClassOnce sync.Once
)

func getContextualEmbeddingResultClass() _ContextualEmbeddingResultClass {
	ContextualEmbeddingResultClassOnce.Do(func() {
		ContextualEmbeddingResultClass = _ContextualEmbeddingResultClass{objc.GetClass("NLContextualEmbeddingResult")}
	})
	return ContextualEmbeddingResultClass
}

type _ContextualEmbeddingResultClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ContextualEmbeddingResult */
// An interface definition for the [ContextualEmbeddingResult] class.
type IContextualEmbeddingResult interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ContextualEmbeddingResult */
	// properties:
	Language() Language /* typedef */
	SequenceLength() uint
	String() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ContextualEmbeddingResult */
	// methods:
	EnumerateTokenVectorsInRangeUsingBlock(range_ corefoundation.Range, block bool)
	TokenVectorAtIndexTokenRange(characterIndex uint, tokenRange RangePointer /* not a class type */) []foundation.Number
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ContextualEmbeddingResult */
// Alloc allocates a new instance without initialization.
func (cc _ContextualEmbeddingResultClass) Alloc() ContextualEmbeddingResult {
	rv := objc.Send[ContextualEmbeddingResult](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContextualEmbeddingResultClass) New() ContextualEmbeddingResult {
	rv := objc.Send[ContextualEmbeddingResult](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContextualEmbeddingResult) Init() ContextualEmbeddingResult {
	rv := objc.Send[ContextualEmbeddingResult](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContextualEmbeddingResult) Autorelease() ContextualEmbeddingResult {
	rv := objc.Send[ContextualEmbeddingResult](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContextualEmbeddingResult creates a new ContextualEmbeddingResult instance.
func NewContextualEmbeddingResult() ContextualEmbeddingResult {
	return getContextualEmbeddingResultClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ContextualEmbeddingResult */
// An object that represents the embedding vector result from applying a contextual embedding to a string.


// An object that represents the embedding vector result from applying a contextual embedding to a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult
type ContextualEmbeddingResult struct {
	objectivec.Object
}

// ContextualEmbeddingResultFrom constructs a [ContextualEmbeddingResult] from an unsafe.Pointer.
//
// An object that represents the embedding vector result from applying a contextual embedding to a string.
func ContextualEmbeddingResultFrom(ptr unsafe.Pointer) ContextualEmbeddingResult {
	return ContextualEmbeddingResult{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ContextualEmbeddingResult *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ContextualEmbeddingResult */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ContextualEmbeddingResult */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ContextualEmbeddingResult */

// Iterates over the embedding vectors for the range you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/enumerateTokenVectorsInRange:usingBlock:
func (c_ ContextualEmbeddingResult) EnumerateTokenVectorsInRangeUsingBlock(range_ corefoundation.Range, block bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("enumerateTokenVectorsInRange:usingBlock:"), range_, block)
}/* debug [instance_methods/method]: EnumerateTokenVectorsInRangeUsingBlock */


// Gets a token vector at the index you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/tokenVectorAtIndex:tokenRange:
func (c_ ContextualEmbeddingResult) TokenVectorAtIndexTokenRange(characterIndex uint, tokenRange RangePointer /* not a class type */) []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("tokenVectorAtIndex:tokenRange:"), characterIndex, tokenRange)
	return rv
}/* debug [instance_methods/method]: TokenVectorAtIndexTokenRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ContextualEmbeddingResult */

// The resulting language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/language
func (c_ ContextualEmbeddingResult) Language() Language /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("language"))
	return rv
}/* debug [instance_properties/getter]: language */


// The number of embedding vectors the request generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/sequenceLength
func (c_ ContextualEmbeddingResult) SequenceLength() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("sequenceLength"))
	return rv
}/* debug [instance_properties/getter]: sequenceLength */


// The string value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/string
func (c_ ContextualEmbeddingResult) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NLContextualEmbeddingResult */



