// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSQuadrilateralAccelerationStructure */


/* debug [class_header]: Header for MPSQuadrilateralAccelerationStructure */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QuadrilateralAccelerationStructure */
// An interface definition for the [QuadrilateralAccelerationStructure] class.
type IQuadrilateralAccelerationStructure interface {
	IPolygonAccelerationStructure
	
/* debug [class_interface_properties]: Properties for QuadrilateralAccelerationStructure */
	// properties:
	QuadrilateralCount() objectivec.IObject
	SetQuadrilateralCount(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QuadrilateralAccelerationStructure */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QuadrilateralAccelerationStructure */
// Alloc allocates a new instance without initialization.
func (qc _QuadrilateralAccelerationStructureClass) Alloc() QuadrilateralAccelerationStructure {
	rv := objc.Send[QuadrilateralAccelerationStructure](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QuadrilateralAccelerationStructure */


// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QuadrilateralAccelerationStructure *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QuadrilateralAccelerationStructure */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QuadrilateralAccelerationStructure */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QuadrilateralAccelerationStructure */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QuadrilateralAccelerationStructure */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsquadrilateralaccelerationstructure/3088909-quadrilateralcount
func (q_ QuadrilateralAccelerationStructure) QuadrilateralCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](q_.ID, objc.Sel("quadrilateralCount"))
	return rv
}/* debug [instance_properties/getter]: quadrilateralCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsquadrilateralaccelerationstructure/3088909-quadrilateralcount
func (q_ QuadrilateralAccelerationStructure) SetQuadrilateralCount(value objectivec.IObject) {
	objc.Send[objc.ID](q_.ID, objc.Sel("setQuadrilateralCount:"), value)
}/* debug [instance_properties/setter]: quadrilateralCount */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSQuadrilateralAccelerationStructure */



