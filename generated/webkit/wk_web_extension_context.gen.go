// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WebExtensionContext] class.
var (
	WebExtensionContextClass     _WebExtensionContextClass
	WebExtensionContextClassOnce sync.Once
)

func getWebExtensionContextClass() _WebExtensionContextClass {
	WebExtensionContextClassOnce.Do(func() {
		WebExtensionContextClass = _WebExtensionContextClass{objc.GetClass("WKWebExtensionContext")}
	})
	return WebExtensionContextClass
}

type _WebExtensionContextClass struct {
	class objc.Class
}

// An interface definition for the [WebExtensionContext] class.
type IWebExtensionContext interface {
	objectivec.IObject
}

// An object that represents the runtime environment for a web extension.
//
// This class provides methods for managing the extension’s permissions, allowing it to inject content, run background logic, show popovers, and display other web-based UI to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionContext
type WebExtensionContext struct {
	objectivec.Object
}

// WebExtensionContextFrom constructs a [WebExtensionContext] from an unsafe.Pointer.
//
// An object that represents the runtime environment for a web extension.
func WebExtensionContextFrom(ptr unsafe.Pointer) WebExtensionContext {
	return WebExtensionContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WebExtensionContextClass) Alloc() WebExtensionContext {
	rv := objc.Send[WebExtensionContext](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WebExtensionContextClass) New() WebExtensionContext {
	rv := objc.Send[WebExtensionContext](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtensionContext) Init() WebExtensionContext {
	rv := objc.Send[WebExtensionContext](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtensionContext) Autorelease() WebExtensionContext {
	rv := objc.Send[WebExtensionContext](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtensionContext creates a new WebExtensionContext instance.
func NewWebExtensionContext() WebExtensionContext {
	return getWebExtensionContextClass().New()
}




