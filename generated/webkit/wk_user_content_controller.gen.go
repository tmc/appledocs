// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKUserContentController */


/* debug [class_header]: Header for WKUserContentController */
// The class instance for the [UserContentController] class.
var (
	UserContentControllerClass     _UserContentControllerClass
	UserContentControllerClassOnce sync.Once
)

func getUserContentControllerClass() _UserContentControllerClass {
	UserContentControllerClassOnce.Do(func() {
		UserContentControllerClass = _UserContentControllerClass{objc.GetClass("WKUserContentController")}
	})
	return UserContentControllerClass
}

type _UserContentControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UserContentController */
// An interface definition for the [UserContentController] class.
type IUserContentController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UserContentController */
	// properties:
	UserScripts() []UserScript
	UserContentController() IWKUserContentController
	SetUserContentController(value IWKUserContentController)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UserContentController */
	// methods:
	AddContentRuleList(contentRuleList IWKContentRuleList)
	AddScriptMessageHandlerContentWorldName(scriptMessageHandler unsafe.Pointer, world IWKContentWorld, name objc.IObject /* cross-framework: NSString */)
	AddScriptMessageHandlerName(scriptMessageHandler unsafe.Pointer, name objc.IObject /* cross-framework: NSString */)
	AddScriptMessageHandlerWithReplyContentWorldName(scriptMessageHandlerWithReply unsafe.Pointer, contentWorld IWKContentWorld, name objc.IObject /* cross-framework: NSString */)
	AddUserScript(userScript IWKUserScript)
	RemoveContentRuleList(contentRuleList IWKContentRuleList)
	RemoveAllContentRuleLists()
	RemoveAllScriptMessageHandlers()
	RemoveAllScriptMessageHandlersFromContentWorld(contentWorld IWKContentWorld)
	RemoveAllUserScripts()
	RemoveScriptMessageHandlerForName(name objc.IObject /* cross-framework: NSString */)
	RemoveScriptMessageHandlerForNameContentWorld(name objc.IObject /* cross-framework: NSString */, contentWorld IWKContentWorld)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UserContentController */
// Alloc allocates a new instance without initialization.
func (uc _UserContentControllerClass) Alloc() UserContentController {
	rv := objc.Send[UserContentController](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UserContentControllerClass) New() UserContentController {
	rv := objc.Send[UserContentController](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserContentController) Init() UserContentController {
	rv := objc.Send[UserContentController](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserContentController) Autorelease() UserContentController {
	rv := objc.Send[UserContentController](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserContentController creates a new UserContentController instance.
func NewUserContentController() UserContentController {
	return getUserContentControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UserContentController */
// An object for managing interactions between JavaScript code and your web view, and for filtering content in your web view.
//
// A object provides a bridge between your app and the JavaScript code running in the web view. Use this object to do the following: Inject JavaScript code into webpages running in your web view. Install custom JavaScript functions that call through to your app’s native code. Specify custom filters to prevent the webpage from loading restricted content. Create and configure a object as part of your overall web view setup. Assign the object to the property of your object before creating your web view.


// An object for managing interactions between JavaScript code and your web view, and for filtering content in your web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController
type UserContentController struct {
	objectivec.Object
}

// UserContentControllerFrom constructs a [UserContentController] from an unsafe.Pointer.
//
// An object for managing interactions between JavaScript code and your web view, and for filtering content in your web view.
func UserContentControllerFrom(ptr unsafe.Pointer) UserContentController {
	return UserContentController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UserContentController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UserContentController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UserContentController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UserContentController */

// Adds the specified content rule list to the content controller object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/add(_:)
func (u_ UserContentController) AddContentRuleList(contentRuleList IWKContentRuleList) {
	objc.Send[objc.ID](u_.ID, objc.Sel("addContentRuleList:"), contentRuleList)
}/* debug [instance_methods/method]: AddContentRuleList */


// Installs a message handler that you can call from the specified content world in your JavaScript code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/add(_:contentWorld:name:)
func (u_ UserContentController) AddScriptMessageHandlerContentWorldName(scriptMessageHandler unsafe.Pointer, world IWKContentWorld, name objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("addScriptMessageHandler:contentWorld:name:"), scriptMessageHandler, world, name)
}/* debug [instance_methods/method]: AddScriptMessageHandlerContentWorldName */


// Installs a message handler that you can call from your JavaScript code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/add(_:name:)
func (u_ UserContentController) AddScriptMessageHandlerName(scriptMessageHandler unsafe.Pointer, name objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("addScriptMessageHandler:name:"), scriptMessageHandler, name)
}/* debug [instance_methods/method]: AddScriptMessageHandlerName */


// Installs a message handler that returns a reply to your JavaScript code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/addScriptMessageHandler(_:contentWorld:name:)
func (u_ UserContentController) AddScriptMessageHandlerWithReplyContentWorldName(scriptMessageHandlerWithReply unsafe.Pointer, contentWorld IWKContentWorld, name objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("addScriptMessageHandlerWithReply:contentWorld:name:"), scriptMessageHandlerWithReply, contentWorld, name)
}/* debug [instance_methods/method]: AddScriptMessageHandlerWithReplyContentWorldName */


// Injects the specified script into the webpage’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/addUserScript(_:)
func (u_ UserContentController) AddUserScript(userScript IWKUserScript) {
	objc.Send[objc.ID](u_.ID, objc.Sel("addUserScript:"), userScript)
}/* debug [instance_methods/method]: AddUserScript */


// Removes the specified rule list from the content controller object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/remove(_:)
func (u_ UserContentController) RemoveContentRuleList(contentRuleList IWKContentRuleList) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeContentRuleList:"), contentRuleList)
}/* debug [instance_methods/method]: RemoveContentRuleList */


// Removes all rules lists from the content controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/removeAllContentRuleLists()
func (u_ UserContentController) RemoveAllContentRuleLists() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllContentRuleLists"))
}/* debug [instance_methods/method]: RemoveAllContentRuleLists */


// Uninstalls all custom message handlers associated with the user content controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/removeAllScriptMessageHandlers()
func (u_ UserContentController) RemoveAllScriptMessageHandlers() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllScriptMessageHandlers"))
}/* debug [instance_methods/method]: RemoveAllScriptMessageHandlers */


// Uninstalls all custom message handlers from the specified content world in your JavaScript code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/removeAllScriptMessageHandlers(from:)
func (u_ UserContentController) RemoveAllScriptMessageHandlersFromContentWorld(contentWorld IWKContentWorld) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllScriptMessageHandlersFromContentWorld:"), contentWorld)
}/* debug [instance_methods/method]: RemoveAllScriptMessageHandlersFromContentWorld */


// Removes all user scripts from the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/removeAllUserScripts()
func (u_ UserContentController) RemoveAllUserScripts() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllUserScripts"))
}/* debug [instance_methods/method]: RemoveAllUserScripts */


// Uninstalls the custom message handler with the specified name from your JavaScript code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/removeScriptMessageHandler(forName:)
func (u_ UserContentController) RemoveScriptMessageHandlerForName(name objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeScriptMessageHandlerForName:"), name)
}/* debug [instance_methods/method]: RemoveScriptMessageHandlerForName */


// Uninstalls a custom message handler from the specified content world in your JavaScript code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/removeScriptMessageHandler(forName:contentWorld:)
func (u_ UserContentController) RemoveScriptMessageHandlerForNameContentWorld(name objc.IObject /* cross-framework: NSString */, contentWorld IWKContentWorld) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeScriptMessageHandlerForName:contentWorld:"), name, contentWorld)
}/* debug [instance_methods/method]: RemoveScriptMessageHandlerForNameContentWorld */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UserContentController */

// The user scripts associated with the user content controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/userScripts
func (u_ UserContentController) UserScripts() []UserScript {
	rv := objc.Send[[]UserScript](u_.ID, objc.Sel("userScripts"))
	return rv
}/* debug [instance_properties/getter]: userScripts */


// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/usercontentcontroller
func (u_ UserContentController) UserContentController() IWKUserContentController {
	rv := objc.Send[UserContentController](u_.ID, objc.Sel("userContentController"))
	return rv
}/* debug [instance_properties/getter]: userContentController */


// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/usercontentcontroller
func (u_ UserContentController) SetUserContentController(value IWKUserContentController) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUserContentController:"), value)
}/* debug [instance_properties/setter]: userContentController */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKUserContentController */



