// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NLContextualEmbeddingResult] class.
var nLContextualEmbeddingResultClass = _NLContextualEmbeddingResultClass{objc.GetClass("NLContextualEmbeddingResult")}

type _NLContextualEmbeddingResultClass struct {
	class objc.Class
}

// An interface definition for the [NLContextualEmbeddingResult] class.
type INLContextualEmbeddingResult interface {
	objectivec.IObject
}

// An object that represents the embedding vector result from applying a contextual embedding to a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbeddingResult

type NLContextualEmbeddingResult struct {
	objectivec.Object
}

// NLContextualEmbeddingResultFrom constructs a [NLContextualEmbeddingResult] from an unsafe.Pointer.
//
// An object that represents the embedding vector result from applying a contextual embedding to a string.
func NLContextualEmbeddingResultFrom(ptr unsafe.Pointer) NLContextualEmbeddingResult {
	return NLContextualEmbeddingResult{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NLContextualEmbeddingResultClass) Alloc() NLContextualEmbeddingResult {
	rv := objc.Send[NLContextualEmbeddingResult](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NLContextualEmbeddingResultClass) New() NLContextualEmbeddingResult {
	rv := objc.Send[NLContextualEmbeddingResult](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NLContextualEmbeddingResult) Init() NLContextualEmbeddingResult {
	rv := objc.Send[NLContextualEmbeddingResult](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NLContextualEmbeddingResult) Autorelease() NLContextualEmbeddingResult {
	rv := objc.Send[NLContextualEmbeddingResult](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLContextualEmbeddingResult creates a new NLContextualEmbeddingResult instance.
func NewNLContextualEmbeddingResult() NLContextualEmbeddingResult {
	return nLContextualEmbeddingResultClass.New()
}




