// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [WebExtensionMessagePort] class.
var (
	WebExtensionMessagePortClass     _WebExtensionMessagePortClass
	WebExtensionMessagePortClassOnce sync.Once
)

func getWebExtensionMessagePortClass() _WebExtensionMessagePortClass {
	WebExtensionMessagePortClassOnce.Do(func() {
		WebExtensionMessagePortClass = _WebExtensionMessagePortClass{objc.GetClass("WKWebExtensionMessagePort")}
	})
	return WebExtensionMessagePortClass
}

type _WebExtensionMessagePortClass struct {
	class objc.Class
}

// An interface definition for the [WebExtensionMessagePort] class.
type IWebExtensionMessagePort interface {
	objectivec.IObject
}

// An object that manages message-based communication with a web extension.
//
// Contains properties and methods to handle message exchanges with a web extension.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort
type WebExtensionMessagePort struct {
	objectivec.Object
}

// WebExtensionMessagePortFrom constructs a [WebExtensionMessagePort] from an unsafe.Pointer.
//
// An object that manages message-based communication with a web extension.
func WebExtensionMessagePortFrom(ptr unsafe.Pointer) WebExtensionMessagePort {
	return WebExtensionMessagePort{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WebExtensionMessagePortClass) Alloc() WebExtensionMessagePort {
	rv := objc.Send[WebExtensionMessagePort](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WebExtensionMessagePortClass) New() WebExtensionMessagePort {
	rv := objc.Send[WebExtensionMessagePort](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtensionMessagePort) Init() WebExtensionMessagePort {
	rv := objc.Send[WebExtensionMessagePort](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtensionMessagePort) Autorelease() WebExtensionMessagePort {
	rv := objc.Send[WebExtensionMessagePort](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtensionMessagePort creates a new WebExtensionMessagePort instance.
func NewWebExtensionMessagePort() WebExtensionMessagePort {
	return getWebExtensionMessagePortClass().New()
}




