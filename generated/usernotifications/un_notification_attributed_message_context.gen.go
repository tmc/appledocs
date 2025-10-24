// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/intents"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UNNotificationAttributedMessageContext */


/* debug [class_header]: Header for UNNotificationAttributedMessageContext */
// The class instance for the [UNNotificationAttributedMessageContext] class.
var (
	UNNotificationAttributedMessageContextClass     _UNNotificationAttributedMessageContextClass
	UNNotificationAttributedMessageContextClassOnce sync.Once
)

func getUNNotificationAttributedMessageContextClass() _UNNotificationAttributedMessageContextClass {
	UNNotificationAttributedMessageContextClassOnce.Do(func() {
		UNNotificationAttributedMessageContextClass = _UNNotificationAttributedMessageContextClass{objc.GetClass("UNNotificationAttributedMessageContext")}
	})
	return UNNotificationAttributedMessageContextClass
}

type _UNNotificationAttributedMessageContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UNNotificationAttributedMessageContext */
// An interface definition for the [UNNotificationAttributedMessageContext] class.
type IUNNotificationAttributedMessageContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UNNotificationAttributedMessageContext */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UNNotificationAttributedMessageContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UNNotificationAttributedMessageContext */
// Alloc allocates a new instance without initialization.
func (uc _UNNotificationAttributedMessageContextClass) Alloc() UNNotificationAttributedMessageContext {
	rv := objc.Send[UNNotificationAttributedMessageContext](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UNNotificationAttributedMessageContextClass) New() UNNotificationAttributedMessageContext {
	rv := objc.Send[UNNotificationAttributedMessageContext](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNNotificationAttributedMessageContext) Init() UNNotificationAttributedMessageContext {
	rv := objc.Send[UNNotificationAttributedMessageContext](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNNotificationAttributedMessageContext) Autorelease() UNNotificationAttributedMessageContext {
	rv := objc.Send[UNNotificationAttributedMessageContext](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationAttributedMessageContext creates a new UNNotificationAttributedMessageContext instance.
func NewUNNotificationAttributedMessageContext() UNNotificationAttributedMessageContext {
	return getUNNotificationAttributedMessageContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UNNotificationAttributedMessageContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttributedMessageContext
type UNNotificationAttributedMessageContext struct {
	objectivec.Object
}

// UNNotificationAttributedMessageContextFrom constructs a [UNNotificationAttributedMessageContext] from an unsafe.Pointer.
func UNNotificationAttributedMessageContextFrom(ptr unsafe.Pointer) UNNotificationAttributedMessageContext {
	return UNNotificationAttributedMessageContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UNNotificationAttributedMessageContext */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttributedMessageContext/init(sendMessageIntent:attributedContent:)
func NewUNNotificationAttributedMessageContextWithSendMessageIntentAttributedContent(sendMessageIntent intents.INSendMessageIntent, attributedContent foundation.AttributedString) UNNotificationAttributedMessageContext {
	rv := objc.Send[UNNotificationAttributedMessageContext](objc.ID(getUNNotificationAttributedMessageContextClass().class), objc.Sel("contextWithSendMessageIntent:attributedContent:"), sendMessageIntent, attributedContent)
	return rv
}/* debug [class_init_methods/constructor]: NewUNNotificationAttributedMessageContextWithSendMessageIntentAttributedContent */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UNNotificationAttributedMessageContext */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttributedMessageContext/init(sendMessageIntent:attributedContent:)
func (uc _UNNotificationAttributedMessageContextClass) ContextWithSendMessageIntentAttributedContent(sendMessageIntent intents.INSendMessageIntent, attributedContent foundation.AttributedString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("contextWithSendMessageIntent:attributedContent:"), sendMessageIntent, attributedContent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContextWithSendMessageIntentAttributedContent) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UNNotificationAttributedMessageContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UNNotificationAttributedMessageContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UNNotificationAttributedMessageContext */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class UNNotificationAttributedMessageContext */


