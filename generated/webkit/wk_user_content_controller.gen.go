// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [UserContentController] class.
type IUserContentController interface {
	objectivec.IObject
	// properties:
	UserScripts() IWKUserScript
	SetUserScripts(value IWKUserScript)
	UserContentController() IWKUserContentController
	SetUserContentController(value IWKUserContentController)
	// methods:
	RemoveScriptMessageHandlerForNameContentWorld(name objc.IObject /* cross-framework: NSString */, contentWorld IWKContentWorld)
}

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

// Alloc allocates a new instance without initialization.
func (uc _UserContentControllerClass) Alloc() UserContentController {
	rv := objc.Send[UserContentController](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Uninstalls a custom message handler from the specified content world in your JavaScript code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserContentController/removeScriptMessageHandler(forName:contentWorld:)
func (u_ UserContentController) RemoveScriptMessageHandlerForNameContentWorld(name objc.IObject /* cross-framework: NSString */, contentWorld IWKContentWorld) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeScriptMessageHandlerForName:contentWorld:"), name, contentWorld)
}


// The user scripts associated with the user content controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/userscripts
func (u_ UserContentController) UserScripts() IWKUserScript {
	rv := objc.Send[UserScript](u_.ID, objc.Sel("userScripts"))
	return rv
}


// The user scripts associated with the user content controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkusercontentcontroller/userscripts
func (u_ UserContentController) SetUserScripts(value IWKUserScript) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUserScripts:"), value)
}


// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/usercontentcontroller
func (u_ UserContentController) UserContentController() IWKUserContentController {
	rv := objc.Send[UserContentController](u_.ID, objc.Sel("userContentController"))
	return rv
}


// The object that coordinates interactions between your app’s native code and the webpage’s scripts and other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/usercontentcontroller
func (u_ UserContentController) SetUserContentController(value IWKUserContentController) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUserContentController:"), value)
}



