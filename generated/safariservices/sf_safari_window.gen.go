// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariWindow */


/* debug [class_header]: Header for SFSafariWindow */
// The class instance for the [SFSafariWindow] class.
var (
	SFSafariWindowClass     _SFSafariWindowClass
	SFSafariWindowClassOnce sync.Once
)

func getSFSafariWindowClass() _SFSafariWindowClass {
	SFSafariWindowClassOnce.Do(func() {
		SFSafariWindowClass = _SFSafariWindowClass{objc.GetClass("SFSafariWindow")}
	})
	return SFSafariWindowClass
}

type _SFSafariWindowClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariWindow */
// An interface definition for the [SFSafariWindow] class.
type ISFSafariWindow interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariWindow */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariWindow */
	// methods:
	Close()
	GetActiveTabWithCompletionHandler(completionHandler unsafe.Pointer)
	GetAllTabsWithCompletionHandler(completionHandler unsafe.Pointer)
	GetToolbarItemWithCompletionHandler(completionHandler unsafe.Pointer)
	OpenTabWithURLMakeActiveIfPossibleCompletionHandler(url objc.IObject /* cross-framework: NSURL */, activateTab bool, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariWindow */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariWindowClass) Alloc() SFSafariWindow {
	rv := objc.Send[SFSafariWindow](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariWindowClass) New() SFSafariWindow {
	rv := objc.Send[SFSafariWindow](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariWindow) Init() SFSafariWindow {
	rv := objc.Send[SFSafariWindow](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariWindow) Autorelease() SFSafariWindow {
	rv := objc.Send[SFSafariWindow](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariWindow creates a new SFSafariWindow instance.
func NewSFSafariWindow() SFSafariWindow {
	return getSFSafariWindowClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariWindow */
// A proxy for a Safari window.


// A proxy for a Safari window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariWindow
type SFSafariWindow struct {
	objectivec.Object
}

// SFSafariWindowFrom constructs a [SFSafariWindow] from an unsafe.Pointer.
//
// A proxy for a Safari window.
func SFSafariWindowFrom(ptr unsafe.Pointer) SFSafariWindow {
	return SFSafariWindow{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariWindow *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariWindow */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariWindow */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariWindow */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariWindow/close()
func (s_ SFSafariWindow) Close() {
	objc.Send[objc.ID](s_.ID, objc.Sel("close"))
}/* debug [instance_methods/method]: Close */


// Calls the completion handler with the active tab in the target window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariWindow/getActiveTab(completionHandler:)
func (s_ SFSafariWindow) GetActiveTabWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getActiveTabWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetActiveTabWithCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariWindow/getAllTabs(completionHandler:)
func (s_ SFSafariWindow) GetAllTabsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getAllTabsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetAllTabsWithCompletionHandler */


// Gets the extension’s toolbar item from the target window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariWindow/getToolbarItem(completionHandler:)
func (s_ SFSafariWindow) GetToolbarItemWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getToolbarItemWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetToolbarItemWithCompletionHandler */


// Opens a tab at the end of the tab bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariWindow/openTab(with:makeActiveIfPossible:completionHandler:)
func (s_ SFSafariWindow) OpenTabWithURLMakeActiveIfPossibleCompletionHandler(url objc.IObject /* cross-framework: NSURL */, activateTab bool, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("openTabWithURL:makeActiveIfPossible:completionHandler:"), url, activateTab, completionHandler)
}/* debug [instance_methods/method]: OpenTabWithURLMakeActiveIfPossibleCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariWindow */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariWindow */



