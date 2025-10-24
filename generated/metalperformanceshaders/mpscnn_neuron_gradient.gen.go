// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronGradient */


/* debug [class_header]: Header for MPSCNNNeuronGradient */
// The class instance for the [CNNNeuronGradient] class.
var (
	CNNNeuronGradientClass     _CNNNeuronGradientClass
	CNNNeuronGradientClassOnce sync.Once
)

func getCNNNeuronGradientClass() _CNNNeuronGradientClass {
	CNNNeuronGradientClassOnce.Do(func() {
		CNNNeuronGradientClass = _CNNNeuronGradientClass{objc.GetClass("MPSCNNNeuronGradient")}
	})
	return CNNNeuronGradientClass
}

type _CNNNeuronGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronGradient */
// An interface definition for the [CNNNeuronGradient] class.
type ICNNNeuronGradient interface {
	ICNNGradientKernel
	
/* debug [class_interface_properties]: Properties for CNNNeuronGradient */
	// properties:
	NeuronType() CNNNeuronType get /* not a class type */
	SetNeuronType(value CNNNeuronType get /* not a class type */)
	C() objectivec.IObject
	SetC(value objectivec.IObject)
	A() objectivec.IObject
	SetA(value objectivec.IObject)
	B() objectivec.IObject
	SetB(value objectivec.IObject)
	Data() objectivec.IObject
	SetData(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronGradientClass) Alloc() CNNNeuronGradient {
	rv := objc.Send[CNNNeuronGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronGradientClass) New() CNNNeuronGradient {
	rv := objc.Send[CNNNeuronGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronGradient) Init() CNNNeuronGradient {
	rv := objc.Send[CNNNeuronGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronGradient) Autorelease() CNNNeuronGradient {
	rv := objc.Send[CNNNeuronGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronGradient creates a new CNNNeuronGradient instance.
func NewCNNNeuronGradient() CNNNeuronGradient {
	return getCNNNeuronGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronGradient */
// A gradient neuron filter.


// A gradient neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradient
type CNNNeuronGradient struct {
	CNNGradientKernel
}

// CNNNeuronGradientFrom constructs a [CNNNeuronGradient] from an unsafe.Pointer.
//
// A gradient neuron filter.
func CNNNeuronGradientFrom(ptr unsafe.Pointer) CNNNeuronGradient {
	return CNNNeuronGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942304-initwithcoder
func NewCNNNeuronGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNNeuronGradient {
	instance := getCNNNeuronGradientClass().Alloc()
	rv := objc.Send[CNNNeuronGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronGradientWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942293-initwithdevice
func NewCNNNeuronGradientWithDeviceNeuronDescriptor(device unsafe.Pointer, neuronDescriptor INeuronDescriptor) CNNNeuronGradient {
	instance := getCNNNeuronGradientClass().Alloc()
	rv := objc.Send[CNNNeuronGradient](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronGradientWithDeviceNeuronDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942300-neurontype
func (c_ CNNNeuronGradient) NeuronType() CNNNeuronType get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("neuronType"))
	return rv
}/* debug [instance_properties/getter]: neuronType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942300-neurontype
func (c_ CNNNeuronGradient) SetNeuronType(value CNNNeuronType get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNeuronType:"), value)
}/* debug [instance_properties/setter]: neuronType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942310-c
func (c_ CNNNeuronGradient) C() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("c"))
	return rv
}/* debug [instance_properties/getter]: c */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942310-c
func (c_ CNNNeuronGradient) SetC(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setC:"), value)
}/* debug [instance_properties/setter]: c */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942312-a
func (c_ CNNNeuronGradient) A() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("a"))
	return rv
}/* debug [instance_properties/getter]: a */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942312-a
func (c_ CNNNeuronGradient) SetA(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setA:"), value)
}/* debug [instance_properties/setter]: a */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942313-b
func (c_ CNNNeuronGradient) B() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("b"))
	return rv
}/* debug [instance_properties/getter]: b */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942313-b
func (c_ CNNNeuronGradient) SetB(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setB:"), value)
}/* debug [instance_properties/setter]: b */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942314-data
func (c_ CNNNeuronGradient) Data() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942314-data
func (c_ CNNNeuronGradient) SetData(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronGradient */


