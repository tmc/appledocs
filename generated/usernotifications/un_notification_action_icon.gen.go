// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UNNotificationActionIcon */

/* debug [class_header]: Header for UNNotificationActionIcon */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for UNNotificationActionIcon */
// An interface definition for the [UNNotificationActionIcon] class.
type IUNNotificationActionIcon interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for UNNotificationActionIcon */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for UNNotificationActionIcon */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for UNNotificationActionIcon */
// Alloc allocates a new instance without initialization.
func (uc _UNNotificationActionIconClass) Alloc() UNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for UNNotificationActionIcon */
// An icon associated with an action.

// An icon associated with an action.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for UNNotificationActionIcon */

// Creates an action icon by using a system symbol image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionIcon/init(systemImageName:)
func NewUNNotificationActionIconWithSystemImageName(systemImageName objc.IObject /* cross-framework: NSString */) UNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](objc.ID(getUNNotificationActionIconClass().class), objc.Sel("iconWithSystemImageName:"), systemImageName)
	return rv
} /* debug [class_init_methods/constructor]: NewUNNotificationActionIconWithSystemImageName */

// Creates an action icon based on an image in your app’s bundle, preferably in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionIcon/init(templateImageName:)
func NewUNNotificationActionIconWithTemplateImageName(templateImageName objc.IObject /* cross-framework: NSString */) UNNotificationActionIcon {
	rv := objc.Send[UNNotificationActionIcon](objc.ID(getUNNotificationActionIconClass().class), objc.Sel("iconWithTemplateImageName:"), templateImageName)
	return rv
} /* debug [class_init_methods/constructor]: NewUNNotificationActionIconWithTemplateImageName */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for UNNotificationActionIcon */

// Creates an action icon by using a system symbol image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionIcon/init(systemImageName:)
func (uc _UNNotificationActionIconClass) IconWithSystemImageName(systemImageName objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("iconWithSystemImageName:"), systemImageName)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=IconWithSystemImageName) */

// Creates an action icon based on an image in your app’s bundle, preferably in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionIcon/init(templateImageName:)
func (uc _UNNotificationActionIconClass) IconWithTemplateImageName(templateImageName objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("iconWithTemplateImageName:"), templateImageName)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=IconWithTemplateImageName) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for UNNotificationActionIcon */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for UNNotificationActionIcon */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for UNNotificationActionIcon */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class UNNotificationActionIcon */
