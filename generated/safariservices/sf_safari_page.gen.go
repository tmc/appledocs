// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariPage */


/* debug [class_header]: Header for SFSafariPage */
// The class instance for the [SFSafariPage] class.
var (
	SFSafariPageClass     _SFSafariPageClass
	SFSafariPageClassOnce sync.Once
)

func getSFSafariPageClass() _SFSafariPageClass {
	SFSafariPageClassOnce.Do(func() {
		SFSafariPageClass = _SFSafariPageClass{objc.GetClass("SFSafariPage")}
	})
	return SFSafariPageClass
}

type _SFSafariPageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariPage */
// An interface definition for the [SFSafariPage] class.
type ISFSafariPage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariPage */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariPage */
	// methods:
	DispatchMessageToScriptWithNameUserInfo(messageName objc.IObject /* cross-framework: NSString */, userInfo foundation.IDictionary)
	GetContainingTabWithCompletionHandler(completionHandler unsafe.Pointer)
	GetPagePropertiesWithCompletionHandler(completionHandler unsafe.Pointer)
	GetScreenshotOfVisibleAreaWithCompletionHandler(completionHandler unsafe.Pointer)
	Reload()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariPage */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariPageClass) Alloc() SFSafariPage {
	rv := objc.Send[SFSafariPage](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariPageClass) New() SFSafariPage {
	rv := objc.Send[SFSafariPage](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariPage) Init() SFSafariPage {
	rv := objc.Send[SFSafariPage](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariPage) Autorelease() SFSafariPage {
	rv := objc.Send[SFSafariPage](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariPage creates a new SFSafariPage instance.
func NewSFSafariPage() SFSafariPage {
	return getSFSafariPageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariPage */
// A proxy for a Safari webpage.
//
// Use an object in your Safari app extension to send messages to injected content scripts, access page properties, and reload the page.


// A proxy for a Safari webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPage
type SFSafariPage struct {
	objectivec.Object
}

// SFSafariPageFrom constructs a [SFSafariPage] from an unsafe.Pointer.
//
// A proxy for a Safari webpage.
func SFSafariPageFrom(ptr unsafe.Pointer) SFSafariPage {
	return SFSafariPage{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariPage *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariPage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariPage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariPage */

// Dispatches a message from the app extension to the content script injected in this page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPage/dispatchMessageToScript(withName:userInfo:)
func (s_ SFSafariPage) DispatchMessageToScriptWithNameUserInfo(messageName objc.IObject /* cross-framework: NSString */, userInfo foundation.IDictionary) {
	objc.Send[objc.ID](s_.ID, objc.Sel("dispatchMessageToScriptWithName:userInfo:"), messageName, userInfo)
}/* debug [instance_methods/method]: DispatchMessageToScriptWithNameUserInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPage/getContainingTab(completionHandler:)
func (s_ SFSafariPage) GetContainingTabWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getContainingTabWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetContainingTabWithCompletionHandler */


// Retrieves the properties of the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPage/getPropertiesWithCompletionHandler(_:)
func (s_ SFSafariPage) GetPagePropertiesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getPagePropertiesWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetPagePropertiesWithCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPage/getScreenshotOfVisibleArea(completionHandler:)
func (s_ SFSafariPage) GetScreenshotOfVisibleAreaWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getScreenshotOfVisibleAreaWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetScreenshotOfVisibleAreaWithCompletionHandler */


// Tells Safari to reload the webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPage/reload()
func (s_ SFSafariPage) Reload() {
	objc.Send[objc.ID](s_.ID, objc.Sel("reload"))
}/* debug [instance_methods/method]: Reload */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariPage */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariPage */



