// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [SFSafariTab] class.
type ISFSafariTab interface {
	objectivec.IObject
	// properties:
	// methods:
	ActivateWithCompletionHandler(completionHandler unsafe.Pointer)
	Close()
	GetActivePageWithCompletionHandler(completionHandler unsafe.Pointer)
	GetContainingWindowWithCompletionHandler(completionHandler unsafe.Pointer)
	GetPagesWithCompletionHandler(completionHandler unsafe.Pointer)
	NavigateToURL(url objc.IObject /* cross-framework: NSURL */)
}

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

// Alloc allocates a new instance without initialization.
func (sc _SFSafariTabClass) Alloc() SFSafariTab {
	rv := objc.Send[SFSafariTab](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Activates the tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariTab/activate(completionHandler:)
func (s_ SFSafariTab) ActivateWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("activateWithCompletionHandler:"), completionHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariTab/close()
func (s_ SFSafariTab) Close() {
	objc.Send[objc.ID](s_.ID, objc.Sel("close"))
}


// Calls the completion handler passing the active page in the tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariTab/getActivePage(completionHandler:)
func (s_ SFSafariTab) GetActivePageWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getActivePageWithCompletionHandler:"), completionHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariTab/getContainingWindow(completionHandler:)
func (s_ SFSafariTab) GetContainingWindowWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getContainingWindowWithCompletionHandler:"), completionHandler)
}


// Calls the completion handler with all of the tab’s active and preloading pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariTab/getPagesWithCompletionHandler(_:)
func (s_ SFSafariTab) GetPagesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getPagesWithCompletionHandler:"), completionHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariTab/navigate(to:)
func (s_ SFSafariTab) NavigateToURL(url objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("navigateToURL:"), url)
}



