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
	// properties:
	Attachments() []IUNNotificationAttachment
	SetAttachments(value []IUNNotificationAttachment)
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
	// methods:
}

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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/attachments
func (u_ UNMutableNotificationContent) Attachments() []IUNNotificationAttachment {
	rv := objc.Send[[]UNNotificationAttachment](u_.ID, objc.Sel("attachments"))
	return rv
}


// The visual and audio attachments to display alongside the notification’s main content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/attachments
func (u_ UNMutableNotificationContent) SetAttachments(value []IUNNotificationAttachment) {
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/badge
func (u_ UNMutableNotificationContent) Badge() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](u_.ID, objc.Sel("badge"))
	return rv
}


// The number that your app’s icon displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/badge
func (u_ UNMutableNotificationContent) SetBadge(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setBadge:"), value)
}


// The localized text that provides the notification’s main content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/body
func (u_ UNMutableNotificationContent) Body() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("body"))
	return rv
}


// The localized text that provides the notification’s main content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/body
func (u_ UNMutableNotificationContent) SetBody(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setBody:"), value)
}


// The identifier of the notification’s category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/categoryIdentifier
func (u_ UNMutableNotificationContent) CategoryIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("categoryIdentifier"))
	return rv
}


// The identifier of the notification’s category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/categoryIdentifier
func (u_ UNMutableNotificationContent) SetCategoryIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCategoryIdentifier:"), value)
}


// The criteria the system evaluates to determine if it displays the notification in the current Focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/filterCriteria
func (u_ UNMutableNotificationContent) FilterCriteria() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("filterCriteria"))
	return rv
}


// The criteria the system evaluates to determine if it displays the notification in the current Focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/filterCriteria
func (u_ UNMutableNotificationContent) SetFilterCriteria(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setFilterCriteria:"), value)
}


// The notification’s importance and required delivery timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/interruptionLevel
func (u_ UNMutableNotificationContent) InterruptionLevel() UNNotificationInterruptionLevel {
	rv := objc.Send[UNNotificationInterruptionLevel](u_.ID, objc.Sel("interruptionLevel"))
	return rv
}


// The notification’s importance and required delivery timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/interruptionLevel
func (u_ UNMutableNotificationContent) SetInterruptionLevel(value UNNotificationInterruptionLevel) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setInterruptionLevel:"), value)
}


// The score the system uses to determine if the notification is the summary’s featured notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/relevanceScore
func (u_ UNMutableNotificationContent) RelevanceScore() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("relevanceScore"))
	return rv
}


// The score the system uses to determine if the notification is the summary’s featured notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/relevanceScore
func (u_ UNMutableNotificationContent) SetRelevanceScore(value float64) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRelevanceScore:"), value)
}


// The sound that plays when the system delivers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/sound
func (u_ UNMutableNotificationContent) Sound() IUNNotificationSound {
	rv := objc.Send[UNNotificationSound](u_.ID, objc.Sel("sound"))
	return rv
}


// The sound that plays when the system delivers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/sound
func (u_ UNMutableNotificationContent) SetSound(value IUNNotificationSound) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSound:"), value)
}


// The localized text that provides the notification’s secondary description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/subtitle
func (u_ UNMutableNotificationContent) Subtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("subtitle"))
	return rv
}


// The localized text that provides the notification’s secondary description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/subtitle
func (u_ UNMutableNotificationContent) SetSubtitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSubtitle:"), value)
}


// The text the system adds to the notification summary to provide additional context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/summaryArgument
func (u_ UNMutableNotificationContent) SummaryArgument() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("summaryArgument"))
	return rv
}


// The text the system adds to the notification summary to provide additional context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/summaryArgument
func (u_ UNMutableNotificationContent) SetSummaryArgument(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSummaryArgument:"), value)
}


// The number the system adds to the notification summary when the notification represents multiple items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/summaryArgumentCount
func (u_ UNMutableNotificationContent) SummaryArgumentCount() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("summaryArgumentCount"))
	return rv
}


// The number the system adds to the notification summary when the notification represents multiple items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/summaryArgumentCount
func (u_ UNMutableNotificationContent) SetSummaryArgumentCount(value uint) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSummaryArgumentCount:"), value)
}


// The value your app uses to determine which scene to display to handle the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/targetContentIdentifier
func (u_ UNMutableNotificationContent) TargetContentIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("targetContentIdentifier"))
	return rv
}


// The value your app uses to determine which scene to display to handle the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/targetContentIdentifier
func (u_ UNMutableNotificationContent) SetTargetContentIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTargetContentIdentifier:"), value)
}


// The identifier that groups related notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/threadIdentifier
func (u_ UNMutableNotificationContent) ThreadIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("threadIdentifier"))
	return rv
}


// The identifier that groups related notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/threadIdentifier
func (u_ UNMutableNotificationContent) SetThreadIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setThreadIdentifier:"), value)
}


// The localized text that provides the notification’s primary description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/title
func (u_ UNMutableNotificationContent) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("title"))
	return rv
}


// The localized text that provides the notification’s primary description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/title
func (u_ UNMutableNotificationContent) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTitle:"), value)
}


// The custom data to associate with the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/userInfo
func (u_ UNMutableNotificationContent) UserInfo() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](u_.ID, objc.Sel("userInfo"))
	return rv
}


// The custom data to associate with the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/userInfo
func (u_ UNMutableNotificationContent) SetUserInfo(value objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUserInfo:"), value)
}


