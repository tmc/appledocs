// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class isa */


/* debug [class_header]: Header for isa */
// The class instance for the [ISA] class.
var (
	IsaClass     _ISAClass
	IsaClassOnce sync.Once
)

func getISAClass() _ISAClass {
	IsaClassOnce.Do(func() {
		IsaClass = _ISAClass{objc.GetClass("isa")}
	})
	return IsaClass
}

type _ISAClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ISA */
// An interface definition for the [ISA] class.
type IISA interface {
	IObject
	
/* debug [class_interface_properties]: Properties for ISA */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ISA */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ISA */
// Alloc allocates a new instance without initialization.
func (ic _ISAClass) Alloc() ISA {
	rv := objc.Send[ISA](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ISAClass) New() ISA {
	rv := objc.Send[ISA](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ISA) Init() ISA {
	rv := objc.Send[ISA](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ISA) Autorelease() ISA {
	rv := objc.Send[ISA](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewISA creates a new ISA instance.
func NewISA() ISA {
	return getISAClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ISA */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isa
type ISA struct {
	Object
}

// ISAFrom constructs a [ISA] from an unsafe.Pointer.
func ISAFrom(ptr unsafe.Pointer) ISA {
	return ISA{Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ISA *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ISA */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ISA */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ISA */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ISA */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class isa */



