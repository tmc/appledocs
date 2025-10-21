// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Workspace] class.
var (
	WorkspaceClass     _WorkspaceClass
	WorkspaceClassOnce sync.Once
)

func getWorkspaceClass() _WorkspaceClass {
	WorkspaceClassOnce.Do(func() {
		WorkspaceClass = _WorkspaceClass{objc.GetClass("NSWorkspace")}
	})
	return WorkspaceClass
}

type _WorkspaceClass struct {
	class objc.Class
}

// An interface definition for the [Workspace] class.
type IWorkspace interface {
	objectivec.IObject
	HideOtherApplications()
	LaunchApplicationAtURLOptionsConfigurationError(url unsafe.Pointer, options unsafe.Pointer, configuration unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	OpenURL(url unsafe.Pointer) bool
	OpenURLConfigurationCompletionHandler(url unsafe.Pointer, configuration unsafe.Pointer, completionHandler unsafe.Pointer)
	OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls unsafe.Pointer, applicationURL unsafe.Pointer, configuration unsafe.Pointer, completionHandler unsafe.Pointer)
	RequestAuthorizationOfTypeCompletionHandler(type_ unsafe.Pointer, completionHandler unsafe.Pointer)
}

// A workspace that can launch other apps and perform a variety of file-handling services.
//
// There is one shared object per app. You use the class method to access it. For example, the following statement uses an object to request that a file be opened in the TextEdit app: You can use the workspace object to: Open, manipulate, and get information about files and devices. Track changes to the file system, devices, and the user database. Get and set Finder information for files. Launch apps.
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


// The shared workspace object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/shared
func (wc _WorkspaceClass) SharedWorkspace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("sharedWorkspace"))
	return rv
}
// Hides all applications other than the sender.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/hideOtherApplications()
func (w_ Workspace) HideOtherApplications() {
	objc.Send[objc.ID](w_.ID, objc.Sel("hideOtherApplications"))
}

// Launches the app at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/launchApplication(at:options:configuration:)
func (w_ Workspace) LaunchApplicationAtURLOptionsConfigurationError(url unsafe.Pointer, options unsafe.Pointer, configuration unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("launchApplicationAtURL:options:configuration:error:"), url, options, configuration, error_)
	return rv
}

// Opens the location at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:)
func (w_ Workspace) OpenURL(url unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("openURL:"), url)
	return rv
}

// Opens a URL asynchronously using the provided options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:configuration:completionHandler:)
func (w_ Workspace) OpenURLConfigurationCompletionHandler(url unsafe.Pointer, configuration unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("openURL:configuration:completionHandler:"), url, configuration, completionHandler)
}

// Opens one or more URLs asynchronously in the specified app using the provided options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:withApplicationAt:configuration:completionHandler:)
func (w_ Workspace) OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls unsafe.Pointer, applicationURL unsafe.Pointer, configuration unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("openURLs:withApplicationAtURL:configuration:completionHandler:"), urls, applicationURL, configuration, completionHandler)
}

// Requests authorization to perform a privileged file operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/requestAuthorization(to:completionHandler:)
func (w_ Workspace) RequestAuthorizationOfTypeCompletionHandler(type_ unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("requestAuthorizationOfType:completionHandler:"), type_, completionHandler)
}

// The notification center for workspace notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/notificationCenter
func (w_ Workspace) NotificationCenter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("notificationCenter"))
	return rv
}

// Returns an array of running apps.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/runningApplications
func (w_ Workspace) RunningApplications() []RunningApplication {
	rv := objc.Send[[]RunningApplication](w_.ID, objc.Sel("runningApplications"))
	return rv
}

// The shared workspace object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/shared
func (w_ Workspace) SharedWorkspace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("sharedWorkspace"))
	return rv
}



