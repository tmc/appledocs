// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMStepCounter */


/* debug [class_header]: Header for CMStepCounter */
// The class instance for the [StepCounter] class.
var (
	StepCounterClass     _StepCounterClass
	StepCounterClassOnce sync.Once
)

func getStepCounterClass() _StepCounterClass {
	StepCounterClassOnce.Do(func() {
		StepCounterClass = _StepCounterClass{objc.GetClass("CMStepCounter")}
	})
	return StepCounterClass
}

type _StepCounterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StepCounter */
// An interface definition for the [StepCounter] class.
type IStepCounter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for StepCounter */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StepCounter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StepCounter */
// Alloc allocates a new instance without initialization.
func (sc _StepCounterClass) Alloc() StepCounter {
	rv := objc.Send[StepCounter](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StepCounterClass) New() StepCounter {
	rv := objc.Send[StepCounter](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StepCounter) Init() StepCounter {
	rv := objc.Send[StepCounter](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StepCounter) Autorelease() StepCounter {
	rv := objc.Send[StepCounter](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStepCounter creates a new StepCounter instance.
func NewStepCounter() StepCounter {
	return getStepCounterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StepCounter */
// The number of steps the user has taken with the device.
//
// Step information is gathered on devices with the appropriate built-in hardware and stored so that you can run queries to determine the user’s recent physical activity. You use this class to gather both current step data and any historical data.


// The number of steps the user has taken with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMStepCounter
type StepCounter struct {
	objectivec.Object
}

// StepCounterFrom constructs a [StepCounter] from an unsafe.Pointer.
//
// The number of steps the user has taken with the device.
func StepCounterFrom(ptr unsafe.Pointer) StepCounter {
	return StepCounter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StepCounter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StepCounter */

// Returns a Boolean indicating whether step-counting support is available on the current device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMStepCounter/isStepCountingAvailable()
func (sc _StepCounterClass) IsStepCountingAvailable() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("isStepCountingAvailable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IsStepCountingAvailable) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StepCounter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StepCounter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StepCounter */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMStepCounter */


