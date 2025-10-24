// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronGradientNode] class.
var (
	CNNNeuronGradientNodeClass     _CNNNeuronGradientNodeClass
	CNNNeuronGradientNodeClassOnce sync.Once
)

func getCNNNeuronGradientNodeClass() _CNNNeuronGradientNodeClass {
	CNNNeuronGradientNodeClassOnce.Do(func() {
		CNNNeuronGradientNodeClass = _CNNNeuronGradientNodeClass{objc.GetClass("MPSCNNNeuronGradientNode")}
	})
	return CNNNeuronGradientNodeClass
}

type _CNNNeuronGradientNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronGradientNode] class.
type ICNNNeuronGradientNode interface {
	IGradientFilterNode
	

	// properties:
	Descriptor() IMPSNNNeuronDescriptor
	SetDescriptor(value IMPSNNNeuronDescriptor)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronGradientNodeClass) Alloc() CNNNeuronGradientNode {
	rv := objc.Send[CNNNeuronGradientNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronGradientNodeClass) New() CNNNeuronGradientNode {
	rv := objc.Send[CNNNeuronGradientNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronGradientNode) Init() CNNNeuronGradientNode {
	rv := objc.Send[CNNNeuronGradientNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronGradientNode) Autorelease() CNNNeuronGradientNode {
	rv := objc.Send[CNNNeuronGradientNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronGradientNode creates a new CNNNeuronGradientNode instance.
func NewCNNNeuronGradientNode() CNNNeuronGradientNode {
	return getCNNNeuronGradientNodeClass().New()
}





// A representation of a gradient exponential neuron filter.


// A representation of a gradient exponential neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronGradientNode
type CNNNeuronGradientNode struct {
	GradientFilterNode
}

// CNNNeuronGradientNodeFrom constructs a [CNNNeuronGradientNode] from an unsafe.Pointer.
//
// A representation of a gradient exponential neuron filter.
func CNNNeuronGradientNodeFrom(ptr unsafe.Pointer) CNNNeuronGradientNode {
	return CNNNeuronGradientNode{
		GradientFilterNode: GradientFilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradientnode/2948031-initwithsourcegradient
func NewCNNNeuronGradientNodeWithSourceGradientSourceImageGradientStateDescriptor(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, descriptor INeuronDescriptor) CNNNeuronGradientNode {
	instance := getCNNNeuronGradientNodeClass().Alloc()
	rv := objc.Send[CNNNeuronGradientNode](instance.ID, objc.Sel("initWithSourceGradient:sourceImage:gradientState:descriptor:"), sourceGradient, sourceImage, gradientState, descriptor)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradientnode/2948028-nodewithsourcegradient
func (cc _CNNNeuronGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateDescriptor(sourceGradient IImageNode, sourceImage IImageNode, gradientState IGradientStateNode, descriptor INeuronDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:descriptor:"), sourceGradient, sourceImage, gradientState, descriptor)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradientnode/2948040-descriptor
func (c_ CNNNeuronGradientNode) Descriptor() IMPSNNNeuronDescriptor {
	rv := objc.Send[NeuronDescriptor](c_.ID, objc.Sel("descriptor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradientnode/2948040-descriptor
func (c_ CNNNeuronGradientNode) SetDescriptor(value IMPSNNNeuronDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDescriptor:"), value)
}







