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
	AbsolutePathForAppBundleWithIdentifier(bundleIdentifier string) foundation.String
	ActivateFileViewerSelectingURLs(fileURLs []foundation.IURL)
	ActiveApplication() foundation.Dictionary
	CheckForRemovableMedia()
	DesktopImageOptionsForScreen(screen IScreen) unsafe.Pointer
	DesktopImageURLForScreen(screen IScreen) foundation.URL
	DuplicateURLsCompletionHandler(URLs []foundation.IURL, handler unsafe.Pointer)
	ExtendPowerOffBy(requested int) int
	FileSystemChanged() bool
	FilenameExtensionIsValidForType(filenameExtension string, typeName string) bool
	FindApplications()
	FullPathForApplication(appName string) foundation.String
	GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType(fullPath string, removableFlag unsafe.Pointer, writableFlag unsafe.Pointer, unmountableFlag unsafe.Pointer, description string, fileSystemType string) bool
	GetInfoForFileApplicationType(fullPath string, appName string, type_ string) bool
	HideOtherApplications()
	IconForContentType(contentType unsafe.Pointer) Image
	IconForFile(fullPath string) Image
	IconForFileType(fileType string) Image
	IconForFiles(fullPaths []string) Image
	IsFilePackageAtPath(fullPath string) bool
	LaunchApplication(appName string) bool
	LaunchApplicationShowIconAutolaunch(appName string, showIcon bool, autolaunch bool) bool
	LaunchApplicationAtURLOptionsConfigurationError(url foundation.IURL, options WorkspaceLaunchOptions, configuration unsafe.Pointer, error_ unsafe.Pointer) RunningApplication
	LaunchAppWithBundleIdentifierOptionsAdditionalEventParamDescriptorLaunchIdentifier(bundleIdentifier string, options WorkspaceLaunchOptions, descriptor foundation.IAppleEventDescriptor, identifier foundation.INumber) bool
	LaunchedApplications() foundation.Array
	LocalizedDescriptionForType(typeName string) foundation.String
	MountNewRemovableMedia() foundation.Array
	MountedLocalVolumePaths() foundation.Array
	MountedRemovableMedia() foundation.Array
	NoteFileSystemChanged(path string)
	NoteUserDefaultsChanged()
	OpenURL(url foundation.IURL) bool
	OpenURLConfigurationCompletionHandler(url foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer)
	OpenURLOptionsConfigurationError(url foundation.IURL, options WorkspaceLaunchOptions, configuration unsafe.Pointer, error_ unsafe.Pointer) RunningApplication
	OpenURLsWithAppBundleIdentifierOptionsAdditionalEventParamDescriptorLaunchIdentifiers(urls []foundation.IURL, bundleIdentifier string, options WorkspaceLaunchOptions, descriptor foundation.IAppleEventDescriptor, identifiers []foundation.INumber) bool
	OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls []foundation.IURL, applicationURL foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer)
	OpenURLsWithApplicationAtURLOptionsConfigurationError(urls []foundation.IURL, applicationURL foundation.IURL, options WorkspaceLaunchOptions, configuration unsafe.Pointer, error_ unsafe.Pointer) RunningApplication
	OpenApplicationAtURLConfigurationCompletionHandler(applicationURL foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer)
	OpenFile(fullPath string) bool
	OpenFileFromImageAtInView(fullPath string, image IImage, point coregraphics.CGPoint, view IView) bool
	OpenFileWithApplication(fullPath string, appName string) bool
	OpenFileWithApplicationAndDeactivate(fullPath string, appName string, flag bool) bool
	OpenTempFile(fullPath string) bool
	PerformFileOperationSourceDestinationFilesTag(operation IWorkspaceFileOperationName, source string, destination string, files objectivec.IObject, tag unsafe.Pointer) bool
	PreferredFilenameExtensionForType(typeName string) foundation.String
	RecycleURLsCompletionHandler(URLs []foundation.IURL, handler unsafe.Pointer)
	RequestAuthorizationOfTypeCompletionHandler(type_ WorkspaceAuthorizationType, completionHandler unsafe.Pointer)
	SelectFileInFileViewerRootedAtPath(fullPath string, rootFullPath string) bool
	SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler(applicationURL foundation.IURL, contentType unsafe.Pointer, completionHandler unsafe.Pointer)
	SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler unsafe.Pointer)
	SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler unsafe.Pointer)
	SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler(applicationURL foundation.IURL, urlScheme string, completionHandler unsafe.Pointer)
	SetDesktopImageURLForScreenOptionsError(url foundation.IURL, screen IScreen, options unsafe.Pointer, error_ unsafe.Pointer) bool
	SetIconForFileOptions(image IImage, fullPath string, options WorkspaceIconCreationOptions) bool
	ShowSearchResultsForQueryString(queryString string) bool
	SlideImageFromTo(image IImage, fromPoint coregraphics.CGPoint, toPoint coregraphics.CGPoint)
	TypeConformsToType(firstTypeName string, secondTypeName string) bool
	TypeOfFileError(absoluteFilePath string, outError unsafe.Pointer) foundation.String
	UnmountAndEjectDeviceAtURLError(url foundation.IURL, error_ unsafe.Pointer) bool
	UnmountAndEjectDeviceAtPath(path string) bool
	URLForApplicationToOpenURL(url foundation.IURL) foundation.URL
	URLForApplicationToOpenContentType(contentType unsafe.Pointer) foundation.URL
	URLForApplicationWithBundleIdentifier(bundleIdentifier string) foundation.URL
	URLsForApplicationsToOpenContentType(contentType unsafe.Pointer) []foundation.URL
	URLsForApplicationsToOpenURL(url foundation.IURL) []foundation.URL
	URLsForApplicationsWithBundleIdentifier(bundleIdentifier string) []foundation.URL
	UserDefaultsChanged() bool
	AccessibilityDisplayShouldDifferentiateWithoutColor() bool
	AccessibilityDisplayShouldIncreaseContrast() bool
	AccessibilityDisplayShouldInvertColors() bool
	AccessibilityDisplayShouldReduceMotion() bool
	AccessibilityDisplayShouldReduceTransparency() bool
	FileLabelColors() []Color
	FileLabels() []string
	FrontmostApplication() NSRunningApplication
	SwitchControlEnabled() bool
	VoiceOverEnabled() bool
	MenuBarOwningApplication() NSRunningApplication
	NotificationCenter() foundation.NotificationCenter
	RunningApplications() []RunningApplication
	IsSwitchControlEnabled() bool
	SetIsSwitchControlEnabled(value bool)
	IsVoiceOverEnabled() bool
	SetIsVoiceOverEnabled(value bool)
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
	rv := objc.Send[NSWorkspace](objc.ID(wc.class), objc.Sel("sharedWorkspace"))
	return rv
}

// Returns the absolute file system path of an app bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/absolutePathForApplication(withBundleIdentifier:)
func (w_ Workspace) AbsolutePathForAppBundleWithIdentifier(bundleIdentifier string) foundation.String {
	rv := objc.Send[foundation.String](w_.ID, objc.Sel("absolutePathForAppBundleWithIdentifier:"), objc.String(bundleIdentifier))
	return rv
}


// Activates the Finder, and opens one or more windows selecting the specified files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/activateFileViewerSelecting(_:)
func (w_ Workspace) ActivateFileViewerSelectingURLs(fileURLs []foundation.IURL) {
	objc.Send[objc.ID](w_.ID, objc.Sel("activateFileViewerSelectingURLs:"), fileURLs)
}


// Returns a dictionary with information about the current active app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/activeApplication()
func (w_ Workspace) ActiveApplication() foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](w_.ID, objc.Sel("activeApplication"))
	return rv
}


// Polls the system’s drives for any disks that have been inserted but not yet mounted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/checkForRemovableMedia
func (w_ Workspace) CheckForRemovableMedia() {
	objc.Send[objc.ID](w_.ID, objc.Sel("checkForRemovableMedia"))
}


// Returns the desktop image options for the given screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/desktopImageOptions(for:)
func (w_ Workspace) DesktopImageOptionsForScreen(screen IScreen) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("desktopImageOptionsForScreen:"), screen)
	return rv
}


// Returns the URL for the desktop image for the given screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/desktopImageURL(for:)
func (w_ Workspace) DesktopImageURLForScreen(screen IScreen) foundation.URL {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("desktopImageURLForScreen:"), screen)
	return rv
}


// Duplicates the specified URLS asynchronously in the same manner as the Finder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/duplicate(_:completionHandler:)
func (w_ Workspace) DuplicateURLsCompletionHandler(URLs []foundation.IURL, handler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("duplicateURLs:completionHandler:"), URLs, handler)
}


// Requests the system wait for the specified amount of time before turning off the power or logging out the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/extendPowerOff(by:)
func (w_ Workspace) ExtendPowerOffBy(requested int) int {
	rv := objc.Send[int](w_.ID, objc.Sel("extendPowerOffBy:"), requested)
	return rv
}


// Returns a Boolean value indicating whether a change to the file system has been registered with a message since the last message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/fileSystemChanged
func (w_ Workspace) FileSystemChanged() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("fileSystemChanged"))
	return rv
}


// Returns whether the specified filename extension is appropriate for the Uniform Type Identifier (UTI).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/filenameExtension(_:isValidForType:)
func (w_ Workspace) FilenameExtensionIsValidForType(filenameExtension string, typeName string) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("filenameExtension:isValidForType:"), objc.String(filenameExtension), objc.String(typeName))
	return rv
}


// Examines all apps and updates the records of registered services and file types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/findApplications
func (w_ Workspace) FindApplications() {
	objc.Send[objc.ID](w_.ID, objc.Sel("findApplications"))
}


// Returns the full path for the specified app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/fullPath(forApplication:)
func (w_ Workspace) FullPathForApplication(appName string) foundation.String {
	rv := objc.Send[foundation.String](w_.ID, objc.Sel("fullPathForApplication:"), objc.String(appName))
	return rv
}


// Returns information about the file system at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/getFileSystemInfo(forPath:isRemovable:isWritable:isUnmountable:description:type:)
func (w_ Workspace) GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType(fullPath string, removableFlag unsafe.Pointer, writableFlag unsafe.Pointer, unmountableFlag unsafe.Pointer, description string, fileSystemType string) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("getFileSystemInfoForPath:isRemovable:isWritable:isUnmountable:description:type:"), objc.String(fullPath), removableFlag, writableFlag, unmountableFlag, objc.String(description), objc.String(fileSystemType))
	return rv
}


// Retrieves information about the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/getInfoForFile(_:application:type:)
func (w_ Workspace) GetInfoForFileApplicationType(fullPath string, appName string, type_ string) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("getInfoForFile:application:type:"), objc.String(fullPath), objc.String(appName), objc.String(type_))
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
func (w_ Workspace) IconForContentType(contentType unsafe.Pointer) Image {
	rv := objc.Send[Image](w_.ID, objc.Sel("iconForContentType:"), contentType)
	return rv
}


// Returns an image containing the icon for the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/icon(forFile:)
func (w_ Workspace) IconForFile(fullPath string) Image {
	rv := objc.Send[Image](w_.ID, objc.Sel("iconForFile:"), objc.String(fullPath))
	return rv
}


// Returns an image containing the icon for files of the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/icon(forFileType:)
func (w_ Workspace) IconForFileType(fileType string) Image {
	rv := objc.Send[Image](w_.ID, objc.Sel("iconForFileType:"), objc.String(fileType))
	return rv
}


// Returns an image containing the icon for the specified files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/icon(forFiles:)
func (w_ Workspace) IconForFiles(fullPaths []string) Image {
	rv := objc.Send[Image](w_.ID, objc.Sel("iconForFiles:"), fullPaths)
	return rv
}


// Determines whether the specified path is a file package.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/isFilePackage(atPath:)
func (w_ Workspace) IsFilePackageAtPath(fullPath string) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isFilePackageAtPath:"), objc.String(fullPath))
	return rv
}


// Launches the specified app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/launchApplication(_:)
func (w_ Workspace) LaunchApplication(appName string) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("launchApplication:"), objc.String(appName))
	return rv
}


// Launches the specified app using additional options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/launchApplication(_:showIcon:autolaunch:)
func (w_ Workspace) LaunchApplicationShowIconAutolaunch(appName string, showIcon bool, autolaunch bool) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("launchApplication:showIcon:autolaunch:"), objc.String(appName), showIcon, autolaunch)
	return rv
}


// Launches the app at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/launchApplication(at:options:configuration:)
func (w_ Workspace) LaunchApplicationAtURLOptionsConfigurationError(url foundation.IURL, options WorkspaceLaunchOptions, configuration unsafe.Pointer, error_ unsafe.Pointer) RunningApplication {
	rv := objc.Send[RunningApplication](w_.ID, objc.Sel("launchApplicationAtURL:options:configuration:error:"), url, options, configuration, error_)
	return rv
}


// Launches the app corresponding to the specified .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/launchApplication(withBundleIdentifier:options:additionalEventParamDescriptor:launchIdentifier:)
func (w_ Workspace) LaunchAppWithBundleIdentifierOptionsAdditionalEventParamDescriptorLaunchIdentifier(bundleIdentifier string, options WorkspaceLaunchOptions, descriptor foundation.IAppleEventDescriptor, identifier foundation.INumber) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("launchAppWithBundleIdentifier:options:additionalEventParamDescriptor:launchIdentifier:"), objc.String(bundleIdentifier), options, descriptor, identifier)
	return rv
}


// Returns an array of dictionaries, one entry for each running app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/launchedApplications
func (w_ Workspace) LaunchedApplications() foundation.Array {
	rv := objc.Send[foundation.Array](w_.ID, objc.Sel("launchedApplications"))
	return rv
}


// Returns the localized description for the specified Uniform Type Identifier (UTI).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/localizedDescription(forType:)
func (w_ Workspace) LocalizedDescriptionForType(typeName string) foundation.String {
	rv := objc.Send[foundation.String](w_.ID, objc.Sel("localizedDescriptionForType:"), objc.String(typeName))
	return rv
}


// Returns the full pathnames of any newly mounted disks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/mountNewRemovableMedia
func (w_ Workspace) MountNewRemovableMedia() foundation.Array {
	rv := objc.Send[foundation.Array](w_.ID, objc.Sel("mountNewRemovableMedia"))
	return rv
}


// Returns the mount points of all local volumes, not just the removable ones returned by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/mountedLocalVolumePaths()
func (w_ Workspace) MountedLocalVolumePaths() foundation.Array {
	rv := objc.Send[foundation.Array](w_.ID, objc.Sel("mountedLocalVolumePaths"))
	return rv
}


// Returns the full pathnames of all currently mounted removable disks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/mountedRemovableMedia()
func (w_ Workspace) MountedRemovableMedia() foundation.Array {
	rv := objc.Send[foundation.Array](w_.ID, objc.Sel("mountedRemovableMedia"))
	return rv
}


// Informs the workspace object that the file system changed at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/noteFileSystemChanged(_:)
func (w_ Workspace) NoteFileSystemChanged(path string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("noteFileSystemChanged:"), objc.String(path))
}


// Informs the object that the defaults database has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/noteUserDefaultsChanged
func (w_ Workspace) NoteUserDefaultsChanged() {
	objc.Send[objc.ID](w_.ID, objc.Sel("noteUserDefaultsChanged"))
}


// Opens the location at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:)
func (w_ Workspace) OpenURL(url foundation.IURL) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("openURL:"), url)
	return rv
}


// Opens a URL asynchronously using the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:configuration:completionHandler:)
func (w_ Workspace) OpenURLConfigurationCompletionHandler(url foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("openURL:configuration:completionHandler:"), url, configuration, completionHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:options:configuration:)
func (w_ Workspace) OpenURLOptionsConfigurationError(url foundation.IURL, options WorkspaceLaunchOptions, configuration unsafe.Pointer, error_ unsafe.Pointer) RunningApplication {
	rv := objc.Send[RunningApplication](w_.ID, objc.Sel("openURL:options:configuration:error:"), url, options, configuration, error_)
	return rv
}


// Opens one or more files from an array of URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:withAppBundleIdentifier:options:additionalEventParamDescriptor:launchIdentifiers:)
func (w_ Workspace) OpenURLsWithAppBundleIdentifierOptionsAdditionalEventParamDescriptorLaunchIdentifiers(urls []foundation.IURL, bundleIdentifier string, options WorkspaceLaunchOptions, descriptor foundation.IAppleEventDescriptor, identifiers []foundation.INumber) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("openURLs:withAppBundleIdentifier:options:additionalEventParamDescriptor:launchIdentifiers:"), urls, objc.String(bundleIdentifier), options, descriptor, identifiers)
	return rv
}


// Opens one or more URLs asynchronously in the specified app using the provided options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:withApplicationAt:configuration:completionHandler:)
func (w_ Workspace) OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls []foundation.IURL, applicationURL foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("openURLs:withApplicationAtURL:configuration:completionHandler:"), urls, applicationURL, configuration, completionHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:withApplicationAt:options:configuration:)
func (w_ Workspace) OpenURLsWithApplicationAtURLOptionsConfigurationError(urls []foundation.IURL, applicationURL foundation.IURL, options WorkspaceLaunchOptions, configuration unsafe.Pointer, error_ unsafe.Pointer) RunningApplication {
	rv := objc.Send[RunningApplication](w_.ID, objc.Sel("openURLs:withApplicationAtURL:options:configuration:error:"), urls, applicationURL, options, configuration, error_)
	return rv
}


// Launches the app at the specified URL and asynchronously reports back on the app’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/openApplication(at:configuration:completionHandler:)
func (w_ Workspace) OpenApplicationAtURLConfigurationCompletionHandler(applicationURL foundation.IURL, configuration IWorkspaceOpenConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("openApplicationAtURL:configuration:completionHandler:"), applicationURL, configuration, completionHandler)
}


// Opens the specified file specified using the default app associated with its type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/openFile(_:)
func (w_ Workspace) OpenFile(fullPath string) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("openFile:"), objc.String(fullPath))
	return rv
}


// Opens a file using the default app for its type and animates the action using a custom icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/openFile(_:from:at:in:)
func (w_ Workspace) OpenFileFromImageAtInView(fullPath string, image IImage, point coregraphics.CGPoint, view IView) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("openFile:fromImage:at:inView:"), objc.String(fullPath), image, point, view)
	return rv
}


// Opens a file using the specified app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/openFile(_:withApplication:)
func (w_ Workspace) OpenFileWithApplication(fullPath string, appName string) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("openFile:withApplication:"), objc.String(fullPath), objc.String(appName))
	return rv
}


// Opens the specified file and optionally deactivates the sending app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/openFile(_:withApplication:andDeactivate:)
func (w_ Workspace) OpenFileWithApplicationAndDeactivate(fullPath string, appName string, flag bool) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("openFile:withApplication:andDeactivate:"), objc.String(fullPath), objc.String(appName), flag)
	return rv
}


// Opens the specified temporary file using the default app for its type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/openTempFile:
func (w_ Workspace) OpenTempFile(fullPath string) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("openTempFile:"), objc.String(fullPath))
	return rv
}


// Performs a file operation on a set of files in a particular directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/performFileOperation(_:source:destination:files:tag:)
func (w_ Workspace) PerformFileOperationSourceDestinationFilesTag(operation IWorkspaceFileOperationName, source string, destination string, files objectivec.IObject, tag unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("performFileOperation:source:destination:files:tag:"), operation, objc.String(source), objc.String(destination), files, tag)
	return rv
}


// Returns the preferred filename extension for the specified Uniform Type Identifier (UTI).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/preferredFilenameExtension(forType:)
func (w_ Workspace) PreferredFilenameExtensionForType(typeName string) foundation.String {
	rv := objc.Send[foundation.String](w_.ID, objc.Sel("preferredFilenameExtensionForType:"), objc.String(typeName))
	return rv
}


// Moves the specified URLs to the trash in the same manner as the Finder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/recycle(_:completionHandler:)
func (w_ Workspace) RecycleURLsCompletionHandler(URLs []foundation.IURL, handler unsafe.Pointer) {
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
func (w_ Workspace) SelectFileInFileViewerRootedAtPath(fullPath string, rootFullPath string) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("selectFile:inFileViewerRootedAtPath:"), objc.String(fullPath), objc.String(rootFullPath))
	return rv
}


// Sets the default app to use when opening files of a specific content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpen:completion:)
func (w_ Workspace) SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler(applicationURL foundation.IURL, contentType unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultApplicationAtURL:toOpenContentType:completionHandler:"), applicationURL, contentType, completionHandler)
}


// Sets the default app to use when opening files of a specific content type defined by a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpenContentTypeOfFileAt:completion:)
func (w_ Workspace) SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultApplicationAtURL:toOpenContentTypeOfFileAtURL:completionHandler:"), applicationURL, url, completionHandler)
}


// Sets the default app to use when opening a specific file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpenFileAt:completion:)
func (w_ Workspace) SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultApplicationAtURL:toOpenFileAtURL:completionHandler:"), applicationURL, url, completionHandler)
}


// Sets the default app to use when opening files of a specific scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpenURLsWithScheme:completion:)
func (w_ Workspace) SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler(applicationURL foundation.IURL, urlScheme string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultApplicationAtURL:toOpenURLsWithScheme:completionHandler:"), applicationURL, objc.String(urlScheme), completionHandler)
}


// Sets the desktop image for the given screen to the image at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDesktopImageURL(_:for:options:)
func (w_ Workspace) SetDesktopImageURLForScreenOptionsError(url foundation.IURL, screen IScreen, options unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("setDesktopImageURL:forScreen:options:error:"), url, screen, options, error_)
	return rv
}


// Sets the icon for the file or directory at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/setIcon(_:forFile:options:)
func (w_ Workspace) SetIconForFileOptions(image IImage, fullPath string, options WorkspaceIconCreationOptions) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("setIcon:forFile:options:"), image, objc.String(fullPath), options)
	return rv
}


// Displays a Spotlight search results window in Finder for the specified query string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/showSearchResults(forQueryString:)
func (w_ Workspace) ShowSearchResultsForQueryString(queryString string) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("showSearchResultsForQueryString:"), objc.String(queryString))
	return rv
}


// Animates a sliding image from one point to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/slideImage:from:to:
func (w_ Workspace) SlideImageFromTo(image IImage, fromPoint coregraphics.CGPoint, toPoint coregraphics.CGPoint) {
	objc.Send[objc.ID](w_.ID, objc.Sel("slideImage:from:to:"), image, fromPoint, toPoint)
}


// Returns a Boolean indicating that the first Uniform Type Identifier (UTI) conforms to the second UTI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/type(_:conformsToType:)
func (w_ Workspace) TypeConformsToType(firstTypeName string, secondTypeName string) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("type:conformsToType:"), objc.String(firstTypeName), objc.String(secondTypeName))
	return rv
}


// Returns the uniform type identifier of the specified file, if it can be determined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/type(ofFile:)
func (w_ Workspace) TypeOfFileError(absoluteFilePath string, outError unsafe.Pointer) foundation.String {
	rv := objc.Send[foundation.String](w_.ID, objc.Sel("typeOfFile:error:"), objc.String(absoluteFilePath), outError)
	return rv
}


// Attempts to eject the volume mounted at the given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/unmountAndEjectDevice(at:)
func (w_ Workspace) UnmountAndEjectDeviceAtURLError(url foundation.IURL, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("unmountAndEjectDeviceAtURL:error:"), url, error_)
	return rv
}


// Unmounts and ejects the device at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/unmountAndEjectDevice(atPath:)
func (w_ Workspace) UnmountAndEjectDeviceAtPath(path string) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("unmountAndEjectDeviceAtPath:"), objc.String(path))
	return rv
}


// Returns the URL to the default app to open the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlForApplication(toOpen:)-7qkzf
func (w_ Workspace) URLForApplicationToOpenURL(url foundation.IURL) foundation.URL {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("URLForApplicationToOpenURL:"), url)
	return rv
}


// Returns the URL to the default app to open the specified content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlForApplication(toOpen:)-95cvp
func (w_ Workspace) URLForApplicationToOpenContentType(contentType unsafe.Pointer) foundation.URL {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("URLForApplicationToOpenContentType:"), contentType)
	return rv
}


// Returns the URL to the default app with the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlForApplication(withBundleIdentifier:)
func (w_ Workspace) URLForApplicationWithBundleIdentifier(bundleIdentifier string) foundation.URL {
	rv := objc.Send[foundation.URL](w_.ID, objc.Sel("URLForApplicationWithBundleIdentifier:"), objc.String(bundleIdentifier))
	return rv
}


// Returns an array of URLs to all available applications that can open the specified content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlsForApplications(toOpen:)-60rkm
func (w_ Workspace) URLsForApplicationsToOpenContentType(contentType unsafe.Pointer) []foundation.URL {
	rv := objc.Send[[]foundation.URL](w_.ID, objc.Sel("URLsForApplicationsToOpenContentType:"), contentType)
	return rv
}


// Returns an array of URLs to all available applications that can open the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlsForApplications(toOpen:)-ualk
func (w_ Workspace) URLsForApplicationsToOpenURL(url foundation.IURL) []foundation.URL {
	rv := objc.Send[[]foundation.URL](w_.ID, objc.Sel("URLsForApplicationsToOpenURL:"), url)
	return rv
}


// Returns an array of URLs to all available applications that can open the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlsForApplications(withBundleIdentifier:)
func (w_ Workspace) URLsForApplicationsWithBundleIdentifier(bundleIdentifier string) []foundation.URL {
	rv := objc.Send[[]foundation.URL](w_.ID, objc.Sel("URLsForApplicationsWithBundleIdentifier:"), objc.String(bundleIdentifier))
	return rv
}


// Returns a Boolean value indicating whether a change to the defaults database has been registered with a message since the last message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/userDefaultsChanged
func (w_ Workspace) UserDefaultsChanged() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("userDefaultsChanged"))
	return rv
}


// A Boolean value that indicates whether the app avoids conveying information through color alone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldDifferentiateWithoutColor
func (w_ Workspace) AccessibilityDisplayShouldDifferentiateWithoutColor() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldDifferentiateWithoutColor"))
	return rv
}


// A Boolean value that indicates whether the app presents a high-contrast user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldIncreaseContrast
func (w_ Workspace) AccessibilityDisplayShouldIncreaseContrast() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldIncreaseContrast"))
	return rv
}


// A Boolean value that indicates whether the accessibility option to invert colors is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldInvertColors
func (w_ Workspace) AccessibilityDisplayShouldInvertColors() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldInvertColors"))
	return rv
}


// A Boolean value that indicates whether the accessibility option to reduce motion is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldReduceMotion
func (w_ Workspace) AccessibilityDisplayShouldReduceMotion() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldReduceMotion"))
	return rv
}


// A Boolean value that indicates whether the app avoids using semitransparent backgrounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldReduceTransparency
func (w_ Workspace) AccessibilityDisplayShouldReduceTransparency() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("accessibilityDisplayShouldReduceTransparency"))
	return rv
}


// The array of colors for the file labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/fileLabelColors
func (w_ Workspace) FileLabelColors() []Color {
	rv := objc.Send[[]Color](w_.ID, objc.Sel("fileLabelColors"))
	return rv
}


// The array of file labels, returned as strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/fileLabels
func (w_ Workspace) FileLabels() []string {
	rv := objc.Send[[]string](w_.ID, objc.Sel("fileLabels"))
	return rv
}


// Returns the frontmost app, which is the app that receives key events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/frontmostApplication
func (w_ Workspace) FrontmostApplication() NSRunningApplication {
	rv := objc.Send[NSRunningApplication](w_.ID, objc.Sel("frontmostApplication"))
	return rv
}


// A Boolean value that indicates whether Switch Control is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/isSwitchControlEnabled
func (w_ Workspace) SwitchControlEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("switchControlEnabled"))
	return rv
}


// A Boolean value that indicates whether VoiceOver is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/isVoiceOverEnabled
func (w_ Workspace) VoiceOverEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("voiceOverEnabled"))
	return rv
}


// Returns the app that owns the currently displayed menu bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/menuBarOwningApplication
func (w_ Workspace) MenuBarOwningApplication() NSRunningApplication {
	rv := objc.Send[NSRunningApplication](w_.ID, objc.Sel("menuBarOwningApplication"))
	return rv
}


// The notification center for workspace notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/notificationCenter
func (w_ Workspace) NotificationCenter() foundation.NotificationCenter {
	rv := objc.Send[foundation.NotificationCenter](w_.ID, objc.Sel("notificationCenter"))
	return rv
}


// Returns an array of running apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/runningApplications
func (w_ Workspace) RunningApplications() []RunningApplication {
	rv := objc.Send[[]RunningApplication](w_.ID, objc.Sel("runningApplications"))
	return rv
}


// The shared workspace object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/shared
func (w_ Workspace) SharedWorkspace() NSWorkspace {
	rv := objc.Send[NSWorkspace](w_.ID, objc.Sel("sharedWorkspace"))
	return rv
}


// A Boolean value that indicates whether Switch Control is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isswitchcontrolenabled
func (w_ Workspace) IsSwitchControlEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isSwitchControlEnabled"))
	return rv
}


// A Boolean value that indicates whether Switch Control is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isswitchcontrolenabled
func (w_ Workspace) SetIsSwitchControlEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsSwitchControlEnabled:"), value)
}


// A Boolean value that indicates whether VoiceOver is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isvoiceoverenabled
func (w_ Workspace) IsVoiceOverEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isVoiceOverEnabled"))
	return rv
}


// A Boolean value that indicates whether VoiceOver is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsworkspace/isvoiceoverenabled
func (w_ Workspace) SetIsVoiceOverEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsVoiceOverEnabled:"), value)
}



