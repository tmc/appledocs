// Code generated from Apple documentation for ScreenTime. DO NOT EDIT.

package screentime

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
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
}

// The controller you use to report web usage and block restricted webpages.
//
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


// The URL for the webpage.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/url
func (s_ STWebpageController) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("URL"))
	return rv
}


// SetURL sets the value of the URL property.
// The URL for the webpage.

//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/url
func (s_ STWebpageController) SetURL(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setURL:"), value)
}

// A Boolean that indicates whether a parent or guardian has blocked the URL.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STWebpageController/urlIsBlocked
func (s_ STWebpageController) URLIsBlocked() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("URLIsBlocked"))
	return rv
}




