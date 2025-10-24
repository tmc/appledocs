// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NLEmbedding */


/* debug [class_header]: Header for NLEmbedding */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Embedding */
// An interface definition for the [Embedding] class.
type IEmbedding interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Embedding */
	// properties:
	Dimension() uint
	Language() Language /* typedef */
	Revision() uint
	VocabularySize() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Embedding */
	// methods:
	ContainsString(string_ objc.IObject /* cross-framework: NSString */) bool
	DistanceBetweenStringAndStringDistanceType(firstString objc.IObject /* cross-framework: NSString */, secondString objc.IObject /* cross-framework: NSString */, distanceType DistanceType) Distance /* typedef */
	EnumerateNeighborsForStringMaximumCountDistanceTypeUsingBlock(string_ objc.IObject /* cross-framework: NSString */, maxCount uint, distanceType DistanceType, block bool)
	EnumerateNeighborsForStringMaximumCountMaximumDistanceDistanceTypeUsingBlock(string_ objc.IObject /* cross-framework: NSString */, maxCount uint, maxDistance Distance /* typedef */, distanceType DistanceType, block bool)
	EnumerateNeighborsForVectorMaximumCountDistanceTypeUsingBlock(vector []foundation.Number, maxCount uint, distanceType DistanceType, block bool)
	EnumerateNeighborsForVectorMaximumCountMaximumDistanceDistanceTypeUsingBlock(vector []foundation.Number, maxCount uint, maxDistance Distance /* typedef */, distanceType DistanceType, block bool)
	GetVectorForString(vector objectivec.IObject, string_ objc.IObject /* cross-framework: NSString */) bool
	NeighborsForStringMaximumCountDistanceType(string_ objc.IObject /* cross-framework: NSString */, maxCount uint, distanceType DistanceType) []string
	NeighborsForStringMaximumCountMaximumDistanceDistanceType(string_ objc.IObject /* cross-framework: NSString */, maxCount uint, maxDistance Distance /* typedef */, distanceType DistanceType) []string
	NeighborsForVectorMaximumCountDistanceType(vector []foundation.Number, maxCount uint, distanceType DistanceType) []string
	NeighborsForVectorMaximumCountMaximumDistanceDistanceType(vector []foundation.Number, maxCount uint, maxDistance Distance /* typedef */, distanceType DistanceType) []string
	VectorForString(string_ objc.IObject /* cross-framework: NSString */) []foundation.Number
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Embedding */
// Alloc allocates a new instance without initialization.
func (ec _EmbeddingClass) Alloc() Embedding {
	rv := objc.Send[Embedding](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Embedding */
// A map of strings to vectors, which locates neighboring, similar strings.
//
// Use an to find similar strings based on the proximity of their vectors. The is the entire set of strings in an embedding. Each string in the vocabulary has a vector, which is an array of doubles, and each double corresponds to a dimension in the embedding. An uses these vectors to determine the distance between two strings, or to find the nearest neighbors of a string in the vocabulary. The higher the similarity of any two strings, the smaller the distance is between them. provides built-in word embeddings that you can retrieve by using the method. You can also compile your own custom embedding into an efficient, searchable, on-disk representation. Typically, you compile an embedding by using Create ML’s and save it as a file for your Xcode project at development time. Alternatively, you can compile an embedding at runtime by using Natural Language’s method. Your custom embedding can use any kind of string that’s useful to your app, such as phrases, brand names, serial numbers, and so on. For example, you could make an embedding of movie titles. Each movie title could have a vector that places similar movies close together in the embedding.


// A map of strings to vectors, which locates neighboring, similar strings.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Embedding */

// Creates a word embedding from a model file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/init(contentsOf:)
func NewEmbeddingWithContentsOfURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) Embedding {
	rv := objc.Send[Embedding](objc.ID(getEmbeddingClass().class), objc.Sel("embeddingWithContentsOfURL:error:"), url, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewEmbeddingWithContentsOfURLError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Embedding */

// Retrieves the current version of a word embedding for the given language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/currentRevision(for:)
func (ec _EmbeddingClass) CurrentRevisionForLanguage(language Language /* typedef */) uint {
	rv := objc.Send[uint](objc.ID(ec.class), objc.Sel("currentRevisionForLanguage:"), language)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CurrentRevisionForLanguage) */


// Retrieves the current version of a sentence embedding for the given language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/currentSentenceEmbeddingRevision(for:)
func (ec _EmbeddingClass) CurrentSentenceEmbeddingRevisionForLanguage(language Language /* typedef */) uint {
	rv := objc.Send[uint](objc.ID(ec.class), objc.Sel("currentSentenceEmbeddingRevisionForLanguage:"), language)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CurrentSentenceEmbeddingRevisionForLanguage) */


// Creates a word embedding from a model file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/init(contentsOf:)
func (ec _EmbeddingClass) EmbeddingWithContentsOfURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ec.class), objc.Sel("embeddingWithContentsOfURL:error:"), url, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EmbeddingWithContentsOfURLError) */


// Retrieves a sentence embedding for a given language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/sentenceEmbedding(for:)
func (ec _EmbeddingClass) SentenceEmbeddingForLanguage(language Language /* typedef */) IEmbedding {
	rv := objc.Send[Embedding](objc.ID(ec.class), objc.Sel("sentenceEmbeddingForLanguage:"), language)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SentenceEmbeddingForLanguage) */


// Retrieves a sentence embedding for a given language and revision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/sentenceEmbedding(for:revision:)
func (ec _EmbeddingClass) SentenceEmbeddingForLanguageRevision(language Language /* typedef */, revision uint) IEmbedding {
	rv := objc.Send[Embedding](objc.ID(ec.class), objc.Sel("sentenceEmbeddingForLanguage:revision:"), language, revision)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SentenceEmbeddingForLanguageRevision) */


// Retrieves all version numbers of a word embedding for the given language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/supportedRevisions(for:)
func (ec _EmbeddingClass) SupportedRevisionsForLanguage(language Language /* typedef */) foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](objc.ID(ec.class), objc.Sel("supportedRevisionsForLanguage:"), language)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedRevisionsForLanguage) */


// Retrieves all version numbers of a sentence embedding for the given language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/supportedSentenceEmbeddingRevisions(for:)
func (ec _EmbeddingClass) SupportedSentenceEmbeddingRevisionsForLanguage(language Language /* typedef */) foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](objc.ID(ec.class), objc.Sel("supportedSentenceEmbeddingRevisionsForLanguage:"), language)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedSentenceEmbeddingRevisionsForLanguage) */


// Retrieves a word embedding for a given language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/wordEmbedding(for:)
func (ec _EmbeddingClass) WordEmbeddingForLanguage(language Language /* typedef */) IEmbedding {
	rv := objc.Send[Embedding](objc.ID(ec.class), objc.Sel("wordEmbeddingForLanguage:"), language)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WordEmbeddingForLanguage) */


// Retrieves a word embedding for a given language and revision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/wordEmbedding(for:revision:)
func (ec _EmbeddingClass) WordEmbeddingForLanguageRevision(language Language /* typedef */, revision uint) IEmbedding {
	rv := objc.Send[Embedding](objc.ID(ec.class), objc.Sel("wordEmbeddingForLanguage:revision:"), language, revision)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WordEmbeddingForLanguageRevision) */


// Exports the word embedding contained within a Core ML model file at the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/writeEmbeddingForDictionary:language:revision:toURL:error:
func (ec _EmbeddingClass) WriteEmbeddingForDictionaryLanguageRevisionToURLError(dictionary foundation.IDictionary, language Language /* typedef */, revision uint, url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(ec.class), objc.Sel("writeEmbeddingForDictionary:language:revision:toURL:error:"), dictionary, language, revision, url, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WriteEmbeddingForDictionaryLanguageRevisionToURLError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Embedding */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Embedding */

// Requests a Boolean value that indicates whether the term is in the vocabulary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/contains(_:)
func (e_ Embedding) ContainsString(string_ objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("containsString:"), string_)
	return rv
}/* debug [instance_methods/method]: ContainsString */


// Calculates the distance between two strings in the vocabulary space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/distanceBetweenString:andString:distanceType:
func (e_ Embedding) DistanceBetweenStringAndStringDistanceType(firstString objc.IObject /* cross-framework: NSString */, secondString objc.IObject /* cross-framework: NSString */, distanceType DistanceType) Distance /* typedef */ {
	rv := objc.Send[float64](e_.ID, objc.Sel("distanceBetweenString:andString:distanceType:"), firstString, secondString, distanceType)
	return rv
}/* debug [instance_methods/method]: DistanceBetweenStringAndStringDistanceType */


// Passes the nearest strings of a string in the vocabulary to a block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/enumerateNeighborsForString:maximumCount:distanceType:usingBlock:
func (e_ Embedding) EnumerateNeighborsForStringMaximumCountDistanceTypeUsingBlock(string_ objc.IObject /* cross-framework: NSString */, maxCount uint, distanceType DistanceType, block bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("enumerateNeighborsForString:maximumCount:distanceType:usingBlock:"), string_, maxCount, distanceType, block)
}/* debug [instance_methods/method]: EnumerateNeighborsForStringMaximumCountDistanceTypeUsingBlock */


// Passes the nearest strings, within a radius of a string in the vocabulary, to a block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/enumerateNeighborsForString:maximumCount:maximumDistance:distanceType:usingBlock:
func (e_ Embedding) EnumerateNeighborsForStringMaximumCountMaximumDistanceDistanceTypeUsingBlock(string_ objc.IObject /* cross-framework: NSString */, maxCount uint, maxDistance Distance /* typedef */, distanceType DistanceType, block bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("enumerateNeighborsForString:maximumCount:maximumDistance:distanceType:usingBlock:"), string_, maxCount, maxDistance, distanceType, block)
}/* debug [instance_methods/method]: EnumerateNeighborsForStringMaximumCountMaximumDistanceDistanceTypeUsingBlock */


// Passes the nearest strings of a location in the vocabulary space to a closure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/enumerateNeighborsForVector:maximumCount:distanceType:usingBlock:
func (e_ Embedding) EnumerateNeighborsForVectorMaximumCountDistanceTypeUsingBlock(vector []foundation.Number, maxCount uint, distanceType DistanceType, block bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("enumerateNeighborsForVector:maximumCount:distanceType:usingBlock:"), vector, maxCount, distanceType, block)
}/* debug [instance_methods/method]: EnumerateNeighborsForVectorMaximumCountDistanceTypeUsingBlock */


// Passes the nearest strings, within a radius of a location in the vocabulary space, to a block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/enumerateNeighborsForVector:maximumCount:maximumDistance:distanceType:usingBlock:
func (e_ Embedding) EnumerateNeighborsForVectorMaximumCountMaximumDistanceDistanceTypeUsingBlock(vector []foundation.Number, maxCount uint, maxDistance Distance /* typedef */, distanceType DistanceType, block bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("enumerateNeighborsForVector:maximumCount:maximumDistance:distanceType:usingBlock:"), vector, maxCount, maxDistance, distanceType, block)
}/* debug [instance_methods/method]: EnumerateNeighborsForVectorMaximumCountMaximumDistanceDistanceTypeUsingBlock */


// Copies a vector into the given a pointer to a float array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/getVector:forString:
func (e_ Embedding) GetVectorForString(vector objectivec.IObject, string_ objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("getVector:forString:"), vector, string_)
	return rv
}/* debug [instance_methods/method]: GetVectorForString */


// Retrieves a limited number of strings near a string in the vocabulary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForString:maximumCount:distanceType:
func (e_ Embedding) NeighborsForStringMaximumCountDistanceType(string_ objc.IObject /* cross-framework: NSString */, maxCount uint, distanceType DistanceType) []string {
	rv := objc.Send[[]string](e_.ID, objc.Sel("neighborsForString:maximumCount:distanceType:"), string_, maxCount, distanceType)
	return rv
}/* debug [instance_methods/method]: NeighborsForStringMaximumCountDistanceType */


// Retrieves a limited number of strings, within a radius of a string, in the vocabulary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForString:maximumCount:maximumDistance:distanceType:
func (e_ Embedding) NeighborsForStringMaximumCountMaximumDistanceDistanceType(string_ objc.IObject /* cross-framework: NSString */, maxCount uint, maxDistance Distance /* typedef */, distanceType DistanceType) []string {
	rv := objc.Send[[]string](e_.ID, objc.Sel("neighborsForString:maximumCount:maximumDistance:distanceType:"), string_, maxCount, maxDistance, distanceType)
	return rv
}/* debug [instance_methods/method]: NeighborsForStringMaximumCountMaximumDistanceDistanceType */


// Retrieves a limited number of strings near a location in the vocabulary space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForVector:maximumCount:distanceType:
func (e_ Embedding) NeighborsForVectorMaximumCountDistanceType(vector []foundation.Number, maxCount uint, distanceType DistanceType) []string {
	rv := objc.Send[[]string](e_.ID, objc.Sel("neighborsForVector:maximumCount:distanceType:"), vector, maxCount, distanceType)
	return rv
}/* debug [instance_methods/method]: NeighborsForVectorMaximumCountDistanceType */


// Retrieves a limited number of strings within a radius of a location in the vocabulary space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/neighborsForVector:maximumCount:maximumDistance:distanceType:
func (e_ Embedding) NeighborsForVectorMaximumCountMaximumDistanceDistanceType(vector []foundation.Number, maxCount uint, maxDistance Distance /* typedef */, distanceType DistanceType) []string {
	rv := objc.Send[[]string](e_.ID, objc.Sel("neighborsForVector:maximumCount:maximumDistance:distanceType:"), vector, maxCount, maxDistance, distanceType)
	return rv
}/* debug [instance_methods/method]: NeighborsForVectorMaximumCountMaximumDistanceDistanceType */


// Requests the vector for the given term.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/vectorForString:
func (e_ Embedding) VectorForString(string_ objc.IObject /* cross-framework: NSString */) []foundation.Number {
	rv := objc.Send[[]foundation.Number](e_.ID, objc.Sel("vectorForString:"), string_)
	return rv
}/* debug [instance_methods/method]: VectorForString */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Embedding */

// The number of dimensions in the vocabulary’s vector space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/dimension
func (e_ Embedding) Dimension() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("dimension"))
	return rv
}/* debug [instance_properties/getter]: dimension */


// The language of the text in the word embedding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/language
func (e_ Embedding) Language() Language /* typedef */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("language"))
	return rv
}/* debug [instance_properties/getter]: language */


// The revision of the word embedding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/revision
func (e_ Embedding) Revision() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("revision"))
	return rv
}/* debug [instance_properties/getter]: revision */


// The number of words in the vocabulary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLEmbedding/vocabularySize
func (e_ Embedding) VocabularySize() uint {
	rv := objc.Send[uint](e_.ID, objc.Sel("vocabularySize"))
	return rv
}/* debug [instance_properties/getter]: vocabularySize */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NLEmbedding */


