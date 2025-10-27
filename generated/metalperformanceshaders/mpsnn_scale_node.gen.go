// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ScaleNode] class.
var (
	ScaleNodeClass     _ScaleNodeClass
	ScaleNodeClassOnce sync.Once
)

func getScaleNodeClass() _ScaleNodeClass {
	ScaleNodeClassOnce.Do(func() {
		ScaleNodeClass = _ScaleNodeClass{objc.GetClass("MPSNNScaleNode")}
	})
	return ScaleNodeClass
}

type _ScaleNodeClass struct {
	class objc.Class
}





// An interface definition for the [ScaleNode] class.
type IScaleNode interface {
	IFilterNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _ScaleNodeClass) Alloc() ScaleNode {
	rv := objc.Send[ScaleNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScaleNodeClass) New() ScaleNode {
	rv := objc.Send[ScaleNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScaleNode) Init() ScaleNode {
	rv := objc.Send[ScaleNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScaleNode) Autorelease() ScaleNode {
	rv := objc.Send[ScaleNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScaleNode creates a new ScaleNode instance.
func NewScaleNode() ScaleNode {
	return getScaleNodeClass().New()
}





// Abstract node representing an image resampling filter.


// Abstract node representing an image resampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNScaleNode
type ScaleNode struct {
	FilterNode
}

// ScaleNodeFrom constructs a [ScaleNode] from an unsafe.Pointer.
//
// Abstract node representing an image resampling filter.
func ScaleNodeFrom(ptr unsafe.Pointer) ScaleNode {
	return ScaleNode{
		FilterNode: FilterNodeFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915285-initwithsource
func NewScaleNodeWithSourceOutputSize(sourceNode IImageNode, size metal.IMTLSize) ScaleNode {
	instance := getScaleNodeClass().Alloc()
	rv := objc.Send[ScaleNode](instance.ID, objc.Sel("initWithSource:outputSize:"), sourceNode, size)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915278-initwithsource
func NewScaleNodeWithSourceTransformProviderOutputSize(sourceNode IImageNode, transformProvider unsafe.Pointer, size metal.IMTLSize) ScaleNode {
	instance := getScaleNodeClass().Alloc()
	rv := objc.Send[ScaleNode](instance.ID, objc.Sel("initWithSource:transformProvider:outputSize:"), sourceNode, transformProvider, size)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915280-nodewithsource
func (sc _ScaleNodeClass) NodeWithSourceOutputSize(sourceNode IImageNode, size metal.IMTLSize) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("nodeWithSource:outputSize:"), sourceNode, size)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915286-nodewithsource
func (sc _ScaleNodeClass) NodeWithSourceTransformProviderOutputSize(sourceNode IImageNode, transformProvider unsafe.Pointer, size metal.IMTLSize) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("nodeWithSource:transformProvider:outputSize:"), sourceNode, transformProvider, size)
	return rv
}






















