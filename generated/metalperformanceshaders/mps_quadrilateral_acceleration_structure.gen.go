// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [QuadrilateralAccelerationStructure] class.
var (
	QuadrilateralAccelerationStructureClass     _QuadrilateralAccelerationStructureClass
	QuadrilateralAccelerationStructureClassOnce sync.Once
)

func getQuadrilateralAccelerationStructureClass() _QuadrilateralAccelerationStructureClass {
	QuadrilateralAccelerationStructureClassOnce.Do(func() {
		QuadrilateralAccelerationStructureClass = _QuadrilateralAccelerationStructureClass{objc.GetClass("MPSQuadrilateralAccelerationStructure")}
	})
	return QuadrilateralAccelerationStructureClass
}

type _QuadrilateralAccelerationStructureClass struct {
	class objc.Class
}

// An interface definition for the [QuadrilateralAccelerationStructure] class.
type IQuadrilateralAccelerationStructure interface {
	IPolygonAccelerationStructure
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSQuadrilateralAccelerationStructure
type QuadrilateralAccelerationStructure struct {
	PolygonAccelerationStructure
}

// QuadrilateralAccelerationStructureFrom constructs a [QuadrilateralAccelerationStructure] from an unsafe.Pointer.
func QuadrilateralAccelerationStructureFrom(ptr unsafe.Pointer) QuadrilateralAccelerationStructure {
	return QuadrilateralAccelerationStructure{
		PolygonAccelerationStructure: PolygonAccelerationStructureFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (qc _QuadrilateralAccelerationStructureClass) Alloc() QuadrilateralAccelerationStructure {
	rv := objc.Send[QuadrilateralAccelerationStructure](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QuadrilateralAccelerationStructureClass) New() QuadrilateralAccelerationStructure {
	rv := objc.Send[QuadrilateralAccelerationStructure](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QuadrilateralAccelerationStructure) Init() QuadrilateralAccelerationStructure {
	rv := objc.Send[QuadrilateralAccelerationStructure](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QuadrilateralAccelerationStructure) Autorelease() QuadrilateralAccelerationStructure {
	rv := objc.Send[QuadrilateralAccelerationStructure](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuadrilateralAccelerationStructure creates a new QuadrilateralAccelerationStructure instance.
func NewQuadrilateralAccelerationStructure() QuadrilateralAccelerationStructure {
	return getQuadrilateralAccelerationStructureClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSQuadrilateralAccelerationStructure/quadrilateralCount
func (q_ QuadrilateralAccelerationStructure) QuadrilateralCount() uint {
	rv := objc.Send[uint](q_.ID, objc.Sel("quadrilateralCount"))
	return rv
}


// SetQuadrilateralCount sets the value of the quadrilateralCount property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSQuadrilateralAccelerationStructure/quadrilateralCount
func (q_ QuadrilateralAccelerationStructure) SetQuadrilateralCount(value uint) {
	objc.Send[objc.ID](q_.ID, objc.Sel("setQuadrilateralCount:"), value)
}


