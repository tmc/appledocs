// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class UNTextInputNotificationResponse */


/* debug [class_header]: Header for UNTextInputNotificationResponse */
// The class instance for the [UNTextInputNotificationResponse] class.
var (
	UNTextInputNotificationResponseClass     _UNTextInputNotificationResponseClass
	UNTextInputNotificationResponseClassOnce sync.Once
)

func getUNTextInputNotificationResponseClass() _UNTextInputNotificationResponseClass {
	UNTextInputNotificationResponseClassOnce.Do(func() {
		UNTextInputNotificationResponseClass = _UNTextInputNotificationResponseClass{objc.GetClass("UNTextInputNotificationResponse")}
	})
	return UNTextInputNotificationResponseClass
}

type _UNTextInputNotificationResponseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UNTextInputNotificationResponse */
// An interface definition for the [UNTextInputNotificationResponse] class.
type IUNTextInputNotificationResponse interface {
	IUNNotificationResponse
	
/* debug [class_interface_properties]: Properties for UNTextInputNotificationResponse */
	// properties:
	UserText() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UNTextInputNotificationResponse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UNTextInputNotificationResponse */
// Alloc allocates a new instance without initialization.
func (uc _UNTextInputNotificationResponseClass) Alloc() UNTextInputNotificationResponse {
	rv := objc.Send[UNTextInputNotificationResponse](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UNTextInputNotificationResponseClass) New() UNTextInputNotificationResponse {
	rv := objc.Send[UNTextInputNotificationResponse](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNTextInputNotificationResponse) Init() UNTextInputNotificationResponse {
	rv := objc.Send[UNTextInputNotificationResponse](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNTextInputNotificationResponse) Autorelease() UNTextInputNotificationResponse {
	rv := objc.Send[UNTextInputNotificationResponse](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNTextInputNotificationResponse creates a new UNTextInputNotificationResponse instance.
func NewUNTextInputNotificationResponse() UNTextInputNotificationResponse {
	return getUNTextInputNotificationResponseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UNTextInputNotificationResponse */
// The user’s response to an actionable notification, including any custom text that the user typed or dictated.
//
// The system delivers a object to your app so that you can process user-provided text content. When defining your categories, you can specify an object instead of an object for your action. If you do, the system creates an object when the user selects the accompanying action, and it fills the property with any user-entered text. You don’t create objects yourself. Instead, the shared user notification center object creates them and delivers them to the method of its delegate object. Use that method to extract any needed information from the response object and take appropriate action. For more information about responding to actions, see .


// The user’s response to an actionable notification, including any custom text that the user typed or dictated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationResponse
type UNTextInputNotificationResponse struct {
	UNNotificationResponse
}

// UNTextInputNotificationResponseFrom constructs a [UNTextInputNotificationResponse] from an unsafe.Pointer.
//
// The user’s response to an actionable notification, including any custom text that the user typed or dictated.
func UNTextInputNotificationResponseFrom(ptr unsafe.Pointer) UNTextInputNotificationResponse {
	return UNTextInputNotificationResponse{
		UNNotificationResponse: UNNotificationResponseFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UNTextInputNotificationResponse *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UNTextInputNotificationResponse */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UNTextInputNotificationResponse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UNTextInputNotificationResponse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UNTextInputNotificationResponse */

// The text response provided by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationResponse/userText
func (u_ UNTextInputNotificationResponse) UserText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("userText"))
	return rv
}/* debug [instance_properties/getter]: userText */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class UNTextInputNotificationResponse */



