// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CInferenceGraph] class.
var (
	CInferenceGraphClass     _CInferenceGraphClass
	CInferenceGraphClassOnce sync.Once
)

func getCInferenceGraphClass() _CInferenceGraphClass {
	CInferenceGraphClassOnce.Do(func() {
		CInferenceGraphClass = _CInferenceGraphClass{objc.GetClass("MLCInferenceGraph")}
	})
	return CInferenceGraphClass
}

type _CInferenceGraphClass struct {
	class objc.Class
}

// An interface definition for the [CInferenceGraph] class.
type ICInferenceGraph interface {
	ICGraph
	DeviceMemorySize() int
	SetDeviceMemorySize(value int)
}

// An inference graph created from one or more MLCGraph instances plus additional layers added directly to the inference graph.
//
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCInferenceGraph
type CInferenceGraph struct {
	CGraph
}

// CInferenceGraphFrom constructs a [CInferenceGraph] from an unsafe.Pointer.
//
// An inference graph created from one or more MLCGraph instances plus additional layers added directly to the inference graph.
func CInferenceGraphFrom(ptr unsafe.Pointer) CInferenceGraph {
	return CInferenceGraph{
		CGraph: CGraphFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CInferenceGraphClass) Alloc() CInferenceGraph {
	rv := objc.Send[CInferenceGraph](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CInferenceGraphClass) New() CInferenceGraph {
	rv := objc.Send[CInferenceGraph](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CInferenceGraph) Init() CInferenceGraph {
	rv := objc.Send[CInferenceGraph](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CInferenceGraph) Autorelease() CInferenceGraph {
	rv := objc.Send[CInferenceGraph](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCInferenceGraph creates a new CInferenceGraph instance.
func NewCInferenceGraph() CInferenceGraph {
	return getCInferenceGraphClass().New()
}


// The device memory size in bytes for all intermediate tensors in the inference graph.
//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinferencegraph/devicememorysize
func (c_ CInferenceGraph) DeviceMemorySize() int {
	rv := objc.Send[int](c_.ID, objc.Sel("deviceMemorySize"))
	return rv
}


// SetDeviceMemorySize sets the value of the deviceMemorySize property.
// The device memory size in bytes for all intermediate tensors in the inference graph.

//
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlcinferencegraph/devicememorysize
func (c_ CInferenceGraph) SetDeviceMemorySize(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeviceMemorySize:"), value)
}



