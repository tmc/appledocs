// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSPredicate */


/* debug [class_header]: Header for MPSPredicate */
// The class instance for the [Predicate] class.
var (
	PredicateClass     _PredicateClass
	PredicateClassOnce sync.Once
)

func getPredicateClass() _PredicateClass {
	PredicateClassOnce.Do(func() {
		PredicateClass = _PredicateClass{objc.GetClass("MPSPredicate")}
	})
	return PredicateClass
}

type _PredicateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Predicate */
// An interface definition for the [Predicate] class.
type IPredicate interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Predicate */
	// properties:
	PredicateBuffer() Buffer get /* not a class type */
	SetPredicateBuffer(value Buffer get /* not a class type */)
	PredicateOffset() objectivec.IObject
	SetPredicateOffset(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Predicate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Predicate */
// Alloc allocates a new instance without initialization.
func (pc _PredicateClass) Alloc() Predicate {
	rv := objc.Send[Predicate](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PredicateClass) New() Predicate {
	rv := objc.Send[Predicate](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Predicate) Init() Predicate {
	rv := objc.Send[Predicate](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Predicate) Autorelease() Predicate {
	rv := objc.Send[Predicate](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPredicate creates a new Predicate instance.
func NewPredicate() Predicate {
	return getPredicateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Predicate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPredicate
type Predicate struct {
	objectivec.Object
}

// PredicateFrom constructs a [Predicate] from an unsafe.Pointer.
func PredicateFrom(ptr unsafe.Pointer) Predicate {
	return Predicate{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Predicate */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/3114034-initwithbuffer
func NewPredicateWithBufferOffset(buffer unsafe.Pointer, offset uint) Predicate {
	instance := getPredicateClass().Alloc()
	rv := objc.Send[Predicate](instance.ID, objc.Sel("initWithBuffer:offset:"), buffer, offset)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPredicateWithBufferOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/3114035-initwithdevice
func NewPredicateWithDevice(device unsafe.Pointer) Predicate {
	instance := getPredicateClass().Alloc()
	rv := objc.Send[Predicate](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPredicateWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Predicate */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/3114038-predicatewithbuffer
func (pc _PredicateClass) PredicateWithBufferOffset(buffer unsafe.Pointer, offset uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("predicateWithBuffer:offset:"), buffer, offset)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PredicateWithBufferOffset) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Predicate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Predicate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Predicate */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/3114036-predicatebuffer
func (p_ Predicate) PredicateBuffer() Buffer get /* not a class type */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("predicateBuffer"))
	return rv
}/* debug [instance_properties/getter]: predicateBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/3114036-predicatebuffer
func (p_ Predicate) SetPredicateBuffer(value Buffer get /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPredicateBuffer:"), value)
}/* debug [instance_properties/setter]: predicateBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/3114037-predicateoffset
func (p_ Predicate) PredicateOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("predicateOffset"))
	return rv
}/* debug [instance_properties/getter]: predicateOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/3114037-predicateoffset
func (p_ Predicate) SetPredicateOffset(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPredicateOffset:"), value)
}/* debug [instance_properties/setter]: predicateOffset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSPredicate */


