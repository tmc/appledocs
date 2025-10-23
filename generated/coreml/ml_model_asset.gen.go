// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelAsset] class.
var (
	ModelAssetClass     _ModelAssetClass
	ModelAssetClassOnce sync.Once
)

func getModelAssetClass() _ModelAssetClass {
	ModelAssetClassOnce.Do(func() {
		ModelAssetClass = _ModelAssetClass{objc.GetClass("MLModelAsset")}
	})
	return ModelAssetClass
}

type _ModelAssetClass struct {
	class objc.Class
}

// An interface definition for the [ModelAsset] class.
type IModelAsset interface {
	objectivec.IObject
	// properties:
	// methods:
	FunctionNamesWithCompletionHandler(handler unsafe.Pointer)
	ModelDescriptionWithCompletionHandler(handler unsafe.Pointer)
	ModelDescriptionOfFunctionNamedCompletionHandler(functionName string /* primitive/slice/pointer. */, handler unsafe.Pointer)
}

// An abstraction of a compiled Core ML model asset.
//
// provides a unified interface by abstracting the compiled model representations for files and in-memory representations. To use an in-memory model, create an with an in-memory model specification, then call .


// An abstraction of a compiled Core ML model asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset
type ModelAsset struct {
	objectivec.Object
}

// ModelAssetFrom constructs a [ModelAsset] from an unsafe.Pointer.
//
// An abstraction of a compiled Core ML model asset.
func ModelAssetFrom(ptr unsafe.Pointer) ModelAsset {
	return ModelAsset{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelAssetClass) Alloc() ModelAsset {
	rv := objc.Send[ModelAsset](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelAssetClass) New() ModelAsset {
	rv := objc.Send[ModelAsset](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelAsset) Init() ModelAsset {
	rv := objc.Send[ModelAsset](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelAsset) Autorelease() ModelAsset {
	rv := objc.Send[ModelAsset](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelAsset creates a new ModelAsset instance.
func NewModelAsset() ModelAsset {
	return getModelAssetClass().New()
}



// Construct a model asset from an ML Program specification by replacing blob file references with corresponding in-memory blobs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(specification:blobMapping:)
func NewModelAssetWithSpecificationDataBlobMappingError(specificationData foundation.objc.IObject /* cross-framework NSData */, blobMapping foundation.IDictionary /* already interface */, error_ unsafe.Pointer) ModelAsset {
	rv := objc.Send[ModelAsset](objc.ID(getModelAssetClass().class), objc.Sel("modelAssetWithSpecificationData:blobMapping:error:"), specificationData, blobMapping, error_)
	return rv
}


// Creates a model asset from an in-memory model specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(specification:)
func NewModelAssetWithSpecificationDataError(specificationData foundation.objc.IObject /* cross-framework NSData */, error_ unsafe.Pointer) ModelAsset {
	rv := objc.Send[ModelAsset](objc.ID(getModelAssetClass().class), objc.Sel("modelAssetWithSpecificationData:error:"), specificationData, error_)
	return rv
}


// Constructs a ModelAsset from a compiled model URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(url:)
func NewModelAssetWithURLError(compiledModelURL foundation.objc.IObject /* cross-framework URL */, error_ unsafe.Pointer) ModelAsset {
	rv := objc.Send[ModelAsset](objc.ID(getModelAssetClass().class), objc.Sel("modelAssetWithURL:error:"), compiledModelURL, error_)
	return rv
}



// Creates a model asset from an in-memory model specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(specification:)
func (mc _ModelAssetClass) ModelAssetWithSpecificationDataError(specificationData foundation.objc.IObject /* cross-framework NSData */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("modelAssetWithSpecificationData:error:"), specificationData, error_)
	return rv
}


// Construct a model asset from an ML Program specification by replacing blob file references with corresponding in-memory blobs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(specification:blobMapping:)
func (mc _ModelAssetClass) ModelAssetWithSpecificationDataBlobMappingError(specificationData foundation.objc.IObject /* cross-framework NSData */, blobMapping foundation.IDictionary /* already interface */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("modelAssetWithSpecificationData:blobMapping:error:"), specificationData, blobMapping, error_)
	return rv
}


// Constructs a ModelAsset from a compiled model URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(url:)
func (mc _ModelAssetClass) ModelAssetWithURLError(compiledModelURL foundation.objc.IObject /* cross-framework URL */, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("modelAssetWithURL:error:"), compiledModelURL, error_)
	return rv
}


// The list of function names in the model asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/functionNames(completionHandler:)
func (m_ ModelAsset) FunctionNamesWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("functionNamesWithCompletionHandler:"), handler)
}


// The default model descripton.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelDescription(completionHandler:)
func (m_ ModelAsset) ModelDescriptionWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("modelDescriptionWithCompletionHandler:"), handler)
}


// The model descripton for a specified function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelDescription(ofFunctionNamed:completionHandler:)
func (m_ ModelAsset) ModelDescriptionOfFunctionNamedCompletionHandler(functionName string /* primitive/slice/pointer. */, handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("modelDescriptionOfFunctionNamed:completionHandler:"), objc.String(functionName), handler)
}


