// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class automaticTerminationOptOutCounter */


/* debug [class_header]: Header for automaticTerminationOptOutCounter */
// The class instance for the [automaticTerminationOptOutCounter] class.
var (
	AutomaticTerminationOptOutCounterClass     _automaticTerminationOptOutCounterClass
	AutomaticTerminationOptOutCounterClassOnce sync.Once
)

func getautomaticTerminationOptOutCounterClass() _automaticTerminationOptOutCounterClass {
	AutomaticTerminationOptOutCounterClassOnce.Do(func() {
		AutomaticTerminationOptOutCounterClass = _automaticTerminationOptOutCounterClass{objc.GetClass("automaticTerminationOptOutCounter")}
	})
	return AutomaticTerminationOptOutCounterClass
}

type _automaticTerminationOptOutCounterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for automaticTerminationOptOutCounter */
// An interface definition for the [automaticTerminationOptOutCounter] class.
type IautomaticTerminationOptOutCounter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for automaticTerminationOptOutCounter */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for automaticTerminationOptOutCounter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for automaticTerminationOptOutCounter */
// Alloc allocates a new instance without initialization.
func (ac _automaticTerminationOptOutCounterClass) Alloc() automaticTerminationOptOutCounter {
	rv := objc.Send[automaticTerminationOptOutCounter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _automaticTerminationOptOutCounterClass) New() automaticTerminationOptOutCounter {
	rv := objc.Send[automaticTerminationOptOutCounter](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ automaticTerminationOptOutCounter) Init() automaticTerminationOptOutCounter {
	rv := objc.Send[automaticTerminationOptOutCounter](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ automaticTerminationOptOutCounter) Autorelease() automaticTerminationOptOutCounter {
	rv := objc.Send[automaticTerminationOptOutCounter](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewautomaticTerminationOptOutCounter creates a new automaticTerminationOptOutCounter instance.
func NewautomaticTerminationOptOutCounter() automaticTerminationOptOutCounter {
	return getautomaticTerminationOptOutCounterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for automaticTerminationOptOutCounter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProcessInfo/automaticTerminationOptOutCounter
type automaticTerminationOptOutCounter struct {
	objectivec.Object
}

// automaticTerminationOptOutCounterFrom constructs a [automaticTerminationOptOutCounter] from an unsafe.Pointer.
func automaticTerminationOptOutCounterFrom(ptr unsafe.Pointer) automaticTerminationOptOutCounter {
	return automaticTerminationOptOutCounter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for automaticTerminationOptOutCounter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for automaticTerminationOptOutCounter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for automaticTerminationOptOutCounter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for automaticTerminationOptOutCounter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for automaticTerminationOptOutCounter */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class automaticTerminationOptOutCounter */



