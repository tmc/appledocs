// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mInactivityTimer */


/* debug [class_header]: Header for mInactivityTimer */
// The class instance for the [mInactivityTimer] class.
var (
	MInactivityTimerClass     _mInactivityTimerClass
	MInactivityTimerClassOnce sync.Once
)

func getmInactivityTimerClass() _mInactivityTimerClass {
	MInactivityTimerClassOnce.Do(func() {
		MInactivityTimerClass = _mInactivityTimerClass{objc.GetClass("mInactivityTimer")}
	})
	return MInactivityTimerClass
}

type _mInactivityTimerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mInactivityTimer */
// An interface definition for the [mInactivityTimer] class.
type ImInactivityTimer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mInactivityTimer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mInactivityTimer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mInactivityTimer */
// Alloc allocates a new instance without initialization.
func (mc _mInactivityTimerClass) Alloc() mInactivityTimer {
	rv := objc.Send[mInactivityTimer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mInactivityTimerClass) New() mInactivityTimer {
	rv := objc.Send[mInactivityTimer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mInactivityTimer) Init() mInactivityTimer {
	rv := objc.Send[mInactivityTimer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mInactivityTimer) Autorelease() mInactivityTimer {
	rv := objc.Send[mInactivityTimer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmInactivityTimer creates a new mInactivityTimer instance.
func NewmInactivityTimer() mInactivityTimer {
	return getmInactivityTimerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mInactivityTimer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mInactivityTimer
type mInactivityTimer struct {
	objectivec.Object
}

// mInactivityTimerFrom constructs a [mInactivityTimer] from an unsafe.Pointer.
func mInactivityTimerFrom(ptr unsafe.Pointer) mInactivityTimer {
	return mInactivityTimer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mInactivityTimer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mInactivityTimer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mInactivityTimer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mInactivityTimer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mInactivityTimer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mInactivityTimer */



