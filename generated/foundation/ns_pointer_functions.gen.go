// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPointerFunctions */


/* debug [class_header]: Header for NSPointerFunctions */
// The class instance for the [PointerFunctions] class.
var (
	PointerFunctionsClass     _PointerFunctionsClass
	PointerFunctionsClassOnce sync.Once
)

func getPointerFunctionsClass() _PointerFunctionsClass {
	PointerFunctionsClassOnce.Do(func() {
		PointerFunctionsClass = _PointerFunctionsClass{objc.GetClass("NSPointerFunctions")}
	})
	return PointerFunctionsClass
}

type _PointerFunctionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PointerFunctions */
// An interface definition for the [PointerFunctions] class.
type IPointerFunctions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PointerFunctions */
	// properties:
	UsesStrongWriteBarrier() bool
	SetUsesStrongWriteBarrier(value bool)
	UsesWeakReadAndWriteBarriers() bool
	SetUsesWeakReadAndWriteBarriers(value bool)
	PointerFunctions() IPointerFunctions
	SetPointerFunctions(value IPointerFunctions)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PointerFunctions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PointerFunctions */
// Alloc allocates a new instance without initialization.
func (pc _PointerFunctionsClass) Alloc() PointerFunctions {
	rv := objc.Send[PointerFunctions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PointerFunctionsClass) New() PointerFunctions {
	rv := objc.Send[PointerFunctions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PointerFunctions) Init() PointerFunctions {
	rv := objc.Send[PointerFunctions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PointerFunctions) Autorelease() PointerFunctions {
	rv := objc.Send[PointerFunctions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPointerFunctions creates a new PointerFunctions instance.
func NewPointerFunctions() PointerFunctions {
	return getPointerFunctionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PointerFunctions */
// An instance of defines callout functions appropriate for managing a pointer reference held somewhere else.
//
// The functions specified by an instance of are separated into two clusters—those that define “personality” such as “object” or “C-string”, and those that describe memory management issues such as a memory deallocation function. There are constants for common personalities and memory manager selections (see ). , , and use an object to define the acquisition and retention behavior for the pointers they manage. Note, however, that not all combinations of personality and memory management behavior are valid for these collections. The pointer collection objects copy the object on input and output, so you cannot usefully subclass .


// An instance of defines callout functions appropriate for managing a pointer reference held somewhere else.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions
type PointerFunctions struct {
	objectivec.Object
}

// PointerFunctionsFrom constructs a [PointerFunctions] from an unsafe.Pointer.
//
// An instance of defines callout functions appropriate for managing a pointer reference held somewhere else.
func PointerFunctionsFrom(ptr unsafe.Pointer) PointerFunctions {
	return PointerFunctions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PointerFunctions */

// Returns an object initialized with the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/init(options:)
func NewPointerFunctionsWithOptions(options PointerFunctionsOptions) PointerFunctions {
	instance := getPointerFunctionsClass().Alloc()
	rv := objc.Send[PointerFunctions](instance.ID, objc.Sel("initWithOptions:"), options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPointerFunctionsWithOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PointerFunctions */

// Returns a new object initialized with the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/pointerFunctionsWithOptions:
func (pc _PointerFunctionsClass) PointerFunctionsWithOptions(options PointerFunctionsOptions) IPointerFunctions {
	rv := objc.Send[PointerFunctions](objc.ID(pc.class), objc.Sel("pointerFunctionsWithOptions:"), options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PointerFunctionsWithOptions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PointerFunctions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PointerFunctions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PointerFunctions */

// Specifies whether, in a garbage collected environment, pointers should be assigned using a strong write barrier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/usesStrongWriteBarrier
func (p_ PointerFunctions) UsesStrongWriteBarrier() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesStrongWriteBarrier"))
	return rv
}/* debug [instance_properties/getter]: usesStrongWriteBarrier */


// Specifies whether, in a garbage collected environment, pointers should be assigned using a strong write barrier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/usesStrongWriteBarrier
func (p_ PointerFunctions) SetUsesStrongWriteBarrier(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesStrongWriteBarrier:"), value)
}/* debug [instance_properties/setter]: usesStrongWriteBarrier */


// Specifies whether, in a garbage collected environment, pointers should use weak read and write barriers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/usesWeakReadAndWriteBarriers
func (p_ PointerFunctions) UsesWeakReadAndWriteBarriers() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesWeakReadAndWriteBarriers"))
	return rv
}/* debug [instance_properties/getter]: usesWeakReadAndWriteBarriers */


// Specifies whether, in a garbage collected environment, pointers should use weak read and write barriers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/usesWeakReadAndWriteBarriers
func (p_ PointerFunctions) SetUsesWeakReadAndWriteBarriers(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesWeakReadAndWriteBarriers:"), value)
}/* debug [instance_properties/setter]: usesWeakReadAndWriteBarriers */


// The pointer functions for the hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/pointerfunctions
func (p_ PointerFunctions) PointerFunctions() IPointerFunctions {
	rv := objc.Send[PointerFunctions](p_.ID, objc.Sel("pointerFunctions"))
	return rv
}/* debug [instance_properties/getter]: pointerFunctions */


// The pointer functions for the hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/pointerfunctions
func (p_ PointerFunctions) SetPointerFunctions(value IPointerFunctions) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPointerFunctions:"), value)
}/* debug [instance_properties/setter]: pointerFunctions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPointerFunctions */


