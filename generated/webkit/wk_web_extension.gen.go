// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WebExtension] class.
var (
	WebExtensionClass     _WebExtensionClass
	WebExtensionClassOnce sync.Once
)

func getWebExtensionClass() _WebExtensionClass {
	WebExtensionClassOnce.Do(func() {
		WebExtensionClass = _WebExtensionClass{objc.GetClass("WKWebExtension")}
	})
	return WebExtensionClass
}

type _WebExtensionClass struct {
	class objc.Class
}

// An interface definition for the [WebExtension] class.
type IWebExtension interface {
	objectivec.IObject
}

// An object that encapsulates a web extension’s resources that the manifest file defines.
//
// This class reads and parses the file along with the supporting resources like icons and localizations.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension
type WebExtension struct {
	objectivec.Object
}

// WebExtensionFrom constructs a [WebExtension] from an unsafe.Pointer.
//
// An object that encapsulates a web extension’s resources that the manifest file defines.
func WebExtensionFrom(ptr unsafe.Pointer) WebExtension {
	return WebExtension{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WebExtensionClass) Alloc() WebExtension {
	rv := objc.Send[WebExtension](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WebExtensionClass) New() WebExtension {
	rv := objc.Send[WebExtension](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtension) Init() WebExtension {
	rv := objc.Send[WebExtension](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtension) Autorelease() WebExtension {
	rv := objc.Send[WebExtension](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtension creates a new WebExtension instance.
func NewWebExtension() WebExtension {
	return getWebExtensionClass().New()
}




