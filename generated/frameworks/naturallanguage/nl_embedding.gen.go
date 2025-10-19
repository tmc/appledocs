// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NLEmbedding] class.
var nLEmbeddingClass = _NLEmbeddingClass{objc.GetClass("NLEmbedding")}

type _NLEmbeddingClass struct {
	class objc.Class
}

// An interface definition for the [NLEmbedding] class.
type INLEmbedding interface {
	objectivec.IObject
	ContainsString(string string) bool
	DistanceBetweenStringAndStringDistanceType(firstString string, secondString string, distanceType unsafe.Pointer) unsafe.Pointer
	EnumerateNeighborsForStringMaximumCountDistanceTypeUsingBlock(string string, maxCount uint, distanceType unsafe.Pointer, block unsafe.Pointer)
	EnumerateNeighborsForStringMaximumCountMaximumDistanceDistanceTypeUsingBlock(string string, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer, block unsafe.Pointer)
	EnumerateNeighborsForVectorMaximumCountDistanceTypeUsingBlock(vector unsafe.Pointer, maxCount uint, distanceType unsafe.Pointer, block unsafe.Pointer)
	EnumerateNeighborsForVectorMaximumCountMaximumDistanceDistanceTypeUsingBlock(vector unsafe.Pointer, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer, block unsafe.Pointer)
	GetVectorForString(vector unsafe.Pointer, string string) bool
	NeighborsForStringMaximumCountDistanceType(string string, maxCount uint, distanceType unsafe.Pointer) unsafe.Pointer
	NeighborsForStringMaximumCountMaximumDistanceDistanceType(string string, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer) unsafe.Pointer
	NeighborsForVectorMaximumCountDistanceType(vector unsafe.Pointer, maxCount uint, distanceType unsafe.Pointer) unsafe.Pointer
	NeighborsForVectorMaximumCountMaximumDistanceDistanceType(vector unsafe.Pointer, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer) unsafe.Pointer
	VectorForString(string string) unsafe.Pointer
}

// A map of strings to vectors, which locates neighboring, similar strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding

type NLEmbedding struct {
	objectivec.Object
}

// NLEmbeddingFrom constructs a [NLEmbedding] from an unsafe.Pointer.
//
// A map of strings to vectors, which locates neighboring, similar strings.
func NLEmbeddingFrom(ptr unsafe.Pointer) NLEmbedding {
	return NLEmbedding{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NLEmbeddingClass) Alloc() NLEmbedding {
	rv := objc.Send[NLEmbedding](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NLEmbeddingClass) New() NLEmbedding {
	rv := objc.Send[NLEmbedding](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NLEmbedding) Init() NLEmbedding {
	rv := objc.Send[NLEmbedding](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NLEmbedding) Autorelease() NLEmbedding {
	rv := objc.Send[NLEmbedding](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLEmbedding creates a new NLEmbedding instance.
func NewNLEmbedding() NLEmbedding {
	return nLEmbeddingClass.New()
}


// Creates a word embedding from a model file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/init(contentsOf:)
func NewEmbeddingWithContentsOfURLError(url unsafe.Pointer, error unsafe.Pointer) NLEmbedding {
	rv := objc.Send[NLEmbedding](objc.ID(nLEmbeddingClass.class), objc.Sel("embeddingWithContentsOfURL:error:"), url, error)
	rv.Autorelease()
	return rv
}


// Retrieves the current version of a word embedding for the given language. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/currentRevision(for:)
func (nc _NLEmbeddingClass) CurrentRevisionForLanguage(language unsafe.Pointer) uint {
	rv := objc.Send[uint](objc.ID(nc.class), objc.Sel("currentRevisionForLanguage:"), language)
	return rv
}
// Retrieves the current version of a sentence embedding for the given language. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/currentSentenceEmbeddingRevision(for:)
func (nc _NLEmbeddingClass) CurrentSentenceEmbeddingRevisionForLanguage(language unsafe.Pointer) uint {
	rv := objc.Send[uint](objc.ID(nc.class), objc.Sel("currentSentenceEmbeddingRevisionForLanguage:"), language)
	return rv
}
// Creates a word embedding from a model file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/init(contentsOf:)
func (nc _NLEmbeddingClass) EmbeddingWithContentsOfURLError(url unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("embeddingWithContentsOfURL:error:"), url, error)
	return rv
}
// Retrieves a sentence embedding for a given language. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/sentenceEmbedding(for:)
func (nc _NLEmbeddingClass) SentenceEmbeddingForLanguage(language unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("sentenceEmbeddingForLanguage:"), language)
	return rv
}
// Retrieves a sentence embedding for a given language and revision. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/sentenceEmbedding(for:revision:)
func (nc _NLEmbeddingClass) SentenceEmbeddingForLanguageRevision(language unsafe.Pointer, revision uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("sentenceEmbeddingForLanguage:revision:"), language, revision)
	return rv
}
// Retrieves all version numbers of a word embedding for the given language. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/supportedRevisions(for:)
func (nc _NLEmbeddingClass) SupportedRevisionsForLanguage(language unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("supportedRevisionsForLanguage:"), language)
	return rv
}
// Retrieves all version numbers of a sentence embedding for the given language. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/supportedSentenceEmbeddingRevisions(for:)
func (nc _NLEmbeddingClass) SupportedSentenceEmbeddingRevisionsForLanguage(language unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("supportedSentenceEmbeddingRevisionsForLanguage:"), language)
	return rv
}
// Retrieves a word embedding for a given language. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/wordEmbedding(for:)
func (nc _NLEmbeddingClass) WordEmbeddingForLanguage(language unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("wordEmbeddingForLanguage:"), language)
	return rv
}
// Retrieves a word embedding for a given language and revision. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/wordEmbedding(for:revision:)
func (nc _NLEmbeddingClass) WordEmbeddingForLanguageRevision(language unsafe.Pointer, revision uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("wordEmbeddingForLanguage:revision:"), language, revision)
	return rv
}
// Exports the word embedding contained within a Core ML model file at the given URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/writeEmbeddingForDictionary:language:revision:toURL:error:
func (nc _NLEmbeddingClass) WriteEmbeddingForDictionaryLanguageRevisionToURLError(dictionary unsafe.Pointer, language unsafe.Pointer, revision uint, url unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(nc.class), objc.Sel("writeEmbeddingForDictionary:language:revision:toURL:error:"), dictionary, language, revision, url, error)
	return rv
}
// Requests a Boolean value that indicates whether the term is in the vocabulary. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/contains(_:)
func (n_ NLEmbedding) ContainsString(string string) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("containsString:"), string)
	return rv
}
// Calculates the distance between two strings in the vocabulary space. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/distanceBetweenString:andString:distanceType:
func (n_ NLEmbedding) DistanceBetweenStringAndStringDistanceType(firstString string, secondString string, distanceType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("distanceBetweenString:andString:distanceType:"), firstString, secondString, distanceType)
	return rv
}
// Passes the nearest strings of a string in the vocabulary to a block. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/enumerateNeighborsForString:maximumCount:distanceType:usingBlock:
func (n_ NLEmbedding) EnumerateNeighborsForStringMaximumCountDistanceTypeUsingBlock(string string, maxCount uint, distanceType unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("enumerateNeighborsForString:maximumCount:distanceType:usingBlock:"), string, maxCount, distanceType, block)
}
// Passes the nearest strings, within a radius of a string in the vocabulary, to a block. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/enumerateNeighborsForString:maximumCount:maximumDistance:distanceType:usingBlock:
func (n_ NLEmbedding) EnumerateNeighborsForStringMaximumCountMaximumDistanceDistanceTypeUsingBlock(string string, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("enumerateNeighborsForString:maximumCount:maximumDistance:distanceType:usingBlock:"), string, maxCount, maxDistance, distanceType, block)
}
// Passes the nearest strings of a location in the vocabulary space to a closure. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/enumerateNeighborsForVector:maximumCount:distanceType:usingBlock:
func (n_ NLEmbedding) EnumerateNeighborsForVectorMaximumCountDistanceTypeUsingBlock(vector unsafe.Pointer, maxCount uint, distanceType unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("enumerateNeighborsForVector:maximumCount:distanceType:usingBlock:"), vector, maxCount, distanceType, block)
}
// Passes the nearest strings, within a radius of a location in the vocabulary space, to a block. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/enumerateNeighborsForVector:maximumCount:maximumDistance:distanceType:usingBlock:
func (n_ NLEmbedding) EnumerateNeighborsForVectorMaximumCountMaximumDistanceDistanceTypeUsingBlock(vector unsafe.Pointer, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("enumerateNeighborsForVector:maximumCount:maximumDistance:distanceType:usingBlock:"), vector, maxCount, maxDistance, distanceType, block)
}
// Copies a vector into the given a pointer to a float array. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/getVector:forString:
func (n_ NLEmbedding) GetVectorForString(vector unsafe.Pointer, string string) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("getVector:forString:"), vector, string)
	return rv
}
// Retrieves a limited number of strings near a string in the vocabulary. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForString:maximumCount:distanceType:
func (n_ NLEmbedding) NeighborsForStringMaximumCountDistanceType(string string, maxCount uint, distanceType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("neighborsForString:maximumCount:distanceType:"), string, maxCount, distanceType)
	return rv
}
// Retrieves a limited number of strings, within a radius of a string, in the vocabulary. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForString:maximumCount:maximumDistance:distanceType:
func (n_ NLEmbedding) NeighborsForStringMaximumCountMaximumDistanceDistanceType(string string, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("neighborsForString:maximumCount:maximumDistance:distanceType:"), string, maxCount, maxDistance, distanceType)
	return rv
}
// Retrieves a limited number of strings near a location in the vocabulary space. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForVector:maximumCount:distanceType:
func (n_ NLEmbedding) NeighborsForVectorMaximumCountDistanceType(vector unsafe.Pointer, maxCount uint, distanceType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("neighborsForVector:maximumCount:distanceType:"), vector, maxCount, distanceType)
	return rv
}
// Retrieves a limited number of strings within a radius of a location in the vocabulary space. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForVector:maximumCount:maximumDistance:distanceType:
func (n_ NLEmbedding) NeighborsForVectorMaximumCountMaximumDistanceDistanceType(vector unsafe.Pointer, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("neighborsForVector:maximumCount:maximumDistance:distanceType:"), vector, maxCount, maxDistance, distanceType)
	return rv
}
// Requests the vector for the given term. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/vectorForString:
func (n_ NLEmbedding) VectorForString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("vectorForString:"), string)
	return rv
}

