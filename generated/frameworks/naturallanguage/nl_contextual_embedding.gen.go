// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NLContextualEmbedding] class.
var nLContextualEmbeddingClass = _NLContextualEmbeddingClass{objc.GetClass("NLContextualEmbedding")}

type _NLContextualEmbeddingClass struct {
	class objc.Class
}

// An interface definition for the [NLContextualEmbedding] class.
type INLContextualEmbedding interface {
	objectivec.IObject
	EmbeddingResultForStringLanguageError(string string, language unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	LoadWithError(error unsafe.Pointer) bool
	RequestEmbeddingAssetsWithCompletionHandler(completionHandler unsafe.Pointer)
	Unload()
}

// A model that computes sequences of embedding vectors for natural language utterances. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding

type NLContextualEmbedding struct {
	objectivec.Object
}

// NLContextualEmbeddingFrom constructs a [NLContextualEmbedding] from an unsafe.Pointer.
//
// A model that computes sequences of embedding vectors for natural language utterances.
func NLContextualEmbeddingFrom(ptr unsafe.Pointer) NLContextualEmbedding {
	return NLContextualEmbedding{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NLContextualEmbeddingClass) Alloc() NLContextualEmbedding {
	rv := objc.Send[NLContextualEmbedding](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NLContextualEmbeddingClass) New() NLContextualEmbedding {
	rv := objc.Send[NLContextualEmbedding](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NLContextualEmbedding) Init() NLContextualEmbedding {
	rv := objc.Send[NLContextualEmbedding](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NLContextualEmbedding) Autorelease() NLContextualEmbedding {
	rv := objc.Send[NLContextualEmbedding](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLContextualEmbedding creates a new NLContextualEmbedding instance.
func NewNLContextualEmbedding() NLContextualEmbedding {
	return nLContextualEmbeddingClass.New()
}


// Creates a contextual embedding from a model identifier. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(modelIdentifier:)
func NewContextualEmbeddingWithModelIdentifier(modelIdentifier string) NLContextualEmbedding {
	rv := objc.Send[NLContextualEmbedding](objc.ID(nLContextualEmbeddingClass.class), objc.Sel("contextualEmbeddingWithModelIdentifier:"), modelIdentifier)
	rv.Autorelease()
	return rv
}
// Creates a contextual embedding from a script. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(script:)
func NewContextualEmbeddingWithScript(script unsafe.Pointer) NLContextualEmbedding {
	rv := objc.Send[NLContextualEmbedding](objc.ID(nLContextualEmbeddingClass.class), objc.Sel("contextualEmbeddingWithScript:"), script)
	rv.Autorelease()
	return rv
}
// Creates a contextual embedding from a language. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(language:)
func NewContextualEmbeddingWithLanguage(language unsafe.Pointer) NLContextualEmbedding {
	rv := objc.Send[NLContextualEmbedding](objc.ID(nLContextualEmbeddingClass.class), objc.Sel("contextualEmbeddingWithLanguage:"), language)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/contextualEmbeddings(forValues:)
func (nc _NLContextualEmbeddingClass) ContextualEmbeddingsForValues(valuesDictionary unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("contextualEmbeddingsForValues:"), valuesDictionary)
	return rv
}
// Creates a contextual embedding from a language. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(language:)
func (nc _NLContextualEmbeddingClass) ContextualEmbeddingWithLanguage(language unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("contextualEmbeddingWithLanguage:"), language)
	return rv
}
// Creates a contextual embedding from a model identifier. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(modelIdentifier:)
func (nc _NLContextualEmbeddingClass) ContextualEmbeddingWithModelIdentifier(modelIdentifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("contextualEmbeddingWithModelIdentifier:"), modelIdentifier)
	return rv
}
// Creates a contextual embedding from a script. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/init(script:)
func (nc _NLContextualEmbeddingClass) ContextualEmbeddingWithScript(script unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("contextualEmbeddingWithScript:"), script)
	return rv
}
// Applies an embedding to a string and obtains the resulting embedding vectors. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/embeddingResult(for:language:)
func (n_ NLContextualEmbedding) EmbeddingResultForStringLanguageError(string string, language unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("embeddingResultForString:language:error:"), string, language, error)
	return rv
}
// Loads the embedding model. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/load()
func (n_ NLContextualEmbedding) LoadWithError(error unsafe.Pointer) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("loadWithError:"), error)
	return rv
}
// Requests assets for an embedding, if available. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/requestAssets(completionHandler:)
func (n_ NLContextualEmbedding) RequestEmbeddingAssetsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("requestEmbeddingAssetsWithCompletionHandler:"), completionHandler)
}
// Unloads the embedding model. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLContextualEmbedding/unload()
func (n_ NLContextualEmbedding) Unload() {
	objc.Send[objc.ID](n_.ID, objc.Sel("unload"))
}

