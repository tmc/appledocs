// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFSafariPageProperties] class.
var (
	SFSafariPagePropertiesClass     _SFSafariPagePropertiesClass
	SFSafariPagePropertiesClassOnce sync.Once
)

func getSFSafariPagePropertiesClass() _SFSafariPagePropertiesClass {
	SFSafariPagePropertiesClassOnce.Do(func() {
		SFSafariPagePropertiesClass = _SFSafariPagePropertiesClass{objc.GetClass("SFSafariPageProperties")}
	})
	return SFSafariPagePropertiesClass
}

type _SFSafariPagePropertiesClass struct {
	class objc.Class
}

// An interface definition for the [SFSafariPageProperties] class.
type ISFSafariPageProperties interface {
	objectivec.IObject
}

// An object that captures information about a webpage.
//
// Use the properties object to retrieve page information, such as the current URL, page title, active status, and private browsing status.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPageProperties
type SFSafariPageProperties struct {
	objectivec.Object
}

// SFSafariPagePropertiesFrom constructs a [SFSafariPageProperties] from an unsafe.Pointer.
//
// An object that captures information about a webpage.
func SFSafariPagePropertiesFrom(ptr unsafe.Pointer) SFSafariPageProperties {
	return SFSafariPageProperties{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSafariPagePropertiesClass) Alloc() SFSafariPageProperties {
	rv := objc.Send[SFSafariPageProperties](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSafariPagePropertiesClass) New() SFSafariPageProperties {
	rv := objc.Send[SFSafariPageProperties](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariPageProperties) Init() SFSafariPageProperties {
	rv := objc.Send[SFSafariPageProperties](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariPageProperties) Autorelease() SFSafariPageProperties {
	rv := objc.Send[SFSafariPageProperties](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariPageProperties creates a new SFSafariPageProperties instance.
func NewSFSafariPageProperties() SFSafariPageProperties {
	return getSFSafariPagePropertiesClass().New()
}


// A Boolean value that indicates whether the page is currently active.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPageProperties/isActive
func (s_ SFSafariPageProperties) Active() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("active"))
	return rv
}

// The title of the page.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPageProperties/title
func (s_ SFSafariPageProperties) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("title"))
	return rv
}

// Indicates the URL of the page.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPageProperties/url
func (s_ SFSafariPageProperties) Url() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("url"))
	return rv
}

// A Boolean value that indicates whether the page is using Safari Private Browsing.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPageProperties/usesPrivateBrowsing
func (s_ SFSafariPageProperties) UsesPrivateBrowsing() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("usesPrivateBrowsing"))
	return rv
}



