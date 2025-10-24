// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mSavedCommand */


/* debug [class_header]: Header for mSavedCommand */
// The class instance for the [mSavedCommand] class.
var (
	MSavedCommandClass     _mSavedCommandClass
	MSavedCommandClassOnce sync.Once
)

func getmSavedCommandClass() _mSavedCommandClass {
	MSavedCommandClassOnce.Do(func() {
		MSavedCommandClass = _mSavedCommandClass{objc.GetClass("mSavedCommand")}
	})
	return MSavedCommandClass
}

type _mSavedCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mSavedCommand */
// An interface definition for the [mSavedCommand] class.
type ImSavedCommand interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mSavedCommand */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mSavedCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mSavedCommand */
// Alloc allocates a new instance without initialization.
func (mc _mSavedCommandClass) Alloc() mSavedCommand {
	rv := objc.Send[mSavedCommand](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mSavedCommandClass) New() mSavedCommand {
	rv := objc.Send[mSavedCommand](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mSavedCommand) Init() mSavedCommand {
	rv := objc.Send[mSavedCommand](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mSavedCommand) Autorelease() mSavedCommand {
	rv := objc.Send[mSavedCommand](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmSavedCommand creates a new mSavedCommand instance.
func NewmSavedCommand() mSavedCommand {
	return getmSavedCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mSavedCommand */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mSavedCommand
type mSavedCommand struct {
	objectivec.Object
}

// mSavedCommandFrom constructs a [mSavedCommand] from an unsafe.Pointer.
func mSavedCommandFrom(ptr unsafe.Pointer) mSavedCommand {
	return mSavedCommand{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mSavedCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mSavedCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mSavedCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mSavedCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mSavedCommand */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mSavedCommand */



