// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mActionCommand */


/* debug [class_header]: Header for mActionCommand */
// The class instance for the [mActionCommand] class.
var (
	MActionCommandClass     _mActionCommandClass
	MActionCommandClassOnce sync.Once
)

func getmActionCommandClass() _mActionCommandClass {
	MActionCommandClassOnce.Do(func() {
		MActionCommandClass = _mActionCommandClass{objc.GetClass("mActionCommand")}
	})
	return MActionCommandClass
}

type _mActionCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mActionCommand */
// An interface definition for the [mActionCommand] class.
type ImActionCommand interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mActionCommand */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mActionCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mActionCommand */
// Alloc allocates a new instance without initialization.
func (mc _mActionCommandClass) Alloc() mActionCommand {
	rv := objc.Send[mActionCommand](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mActionCommandClass) New() mActionCommand {
	rv := objc.Send[mActionCommand](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mActionCommand) Init() mActionCommand {
	rv := objc.Send[mActionCommand](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mActionCommand) Autorelease() mActionCommand {
	rv := objc.Send[mActionCommand](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmActionCommand creates a new mActionCommand instance.
func NewmActionCommand() mActionCommand {
	return getmActionCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mActionCommand */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/OBEXFileTransferServices/mActionCommand
type mActionCommand struct {
	objectivec.Object
}

// mActionCommandFrom constructs a [mActionCommand] from an unsafe.Pointer.
func mActionCommandFrom(ptr unsafe.Pointer) mActionCommand {
	return mActionCommand{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mActionCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mActionCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mActionCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mActionCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mActionCommand */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mActionCommand */



