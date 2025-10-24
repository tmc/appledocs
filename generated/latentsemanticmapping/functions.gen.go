// Code generated from Apple documentation for LatentSemanticMapping. DO NOT EDIT.

package latentsemanticmapping

/* debug [functions.gen.go]: Generating 30 functions for LatentSemanticMapping */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// LatentSemanticMapping Functions (30 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_LSMMapAddCategory func(LSMMapRef) LSMCategory
	_LSMMapAddText func(LSMMapRef, LSMTextRef, LSMCategory) unsafe.Pointer
	_LSMMapAddTextWithWeight func(LSMMapRef, LSMTextRef, LSMCategory, float32) unsafe.Pointer
	_LSMMapApplyClusters func(LSMMapRef, ArrayRef) unsafe.Pointer
	_LSMMapCompile func(LSMMapRef) unsafe.Pointer
	_LSMMapCreate func(AllocatorRef, OptionFlags) LSMMapRef
	_LSMMapCreateClusters func(AllocatorRef, LSMMapRef, ArrayRef, Index, OptionFlags) ArrayRef
	_LSMMapCreateFromURL func(AllocatorRef, URLRef, OptionFlags) LSMMapRef
	_LSMMapGetCategoryCount func(LSMMapRef) Index
	_LSMMapGetProperties func(LSMMapRef) DictionaryRef
	_LSMMapGetTypeID func() TypeID
	_LSMMapSetProperties func(LSMMapRef, DictionaryRef)
	_LSMMapSetStopWords func(LSMMapRef, LSMTextRef) unsafe.Pointer
	_LSMMapStartTraining func(LSMMapRef) unsafe.Pointer
	_LSMMapWriteToStream func(LSMMapRef, LSMTextRef, WriteStreamRef, OptionFlags) unsafe.Pointer
	_LSMMapWriteToURL func(LSMMapRef, URLRef, OptionFlags) unsafe.Pointer
	_LSMResultCopyToken func(LSMResultRef, Index) DataRef
	_LSMResultCopyTokenCluster func(LSMResultRef, Index) ArrayRef
	_LSMResultCopyWord func(LSMResultRef, Index) StringRef
	_LSMResultCopyWordCluster func(LSMResultRef, Index) ArrayRef
	_LSMResultCreate func(AllocatorRef, LSMMapRef, LSMTextRef, Index, OptionFlags) LSMResultRef
	_LSMResultGetCategory func(LSMResultRef, Index) LSMCategory
	_LSMResultGetCount func(LSMResultRef) Index
	_LSMResultGetScore func(LSMResultRef, Index) float32
	_LSMResultGetTypeID func() TypeID
	_LSMTextAddToken func(LSMTextRef, DataRef) unsafe.Pointer
	_LSMTextAddWord func(LSMTextRef, StringRef) unsafe.Pointer
	_LSMTextAddWords func(LSMTextRef, StringRef, LocaleRef, OptionFlags) unsafe.Pointer
	_LSMTextCreate func(AllocatorRef, LSMMapRef) LSMTextRef
	_LSMTextGetTypeID func() TypeID
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
//
// Added in macOS .
// Adds another category and returns its category identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapAddCategory(_:)
func LSMMapAddCategory(mapref LSMMapRef) LSMCategory {
	return _LSMMapAddCategory(mapref)
}/* debug [functions.gen.go/function]: LSMMapAddCategory */

// Adds a training text to the specified category.
//
// Added in macOS .
// Adds a training text to the specified category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapAddText(_:_:_:)
func LSMMapAddText(mapref LSMMapRef, textref LSMTextRef, category LSMCategory) unsafe.Pointer {
	return _LSMMapAddText(mapref, textref, category)
}/* debug [functions.gen.go/function]: LSMMapAddText */

// Adds a training text to the specified category with a weight other than 1.
//
// Added in macOS .
// Adds a training text to the specified category with a weight other than 1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapAddTextWithWeight(_:_:_:_:)
func LSMMapAddTextWithWeight(mapref LSMMapRef, textref LSMTextRef, category LSMCategory, weight float32) unsafe.Pointer {
	return _LSMMapAddTextWithWeight(mapref, textref, category, weight)
}/* debug [functions.gen.go/function]: LSMMapAddTextWithWeight */

// Groups categories or words (tokens) into the specified sets of clusters.
//
// Added in macOS .
// Groups categories or words (tokens) into the specified sets of clusters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapApplyClusters(_:_:)
func LSMMapApplyClusters(mapref LSMMapRef, clusters ArrayRef) unsafe.Pointer {
	return _LSMMapApplyClusters(mapref, clusters)
}/* debug [functions.gen.go/function]: LSMMapApplyClusters */

// Compiles the map into executable form and puts it into mapping mode, preparing it for the classification of texts.
//
// Added in macOS .
// Compiles the map into executable form and puts it into mapping mode, preparing it for the classification of texts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapCompile(_:)
func LSMMapCompile(mapref LSMMapRef) unsafe.Pointer {
	return _LSMMapCompile(mapref)
}/* debug [functions.gen.go/function]: LSMMapCompile */

// Creates a new Latent Semantic Mapping map.
//
// Added in macOS .
// Creates a new Latent Semantic Mapping map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapCreate(_:_:)
func LSMMapCreate(alloc AllocatorRef, flags OptionFlags) LSMMapRef {
	return _LSMMapCreate(alloc, flags)
}/* debug [functions.gen.go/function]: LSMMapCreate */

// Computes a set of clusters that group similar categories or words.
//
// Added in macOS .
// Computes a set of clusters that group similar categories or words.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapCreateClusters(_:_:_:_:_:)
func LSMMapCreateClusters(alloc AllocatorRef, mapref LSMMapRef, subset ArrayRef, numClusters Index, flags OptionFlags) ArrayRef {
	return _LSMMapCreateClusters(alloc, mapref, subset, numClusters, flags)
}/* debug [functions.gen.go/function]: LSMMapCreateClusters */

// Loads a map from the specified file.
//
// Added in macOS .
// Loads a map from the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapCreateFromURL(_:_:_:)
func LSMMapCreateFromURL(alloc AllocatorRef, file URLRef, flags OptionFlags) LSMMapRef {
	return _LSMMapCreateFromURL(alloc, file, flags)
}/* debug [functions.gen.go/function]: LSMMapCreateFromURL */

// Returns the number of categories in the map.
//
// Added in macOS .
// Returns the number of categories in the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapGetCategoryCount(_:)
func LSMMapGetCategoryCount(mapref LSMMapRef) Index {
	return _LSMMapGetCategoryCount(mapref)
}/* debug [functions.gen.go/function]: LSMMapGetCategoryCount */

// Gets a dictionary of properties for the map.
//
// Added in macOS .
// Gets a dictionary of properties for the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapGetProperties(_:)
func LSMMapGetProperties(mapref LSMMapRef) DictionaryRef {
	return _LSMMapGetProperties(mapref)
}/* debug [functions.gen.go/function]: LSMMapGetProperties */

// Returns the Core Foundation type identifier for Latent Semantic Mapping maps.
//
// Added in macOS .
// Returns the Core Foundation type identifier for Latent Semantic Mapping maps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapGetTypeID()
func LSMMapGetTypeID() TypeID {
	return _LSMMapGetTypeID()
}/* debug [functions.gen.go/function]: LSMMapGetTypeID */

// Sets a dictionary of properties for the map.
//
// Added in macOS .
// Sets a dictionary of properties for the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapSetProperties(_:_:)
func LSMMapSetProperties(mapref LSMMapRef, properties DictionaryRef) {
	_LSMMapSetProperties(mapref, properties)
}/* debug [functions.gen.go/function]: LSMMapSetProperties */

// Specifies which words to omit from all classification efforts.
//
// Added in macOS .
// Specifies which words to omit from all classification efforts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapSetStopWords(_:_:)
func LSMMapSetStopWords(mapref LSMMapRef, textref LSMTextRef) unsafe.Pointer {
	return _LSMMapSetStopWords(mapref, textref)
}/* debug [functions.gen.go/function]: LSMMapSetStopWords */

// Puts the map into training mode, preparing it for the addition of more categories or texts.
//
// Added in macOS .
// Puts the map into training mode, preparing it for the addition of more categories or texts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapStartTraining(_:)
func LSMMapStartTraining(mapref LSMMapRef) unsafe.Pointer {
	return _LSMMapStartTraining(mapref)
}/* debug [functions.gen.go/function]: LSMMapStartTraining */

// Writes information about a map or text to a stream in text form.
//
// Added in macOS .
// Writes information about a map or text to a stream in text form.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapWriteToStream(_:_:_:_:)
func LSMMapWriteToStream(mapref LSMMapRef, textref LSMTextRef, stream WriteStreamRef, options OptionFlags) unsafe.Pointer {
	return _LSMMapWriteToStream(mapref, textref, stream, options)
}/* debug [functions.gen.go/function]: LSMMapWriteToStream */

// Compiles the map, if necessary, and stores it into the specified file.
//
// Added in macOS .
// Compiles the map, if necessary, and stores it into the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMMapWriteToURL(_:_:_:)
func LSMMapWriteToURL(mapref LSMMapRef, file URLRef, flags OptionFlags) unsafe.Pointer {
	return _LSMMapWriteToURL(mapref, file, flags)
}/* debug [functions.gen.go/function]: LSMMapWriteToURL */

// Returns the token for the n-th best (zero-based) result.
//
// Added in macOS .
// Returns the token for the n-th best (zero-based) result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultCopyToken(_:_:)
func LSMResultCopyToken(result LSMResultRef, n Index) DataRef {
	return _LSMResultCopyToken(result, n)
}/* debug [functions.gen.go/function]: LSMResultCopyToken */

// Returns the cluster of tokens for the n-th best (zero-based) result.
//
// Added in macOS .
// Returns the cluster of tokens for the n-th best (zero-based) result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultCopyTokenCluster(_:_:)
func LSMResultCopyTokenCluster(result LSMResultRef, n Index) ArrayRef {
	return _LSMResultCopyTokenCluster(result, n)
}/* debug [functions.gen.go/function]: LSMResultCopyTokenCluster */

// Returns the word for the n-th best (zero-based) result.
//
// Added in macOS .
// Returns the word for the n-th best (zero-based) result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultCopyWord(_:_:)
func LSMResultCopyWord(result LSMResultRef, n Index) StringRef {
	return _LSMResultCopyWord(result, n)
}/* debug [functions.gen.go/function]: LSMResultCopyWord */

// Returns the cluster of words for the n-th best (zero-based) result.
//
// Added in macOS .
// Returns the cluster of words for the n-th best (zero-based) result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultCopyWordCluster(_:_:)
func LSMResultCopyWordCluster(result LSMResultRef, n Index) ArrayRef {
	return _LSMResultCopyWordCluster(result, n)
}/* debug [functions.gen.go/function]: LSMResultCopyWordCluster */

// Returns the categories or words that best match when a text is mapped into a map, in decreasing order of likelihood.
//
// Added in macOS .
// Returns the categories or words that best match when a text is mapped into a map, in decreasing order of likelihood.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultCreate(_:_:_:_:_:)
func LSMResultCreate(alloc AllocatorRef, mapref LSMMapRef, textref LSMTextRef, numResults Index, flags OptionFlags) LSMResultRef {
	return _LSMResultCreate(alloc, mapref, textref, numResults, flags)
}/* debug [functions.gen.go/function]: LSMResultCreate */

// Returns the category of the specified result.
//
// Added in macOS .
// Returns the category of the specified result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultGetCategory(_:_:)
func LSMResultGetCategory(result LSMResultRef, n Index) LSMCategory {
	return _LSMResultGetCategory(result, n)
}/* debug [functions.gen.go/function]: LSMResultGetCategory */

// Returns the number of results.
//
// Added in macOS .
// Returns the number of results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultGetCount(_:)
func LSMResultGetCount(result LSMResultRef) Index {
	return _LSMResultGetCount(result)
}/* debug [functions.gen.go/function]: LSMResultGetCount */

// Returns the likelihood of the specified result.
//
// Added in macOS .
// Returns the likelihood of the specified result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultGetScore(_:_:)
func LSMResultGetScore(result LSMResultRef, n Index) float32 {
	return _LSMResultGetScore(result, n)
}/* debug [functions.gen.go/function]: LSMResultGetScore */

// Returns the Core Foundation type identifier for Latent Semantic Mapping results.
//
// Added in macOS .
// Returns the Core Foundation type identifier for Latent Semantic Mapping results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMResultGetTypeID()
func LSMResultGetTypeID() TypeID {
	return _LSMResultGetTypeID()
}/* debug [functions.gen.go/function]: LSMResultGetTypeID */

// Adds an arbitrary binary token to the text.
//
// Added in macOS .
// Adds an arbitrary binary token to the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMTextAddToken(_:_:)
func LSMTextAddToken(textref LSMTextRef, token DataRef) unsafe.Pointer {
	return _LSMTextAddToken(textref, token)
}/* debug [functions.gen.go/function]: LSMTextAddToken */

// Adds a word to the text.
//
// Added in macOS .
// Adds a word to the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMTextAddWord(_:_:)
func LSMTextAddWord(textref LSMTextRef, word StringRef) unsafe.Pointer {
	return _LSMTextAddWord(textref, word)
}/* debug [functions.gen.go/function]: LSMTextAddWord */

// Breaks a string into words using the specified locale, and adds the words to the text.
//
// Added in macOS .
// Breaks a string into words using the specified locale, and adds the words to the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMTextAddWords(_:_:_:_:)
func LSMTextAddWords(textref LSMTextRef, words StringRef, locale LocaleRef, flags OptionFlags) unsafe.Pointer {
	return _LSMTextAddWords(textref, words, locale, flags)
}/* debug [functions.gen.go/function]: LSMTextAddWords */

// Creates a new text.
//
// Added in macOS .
// Creates a new text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMTextCreate(_:_:)
func LSMTextCreate(alloc AllocatorRef, mapref LSMMapRef) LSMTextRef {
	return _LSMTextCreate(alloc, mapref)
}/* debug [functions.gen.go/function]: LSMTextCreate */

// Returns the Core Foundation type identifier for Latent Semantic Mapping texts.
//
// Added in macOS .
// Returns the Core Foundation type identifier for Latent Semantic Mapping texts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LatentSemanticMapping/LSMTextGetTypeID()
func LSMTextGetTypeID() TypeID {
	return _LSMTextGetTypeID()
}/* debug [functions.gen.go/function]: LSMTextGetTypeID */




