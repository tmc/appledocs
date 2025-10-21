// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [SFSafariPage] class.
type ISFSafariPage interface {
	objectivec.IObject
	DispatchMessageToScriptWithNameUserInfo(messageName appkit.string, userInfo unsafe.Pointer)
	GetContainingTabWithCompletionHandler(completionHandler unsafe.Pointer)
	GetPagePropertiesWithCompletionHandler(completionHandler unsafe.Pointer)
	GetScreenshotOfVisibleAreaWithCompletionHandler(completionHandler unsafe.Pointer)
	Reload()
}

// A proxy for a Safari webpage.
//
// Use an object in your Safari app extension to send messages to injected content scripts, access page properties, and reload the page.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SFSafariPageClass) Alloc() SFSafariPage {
	rv := objc.Send[SFSafariPage](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Dispatches a message from the app extension to the content script injected in this page.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPage/dispatchMessageToScript(withName:userInfo:)
func (s_ SFSafariPage) DispatchMessageToScriptWithNameUserInfo(messageName appkit.string, userInfo unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("dispatchMessageToScriptWithName:userInfo:"), messageName, userInfo)
}

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPage/getContainingTab(completionHandler:)
func (s_ SFSafariPage) GetContainingTabWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getContainingTabWithCompletionHandler:"), completionHandler)
}

// Retrieves the properties of the webpage.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPage/getPropertiesWithCompletionHandler(_:)
func (s_ SFSafariPage) GetPagePropertiesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getPagePropertiesWithCompletionHandler:"), completionHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPage/getScreenshotOfVisibleArea(completionHandler:)
func (s_ SFSafariPage) GetScreenshotOfVisibleAreaWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("getScreenshotOfVisibleAreaWithCompletionHandler:"), completionHandler)
}

// Tells Safari to reload the webpage.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPage/reload()
func (s_ SFSafariPage) Reload() {
	objc.Send[objc.ID](s_.ID, objc.Sel("reload"))
}



