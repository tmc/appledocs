// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ModelStructureNeuralNetwork] class.
var (
	ModelStructureNeuralNetworkClass     _ModelStructureNeuralNetworkClass
	ModelStructureNeuralNetworkClassOnce sync.Once
)

func getModelStructureNeuralNetworkClass() _ModelStructureNeuralNetworkClass {
	ModelStructureNeuralNetworkClassOnce.Do(func() {
		ModelStructureNeuralNetworkClass = _ModelStructureNeuralNetworkClass{objc.GetClass("MLModelStructureNeuralNetwork")}
	})
	return ModelStructureNeuralNetworkClass
}

type _ModelStructureNeuralNetworkClass struct {
	class objc.Class
}





// An interface definition for the [ModelStructureNeuralNetwork] class.
type IModelStructureNeuralNetwork interface {
	objectivec.IObject
	

	// properties:
	Layers() []ModelStructureNeuralNetworkLayer


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _ModelStructureNeuralNetworkClass) Alloc() ModelStructureNeuralNetwork {
	rv := objc.Send[ModelStructureNeuralNetwork](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _ModelStructureNeuralNetworkClass) New() ModelStructureNeuralNetwork {
	rv := objc.Send[ModelStructureNeuralNetwork](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelStructureNeuralNetwork) Init() ModelStructureNeuralNetwork {
	rv := objc.Send[ModelStructureNeuralNetwork](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelStructureNeuralNetwork) Autorelease() ModelStructureNeuralNetwork {
	rv := objc.Send[ModelStructureNeuralNetwork](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelStructureNeuralNetwork creates a new ModelStructureNeuralNetwork instance.
func NewModelStructureNeuralNetwork() ModelStructureNeuralNetwork {
	return getModelStructureNeuralNetworkClass().New()
}





// A class representing the structure of a NeuralNetwork model.


// A class representing the structure of a NeuralNetwork model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetwork
type ModelStructureNeuralNetwork struct {
	objectivec.Object
}

// ModelStructureNeuralNetworkFrom constructs a [ModelStructureNeuralNetwork] from an unsafe.Pointer.
//
// A class representing the structure of a NeuralNetwork model.
func ModelStructureNeuralNetworkFrom(ptr unsafe.Pointer) ModelStructureNeuralNetwork {
	return ModelStructureNeuralNetwork{objectivec.Object{objc.ID(ptr)}}
}

























// The topologically sorted layers in the NeuralNetwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetwork/layers
func (m_ ModelStructureNeuralNetwork) Layers() []ModelStructureNeuralNetworkLayer {
	rv := objc.Send[[]ModelStructureNeuralNetworkLayer](m_.ID, objc.Sel("layers"))
	return rv
}








