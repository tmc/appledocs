// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Embedding] class.
var (
	EmbeddingClass     _EmbeddingClass
	EmbeddingClassOnce sync.Once
)

func getEmbeddingClass() _EmbeddingClass {
	EmbeddingClassOnce.Do(func() {
		EmbeddingClass = _EmbeddingClass{objc.GetClass("NLEmbedding")}
	})
	return EmbeddingClass
}

type _EmbeddingClass struct {
	class objc.Class
}

// An interface definition for the [Embedding] class.
type IEmbedding interface {
	objectivec.IObject
	ContainsString(string_ string) bool
	DistanceBetweenStringAndStringDistanceType(firstString string, secondString string, distanceType unsafe.Pointer) unsafe.Pointer
	EnumerateNeighborsForStringMaximumCountDistanceTypeUsingBlock(string_ string, maxCount uint, distanceType unsafe.Pointer, block unsafe.Pointer)
	EnumerateNeighborsForStringMaximumCountMaximumDistanceDistanceTypeUsingBlock(string_ string, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer, block unsafe.Pointer)
	EnumerateNeighborsForVectorMaximumCountDistanceTypeUsingBlock(vector unsafe.Pointer, maxCount uint, distanceType unsafe.Pointer, block unsafe.Pointer)
	EnumerateNeighborsForVectorMaximumCountMaximumDistanceDistanceTypeUsingBlock(vector unsafe.Pointer, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer, block unsafe.Pointer)
	GetVectorForString(vector unsafe.Pointer, string_ string) bool
	NeighborsForStringMaximumCountDistanceType(string_ string, maxCount uint, distanceType unsafe.Pointer) []string
	NeighborsForStringMaximumCountMaximumDistanceDistanceType(string_ string, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer) []string
	NeighborsForVectorMaximumCountDistanceType(vector unsafe.Pointer, maxCount uint, distanceType unsafe.Pointer) []string
	NeighborsForVectorMaximumCountMaximumDistanceDistanceType(vector unsafe.Pointer, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer) []string
	VectorForString(string_ string) []accessibility.NSNumber
}

// A map of strings to vectors, which locates neighboring, similar strings.
//
// Use an to find similar strings based on the proximity of their vectors. The is the entire set of strings in an embedding. Each string in the vocabulary has a vector, which is an array of doubles, and each double corresponds to a dimension in the embedding. An uses these vectors to determine the distance between two strings, or to find the nearest neighbors of a string in the vocabulary. The higher the similarity of any two strings, the smaller the distance is between them. provides built-in word embeddings that you can retrieve by using the method. You can also compile your own custom embedding into an efficient, searchable, on-disk representation. Typically, you compile an embedding by using Create ML’s and save it as a file for your Xcode project at development time. Alternatively, you can compile an embedding at runtime by using Natural Language’s method. Your custom embedding can use any kind of string that’s useful to your app, such as phrases, brand names, serial numbers, and so on. For example, you could make an embedding of movie titles. Each movie title could have a vector that places similar movies close together in the embedding.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding
type Embedding struct {
	objectivec.Object
}

// EmbeddingFrom constructs a [Embedding] from an unsafe.Pointer.
//
// A map of strings to vectors, which locates neighboring, similar strings.
func EmbeddingFrom(ptr unsafe.Pointer) Embedding {
	return Embedding{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EmbeddingClass) Alloc() Embedding {
	rv := objc.Send[Embedding](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EmbeddingClass) New() Embedding {
	rv := objc.Send[Embedding](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ Embedding) Init() Embedding {
	rv := objc.Send[Embedding](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ Embedding) Autorelease() Embedding {
	rv := objc.Send[Embedding](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEmbedding creates a new Embedding instance.
func NewEmbedding() Embedding {
	return getEmbeddingClass().New()
}


// Creates a word embedding from a model file.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/init(contentsOf:)
func NewEmbeddingWithContentsOfURLError(url unsafe.Pointer, error_ unsafe.Pointer) Embedding {
	rv := objc.Send[Embedding](objc.ID(getEmbeddingClass().class), objc.Sel("embeddingWithContentsOfURL:error:"), url, error_)
	return rv
}


// Retrieves the current version of a word embedding for the given language.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/currentRevision(for:)
func (ec _EmbeddingClass) CurrentRevisionForLanguage(language unsafe.Pointer) uint {
	rv := objc.Send[uint](objc.ID(ec.class), objc.Sel("currentRevisionForLanguage:"), language)
	return rv
}

// Retrieves the current version of a sentence embedding for the given language.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/currentSentenceEmbeddingRevision(for:)
func (ec _EmbeddingClass) CurrentSentenceEmbeddingRevisionForLanguage(language unsafe.Pointer) uint {
	rv := objc.Send[uint](objc.ID(ec.class), objc.Sel("currentSentenceEmbeddingRevisionForLanguage:"), language)
	return rv
}

// Creates a word embedding from a model file.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/init(contentsOf:)
func (ec _EmbeddingClass) EmbeddingWithContentsOfURLError(url unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("embeddingWithContentsOfURL:error:"), url, error_)
	return rv
}

// Retrieves a sentence embedding for a given language.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/sentenceEmbedding(for:)
func (ec _EmbeddingClass) SentenceEmbeddingForLanguage(language unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("sentenceEmbeddingForLanguage:"), language)
	return rv
}

// Retrieves a sentence embedding for a given language and revision.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/sentenceEmbedding(for:revision:)
func (ec _EmbeddingClass) SentenceEmbeddingForLanguageRevision(language unsafe.Pointer, revision uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("sentenceEmbeddingForLanguage:revision:"), language, revision)
	return rv
}

// Retrieves all version numbers of a word embedding for the given language.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/supportedRevisions(for:)
func (ec _EmbeddingClass) SupportedRevisionsForLanguage(language unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("supportedRevisionsForLanguage:"), language)
	return rv
}

// Retrieves all version numbers of a sentence embedding for the given language.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/supportedSentenceEmbeddingRevisions(for:)
func (ec _EmbeddingClass) SupportedSentenceEmbeddingRevisionsForLanguage(language unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("supportedSentenceEmbeddingRevisionsForLanguage:"), language)
	return rv
}

// Retrieves a word embedding for a given language.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/wordEmbedding(for:)
func (ec _EmbeddingClass) WordEmbeddingForLanguage(language unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("wordEmbeddingForLanguage:"), language)
	return rv
}

// Retrieves a word embedding for a given language and revision.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/wordEmbedding(for:revision:)
func (ec _EmbeddingClass) WordEmbeddingForLanguageRevision(language unsafe.Pointer, revision uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("wordEmbeddingForLanguage:revision:"), language, revision)
	return rv
}

// Exports the word embedding contained within a Core ML model file at the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/writeEmbeddingForDictionary:language:revision:toURL:error:
func (ec _EmbeddingClass) WriteEmbeddingForDictionaryLanguageRevisionToURLError(dictionary unsafe.Pointer, language unsafe.Pointer, revision uint, url unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(ec.class), objc.Sel("writeEmbeddingForDictionary:language:revision:toURL:error:"), dictionary, language, revision, url, error_)
	return rv
}

// Requests a Boolean value that indicates whether the term is in the vocabulary.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/contains(_:)
func (e_ Embedding) ContainsString(string_ string) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("containsString:"), objc.String(string_))
	return rv
}

// Calculates the distance between two strings in the vocabulary space.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/distanceBetweenString:andString:distanceType:
func (e_ Embedding) DistanceBetweenStringAndStringDistanceType(firstString string, secondString string, distanceType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("distanceBetweenString:andString:distanceType:"), objc.String(firstString), objc.String(secondString), distanceType)
	return rv
}

// Passes the nearest strings of a string in the vocabulary to a block.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/enumerateNeighborsForString:maximumCount:distanceType:usingBlock:
func (e_ Embedding) EnumerateNeighborsForStringMaximumCountDistanceTypeUsingBlock(string_ string, maxCount uint, distanceType unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("enumerateNeighborsForString:maximumCount:distanceType:usingBlock:"), objc.String(string_), maxCount, distanceType, block)
}

// Passes the nearest strings, within a radius of a string in the vocabulary, to a block.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/enumerateNeighborsForString:maximumCount:maximumDistance:distanceType:usingBlock:
func (e_ Embedding) EnumerateNeighborsForStringMaximumCountMaximumDistanceDistanceTypeUsingBlock(string_ string, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("enumerateNeighborsForString:maximumCount:maximumDistance:distanceType:usingBlock:"), objc.String(string_), maxCount, maxDistance, distanceType, block)
}

// Passes the nearest strings of a location in the vocabulary space to a closure.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/enumerateNeighborsForVector:maximumCount:distanceType:usingBlock:
func (e_ Embedding) EnumerateNeighborsForVectorMaximumCountDistanceTypeUsingBlock(vector unsafe.Pointer, maxCount uint, distanceType unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("enumerateNeighborsForVector:maximumCount:distanceType:usingBlock:"), vector, maxCount, distanceType, block)
}

// Passes the nearest strings, within a radius of a location in the vocabulary space, to a block.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/enumerateNeighborsForVector:maximumCount:maximumDistance:distanceType:usingBlock:
func (e_ Embedding) EnumerateNeighborsForVectorMaximumCountMaximumDistanceDistanceTypeUsingBlock(vector unsafe.Pointer, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("enumerateNeighborsForVector:maximumCount:maximumDistance:distanceType:usingBlock:"), vector, maxCount, maxDistance, distanceType, block)
}

// Copies a vector into the given a pointer to a float array.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/getVector:forString:
func (e_ Embedding) GetVectorForString(vector unsafe.Pointer, string_ string) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("getVector:forString:"), vector, objc.String(string_))
	return rv
}

// Retrieves a limited number of strings near a string in the vocabulary.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForString:maximumCount:distanceType:
func (e_ Embedding) NeighborsForStringMaximumCountDistanceType(string_ string, maxCount uint, distanceType unsafe.Pointer) []string {
	rv := objc.Send[[]string](e_.ID, objc.Sel("neighborsForString:maximumCount:distanceType:"), objc.String(string_), maxCount, distanceType)
	return rv
}

// Retrieves a limited number of strings, within a radius of a string, in the vocabulary.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForString:maximumCount:maximumDistance:distanceType:
func (e_ Embedding) NeighborsForStringMaximumCountMaximumDistanceDistanceType(string_ string, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer) []string {
	rv := objc.Send[[]string](e_.ID, objc.Sel("neighborsForString:maximumCount:maximumDistance:distanceType:"), objc.String(string_), maxCount, maxDistance, distanceType)
	return rv
}

// Retrieves a limited number of strings near a location in the vocabulary space.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForVector:maximumCount:distanceType:
func (e_ Embedding) NeighborsForVectorMaximumCountDistanceType(vector unsafe.Pointer, maxCount uint, distanceType unsafe.Pointer) []string {
	rv := objc.Send[[]string](e_.ID, objc.Sel("neighborsForVector:maximumCount:distanceType:"), vector, maxCount, distanceType)
	return rv
}

// Retrieves a limited number of strings within a radius of a location in the vocabulary space.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForVector:maximumCount:maximumDistance:distanceType:
func (e_ Embedding) NeighborsForVectorMaximumCountMaximumDistanceDistanceType(vector unsafe.Pointer, maxCount uint, maxDistance unsafe.Pointer, distanceType unsafe.Pointer) []string {
	rv := objc.Send[[]string](e_.ID, objc.Sel("neighborsForVector:maximumCount:maximumDistance:distanceType:"), vector, maxCount, maxDistance, distanceType)
	return rv
}

// Requests the vector for the given term.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/vectorForString:
func (e_ Embedding) VectorForString(string_ string) []accessibility.NSNumber {
	rv := objc.Send[[]accessibility.NSNumber](e_.ID, objc.Sel("vectorForString:"), objc.String(string_))
	return rv
}

// The number of dimensions in the vocabulary’s vector space.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/dimension
func (e_ Embedding) Dimension() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("dimension"))
	return rv
}

// The language of the text in the word embedding.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/language
func (e_ Embedding) Language() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("language"))
	return rv
}

// The revision of the word embedding.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/revision
func (e_ Embedding) Revision() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("revision"))
	return rv
}

// The number of words in the vocabulary.
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/vocabularySize
func (e_ Embedding) VocabularySize() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("vocabularySize"))
	return rv
}


