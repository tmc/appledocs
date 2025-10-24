// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [NeuronDescriptor] class.
var (
	NeuronDescriptorClass     _NeuronDescriptorClass
	NeuronDescriptorClassOnce sync.Once
)

func getNeuronDescriptorClass() _NeuronDescriptorClass {
	NeuronDescriptorClassOnce.Do(func() {
		NeuronDescriptorClass = _NeuronDescriptorClass{objc.GetClass("MPSNNNeuronDescriptor")}
	})
	return NeuronDescriptorClass
}

type _NeuronDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [NeuronDescriptor] class.
type INeuronDescriptor interface {
	objectivec.IObject
	

	// properties:
	NeuronType() CNNNeuronType get set /* not a class type */
	SetNeuronType(value CNNNeuronType get set /* not a class type */)
	Data() objectivec.IObject
	SetData(value objectivec.IObject)
	B() objectivec.IObject
	SetB(value objectivec.IObject)
	C() objectivec.IObject
	SetC(value objectivec.IObject)
	A() objectivec.IObject
	SetA(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (nc _NeuronDescriptorClass) Alloc() NeuronDescriptor {
	rv := objc.Send[NeuronDescriptor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NeuronDescriptorClass) New() NeuronDescriptor {
	rv := objc.Send[NeuronDescriptor](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NeuronDescriptor) Init() NeuronDescriptor {
	rv := objc.Send[NeuronDescriptor](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NeuronDescriptor) Autorelease() NeuronDescriptor {
	rv := objc.Send[NeuronDescriptor](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuronDescriptor creates a new NeuronDescriptor instance.
func NewNeuronDescriptor() NeuronDescriptor {
	return getNeuronDescriptorClass().New()
}





// An object that specifies properties used by a neuron kernel.


// An object that specifies properties used by a neuron kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNNeuronDescriptor
type NeuronDescriptor struct {
	objectivec.Object
}

// NeuronDescriptorFrom constructs a [NeuronDescriptor] from an unsafe.Pointer.
//
// An object that specifies properties used by a neuron kernel.
func NeuronDescriptorFrom(ptr unsafe.Pointer) NeuronDescriptor {
	return NeuronDescriptor{objectivec.Object{objc.ID(ptr)}}
}










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942295-cnnneurondescriptor
func (nc _NeuronDescriptorClass) CnnNeuronDescriptor() {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("cnnNeuronDescriptor"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942295-cnnneurondescriptorwithtype
func (nc _NeuronDescriptorClass) CnnNeuronDescriptorWithTypeAB(neuronType CNNNeuronType, a float32, b float32) INeuronDescriptor {
	rv := objc.Send[NeuronDescriptor](objc.ID(nc.class), objc.Sel("cnnNeuronDescriptorWithType:a:b:"), neuronType, a, b)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942296-cnnneurondescriptorwithtype
func (nc _NeuronDescriptorClass) CnnNeuronDescriptorWithTypeABC(neuronType CNNNeuronType, a float32, b float32, c float32) INeuronDescriptor {
	rv := objc.Send[NeuronDescriptor](objc.ID(nc.class), objc.Sel("cnnNeuronDescriptorWithType:a:b:c:"), neuronType, a, b, c)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942301-cnnneurondescriptorwithtype
func (nc _NeuronDescriptorClass) CnnNeuronDescriptorWithTypeA(neuronType CNNNeuronType, a float32) INeuronDescriptor {
	rv := objc.Send[NeuronDescriptor](objc.ID(nc.class), objc.Sel("cnnNeuronDescriptorWithType:a:"), neuronType, a)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942307-cnnneurondescriptorwithtype
func (nc _NeuronDescriptorClass) CnnNeuronDescriptorWithType(neuronType CNNNeuronType) INeuronDescriptor {
	rv := objc.Send[NeuronDescriptor](objc.ID(nc.class), objc.Sel("cnnNeuronDescriptorWithType:"), neuronType)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942317-cnnneuronpreludescriptor
func (nc _NeuronDescriptorClass) CnnNeuronPReLUDescriptor() {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("cnnNeuronPReLUDescriptor"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942317-cnnneuronpreludescriptorwithdata
func (nc _NeuronDescriptorClass) CnnNeuronPReLUDescriptorWithDataNoCopy(data foundation.Data, noCopy bool) INeuronDescriptor {
	rv := objc.Send[NeuronDescriptor](objc.ID(nc.class), objc.Sel("cnnNeuronPReLUDescriptorWithData:noCopy:"), data, noCopy)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942292-neurontype
func (n_ NeuronDescriptor) NeuronType() CNNNeuronType get set /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("neuronType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942292-neurontype
func (n_ NeuronDescriptor) SetNeuronType(value CNNNeuronType get set /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNeuronType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942299-data
func (n_ NeuronDescriptor) Data() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942299-data
func (n_ NeuronDescriptor) SetData(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942302-b
func (n_ NeuronDescriptor) B() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("b"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942302-b
func (n_ NeuronDescriptor) SetB(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setB:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942305-c
func (n_ NeuronDescriptor) C() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("c"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942305-c
func (n_ NeuronDescriptor) SetC(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setC:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942316-a
func (n_ NeuronDescriptor) A() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("a"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942316-a
func (n_ NeuronDescriptor) SetA(value objectivec.IObject) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setA:"), value)
}








