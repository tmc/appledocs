// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UNNotificationContent] class.
var (
	UNNotificationContentClass     _UNNotificationContentClass
	UNNotificationContentClassOnce sync.Once
)

func getUNNotificationContentClass() _UNNotificationContentClass {
	UNNotificationContentClassOnce.Do(func() {
		UNNotificationContentClass = _UNNotificationContentClass{objc.GetClass("UNNotificationContent")}
	})
	return UNNotificationContentClass
}

type _UNNotificationContentClass struct {
	class objc.Class
}

// An interface definition for the [UNNotificationContent] class.
type IUNNotificationContent interface {
	objectivec.IObject
	// properties:
	Attachments() []IUNNotificationAttachment
	Badge() objc.IObject /* cross-framework: NSNumber */
	Body() objc.IObject /* cross-framework: NSString */
	CategoryIdentifier() objc.IObject /* cross-framework: NSString */
	FilterCriteria() objc.IObject /* cross-framework: NSString */
	InterruptionLevel() UNNotificationInterruptionLevel
	RelevanceScore() float64
	Sound() IUNNotificationSound
	Subtitle() objc.IObject /* cross-framework: NSString */
	SummaryArgument() objc.IObject /* cross-framework: NSString */
	SummaryArgumentCount() uint
	TargetContentIdentifier() objc.IObject /* cross-framework: NSString */
	ThreadIdentifier() objc.IObject /* cross-framework: NSString */
	Title() objc.IObject /* cross-framework: NSString */
	UserInfo() objc.IObject /* cross-framework: NSDictionary */
	// methods:
	ContentByUpdatingWithProviderError(provider objectivec.IObject, outError unsafe.Pointer) IUNNotificationContent
}

// The uneditable content of a notification.
//
// A object contains the data associated with a notification. When your app receives a notification, the associated object contains an object of this type with the content that your app received. Use the content object to get the details of the notification, including the type of notification that the system delivered, any custom data you stored in the dictionary before scheduling the notification, and any attachments. Don’t create instances of this class directly. For remote notifications, the system derives the contents of this object from the JSON payload that your server sends to the APNS server. For local notifications, create a object, and configure the contents of that object instead.


// The uneditable content of a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent
type UNNotificationContent struct {
	objectivec.Object
}

// UNNotificationContentFrom constructs a [UNNotificationContent] from an unsafe.Pointer.
//
// The uneditable content of a notification.
func UNNotificationContentFrom(ptr unsafe.Pointer) UNNotificationContent {
	return UNNotificationContent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UNNotificationContentClass) Alloc() UNNotificationContent {
	rv := objc.Send[UNNotificationContent](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UNNotificationContentClass) New() UNNotificationContent {
	rv := objc.Send[UNNotificationContent](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNNotificationContent) Init() UNNotificationContent {
	rv := objc.Send[UNNotificationContent](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNNotificationContent) Autorelease() UNNotificationContent {
	rv := objc.Send[UNNotificationContent](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationContent creates a new UNNotificationContent instance.
func NewUNNotificationContent() UNNotificationContent {
	return getUNNotificationContentClass().New()
}



// Returns a copy of the notification that includes content from the specified provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/updating(from:)
func (u_ UNNotificationContent) ContentByUpdatingWithProviderError(provider objectivec.IObject, outError unsafe.Pointer) UNNotificationContent {
	rv := objc.Send[UNNotificationContent](u_.ID, objc.Sel("contentByUpdatingWithProvider:error:"), provider, outError)
	return rv
}


// The visual and audio attachments to display alongside the notification’s main content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/attachments
func (u_ UNNotificationContent) Attachments() []IUNNotificationAttachment {
	rv := objc.Send[[]UNNotificationAttachment](u_.ID, objc.Sel("attachments"))
	return rv
}


// The number that your app’s icon displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/badge
func (u_ UNNotificationContent) Badge() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](u_.ID, objc.Sel("badge"))
	return rv
}


// The localized text that provides the notification’s main content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/body
func (u_ UNNotificationContent) Body() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("body"))
	return rv
}


// The identifier of the notification’s category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/categoryIdentifier
func (u_ UNNotificationContent) CategoryIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("categoryIdentifier"))
	return rv
}


// The criteria the system evaluates to determine if it displays the notification in the current Focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/filterCriteria
func (u_ UNNotificationContent) FilterCriteria() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("filterCriteria"))
	return rv
}


// The notification’s importance and required delivery timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/interruptionLevel
func (u_ UNNotificationContent) InterruptionLevel() UNNotificationInterruptionLevel {
	rv := objc.Send[UNNotificationInterruptionLevel](u_.ID, objc.Sel("interruptionLevel"))
	return rv
}


// The score the system uses to determine if the notification is the summary’s featured notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/relevanceScore
func (u_ UNNotificationContent) RelevanceScore() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("relevanceScore"))
	return rv
}


// The sound that plays when the system delivers the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/sound
func (u_ UNNotificationContent) Sound() IUNNotificationSound {
	rv := objc.Send[UNNotificationSound](u_.ID, objc.Sel("sound"))
	return rv
}


// The localized text that provides the notification’s secondary description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/subtitle
func (u_ UNNotificationContent) Subtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("subtitle"))
	return rv
}


// The text the system adds to the notification summary to provide additional context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/summaryArgument
func (u_ UNNotificationContent) SummaryArgument() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("summaryArgument"))
	return rv
}


// The number the system adds to the notification summary when the notification represents multiple items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/summaryArgumentCount
func (u_ UNNotificationContent) SummaryArgumentCount() uint {
	rv := objc.Send[uint](u_.ID, objc.Sel("summaryArgumentCount"))
	return rv
}


// The value your app uses to determine which scene to display to handle the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/targetContentIdentifier
func (u_ UNNotificationContent) TargetContentIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("targetContentIdentifier"))
	return rv
}


// The identifier that groups related notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/threadIdentifier
func (u_ UNNotificationContent) ThreadIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("threadIdentifier"))
	return rv
}


// The localized text that provides the notification’s primary description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/title
func (u_ UNNotificationContent) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("title"))
	return rv
}


// The custom data to associate with the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/userInfo
func (u_ UNNotificationContent) UserInfo() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](u_.ID, objc.Sel("userInfo"))
	return rv
}


