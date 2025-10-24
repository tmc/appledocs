// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TriangleAccelerationStructure] class.
var (
	TriangleAccelerationStructureClass     _TriangleAccelerationStructureClass
	TriangleAccelerationStructureClassOnce sync.Once
)

func getTriangleAccelerationStructureClass() _TriangleAccelerationStructureClass {
	TriangleAccelerationStructureClassOnce.Do(func() {
		TriangleAccelerationStructureClass = _TriangleAccelerationStructureClass{objc.GetClass("MPSTriangleAccelerationStructure")}
	})
	return TriangleAccelerationStructureClass
}

type _TriangleAccelerationStructureClass struct {
	class objc.Class
}

// An interface definition for the [TriangleAccelerationStructure] class.
type ITriangleAccelerationStructure interface {
	IPolygonAccelerationStructure
	// properties:
	TriangleCount() uint
	SetTriangleCount(value uint)
	// methods:
}

// An acceleration structure built over triangles.


// An acceleration structure built over triangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTriangleAccelerationStructure
type TriangleAccelerationStructure struct {
	PolygonAccelerationStructure
}

// TriangleAccelerationStructureFrom constructs a [TriangleAccelerationStructure] from an unsafe.Pointer.
//
// An acceleration structure built over triangles.
func TriangleAccelerationStructureFrom(ptr unsafe.Pointer) TriangleAccelerationStructure {
	return TriangleAccelerationStructure{
		PolygonAccelerationStructure: PolygonAccelerationStructureFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TriangleAccelerationStructureClass) Alloc() TriangleAccelerationStructure {
	rv := objc.Send[TriangleAccelerationStructure](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TriangleAccelerationStructureClass) New() TriangleAccelerationStructure {
	rv := objc.Send[TriangleAccelerationStructure](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TriangleAccelerationStructure) Init() TriangleAccelerationStructure {
	rv := objc.Send[TriangleAccelerationStructure](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TriangleAccelerationStructure) Autorelease() TriangleAccelerationStructure {
	rv := objc.Send[TriangleAccelerationStructure](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTriangleAccelerationStructure creates a new TriangleAccelerationStructure instance.
func NewTriangleAccelerationStructure() TriangleAccelerationStructure {
	return getTriangleAccelerationStructureClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTriangleAccelerationStructure/triangleCount
func (t_ TriangleAccelerationStructure) TriangleCount() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("triangleCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSTriangleAccelerationStructure/triangleCount
func (t_ TriangleAccelerationStructure) SetTriangleCount(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTriangleCount:"), value)
}



