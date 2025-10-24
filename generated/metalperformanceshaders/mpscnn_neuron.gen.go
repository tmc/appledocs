// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuron */


/* debug [class_header]: Header for MPSCNNNeuron */
// The class instance for the [CNNNeuron] class.
var (
	CNNNeuronClass     _CNNNeuronClass
	CNNNeuronClassOnce sync.Once
)

func getCNNNeuronClass() _CNNNeuronClass {
	CNNNeuronClassOnce.Do(func() {
		CNNNeuronClass = _CNNNeuronClass{objc.GetClass("MPSCNNNeuron")}
	})
	return CNNNeuronClass
}

type _CNNNeuronClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuron */
// An interface definition for the [CNNNeuron] class.
type ICNNNeuron interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for CNNNeuron */
	// properties:
	A() objectivec.IObject
	SetA(value objectivec.IObject)
	C() objectivec.IObject
	SetC(value objectivec.IObject)
	B() objectivec.IObject
	SetB(value objectivec.IObject)
	Data() objectivec.IObject
	SetData(value objectivec.IObject)
	NeuronType() CNNNeuronType get /* not a class type */
	SetNeuronType(value CNNNeuronType get /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuron */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuron */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronClass) Alloc() CNNNeuron {
	rv := objc.Send[CNNNeuron](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronClass) New() CNNNeuron {
	rv := objc.Send[CNNNeuron](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuron) Init() CNNNeuron {
	rv := objc.Send[CNNNeuron](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuron) Autorelease() CNNNeuron {
	rv := objc.Send[CNNNeuron](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuron creates a new CNNNeuron instance.
func NewCNNNeuron() CNNNeuron {
	return getCNNNeuronClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuron */
// A filter that applies a neuron activation function.
//
// Do not use this class directly; use one of the subclasses instead.


// A filter that applies a neuron activation function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuron
type CNNNeuron struct {
	CNNKernel
}

// CNNNeuronFrom constructs a [CNNNeuron] from an unsafe.Pointer.
//
// A filter that applies a neuron activation function.
func CNNNeuronFrom(ptr unsafe.Pointer) CNNNeuron {
	return CNNNeuron{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuron */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2867058-initwithcoder
func NewCNNNeuronWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNNeuron {
	instance := getCNNNeuronClass().Alloc()
	rv := objc.Send[CNNNeuron](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice
func NewCNNNeuronWithDeviceNeuronDescriptor(device unsafe.Pointer, neuronDescriptor INeuronDescriptor) CNNNeuron {
	instance := getCNNNeuronClass().Alloc()
	rv := objc.Send[CNNNeuron](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronWithDeviceNeuronDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuron */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuron */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuron */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuron */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942297-a
func (c_ CNNNeuron) A() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("a"))
	return rv
}/* debug [instance_properties/getter]: a */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942297-a
func (c_ CNNNeuron) SetA(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setA:"), value)
}/* debug [instance_properties/setter]: a */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942303-c
func (c_ CNNNeuron) C() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("c"))
	return rv
}/* debug [instance_properties/getter]: c */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942303-c
func (c_ CNNNeuron) SetC(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setC:"), value)
}/* debug [instance_properties/setter]: c */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942306-b
func (c_ CNNNeuron) B() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("b"))
	return rv
}/* debug [instance_properties/getter]: b */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942306-b
func (c_ CNNNeuron) SetB(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setB:"), value)
}/* debug [instance_properties/setter]: b */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942308-data
func (c_ CNNNeuron) Data() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942308-data
func (c_ CNNNeuron) SetData(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942309-neurontype
func (c_ CNNNeuron) NeuronType() CNNNeuronType get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("neuronType"))
	return rv
}/* debug [instance_properties/getter]: neuronType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942309-neurontype
func (c_ CNNNeuron) SetNeuronType(value CNNNeuronType get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNeuronType:"), value)
}/* debug [instance_properties/setter]: neuronType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuron */


