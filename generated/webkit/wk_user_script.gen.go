// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKUserScript */

/* debug [class_header]: Header for WKUserScript */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for UserScript */
// An interface definition for the [UserScript] class.
type IUserScript interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for UserScript */
	// properties:
	InjectionTime() UserScriptInjectionTime
	ForMainFrameOnly() bool
	Source() objc.IObject /* cross-framework: NSString */
	IsForMainFrameOnly() bool
	SetIsForMainFrameOnly(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for UserScript */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for UserScript */
// Alloc allocates a new instance without initialization.
func (uc _UserScriptClass) Alloc() UserScript {
	rv := objc.Send[UserScript](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for UserScript */
// A script that the web view injects into a webpage.
//
// Create a object when you want to inject custom script code into the pages of your web view. Use this object to specify the JavaScript code to inject, and parameters relating to when and how to inject that code. Before you create the web view, add this object to the object associated with your web view’s configuration.

// A script that the web view injects into a webpage.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for UserScript */

// Creates a user script object that contains the specified source code and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScript/init(source:injectionTime:forMainFrameOnly:)
func NewUserScriptWithSourceInjectionTimeForMainFrameOnly(source objc.IObject /* cross-framework: NSString */, injectionTime UserScriptInjectionTime, forMainFrameOnly bool) UserScript {
	instance := getUserScriptClass().Alloc()
	rv := objc.Send[UserScript](instance.ID, objc.Sel("initWithSource:injectionTime:forMainFrameOnly:"), source, injectionTime, forMainFrameOnly)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewUserScriptWithSourceInjectionTimeForMainFrameOnly */

// Creates a user script object that is scoped to a particular content world.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScript/init(source:injectionTime:forMainFrameOnly:in:)
func NewUserScriptWithSourceInjectionTimeForMainFrameOnlyInContentWorld(source objc.IObject /* cross-framework: NSString */, injectionTime UserScriptInjectionTime, forMainFrameOnly bool, contentWorld IWKContentWorld) UserScript {
	instance := getUserScriptClass().Alloc()
	rv := objc.Send[UserScript](instance.ID, objc.Sel("initWithSource:injectionTime:forMainFrameOnly:inContentWorld:"), source, injectionTime, forMainFrameOnly, contentWorld)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewUserScriptWithSourceInjectionTimeForMainFrameOnlyInContentWorld */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for UserScript */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for UserScript */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for UserScript */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for UserScript */

// The time at which to inject the script into the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScript/injectionTime
func (u_ UserScript) InjectionTime() UserScriptInjectionTime {
	rv := objc.Send[UserScriptInjectionTime](u_.ID, objc.Sel("injectionTime"))
	return rv
} /* debug [instance_properties/getter]: injectionTime */

// A Boolean value that indicates whether to inject the script into the main frame or all frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScript/isForMainFrameOnly
func (u_ UserScript) ForMainFrameOnly() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("forMainFrameOnly"))
	return rv
} /* debug [instance_properties/getter]: forMainFrameOnly */

// The script’s source code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKUserScript/source
func (u_ UserScript) Source() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("source"))
	return rv
} /* debug [instance_properties/getter]: source */

// A Boolean value that indicates whether to inject the script into the main frame or all frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuserscript/isformainframeonly
func (u_ UserScript) IsForMainFrameOnly() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isForMainFrameOnly"))
	return rv
} /* debug [instance_properties/getter]: isForMainFrameOnly */

// A Boolean value that indicates whether to inject the script into the main frame or all frames.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuserscript/isformainframeonly
func (u_ UserScript) SetIsForMainFrameOnly(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsForMainFrameOnly:"), value)
} /* debug [instance_properties/setter]: isForMainFrameOnly */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKUserScript */
