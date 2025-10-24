// Code generated from Apple documentation for IOBluetoothUI. DO NOT EDIT.

package iobluetoothui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class isIncomingRequest */


/* debug [class_header]: Header for isIncomingRequest */
// The class instance for the [isIncomingRequest] class.
var (
	IsIncomingRequestClass     _isIncomingRequestClass
	IsIncomingRequestClassOnce sync.Once
)

func getisIncomingRequestClass() _isIncomingRequestClass {
	IsIncomingRequestClassOnce.Do(func() {
		IsIncomingRequestClass = _isIncomingRequestClass{objc.GetClass("isIncomingRequest")}
	})
	return IsIncomingRequestClass
}

type _isIncomingRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for isIncomingRequest */
// An interface definition for the [isIncomingRequest] class.
type IisIncomingRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for isIncomingRequest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for isIncomingRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for isIncomingRequest */
// Alloc allocates a new instance without initialization.
func (ic _isIncomingRequestClass) Alloc() isIncomingRequest {
	rv := objc.Send[isIncomingRequest](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _isIncomingRequestClass) New() isIncomingRequest {
	rv := objc.Send[isIncomingRequest](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ isIncomingRequest) Init() isIncomingRequest {
	rv := objc.Send[isIncomingRequest](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ isIncomingRequest) Autorelease() isIncomingRequest {
	rv := objc.Send[isIncomingRequest](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewisIncomingRequest creates a new isIncomingRequest instance.
func NewisIncomingRequest() isIncomingRequest {
	return getisIncomingRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for isIncomingRequest */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetoothUI/IOBluetoothPasskeyDisplay/isIncomingRequest-c.ivar
type isIncomingRequest struct {
	objectivec.Object
}

// isIncomingRequestFrom constructs a [isIncomingRequest] from an unsafe.Pointer.
func isIncomingRequestFrom(ptr unsafe.Pointer) isIncomingRequest {
	return isIncomingRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for isIncomingRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for isIncomingRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for isIncomingRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for isIncomingRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for isIncomingRequest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class isIncomingRequest */



