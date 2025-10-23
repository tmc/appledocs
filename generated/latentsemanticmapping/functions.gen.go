// Code generated from Apple documentation for LatentSemanticMapping. DO NOT EDIT.

package latentsemanticmapping

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// LatentSemanticMapping Functions (30 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_LSMMapAddCategory func(unsafe.Pointer) unsafe.Pointer
	_LSMMapAddText func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMMapAddTextWithWeight func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, float32) unsafe.Pointer
	_LSMMapApplyClusters func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMMapCompile func(unsafe.Pointer) unsafe.Pointer
	_LSMMapCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMMapCreateClusters func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMMapCreateFromURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMMapGetCategoryCount func(unsafe.Pointer) unsafe.Pointer
	_LSMMapGetProperties func(unsafe.Pointer) unsafe.Pointer
	_LSMMapGetTypeID func() unsafe.Pointer
	_LSMMapSetProperties func(unsafe.Pointer, unsafe.Pointer)
	_LSMMapSetStopWords func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMMapStartTraining func(unsafe.Pointer) unsafe.Pointer
	_LSMMapWriteToStream func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMMapWriteToURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMResultCopyToken func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMResultCopyTokenCluster func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMResultCopyWord func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMResultCopyWordCluster func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMResultCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMResultGetCategory func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMResultGetCount func(unsafe.Pointer) unsafe.Pointer
	_LSMResultGetScore func(unsafe.Pointer, unsafe.Pointer) float32
	_LSMResultGetTypeID func() unsafe.Pointer
	_LSMTextAddToken func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMTextAddWord func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMTextAddWords func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMTextCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSMTextGetTypeID func() unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_LSMMapAddCategory, lib, "LSMMapAddCategory")
	tryRegister(&_LSMMapAddText, lib, "LSMMapAddText")
	tryRegister(&_LSMMapAddTextWithWeight, lib, "LSMMapAddTextWithWeight")
	tryRegister(&_LSMMapApplyClusters, lib, "LSMMapApplyClusters")
	tryRegister(&_LSMMapCompile, lib, "LSMMapCompile")
	tryRegister(&_LSMMapCreate, lib, "LSMMapCreate")
	tryRegister(&_LSMMapCreateClusters, lib, "LSMMapCreateClusters")
	tryRegister(&_LSMMapCreateFromURL, lib, "LSMMapCreateFromURL")
	tryRegister(&_LSMMapGetCategoryCount, lib, "LSMMapGetCategoryCount")
	tryRegister(&_LSMMapGetProperties, lib, "LSMMapGetProperties")
	tryRegister(&_LSMMapGetTypeID, lib, "LSMMapGetTypeID")
	tryRegister(&_LSMMapSetProperties, lib, "LSMMapSetProperties")
	tryRegister(&_LSMMapSetStopWords, lib, "LSMMapSetStopWords")
	tryRegister(&_LSMMapStartTraining, lib, "LSMMapStartTraining")
	tryRegister(&_LSMMapWriteToStream, lib, "LSMMapWriteToStream")
	tryRegister(&_LSMMapWriteToURL, lib, "LSMMapWriteToURL")
	tryRegister(&_LSMResultCopyToken, lib, "LSMResultCopyToken")
	tryRegister(&_LSMResultCopyTokenCluster, lib, "LSMResultCopyTokenCluster")
	tryRegister(&_LSMResultCopyWord, lib, "LSMResultCopyWord")
	tryRegister(&_LSMResultCopyWordCluster, lib, "LSMResultCopyWordCluster")
	tryRegister(&_LSMResultCreate, lib, "LSMResultCreate")
	tryRegister(&_LSMResultGetCategory, lib, "LSMResultGetCategory")
	tryRegister(&_LSMResultGetCount, lib, "LSMResultGetCount")
	tryRegister(&_LSMResultGetScore, lib, "LSMResultGetScore")
	tryRegister(&_LSMResultGetTypeID, lib, "LSMResultGetTypeID")
	tryRegister(&_LSMTextAddToken, lib, "LSMTextAddToken")
	tryRegister(&_LSMTextAddWord, lib, "LSMTextAddWord")
	tryRegister(&_LSMTextAddWords, lib, "LSMTextAddWords")
	tryRegister(&_LSMTextCreate, lib, "LSMTextCreate")
	tryRegister(&_LSMTextGetTypeID, lib, "LSMTextGetTypeID")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Adds another category and returns its category identifier.

// Adds another category and returns its category identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapAddCategory(_:)
func LSMMapAddCategory(mapref unsafe.Pointer) unsafe.Pointer {
	return _LSMMapAddCategory(mapref)
}

// Adds a training text to the specified category.

// Adds a training text to the specified category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapAddText(_:_:_:)
func LSMMapAddText(mapref unsafe.Pointer, textref unsafe.Pointer, category unsafe.Pointer) unsafe.Pointer {
	return _LSMMapAddText(mapref, textref, category)
}

// Adds a training text to the specified category with a weight other than 1.

// Adds a training text to the specified category with a weight other than 1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapAddTextWithWeight(_:_:_:_:)
func LSMMapAddTextWithWeight(mapref unsafe.Pointer, textref unsafe.Pointer, category unsafe.Pointer, weight float32) unsafe.Pointer {
	return _LSMMapAddTextWithWeight(mapref, textref, category, weight)
}

// Groups categories or words (tokens) into the specified sets of clusters.

// Groups categories or words (tokens) into the specified sets of clusters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapApplyClusters(_:_:)
func LSMMapApplyClusters(mapref unsafe.Pointer, clusters unsafe.Pointer) unsafe.Pointer {
	return _LSMMapApplyClusters(mapref, clusters)
}

// Compiles the map into executable form and puts it into mapping mode, preparing it for the classification of texts.

// Compiles the map into executable form and puts it into mapping mode, preparing it for the classification of texts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapCompile(_:)
func LSMMapCompile(mapref unsafe.Pointer) unsafe.Pointer {
	return _LSMMapCompile(mapref)
}

// Creates a new Latent Semantic Mapping map.

// Creates a new Latent Semantic Mapping map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapCreate(_:_:)
func LSMMapCreate(alloc unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _LSMMapCreate(alloc, flags)
}

// Computes a set of clusters that group similar categories or words.

// Computes a set of clusters that group similar categories or words.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapCreateClusters(_:_:_:_:_:)
func LSMMapCreateClusters(alloc unsafe.Pointer, mapref unsafe.Pointer, subset unsafe.Pointer, numClusters unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _LSMMapCreateClusters(alloc, mapref, subset, numClusters, flags)
}

// Loads a map from the specified file.

// Loads a map from the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapCreateFromURL(_:_:_:)
func LSMMapCreateFromURL(alloc unsafe.Pointer, file unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _LSMMapCreateFromURL(alloc, file, flags)
}

// Returns the number of categories in the map.

// Returns the number of categories in the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapGetCategoryCount(_:)
func LSMMapGetCategoryCount(mapref unsafe.Pointer) unsafe.Pointer {
	return _LSMMapGetCategoryCount(mapref)
}

// Gets a dictionary of properties for the map.

// Gets a dictionary of properties for the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapGetProperties(_:)
func LSMMapGetProperties(mapref unsafe.Pointer) unsafe.Pointer {
	return _LSMMapGetProperties(mapref)
}

// Returns the Core Foundation type identifier for Latent Semantic Mapping maps.

// Returns the Core Foundation type identifier for Latent Semantic Mapping maps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapGetTypeID()
func LSMMapGetTypeID() unsafe.Pointer {
	return _LSMMapGetTypeID()
}

// Sets a dictionary of properties for the map.

// Sets a dictionary of properties for the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapSetProperties(_:_:)
func LSMMapSetProperties(mapref unsafe.Pointer, properties unsafe.Pointer) {
	_LSMMapSetProperties(mapref, properties)
}

// Specifies which words to omit from all classification efforts.

// Specifies which words to omit from all classification efforts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapSetStopWords(_:_:)
func LSMMapSetStopWords(mapref unsafe.Pointer, textref unsafe.Pointer) unsafe.Pointer {
	return _LSMMapSetStopWords(mapref, textref)
}

// Puts the map into training mode, preparing it for the addition of more categories or texts.

// Puts the map into training mode, preparing it for the addition of more categories or texts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapStartTraining(_:)
func LSMMapStartTraining(mapref unsafe.Pointer) unsafe.Pointer {
	return _LSMMapStartTraining(mapref)
}

// Writes information about a map or text to a stream in text form.

// Writes information about a map or text to a stream in text form.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapWriteToStream(_:_:_:_:)
func LSMMapWriteToStream(mapref unsafe.Pointer, textref unsafe.Pointer, stream unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _LSMMapWriteToStream(mapref, textref, stream, options)
}

// Compiles the map, if necessary, and stores it into the specified file.

// Compiles the map, if necessary, and stores it into the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapWriteToURL(_:_:_:)
func LSMMapWriteToURL(mapref unsafe.Pointer, file unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _LSMMapWriteToURL(mapref, file, flags)
}

// Returns the token for the n-th best (zero-based) result.

// Returns the token for the n-th best (zero-based) result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultCopyToken(_:_:)
func LSMResultCopyToken(result unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	return _LSMResultCopyToken(result, n)
}

// Returns the cluster of tokens for the n-th best (zero-based) result.

// Returns the cluster of tokens for the n-th best (zero-based) result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultCopyTokenCluster(_:_:)
func LSMResultCopyTokenCluster(result unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	return _LSMResultCopyTokenCluster(result, n)
}

// Returns the word for the n-th best (zero-based) result.

// Returns the word for the n-th best (zero-based) result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultCopyWord(_:_:)
func LSMResultCopyWord(result unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	return _LSMResultCopyWord(result, n)
}

// Returns the cluster of words for the n-th best (zero-based) result.

// Returns the cluster of words for the n-th best (zero-based) result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultCopyWordCluster(_:_:)
func LSMResultCopyWordCluster(result unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	return _LSMResultCopyWordCluster(result, n)
}

// Returns the categories or words that best match when a text is mapped into a map, in decreasing order of likelihood.

// Returns the categories or words that best match when a text is mapped into a map, in decreasing order of likelihood.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultCreate(_:_:_:_:_:)
func LSMResultCreate(alloc unsafe.Pointer, mapref unsafe.Pointer, textref unsafe.Pointer, numResults unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _LSMResultCreate(alloc, mapref, textref, numResults, flags)
}

// Returns the category of the specified result.

// Returns the category of the specified result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultGetCategory(_:_:)
func LSMResultGetCategory(result unsafe.Pointer, n unsafe.Pointer) unsafe.Pointer {
	return _LSMResultGetCategory(result, n)
}

// Returns the number of results.

// Returns the number of results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultGetCount(_:)
func LSMResultGetCount(result unsafe.Pointer) unsafe.Pointer {
	return _LSMResultGetCount(result)
}

// Returns the likelihood of the specified result.

// Returns the likelihood of the specified result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultGetScore(_:_:)
func LSMResultGetScore(result unsafe.Pointer, n unsafe.Pointer) float32 {
	return _LSMResultGetScore(result, n)
}

// Returns the Core Foundation type identifier for Latent Semantic Mapping results.

// Returns the Core Foundation type identifier for Latent Semantic Mapping results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultGetTypeID()
func LSMResultGetTypeID() unsafe.Pointer {
	return _LSMResultGetTypeID()
}

// Adds an arbitrary binary token to the text.

// Adds an arbitrary binary token to the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMTextAddToken(_:_:)
func LSMTextAddToken(textref unsafe.Pointer, token unsafe.Pointer) unsafe.Pointer {
	return _LSMTextAddToken(textref, token)
}

// Adds a word to the text.

// Adds a word to the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMTextAddWord(_:_:)
func LSMTextAddWord(textref unsafe.Pointer, word unsafe.Pointer) unsafe.Pointer {
	return _LSMTextAddWord(textref, word)
}

// Breaks a string into words using the specified locale, and adds the words to the text.

// Breaks a string into words using the specified locale, and adds the words to the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMTextAddWords(_:_:_:_:)
func LSMTextAddWords(textref unsafe.Pointer, words unsafe.Pointer, locale unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _LSMTextAddWords(textref, words, locale, flags)
}

// Creates a new text.

// Creates a new text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMTextCreate(_:_:)
func LSMTextCreate(alloc unsafe.Pointer, mapref unsafe.Pointer) unsafe.Pointer {
	return _LSMTextCreate(alloc, mapref)
}

// Returns the Core Foundation type identifier for Latent Semantic Mapping texts.

// Returns the Core Foundation type identifier for Latent Semantic Mapping texts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMTextGetTypeID()
func LSMTextGetTypeID() unsafe.Pointer {
	return _LSMTextGetTypeID()
}



