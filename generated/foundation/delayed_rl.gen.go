// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class delayedRL */


/* debug [class_header]: Header for delayedRL */
// The class instance for the [delayedRL] class.
var (
	DelayedRLClass     _delayedRLClass
	DelayedRLClassOnce sync.Once
)

func getdelayedRLClass() _delayedRLClass {
	DelayedRLClassOnce.Do(func() {
		DelayedRLClass = _delayedRLClass{objc.GetClass("delayedRL")}
	})
	return DelayedRLClass
}

type _delayedRLClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for delayedRL */
// An interface definition for the [delayedRL] class.
type IdelayedRL interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for delayedRL */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for delayedRL */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for delayedRL */
// Alloc allocates a new instance without initialization.
func (dc _delayedRLClass) Alloc() delayedRL {
	rv := objc.Send[delayedRL](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _delayedRLClass) New() delayedRL {
	rv := objc.Send[delayedRL](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ delayedRL) Init() delayedRL {
	rv := objc.Send[delayedRL](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ delayedRL) Autorelease() delayedRL {
	rv := objc.Send[delayedRL](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewdelayedRL creates a new delayedRL instance.
func NewdelayedRL() delayedRL {
	return getdelayedRLClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for delayedRL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/delayedRL
type delayedRL struct {
	objectivec.Object
}

// delayedRLFrom constructs a [delayedRL] from an unsafe.Pointer.
func delayedRLFrom(ptr unsafe.Pointer) delayedRL {
	return delayedRL{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for delayedRL *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for delayedRL */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for delayedRL */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for delayedRL */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for delayedRL */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class delayedRL */



