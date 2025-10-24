// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelStructure */


/* debug [class_header]: Header for MLModelStructure */
// The class instance for the [ModelStructure] class.
var (
	ModelStructureClass     _ModelStructureClass
	ModelStructureClassOnce sync.Once
)

func getModelStructureClass() _ModelStructureClass {
	ModelStructureClassOnce.Do(func() {
		ModelStructureClass = _ModelStructureClass{objc.GetClass("MLModelStructure")}
	})
	return ModelStructureClass
}

type _ModelStructureClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelStructure */
// An interface definition for the [ModelStructure] class.
type IModelStructure interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelStructure */
	// properties:
	NeuralNetwork() IMLModelStructureNeuralNetwork
	Pipeline() IMLModelStructurePipeline
	Program() IMLModelStructureProgram
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelStructure */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelStructure */
// Alloc allocates a new instance without initialization.
func (mc _ModelStructureClass) Alloc() ModelStructure {
	rv := objc.Send[ModelStructure](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _ModelStructureClass) New() ModelStructure {
	rv := objc.Send[ModelStructure](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructure) Init() ModelStructure {
	rv := objc.Send[ModelStructure](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructure) Autorelease() ModelStructure {
	rv := objc.Send[ModelStructure](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructure creates a new ModelStructure instance.
func NewModelStructure() ModelStructure {
	return getModelStructureClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelStructure */
// A class representing the structure of a model.


// A class representing the structure of a model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class
type ModelStructure struct {
	objectivec.Object
}

// ModelStructureFrom constructs a [ModelStructure] from an unsafe.Pointer.
//
// A class representing the structure of a model.
func ModelStructureFrom(ptr unsafe.Pointer) ModelStructure {
	return ModelStructure{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelStructure *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelStructure */

// Construct the model structure asynchronously given the location of its on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/loadContentsOfURL:completionHandler:
func (mc _ModelStructureClass) LoadContentsOfURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("loadContentsOfURL:completionHandler:"), url, handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadContentsOfURLCompletionHandler) */


// Construct the model structure asynchronously given the model asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/loadModelAsset:completionHandler:
func (mc _ModelStructureClass) LoadModelAssetCompletionHandler(asset IMLModelAsset, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("loadModelAsset:completionHandler:"), asset, handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadModelAssetCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelStructure */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelStructure */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelStructure */

// If the model is of NeuralNetwork type then it is the structure of the NeuralNetwork otherwise .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/neuralNetwork
func (m_ ModelStructure) NeuralNetwork() IMLModelStructureNeuralNetwork {
	rv := objc.Send[ModelStructureNeuralNetwork](m_.ID, objc.Sel("neuralNetwork"))
	return rv
}/* debug [instance_properties/getter]: neuralNetwork */


// If the model is of Pipeline type then it is the structure of the Pipeline otherwise .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/pipeline
func (m_ ModelStructure) Pipeline() IMLModelStructurePipeline {
	rv := objc.Send[ModelStructurePipeline](m_.ID, objc.Sel("pipeline"))
	return rv
}/* debug [instance_properties/getter]: pipeline */


// If the model is of ML Program type then it is the structure of the ML Program otherwise .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/program
func (m_ ModelStructure) Program() IMLModelStructureProgram {
	rv := objc.Send[ModelStructureProgram](m_.ID, objc.Sel("program"))
	return rv
}/* debug [instance_properties/getter]: program */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelStructure */



