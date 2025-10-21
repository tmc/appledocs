// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UNNotificationServiceExtension] class.
var (
	UNNotificationServiceExtensionClass     _UNNotificationServiceExtensionClass
	UNNotificationServiceExtensionClassOnce sync.Once
)

func getUNNotificationServiceExtensionClass() _UNNotificationServiceExtensionClass {
	UNNotificationServiceExtensionClassOnce.Do(func() {
		UNNotificationServiceExtensionClass = _UNNotificationServiceExtensionClass{objc.GetClass("UNNotificationServiceExtension")}
	})
	return UNNotificationServiceExtensionClass
}

type _UNNotificationServiceExtensionClass struct {
	class objc.Class
}

// An interface definition for the [UNNotificationServiceExtension] class.
type IUNNotificationServiceExtension interface {
	objectivec.IObject
	DidReceiveNotificationRequestWithContentHandler(request unsafe.Pointer, contentHandler unsafe.Pointer)
	ServiceExtensionTimeWillExpire()
}

// An object that modifies the content of a remote notification before it’s delivered to the user.
//
// A object provides the entry point for a notification service app extension. This object lets you customize the content of a remote notification before the system delivers it to the user. A notification service app extension doesn’t present any UI of its own. Instead, it’s launched on demand when the system delivers a notification of the appropriate type to the user’s device. You use this extension to modify the notification’s content or download content related to the extension. For example, you could use the extension to decrypt an encrypted data block or to download images associated with the notification. You don’t create objects yourself. Instead, the Xcode template for a notification service extension target contains a subclass for you to modify. Use the methods of that subclass to implement your app extension’s behavior. When your app receives a remote notification for your app, the system loads your extension and calls its method given the following conditions: Your app has configured the remote notification to display an alert. The remote notification’s dictionary includes the key with the value set to . The method performs the main work of your extension. You use that method to make any changes to the notification’s content. That method has a limited amount of time to perform its task and execute the provided completion block. If your method doesn’t finish in time, the system calls the method to give you one last chance to submit your changes. If you don’t update the notification content before time expires, the system displays the original content. As for any app extension, you deliver a notification service app extension class as a bundle inside your app. The template that Xcode provides configures the file automatically for this app extension type. Specifically, it sets the value of the key to and sets the value of the key to the name of your subclass. For information about how to set up and send remote notifications, see .
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationServiceExtension
type UNNotificationServiceExtension struct {
	objectivec.Object
}

// UNNotificationServiceExtensionFrom constructs a [UNNotificationServiceExtension] from an unsafe.Pointer.
//
// An object that modifies the content of a remote notification before it’s delivered to the user.
func UNNotificationServiceExtensionFrom(ptr unsafe.Pointer) UNNotificationServiceExtension {
	return UNNotificationServiceExtension{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UNNotificationServiceExtensionClass) Alloc() UNNotificationServiceExtension {
	rv := objc.Send[UNNotificationServiceExtension](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UNNotificationServiceExtensionClass) New() UNNotificationServiceExtension {
	rv := objc.Send[UNNotificationServiceExtension](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNNotificationServiceExtension) Init() UNNotificationServiceExtension {
	rv := objc.Send[UNNotificationServiceExtension](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNNotificationServiceExtension) Autorelease() UNNotificationServiceExtension {
	rv := objc.Send[UNNotificationServiceExtension](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationServiceExtension creates a new UNNotificationServiceExtension instance.
func NewUNNotificationServiceExtension() UNNotificationServiceExtension {
	return getUNNotificationServiceExtensionClass().New()
}


// Asks you to make any needed changes to the notification and notify the system when you’re done.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationServiceExtension/didReceive(_:withContentHandler:)
func (u_ UNNotificationServiceExtension) DidReceiveNotificationRequestWithContentHandler(request unsafe.Pointer, contentHandler unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("didReceiveNotificationRequest:withContentHandler:"), request, contentHandler)
}

// Tells you that the system is terminating your extension.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationServiceExtension/serviceExtensionTimeWillExpire()
func (u_ UNNotificationServiceExtension) ServiceExtensionTimeWillExpire() {
	objc.Send[objc.ID](u_.ID, objc.Sel("serviceExtensionTimeWillExpire"))
}



