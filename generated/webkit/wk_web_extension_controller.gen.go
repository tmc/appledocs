// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WebExtensionController] class.
var (
	WebExtensionControllerClass     _WebExtensionControllerClass
	WebExtensionControllerClassOnce sync.Once
)

func getWebExtensionControllerClass() _WebExtensionControllerClass {
	WebExtensionControllerClassOnce.Do(func() {
		WebExtensionControllerClass = _WebExtensionControllerClass{objc.GetClass("WKWebExtensionController")}
	})
	return WebExtensionControllerClass
}

type _WebExtensionControllerClass struct {
	class objc.Class
}

// An interface definition for the [WebExtensionController] class.
type IWebExtensionController interface {
	objectivec.IObject
	UnloadExtensionContextError(extensionContext unsafe.Pointer, error_ unsafe.Pointer) bool
}

// An object that manages a set of loaded extension contexts.
//
// You can have one or more extension controller instances, allowing different parts of the app to use different sets of extensions. You can associate a controller with using the property on .
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController
type WebExtensionController struct {
	objectivec.Object
}

// WebExtensionControllerFrom constructs a [WebExtensionController] from an unsafe.Pointer.
//
// An object that manages a set of loaded extension contexts.
func WebExtensionControllerFrom(ptr unsafe.Pointer) WebExtensionController {
	return WebExtensionController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WebExtensionControllerClass) Alloc() WebExtensionController {
	rv := objc.Send[WebExtensionController](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WebExtensionControllerClass) New() WebExtensionController {
	rv := objc.Send[WebExtensionController](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtensionController) Init() WebExtensionController {
	rv := objc.Send[WebExtensionController](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtensionController) Autorelease() WebExtensionController {
	rv := objc.Send[WebExtensionController](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtensionController creates a new WebExtensionController instance.
func NewWebExtensionController() WebExtensionController {
	return getWebExtensionControllerClass().New()
}


// Unloads the specified extension context.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtensionController/unload(_:)
func (w_ WebExtensionController) UnloadExtensionContextError(extensionContext unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("unloadExtensionContext:error:"), extensionContext, error_)
	return rv
}

// The extension controller delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontroller/delegate
func (w_ WebExtensionController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The extension controller delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontroller/delegate
func (w_ WebExtensionController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDelegate:"), value)
}

// A copy of the configuration with which the web extension controller was initialized.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontroller/configuration-swift.property
func (w_ WebExtensionController) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("configuration"))
	return rv
}


// SetConfiguration sets the value of the configuration property.
// A copy of the configuration with which the web extension controller was initialized.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontroller/configuration-swift.property
func (w_ WebExtensionController) SetConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setConfiguration:"), value)
}

// A set of all the currently loaded extension contexts.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontroller/extensioncontexts
func (w_ WebExtensionController) ExtensionContexts() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("extensionContexts"))
	return rv
}


// SetExtensionContexts sets the value of the extensionContexts property.
// A set of all the currently loaded extension contexts.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontroller/extensioncontexts
func (w_ WebExtensionController) SetExtensionContexts(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setExtensionContexts:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/webextensioncontroller
func (w_ WebExtensionController) WebExtensionController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("webExtensionController"))
	return rv
}


// SetWebExtensionController sets the value of the webExtensionController property.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/webextensioncontroller
func (w_ WebExtensionController) SetWebExtensionController(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWebExtensionController:"), value)
}

// A set of all the currently loaded extensions.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontroller/extensions
func (w_ WebExtensionController) Extensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("extensions"))
	return rv
}


// SetExtensions sets the value of the extensions property.
// A set of all the currently loaded extensions.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextensioncontroller/extensions
func (w_ WebExtensionController) SetExtensions(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setExtensions:"), value)
}



