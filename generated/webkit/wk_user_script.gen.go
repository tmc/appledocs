// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UserScript] class.
var (
	UserScriptClass     _UserScriptClass
	UserScriptClassOnce sync.Once
)

func getUserScriptClass() _UserScriptClass {
	UserScriptClassOnce.Do(func() {
		UserScriptClass = _UserScriptClass{objc.GetClass("WKUserScript")}
	})
	return UserScriptClass
}

type _UserScriptClass struct {
	class objc.Class
}

// An interface definition for the [UserScript] class.
type IUserScript interface {
	objectivec.IObject
}

// A script that the web view injects into a webpage.
//
// Create a object when you want to inject custom script code into the pages of your web view. Use this object to specify the JavaScript code to inject, and parameters relating to when and how to inject that code. Before you create the web view, add this object to the object associated with your web view’s configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScript
type UserScript struct {
	objectivec.Object
}

// UserScriptFrom constructs a [UserScript] from an unsafe.Pointer.
//
// A script that the web view injects into a webpage.
func UserScriptFrom(ptr unsafe.Pointer) UserScript {
	return UserScript{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UserScriptClass) Alloc() UserScript {
	rv := objc.Send[UserScript](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UserScriptClass) New() UserScript {
	rv := objc.Send[UserScript](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserScript) Init() UserScript {
	rv := objc.Send[UserScript](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserScript) Autorelease() UserScript {
	rv := objc.Send[UserScript](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserScript creates a new UserScript instance.
func NewUserScript() UserScript {
	return getUserScriptClass().New()
}




// Creates a user script object that contains the specified source code and attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScript/init(source:injectionTime:forMainFrameOnly:)
func NewUserScriptWithSourceInjectionTimeForMainFrameOnly(source string, injectionTime unsafe.Pointer, forMainFrameOnly bool) UserScript {
	instance := getUserScriptClass().Alloc()
	rv := objc.Send[UserScript](instance.ID, objc.Sel("initWithSource:injectionTime:forMainFrameOnly:"), objc.String(source), injectionTime, forMainFrameOnly)
	rv.Autorelease()
	return rv
}



// Creates a user script object that is scoped to a particular content world.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScript/init(source:injectionTime:forMainFrameOnly:in:)
func NewUserScriptWithSourceInjectionTimeForMainFrameOnlyInContentWorld(source string, injectionTime unsafe.Pointer, forMainFrameOnly bool, contentWorld unsafe.Pointer) UserScript {
	instance := getUserScriptClass().Alloc()
	rv := objc.Send[UserScript](instance.ID, objc.Sel("initWithSource:injectionTime:forMainFrameOnly:inContentWorld:"), objc.String(source), injectionTime, forMainFrameOnly, contentWorld)
	rv.Autorelease()
	return rv
}


// The time at which to inject the script into the webpage.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScript/injectionTime
func (u_ UserScript) InjectionTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("injectionTime"))
	return rv
}

// A Boolean value that indicates whether to inject the script into the main frame or all frames.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScript/isForMainFrameOnly
func (u_ UserScript) ForMainFrameOnly() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("forMainFrameOnly"))
	return rv
}

// The script’s source code.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScript/source
func (u_ UserScript) Source() string {
	rv := objc.Send[string](u_.ID, objc.Sel("source"))
	return rv
}

// A Boolean value that indicates whether to inject the script into the main frame or all frames.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuserscript/isformainframeonly
func (u_ UserScript) IsForMainFrameOnly() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isForMainFrameOnly"))
	return rv
}


// SetIsForMainFrameOnly sets the value of the isForMainFrameOnly property.
// A Boolean value that indicates whether to inject the script into the main frame or all frames.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuserscript/isformainframeonly
func (u_ UserScript) SetIsForMainFrameOnly(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsForMainFrameOnly:"), value)
}


