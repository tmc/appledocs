// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NLContextualEmbedding */


/* debug [class_header]: Header for NLContextualEmbedding */
// The class instance for the [ContextualEmbedding] class.
var (
	ContextualEmbeddingClass     _ContextualEmbeddingClass
	ContextualEmbeddingClassOnce sync.Once
)

func getContextualEmbeddingClass() _ContextualEmbeddingClass {
	ContextualEmbeddingClassOnce.Do(func() {
		ContextualEmbeddingClass = _ContextualEmbeddingClass{objc.GetClass("NLContextualEmbedding")}
	})
	return ContextualEmbeddingClass
}

type _ContextualEmbeddingClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ContextualEmbedding */
// An interface definition for the [ContextualEmbedding] class.
type IContextualEmbedding interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ContextualEmbedding */
	// properties:
	Dimension() uint
	HasAvailableAssets() bool
	Languages() []string
	MaximumSequenceLength() uint
	ModelIdentifier() objc.IObject /* cross-framework: NSString */
	Revision() uint
	Scripts() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ContextualEmbedding */
	// methods:
	EmbeddingResultForStringLanguageError(string_ objc.IObject /* cross-framework: NSString */, language Language /* typedef */, error_ objectivec.IObject) IContextualEmbeddingResult
	LoadWithError(error_ objectivec.IObject) bool
	RequestEmbeddingAssetsWithCompletionHandler(completionHandler unsafe.Pointer)
	Unload()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ContextualEmbedding */
// Alloc allocates a new instance without initialization.
func (cc _ContextualEmbeddingClass) Alloc() ContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContextualEmbeddingClass) New() ContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContextualEmbedding) Init() ContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContextualEmbedding) Autorelease() ContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContextualEmbedding creates a new ContextualEmbedding instance.
func NewContextualEmbedding() ContextualEmbedding {
	return getContextualEmbeddingClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ContextualEmbedding */
// A model that computes sequences of embedding vectors for natural language utterances.
//
// Starting in iOS 17 and macOS 14, the framework supports 27 languages across three models: Latin — including Croatian, Czech, Danish, Dutch, English, Finnish, French, German, Hungarian, Indonesian, Italian, Norwegian, Polish, Portuguese, Romanian, Slovak, Swedish, Spanish, Turkish, and Vietnamese Cyrillic — including Bulgarian, Kazakh, Russian, and Ukrainian Chinese, Japanese, and Korean In iOS 18 and macOS 15, the framework expands language support to include three additional models: Arabic Thai Indic — including Hindi, Marathi, Bangla, Urdu, Punjabi, Gujarati, Tamil, Telugu, Kannada, and Malayalam


// A model that computes sequences of embedding vectors for natural language utterances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding
type ContextualEmbedding struct {
	objectivec.Object
}

// ContextualEmbeddingFrom constructs a [ContextualEmbedding] from an unsafe.Pointer.
//
// A model that computes sequences of embedding vectors for natural language utterances.
func ContextualEmbeddingFrom(ptr unsafe.Pointer) ContextualEmbedding {
	return ContextualEmbedding{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ContextualEmbedding */

// Creates a contextual embedding from a language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(language:)
func NewContextualEmbeddingWithLanguage(language Language /* typedef */) ContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](objc.ID(getContextualEmbeddingClass().class), objc.Sel("contextualEmbeddingWithLanguage:"), language)
	return rv
}/* debug [class_init_methods/constructor]: NewContextualEmbeddingWithLanguage */


// Creates a contextual embedding from a model identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(modelIdentifier:)
func NewContextualEmbeddingWithModelIdentifier(modelIdentifier objc.IObject /* cross-framework: NSString */) ContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](objc.ID(getContextualEmbeddingClass().class), objc.Sel("contextualEmbeddingWithModelIdentifier:"), modelIdentifier)
	return rv
}/* debug [class_init_methods/constructor]: NewContextualEmbeddingWithModelIdentifier */


// Creates a contextual embedding from a script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(script:)
func NewContextualEmbeddingWithScript(script Script /* typedef */) ContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](objc.ID(getContextualEmbeddingClass().class), objc.Sel("contextualEmbeddingWithScript:"), script)
	return rv
}/* debug [class_init_methods/constructor]: NewContextualEmbeddingWithScript */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ContextualEmbedding */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/contextualEmbeddings(forValues:)
func (cc _ContextualEmbeddingClass) ContextualEmbeddingsForValues(valuesDictionary foundation.IDictionary) []ContextualEmbedding {
	rv := objc.Send[[]ContextualEmbedding](objc.ID(cc.class), objc.Sel("contextualEmbeddingsForValues:"), valuesDictionary)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContextualEmbeddingsForValues) */


// Creates a contextual embedding from a language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(language:)
func (cc _ContextualEmbeddingClass) ContextualEmbeddingWithLanguage(language Language /* typedef */) IContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](objc.ID(cc.class), objc.Sel("contextualEmbeddingWithLanguage:"), language)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContextualEmbeddingWithLanguage) */


// Creates a contextual embedding from a model identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(modelIdentifier:)
func (cc _ContextualEmbeddingClass) ContextualEmbeddingWithModelIdentifier(modelIdentifier objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("contextualEmbeddingWithModelIdentifier:"), modelIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContextualEmbeddingWithModelIdentifier) */


// Creates a contextual embedding from a script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(script:)
func (cc _ContextualEmbeddingClass) ContextualEmbeddingWithScript(script Script /* typedef */) IContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](objc.ID(cc.class), objc.Sel("contextualEmbeddingWithScript:"), script)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContextualEmbeddingWithScript) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ContextualEmbedding */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ContextualEmbedding */

// Applies an embedding to a string and obtains the resulting embedding vectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/embeddingResult(for:language:)
func (c_ ContextualEmbedding) EmbeddingResultForStringLanguageError(string_ objc.IObject /* cross-framework: NSString */, language Language /* typedef */, error_ objectivec.IObject) IContextualEmbeddingResult {
	rv := objc.Send[ContextualEmbeddingResult](c_.ID, objc.Sel("embeddingResultForString:language:error:"), string_, language, error_)
	return rv
}/* debug [instance_methods/method]: EmbeddingResultForStringLanguageError */


// Loads the embedding model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/load()
func (c_ ContextualEmbedding) LoadWithError(error_ objectivec.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("loadWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: LoadWithError */


// Requests assets for an embedding, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/requestAssets(completionHandler:)
func (c_ ContextualEmbedding) RequestEmbeddingAssetsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("requestEmbeddingAssetsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: RequestEmbeddingAssetsWithCompletionHandler */


// Unloads the embedding model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/unload()
func (c_ ContextualEmbedding) Unload() {
	objc.Send[objc.ID](c_.ID, objc.Sel("unload"))
}/* debug [instance_methods/method]: Unload */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ContextualEmbedding */

// The number of dimensions in the script’s vector space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/dimension
func (c_ ContextualEmbedding) Dimension() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dimension"))
	return rv
}/* debug [instance_properties/getter]: dimension */


// A Boolean value that indicates whether assets are available to load.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/hasAvailableAssets
func (c_ ContextualEmbedding) HasAvailableAssets() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasAvailableAssets"))
	return rv
}/* debug [instance_properties/getter]: hasAvailableAssets */


// The languages of the text in the contextual embedding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/languages
func (c_ ContextualEmbedding) Languages() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("languages"))
	return rv
}/* debug [instance_properties/getter]: languages */


// The maximum number of embedding vectors the model generates, in sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/maximumSequenceLength
func (c_ ContextualEmbedding) MaximumSequenceLength() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maximumSequenceLength"))
	return rv
}/* debug [instance_properties/getter]: maximumSequenceLength */


// The model identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/modelIdentifier
func (c_ ContextualEmbedding) ModelIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("modelIdentifier"))
	return rv
}/* debug [instance_properties/getter]: modelIdentifier */


// The revision of the contextual embedding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/revision
func (c_ ContextualEmbedding) Revision() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("revision"))
	return rv
}/* debug [instance_properties/getter]: revision */


// The scripts of the text in the contextual embedding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/scripts
func (c_ ContextualEmbedding) Scripts() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("scripts"))
	return rv
}/* debug [instance_properties/getter]: scripts */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NLContextualEmbedding */


