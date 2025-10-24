// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKNotificationBanner */


/* debug [class_header]: Header for GKNotificationBanner */
// The class instance for the [NotificationBanner] class.
var (
	NotificationBannerClass     _NotificationBannerClass
	NotificationBannerClassOnce sync.Once
)

func getNotificationBannerClass() _NotificationBannerClass {
	NotificationBannerClassOnce.Do(func() {
		NotificationBannerClass = _NotificationBannerClass{objc.GetClass("GKNotificationBanner")}
	})
	return NotificationBannerClass
}

type _NotificationBannerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NotificationBanner */
// An interface definition for the [NotificationBanner] class.
type INotificationBanner interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NotificationBanner */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NotificationBanner */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NotificationBanner */
// Alloc allocates a new instance without initialization.
func (nc _NotificationBannerClass) Alloc() NotificationBanner {
	rv := objc.Send[NotificationBanner](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NotificationBannerClass) New() NotificationBanner {
	rv := objc.Send[NotificationBanner](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NotificationBanner) Init() NotificationBanner {
	rv := objc.Send[NotificationBanner](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NotificationBanner) Autorelease() NotificationBanner {
	rv := objc.Send[NotificationBanner](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNotificationBanner creates a new NotificationBanner instance.
func NewNotificationBanner() NotificationBanner {
	return getNotificationBannerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NotificationBanner */
// A Game Center-style banner that displays a message to the local player.
//
// This class displays a message in a banner to the local player, similar to the banner that GameKit displays when a player earns an achievement. If the game is in the foreground, the banner appears immediately. If the game is in the background, the banner appears when the game becomes active. To display the banner with your message, use the method. To specify a duration that GameKit presents the banner, use the method instead. Optionally, pass these methods a completion handler that GameKit calls after it dismisses the banner.


// A Game Center-style banner that displays a message to the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKNotificationBanner
type NotificationBanner struct {
	objectivec.Object
}

// NotificationBannerFrom constructs a [NotificationBanner] from an unsafe.Pointer.
//
// A Game Center-style banner that displays a message to the local player.
func NotificationBannerFrom(ptr unsafe.Pointer) NotificationBanner {
	return NotificationBanner{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NotificationBanner *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NotificationBanner */

// Displays a banner with a title and message to the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKNotificationBanner/show(withTitle:message:completionHandler:)
func (nc _NotificationBannerClass) ShowBannerWithTitleMessageCompletionHandler(title objc.IObject /* cross-framework: NSString */, message objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("showBannerWithTitle:message:completionHandler:"), title, message, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ShowBannerWithTitleMessageCompletionHandler) */


// Displays a banner to the player for a specified period of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKNotificationBanner/show(withTitle:message:duration:completionHandler:)
func (nc _NotificationBannerClass) ShowBannerWithTitleMessageDurationCompletionHandler(title objc.IObject /* cross-framework: NSString */, message objc.IObject /* cross-framework: NSString */, duration float64, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("showBannerWithTitle:message:duration:completionHandler:"), title, message, duration, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ShowBannerWithTitleMessageDurationCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NotificationBanner */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NotificationBanner */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NotificationBanner */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKNotificationBanner */



