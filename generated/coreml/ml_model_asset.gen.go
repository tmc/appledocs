// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelAsset */


/* debug [class_header]: Header for MLModelAsset */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelAsset */
// An interface definition for the [ModelAsset] class.
type IModelAsset interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelAsset */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelAsset */
	// methods:
	FunctionNamesWithCompletionHandler(handler unsafe.Pointer)
	ModelDescriptionWithCompletionHandler(handler unsafe.Pointer)
	ModelDescriptionOfFunctionNamedCompletionHandler(functionName objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelAsset */
// Alloc allocates a new instance without initialization.
func (mc _ModelAssetClass) Alloc() ModelAsset {
	rv := objc.Send[ModelAsset](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelAsset */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelAsset */

// Construct a model asset from an ML Program specification by replacing blob file references with corresponding in-memory blobs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(specification:blobMapping:)
func NewModelAssetWithSpecificationDataBlobMappingError(specificationData objc.IObject /* cross-framework: NSData */, blobMapping foundation.IDictionary, error_ objectivec.IObject) ModelAsset {
	rv := objc.Send[ModelAsset](objc.ID(getModelAssetClass().class), objc.Sel("modelAssetWithSpecificationData:blobMapping:error:"), specificationData, blobMapping, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewModelAssetWithSpecificationDataBlobMappingError */


// Creates a model asset from an in-memory model specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(specification:)
func NewModelAssetWithSpecificationDataError(specificationData objc.IObject /* cross-framework: NSData */, error_ objectivec.IObject) ModelAsset {
	rv := objc.Send[ModelAsset](objc.ID(getModelAssetClass().class), objc.Sel("modelAssetWithSpecificationData:error:"), specificationData, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewModelAssetWithSpecificationDataError */


// Constructs a ModelAsset from a compiled model URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(url:)
func NewModelAssetWithURLError(compiledModelURL objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) ModelAsset {
	rv := objc.Send[ModelAsset](objc.ID(getModelAssetClass().class), objc.Sel("modelAssetWithURL:error:"), compiledModelURL, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewModelAssetWithURLError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelAsset */

// Creates a model asset from an in-memory model specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(specification:)
func (mc _ModelAssetClass) ModelAssetWithSpecificationDataError(specificationData objc.IObject /* cross-framework: NSData */, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("modelAssetWithSpecificationData:error:"), specificationData, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ModelAssetWithSpecificationDataError) */


// Construct a model asset from an ML Program specification by replacing blob file references with corresponding in-memory blobs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(specification:blobMapping:)
func (mc _ModelAssetClass) ModelAssetWithSpecificationDataBlobMappingError(specificationData objc.IObject /* cross-framework: NSData */, blobMapping foundation.IDictionary, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("modelAssetWithSpecificationData:blobMapping:error:"), specificationData, blobMapping, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ModelAssetWithSpecificationDataBlobMappingError) */


// Constructs a ModelAsset from a compiled model URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(url:)
func (mc _ModelAssetClass) ModelAssetWithURLError(compiledModelURL objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("modelAssetWithURL:error:"), compiledModelURL, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ModelAssetWithURLError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelAsset */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelAsset */

// The list of function names in the model asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/functionNames(completionHandler:)
func (m_ ModelAsset) FunctionNamesWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("functionNamesWithCompletionHandler:"), handler)
}/* debug [instance_methods/method]: FunctionNamesWithCompletionHandler */


// The default model descripton.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelDescription(completionHandler:)
func (m_ ModelAsset) ModelDescriptionWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("modelDescriptionWithCompletionHandler:"), handler)
}/* debug [instance_methods/method]: ModelDescriptionWithCompletionHandler */


// The model descripton for a specified function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelDescription(ofFunctionNamed:completionHandler:)
func (m_ ModelAsset) ModelDescriptionOfFunctionNamedCompletionHandler(functionName objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("modelDescriptionOfFunctionNamed:completionHandler:"), functionName, handler)
}/* debug [instance_methods/method]: ModelDescriptionOfFunctionNamedCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelAsset */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelAsset */


