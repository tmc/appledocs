// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKScriptMessage */


/* debug [class_header]: Header for WKScriptMessage */
// The class instance for the [ScriptMessage] class.
var (
	ScriptMessageClass     _ScriptMessageClass
	ScriptMessageClassOnce sync.Once
)

func getScriptMessageClass() _ScriptMessageClass {
	ScriptMessageClassOnce.Do(func() {
		ScriptMessageClass = _ScriptMessageClass{objc.GetClass("WKScriptMessage")}
	})
	return ScriptMessageClass
}

type _ScriptMessageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScriptMessage */
// An interface definition for the [ScriptMessage] class.
type IScriptMessage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScriptMessage */
	// properties:
	Body() objc.ID
	FrameInfo() IWKFrameInfo
	Name() objc.IObject /* cross-framework: NSString */
	WebView() IWKWebView
	World() IWKContentWorld
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScriptMessage */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScriptMessage */
// Alloc allocates a new instance without initialization.
func (sc _ScriptMessageClass) Alloc() ScriptMessage {
	rv := objc.Send[ScriptMessage](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScriptMessageClass) New() ScriptMessage {
	rv := objc.Send[ScriptMessage](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptMessage) Init() ScriptMessage {
	rv := objc.Send[ScriptMessage](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptMessage) Autorelease() ScriptMessage {
	rv := objc.Send[ScriptMessage](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptMessage creates a new ScriptMessage instance.
func NewScriptMessage() ScriptMessage {
	return getScriptMessageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScriptMessage */
// An object that encapsulates a message sent by JavaScript code from a webpage.
//
// Use a object to get details about a JavaScript message sent to a custom message handler in your app. You don’t create objects directly. When JavaScript code targets one of your app’s message handlers, the object of the web view creates a object and delivers it to the message handler’s delegate method. Use the object you’re provided to process the message and provide an appropriate response. For more information about handling script messages, see the and protocols. For information about how to register message handlers, see the methods of .


// An object that encapsulates a message sent by JavaScript code from a webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKScriptMessage
type ScriptMessage struct {
	objectivec.Object
}

// ScriptMessageFrom constructs a [ScriptMessage] from an unsafe.Pointer.
//
// An object that encapsulates a message sent by JavaScript code from a webpage.
func ScriptMessageFrom(ptr unsafe.Pointer) ScriptMessage {
	return ScriptMessage{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScriptMessage *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScriptMessage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScriptMessage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScriptMessage */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScriptMessage */

// The body of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKScriptMessage/body
func (s_ ScriptMessage) Body() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("body"))
	return rv
}/* debug [instance_properties/getter]: body */


// The frame that sent the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKScriptMessage/frameInfo
func (s_ ScriptMessage) FrameInfo() IWKFrameInfo {
	rv := objc.Send[FrameInfo](s_.ID, objc.Sel("frameInfo"))
	return rv
}/* debug [instance_properties/getter]: frameInfo */


// The name of the message handler to which the message is sent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKScriptMessage/name
func (s_ ScriptMessage) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The web view that sent the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKScriptMessage/webView
func (s_ ScriptMessage) WebView() IWKWebView {
	rv := objc.Send[WebView](s_.ID, objc.Sel("webView"))
	return rv
}/* debug [instance_properties/getter]: webView */


// The namespace in which the JavaScript code executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKScriptMessage/world
func (s_ ScriptMessage) World() IWKContentWorld {
	rv := objc.Send[ContentWorld](s_.ID, objc.Sel("world"))
	return rv
}/* debug [instance_properties/getter]: world */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKScriptMessage */



