// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CXCallController */


/* debug [class_header]: Header for CXCallController */
// The class instance for the [CXCallController] class.
var (
	CXCallControllerClass     _CXCallControllerClass
	CXCallControllerClassOnce sync.Once
)

func getCXCallControllerClass() _CXCallControllerClass {
	CXCallControllerClassOnce.Do(func() {
		CXCallControllerClass = _CXCallControllerClass{objc.GetClass("CXCallController")}
	})
	return CXCallControllerClass
}

type _CXCallControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXCallController */
// An interface definition for the [CXCallController] class.
type ICXCallController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CXCallController */
	// properties:
	CXErrorDomainRequestTransaction() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXCallController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXCallController */
// Alloc allocates a new instance without initialization.
func (cc _CXCallControllerClass) Alloc() CXCallController {
	rv := objc.Send[CXCallController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXCallControllerClass) New() CXCallController {
	rv := objc.Send[CXCallController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallController) Init() CXCallController {
	rv := objc.Send[CXCallController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallController) Autorelease() CXCallController {
	rv := objc.Send[CXCallController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallController creates a new CXCallController instance.
func NewCXCallController() CXCallController {
	return getCXCallControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXCallController */
// A programmatic interface for interacting with and observing calls.
//
// A object interacts with calls by performing actions, which are represented by instances of subclasses. You can request that one or more actions be performed in a single object using the method. A transaction may be rejected by the system for one of the reasons listed in the enumeration. Each object manages a object, which can be accessed using the property. You can provide an object conforming to the protocol to the call observer in order to be notified of any changes to active calls.


// A programmatic interface for interacting with and observing calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallController
type CXCallController struct {
	objectivec.Object
}

// CXCallControllerFrom constructs a [CXCallController] from an unsafe.Pointer.
//
// A programmatic interface for interacting with and observing calls.
func CXCallControllerFrom(ptr unsafe.Pointer) CXCallController {
	return CXCallController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXCallController */

// Initializes a new call controller with a specified queue, which is used for calling completion blocks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallController/init(queue:)
func NewCXCallControllerWithQueue(queue unsafe.Pointer) CXCallController {
	instance := getCXCallControllerClass().Alloc()
	rv := objc.Send[CXCallController](instance.ID, objc.Sel("initWithQueue:"), queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXCallControllerWithQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXCallController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXCallController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXCallController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXCallController */

// Domain for errors when requesting a transaction from a call controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxerrordomainrequesttransaction
func (c_ CXCallController) CXErrorDomainRequestTransaction() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CXErrorDomainRequestTransaction"))
	return rv
}/* debug [instance_properties/getter]: CXErrorDomainRequestTransaction */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXCallController */


