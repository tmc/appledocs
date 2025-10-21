// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CSIndexExtensionRequestHandler] class.
var (
	CSIndexExtensionRequestHandlerClass     _CSIndexExtensionRequestHandlerClass
	CSIndexExtensionRequestHandlerClassOnce sync.Once
)

func getCSIndexExtensionRequestHandlerClass() _CSIndexExtensionRequestHandlerClass {
	CSIndexExtensionRequestHandlerClassOnce.Do(func() {
		CSIndexExtensionRequestHandlerClass = _CSIndexExtensionRequestHandlerClass{objc.GetClass("CSIndexExtensionRequestHandler")}
	})
	return CSIndexExtensionRequestHandlerClass
}

type _CSIndexExtensionRequestHandlerClass struct {
	class objc.Class
}

// An interface definition for the [CSIndexExtensionRequestHandler] class.
type ICSIndexExtensionRequestHandler interface {
	objectivec.IObject
}

// An interface that implements an index-maintenance app extension.
//
// The class provides the main entry point for an index-maintenance app extension. If any issues arise with your app’s indexes and your app isn’t running, the system loads your app extension and looks for an implementation of this class. It instantiates the class it finds and uses it to perform any index-related maintenance. Define a custom subclass of in your app extension and implement methods of the protocol in it. Use those methods to perform any required updates to your app’s index files. For example, use the method to reindex all items in your app.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSIndexExtensionRequestHandler
type CSIndexExtensionRequestHandler struct {
	objectivec.Object
}

// CSIndexExtensionRequestHandlerFrom constructs a [CSIndexExtensionRequestHandler] from an unsafe.Pointer.
//
// An interface that implements an index-maintenance app extension.
func CSIndexExtensionRequestHandlerFrom(ptr unsafe.Pointer) CSIndexExtensionRequestHandler {
	return CSIndexExtensionRequestHandler{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CSIndexExtensionRequestHandlerClass) Alloc() CSIndexExtensionRequestHandler {
	rv := objc.Send[CSIndexExtensionRequestHandler](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CSIndexExtensionRequestHandlerClass) New() CSIndexExtensionRequestHandler {
	rv := objc.Send[CSIndexExtensionRequestHandler](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSIndexExtensionRequestHandler) Init() CSIndexExtensionRequestHandler {
	rv := objc.Send[CSIndexExtensionRequestHandler](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSIndexExtensionRequestHandler) Autorelease() CSIndexExtensionRequestHandler {
	rv := objc.Send[CSIndexExtensionRequestHandler](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSIndexExtensionRequestHandler creates a new CSIndexExtensionRequestHandler instance.
func NewCSIndexExtensionRequestHandler() CSIndexExtensionRequestHandler {
	return getCSIndexExtensionRequestHandlerClass().New()
}




