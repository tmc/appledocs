// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ExtensionContext] class.
var (
	ExtensionContextClass     _ExtensionContextClass
	ExtensionContextClassOnce sync.Once
)

func getExtensionContextClass() _ExtensionContextClass {
	ExtensionContextClassOnce.Do(func() {
		ExtensionContextClass = _ExtensionContextClass{objc.GetClass("NSExtensionContext")}
	})
	return ExtensionContextClass
}

type _ExtensionContextClass struct {
	class objc.Class
}

// An interface definition for the [ExtensionContext] class.
type IExtensionContext interface {
	objectivec.IObject
	// properties:
	HostedViewMaximumAllowedSize() objc.IObject /* cross-framework: Size */
	SetHostedViewMaximumAllowedSize(value objc.IObject /* cross-framework: Size */)
	WidgetActiveDisplayMode() unsafe.Pointer
	SetWidgetActiveDisplayMode(value unsafe.Pointer)
	WidgetLargestAvailableDisplayMode() unsafe.Pointer
	SetWidgetLargestAvailableDisplayMode(value unsafe.Pointer)
	NSExtensionItemsAndErrorsKey() IString
	// methods:
}

// The host app context from which an app extension is invoked.
//
// When a host app sends a request to an app extension, it provides an extension context. For many app extensions, the most important part of the context is the data the user wants to work with, which is contained in the property.


// The host app context from which an app extension is invoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext
type ExtensionContext struct {
	objectivec.Object
}

// ExtensionContextFrom constructs a [ExtensionContext] from an unsafe.Pointer.
//
// The host app context from which an app extension is invoked.
func ExtensionContextFrom(ptr unsafe.Pointer) ExtensionContext {
	return ExtensionContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _ExtensionContextClass) Alloc() ExtensionContext {
	rv := objc.Send[ExtensionContext](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _ExtensionContextClass) New() ExtensionContext {
	rv := objc.Send[ExtensionContext](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExtensionContext) Init() ExtensionContext {
	rv := objc.Send[ExtensionContext](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExtensionContext) Autorelease() ExtensionContext {
	rv := objc.Send[ExtensionContext](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExtensionContext creates a new ExtensionContext instance.
func NewExtensionContext() ExtensionContext {
	return getExtensionContextClass().New()
}



// The maximum size for a Siri hosted view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/hostedviewmaximumallowedsize
func (e_ ExtensionContext) HostedViewMaximumAllowedSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("hostedViewMaximumAllowedSize"))
	return rv
}


// The maximum size for a Siri hosted view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/hostedviewmaximumallowedsize
func (e_ ExtensionContext) SetHostedViewMaximumAllowedSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setHostedViewMaximumAllowedSize:"), value)
}


// The active display mode of the widget.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/widgetactivedisplaymode
func (e_ ExtensionContext) WidgetActiveDisplayMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("widgetActiveDisplayMode"))
	return rv
}


// The active display mode of the widget.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/widgetactivedisplaymode
func (e_ ExtensionContext) SetWidgetActiveDisplayMode(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setWidgetActiveDisplayMode:"), value)
}


// The largest display mode the widget supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/widgetlargestavailabledisplaymode
func (e_ ExtensionContext) WidgetLargestAvailableDisplayMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("widgetLargestAvailableDisplayMode"))
	return rv
}


// The largest display mode the widget supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/widgetlargestavailabledisplaymode
func (e_ ExtensionContext) SetWidgetLargestAvailableDisplayMode(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setWidgetLargestAvailableDisplayMode:"), value)
}


// The extension items and errors key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitemsanderrorskey
func (e_ ExtensionContext) NSExtensionItemsAndErrorsKey() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("NSExtensionItemsAndErrorsKey"))
	return rv
}


