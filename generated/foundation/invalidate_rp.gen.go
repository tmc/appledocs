// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class invalidateRP */


/* debug [class_header]: Header for invalidateRP */
// The class instance for the [invalidateRP] class.
var (
	InvalidateRPClass     _invalidateRPClass
	InvalidateRPClassOnce sync.Once
)

func getinvalidateRPClass() _invalidateRPClass {
	InvalidateRPClassOnce.Do(func() {
		InvalidateRPClass = _invalidateRPClass{objc.GetClass("invalidateRP")}
	})
	return InvalidateRPClass
}

type _invalidateRPClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for invalidateRP */
// An interface definition for the [invalidateRP] class.
type IinvalidateRP interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for invalidateRP */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for invalidateRP */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for invalidateRP */
// Alloc allocates a new instance without initialization.
func (ic _invalidateRPClass) Alloc() invalidateRP {
	rv := objc.Send[invalidateRP](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _invalidateRPClass) New() invalidateRP {
	rv := objc.Send[invalidateRP](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ invalidateRP) Init() invalidateRP {
	rv := objc.Send[invalidateRP](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ invalidateRP) Autorelease() invalidateRP {
	rv := objc.Send[invalidateRP](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewinvalidateRP creates a new invalidateRP instance.
func NewinvalidateRP() invalidateRP {
	return getinvalidateRPClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for invalidateRP */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/invalidateRP
type invalidateRP struct {
	objectivec.Object
}

// invalidateRPFrom constructs a [invalidateRP] from an unsafe.Pointer.
func invalidateRPFrom(ptr unsafe.Pointer) invalidateRP {
	return invalidateRP{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for invalidateRP *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for invalidateRP */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for invalidateRP */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for invalidateRP */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for invalidateRP */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class invalidateRP */



