// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [UNTextInputNotificationResponse] class.
type IUNTextInputNotificationResponse interface {
	IUNNotificationResponse
}

// The user’s response to an actionable notification, including any custom text that the user typed or dictated.
//
// The system delivers a object to your app so that you can process user-provided text content. When defining your categories, you can specify an object instead of an object for your action. If you do, the system creates an object when the user selects the accompanying action, and it fills the property with any user-entered text. You don’t create objects yourself. Instead, the shared user notification center object creates them and delivers them to the method of its delegate object. Use that method to extract any needed information from the response object and take appropriate action. For more information about responding to actions, see .
//
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

// Alloc allocates a new instance without initialization.
func (uc _UNTextInputNotificationResponseClass) Alloc() UNTextInputNotificationResponse {
	rv := objc.Send[UNTextInputNotificationResponse](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The text response provided by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationResponse/userText
func (u_ UNTextInputNotificationResponse) UserText() string {
	rv := objc.Send[string](u_.ID, objc.Sel("userText"))
	return rv
}



