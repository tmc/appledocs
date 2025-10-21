// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFAddToHomeScreenInfo] class.
var (
	SFAddToHomeScreenInfoClass     _SFAddToHomeScreenInfoClass
	SFAddToHomeScreenInfoClassOnce sync.Once
)

func getSFAddToHomeScreenInfoClass() _SFAddToHomeScreenInfoClass {
	SFAddToHomeScreenInfoClassOnce.Do(func() {
		SFAddToHomeScreenInfoClass = _SFAddToHomeScreenInfoClass{objc.GetClass("SFAddToHomeScreenInfo")}
	})
	return SFAddToHomeScreenInfoClass
}

type _SFAddToHomeScreenInfoClass struct {
	class objc.Class
}

// An interface definition for the [SFAddToHomeScreenInfo] class.
type ISFAddToHomeScreenInfo interface {
	objectivec.IObject
}

// A class that provides information about a web app that someone adds to their Home Screen.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAddToHomeScreenInfo
type SFAddToHomeScreenInfo struct {
	objectivec.Object
}

// SFAddToHomeScreenInfoFrom constructs a [SFAddToHomeScreenInfo] from an unsafe.Pointer.
//
// A class that provides information about a web app that someone adds to their Home Screen.
func SFAddToHomeScreenInfoFrom(ptr unsafe.Pointer) SFAddToHomeScreenInfo {
	return SFAddToHomeScreenInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFAddToHomeScreenInfoClass) Alloc() SFAddToHomeScreenInfo {
	rv := objc.Send[SFAddToHomeScreenInfo](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFAddToHomeScreenInfoClass) New() SFAddToHomeScreenInfo {
	rv := objc.Send[SFAddToHomeScreenInfo](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFAddToHomeScreenInfo) Init() SFAddToHomeScreenInfo {
	rv := objc.Send[SFAddToHomeScreenInfo](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFAddToHomeScreenInfo) Autorelease() SFAddToHomeScreenInfo {
	rv := objc.Send[SFAddToHomeScreenInfo](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFAddToHomeScreenInfo creates a new SFAddToHomeScreenInfo instance.
func NewSFAddToHomeScreenInfo() SFAddToHomeScreenInfo {
	return getSFAddToHomeScreenInfoClass().New()
}




// Initializes a Home Screen information object with the supplied web app manifest.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAddToHomeScreenInfo/init(manifest:)
func NewSFAddToHomeScreenInfoWithManifest(manifest unsafe.Pointer) SFAddToHomeScreenInfo {
	instance := getSFAddToHomeScreenInfoClass().Alloc()
	rv := objc.Send[SFAddToHomeScreenInfo](instance.ID, objc.Sel("initWithManifest:"), manifest)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAddToHomeScreenInfo/manifest
func (s_ SFAddToHomeScreenInfo) Manifest() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("manifest"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAddToHomeScreenInfo/websiteCookies
func (s_ SFAddToHomeScreenInfo) WebsiteCookies() []foundation.HTTPCookie {
	rv := objc.Send[[]foundation.HTTPCookie](s_.ID, objc.Sel("websiteCookies"))
	return rv
}


// SetWebsiteCookies sets the value of the websiteCookies property.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAddToHomeScreenInfo/websiteCookies
func (s_ SFAddToHomeScreenInfo) SetWebsiteCookies(value []foundation.IHTTPCookie) {
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
	objc.Send[objc.ID](s_.ID, objc.Sel("setWebsiteCookies:"), nsArray)
}


