
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Workspace] class.
var WorkspaceClass _WorkspaceClass

func init() {
	WorkspaceClass = _WorkspaceClass{objc.GetClass("NSWorkspace")}
}

type _WorkspaceClass struct {
	objc.Class
}

// An interface definition for the [Workspace] class.
type IWorkspace interface {
	ID() objc.ID
	LaunchApplicationAtURLOptionsConfigurationError(url unsafe.Pointer, options unsafe.Pointer, configuration unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	OpenURL(url unsafe.Pointer) bool
	OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls unsafe.Pointer, applicationURL unsafe.Pointer, configuration unsafe.Pointer, completionHandler unsafe.Pointer)
}

type Workspace struct {
	id objc.ID
}

func WorkspaceFrom(ptr unsafe.Pointer) Workspace {
	return Workspace{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ Workspace) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _WorkspaceClass) Alloc() Workspace {
	rv := objc.Send[Workspace](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _WorkspaceClass) New() Workspace {
	rv := objc.Send[Workspace](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewWorkspace creates and returns a new initialized instance.
func NewWorkspace() Workspace {
	return WorkspaceClass.New()
}

// Init initializes the instance.
func (w_ Workspace) Init() Workspace {
	rv := objc.Send[Workspace](w_.ID(), selInit)
	return rv
}
// Launches the app at the specified URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWorkspace/launchApplication(at:options:configuration:)
func (w_ Workspace) LaunchApplicationAtURLOptionsConfigurationError(url unsafe.Pointer, options unsafe.Pointer, configuration unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("launchApplicationAtURL:options:configuration:error:"), url, options, configuration, error)
	return rv
}
// Opens the location at the specified URL. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWorkspace/open(_:)
func (w_ Workspace) OpenURL(url unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("openURL:"), url)
	return rv
}
// Opens one or more URLs asynchronously in the specified app using the provided options. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWorkspace/open(_:withApplicationAt:configuration:completionHandler:)
func (w_ Workspace) OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls unsafe.Pointer, applicationURL unsafe.Pointer, configuration unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("openURLs:withApplicationAtURL:configuration:completionHandler:"), urls, applicationURL, configuration, completionHandler)
}
