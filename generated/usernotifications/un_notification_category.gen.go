// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UNNotificationCategory] class.
var (
	UNNotificationCategoryClass     _UNNotificationCategoryClass
	UNNotificationCategoryClassOnce sync.Once
)

func getUNNotificationCategoryClass() _UNNotificationCategoryClass {
	UNNotificationCategoryClassOnce.Do(func() {
		UNNotificationCategoryClass = _UNNotificationCategoryClass{objc.GetClass("UNNotificationCategory")}
	})
	return UNNotificationCategoryClass
}

type _UNNotificationCategoryClass struct {
	class objc.Class
}

// An interface definition for the [UNNotificationCategory] class.
type IUNNotificationCategory interface {
	objectivec.IObject
}

// A type of notification your app supports and the custom actions that the system displays.
//
// A object defines a type of notification that your executable can receive. You create category objects to define your app’s — notifications that have action buttons the user can select in response to the notification. Each category object you create stores the actions and other behaviors associated with a specific type of notification. Register your category objects using the method of . You can register as many category objects as you need. To apply category objects to your notifications, include the category’s identifier string in the payload of any notifications you create. For local notifications, put this string in the property of the object that you use to specify the notification’s content. For remote notifications, use this string as the value of the key in the dictionary of your payload. Categories can have associated actions, which define custom buttons the system displays for notifications of that category. When the system has unlimited space, the system displays up to 10 actions. When the system has limited space, the system displays at most two actions.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory
type UNNotificationCategory struct {
	objectivec.Object
}

// UNNotificationCategoryFrom constructs a [UNNotificationCategory] from an unsafe.Pointer.
//
// A type of notification your app supports and the custom actions that the system displays.
func UNNotificationCategoryFrom(ptr unsafe.Pointer) UNNotificationCategory {
	return UNNotificationCategory{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UNNotificationCategoryClass) Alloc() UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UNNotificationCategoryClass) New() UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNNotificationCategory) Init() UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNNotificationCategory) Autorelease() UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationCategory creates a new UNNotificationCategory instance.
func NewUNNotificationCategory() UNNotificationCategory {
	return getUNNotificationCategoryClass().New()
}




// Creates a category object containing the specified actions, options, placeholder text used when previews aren’t shown, and summary format string.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:categorySummaryFormat:options:)
func NewUNNotificationCategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderCategorySummaryFormatOptions(identifier appkit.string, actions []objc.ID, intentIdentifiers []string, hiddenPreviewsBodyPlaceholder appkit.string, categorySummaryFormat appkit.string, options UNNotificationCategoryOptions) UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](objc.ID(getUNNotificationCategoryClass().class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:categorySummaryFormat:options:"), identifier, actions, intentIdentifiers, hiddenPreviewsBodyPlaceholder, categorySummaryFormat, options)
	return rv
}



// Creates a category object containing the specified actions, options, and placeholder text used when previews aren’t shown.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:options:)
func NewUNNotificationCategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderOptions(identifier appkit.string, actions []objc.ID, intentIdentifiers []string, hiddenPreviewsBodyPlaceholder appkit.string, options UNNotificationCategoryOptions) UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](objc.ID(getUNNotificationCategoryClass().class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:options:"), identifier, actions, intentIdentifiers, hiddenPreviewsBodyPlaceholder, options)
	return rv
}



// Creates a category object containing the specified actions and options.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:options:)
func NewUNNotificationCategoryWithIdentifierActionsIntentIdentifiersOptions(identifier appkit.string, actions []objc.ID, intentIdentifiers []string, options UNNotificationCategoryOptions) UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](objc.ID(getUNNotificationCategoryClass().class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:options:"), identifier, actions, intentIdentifiers, options)
	return rv
}


// Creates a category object containing the specified actions, options, placeholder text used when previews aren’t shown, and summary format string.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:categorySummaryFormat:options:)
func (uc _UNNotificationCategoryClass) CategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderCategorySummaryFormatOptions(identifier appkit.string, actions []objc.ID, intentIdentifiers []string, hiddenPreviewsBodyPlaceholder appkit.string, categorySummaryFormat appkit.string, options UNNotificationCategoryOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:categorySummaryFormat:options:"), identifier, actions, intentIdentifiers, hiddenPreviewsBodyPlaceholder, categorySummaryFormat, options)
	return rv
}

// Creates a category object containing the specified actions, options, and placeholder text used when previews aren’t shown.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:options:)
func (uc _UNNotificationCategoryClass) CategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderOptions(identifier appkit.string, actions []objc.ID, intentIdentifiers []string, hiddenPreviewsBodyPlaceholder appkit.string, options UNNotificationCategoryOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:options:"), identifier, actions, intentIdentifiers, hiddenPreviewsBodyPlaceholder, options)
	return rv
}

// Creates a category object containing the specified actions and options.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:options:)
func (uc _UNNotificationCategoryClass) CategoryWithIdentifierActionsIntentIdentifiersOptions(identifier appkit.string, actions []objc.ID, intentIdentifiers []string, options UNNotificationCategoryOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:options:"), identifier, actions, intentIdentifiers, options)
	return rv
}

// The actions to display when the system delivers notifications of this type.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/actions
func (u_ UNNotificationCategory) Actions() []objc.ID {
	rv := objc.Send[[]objc.ID](u_.ID, objc.Sel("actions"))
	return rv
}

// A format string for the summary description used when the system groups the category’s notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/categorySummaryFormat
func (u_ UNNotificationCategory) CategorySummaryFormat() appkit.string {
	rv := objc.Send[appkit.string](u_.ID, objc.Sel("categorySummaryFormat"))
	return rv
}

// The placeholder text to display when the system disables notification previews for the app.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/hiddenPreviewsBodyPlaceholder
func (u_ UNNotificationCategory) HiddenPreviewsBodyPlaceholder() appkit.string {
	rv := objc.Send[appkit.string](u_.ID, objc.Sel("hiddenPreviewsBodyPlaceholder"))
	return rv
}

// The unique string assigned to the category.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/identifier
func (u_ UNNotificationCategory) Identifier() appkit.string {
	rv := objc.Send[appkit.string](u_.ID, objc.Sel("identifier"))
	return rv
}

// The intents related to notifications of this category.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/intentIdentifiers
func (u_ UNNotificationCategory) IntentIdentifiers() []string {
	rv := objc.Send[[]string](u_.ID, objc.Sel("intentIdentifiers"))
	return rv
}

// Options for how to handle notifications of this type.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/options
func (u_ UNNotificationCategory) Options() UNNotificationCategoryOptions {
	rv := objc.Send[UNNotificationCategoryOptions](u_.ID, objc.Sel("options"))
	return rv
}

// The identifier of the notification’s category.
//
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unmutablenotificationcontent/categoryidentifier
func (u_ UNNotificationCategory) CategoryIdentifier() appkit.string {
	rv := objc.Send[appkit.string](u_.ID, objc.Sel("categoryIdentifier"))
	return rv
}


// SetCategoryIdentifier sets the value of the categoryIdentifier property.
// The identifier of the notification’s category.

//
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unmutablenotificationcontent/categoryidentifier
func (u_ UNNotificationCategory) SetCategoryIdentifier(value appkit.string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCategoryIdentifier:"), value)
}


