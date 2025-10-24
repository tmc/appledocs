// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSTriangleAccelerationStructure */


/* debug [class_header]: Header for MPSTriangleAccelerationStructure */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TriangleAccelerationStructure */
// An interface definition for the [TriangleAccelerationStructure] class.
type ITriangleAccelerationStructure interface {
	IPolygonAccelerationStructure
	
/* debug [class_interface_properties]: Properties for TriangleAccelerationStructure */
	// properties:
	TriangleCount() objectivec.IObject
	SetTriangleCount(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TriangleAccelerationStructure */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TriangleAccelerationStructure */
// Alloc allocates a new instance without initialization.
func (tc _TriangleAccelerationStructureClass) Alloc() TriangleAccelerationStructure {
	rv := objc.Send[TriangleAccelerationStructure](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TriangleAccelerationStructure */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TriangleAccelerationStructure *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TriangleAccelerationStructure */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TriangleAccelerationStructure */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TriangleAccelerationStructure */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TriangleAccelerationStructure */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstriangleaccelerationstructure/2980881-trianglecount
func (t_ TriangleAccelerationStructure) TriangleCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("triangleCount"))
	return rv
}/* debug [instance_properties/getter]: triangleCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstriangleaccelerationstructure/2980881-trianglecount
func (t_ TriangleAccelerationStructure) SetTriangleCount(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTriangleCount:"), value)
}/* debug [instance_properties/setter]: triangleCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSTriangleAccelerationStructure */



