// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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


// The unique identifier for the app to which this port should be connected.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/messageport/applicationidentifier
func (w_ WebExtensionMessagePort) ApplicationIdentifier() string {
	rv := objc.Send[string](w_.ID, objc.Sel("applicationIdentifier"))
	return rv
}


// SetApplicationIdentifier sets the value of the applicationIdentifier property.
// The unique identifier for the app to which this port should be connected.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/messageport/applicationidentifier
func (w_ WebExtensionMessagePort) SetApplicationIdentifier(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setApplicationIdentifier:"), objc.String(value))
}

// The block to be executed when the port disconnects.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/messageport/disconnecthandler
func (w_ WebExtensionMessagePort) DisconnectHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("disconnectHandler"))
	return rv
}


// SetDisconnectHandler sets the value of the disconnectHandler property.
// The block to be executed when the port disconnects.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/messageport/disconnecthandler
func (w_ WebExtensionMessagePort) SetDisconnectHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisconnectHandler:"), value)
}

// Indicates whether the message port is disconnected.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/messageport/isdisconnected
func (w_ WebExtensionMessagePort) IsDisconnected() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isDisconnected"))
	return rv
}


// SetIsDisconnected sets the value of the isDisconnected property.
// Indicates whether the message port is disconnected.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/messageport/isdisconnected
func (w_ WebExtensionMessagePort) SetIsDisconnected(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsDisconnected:"), value)
}

// The block to be executed when a message is received from the web extension.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/messageport/messagehandler
func (w_ WebExtensionMessagePort) MessageHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("messageHandler"))
	return rv
}


// SetMessageHandler sets the value of the messageHandler property.
// The block to be executed when a message is received from the web extension.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/messageport/messagehandler
func (w_ WebExtensionMessagePort) SetMessageHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMessageHandler:"), value)
}



