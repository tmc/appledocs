// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSAccelerationStructure */


/* debug [class_header]: Header for MPSAccelerationStructure */
// The class instance for the [AccelerationStructure] class.
var (
	AccelerationStructureClass     _AccelerationStructureClass
	AccelerationStructureClassOnce sync.Once
)

func getAccelerationStructureClass() _AccelerationStructureClass {
	AccelerationStructureClassOnce.Do(func() {
		AccelerationStructureClass = _AccelerationStructureClass{objc.GetClass("MPSAccelerationStructure")}
	})
	return AccelerationStructureClass
}

type _AccelerationStructureClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccelerationStructure */
// An interface definition for the [AccelerationStructure] class.
type IAccelerationStructure interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for AccelerationStructure */
	// properties:
	BoundingBox() AxisAlignedBoundingBox get /* not a class type */
	SetBoundingBox(value AxisAlignedBoundingBox get /* not a class type */)
	Group() IMPSAccelerationStructureGroup
	SetGroup(value IMPSAccelerationStructureGroup)
	Status() AccelerationStructureStatus get /* not a class type */
	SetStatus(value AccelerationStructureStatus get /* not a class type */)
	Usage() AccelerationStructureUsage get set /* not a class type */
	SetUsage(value AccelerationStructureUsage get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccelerationStructure */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccelerationStructure */
// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureClass) Alloc() AccelerationStructure {
	rv := objc.Send[AccelerationStructure](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccelerationStructureClass) New() AccelerationStructure {
	rv := objc.Send[AccelerationStructure](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructure) Init() AccelerationStructure {
	rv := objc.Send[AccelerationStructure](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructure) Autorelease() AccelerationStructure {
	rv := objc.Send[AccelerationStructure](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructure creates a new AccelerationStructure instance.
func NewAccelerationStructure() AccelerationStructure {
	return getAccelerationStructureClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccelerationStructure */
// The base class for data structures that are built over geometry and used to accelerate ray tracing.


// The base class for data structures that are built over geometry and used to accelerate ray tracing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure
type AccelerationStructure struct {
	Kernel
}

// AccelerationStructureFrom constructs a [AccelerationStructure] from an unsafe.Pointer.
//
// The base class for data structures that are built over geometry and used to accelerate ray tracing.
func AccelerationStructureFrom(ptr unsafe.Pointer) AccelerationStructure {
	return AccelerationStructure{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccelerationStructure */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructure/2980765-initwithcoder
func NewAccelerationStructureWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) AccelerationStructure {
	instance := getAccelerationStructureClass().Alloc()
	rv := objc.Send[AccelerationStructure](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAccelerationStructureWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructure/2980766-initwithcoder
func NewAccelerationStructureWithCoderGroup(aDecoder Coder /* not a class type */, group IAccelerationStructureGroup) AccelerationStructure {
	instance := getAccelerationStructureClass().Alloc()
	rv := objc.Send[AccelerationStructure](instance.ID, objc.Sel("initWithCoder:group:"), aDecoder, group)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAccelerationStructureWithCoderGroup */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructure/2980767-initwithdevice
func NewAccelerationStructureWithDevice(device unsafe.Pointer) AccelerationStructure {
	instance := getAccelerationStructureClass().Alloc()
	rv := objc.Send[AccelerationStructure](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAccelerationStructureWithDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructure/2980768-initwithgroup
func NewAccelerationStructureWithGroup(group IAccelerationStructureGroup) AccelerationStructure {
	instance := getAccelerationStructureClass().Alloc()
	rv := objc.Send[AccelerationStructure](instance.ID, objc.Sel("initWithGroup:"), group)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAccelerationStructureWithGroup */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccelerationStructure */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccelerationStructure */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccelerationStructure */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccelerationStructure */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructure/2980759-boundingbox
func (a_ AccelerationStructure) BoundingBox() AxisAlignedBoundingBox get /* not a class type */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("boundingBox"))
	return rv
}/* debug [instance_properties/getter]: boundingBox */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructure/2980759-boundingbox
func (a_ AccelerationStructure) SetBoundingBox(value AxisAlignedBoundingBox get /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBoundingBox:"), value)
}/* debug [instance_properties/setter]: boundingBox */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructure/2980764-group
func (a_ AccelerationStructure) Group() IMPSAccelerationStructureGroup {
	rv := objc.Send[AccelerationStructureGroup](a_.ID, objc.Sel("group"))
	return rv
}/* debug [instance_properties/getter]: group */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructure/2980764-group
func (a_ AccelerationStructure) SetGroup(value IMPSAccelerationStructureGroup) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setGroup:"), value)
}/* debug [instance_properties/setter]: group */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructure/2980771-status
func (a_ AccelerationStructure) Status() AccelerationStructureStatus get /* not a class type */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructure/2980771-status
func (a_ AccelerationStructure) SetStatus(value AccelerationStructureStatus get /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructure/2980772-usage
func (a_ AccelerationStructure) Usage() AccelerationStructureUsage get set /* not a class type */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("usage"))
	return rv
}/* debug [instance_properties/getter]: usage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructure/2980772-usage
func (a_ AccelerationStructure) SetUsage(value AccelerationStructureUsage get set /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUsage:"), value)
}/* debug [instance_properties/setter]: usage */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSAccelerationStructure */


