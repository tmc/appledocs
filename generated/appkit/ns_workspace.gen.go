// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSWorkspace */


/* debug [class_header]: Header for NSWorkspace */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Workspace */
// An interface definition for the [Workspace] class.
type IWorkspace interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Workspace */
	// properties:
	AccessibilityDisplayShouldDifferentiateWithoutColor() bool
	AccessibilityDisplayShouldIncreaseContrast() bool
	AccessibilityDisplayShouldInvertColors() bool
	AccessibilityDisplayShouldReduceMotion() bool
	AccessibilityDisplayShouldReduceTransparency() bool
	FileLabelColors() []Color
	FileLabels() []string
	FrontmostApplication() IRunningApplication
	SwitchControlEnabled() bool
	VoiceOverEnabled() bool
	MenuBarOwningApplication() IRunningApplication
	NotificationCenter() foundation.NotificationCenter
	RunningApplications() []RunningApplication
	IsSwitchControlEnabled() bool
	SetIsSwitchControlEnabled(value bool)
	IsVoiceOverEnabled() bool
	SetIsVoiceOverEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Workspace */
	// methods:
	ActivateFileViewerSelectingURLs(fileURLs []foundation.URL)
	DesktopImageOptionsForScreen(screen IScreen) foundation.IDictionary
	DesktopImageURLForScreen(screen IScreen) foundation.URL
	DuplicateURLsCompletionHandler(URLs []foundation.URL, handler unsafe.Pointer)
	ExtendPowerOffBy(requested int) int
	GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType(fullPath objc.IObject /* cross-framework: NSString */, removableFlag objectivec.IObject, writableFlag objectivec.IObject, unmountableFlag objectivec.IObject, description objc.IObject /* cross-framework: NSString */, fileSystemType objc.IObject /* cross-framework: NSString */) bool
	HideOtherApplications()
	IconForContentType(contentType uniformtypeidentifiers.UTType) IImage
	IconForFile(fullPath objc.IObject /* cross-framework: NSString */) IImage
	IconForFiles(fullPaths []string) IImage
	IsFilePackageAtPath(fullPath objc.IObject /* cross-framework: NSString */) bool
	NoteFileSystemChangedWithPath(path objc.IObject /* cross-framework: NSString */)
	OpenURL(url objc.IObject /* cross-framework: NSURL */) bool
	OpenURLConfigurationCompletionHandler(url objc.IObject /* cross-framework: NSURL */, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer)
	OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls []foundation.URL, applicationURL objc.IObject /* cross-framework: NSURL */, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer)
	OpenApplicationAtURLConfigurationCompletionHandler(applicationURL objc.IObject /* cross-framework: NSURL */, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer)
	RecycleURLsCompletionHandler(URLs []foundation.URL, handler unsafe.Pointer)
	RequestAuthorizationOfTypeCompletionHandler(type_ WorkspaceAuthorizationType, completionHandler unsafe.Pointer)
	SelectFileInFileViewerRootedAtPath(fullPath objc.IObject /* cross-framework: NSString */, rootFullPath objc.IObject /* cross-framework: NSString */) bool
	SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler(applicationURL objc.IObject /* cross-framework: NSURL */, contentType uniformtypeidentifiers.UTType, completionHandler unsafe.Pointer)
	SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler(applicationURL objc.IObject /* cross-framework: NSURL */, url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer)
	SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler(applicationURL objc.IObject /* cross-framework: NSURL */, url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer)
	SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler(applicationURL objc.IObject /* cross-framework: NSURL */, urlScheme objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	SetDesktopImageURLForScreenOptionsError(url objc.IObject /* cross-framework: NSURL */, screen IScreen, options foundation.IDictionary, error_ objectivec.IObject) bool
	SetIconForFileOptions(image IImage, fullPath objc.IObject /* cross-framework: NSString */, options WorkspaceIconCreationOptions) bool
	ShowSearchResultsForQueryString(queryString objc.IObject /* cross-framework: NSString */) bool
	UnmountAndEjectDeviceAtURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) bool
	UnmountAndEjectDeviceAtPath(path objc.IObject /* cross-framework: NSString */) bool
	URLForApplicationToOpenURL(url objc.IObject /* cross-framework: NSURL */) foundation.URL
	URLForApplicationToOpenContentType(contentType uniformtypeidentifiers.UTType) foundation.URL
	URLForApplicationWithBundleIdentifier(bundleIdentifier objc.IObject /* cross-framework: NSString */) foundation.URL
	URLsForApplicationsToOpenContentType(contentType uniformtypeidentifiers.UTType) []foundation.URL
	URLsForApplicationsToOpenURL(url objc.IObject /* cross-framework: NSURL */) []foundation.URL
	URLsForApplicationsWithBundleIdentifier(bundleIdentifier objc.IObject /* cross-framework: NSString */) []foundation.URL
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Workspace */
// Alloc allocates a new instance without initialization.
func (wc _WorkspaceClass) Alloc() Workspace {
	rv := objc.Send[Workspace](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Workspace */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Workspace *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Workspace */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Workspace */

// The shared workspace object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/shared
func (wc _WorkspaceClass) SharedWorkspace() Workspace {
	rv := objc.Send[Workspace](objc.ID(wc.class), objc.Sel("sharedWorkspace"))
	return rv
}/* debug [class_properties_class/property]: sharedWorkspace */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Workspace */

// Activates the Finder, and opens one or more windows selecting the specified files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/activateFileViewerSelecting(_:)
func (w_ Workspace) ActivateFileViewerSelectingURLs(fileURLs []foundation.URL) {
	objc.Send[objc.ID](w_.ID, objc.Sel("activateFileViewerSelectingURLs:"), fileURLs)
}/* debug [instance_methods/method]: ActivateFileViewerSelectingURLs */


// Returns the desktop image options for the given screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/desktopImageOptions(for:)
func (w_ Workspace) DesktopImageOptionsForScreen(screen IScreen) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](w_.ID, objc.Sel("desktopImageOptionsForScreen:"), screen)
	return rv
}/* debug [instance_methods/method]: DesktopImageOptionsForScreen */


// Returns the URL for the desktop image for the given screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/desktopImageURL(for:)
func (w_ Workspace) DesktopImageURLForScreen(screen IScreen) foundation.URL {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("desktopImageURLForScreen:"), screen)
	return rv
}/* debug [instance_methods/method]: DesktopImageURLForScreen */


// Duplicates the specified URLS asynchronously in the same manner as the Finder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/duplicate(_:completionHandler:)
func (w_ Workspace) DuplicateURLsCompletionHandler(URLs []foundation.URL, handler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("duplicateURLs:completionHandler:"), URLs, handler)
}/* debug [instance_methods/method]: DuplicateURLsCompletionHandler */


// Requests the system wait for the specified amount of time before turning off the power or logging out the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/extendPowerOff(by:)
func (w_ Workspace) ExtendPowerOffBy(requested int) int {
	rv := objc.Send[int](w_.ID, objc.Sel("extendPowerOffBy:"), requested)
	return rv
}/* debug [instance_methods/method]: ExtendPowerOffBy */


// Returns information about the file system at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/getFileSystemInfo(forPath:isRemovable:isWritable:isUnmountable:description:type:)
func (w_ Workspace) GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType(fullPath objc.IObject /* cross-framework: NSString */, removableFlag objectivec.IObject, writableFlag objectivec.IObject, unmountableFlag objectivec.IObject, description objc.IObject /* cross-framework: NSString */, fileSystemType objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("getFileSystemInfoForPath:isRemovable:isWritable:isUnmountable:description:type:"), fullPath, removableFlag, writableFlag, unmountableFlag, description, fileSystemType)
	return rv
}/* debug [instance_methods/method]: GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType */


// Hides all applications other than the sender.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/hideOtherApplications()
func (w_ Workspace) HideOtherApplications() {
	objc.Send[objc.ID](w_.ID, objc.Sel("hideOtherApplications"))
}/* debug [instance_methods/method]: HideOtherApplications */


// Returns an image containing the icon for the specified content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/icon(for:)
func (w_ Workspace) IconForContentType(contentType uniformtypeidentifiers.UTType) IImage {
	rv := objc.Send[Image](w_.ID, objc.Sel("iconForContentType:"), contentType)
	return rv
}/* debug [instance_methods/method]: IconForContentType */


// Returns an image containing the icon for the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/icon(forFile:)
func (w_ Workspace) IconForFile(fullPath objc.IObject /* cross-framework: NSString */) IImage {
	rv := objc.Send[Image](w_.ID, objc.Sel("iconForFile:"), fullPath)
	return rv
}/* debug [instance_methods/method]: IconForFile */


// Returns an image containing the icon for the specified files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/icon(forFiles:)
func (w_ Workspace) IconForFiles(fullPaths []string) IImage {
	rv := objc.Send[Image](w_.ID, objc.Sel("iconForFiles:"), fullPaths)
	return rv
}/* debug [instance_methods/method]: IconForFiles */


// Determines whether the specified path is a file package.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/isFilePackage(atPath:)
func (w_ Workspace) IsFilePackageAtPath(fullPath objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isFilePackageAtPath:"), fullPath)
	return rv
}/* debug [instance_methods/method]: IsFilePackageAtPath */


// Informs the workspace object that the file system changed at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/noteFileSystemChanged(_:)
func (w_ Workspace) NoteFileSystemChangedWithPath(path objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("noteFileSystemChanged:"), path)
}/* debug [instance_methods/method]: NoteFileSystemChangedWithPath */


// Opens the location at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:)
func (w_ Workspace) OpenURL(url objc.IObject /* cross-framework: NSURL */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("openURL:"), url)
	return rv
}/* debug [instance_methods/method]: OpenURL */


// Opens a URL asynchronously using the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:configuration:completionHandler:)
func (w_ Workspace) OpenURLConfigurationCompletionHandler(url objc.IObject /* cross-framework: NSURL */, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("openURL:configuration:completionHandler:"), url, configuration, completionHandler)
}/* debug [instance_methods/method]: OpenURLConfigurationCompletionHandler */


// Opens one or more URLs asynchronously in the specified app using the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:withApplicationAt:configuration:completionHandler:)
func (w_ Workspace) OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls []foundation.URL, applicationURL objc.IObject /* cross-framework: NSURL */, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("openURLs:withApplicationAtURL:configuration:completionHandler:"), urls, applicationURL, configuration, completionHandler)
}/* debug [instance_methods/method]: OpenURLsWithApplicationAtURLConfigurationCompletionHandler */


// Launches the app at the specified URL and asynchronously reports back on the app’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/openApplication(at:configuration:completionHandler:)
func (w_ Workspace) OpenApplicationAtURLConfigurationCompletionHandler(applicationURL objc.IObject /* cross-framework: NSURL */, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("openApplicationAtURL:configuration:completionHandler:"), applicationURL, configuration, completionHandler)
}/* debug [instance_methods/method]: OpenApplicationAtURLConfigurationCompletionHandler */


// Moves the specified URLs to the trash in the same manner as the Finder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/recycle(_:completionHandler:)
func (w_ Workspace) RecycleURLsCompletionHandler(URLs []foundation.URL, handler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("recycleURLs:completionHandler:"), URLs, handler)
}/* debug [instance_methods/method]: RecycleURLsCompletionHandler */


// Requests authorization to perform a privileged file operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/requestAuthorization(to:completionHandler:)
func (w_ Workspace) RequestAuthorizationOfTypeCompletionHandler(type_ WorkspaceAuthorizationType, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("requestAuthorizationOfType:completionHandler:"), type_, completionHandler)
}/* debug [instance_methods/method]: RequestAuthorizationOfTypeCompletionHandler */


// Selects the file at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/selectFile(_:inFileViewerRootedAtPath:)
func (w_ Workspace) SelectFileInFileViewerRootedAtPath(fullPath objc.IObject /* cross-framework: NSString */, rootFullPath objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("selectFile:inFileViewerRootedAtPath:"), fullPath, rootFullPath)
	return rv
}/* debug [instance_methods/method]: SelectFileInFileViewerRootedAtPath */


// Sets the default app to use when opening files of a specific content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpen:completion:)
func (w_ Workspace) SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler(applicationURL objc.IObject /* cross-framework: NSURL */, contentType uniformtypeidentifiers.UTType, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultApplicationAtURL:toOpenContentType:completionHandler:"), applicationURL, contentType, completionHandler)
}/* debug [instance_methods/method]: SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler */


// Sets the default app to use when opening files of a specific content type defined by a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpenContentTypeOfFileAt:completion:)
func (w_ Workspace) SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler(applicationURL objc.IObject /* cross-framework: NSURL */, url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultApplicationAtURL:toOpenContentTypeOfFileAtURL:completionHandler:"), applicationURL, url, completionHandler)
}/* debug [instance_methods/method]: SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler */


// Sets the default app to use when opening a specific file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpenFileAt:completion:)
func (w_ Workspace) SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler(applicationURL objc.IObject /* cross-framework: NSURL */, url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultApplicationAtURL:toOpenFileAtURL:completionHandler:"), applicationURL, url, completionHandler)
}/* debug [instance_methods/method]: SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler */


// Sets the default app to use when opening files of a specific scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpenURLsWithScheme:completion:)
func (w_ Workspace) SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler(applicationURL objc.IObject /* cross-framework: NSURL */, urlScheme objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultApplicationAtURL:toOpenURLsWithScheme:completionHandler:"), applicationURL, urlScheme, completionHandler)
}/* debug [instance_methods/method]: SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler */


// Sets the desktop image for the given screen to the image at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDesktopImageURL(_:for:options:)
func (w_ Workspace) SetDesktopImageURLForScreenOptionsError(url objc.IObject /* cross-framework: NSURL */, screen IScreen, options foundation.IDictionary, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("setDesktopImageURL:forScreen:options:error:"), url, screen, options, error_)
	return rv
}/* debug [instance_methods/method]: SetDesktopImageURLForScreenOptionsError */


// Sets the icon for the file or directory at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setIcon(_:forFile:options:)
func (w_ Workspace) SetIconForFileOptions(image IImage, fullPath objc.IObject /* cross-framework: NSString */, options WorkspaceIconCreationOptions) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("setIcon:forFile:options:"), image, fullPath, options)
	return rv
}/* debug [instance_methods/method]: SetIconForFileOptions */


// Displays a Spotlight search results window in Finder for the specified query string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/showSearchResults(forQueryString:)
func (w_ Workspace) ShowSearchResultsForQueryString(queryString objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("showSearchResultsForQueryString:"), queryString)
	return rv
}/* debug [instance_methods/method]: ShowSearchResultsForQueryString */


// Attempts to eject the volume mounted at the given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/unmountAndEjectDevice(at:)
func (w_ Workspace) UnmountAndEjectDeviceAtURLError(url objc.IObject /* cross-framework: NSURL */, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("unmountAndEjectDeviceAtURL:error:"), url, error_)
	return rv
}/* debug [instance_methods/method]: UnmountAndEjectDeviceAtURLError */


// Unmounts and ejects the device at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/unmountAndEjectDevice(atPath:)
func (w_ Workspace) UnmountAndEjectDeviceAtPath(path objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("unmountAndEjectDeviceAtPath:"), path)
	return rv
}/* debug [instance_methods/method]: UnmountAndEjectDeviceAtPath */


// Returns the URL to the default app to open the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlForApplication(toOpen:)-7qkzf
func (w_ Workspace) URLForApplicationToOpenURL(url objc.IObject /* cross-framework: NSURL */) foundation.URL {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("URLForApplicationToOpenURL:"), url)
	return rv
}/* debug [instance_methods/method]: URLForApplicationToOpenURL */


// Returns the URL to the default app to open the specified content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlForApplication(toOpen:)-95cvp
func (w_ Workspace) URLForApplicationToOpenContentType(contentType uniformtypeidentifiers.UTType) foundation.URL {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("URLForApplicationToOpenContentType:"), contentType)
	return rv
}/* debug [instance_methods/method]: URLForApplicationToOpenContentType */


// Returns the URL to the default app with the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlForApplication(withBundleIdentifier:)
func (w_ Workspace) URLForApplicationWithBundleIdentifier(bundleIdentifier objc.IObject /* cross-framework: NSString */) foundation.URL {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("URLForApplicationWithBundleIdentifier:"), bundleIdentifier)
	return rv
}/* debug [instance_methods/method]: URLForApplicationWithBundleIdentifier */


// Returns an array of URLs to all available applications that can open the specified content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlsForApplications(toOpen:)-60rkm
func (w_ Workspace) URLsForApplicationsToOpenContentType(contentType uniformtypeidentifiers.UTType) []foundation.URL {
	rv := objc.Send[[]foundation.URL](w_.ID, objc.Sel("URLsForApplicationsToOpenContentType:"), contentType)
	return rv
}/* debug [instance_methods/method]: URLsForApplicationsToOpenContentType */


// Returns an array of URLs to all available applications that can open the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlsForApplications(toOpen:)-ualk
func (w_ Workspace) URLsForApplicationsToOpenURL(url objc.IObject /* cross-framework: NSURL */) []foundation.URL {
	rv := objc.Send[[]foundation.URL](w_.ID, objc.Sel("URLsForApplicationsToOpenURL:"), url)
	return rv
}/* debug [instance_methods/method]: URLsForApplicationsToOpenURL */


// Returns an array of URLs to all available applications that can open the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlsForApplications(withBundleIdentifier:)
func (w_ Workspace) URLsForApplicationsWithBundleIdentifier(bundleIdentifier objc.IObject /* cross-framework: NSString */) []foundation.URL {
	rv := objc.Send[[]foundation.URL](w_.ID, objc.Sel("URLsForApplicationsWithBundleIdentifier:"), bundleIdentifier)
	return rv
}/* debug [instance_methods/method]: URLsForApplicationsWithBundleIdentifier */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Workspace */

// A Boolean value that indicates whether the app avoids conveying information through color alone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldDifferentiateWithoutColor
func (w_ Workspace) AccessibilityDisplayShouldDifferentiateWithoutColor() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldDifferentiateWithoutColor"))
	return rv
}/* debug [instance_properties/getter]: accessibilityDisplayShouldDifferentiateWithoutColor */


// A Boolean value that indicates whether the app presents a high-contrast user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldIncreaseContrast
func (w_ Workspace) AccessibilityDisplayShouldIncreaseContrast() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldIncreaseContrast"))
	return rv
}/* debug [instance_properties/getter]: accessibilityDisplayShouldIncreaseContrast */


// A Boolean value that indicates whether the accessibility option to invert colors is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldInvertColors
func (w_ Workspace) AccessibilityDisplayShouldInvertColors() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldInvertColors"))
	return rv
}/* debug [instance_properties/getter]: accessibilityDisplayShouldInvertColors */


// A Boolean value that indicates whether the accessibility option to reduce motion is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldReduceMotion
func (w_ Workspace) AccessibilityDisplayShouldReduceMotion() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldReduceMotion"))
	return rv
}/* debug [instance_properties/getter]: accessibilityDisplayShouldReduceMotion */


// A Boolean value that indicates whether the app avoids using semitransparent backgrounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldReduceTransparency
func (w_ Workspace) AccessibilityDisplayShouldReduceTransparency() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldReduceTransparency"))
	return rv
}/* debug [instance_properties/getter]: accessibilityDisplayShouldReduceTransparency */


// The array of colors for the file labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/fileLabelColors
func (w_ Workspace) FileLabelColors() []Color {
	rv := objc.Send[[]Color](w_.ID, objc.Sel("fileLabelColors"))
	return rv
}/* debug [instance_properties/getter]: fileLabelColors */


// The array of file labels, returned as strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/fileLabels
func (w_ Workspace) FileLabels() []string {
	rv := objc.Send[[]string](w_.ID, objc.Sel("fileLabels"))
	return rv
}/* debug [instance_properties/getter]: fileLabels */


// Returns the frontmost app, which is the app that receives key events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/frontmostApplication
func (w_ Workspace) FrontmostApplication() IRunningApplication {
	rv := objc.Send[RunningApplication](w_.ID, objc.Sel("frontmostApplication"))
	return rv
}/* debug [instance_properties/getter]: frontmostApplication */


// A Boolean value that indicates whether Switch Control is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/isSwitchControlEnabled
func (w_ Workspace) SwitchControlEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("switchControlEnabled"))
	return rv
}/* debug [instance_properties/getter]: switchControlEnabled */


// A Boolean value that indicates whether VoiceOver is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/isVoiceOverEnabled
func (w_ Workspace) VoiceOverEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("voiceOverEnabled"))
	return rv
}/* debug [instance_properties/getter]: voiceOverEnabled */


// Returns the app that owns the currently displayed menu bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/menuBarOwningApplication
func (w_ Workspace) MenuBarOwningApplication() IRunningApplication {
	rv := objc.Send[RunningApplication](w_.ID, objc.Sel("menuBarOwningApplication"))
	return rv
}/* debug [instance_properties/getter]: menuBarOwningApplication */


// The notification center for workspace notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/notificationCenter
func (w_ Workspace) NotificationCenter() foundation.NotificationCenter {
	rv := objc.Send[foundation.NotificationCenter](w_.ID, objc.Sel("notificationCenter"))
	return rv
}/* debug [instance_properties/getter]: notificationCenter */


// Returns an array of running apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/runningApplications
func (w_ Workspace) RunningApplications() []RunningApplication {
	rv := objc.Send[[]RunningApplication](w_.ID, objc.Sel("runningApplications"))
	return rv
}/* debug [instance_properties/getter]: runningApplications */


// The shared workspace object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/shared
func (w_ Workspace) SharedWorkspace() IWorkspace {
	rv := objc.Send[Workspace](w_.ID, objc.Sel("sharedWorkspace"))
	return rv
}/* debug [instance_properties/getter]: sharedWorkspace */


// A Boolean value that indicates whether Switch Control is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isswitchcontrolenabled
func (w_ Workspace) IsSwitchControlEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isSwitchControlEnabled"))
	return rv
}/* debug [instance_properties/getter]: isSwitchControlEnabled */


// A Boolean value that indicates whether Switch Control is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isswitchcontrolenabled
func (w_ Workspace) SetIsSwitchControlEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsSwitchControlEnabled:"), value)
}/* debug [instance_properties/setter]: isSwitchControlEnabled */


// A Boolean value that indicates whether VoiceOver is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isvoiceoverenabled
func (w_ Workspace) IsVoiceOverEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isVoiceOverEnabled"))
	return rv
}/* debug [instance_properties/getter]: isVoiceOverEnabled */


// A Boolean value that indicates whether VoiceOver is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isvoiceoverenabled
func (w_ Workspace) SetIsVoiceOverEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsVoiceOverEnabled:"), value)
}/* debug [instance_properties/setter]: isVoiceOverEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSWorkspace */



