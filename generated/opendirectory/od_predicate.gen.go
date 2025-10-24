// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class odPredicate */


/* debug [class_header]: Header for odPredicate */
// The class instance for the [odPredicate] class.
var (
	OdPredicateClass     _odPredicateClass
	OdPredicateClassOnce sync.Once
)

func getodPredicateClass() _odPredicateClass {
	OdPredicateClassOnce.Do(func() {
		OdPredicateClass = _odPredicateClass{objc.GetClass("odPredicate")}
	})
	return OdPredicateClass
}

type _odPredicateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for odPredicate */
// An interface definition for the [odPredicate] class.
type IodPredicate interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for odPredicate */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for odPredicate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for odPredicate */
// Alloc allocates a new instance without initialization.
func (oc _odPredicateClass) Alloc() odPredicate {
	rv := objc.Send[odPredicate](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _odPredicateClass) New() odPredicate {
	rv := objc.Send[odPredicate](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ odPredicate) Init() odPredicate {
	rv := objc.Send[odPredicate](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ odPredicate) Autorelease() odPredicate {
	rv := objc.Send[odPredicate](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewodPredicate creates a new odPredicate instance.
func NewodPredicate() odPredicate {
	return getodPredicateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for odPredicate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/odPredicate-c.ivar
type odPredicate struct {
	objectivec.Object
}

// odPredicateFrom constructs a [odPredicate] from an unsafe.Pointer.
func odPredicateFrom(ptr unsafe.Pointer) odPredicate {
	return odPredicate{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for odPredicate *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for odPredicate */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for odPredicate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for odPredicate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for odPredicate */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class odPredicate */



