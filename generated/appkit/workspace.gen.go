// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Workspace] class.
var (
	workspaceClass     _WorkspaceClass
	workspaceClassOnce sync.Once
)

func getWorkspaceClass() _WorkspaceClass {
	workspaceClassOnce.Do(func() {
		workspaceClass = _WorkspaceClass{objc.GetClass("NSWorkspace")}
	})
	return workspaceClass
}

type _WorkspaceClass struct {
	class objc.Class
}

// An interface definition for the [Workspace] class.
type IWorkspace interface {
	objectivec.IObject
	LaunchApplicationAtURLOptionsConfigurationError(url unsafe.Pointer, options unsafe.Pointer, configuration unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	OpenURL(url unsafe.Pointer) bool
	OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls unsafe.Pointer, applicationURL unsafe.Pointer, configuration unsafe.Pointer, completionHandler unsafe.Pointer)
}

// A workspace that can launch other apps and perform a variety of file-handling services.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace
type Workspace struct {
	objectivec.Object
}

// WorkspaceFrom constructs a [Workspace] from an unsafe.Pointer.
//
// A workspace that can launch other apps and perform a variety of file-handling services.
func WorkspaceFrom(ptr unsafe.Pointer) Workspace {
	return Workspace{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WorkspaceClass) Alloc() Workspace {
	rv := objc.Send[Workspace](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WorkspaceClass) New() Workspace {
	rv := objc.Send[Workspace](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ Workspace) Init() Workspace {
	rv := objc.Send[Workspace](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ Workspace) Autorelease() Workspace {
	rv := objc.Send[Workspace](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWorkspace creates a new Workspace instance.
func NewWorkspace() Workspace {
	return getWorkspaceClass().New()
}


// Launches the app at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/launchApplication(at:options:configuration:)
func (w_ Workspace) LaunchApplicationAtURLOptionsConfigurationError(url unsafe.Pointer, options unsafe.Pointer, configuration unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("launchApplicationAtURL:options:configuration:error:"), url, options, configuration, error)
	return rv
}

// Opens the location at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:)
func (w_ Workspace) OpenURL(url unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("openURL:"), url)
	return rv
}

// Opens one or more URLs asynchronously in the specified app using the provided options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:withApplicationAt:configuration:completionHandler:)
func (w_ Workspace) OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls unsafe.Pointer, applicationURL unsafe.Pointer, configuration unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("openURLs:withApplicationAtURL:configuration:completionHandler:"), urls, applicationURL, configuration, completionHandler)
}



