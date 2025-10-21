// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [UNMutableNotificationContent] class.
type IUNMutableNotificationContent interface {
	IUNNotificationContent
}

// The editable content for a notification.
//
// Create a object when you want to specify the payload for a local notification. Specifically, use this object to specify the title and message for an alert, the sound to play, or the value to assign to your app’s badge. You might also provide details about how the system handles the notification. For example, you can specify a custom launch image and a thread identifier for visually grouping related notifications. After creating your content object, assign it to a object, add a trigger condition, and schedule your notification. The trigger condition defines when the system delivers the notification to the user. Listing 1 shows the scheduling of a local notification that displays an alert and plays a sound after a delay of five seconds. Store the strings for the alert’s title and body in the app’s file. Listing 1. Creating the content for a local notification
//
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

// Alloc allocates a new instance without initialization.
func (uc _UNMutableNotificationContentClass) Alloc() UNMutableNotificationContent {
	rv := objc.Send[UNMutableNotificationContent](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The visual and audio attachments to display alongside the notification’s main content.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/attachments
func (u_ UNMutableNotificationContent) Attachments() []UNNotificationAttachment {
	rv := objc.Send[[]UNNotificationAttachment](u_.ID, objc.Sel("attachments"))
	return rv
}


// SetAttachments sets the value of the attachments property.
// The visual and audio attachments to display alongside the notification’s main content.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/attachments
func (u_ UNMutableNotificationContent) SetAttachments(value []UNNotificationAttachment) {
	// Convert Go slice to NSArray
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
}

// The number that your app’s icon displays.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/badge
func (u_ UNMutableNotificationContent) Badge() foundation.Number {
	rv := objc.Send[foundation.Number](u_.ID, objc.Sel("badge"))
	return rv
}


// SetBadge sets the value of the badge property.
// The number that your app’s icon displays.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/badge
func (u_ UNMutableNotificationContent) SetBadge(value foundation.Number) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setBadge:"), value)
}

// The localized text that provides the notification’s main content.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/body
func (u_ UNMutableNotificationContent) Body() string {
	rv := objc.Send[string](u_.ID, objc.Sel("body"))
	return rv
}


// SetBody sets the value of the body property.
// The localized text that provides the notification’s main content.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/body
func (u_ UNMutableNotificationContent) SetBody(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setBody:"), objc.String(value))
}

// The identifier of the notification’s category.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/categoryIdentifier
func (u_ UNMutableNotificationContent) CategoryIdentifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("categoryIdentifier"))
	return rv
}


// SetCategoryIdentifier sets the value of the categoryIdentifier property.
// The identifier of the notification’s category.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/categoryIdentifier
func (u_ UNMutableNotificationContent) SetCategoryIdentifier(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCategoryIdentifier:"), objc.String(value))
}

// The criteria the system evaluates to determine if it displays the notification in the current Focus.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/filterCriteria
func (u_ UNMutableNotificationContent) FilterCriteria() string {
	rv := objc.Send[string](u_.ID, objc.Sel("filterCriteria"))
	return rv
}


// SetFilterCriteria sets the value of the filterCriteria property.
// The criteria the system evaluates to determine if it displays the notification in the current Focus.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/filterCriteria
func (u_ UNMutableNotificationContent) SetFilterCriteria(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setFilterCriteria:"), objc.String(value))
}

// The notification’s importance and required delivery timing.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/interruptionLevel
func (u_ UNMutableNotificationContent) InterruptionLevel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("interruptionLevel"))
	return rv
}


// SetInterruptionLevel sets the value of the interruptionLevel property.
// The notification’s importance and required delivery timing.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/interruptionLevel
func (u_ UNMutableNotificationContent) SetInterruptionLevel(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInterruptionLevel:"), value)
}

// The name of the image or storyboard to use when your app launches because of the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/launchImageName
func (u_ UNMutableNotificationContent) LaunchImageName() string {
	rv := objc.Send[string](u_.ID, objc.Sel("launchImageName"))
	return rv
}


// SetLaunchImageName sets the value of the launchImageName property.
// The name of the image or storyboard to use when your app launches because of the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/launchImageName
func (u_ UNMutableNotificationContent) SetLaunchImageName(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setLaunchImageName:"), objc.String(value))
}

// The score the system uses to determine if the notification is the summary’s featured notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/relevanceScore
func (u_ UNMutableNotificationContent) RelevanceScore() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("relevanceScore"))
	return rv
}


// SetRelevanceScore sets the value of the relevanceScore property.
// The score the system uses to determine if the notification is the summary’s featured notification.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/relevanceScore
func (u_ UNMutableNotificationContent) SetRelevanceScore(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRelevanceScore:"), value)
}

// The sound that plays when the system delivers the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/sound
func (u_ UNMutableNotificationContent) Sound() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("sound"))
	return rv
}


// SetSound sets the value of the sound property.
// The sound that plays when the system delivers the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/sound
func (u_ UNMutableNotificationContent) SetSound(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSound:"), value)
}

// The localized text that provides the notification’s secondary description.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/subtitle
func (u_ UNMutableNotificationContent) Subtitle() string {
	rv := objc.Send[string](u_.ID, objc.Sel("subtitle"))
	return rv
}


// SetSubtitle sets the value of the subtitle property.
// The localized text that provides the notification’s secondary description.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/subtitle
func (u_ UNMutableNotificationContent) SetSubtitle(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSubtitle:"), objc.String(value))
}

// The text the system adds to the notification summary to provide additional context.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/summaryArgument
func (u_ UNMutableNotificationContent) SummaryArgument() string {
	rv := objc.Send[string](u_.ID, objc.Sel("summaryArgument"))
	return rv
}


// SetSummaryArgument sets the value of the summaryArgument property.
// The text the system adds to the notification summary to provide additional context.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/summaryArgument
func (u_ UNMutableNotificationContent) SetSummaryArgument(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSummaryArgument:"), objc.String(value))
}

// The number the system adds to the notification summary when the notification represents multiple items.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/summaryArgumentCount
func (u_ UNMutableNotificationContent) SummaryArgumentCount() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("summaryArgumentCount"))
	return rv
}


// SetSummaryArgumentCount sets the value of the summaryArgumentCount property.
// The number the system adds to the notification summary when the notification represents multiple items.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/summaryArgumentCount
func (u_ UNMutableNotificationContent) SetSummaryArgumentCount(value uint) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSummaryArgumentCount:"), value)
}

// The value your app uses to determine which scene to display to handle the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/targetContentIdentifier
func (u_ UNMutableNotificationContent) TargetContentIdentifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("targetContentIdentifier"))
	return rv
}


// SetTargetContentIdentifier sets the value of the targetContentIdentifier property.
// The value your app uses to determine which scene to display to handle the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/targetContentIdentifier
func (u_ UNMutableNotificationContent) SetTargetContentIdentifier(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTargetContentIdentifier:"), objc.String(value))
}

// The identifier that groups related notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/threadIdentifier
func (u_ UNMutableNotificationContent) ThreadIdentifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("threadIdentifier"))
	return rv
}


// SetThreadIdentifier sets the value of the threadIdentifier property.
// The identifier that groups related notifications.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/threadIdentifier
func (u_ UNMutableNotificationContent) SetThreadIdentifier(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setThreadIdentifier:"), objc.String(value))
}

// The localized text that provides the notification’s primary description.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/title
func (u_ UNMutableNotificationContent) Title() string {
	rv := objc.Send[string](u_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The localized text that provides the notification’s primary description.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/title
func (u_ UNMutableNotificationContent) SetTitle(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// The custom data to associate with the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/userInfo
func (u_ UNMutableNotificationContent) UserInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("userInfo"))
	return rv
}


// SetUserInfo sets the value of the userInfo property.
// The custom data to associate with the notification.

//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/userInfo
func (u_ UNMutableNotificationContent) SetUserInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUserInfo:"), value)
}



