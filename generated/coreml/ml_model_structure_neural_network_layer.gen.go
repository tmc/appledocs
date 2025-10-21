// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelStructureNeuralNetworkLayer] class.
var (
	ModelStructureNeuralNetworkLayerClass     _ModelStructureNeuralNetworkLayerClass
	ModelStructureNeuralNetworkLayerClassOnce sync.Once
)

func getModelStructureNeuralNetworkLayerClass() _ModelStructureNeuralNetworkLayerClass {
	ModelStructureNeuralNetworkLayerClassOnce.Do(func() {
		ModelStructureNeuralNetworkLayerClass = _ModelStructureNeuralNetworkLayerClass{objc.GetClass("MLModelStructureNeuralNetworkLayer")}
	})
	return ModelStructureNeuralNetworkLayerClass
}

type _ModelStructureNeuralNetworkLayerClass struct {
	class objc.Class
}

// An interface definition for the [ModelStructureNeuralNetworkLayer] class.
type IModelStructureNeuralNetworkLayer interface {
	objectivec.IObject
}

// A class representing a layer in a NeuralNetwork.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer
type ModelStructureNeuralNetworkLayer struct {
	objectivec.Object
}

// ModelStructureNeuralNetworkLayerFrom constructs a [ModelStructureNeuralNetworkLayer] from an unsafe.Pointer.
//
// A class representing a layer in a NeuralNetwork.
func ModelStructureNeuralNetworkLayerFrom(ptr unsafe.Pointer) ModelStructureNeuralNetworkLayer {
	return ModelStructureNeuralNetworkLayer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelStructureNeuralNetworkLayerClass) Alloc() ModelStructureNeuralNetworkLayer {
	rv := objc.Send[ModelStructureNeuralNetworkLayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelStructureNeuralNetworkLayerClass) New() ModelStructureNeuralNetworkLayer {
	rv := objc.Send[ModelStructureNeuralNetworkLayer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureNeuralNetworkLayer) Init() ModelStructureNeuralNetworkLayer {
	rv := objc.Send[ModelStructureNeuralNetworkLayer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureNeuralNetworkLayer) Autorelease() ModelStructureNeuralNetworkLayer {
	rv := objc.Send[ModelStructureNeuralNetworkLayer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureNeuralNetworkLayer creates a new ModelStructureNeuralNetworkLayer instance.
func NewModelStructureNeuralNetworkLayer() ModelStructureNeuralNetworkLayer {
	return getModelStructureNeuralNetworkLayerClass().New()
}


// The input names.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/inputNames
func (m_ ModelStructureNeuralNetworkLayer) InputNames() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("inputNames"))
	return rv
}

// The layer name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/name
func (m_ ModelStructureNeuralNetworkLayer) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}

// The output names.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/outputNames
func (m_ ModelStructureNeuralNetworkLayer) OutputNames() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("outputNames"))
	return rv
}

// The type of the layer, e,g, “elementwise”, “pooling”, etc.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/type
func (m_ ModelStructureNeuralNetworkLayer) Type() string {
	rv := objc.Send[string](m_.ID, objc.Sel("type"))
	return rv
}



