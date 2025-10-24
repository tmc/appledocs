// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CNNNeuronGradient] class.
type ICNNNeuronGradient interface {
	ICNNGradientKernel
	

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


	

	// methods:


}





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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942304-initwithcoder
func NewCNNNeuronGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNNeuronGradient {
	instance := getCNNNeuronGradientClass().Alloc()
	rv := objc.Send[CNNNeuronGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942293-initwithdevice
func NewCNNNeuronGradientWithDeviceNeuronDescriptor(device unsafe.Pointer, neuronDescriptor INeuronDescriptor) CNNNeuronGradient {
	instance := getCNNNeuronGradientClass().Alloc()
	rv := objc.Send[CNNNeuronGradient](instance.ID, objc.Sel("initWithDevice:neuronDescriptor:"), device, neuronDescriptor)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942300-neurontype
func (c_ CNNNeuronGradient) NeuronType() CNNNeuronType get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("neuronType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942300-neurontype
func (c_ CNNNeuronGradient) SetNeuronType(value CNNNeuronType get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNeuronType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942310-c
func (c_ CNNNeuronGradient) C() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("c"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942310-c
func (c_ CNNNeuronGradient) SetC(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setC:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942312-a
func (c_ CNNNeuronGradient) A() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("a"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942312-a
func (c_ CNNNeuronGradient) SetA(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setA:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942313-b
func (c_ CNNNeuronGradient) B() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("b"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942313-b
func (c_ CNNNeuronGradient) SetB(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setB:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942314-data
func (c_ CNNNeuronGradient) Data() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942314-data
func (c_ CNNNeuronGradient) SetData(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setData:"), value)
}







