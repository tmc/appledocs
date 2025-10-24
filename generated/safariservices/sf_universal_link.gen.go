// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFUniversalLink] class.
var (
	SFUniversalLinkClass     _SFUniversalLinkClass
	SFUniversalLinkClassOnce sync.Once
)

func getSFUniversalLinkClass() _SFUniversalLinkClass {
	SFUniversalLinkClassOnce.Do(func() {
		SFUniversalLinkClass = _SFUniversalLinkClass{objc.GetClass("SFUniversalLink")}
	})
	return SFUniversalLinkClass
}

type _SFUniversalLinkClass struct {
	class objc.Class
}

// An interface definition for the [SFUniversalLink] class.
type ISFUniversalLink interface {
	objectivec.IObject
	// properties:
	ApplicationURL() objc.IObject /* cross-framework: NSURL */
	Enabled() bool
	SetEnabled(value bool)
	WebpageURL() objc.IObject /* cross-framework: NSURL */
	IsEnabled() bool
	SetIsEnabled(value bool)
	// methods:
}

// An object that provides browsers with the ability to discover associations between an app and a website.
//
// Universal links are a bridge between an app and a website that have related content, such as products or services. Typically, clicking a link in a browser takes a person to a website. However, the person may have an app that provides the same content and a better experience. A web browser uses the class to discover such applications and provide the person with additional options for interaction beyond the default browser behavior. In order to use universal links, you need to use the entitlement with a value of . Before you submit an app with the entitlement to the App Store, you need to get permission to use the entitlement. Request permission at .


// An object that provides browsers with the ability to discover associations between an app and a website.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFUniversalLink
type SFUniversalLink struct {
	objectivec.Object
}

// SFUniversalLinkFrom constructs a [SFUniversalLink] from an unsafe.Pointer.
//
// An object that provides browsers with the ability to discover associations between an app and a website.
func SFUniversalLinkFrom(ptr unsafe.Pointer) SFUniversalLink {
	return SFUniversalLink{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFUniversalLinkClass) Alloc() SFUniversalLink {
	rv := objc.Send[SFUniversalLink](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFUniversalLinkClass) New() SFUniversalLink {
	rv := objc.Send[SFUniversalLink](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFUniversalLink) Init() SFUniversalLink {
	rv := objc.Send[SFUniversalLink](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFUniversalLink) Autorelease() SFUniversalLink {
	rv := objc.Send[SFUniversalLink](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFUniversalLink creates a new SFUniversalLink instance.
func NewSFUniversalLink() SFUniversalLink {
	return getSFUniversalLinkClass().New()
}



// Creates a universal link object with the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFUniversalLink/init(webpageURL:)
func NewSFUniversalLinkWithWebpageURL(url objc.IObject /* cross-framework: NSURL */) SFUniversalLink {
	instance := getSFUniversalLinkClass().Alloc()
	rv := objc.Send[SFUniversalLink](instance.ID, objc.Sel("initWithWebpageURL:"), url)
	rv.Autorelease()
	return rv
}



// The URL to the app that can open this universal link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFUniversalLink/applicationURL
func (s_ SFUniversalLink) ApplicationURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("applicationURL"))
	return rv
}


// A flag that indicates whether the universal link is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFUniversalLink/isEnabled
func (s_ SFUniversalLink) Enabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enabled"))
	return rv
}


// A flag that indicates whether the universal link is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFUniversalLink/isEnabled
func (s_ SFUniversalLink) SetEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnabled:"), value)
}


// The URL specified when initializing the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFUniversalLink/webpageURL
func (s_ SFUniversalLink) WebpageURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("webpageURL"))
	return rv
}


// A flag that indicates whether the universal link is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfuniversallink/isenabled
func (s_ SFUniversalLink) IsEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabled"))
	return rv
}


// A flag that indicates whether the universal link is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfuniversallink/isenabled
func (s_ SFUniversalLink) SetIsEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsEnabled:"), value)
}


