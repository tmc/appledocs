// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [SFSafariWindow] class.
type ISFSafariWindow interface {
	objectivec.IObject
	Close()
	GetActiveTabWithCompletionHandler(completionHandler unsafe.Pointer)
	GetAllTabsWithCompletionHandler(completionHandler unsafe.Pointer)
	GetToolbarItemWithCompletionHandler(completionHandler unsafe.Pointer)
	OpenTabWithURLMakeActiveIfPossibleCompletionHandler(url unsafe.Pointer, activateTab bool, completionHandler unsafe.Pointer)
}

// A proxy for a Safari window.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SFSafariWindowClass) Alloc() SFSafariWindow {
	rv := objc.Send[SFSafariWindow](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariWindow/close()
func (s_ SFSafariWindow) Close() {
	objc.Send[objc.ID](s_.ID, objc.Sel("close"))
}

// Calls the completion handler with the active tab in the target window.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariWindow/getActiveTab(completionHandler:)
func (s_ SFSafariWindow) GetActiveTabWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getActiveTabWithCompletionHandler:"), completionHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariWindow/getAllTabs(completionHandler:)
func (s_ SFSafariWindow) GetAllTabsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getAllTabsWithCompletionHandler:"), completionHandler)
}

// Gets the extension’s toolbar item from the target window.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariWindow/getToolbarItem(completionHandler:)
func (s_ SFSafariWindow) GetToolbarItemWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getToolbarItemWithCompletionHandler:"), completionHandler)
}

// Opens a tab at the end of the tab bar.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariWindow/openTab(with:makeActiveIfPossible:completionHandler:)
func (s_ SFSafariWindow) OpenTabWithURLMakeActiveIfPossibleCompletionHandler(url unsafe.Pointer, activateTab bool, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("openTabWithURL:makeActiveIfPossible:completionHandler:"), url, activateTab, completionHandler)
}



