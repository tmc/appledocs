// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class RPBroadcastActivityController */


/* debug [class_header]: Header for RPBroadcastActivityController */
// The class instance for the [RPBroadcastActivityController] class.
var (
	RPBroadcastActivityControllerClass     _RPBroadcastActivityControllerClass
	RPBroadcastActivityControllerClassOnce sync.Once
)

func getRPBroadcastActivityControllerClass() _RPBroadcastActivityControllerClass {
	RPBroadcastActivityControllerClassOnce.Do(func() {
		RPBroadcastActivityControllerClass = _RPBroadcastActivityControllerClass{objc.GetClass("RPBroadcastActivityController")}
	})
	return RPBroadcastActivityControllerClass
}

type _RPBroadcastActivityControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RPBroadcastActivityController */
// An interface definition for the [RPBroadcastActivityController] class.
type IRPBroadcastActivityController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RPBroadcastActivityController */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RPBroadcastActivityController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RPBroadcastActivityController */
// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastActivityControllerClass) Alloc() RPBroadcastActivityController {
	rv := objc.Send[RPBroadcastActivityController](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RPBroadcastActivityControllerClass) New() RPBroadcastActivityController {
	rv := objc.Send[RPBroadcastActivityController](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPBroadcastActivityController) Init() RPBroadcastActivityController {
	rv := objc.Send[RPBroadcastActivityController](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPBroadcastActivityController) Autorelease() RPBroadcastActivityController {
	rv := objc.Send[RPBroadcastActivityController](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPBroadcastActivityController creates a new RPBroadcastActivityController instance.
func NewRPBroadcastActivityController() RPBroadcastActivityController {
	return getRPBroadcastActivityControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RPBroadcastActivityController */
// A controller object that presents the macOS broadcast picker.


// A controller object that presents the macOS broadcast picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityController
type RPBroadcastActivityController struct {
	objectivec.Object
}

// RPBroadcastActivityControllerFrom constructs a [RPBroadcastActivityController] from an unsafe.Pointer.
//
// A controller object that presents the macOS broadcast picker.
func RPBroadcastActivityControllerFrom(ptr unsafe.Pointer) RPBroadcastActivityController {
	return RPBroadcastActivityController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RPBroadcastActivityController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RPBroadcastActivityController */

// Presents a list of available broadcast services for the user to select.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityController/showBroadcastPicker(at:from:preferredExtensionIdentifier:completionHandler:)
func (rc _RPBroadcastActivityControllerClass) ShowBroadcastPickerAtPointFromWindowPreferredExtensionIdentifierCompletionHandler(point corefoundation.CGPoint, window appkit.Window, preferredExtension objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(rc.class), objc.Sel("showBroadcastPickerAtPoint:fromWindow:preferredExtensionIdentifier:completionHandler:"), point, window, preferredExtension, handler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ShowBroadcastPickerAtPointFromWindowPreferredExtensionIdentifierCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RPBroadcastActivityController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RPBroadcastActivityController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RPBroadcastActivityController */

// The broadcast activity controller’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityController/delegate
func (r_ RPBroadcastActivityController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The broadcast activity controller’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityController/delegate
func (r_ RPBroadcastActivityController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class RPBroadcastActivityController */



