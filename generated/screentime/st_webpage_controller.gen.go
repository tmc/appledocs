// Code generated from Apple documentation for ScreenTime. DO NOT EDIT.

package screentime

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [STWebpageController] class.
type ISTWebpageController interface {
	appkit.IViewController
	// properties:
	ProfileIdentifier() unsafe.Pointer
	SetProfileIdentifier(value unsafe.Pointer)
	SuppressUsageRecording() bool
	SetSuppressUsageRecording(value bool)
	Url() objc.IObject /* cross-framework: URL */
	SetUrl(value objc.IObject /* cross-framework: URL */)
	UrlIsBlocked() bool
	SetUrlIsBlocked(value bool)
	UrlIsPictureInPicture() bool
	SetUrlIsPictureInPicture(value bool)
	UrlIsPlayingVideo() bool
	SetUrlIsPlayingVideo(value bool)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (sc _STWebpageControllerClass) Alloc() STWebpageController {
	rv := objc.Send[STWebpageController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An optional identifier for the current browsing profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stwebpagecontroller/profileidentifier
func (s_ STWebpageController) ProfileIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("profileIdentifier"))
	return rv
}


// An optional identifier for the current browsing profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stwebpagecontroller/profileidentifier
func (s_ STWebpageController) SetProfileIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setProfileIdentifier:"), value)
}


// A Boolean that indicates whether the webpage controller is not recording web
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stwebpagecontroller/suppressusagerecording
func (s_ STWebpageController) SuppressUsageRecording() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("suppressUsageRecording"))
	return rv
}


// A Boolean that indicates whether the webpage controller is not recording web
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stwebpagecontroller/suppressusagerecording
func (s_ STWebpageController) SetSuppressUsageRecording(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuppressUsageRecording:"), value)
}


// The URL for the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stwebpagecontroller/url
func (s_ STWebpageController) Url() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](s_.ID, objc.Sel("url"))
	return rv
}


// The URL for the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stwebpagecontroller/url
func (s_ STWebpageController) SetUrl(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUrl:"), value)
}


// A Boolean that indicates whether a parent or guardian has blocked the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stwebpagecontroller/urlisblocked
func (s_ STWebpageController) UrlIsBlocked() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("urlIsBlocked"))
	return rv
}


// A Boolean that indicates whether a parent or guardian has blocked the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stwebpagecontroller/urlisblocked
func (s_ STWebpageController) SetUrlIsBlocked(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUrlIsBlocked:"), value)
}


// A Boolean that indicates whether the webpage is currently displaying a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stwebpagecontroller/urlispictureinpicture
func (s_ STWebpageController) UrlIsPictureInPicture() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("urlIsPictureInPicture"))
	return rv
}


// A Boolean that indicates whether the webpage is currently displaying a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stwebpagecontroller/urlispictureinpicture
func (s_ STWebpageController) SetUrlIsPictureInPicture(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUrlIsPictureInPicture:"), value)
}


// A Boolean that indicates whether there are one or more videos currently
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stwebpagecontroller/urlisplayingvideo
func (s_ STWebpageController) UrlIsPlayingVideo() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("urlIsPlayingVideo"))
	return rv
}


// A Boolean that indicates whether there are one or more videos currently
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stwebpagecontroller/urlisplayingvideo
func (s_ STWebpageController) SetUrlIsPlayingVideo(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUrlIsPlayingVideo:"), value)
}




