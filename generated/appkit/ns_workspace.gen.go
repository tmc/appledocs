// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	LaunchApplicationAtURLOptionsConfigurationError(url foundation.IURL, options unsafe.Pointer, configuration unsafe.Pointer, error_ unsafe.Pointer) RunningApplication
	OpenURL(url foundation.IURL) bool
	OpenURLConfigurationCompletionHandler(url foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer)
	OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls []foundation.IURL, applicationURL foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer)
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
func (wc _WorkspaceClass) SharedWorkspace() Workspace {
	rv := objc.Send[NSWorkspace](objc.ID(wc.class), objc.Sel("sharedWorkspace"))
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
func (w_ Workspace) LaunchApplicationAtURLOptionsConfigurationError(url foundation.IURL, options unsafe.Pointer, configuration unsafe.Pointer, error_ unsafe.Pointer) RunningApplication {
	rv := objc.Send[RunningApplication](w_.ID, objc.Sel("launchApplicationAtURL:options:configuration:error:"), url, options, configuration, error_)
	return rv
}

// Opens the location at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:)
func (w_ Workspace) OpenURL(url foundation.IURL) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("openURL:"), url)
	return rv
}

// Opens a URL asynchronously using the provided options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:configuration:completionHandler:)
func (w_ Workspace) OpenURLConfigurationCompletionHandler(url foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("openURL:configuration:completionHandler:"), url, configuration, completionHandler)
}

// Opens one or more URLs asynchronously in the specified app using the provided options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:withApplicationAt:configuration:completionHandler:)
func (w_ Workspace) OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls []foundation.IURL, applicationURL foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer) {
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
func (w_ Workspace) NotificationCenter() foundation.NotificationCenter {
	rv := objc.Send[foundation.NotificationCenter](w_.ID, objc.Sel("notificationCenter"))
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
func (w_ Workspace) SharedWorkspace() NSWorkspace {
	rv := objc.Send[NSWorkspace](w_.ID, objc.Sel("sharedWorkspace"))
	return rv
}

// A Boolean value that indicates whether the app avoids conveying information through color alone.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/accessibilitydisplayshoulddifferentiatewithoutcolor
func (w_ Workspace) AccessibilityDisplayShouldDifferentiateWithoutColor() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldDifferentiateWithoutColor"))
	return rv
}


// SetAccessibilityDisplayShouldDifferentiateWithoutColor sets the value of the accessibilityDisplayShouldDifferentiateWithoutColor property.
// A Boolean value that indicates whether the app avoids conveying information through color alone.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/accessibilitydisplayshoulddifferentiatewithoutcolor
func (w_ Workspace) SetAccessibilityDisplayShouldDifferentiateWithoutColor(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAccessibilityDisplayShouldDifferentiateWithoutColor:"), value)
}

// A Boolean value that indicates whether the app presents a high-contrast user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/accessibilitydisplayshouldincreasecontrast
func (w_ Workspace) AccessibilityDisplayShouldIncreaseContrast() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldIncreaseContrast"))
	return rv
}


// SetAccessibilityDisplayShouldIncreaseContrast sets the value of the accessibilityDisplayShouldIncreaseContrast property.
// A Boolean value that indicates whether the app presents a high-contrast user interface.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/accessibilitydisplayshouldincreasecontrast
func (w_ Workspace) SetAccessibilityDisplayShouldIncreaseContrast(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAccessibilityDisplayShouldIncreaseContrast:"), value)
}

// A Boolean value that indicates whether the accessibility option to invert colors is in an enabled state.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/accessibilitydisplayshouldinvertcolors
func (w_ Workspace) AccessibilityDisplayShouldInvertColors() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldInvertColors"))
	return rv
}


// SetAccessibilityDisplayShouldInvertColors sets the value of the accessibilityDisplayShouldInvertColors property.
// A Boolean value that indicates whether the accessibility option to invert colors is in an enabled state.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/accessibilitydisplayshouldinvertcolors
func (w_ Workspace) SetAccessibilityDisplayShouldInvertColors(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAccessibilityDisplayShouldInvertColors:"), value)
}

// A Boolean value that indicates whether the accessibility option to reduce motion is in an enabled state.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/accessibilitydisplayshouldreducemotion
func (w_ Workspace) AccessibilityDisplayShouldReduceMotion() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldReduceMotion"))
	return rv
}


// SetAccessibilityDisplayShouldReduceMotion sets the value of the accessibilityDisplayShouldReduceMotion property.
// A Boolean value that indicates whether the accessibility option to reduce motion is in an enabled state.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/accessibilitydisplayshouldreducemotion
func (w_ Workspace) SetAccessibilityDisplayShouldReduceMotion(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAccessibilityDisplayShouldReduceMotion:"), value)
}

// A Boolean value that indicates whether the app avoids using semitransparent backgrounds.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/accessibilitydisplayshouldreducetransparency
func (w_ Workspace) AccessibilityDisplayShouldReduceTransparency() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldReduceTransparency"))
	return rv
}


// SetAccessibilityDisplayShouldReduceTransparency sets the value of the accessibilityDisplayShouldReduceTransparency property.
// A Boolean value that indicates whether the app avoids using semitransparent backgrounds.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/accessibilitydisplayshouldreducetransparency
func (w_ Workspace) SetAccessibilityDisplayShouldReduceTransparency(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAccessibilityDisplayShouldReduceTransparency:"), value)
}

// The array of colors for the file labels.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/filelabelcolors
func (w_ Workspace) FileLabelColors() NSColor {
	rv := objc.Send[NSColor](w_.ID, objc.Sel("fileLabelColors"))
	return rv
}


// SetFileLabelColors sets the value of the fileLabelColors property.
// The array of colors for the file labels.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/filelabelcolors
func (w_ Workspace) SetFileLabelColors(value IColor) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFileLabelColors:"), value)
}

// The array of file labels, returned as strings.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/filelabels
func (w_ Workspace) FileLabels() string {
	rv := objc.Send[string](w_.ID, objc.Sel("fileLabels"))
	return rv
}


// SetFileLabels sets the value of the fileLabels property.
// The array of file labels, returned as strings.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/filelabels
func (w_ Workspace) SetFileLabels(value string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFileLabels:"), objc.String(value))
}

// Returns the frontmost app, which is the app that receives key events.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/frontmostapplication
func (w_ Workspace) FrontmostApplication() NSRunningApplication {
	rv := objc.Send[NSRunningApplication](w_.ID, objc.Sel("frontmostApplication"))
	return rv
}


// SetFrontmostApplication sets the value of the frontmostApplication property.
// Returns the frontmost app, which is the app that receives key events.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/frontmostapplication
func (w_ Workspace) SetFrontmostApplication(value IRunningApplication) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrontmostApplication:"), value)
}

// A Boolean value that indicates whether Switch Control is currently running.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isswitchcontrolenabled
func (w_ Workspace) IsSwitchControlEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isSwitchControlEnabled"))
	return rv
}


// SetIsSwitchControlEnabled sets the value of the isSwitchControlEnabled property.
// A Boolean value that indicates whether Switch Control is currently running.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isswitchcontrolenabled
func (w_ Workspace) SetIsSwitchControlEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsSwitchControlEnabled:"), value)
}

// A Boolean value that indicates whether VoiceOver is currently running.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isvoiceoverenabled
func (w_ Workspace) IsVoiceOverEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isVoiceOverEnabled"))
	return rv
}


// SetIsVoiceOverEnabled sets the value of the isVoiceOverEnabled property.
// A Boolean value that indicates whether VoiceOver is currently running.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isvoiceoverenabled
func (w_ Workspace) SetIsVoiceOverEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsVoiceOverEnabled:"), value)
}

// Returns the app that owns the currently displayed menu bar.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/menubarowningapplication
func (w_ Workspace) MenuBarOwningApplication() NSRunningApplication {
	rv := objc.Send[NSRunningApplication](w_.ID, objc.Sel("menuBarOwningApplication"))
	return rv
}


// SetMenuBarOwningApplication sets the value of the menuBarOwningApplication property.
// Returns the app that owns the currently displayed menu bar.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/menubarowningapplication
func (w_ Workspace) SetMenuBarOwningApplication(value IRunningApplication) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMenuBarOwningApplication:"), value)
}



