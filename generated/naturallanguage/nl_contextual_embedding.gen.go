// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [ContextualEmbedding] class.
type IContextualEmbedding interface {
	objectivec.IObject
	

	// properties:
	Dimension() uint
	HasAvailableAssets() bool
	Languages() []string
	MaximumSequenceLength() uint
	ModelIdentifier() foundation.foundation.INSString
	Revision() uint
	Scripts() []string


	

	// methods:
	EmbeddingResultForStringLanguageError(string_ foundation.foundation.INSString, language Language, error_ foundation.foundation.INSError) IContextualEmbeddingResult
	LoadWithError(error_ foundation.foundation.INSError) bool
	RequestEmbeddingAssetsWithCompletionHandler(completionHandler unsafe.Pointer)
	Unload()


}





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






// Creates a contextual embedding from a language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(language:)
func NewContextualEmbeddingWithLanguage(language Language) ContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](objc.ID(getContextualEmbeddingClass().class), objc.Sel("contextualEmbeddingWithLanguage:"), language)
	return rv
}


// Creates a contextual embedding from a model identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(modelIdentifier:)
func NewContextualEmbeddingWithModelIdentifier(modelIdentifier foundation.foundation.INSString) ContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](objc.ID(getContextualEmbeddingClass().class), objc.Sel("contextualEmbeddingWithModelIdentifier:"), modelIdentifier)
	return rv
}


// Creates a contextual embedding from a script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(script:)
func NewContextualEmbeddingWithScript(script Script) ContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](objc.ID(getContextualEmbeddingClass().class), objc.Sel("contextualEmbeddingWithScript:"), script)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/contextualEmbeddings(forValues:)
func (cc _ContextualEmbeddingClass) ContextualEmbeddingsForValues(valuesDictionary foundation.IDictionary) []ContextualEmbedding {
	rv := objc.Send[[]ContextualEmbedding](objc.ID(cc.class), objc.Sel("contextualEmbeddingsForValues:"), valuesDictionary)
	return rv
}


// Creates a contextual embedding from a language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(language:)
func (cc _ContextualEmbeddingClass) ContextualEmbeddingWithLanguage(language Language) IContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](objc.ID(cc.class), objc.Sel("contextualEmbeddingWithLanguage:"), language)
	return rv
}


// Creates a contextual embedding from a model identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(modelIdentifier:)
func (cc _ContextualEmbeddingClass) ContextualEmbeddingWithModelIdentifier(modelIdentifier foundation.foundation.INSString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("contextualEmbeddingWithModelIdentifier:"), modelIdentifier)
	return rv
}


// Creates a contextual embedding from a script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(script:)
func (cc _ContextualEmbeddingClass) ContextualEmbeddingWithScript(script Script) IContextualEmbedding {
	rv := objc.Send[ContextualEmbedding](objc.ID(cc.class), objc.Sel("contextualEmbeddingWithScript:"), script)
	return rv
}












// Applies an embedding to a string and obtains the resulting embedding vectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/embeddingResult(for:language:)
func (c_ ContextualEmbedding) EmbeddingResultForStringLanguageError(string_ foundation.foundation.INSString, language Language, error_ foundation.foundation.INSError) IContextualEmbeddingResult {
	rv := objc.Send[ContextualEmbeddingResult](c_.ID, objc.Sel("embeddingResultForString:language:error:"), string_, language, error_)
	return rv
}


// Loads the embedding model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/load()
func (c_ ContextualEmbedding) LoadWithError(error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("loadWithError:"), error_)
	return rv
}


// Requests assets for an embedding, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/requestAssets(completionHandler:)
func (c_ ContextualEmbedding) RequestEmbeddingAssetsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("requestEmbeddingAssetsWithCompletionHandler:"), completionHandler)
}


// Unloads the embedding model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/unload()
func (c_ ContextualEmbedding) Unload() {
	objc.Send[objc.ID](c_.ID, objc.Sel("unload"))
}







// The number of dimensions in the script’s vector space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/dimension
func (c_ ContextualEmbedding) Dimension() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dimension"))
	return rv
}


// A Boolean value that indicates whether assets are available to load.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/hasAvailableAssets
func (c_ ContextualEmbedding) HasAvailableAssets() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasAvailableAssets"))
	return rv
}


// The languages of the text in the contextual embedding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/languages
func (c_ ContextualEmbedding) Languages() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("languages"))
	return rv
}


// The maximum number of embedding vectors the model generates, in sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/maximumSequenceLength
func (c_ ContextualEmbedding) MaximumSequenceLength() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("maximumSequenceLength"))
	return rv
}


// The model identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/modelIdentifier
func (c_ ContextualEmbedding) ModelIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("modelIdentifier"))
	return rv
}


// The revision of the contextual embedding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/revision
func (c_ ContextualEmbedding) Revision() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("revision"))
	return rv
}


// The scripts of the text in the contextual embedding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/scripts
func (c_ ContextualEmbedding) Scripts() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("scripts"))
	return rv
}







