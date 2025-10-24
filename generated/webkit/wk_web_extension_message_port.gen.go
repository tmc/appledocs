// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebExtensionMessagePort */


/* debug [class_header]: Header for WKWebExtensionMessagePort */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebExtensionMessagePort */
// An interface definition for the [WebExtensionMessagePort] class.
type IWebExtensionMessagePort interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebExtensionMessagePort */
	// properties:
	ApplicationIdentifier() objc.IObject /* cross-framework: NSString */
	DisconnectHandler() unsafe.Pointer
	SetDisconnectHandler(value unsafe.Pointer)
	Disconnected() bool
	MessageHandler() unsafe.Pointer
	SetMessageHandler(value unsafe.Pointer)
	IsDisconnected() bool
	SetIsDisconnected(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebExtensionMessagePort */
	// methods:
	Disconnect()
	DisconnectWithError(error_ objc.IObject /* cross-framework: Error */)
	SendMessageCompletionHandler(message objc.IObject, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebExtensionMessagePort */
// Alloc allocates a new instance without initialization.
func (wc _WebExtensionMessagePortClass) Alloc() WebExtensionMessagePort {
	rv := objc.Send[WebExtensionMessagePort](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebExtensionMessagePort */
// An object that manages message-based communication with a web extension.
//
// Contains properties and methods to handle message exchanges with a web extension.


// An object that manages message-based communication with a web extension.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebExtensionMessagePort *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebExtensionMessagePort */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebExtensionMessagePort */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebExtensionMessagePort */

// Disconnects the port, terminating all further messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/disconnect()
func (w_ WebExtensionMessagePort) Disconnect() {
	objc.Send[objc.ID](w_.ID, objc.Sel("disconnect"))
}/* debug [instance_methods/method]: Disconnect */


// Disconnects the port, terminating all further messages with an optional error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/disconnect(throwing:)
func (w_ WebExtensionMessagePort) DisconnectWithError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("disconnectWithError:"), error_)
}/* debug [instance_methods/method]: DisconnectWithError */


// Sends a message to the connected web extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/sendMessage(_:completionHandler:)
func (w_ WebExtensionMessagePort) SendMessageCompletionHandler(message objc.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("sendMessage:completionHandler:"), message, completionHandler)
}/* debug [instance_methods/method]: SendMessageCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebExtensionMessagePort */

// The unique identifier for the app to which this port should be connected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/applicationIdentifier
func (w_ WebExtensionMessagePort) ApplicationIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("applicationIdentifier"))
	return rv
}/* debug [instance_properties/getter]: applicationIdentifier */


// The block to be executed when the port disconnects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/disconnectHandler
func (w_ WebExtensionMessagePort) DisconnectHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("disconnectHandler"))
	return rv
}/* debug [instance_properties/getter]: disconnectHandler */


// The block to be executed when the port disconnects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/disconnectHandler
func (w_ WebExtensionMessagePort) SetDisconnectHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisconnectHandler:"), value)
}/* debug [instance_properties/setter]: disconnectHandler */


// Indicates whether the message port is disconnected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/isDisconnected
func (w_ WebExtensionMessagePort) Disconnected() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("disconnected"))
	return rv
}/* debug [instance_properties/getter]: disconnected */


// The block to be executed when a message is received from the web extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/messageHandler
func (w_ WebExtensionMessagePort) MessageHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("messageHandler"))
	return rv
}/* debug [instance_properties/getter]: messageHandler */


// The block to be executed when a message is received from the web extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/MessagePort/messageHandler
func (w_ WebExtensionMessagePort) SetMessageHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMessageHandler:"), value)
}/* debug [instance_properties/setter]: messageHandler */


// Indicates whether the message port is disconnected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/messageport/isdisconnected
func (w_ WebExtensionMessagePort) IsDisconnected() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isDisconnected"))
	return rv
}/* debug [instance_properties/getter]: isDisconnected */


// Indicates whether the message port is disconnected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/messageport/isdisconnected
func (w_ WebExtensionMessagePort) SetIsDisconnected(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsDisconnected:"), value)
}/* debug [instance_properties/setter]: isDisconnected */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKWebExtensionMessagePort */



