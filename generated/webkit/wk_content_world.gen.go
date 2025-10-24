// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKContentWorld */


/* debug [class_header]: Header for WKContentWorld */
// The class instance for the [ContentWorld] class.
var (
	ContentWorldClass     _ContentWorldClass
	ContentWorldClassOnce sync.Once
)

func getContentWorldClass() _ContentWorldClass {
	ContentWorldClassOnce.Do(func() {
		ContentWorldClass = _ContentWorldClass{objc.GetClass("WKContentWorld")}
	})
	return ContentWorldClass
}

type _ContentWorldClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ContentWorld */
// An interface definition for the [ContentWorld] class.
type IContentWorld interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ContentWorld */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ContentWorld */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ContentWorld */
// Alloc allocates a new instance without initialization.
func (cc _ContentWorldClass) Alloc() ContentWorld {
	rv := objc.Send[ContentWorld](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContentWorldClass) New() ContentWorld {
	rv := objc.Send[ContentWorld](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentWorld) Init() ContentWorld {
	rv := objc.Send[ContentWorld](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentWorld) Autorelease() ContentWorld {
	rv := objc.Send[ContentWorld](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentWorld creates a new ContentWorld instance.
func NewContentWorld() ContentWorld {
	return getContentWorldClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ContentWorld */
// An object that defines a scope of execution for JavaScript code, and which you use to prevent conflicts between different scripts.
//
// Use a object as a namespace to separate your app’s web environment from the environment of individual webpages or scripts you execute. Content worlds help prevent issues that occur when two scripts modify environment variables in conflicting ways. Executing a script in its own content world effectively gives it a separate copy of the environment variables to modify. You might use this support in the following scenarios: You have complex script logic to bridge your web content to your app, but your web content has complex script libraries of its own. In that scenario, use one content world for your app-specific scripts and a separate content world for your content-specific scripts. You implement a web browser that supports JavaScript extensions. In that scenario, create a unique content world for each extension to prevent conflicts between the extensions. A object is a namespace and doesn’t persist data outside of the current web view or webpage. If you use the same content world in two objects, variables in one web view’s content world don’t appear in the other web view. Similarly, when the user or your app navigates to a new webpage, variables from the previous page are gone, even if both pages share the same content world. Use the methods and properties of this class to fetch the content world you need. provides a default content world for your app and a content world for the current web page. You can also create new content worlds. For example, you might create a custom content world for each JavaScript extension you manage. Specify the content world object when configuring or executing scripts associated with your content.


// An object that defines a scope of execution for JavaScript code, and which you use to prevent conflicts between different scripts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentWorld
type ContentWorld struct {
	objectivec.Object
}

// ContentWorldFrom constructs a [ContentWorld] from an unsafe.Pointer.
//
// An object that defines a scope of execution for JavaScript code, and which you use to prevent conflicts between different scripts.
func ContentWorldFrom(ptr unsafe.Pointer) ContentWorld {
	return ContentWorld{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ContentWorld *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ContentWorld */

// Returns the custom content world with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentWorld/world(name:)
func (cc _ContentWorldClass) WorldWithName(name objc.IObject /* cross-framework: NSString */) IContentWorld {
	rv := objc.Send[ContentWorld](objc.ID(cc.class), objc.Sel("worldWithName:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WorldWithName) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ContentWorld */

// The default world for clients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentWorld/defaultClient
func (cc _ContentWorldClass) DefaultClientWorld() ContentWorld {
	rv := objc.Send[ContentWorld](objc.ID(cc.class), objc.Sel("defaultClientWorld"))
	return rv
}/* debug [class_properties_class/property]: defaultClientWorld */

// The content world for the current webpage’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentWorld/page
func (cc _ContentWorldClass) PageWorld() ContentWorld {
	rv := objc.Send[ContentWorld](objc.ID(cc.class), objc.Sel("pageWorld"))
	return rv
}/* debug [class_properties_class/property]: pageWorld */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ContentWorld */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ContentWorld */

// The default world for clients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentWorld/defaultClient
func (c_ ContentWorld) DefaultClientWorld() IWKContentWorld {
	rv := objc.Send[ContentWorld](c_.ID, objc.Sel("defaultClientWorld"))
	return rv
}/* debug [instance_properties/getter]: defaultClientWorld */


// The name of a custom content world.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentWorld/name
func (c_ ContentWorld) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The content world for the current webpage’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentWorld/page
func (c_ ContentWorld) PageWorld() IWKContentWorld {
	rv := objc.Send[ContentWorld](c_.ID, objc.Sel("pageWorld"))
	return rv
}/* debug [instance_properties/getter]: pageWorld */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKContentWorld */



