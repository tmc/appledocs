// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class isDead */


/* debug [class_header]: Header for isDead */
// The class instance for the [isDead] class.
var (
	IsDeadClass     _isDeadClass
	IsDeadClassOnce sync.Once
)

func getisDeadClass() _isDeadClass {
	IsDeadClassOnce.Do(func() {
		IsDeadClass = _isDeadClass{objc.GetClass("isDead")}
	})
	return IsDeadClass
}

type _isDeadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for isDead */
// An interface definition for the [isDead] class.
type IisDead interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for isDead */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for isDead */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for isDead */
// Alloc allocates a new instance without initialization.
func (ic _isDeadClass) Alloc() isDead {
	rv := objc.Send[isDead](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _isDeadClass) New() isDead {
	rv := objc.Send[isDead](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ isDead) Init() isDead {
	rv := objc.Send[isDead](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ isDead) Autorelease() isDead {
	rv := objc.Send[isDead](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewisDead creates a new isDead instance.
func NewisDead() isDead {
	return getisDeadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for isDead */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/isDead
type isDead struct {
	objectivec.Object
}

// isDeadFrom constructs a [isDead] from an unsafe.Pointer.
func isDeadFrom(ptr unsafe.Pointer) isDead {
	return isDead{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for isDead *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for isDead */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for isDead */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for isDead */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for isDead */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class isDead */



