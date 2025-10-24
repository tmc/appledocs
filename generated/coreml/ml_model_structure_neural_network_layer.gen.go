// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelStructureNeuralNetworkLayer */


/* debug [class_header]: Header for MLModelStructureNeuralNetworkLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelStructureNeuralNetworkLayer */
// An interface definition for the [ModelStructureNeuralNetworkLayer] class.
type IModelStructureNeuralNetworkLayer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelStructureNeuralNetworkLayer */
	// properties:
	InputNames() []string
	Name() objc.IObject /* cross-framework: NSString */
	OutputNames() []string
	Type() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelStructureNeuralNetworkLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelStructureNeuralNetworkLayer */
// Alloc allocates a new instance without initialization.
func (mc _ModelStructureNeuralNetworkLayerClass) Alloc() ModelStructureNeuralNetworkLayer {
	rv := objc.Send[ModelStructureNeuralNetworkLayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelStructureNeuralNetworkLayer */
// A class representing a layer in a NeuralNetwork.


// A class representing a layer in a NeuralNetwork.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelStructureNeuralNetworkLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelStructureNeuralNetworkLayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelStructureNeuralNetworkLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelStructureNeuralNetworkLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelStructureNeuralNetworkLayer */

// The input names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/inputNames
func (m_ ModelStructureNeuralNetworkLayer) InputNames() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("inputNames"))
	return rv
}/* debug [instance_properties/getter]: inputNames */


// The layer name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/name
func (m_ ModelStructureNeuralNetworkLayer) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The output names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/outputNames
func (m_ ModelStructureNeuralNetworkLayer) OutputNames() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("outputNames"))
	return rv
}/* debug [instance_properties/getter]: outputNames */


// The type of the layer, e,g, “elementwise”, “pooling”, etc.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelStructureNeuralNetworkLayer/type
func (m_ ModelStructureNeuralNetworkLayer) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelStructureNeuralNetworkLayer */



