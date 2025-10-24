// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class UNMutableNotificationContent */

/* debug [class_header]: Header for UNMutableNotificationContent */
// The class instance for the [UNMutableNotificationContent] class.
var (
	UNMutableNotificationContentClass     _UNMutableNotificationContentClass
	UNMutableNotificationContentClassOnce sync.Once
)

func getUNMutableNotificationContentClass() _UNMutableNotificationContentClass {
	UNMutableNotificationContentClassOnce.Do(func() {
		UNMutableNotificationContentClass = _UNMutableNotificationContentClass{objc.GetClass("UNMutableNotificationContent")}
	})
	return UNMutableNotificationContentClass
}

type _UNMutableNotificationContentClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for UNMutableNotificationContent */
// An interface definition for the [UNMutableNotificationContent] class.
type IUNMutableNotificationContent interface {
	IUNNotificationContent

	/* debug [class_interface_properties]: Properties for UNMutableNotificationContent */
	// properties:
	Attachments() []objc.IObject /* cross-framework: UNNotificationAttachment */
	SetAttachments(value []objc.IObject /* cross-framework: UNNotificationAttachment */)
	Badge() objc.IObject /* cross-framework: NSNumber */
	SetBadge(value objc.IObject /* cross-framework: NSNumber */)
	Body() objc.IObject /* cross-framework: NSString */
	SetBody(value objc.IObject /* cross-framework: NSString */)
	CategoryIdentifier() objc.IObject /* cross-framework: NSString */
	SetCategoryIdentifier(value objc.IObject /* cross-framework: NSString */)
	FilterCriteria() objc.IObject /* cross-framework: NSString */
	SetFilterCriteria(value objc.IObject /* cross-framework: NSString */)
	InterruptionLevel() UNNotificationInterruptionLevel
	SetInterruptionLevel(value UNNotificationInterruptionLevel)
	RelevanceScore() float64
	SetRelevanceScore(value float64)
	Sound() IUNNotificationSound
	SetSound(value IUNNotificationSound)
	Subtitle() objc.IObject /* cross-framework: NSString */
	SetSubtitle(value objc.IObject /* cross-framework: NSString */)
	SummaryArgument() objc.IObject /* cross-framework: NSString */
	SetSummaryArgument(value objc.IObject /* cross-framework: NSString */)
	SummaryArgumentCount() uint
	SetSummaryArgumentCount(value uint)
	TargetContentIdentifier() objc.IObject /* cross-framework: NSString */
	SetTargetContentIdentifier(value objc.IObject /* cross-framework: NSString */)
	ThreadIdentifier() objc.IObject /* cross-framework: NSString */
	SetThreadIdentifier(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	UserInfo() objc.IObject /* cross-framework: NSDictionary */
	SetUserInfo(value objc.IObject /* cross-framework: NSDictionary */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for UNMutableNotificationContent */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for UNMutableNotificationContent */
// Alloc allocates a new instance without initialization.
func (uc _UNMutableNotificationContentClass) Alloc() UNMutableNotificationContent {
	rv := objc.Send[UNMutableNotificationContent](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UNMutableNotificationContentClass) New() UNMutableNotificationContent {
	rv := objc.Send[UNMutableNotificationContent](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNMutableNotificationContent) Init() UNMutableNotificationContent {
	rv := objc.Send[UNMutableNotificationContent](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNMutableNotificationContent) Autorelease() UNMutableNotificationContent {
	rv := objc.Send[UNMutableNotificationContent](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNMutableNotificationContent creates a new UNMutableNotificationContent instance.
func NewUNMutableNotificationContent() UNMutableNotificationContent {
	return getUNMutableNotificationContentClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for UNMutableNotificationContent */
// The editable content for a notification.
//
// Create a object when you want to specify the payload for a local notification. Specifically, use this object to specify the title and message for an alert, the sound to play, or the value to assign to your app’s badge. You might also provide details about how the system handles the notification. For example, you can specify a custom launch image and a thread identifier for visually grouping related notifications. After creating your content object, assign it to a object, add a trigger condition, and schedule your notification. The trigger condition defines when the system delivers the notification to the user. Listing 1 shows the scheduling of a local notification that displays an alert and plays a sound after a delay of five seconds. Store the strings for the alert’s title and body in the app’s file. Listing 1. Creating the content for a local notification

// The editable content for a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent
type UNMutableNotificationContent struct {
	UNNotificationContent
}

// UNMutableNotificationContentFrom constructs a [UNMutableNotificationContent] from an unsafe.Pointer.
//
// The editable content for a notification.
func UNMutableNotificationContentFrom(ptr unsafe.Pointer) UNMutableNotificationContent {
	return UNMutableNotificationContent{
		UNNotificationContent: UNNotificationContentFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for UNMutableNotificationContent */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for UNMutableNotificationContent */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for UNMutableNotificationContent */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for UNMutableNotificationContent */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for UNMutableNotificationContent */

// The visual and audio attachments to display alongside the notification’s main content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/attachments
func (u_ UNMutableNotificationContent) Attachments() []objc.IObject /* cross-framework: UNNotificationAttachment */ {
	rv := objc.Send[[]UNNotificationAttachment](u_.ID, objc.Sel("attachments"))
	return rv
} /* debug [instance_properties/getter]: attachments */

// The visual and audio attachments to display alongside the notification’s main content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/attachments
func (u_ UNMutableNotificationContent) SetAttachments(value []objc.IObject /* cross-framework: UNNotificationAttachment */) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](u_.ID, objc.Sel("setAttachments:"), nsArray)
} /* debug [instance_properties/setter]: attachments */

// The number that your app’s icon displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/badge
func (u_ UNMutableNotificationContent) Badge() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](u_.ID, objc.Sel("badge"))
	return rv
} /* debug [instance_properties/getter]: badge */

// The number that your app’s icon displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/badge
func (u_ UNMutableNotificationContent) SetBadge(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setBadge:"), value)
} /* debug [instance_properties/setter]: badge */

// The localized text that provides the notification’s main content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/body
func (u_ UNMutableNotificationContent) Body() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("body"))
	return rv
} /* debug [instance_properties/getter]: body */

// The localized text that provides the notification’s main content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/body
func (u_ UNMutableNotificationContent) SetBody(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setBody:"), value)
} /* debug [instance_properties/setter]: body */

// The identifier of the notification’s category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/categoryIdentifier
func (u_ UNMutableNotificationContent) CategoryIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("categoryIdentifier"))
	return rv
} /* debug [instance_properties/getter]: categoryIdentifier */

// The identifier of the notification’s category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/categoryIdentifier
func (u_ UNMutableNotificationContent) SetCategoryIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCategoryIdentifier:"), value)
} /* debug [instance_properties/setter]: categoryIdentifier */

// The criteria the system evaluates to determine if it displays the notification in the current Focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/filterCriteria
func (u_ UNMutableNotificationContent) FilterCriteria() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("filterCriteria"))
	return rv
} /* debug [instance_properties/getter]: filterCriteria */

// The criteria the system evaluates to determine if it displays the notification in the current Focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/filterCriteria
func (u_ UNMutableNotificationContent) SetFilterCriteria(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setFilterCriteria:"), value)
} /* debug [instance_properties/setter]: filterCriteria */

// The notification’s importance and required delivery timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/interruptionLevel
func (u_ UNMutableNotificationContent) InterruptionLevel() UNNotificationInterruptionLevel {
	rv := objc.Send[UNNotificationInterruptionLevel](u_.ID, objc.Sel("interruptionLevel"))
	return rv
} /* debug [instance_properties/getter]: interruptionLevel */

// The notification’s importance and required delivery timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/interruptionLevel
func (u_ UNMutableNotificationContent) SetInterruptionLevel(value UNNotificationInterruptionLevel) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInterruptionLevel:"), value)
} /* debug [instance_properties/setter]: interruptionLevel */

// The score the system uses to determine if the notification is the summary’s featured notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/relevanceScore
func (u_ UNMutableNotificationContent) RelevanceScore() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("relevanceScore"))
	return rv
} /* debug [instance_properties/getter]: relevanceScore */

// The score the system uses to determine if the notification is the summary’s featured notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/relevanceScore
func (u_ UNMutableNotificationContent) SetRelevanceScore(value float64) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRelevanceScore:"), value)
} /* debug [instance_properties/setter]: relevanceScore */

// The sound that plays when the system delivers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/sound
func (u_ UNMutableNotificationContent) Sound() IUNNotificationSound {
	rv := objc.Send[UNNotificationSound](u_.ID, objc.Sel("sound"))
	return rv
} /* debug [instance_properties/getter]: sound */

// The sound that plays when the system delivers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/sound
func (u_ UNMutableNotificationContent) SetSound(value IUNNotificationSound) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSound:"), value)
} /* debug [instance_properties/setter]: sound */

// The localized text that provides the notification’s secondary description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/subtitle
func (u_ UNMutableNotificationContent) Subtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("subtitle"))
	return rv
} /* debug [instance_properties/getter]: subtitle */

// The localized text that provides the notification’s secondary description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/subtitle
func (u_ UNMutableNotificationContent) SetSubtitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSubtitle:"), value)
} /* debug [instance_properties/setter]: subtitle */

// The text the system adds to the notification summary to provide additional context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/summaryArgument
func (u_ UNMutableNotificationContent) SummaryArgument() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("summaryArgument"))
	return rv
} /* debug [instance_properties/getter]: summaryArgument */

// The text the system adds to the notification summary to provide additional context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/summaryArgument
func (u_ UNMutableNotificationContent) SetSummaryArgument(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSummaryArgument:"), value)
} /* debug [instance_properties/setter]: summaryArgument */

// The number the system adds to the notification summary when the notification represents multiple items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/summaryArgumentCount
func (u_ UNMutableNotificationContent) SummaryArgumentCount() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("summaryArgumentCount"))
	return rv
} /* debug [instance_properties/getter]: summaryArgumentCount */

// The number the system adds to the notification summary when the notification represents multiple items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/summaryArgumentCount
func (u_ UNMutableNotificationContent) SetSummaryArgumentCount(value uint) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSummaryArgumentCount:"), value)
} /* debug [instance_properties/setter]: summaryArgumentCount */

// The value your app uses to determine which scene to display to handle the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/targetContentIdentifier
func (u_ UNMutableNotificationContent) TargetContentIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("targetContentIdentifier"))
	return rv
} /* debug [instance_properties/getter]: targetContentIdentifier */

// The value your app uses to determine which scene to display to handle the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/targetContentIdentifier
func (u_ UNMutableNotificationContent) SetTargetContentIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTargetContentIdentifier:"), value)
} /* debug [instance_properties/setter]: targetContentIdentifier */

// The identifier that groups related notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/threadIdentifier
func (u_ UNMutableNotificationContent) ThreadIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("threadIdentifier"))
	return rv
} /* debug [instance_properties/getter]: threadIdentifier */

// The identifier that groups related notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/threadIdentifier
func (u_ UNMutableNotificationContent) SetThreadIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setThreadIdentifier:"), value)
} /* debug [instance_properties/setter]: threadIdentifier */

// The localized text that provides the notification’s primary description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/title
func (u_ UNMutableNotificationContent) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("title"))
	return rv
} /* debug [instance_properties/getter]: title */

// The localized text that provides the notification’s primary description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/title
func (u_ UNMutableNotificationContent) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTitle:"), value)
} /* debug [instance_properties/setter]: title */

// The custom data to associate with the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/userInfo
func (u_ UNMutableNotificationContent) UserInfo() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](u_.ID, objc.Sel("userInfo"))
	return rv
} /* debug [instance_properties/getter]: userInfo */

// The custom data to associate with the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/userInfo
func (u_ UNMutableNotificationContent) SetUserInfo(value objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUserInfo:"), value)
} /* debug [instance_properties/setter]: userInfo */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class UNMutableNotificationContent */
