// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mL2CAPChannelConnectionHandler */


/* debug [class_header]: Header for mL2CAPChannelConnectionHandler */
// The class instance for the [mL2CAPChannelConnectionHandler] class.
var (
	ML2CAPChannelConnectionHandlerClass     _mL2CAPChannelConnectionHandlerClass
	ML2CAPChannelConnectionHandlerClassOnce sync.Once
)

func getmL2CAPChannelConnectionHandlerClass() _mL2CAPChannelConnectionHandlerClass {
	ML2CAPChannelConnectionHandlerClassOnce.Do(func() {
		ML2CAPChannelConnectionHandlerClass = _mL2CAPChannelConnectionHandlerClass{objc.GetClass("mL2CAPChannelConnectionHandler")}
	})
	return ML2CAPChannelConnectionHandlerClass
}

type _mL2CAPChannelConnectionHandlerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mL2CAPChannelConnectionHandler */
// An interface definition for the [mL2CAPChannelConnectionHandler] class.
type ImL2CAPChannelConnectionHandler interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mL2CAPChannelConnectionHandler */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mL2CAPChannelConnectionHandler */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mL2CAPChannelConnectionHandler */
// Alloc allocates a new instance without initialization.
func (mc _mL2CAPChannelConnectionHandlerClass) Alloc() mL2CAPChannelConnectionHandler {
	rv := objc.Send[mL2CAPChannelConnectionHandler](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mL2CAPChannelConnectionHandlerClass) New() mL2CAPChannelConnectionHandler {
	rv := objc.Send[mL2CAPChannelConnectionHandler](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mL2CAPChannelConnectionHandler) Init() mL2CAPChannelConnectionHandler {
	rv := objc.Send[mL2CAPChannelConnectionHandler](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mL2CAPChannelConnectionHandler) Autorelease() mL2CAPChannelConnectionHandler {
	rv := objc.Send[mL2CAPChannelConnectionHandler](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmL2CAPChannelConnectionHandler creates a new mL2CAPChannelConnectionHandler instance.
func NewmL2CAPChannelConnectionHandler() mL2CAPChannelConnectionHandler {
	return getmL2CAPChannelConnectionHandlerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mL2CAPChannelConnectionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/mL2CAPChannelConnectionHandler
type mL2CAPChannelConnectionHandler struct {
	objectivec.Object
}

// mL2CAPChannelConnectionHandlerFrom constructs a [mL2CAPChannelConnectionHandler] from an unsafe.Pointer.
func mL2CAPChannelConnectionHandlerFrom(ptr unsafe.Pointer) mL2CAPChannelConnectionHandler {
	return mL2CAPChannelConnectionHandler{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mL2CAPChannelConnectionHandler *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mL2CAPChannelConnectionHandler */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mL2CAPChannelConnectionHandler */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mL2CAPChannelConnectionHandler */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mL2CAPChannelConnectionHandler */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mL2CAPChannelConnectionHandler */



