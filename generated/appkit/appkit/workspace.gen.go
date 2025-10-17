// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Workspace] class.
var WorkspaceClass objc.Class

func init() {
	WorkspaceClass = objc.GetClass("NSWorkspace")
}

type Workspace struct {
	objc.ID
}

func WorkspaceFrom(ptr unsafe.Pointer) Workspace {
	return Workspace{
		ID: objc.ID(ptr),
	}
}


// Launches the app at the specified URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWorkspace/launchApplication(at:options:configuration:)
func (w_ Workspace) LaunchApplicationAtURLOptionsConfigurationError(url unsafe.Pointer, options unsafe.Pointer, configuration unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("launchApplicationAtURL:options:configuration:error:")
	ret := w_.ID.Send(sel, url, options, configuration, error)
	return unsafe.Pointer(ret)
}
// Opens the location at the specified URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWorkspace/open(_:)
func (w_ Workspace) OpenURL(url unsafe.Pointer) bool {
	sel := objc.RegisterName("openURL:")
	ret := w_.ID.Send(sel, url)
	return ret != 0
}
// Opens one or more URLs asynchronously in the specified app using the provided options. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWorkspace/open(_:withApplicationAt:configuration:completionHandler:)
func (w_ Workspace) OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls unsafe.Pointer, applicationURL unsafe.Pointer, configuration unsafe.Pointer, completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("openURLs:withApplicationAtURL:configuration:completionHandler:")
	w_.ID.Send(sel, urls, applicationURL, configuration, completionHandler)
}

