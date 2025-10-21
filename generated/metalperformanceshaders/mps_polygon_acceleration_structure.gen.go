// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PolygonAccelerationStructure] class.
var (
	PolygonAccelerationStructureClass     _PolygonAccelerationStructureClass
	PolygonAccelerationStructureClassOnce sync.Once
)

func getPolygonAccelerationStructureClass() _PolygonAccelerationStructureClass {
	PolygonAccelerationStructureClassOnce.Do(func() {
		PolygonAccelerationStructureClass = _PolygonAccelerationStructureClass{objc.GetClass("MPSPolygonAccelerationStructure")}
	})
	return PolygonAccelerationStructureClass
}

type _PolygonAccelerationStructureClass struct {
	class objc.Class
}

// An interface definition for the [PolygonAccelerationStructure] class.
type IPolygonAccelerationStructure interface {
	IAccelerationStructure
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure
type PolygonAccelerationStructure struct {
	AccelerationStructure
}

// PolygonAccelerationStructureFrom constructs a [PolygonAccelerationStructure] from an unsafe.Pointer.
func PolygonAccelerationStructureFrom(ptr unsafe.Pointer) PolygonAccelerationStructure {
	return PolygonAccelerationStructure{
		AccelerationStructure: AccelerationStructureFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PolygonAccelerationStructureClass) Alloc() PolygonAccelerationStructure {
	rv := objc.Send[PolygonAccelerationStructure](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PolygonAccelerationStructureClass) New() PolygonAccelerationStructure {
	rv := objc.Send[PolygonAccelerationStructure](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PolygonAccelerationStructure) Init() PolygonAccelerationStructure {
	rv := objc.Send[PolygonAccelerationStructure](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PolygonAccelerationStructure) Autorelease() PolygonAccelerationStructure {
	rv := objc.Send[PolygonAccelerationStructure](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPolygonAccelerationStructure creates a new PolygonAccelerationStructure instance.
func NewPolygonAccelerationStructure() PolygonAccelerationStructure {
	return getPolygonAccelerationStructureClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/polygonBuffers
func (p_ PolygonAccelerationStructure) PolygonBuffers() []PolygonBuffer {
	rv := objc.Send[[]PolygonBuffer](p_.ID, objc.Sel("polygonBuffers"))
	return rv
}


// SetPolygonBuffers sets the value of the polygonBuffers property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPolygonAccelerationStructure/polygonBuffers
func (p_ PolygonAccelerationStructure) SetPolygonBuffers(value []PolygonBuffer) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setPolygonBuffers:"), nsArray)
}



