// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [LanczosScaleNode] class.
var (
	LanczosScaleNodeClass     _LanczosScaleNodeClass
	LanczosScaleNodeClassOnce sync.Once
)

func getLanczosScaleNodeClass() _LanczosScaleNodeClass {
	LanczosScaleNodeClassOnce.Do(func() {
		LanczosScaleNodeClass = _LanczosScaleNodeClass{objc.GetClass("MPSNNLanczosScaleNode")}
	})
	return LanczosScaleNodeClass
}

type _LanczosScaleNodeClass struct {
	class objc.Class
}





// An interface definition for the [LanczosScaleNode] class.
type ILanczosScaleNode interface {
	IScaleNode
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (lc _LanczosScaleNodeClass) Alloc() LanczosScaleNode {
	rv := objc.Send[LanczosScaleNode](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LanczosScaleNodeClass) New() LanczosScaleNode {
	rv := objc.Send[LanczosScaleNode](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LanczosScaleNode) Init() LanczosScaleNode {
	rv := objc.Send[LanczosScaleNode](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LanczosScaleNode) Autorelease() LanczosScaleNode {
	rv := objc.Send[LanczosScaleNode](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLanczosScaleNode creates a new LanczosScaleNode instance.
func NewLanczosScaleNode() LanczosScaleNode {
	return getLanczosScaleNodeClass().New()
}





// A representation of a Lanczos resampling filter.


// A representation of a Lanczos resampling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNLanczosScaleNode
type LanczosScaleNode struct {
	ScaleNode
}

// LanczosScaleNodeFrom constructs a [LanczosScaleNode] from an unsafe.Pointer.
//
// A representation of a Lanczos resampling filter.
func LanczosScaleNodeFrom(ptr unsafe.Pointer) LanczosScaleNode {
	return LanczosScaleNode{
		ScaleNode: ScaleNodeFrom(ptr),
	}
}































