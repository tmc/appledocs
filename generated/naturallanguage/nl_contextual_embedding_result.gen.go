// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ContextualEmbeddingResult] class.
type IContextualEmbeddingResult interface {
	objectivec.IObject
	EnumerateTokenVectorsInRangeUsingBlock(range_ Range, block unsafe.Pointer)
	TokenVectorAtIndexTokenRange(characterIndex uint, tokenRange unsafe.Pointer) []accessibility.NSNumber
}

// An object that represents the embedding vector result from applying a contextual embedding to a string.
//
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

// Alloc allocates a new instance without initialization.
func (cc _ContextualEmbeddingResultClass) Alloc() ContextualEmbeddingResult {
	rv := objc.Send[ContextualEmbeddingResult](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Iterates over the embedding vectors for the range you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/enumerateTokenVectorsInRange:usingBlock:
func (c_ ContextualEmbeddingResult) EnumerateTokenVectorsInRangeUsingBlock(range_ Range, block unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("enumerateTokenVectorsInRange:usingBlock:"), range_, block)
}

// Gets a token vector at the index you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/tokenVectorAtIndex:tokenRange:
func (c_ ContextualEmbeddingResult) TokenVectorAtIndexTokenRange(characterIndex uint, tokenRange unsafe.Pointer) []accessibility.NSNumber {
	rv := objc.Send[[]accessibility.NSNumber](c_.ID, objc.Sel("tokenVectorAtIndex:tokenRange:"), characterIndex, tokenRange)
	return rv
}

// The resulting language.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/language
func (c_ ContextualEmbeddingResult) Language() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("language"))
	return rv
}

// The number of embedding vectors the request generates.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/sequenceLength
func (c_ ContextualEmbeddingResult) SequenceLength() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("sequenceLength"))
	return rv
}

// The string value.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult/string
func (c_ ContextualEmbeddingResult) String() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("string"))
	return rv
}



