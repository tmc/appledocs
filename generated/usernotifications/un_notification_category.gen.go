// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UNNotificationCategory */

/* debug [class_header]: Header for UNNotificationCategory */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for UNNotificationCategory */
// An interface definition for the [UNNotificationCategory] class.
type IUNNotificationCategory interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for UNNotificationCategory */
	// properties:
	Actions() []UNNotificationAction
	CategorySummaryFormat() objc.IObject         /* cross-framework: NSString */
	HiddenPreviewsBodyPlaceholder() objc.IObject /* cross-framework: NSString */
	Identifier() objc.IObject                    /* cross-framework: NSString */
	IntentIdentifiers() []string
	Options() UNNotificationCategoryOptions
	CategoryIdentifier() objc.IObject /* cross-framework: NSString */
	SetCategoryIdentifier(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for UNNotificationCategory */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for UNNotificationCategory */
// Alloc allocates a new instance without initialization.
func (uc _UNNotificationCategoryClass) Alloc() UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for UNNotificationCategory */
// A type of notification your app supports and the custom actions that the system displays.
//
// A object defines a type of notification that your executable can receive. You create category objects to define your app’s — notifications that have action buttons the user can select in response to the notification. Each category object you create stores the actions and other behaviors associated with a specific type of notification. Register your category objects using the method of . You can register as many category objects as you need. To apply category objects to your notifications, include the category’s identifier string in the payload of any notifications you create. For local notifications, put this string in the property of the object that you use to specify the notification’s content. For remote notifications, use this string as the value of the key in the dictionary of your payload. Categories can have associated actions, which define custom buttons the system displays for notifications of that category. When the system has unlimited space, the system displays up to 10 actions. When the system has limited space, the system displays at most two actions.

// A type of notification your app supports and the custom actions that the system displays.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for UNNotificationCategory */

// Creates a category object containing the specified actions, options, placeholder text used when previews aren’t shown, and summary format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:categorySummaryFormat:options:)
func NewUNNotificationCategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderCategorySummaryFormatOptions(identifier objc.IObject /* cross-framework: NSString */, actions []UNNotificationAction, intentIdentifiers []string, hiddenPreviewsBodyPlaceholder objc.IObject /* cross-framework: NSString */, categorySummaryFormat objc.IObject /* cross-framework: NSString */, options UNNotificationCategoryOptions) UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](objc.ID(getUNNotificationCategoryClass().class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:categorySummaryFormat:options:"), identifier, actions, intentIdentifiers, hiddenPreviewsBodyPlaceholder, categorySummaryFormat, options)
	return rv
} /* debug [class_init_methods/constructor]: NewUNNotificationCategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderCategorySummaryFormatOptions */

// Creates a category object containing the specified actions, options, and placeholder text used when previews aren’t shown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:options:)
func NewUNNotificationCategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderOptions(identifier objc.IObject /* cross-framework: NSString */, actions []UNNotificationAction, intentIdentifiers []string, hiddenPreviewsBodyPlaceholder objc.IObject /* cross-framework: NSString */, options UNNotificationCategoryOptions) UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](objc.ID(getUNNotificationCategoryClass().class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:options:"), identifier, actions, intentIdentifiers, hiddenPreviewsBodyPlaceholder, options)
	return rv
} /* debug [class_init_methods/constructor]: NewUNNotificationCategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderOptions */

// Creates a category object containing the specified actions and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:options:)
func NewUNNotificationCategoryWithIdentifierActionsIntentIdentifiersOptions(identifier objc.IObject /* cross-framework: NSString */, actions []UNNotificationAction, intentIdentifiers []string, options UNNotificationCategoryOptions) UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](objc.ID(getUNNotificationCategoryClass().class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:options:"), identifier, actions, intentIdentifiers, options)
	return rv
} /* debug [class_init_methods/constructor]: NewUNNotificationCategoryWithIdentifierActionsIntentIdentifiersOptions */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for UNNotificationCategory */

// Creates a category object containing the specified actions, options, placeholder text used when previews aren’t shown, and summary format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:categorySummaryFormat:options:)
func (uc _UNNotificationCategoryClass) CategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderCategorySummaryFormatOptions(identifier objc.IObject /* cross-framework: NSString */, actions []UNNotificationAction, intentIdentifiers []string, hiddenPreviewsBodyPlaceholder objc.IObject /* cross-framework: NSString */, categorySummaryFormat objc.IObject /* cross-framework: NSString */, options UNNotificationCategoryOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:categorySummaryFormat:options:"), identifier, actions, intentIdentifiers, hiddenPreviewsBodyPlaceholder, categorySummaryFormat, options)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=CategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderCategorySummaryFormatOptions) */

// Creates a category object containing the specified actions, options, and placeholder text used when previews aren’t shown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:options:)
func (uc _UNNotificationCategoryClass) CategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderOptions(identifier objc.IObject /* cross-framework: NSString */, actions []UNNotificationAction, intentIdentifiers []string, hiddenPreviewsBodyPlaceholder objc.IObject /* cross-framework: NSString */, options UNNotificationCategoryOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:options:"), identifier, actions, intentIdentifiers, hiddenPreviewsBodyPlaceholder, options)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=CategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderOptions) */

// Creates a category object containing the specified actions and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:options:)
func (uc _UNNotificationCategoryClass) CategoryWithIdentifierActionsIntentIdentifiersOptions(identifier objc.IObject /* cross-framework: NSString */, actions []UNNotificationAction, intentIdentifiers []string, options UNNotificationCategoryOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:options:"), identifier, actions, intentIdentifiers, options)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=CategoryWithIdentifierActionsIntentIdentifiersOptions) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for UNNotificationCategory */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for UNNotificationCategory */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for UNNotificationCategory */

// The actions to display when the system delivers notifications of this type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/actions
func (u_ UNNotificationCategory) Actions() []UNNotificationAction {
	rv := objc.Send[[]UNNotificationAction](u_.ID, objc.Sel("actions"))
	return rv
} /* debug [instance_properties/getter]: actions */

// A format string for the summary description used when the system groups the category’s notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/categorySummaryFormat
func (u_ UNNotificationCategory) CategorySummaryFormat() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("categorySummaryFormat"))
	return rv
} /* debug [instance_properties/getter]: categorySummaryFormat */

// The placeholder text to display when the system disables notification previews for the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/hiddenPreviewsBodyPlaceholder
func (u_ UNNotificationCategory) HiddenPreviewsBodyPlaceholder() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("hiddenPreviewsBodyPlaceholder"))
	return rv
} /* debug [instance_properties/getter]: hiddenPreviewsBodyPlaceholder */

// The unique string assigned to the category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/identifier
func (u_ UNNotificationCategory) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("identifier"))
	return rv
} /* debug [instance_properties/getter]: identifier */

// The intents related to notifications of this category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/intentIdentifiers
func (u_ UNNotificationCategory) IntentIdentifiers() []string {
	rv := objc.Send[[]string](u_.ID, objc.Sel("intentIdentifiers"))
	return rv
} /* debug [instance_properties/getter]: intentIdentifiers */

// Options for how to handle notifications of this type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/options
func (u_ UNNotificationCategory) Options() UNNotificationCategoryOptions {
	rv := objc.Send[UNNotificationCategoryOptions](u_.ID, objc.Sel("options"))
	return rv
} /* debug [instance_properties/getter]: options */

// The identifier of the notification’s category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unmutablenotificationcontent/categoryidentifier
func (u_ UNNotificationCategory) CategoryIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("categoryIdentifier"))
	return rv
} /* debug [instance_properties/getter]: categoryIdentifier */

// The identifier of the notification’s category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unmutablenotificationcontent/categoryidentifier
func (u_ UNNotificationCategory) SetCategoryIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCategoryIdentifier:"), value)
} /* debug [instance_properties/setter]: categoryIdentifier */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class UNNotificationCategory */
