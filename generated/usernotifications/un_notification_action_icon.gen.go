// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [UNNotificationActionIcon] class.
var (
	UNNotificationActionIconClass     _UNNotificationActionIconClass
	UNNotificationActionIconClassOnce sync.Once
)

func getUNNotificationActionIconClass() _UNNotificationActionIconClass {
	UNNotificationActionIconClassOnce.Do(func() {
		UNNotificationActionIconClass = _UNNotificationActionIconClass{objc.GetClass("UNNotificationActionIcon")}
	})
	return UNNotificationActionIconClass
}

type _UNNotificationActionIconClass struct {
	class objc.Class
}

// An interface definition for the [UNNotificationActionIcon] class.
type IUNNotificationActionIcon interface {
	objectivec.IObject
}

// An icon associated with an action.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionIcon
type UNNotificationActionIcon struct {
	objectivec.Object
}

// UNNotificationActionIconFrom constructs a [UNNotificationActionIcon] from an unsafe.Pointer.
//
// An icon associated with an action.
func UNNotificationActionIconFrom(ptr unsafe.Pointer) UNNotificationActionIcon {
	return UNNotificationActionIcon{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UNNotificationActionIconClass) Alloc() UNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UNNotificationActionIconClass) New() UNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNNotificationActionIcon) Init() UNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNNotificationActionIcon) Autorelease() UNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationActionIcon creates a new UNNotificationActionIcon instance.
func NewUNNotificationActionIcon() UNNotificationActionIcon {
	return getUNNotificationActionIconClass().New()
}


// Creates an action icon by using a system symbol image.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionIcon/init(systemImageName:)
func NewUNNotificationActionIconWithSystemImageName(systemImageName string) UNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](objc.ID(getUNNotificationActionIconClass().class), objc.Sel("iconWithSystemImageName:"), objc.String(systemImageName))
	return rv
}

// Creates an action icon based on an image in your app’s bundle, preferably in an asset catalog.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionIcon/init(templateImageName:)
func NewUNNotificationActionIconWithTemplateImageName(templateImageName string) UNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](objc.ID(getUNNotificationActionIconClass().class), objc.Sel("iconWithTemplateImageName:"), objc.String(templateImageName))
	return rv
}


// Creates an action icon by using a system symbol image.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionIcon/init(systemImageName:)
func (uc _UNNotificationActionIconClass) IconWithSystemImageName(systemImageName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("iconWithSystemImageName:"), objc.String(systemImageName))
	return rv
}

// Creates an action icon based on an image in your app’s bundle, preferably in an asset catalog.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionIcon/init(templateImageName:)
func (uc _UNNotificationActionIconClass) IconWithTemplateImageName(templateImageName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("iconWithTemplateImageName:"), objc.String(templateImageName))
	return rv
}


