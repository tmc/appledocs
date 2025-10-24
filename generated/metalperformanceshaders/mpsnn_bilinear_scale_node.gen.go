// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [BilinearScaleNode] class.
var (
	BilinearScaleNodeClass     _BilinearScaleNodeClass
	BilinearScaleNodeClassOnce sync.Once
)

func getBilinearScaleNodeClass() _BilinearScaleNodeClass {
	BilinearScaleNodeClassOnce.Do(func() {
		BilinearScaleNodeClass = _BilinearScaleNodeClass{objc.GetClass("MPSNNBilinearScaleNode")}
	})
	return BilinearScaleNodeClass
}

type _BilinearScaleNodeClass struct {
	class objc.Class
}





// An interface definition for the [BilinearScaleNode] class.
type IBilinearScaleNode interface {
	IScaleNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (bc _BilinearScaleNodeClass) Alloc() BilinearScaleNode {
	rv := objc.Send[BilinearScaleNode](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BilinearScaleNodeClass) New() BilinearScaleNode {
	rv := objc.Send[BilinearScaleNode](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BilinearScaleNode) Init() BilinearScaleNode {
	rv := objc.Send[BilinearScaleNode](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BilinearScaleNode) Autorelease() BilinearScaleNode {
	rv := objc.Send[BilinearScaleNode](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBilinearScaleNode creates a new BilinearScaleNode instance.
func NewBilinearScaleNode() BilinearScaleNode {
	return getBilinearScaleNodeClass().New()
}





// A representation of a bilinear resampling filter.


// A representation of a bilinear resampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBilinearScaleNode
type BilinearScaleNode struct {
	ScaleNode
}

// BilinearScaleNodeFrom constructs a [BilinearScaleNode] from an unsafe.Pointer.
//
// A representation of a bilinear resampling filter.
func BilinearScaleNodeFrom(ptr unsafe.Pointer) BilinearScaleNode {
	return BilinearScaleNode{
		ScaleNode: ScaleNodeFrom(ptr),
	}
}































