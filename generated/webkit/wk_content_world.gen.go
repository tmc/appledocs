// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [ContentWorld] class.
type IContentWorld interface {
	objectivec.IObject
}

// An object that defines a scope of execution for JavaScript code, and which you use to prevent conflicts between different scripts.
//
// Use a object as a namespace to separate your app’s web environment from the environment of individual webpages or scripts you execute. Content worlds help prevent issues that occur when two scripts modify environment variables in conflicting ways. Executing a script in its own content world effectively gives it a separate copy of the environment variables to modify. You might use this support in the following scenarios: You have complex script logic to bridge your web content to your app, but your web content has complex script libraries of its own. In that scenario, use one content world for your app-specific scripts and a separate content world for your content-specific scripts. You implement a web browser that supports JavaScript extensions. In that scenario, create a unique content world for each extension to prevent conflicts between the extensions. A object is a namespace and doesn’t persist data outside of the current web view or webpage. If you use the same content world in two objects, variables in one web view’s content world don’t appear in the other web view. Similarly, when the user or your app navigates to a new webpage, variables from the previous page are gone, even if both pages share the same content world. Use the methods and properties of this class to fetch the content world you need. provides a default content world for your app and a content world for the current web page. You can also create new content worlds. For example, you might create a custom content world for each JavaScript extension you manage. Specify the content world object when configuring or executing scripts associated with your content.
//
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

// Alloc allocates a new instance without initialization.
func (cc _ContentWorldClass) Alloc() ContentWorld {
	rv := objc.Send[ContentWorld](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The content world for the current webpage’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentWorld/page
func (cc _ContentWorldClass) PageWorld() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("pageWorld"))
	return rv
}
// The content world for the current webpage’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKContentWorld/page
func (c_ ContentWorld) PageWorld() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("pageWorld"))
	return rv
}



