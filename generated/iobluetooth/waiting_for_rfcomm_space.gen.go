// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class waitingForRfcommSpace */


/* debug [class_header]: Header for waitingForRfcommSpace */
// The class instance for the [waitingForRfcommSpace] class.
var (
	WaitingForRfcommSpaceClass     _waitingForRfcommSpaceClass
	WaitingForRfcommSpaceClassOnce sync.Once
)

func getwaitingForRfcommSpaceClass() _waitingForRfcommSpaceClass {
	WaitingForRfcommSpaceClassOnce.Do(func() {
		WaitingForRfcommSpaceClass = _waitingForRfcommSpaceClass{objc.GetClass("waitingForRfcommSpace")}
	})
	return WaitingForRfcommSpaceClass
}

type _waitingForRfcommSpaceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for waitingForRfcommSpace */
// An interface definition for the [waitingForRfcommSpace] class.
type IwaitingForRfcommSpace interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for waitingForRfcommSpace */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for waitingForRfcommSpace */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for waitingForRfcommSpace */
// Alloc allocates a new instance without initialization.
func (wc _waitingForRfcommSpaceClass) Alloc() waitingForRfcommSpace {
	rv := objc.Send[waitingForRfcommSpace](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _waitingForRfcommSpaceClass) New() waitingForRfcommSpace {
	rv := objc.Send[waitingForRfcommSpace](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ waitingForRfcommSpace) Init() waitingForRfcommSpace {
	rv := objc.Send[waitingForRfcommSpace](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ waitingForRfcommSpace) Autorelease() waitingForRfcommSpace {
	rv := objc.Send[waitingForRfcommSpace](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewwaitingForRfcommSpace creates a new waitingForRfcommSpace instance.
func NewwaitingForRfcommSpace() waitingForRfcommSpace {
	return getwaitingForRfcommSpaceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for waitingForRfcommSpace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/waitingForRfcommSpace
type waitingForRfcommSpace struct {
	objectivec.Object
}

// waitingForRfcommSpaceFrom constructs a [waitingForRfcommSpace] from an unsafe.Pointer.
func waitingForRfcommSpaceFrom(ptr unsafe.Pointer) waitingForRfcommSpace {
	return waitingForRfcommSpace{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for waitingForRfcommSpace *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for waitingForRfcommSpace */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for waitingForRfcommSpace */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for waitingForRfcommSpace */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for waitingForRfcommSpace */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class waitingForRfcommSpace */



