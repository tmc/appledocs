// Code generated from Apple documentation for ScreenTime. DO NOT EDIT.

package screentime

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class STWebpageController */


/* debug [class_header]: Header for STWebpageController */
// The class instance for the [STWebpageController] class.
var (
	STWebpageControllerClass     _STWebpageControllerClass
	STWebpageControllerClassOnce sync.Once
)

func getSTWebpageControllerClass() _STWebpageControllerClass {
	STWebpageControllerClassOnce.Do(func() {
		STWebpageControllerClass = _STWebpageControllerClass{objc.GetClass("STWebpageController")}
	})
	return STWebpageControllerClass
}

type _STWebpageControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for STWebpageController */
// An interface definition for the [STWebpageController] class.
type ISTWebpageController interface {
	appkit.IViewController
	
/* debug [class_interface_properties]: Properties for STWebpageController */
	// properties:
	ProfileIdentifier() STWebHistoryProfileIdentifier /* typedef */
	SetProfileIdentifier(value STWebHistoryProfileIdentifier /* typedef */)
	SuppressUsageRecording() bool
	SetSuppressUsageRecording(value bool)
	URL() objc.IObject /* cross-framework: NSURL */
	SetURL(value objc.IObject /* cross-framework: NSURL */)
	URLIsBlocked() bool
	URLIsPictureInPicture() bool
	SetURLIsPictureInPicture(value bool)
	URLIsPlayingVideo() bool
	SetURLIsPlayingVideo(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for STWebpageController */
	// methods:
	SetBundleIdentifierError(bundleIdentifier objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for STWebpageController */
// Alloc allocates a new instance without initialization.
func (sc _STWebpageControllerClass) Alloc() STWebpageController {
	rv := objc.Send[STWebpageController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _STWebpageControllerClass) New() STWebpageController {
	rv := objc.Send[STWebpageController](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ STWebpageController) Init() STWebpageController {
	rv := objc.Send[STWebpageController](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ STWebpageController) Autorelease() STWebpageController {
	rv := objc.Send[STWebpageController](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSTWebpageController creates a new STWebpageController instance.
func NewSTWebpageController() STWebpageController {
	return getSTWebpageControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for STWebpageController */
// The controller you use to report web usage and block restricted webpages.


// The controller you use to report web usage and block restricted webpages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController
type STWebpageController struct {
	appkit.ViewController
}

// STWebpageControllerFrom constructs a [STWebpageController] from an unsafe.Pointer.
//
// The controller you use to report web usage and block restricted webpages.
func STWebpageControllerFrom(ptr unsafe.Pointer) STWebpageController {
	return STWebpageController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for STWebpageController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for STWebpageController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for STWebpageController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for STWebpageController */

// Changes the bundle identifier used to report web usage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/setBundleIdentifier(_:)
func (s_ STWebpageController) SetBundleIdentifierError(bundleIdentifier objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("setBundleIdentifier:error:"), bundleIdentifier, error_)
	return rv
}/* debug [instance_methods/method]: SetBundleIdentifierError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for STWebpageController */

// An optional identifier for the current browsing profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/profileIdentifier
func (s_ STWebpageController) ProfileIdentifier() STWebHistoryProfileIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("profileIdentifier"))
	return rv
}/* debug [instance_properties/getter]: profileIdentifier */


// An optional identifier for the current browsing profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/profileIdentifier
func (s_ STWebpageController) SetProfileIdentifier(value STWebHistoryProfileIdentifier /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setProfileIdentifier:"), value)
}/* debug [instance_properties/setter]: profileIdentifier */


// A Boolean that indicates whether the webpage controller is not recording web usage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/suppressUsageRecording
func (s_ STWebpageController) SuppressUsageRecording() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("suppressUsageRecording"))
	return rv
}/* debug [instance_properties/getter]: suppressUsageRecording */


// A Boolean that indicates whether the webpage controller is not recording web usage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/suppressUsageRecording
func (s_ STWebpageController) SetSuppressUsageRecording(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuppressUsageRecording:"), value)
}/* debug [instance_properties/setter]: suppressUsageRecording */


// The URL for the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/url
func (s_ STWebpageController) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// The URL for the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/url
func (s_ STWebpageController) SetURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setURL:"), value)
}/* debug [instance_properties/setter]: URL */


// A Boolean that indicates whether a parent or guardian has blocked the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/urlIsBlocked
func (s_ STWebpageController) URLIsBlocked() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("URLIsBlocked"))
	return rv
}/* debug [instance_properties/getter]: URLIsBlocked */


// A Boolean that indicates whether the webpage is currently displaying a floating picture in picture window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/urlIsPictureInPicture
func (s_ STWebpageController) URLIsPictureInPicture() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("URLIsPictureInPicture"))
	return rv
}/* debug [instance_properties/getter]: URLIsPictureInPicture */


// A Boolean that indicates whether the webpage is currently displaying a floating picture in picture window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/urlIsPictureInPicture
func (s_ STWebpageController) SetURLIsPictureInPicture(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setURLIsPictureInPicture:"), value)
}/* debug [instance_properties/setter]: URLIsPictureInPicture */


// A Boolean that indicates whether there are one or more videos currently playing in the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/urlIsPlayingVideo
func (s_ STWebpageController) URLIsPlayingVideo() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("URLIsPlayingVideo"))
	return rv
}/* debug [instance_properties/getter]: URLIsPlayingVideo */


// A Boolean that indicates whether there are one or more videos currently playing in the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/urlIsPlayingVideo
func (s_ STWebpageController) SetURLIsPlayingVideo(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setURLIsPlayingVideo:"), value)
}/* debug [instance_properties/setter]: URLIsPlayingVideo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class STWebpageController */






