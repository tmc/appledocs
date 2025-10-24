// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelStructureNeuralNetwork */


/* debug [class_header]: Header for MLModelStructureNeuralNetwork */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelStructureNeuralNetwork */
// An interface definition for the [ModelStructureNeuralNetwork] class.
type IModelStructureNeuralNetwork interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelStructureNeuralNetwork */
	// properties:
	Layers() []ModelStructureNeuralNetworkLayer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelStructureNeuralNetwork */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelStructureNeuralNetwork */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelStructureNeuralNetwork */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelStructureNeuralNetwork *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelStructureNeuralNetwork */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelStructureNeuralNetwork */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelStructureNeuralNetwork */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelStructureNeuralNetwork */

// The topologically sorted layers in the NeuralNetwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetwork/layers
func (m_ ModelStructureNeuralNetwork) Layers() []ModelStructureNeuralNetworkLayer {
	rv := objc.Send[[]ModelStructureNeuralNetworkLayer](m_.ID, objc.Sel("layers"))
	return rv
}/* debug [instance_properties/getter]: layers */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelStructureNeuralNetwork */



