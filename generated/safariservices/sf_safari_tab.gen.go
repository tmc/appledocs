// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariTab */


/* debug [class_header]: Header for SFSafariTab */
// The class instance for the [SFSafariTab] class.
var (
	SFSafariTabClass     _SFSafariTabClass
	SFSafariTabClassOnce sync.Once
)

func getSFSafariTabClass() _SFSafariTabClass {
	SFSafariTabClassOnce.Do(func() {
		SFSafariTabClass = _SFSafariTabClass{objc.GetClass("SFSafariTab")}
	})
	return SFSafariTabClass
}

type _SFSafariTabClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariTab */
// An interface definition for the [SFSafariTab] class.
type ISFSafariTab interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariTab */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariTab */
	// methods:
	ActivateWithCompletionHandler(completionHandler unsafe.Pointer)
	Close()
	GetActivePageWithCompletionHandler(completionHandler unsafe.Pointer)
	GetContainingWindowWithCompletionHandler(completionHandler unsafe.Pointer)
	GetPagesWithCompletionHandler(completionHandler unsafe.Pointer)
	NavigateToURL(url objc.IObject /* cross-framework: NSURL */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariTab */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariTabClass) Alloc() SFSafariTab {
	rv := objc.Send[SFSafariTab](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariTabClass) New() SFSafariTab {
	rv := objc.Send[SFSafariTab](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariTab) Init() SFSafariTab {
	rv := objc.Send[SFSafariTab](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariTab) Autorelease() SFSafariTab {
	rv := objc.Send[SFSafariTab](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariTab creates a new SFSafariTab instance.
func NewSFSafariTab() SFSafariTab {
	return getSFSafariTabClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariTab */
// A proxy for a tab in a Safari window.


// A proxy for a tab in a Safari window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariTab
type SFSafariTab struct {
	objectivec.Object
}

// SFSafariTabFrom constructs a [SFSafariTab] from an unsafe.Pointer.
//
// A proxy for a tab in a Safari window.
func SFSafariTabFrom(ptr unsafe.Pointer) SFSafariTab {
	return SFSafariTab{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariTab *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariTab */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariTab */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariTab */

// Activates the tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariTab/activate(completionHandler:)
func (s_ SFSafariTab) ActivateWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("activateWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: ActivateWithCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariTab/close()
func (s_ SFSafariTab) Close() {
	objc.Send[objc.ID](s_.ID, objc.Sel("close"))
}/* debug [instance_methods/method]: Close */


// Calls the completion handler passing the active page in the tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariTab/getActivePage(completionHandler:)
func (s_ SFSafariTab) GetActivePageWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getActivePageWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetActivePageWithCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariTab/getContainingWindow(completionHandler:)
func (s_ SFSafariTab) GetContainingWindowWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getContainingWindowWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetContainingWindowWithCompletionHandler */


// Calls the completion handler with all of the tab’s active and preloading pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariTab/getPagesWithCompletionHandler(_:)
func (s_ SFSafariTab) GetPagesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getPagesWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: GetPagesWithCompletionHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariTab/navigate(to:)
func (s_ SFSafariTab) NavigateToURL(url objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("navigateToURL:"), url)
}/* debug [instance_methods/method]: NavigateToURL */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariTab */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariTab */



