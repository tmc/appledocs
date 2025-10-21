// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ModelStructure] class.
type IModelStructure interface {
	objectivec.IObject
}

// A class representing the structure of a model.
//
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

// Alloc allocates a new instance without initialization.
func (mc _ModelStructureClass) Alloc() ModelStructure {
	rv := objc.Send[ModelStructure](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Construct the model structure asynchronously given the location of its on-disk representation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/loadContentsOfURL:completionHandler:
func (mc _ModelStructureClass) LoadContentsOfURLCompletionHandler(url unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("loadContentsOfURL:completionHandler:"), url, handler)
}

// Construct the model structure asynchronously given the model asset.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/loadModelAsset:completionHandler:
func (mc _ModelStructureClass) LoadModelAssetCompletionHandler(asset unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("loadModelAsset:completionHandler:"), asset, handler)
}

// If the model is of NeuralNetwork type then it is the structure of the NeuralNetwork otherwise .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/neuralNetwork
func (m_ ModelStructure) NeuralNetwork() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("neuralNetwork"))
	return rv
}

// If the model is of Pipeline type then it is the structure of the Pipeline otherwise .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/pipeline
func (m_ ModelStructure) Pipeline() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pipeline"))
	return rv
}

// If the model is of ML Program type then it is the structure of the ML Program otherwise .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/program
func (m_ ModelStructure) Program() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("program"))
	return rv
}



