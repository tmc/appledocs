// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRAttributeValueWaiter */


/* debug [class_header]: Header for MTRAttributeValueWaiter */
// The class instance for the [MTRAttributeValueWaiter] class.
var (
	MTRAttributeValueWaiterClass     _MTRAttributeValueWaiterClass
	MTRAttributeValueWaiterClassOnce sync.Once
)

func getMTRAttributeValueWaiterClass() _MTRAttributeValueWaiterClass {
	MTRAttributeValueWaiterClassOnce.Do(func() {
		MTRAttributeValueWaiterClass = _MTRAttributeValueWaiterClass{objc.GetClass("MTRAttributeValueWaiter")}
	})
	return MTRAttributeValueWaiterClass
}

type _MTRAttributeValueWaiterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAttributeValueWaiter */
// An interface definition for the [MTRAttributeValueWaiter] class.
type IMTRAttributeValueWaiter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRAttributeValueWaiter */
	// properties:
	UUID() foundation.UUID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAttributeValueWaiter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAttributeValueWaiter */
// Alloc allocates a new instance without initialization.
func (mc _MTRAttributeValueWaiterClass) Alloc() MTRAttributeValueWaiter {
	rv := objc.Send[MTRAttributeValueWaiter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAttributeValueWaiterClass) New() MTRAttributeValueWaiter {
	rv := objc.Send[MTRAttributeValueWaiter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAttributeValueWaiter) Init() MTRAttributeValueWaiter {
	rv := objc.Send[MTRAttributeValueWaiter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAttributeValueWaiter) Autorelease() MTRAttributeValueWaiter {
	rv := objc.Send[MTRAttributeValueWaiter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAttributeValueWaiter creates a new MTRAttributeValueWaiter instance.
func NewMTRAttributeValueWaiter() MTRAttributeValueWaiter {
	return getMTRAttributeValueWaiterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAttributeValueWaiter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeValueWaiter
type MTRAttributeValueWaiter struct {
	objectivec.Object
}

// MTRAttributeValueWaiterFrom constructs a [MTRAttributeValueWaiter] from an unsafe.Pointer.
func MTRAttributeValueWaiterFrom(ptr unsafe.Pointer) MTRAttributeValueWaiter {
	return MTRAttributeValueWaiter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAttributeValueWaiter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAttributeValueWaiter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAttributeValueWaiter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAttributeValueWaiter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAttributeValueWaiter */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeValueWaiter/uuid
func (m_ MTRAttributeValueWaiter) UUID() foundation.UUID {
	rv := objc.Send[foundation.UUID](m_.ID, objc.Sel("UUID"))
	return rv
}/* debug [instance_properties/getter]: UUID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAttributeValueWaiter */



