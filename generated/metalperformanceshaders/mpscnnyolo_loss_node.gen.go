// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNYOLOLossNode] class.
var (
	CNNYOLOLossNodeClass     _CNNYOLOLossNodeClass
	CNNYOLOLossNodeClassOnce sync.Once
)

func getCNNYOLOLossNodeClass() _CNNYOLOLossNodeClass {
	CNNYOLOLossNodeClassOnce.Do(func() {
		CNNYOLOLossNodeClass = _CNNYOLOLossNodeClass{objc.GetClass("MPSCNNYOLOLossNode")}
	})
	return CNNYOLOLossNodeClass
}

type _CNNYOLOLossNodeClass struct {
	class objc.Class
}





// An interface definition for the [CNNYOLOLossNode] class.
type ICNNYOLOLossNode interface {
	IFilterNode
	

	// properties:
	InputLabels() IMPSNNLabelsNode
	SetInputLabels(value IMPSNNLabelsNode)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNYOLOLossNodeClass) Alloc() CNNYOLOLossNode {
	rv := objc.Send[CNNYOLOLossNode](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNYOLOLossNodeClass) New() CNNYOLOLossNode {
	rv := objc.Send[CNNYOLOLossNode](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNYOLOLossNode) Init() CNNYOLOLossNode {
	rv := objc.Send[CNNYOLOLossNode](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNYOLOLossNode) Autorelease() CNNYOLOLossNode {
	rv := objc.Send[CNNYOLOLossNode](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNYOLOLossNode creates a new CNNYOLOLossNode instance.
func NewCNNYOLOLossNode() CNNYOLOLossNode {
	return getCNNYOLOLossNodeClass().New()
}





// A representation of a YOLO loss kernel.


// A representation of a YOLO loss kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNYOLOLossNode
type CNNYOLOLossNode struct {
	FilterNode
}

// CNNYOLOLossNodeFrom constructs a [CNNYOLOLossNode] from an unsafe.Pointer.
//
// A representation of a YOLO loss kernel.
func CNNYOLOLossNodeFrom(ptr unsafe.Pointer) CNNYOLOLossNode {
	return CNNYOLOLossNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossnode/2976514-initwithsource
func NewCNNYOLOLossNodeWithSourceLossDescriptor(source IImageNode, descriptor ICNNYOLOLossDescriptor) CNNYOLOLossNode {
	instance := getCNNYOLOLossNodeClass().Alloc()
	rv := objc.Send[CNNYOLOLossNode](instance.ID, objc.Sel("initWithSource:lossDescriptor:"), source, descriptor)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossnode/2976516-nodewithsource
func (cc _CNNYOLOLossNodeClass) NodeWithSourceLossDescriptor(source IImageNode, descriptor ICNNYOLOLossDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("nodeWithSource:lossDescriptor:"), source, descriptor)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossnode/2976515-inputlabels
func (c_ CNNYOLOLossNode) InputLabels() IMPSNNLabelsNode {
	rv := objc.Send[LabelsNode](c_.ID, objc.Sel("inputLabels"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossnode/2976515-inputlabels
func (c_ CNNYOLOLossNode) SetInputLabels(value IMPSNNLabelsNode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInputLabels:"), value)
}







