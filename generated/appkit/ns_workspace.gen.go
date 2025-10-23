// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	// properties:
	AccessibilityDisplayShouldDifferentiateWithoutColor() bool /* primitive/slice/pointer. */
	AccessibilityDisplayShouldIncreaseContrast() bool /* primitive/slice/pointer. */
	AccessibilityDisplayShouldInvertColors() bool /* primitive/slice/pointer. */
	AccessibilityDisplayShouldReduceMotion() bool /* primitive/slice/pointer. */
	AccessibilityDisplayShouldReduceTransparency() bool /* primitive/slice/pointer. */
	FileLabelColors() []Color /* primitive/slice/pointer. */
	FileLabels() []string /* primitive/slice/pointer. */
	FrontmostApplication() IRunningApplication
	SwitchControlEnabled() bool /* primitive/slice/pointer. */
	VoiceOverEnabled() bool /* primitive/slice/pointer. */
	MenuBarOwningApplication() IRunningApplication
	NotificationCenter() NotificationCenter /* not a class type */
	RunningApplications() []RunningApplication /* primitive/slice/pointer. */
	IsSwitchControlEnabled() bool /* primitive/slice/pointer. */
	SetIsSwitchControlEnabled(value bool /* primitive/slice/pointer. */)
	IsVoiceOverEnabled() bool /* primitive/slice/pointer. */
	SetIsVoiceOverEnabled(value bool /* primitive/slice/pointer. */)
	// methods:
	ActivateFileViewerSelectingURLs(fileURLs []foundation.objc.IObject /* cross-framework URL */)
	DesktopImageOptionsForScreen(screen IScreen) foundation.IDictionary /* already interface */
	DesktopImageURLForScreen(screen IScreen) foundation.objc.IObject /* cross-framework: URL */
	DuplicateURLsCompletionHandler(URLs []foundation.objc.IObject /* cross-framework URL */, handler foundation.IDictionary /* already interface */)
	ExtendPowerOffBy(requested int /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */
	GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType(fullPath string /* primitive/slice/pointer. */, removableFlag unsafe.Pointer, writableFlag unsafe.Pointer, unmountableFlag unsafe.Pointer, description string /* primitive/slice/pointer. */, fileSystemType string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	HideOtherApplications()
	IconForContentType(contentType objectivec.IObject) IImage
	IconForFile(fullPath string /* primitive/slice/pointer. */) IImage
	IconForFiles(fullPaths []string /* primitive/slice/pointer. */) IImage
	IsFilePackageAtPath(fullPath string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	NoteFileSystemChanged(path string /* primitive/slice/pointer. */)
	OpenURL(url foundation.objc.IObject /* cross-framework URL */) bool /* primitive/slice/pointer. */
	OpenURLConfigurationCompletionHandler(url foundation.objc.IObject /* cross-framework URL */, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer)
	OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls []foundation.objc.IObject /* cross-framework URL */, applicationURL foundation.objc.IObject /* cross-framework URL */, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer)
	OpenApplicationAtURLConfigurationCompletionHandler(applicationURL foundation.objc.IObject /* cross-framework URL */, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer)
	RecycleURLsCompletionHandler(URLs []foundation.objc.IObject /* cross-framework URL */, handler foundation.IDictionary /* already interface */)
	RequestAuthorizationOfTypeCompletionHandler(type_ WorkspaceAuthorizationType, completionHandler unsafe.Pointer)
	SelectFileInFileViewerRootedAtPath(fullPath string /* primitive/slice/pointer. */, rootFullPath string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler(applicationURL foundation.objc.IObject /* cross-framework URL */, contentType objectivec.IObject, completionHandler unsafe.Pointer)
	SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler(applicationURL foundation.objc.IObject /* cross-framework URL */, url foundation.objc.IObject /* cross-framework URL */, completionHandler unsafe.Pointer)
	SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler(applicationURL foundation.objc.IObject /* cross-framework URL */, url foundation.objc.IObject /* cross-framework URL */, completionHandler unsafe.Pointer)
	SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler(applicationURL foundation.objc.IObject /* cross-framework URL */, urlScheme string /* primitive/slice/pointer. */, completionHandler unsafe.Pointer)
	SetDesktopImageURLForScreenOptionsError(url foundation.objc.IObject /* cross-framework URL */, screen IScreen, options foundation.IDictionary /* already interface */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	SetIconForFileOptions(image IImage, fullPath string /* primitive/slice/pointer. */, options WorkspaceIconCreationOptions) bool /* primitive/slice/pointer. */
	ShowSearchResultsForQueryString(queryString string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	UnmountAndEjectDeviceAtURLError(url foundation.objc.IObject /* cross-framework URL */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
	UnmountAndEjectDeviceAtPath(path string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	URLForApplicationToOpenURL(url foundation.objc.IObject /* cross-framework URL */) foundation.objc.IObject /* cross-framework: URL */
	URLForApplicationToOpenContentType(contentType objectivec.IObject) foundation.objc.IObject /* cross-framework: URL */
	URLForApplicationWithBundleIdentifier(bundleIdentifier string /* primitive/slice/pointer. */) foundation.objc.IObject /* cross-framework: URL */
	URLsForApplicationsToOpenContentType(contentType objectivec.IObject) []foundation.objc.IObject /* cross-framework: URL */
	URLsForApplicationsToOpenURL(url foundation.objc.IObject /* cross-framework URL */) []foundation.objc.IObject /* cross-framework: URL */
	URLsForApplicationsWithBundleIdentifier(bundleIdentifier string /* primitive/slice/pointer. */) []foundation.objc.IObject /* cross-framework: URL */
}

// A workspace that can launch other apps and perform a variety of file-handling services.
//
// There is one shared object per app. You use the class method to access it. For example, the following statement uses an object to request that a file be opened in the TextEdit app: You can use the workspace object to: Open, manipulate, and get information about files and devices. Track changes to the file system, devices, and the user database. Get and set Finder information for files. Launch apps.


// A workspace that can launch other apps and perform a variety of file-handling services.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/shared
func (wc _WorkspaceClass) SharedWorkspace() Workspace {
	rv := objc.Send[Workspace](objc.ID(wc.class), objc.Sel("sharedWorkspace"))
	return rv
}

// Activates the Finder, and opens one or more windows selecting the specified files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/activateFileViewerSelecting(_:)
func (w_ Workspace) ActivateFileViewerSelectingURLs(fileURLs []foundation.objc.IObject /* cross-framework URL */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("activateFileViewerSelectingURLs:"), fileURLs)
}


// Returns the desktop image options for the given screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/desktopImageOptions(for:)
func (w_ Workspace) DesktopImageOptionsForScreen(screen IScreen) foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](w_.ID, objc.Sel("desktopImageOptionsForScreen:"), screen)
	return rv
}


// Returns the URL for the desktop image for the given screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/desktopImageURL(for:)
func (w_ Workspace) DesktopImageURLForScreen(screen IScreen) foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("desktopImageURLForScreen:"), screen)
	return rv
}


// Duplicates the specified URLS asynchronously in the same manner as the Finder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/duplicate(_:completionHandler:)
func (w_ Workspace) DuplicateURLsCompletionHandler(URLs []foundation.objc.IObject /* cross-framework URL */, handler foundation.IDictionary /* already interface */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("duplicateURLs:completionHandler:"), URLs, handler)
}


// Requests the system wait for the specified amount of time before turning off the power or logging out the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/extendPowerOff(by:)
func (w_ Workspace) ExtendPowerOffBy(requested int /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](w_.ID, objc.Sel("extendPowerOffBy:"), requested)
	return rv
}


// Returns information about the file system at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/getFileSystemInfo(forPath:isRemovable:isWritable:isUnmountable:description:type:)
func (w_ Workspace) GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType(fullPath string /* primitive/slice/pointer. */, removableFlag unsafe.Pointer, writableFlag unsafe.Pointer, unmountableFlag unsafe.Pointer, description string /* primitive/slice/pointer. */, fileSystemType string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("getFileSystemInfoForPath:isRemovable:isWritable:isUnmountable:description:type:"), objc.String(fullPath), removableFlag, writableFlag, unmountableFlag, objc.String(description), objc.String(fileSystemType))
	return rv
}


// Hides all applications other than the sender.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/hideOtherApplications()
func (w_ Workspace) HideOtherApplications() {
	objc.Send[objc.ID](w_.ID, objc.Sel("hideOtherApplications"))
}


// Returns an image containing the icon for the specified content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/icon(for:)
func (w_ Workspace) IconForContentType(contentType objectivec.IObject) IImage {
	rv := objc.Send[Image](w_.ID, objc.Sel("iconForContentType:"), contentType)
	return rv
}


// Returns an image containing the icon for the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/icon(forFile:)
func (w_ Workspace) IconForFile(fullPath string /* primitive/slice/pointer. */) IImage {
	rv := objc.Send[Image](w_.ID, objc.Sel("iconForFile:"), objc.String(fullPath))
	return rv
}


// Returns an image containing the icon for the specified files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/icon(forFiles:)
func (w_ Workspace) IconForFiles(fullPaths []string /* primitive/slice/pointer. */) IImage {
	rv := objc.Send[Image](w_.ID, objc.Sel("iconForFiles:"), fullPaths)
	return rv
}


// Determines whether the specified path is a file package.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/isFilePackage(atPath:)
func (w_ Workspace) IsFilePackageAtPath(fullPath string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("isFilePackageAtPath:"), objc.String(fullPath))
	return rv
}


// Informs the workspace object that the file system changed at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/noteFileSystemChanged(_:)
func (w_ Workspace) NoteFileSystemChanged(path string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("noteFileSystemChanged:"), objc.String(path))
}


// Opens the location at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:)
func (w_ Workspace) OpenURL(url foundation.objc.IObject /* cross-framework URL */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("openURL:"), url)
	return rv
}


// Opens a URL asynchronously using the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:configuration:completionHandler:)
func (w_ Workspace) OpenURLConfigurationCompletionHandler(url foundation.objc.IObject /* cross-framework URL */, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("openURL:configuration:completionHandler:"), url, configuration, completionHandler)
}


// Opens one or more URLs asynchronously in the specified app using the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:withApplicationAt:configuration:completionHandler:)
func (w_ Workspace) OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls []foundation.objc.IObject /* cross-framework URL */, applicationURL foundation.objc.IObject /* cross-framework URL */, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("openURLs:withApplicationAtURL:configuration:completionHandler:"), urls, applicationURL, configuration, completionHandler)
}


// Launches the app at the specified URL and asynchronously reports back on the app’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/openApplication(at:configuration:completionHandler:)
func (w_ Workspace) OpenApplicationAtURLConfigurationCompletionHandler(applicationURL foundation.objc.IObject /* cross-framework URL */, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("openApplicationAtURL:configuration:completionHandler:"), applicationURL, configuration, completionHandler)
}


// Moves the specified URLs to the trash in the same manner as the Finder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/recycle(_:completionHandler:)
func (w_ Workspace) RecycleURLsCompletionHandler(URLs []foundation.objc.IObject /* cross-framework URL */, handler foundation.IDictionary /* already interface */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("recycleURLs:completionHandler:"), URLs, handler)
}


// Requests authorization to perform a privileged file operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/requestAuthorization(to:completionHandler:)
func (w_ Workspace) RequestAuthorizationOfTypeCompletionHandler(type_ WorkspaceAuthorizationType, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("requestAuthorizationOfType:completionHandler:"), type_, completionHandler)
}


// Selects the file at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/selectFile(_:inFileViewerRootedAtPath:)
func (w_ Workspace) SelectFileInFileViewerRootedAtPath(fullPath string /* primitive/slice/pointer. */, rootFullPath string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("selectFile:inFileViewerRootedAtPath:"), objc.String(fullPath), objc.String(rootFullPath))
	return rv
}


// Sets the default app to use when opening files of a specific content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpen:completion:)
func (w_ Workspace) SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler(applicationURL foundation.objc.IObject /* cross-framework URL */, contentType objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultApplicationAtURL:toOpenContentType:completionHandler:"), applicationURL, contentType, completionHandler)
}


// Sets the default app to use when opening files of a specific content type defined by a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpenContentTypeOfFileAt:completion:)
func (w_ Workspace) SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler(applicationURL foundation.objc.IObject /* cross-framework URL */, url foundation.objc.IObject /* cross-framework URL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultApplicationAtURL:toOpenContentTypeOfFileAtURL:completionHandler:"), applicationURL, url, completionHandler)
}


// Sets the default app to use when opening a specific file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpenFileAt:completion:)
func (w_ Workspace) SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler(applicationURL foundation.objc.IObject /* cross-framework URL */, url foundation.objc.IObject /* cross-framework URL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultApplicationAtURL:toOpenFileAtURL:completionHandler:"), applicationURL, url, completionHandler)
}


// Sets the default app to use when opening files of a specific scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpenURLsWithScheme:completion:)
func (w_ Workspace) SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler(applicationURL foundation.objc.IObject /* cross-framework URL */, urlScheme string /* primitive/slice/pointer. */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultApplicationAtURL:toOpenURLsWithScheme:completionHandler:"), applicationURL, objc.String(urlScheme), completionHandler)
}


// Sets the desktop image for the given screen to the image at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDesktopImageURL(_:for:options:)
func (w_ Workspace) SetDesktopImageURLForScreenOptionsError(url foundation.objc.IObject /* cross-framework URL */, screen IScreen, options foundation.IDictionary /* already interface */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("setDesktopImageURL:forScreen:options:error:"), url, screen, options, error_)
	return rv
}


// Sets the icon for the file or directory at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setIcon(_:forFile:options:)
func (w_ Workspace) SetIconForFileOptions(image IImage, fullPath string /* primitive/slice/pointer. */, options WorkspaceIconCreationOptions) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("setIcon:forFile:options:"), image, objc.String(fullPath), options)
	return rv
}


// Displays a Spotlight search results window in Finder for the specified query string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/showSearchResults(forQueryString:)
func (w_ Workspace) ShowSearchResultsForQueryString(queryString string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("showSearchResultsForQueryString:"), objc.String(queryString))
	return rv
}


// Attempts to eject the volume mounted at the given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/unmountAndEjectDevice(at:)
func (w_ Workspace) UnmountAndEjectDeviceAtURLError(url foundation.objc.IObject /* cross-framework URL */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("unmountAndEjectDeviceAtURL:error:"), url, error_)
	return rv
}


// Unmounts and ejects the device at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/unmountAndEjectDevice(atPath:)
func (w_ Workspace) UnmountAndEjectDeviceAtPath(path string /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("unmountAndEjectDeviceAtPath:"), objc.String(path))
	return rv
}


// Returns the URL to the default app to open the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlForApplication(toOpen:)-7qkzf
func (w_ Workspace) URLForApplicationToOpenURL(url foundation.objc.IObject /* cross-framework URL */) foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("URLForApplicationToOpenURL:"), url)
	return rv
}


// Returns the URL to the default app to open the specified content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlForApplication(toOpen:)-95cvp
func (w_ Workspace) URLForApplicationToOpenContentType(contentType objectivec.IObject) foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("URLForApplicationToOpenContentType:"), contentType)
	return rv
}


// Returns the URL to the default app with the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlForApplication(withBundleIdentifier:)
func (w_ Workspace) URLForApplicationWithBundleIdentifier(bundleIdentifier string /* primitive/slice/pointer. */) foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("URLForApplicationWithBundleIdentifier:"), objc.String(bundleIdentifier))
	return rv
}


// Returns an array of URLs to all available applications that can open the specified content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlsForApplications(toOpen:)-60rkm
func (w_ Workspace) URLsForApplicationsToOpenContentType(contentType objectivec.IObject) []foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[[]foundation.URL](w_.ID, objc.Sel("URLsForApplicationsToOpenContentType:"), contentType)
	return rv
}


// Returns an array of URLs to all available applications that can open the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlsForApplications(toOpen:)-ualk
func (w_ Workspace) URLsForApplicationsToOpenURL(url foundation.objc.IObject /* cross-framework URL */) []foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[[]foundation.URL](w_.ID, objc.Sel("URLsForApplicationsToOpenURL:"), url)
	return rv
}


// Returns an array of URLs to all available applications that can open the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlsForApplications(withBundleIdentifier:)
func (w_ Workspace) URLsForApplicationsWithBundleIdentifier(bundleIdentifier string /* primitive/slice/pointer. */) []foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[[]foundation.URL](w_.ID, objc.Sel("URLsForApplicationsWithBundleIdentifier:"), objc.String(bundleIdentifier))
	return rv
}


// A Boolean value that indicates whether the app avoids conveying information through color alone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldDifferentiateWithoutColor
func (w_ Workspace) AccessibilityDisplayShouldDifferentiateWithoutColor() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldDifferentiateWithoutColor"))
	return rv
}


// A Boolean value that indicates whether the app presents a high-contrast user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldIncreaseContrast
func (w_ Workspace) AccessibilityDisplayShouldIncreaseContrast() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldIncreaseContrast"))
	return rv
}


// A Boolean value that indicates whether the accessibility option to invert colors is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldInvertColors
func (w_ Workspace) AccessibilityDisplayShouldInvertColors() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldInvertColors"))
	return rv
}


// A Boolean value that indicates whether the accessibility option to reduce motion is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldReduceMotion
func (w_ Workspace) AccessibilityDisplayShouldReduceMotion() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldReduceMotion"))
	return rv
}


// A Boolean value that indicates whether the app avoids using semitransparent backgrounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldReduceTransparency
func (w_ Workspace) AccessibilityDisplayShouldReduceTransparency() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldReduceTransparency"))
	return rv
}


// The array of colors for the file labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/fileLabelColors
func (w_ Workspace) FileLabelColors() []Color /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Color](w_.ID, objc.Sel("fileLabelColors"))
	return rv
}


// The array of file labels, returned as strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/fileLabels
func (w_ Workspace) FileLabels() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](w_.ID, objc.Sel("fileLabels"))
	return rv
}


// Returns the frontmost app, which is the app that receives key events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/frontmostApplication
func (w_ Workspace) FrontmostApplication() IRunningApplication {
	rv := objc.Send[RunningApplication](w_.ID, objc.Sel("frontmostApplication"))
	return rv
}


// A Boolean value that indicates whether Switch Control is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/isSwitchControlEnabled
func (w_ Workspace) SwitchControlEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("switchControlEnabled"))
	return rv
}


// A Boolean value that indicates whether VoiceOver is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/isVoiceOverEnabled
func (w_ Workspace) VoiceOverEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("voiceOverEnabled"))
	return rv
}


// Returns the app that owns the currently displayed menu bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/menuBarOwningApplication
func (w_ Workspace) MenuBarOwningApplication() IRunningApplication {
	rv := objc.Send[RunningApplication](w_.ID, objc.Sel("menuBarOwningApplication"))
	return rv
}


// The notification center for workspace notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/notificationCenter
func (w_ Workspace) NotificationCenter() NotificationCenter /* not a class type */ {
	rv := objc.Send[NotificationCenter](w_.ID, objc.Sel("notificationCenter"))
	return rv
}


// Returns an array of running apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/runningApplications
func (w_ Workspace) RunningApplications() []RunningApplication /* primitive/slice/pointer. */ {
	rv := objc.Send[[]RunningApplication](w_.ID, objc.Sel("runningApplications"))
	return rv
}


// The shared workspace object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/shared
func (w_ Workspace) SharedWorkspace() IWorkspace {
	rv := objc.Send[Workspace](w_.ID, objc.Sel("sharedWorkspace"))
	return rv
}


// A Boolean value that indicates whether Switch Control is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isswitchcontrolenabled
func (w_ Workspace) IsSwitchControlEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("isSwitchControlEnabled"))
	return rv
}


// A Boolean value that indicates whether Switch Control is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isswitchcontrolenabled
func (w_ Workspace) SetIsSwitchControlEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsSwitchControlEnabled:"), value)
}


// A Boolean value that indicates whether VoiceOver is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isvoiceoverenabled
func (w_ Workspace) IsVoiceOverEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("isVoiceOverEnabled"))
	return rv
}


// A Boolean value that indicates whether VoiceOver is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isvoiceoverenabled
func (w_ Workspace) SetIsVoiceOverEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsVoiceOverEnabled:"), value)
}



