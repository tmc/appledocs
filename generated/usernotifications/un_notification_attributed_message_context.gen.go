// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [UNNotificationAttributedMessageContext] class.
type IUNNotificationAttributedMessageContext interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttributedMessageContext
type UNNotificationAttributedMessageContext struct {
	objectivec.Object
}

// UNNotificationAttributedMessageContextFrom constructs a [UNNotificationAttributedMessageContext] from an unsafe.Pointer.
func UNNotificationAttributedMessageContextFrom(ptr unsafe.Pointer) UNNotificationAttributedMessageContext {
	return UNNotificationAttributedMessageContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UNNotificationAttributedMessageContextClass) Alloc() UNNotificationAttributedMessageContext {
	rv := objc.Send[UNNotificationAttributedMessageContext](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttributedMessageContext/init(sendMessageIntent:attributedContent:)
func NewUNNotificationAttributedMessageContextWithSendMessageIntentAttributedContent(sendMessageIntent unsafe.Pointer, attributedContent unsafe.Pointer) UNNotificationAttributedMessageContext {
	rv := objc.Send[UNNotificationAttributedMessageContext](objc.ID(getUNNotificationAttributedMessageContextClass().class), objc.Sel("contextWithSendMessageIntent:attributedContent:"), sendMessageIntent, attributedContent)
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttributedMessageContext/init(sendMessageIntent:attributedContent:)
func (uc _UNNotificationAttributedMessageContextClass) ContextWithSendMessageIntentAttributedContent(sendMessageIntent unsafe.Pointer, attributedContent unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("contextWithSendMessageIntent:attributedContent:"), sendMessageIntent, attributedContent)
	return rv
}


