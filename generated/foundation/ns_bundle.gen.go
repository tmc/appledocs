// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Bundle] class.
var (
	BundleClass     _BundleClass
	BundleClassOnce sync.Once
)

func getBundleClass() _BundleClass {
	BundleClassOnce.Do(func() {
		BundleClass = _BundleClass{objc.GetClass("NSBundle")}
	})
	return BundleClass
}

type _BundleClass struct {
	class objc.Class
}

// An interface definition for the [Bundle] class.
type IBundle interface {
	objectivec.IObject
	ClassNamed(className string) objc.Class
	ContextHelpForKey(key unsafe.Pointer) unsafe.Pointer
	ImageForResource(name unsafe.Pointer) unsafe.Pointer
	Load() bool
	LoadAndReturnError(error_ unsafe.Pointer) bool
	LoadAppleScriptObjectiveCScripts()
	LoadNibNamedOwnerOptions(name string, owner objc.ID, options unsafe.Pointer) unsafe.Pointer
	LoadNibNamedOwnerTopLevelObjects(nibName unsafe.Pointer, owner objc.ID, topLevelObjects objc.ID) bool
	LocalizedStringForKeyValueTable(key string, value string, tableName string) string
	ObjectForInfoDictionaryKey(key string) objc.ID
	PathForAuxiliaryExecutable(executableName string) string
	PathForResourceOfType(name string, ext string) string
	PathForResourceOfTypeInDirectory(name string, ext string, subpath string) string
	PathForResourceOfTypeInDirectoryForLocalization(name string, ext string, subpath string, localizationName string) string
	PathForSoundResource(name unsafe.Pointer) string
	PathForImageResource(name unsafe.Pointer) string
	PathsForResourcesOfTypeInDirectory(ext string, subpath string) []string
	PathsForResourcesOfTypeInDirectoryForLocalization(ext string, subpath string, localizationName string) []string
	PreflightAndReturnError(error_ unsafe.Pointer) bool
	PreservationPriorityForTag(tag string) unsafe.Pointer
	SetPreservationPriorityForTags(priority unsafe.Pointer, tags unsafe.Pointer)
	Unload() bool
	URLForAuxiliaryExecutable(executableName string) unsafe.Pointer
	URLForResourceWithExtension(name string, ext string) unsafe.Pointer
	URLForResourceWithExtensionSubdirectory(name string, ext string, subpath string) unsafe.Pointer
	URLForResourceWithExtensionSubdirectoryLocalization(name string, ext string, subpath string, localizationName string) unsafe.Pointer
	URLForImageResource(name unsafe.Pointer) unsafe.Pointer
	URLsForResourcesWithExtensionSubdirectory(ext string, subpath string) []URL
	URLsForResourcesWithExtensionSubdirectoryLocalization(ext string, subpath string, localizationName string) []URL
	LoadNibFileExternalNameTableWithZone(fileName string, context objc.ID, zone unsafe.Pointer) bool
	LocalizedAttributedStringForKeyValueTable(key string, value string, tableName string) unsafe.Pointer
	LocalizedStringForKeyValueTableLocalizations(key string, value string, tableName string, localizations unsafe.Pointer) string
}

// A representation of the code and resources stored in a bundle directory on disk.
//
// Apple uses bundles to represent apps, frameworks, plug-ins, and many other specific types of content. Bundles organize their contained resources into well-defined subdirectories, and bundle structures vary depending on the platform and the type of the bundle. By using a bundle object, you can access a bundle’s resources without knowing the structure of the bundle. The bundle object provides a single interface for locating items, taking into account the bundle structure, user preferences, available localizations, and other relevant factors. Any executable can use a bundle object to locate resources, either inside an app’s bundle or in a known bundle located elsewhere. You don’t use a bundle object to locate files in a container directory or in other parts of the file system. The general pattern for using a bundle object is as follows: Create a bundle object for the intended bundle directory. Use the methods of the bundle object to locate or load the needed resource. Use other system APIs to interact with the resource. Some types of frequently used resources can be located and opened without a bundle. For example, when loading images, you store images in asset catalogs and load them using the methods of or . Similarly, for string resources, you use to load individual strings instead of loading the entire file yourself.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle
type Bundle struct {
	objectivec.Object
}

// BundleFrom constructs a [Bundle] from an unsafe.Pointer.
//
// A representation of the code and resources stored in a bundle directory on disk.
func BundleFrom(ptr unsafe.Pointer) Bundle {
	return Bundle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BundleClass) Alloc() Bundle {
	rv := objc.Send[Bundle](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BundleClass) New() Bundle {
	rv := objc.Send[Bundle](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ Bundle) Init() Bundle {
	rv := objc.Send[Bundle](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ Bundle) Autorelease() Bundle {
	rv := objc.Send[Bundle](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBundle creates a new Bundle instance.
func NewBundle() Bundle {
	return getBundleClass().New()
}


// Returns the object with which the specified class is associated.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(for:)
func NewBundleForClass(aClass objc.Class) Bundle {
	rv := objc.Send[Bundle](objc.ID(getBundleClass().class), objc.Sel("bundleForClass:"), aClass)
	return rv
}

// Returns the instance that has the specified bundle identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(identifier:)
func NewBundleWithIdentifier(identifier string) Bundle {
	rv := objc.Send[Bundle](objc.ID(getBundleClass().class), objc.Sel("bundleWithIdentifier:"), objc.String(identifier))
	return rv
}

// Returns an object initialized to correspond to the specified directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(path:)
func NewBundleWithPath(path string) Bundle {
	instance := getBundleClass().Alloc()
	rv := objc.Send[Bundle](instance.ID, objc.Sel("initWithPath:"), objc.String(path))
	rv.Autorelease()
	return rv
}

// Returns an object initialized to correspond to the specified file URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(url:)
func NewBundleWithURL(url unsafe.Pointer) Bundle {
	instance := getBundleClass().Alloc()
	rv := objc.Send[Bundle](instance.ID, objc.Sel("initWithURL:"), url)
	rv.Autorelease()
	return rv
}


// Returns the object with which the specified class is associated.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(for:)
func (bc _BundleClass) BundleForClass(aClass objc.Class) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("bundleForClass:"), aClass)
	return rv
}

// Returns the instance that has the specified bundle identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(identifier:)
func (bc _BundleClass) BundleWithIdentifier(identifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("bundleWithIdentifier:"), objc.String(identifier))
	return rv
}

// Returns the full pathname for the resource file identified by the specified name and extension and residing in a given bundle directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/path(forResource:ofType:inDirectory:)-swift.type.method
func (bc _BundleClass) PathForResourceOfTypeInDirectory(name string, ext string, bundlePath string) string {
	rv := objc.Send[string](objc.ID(bc.class), objc.Sel("pathForResource:ofType:inDirectory:"), objc.String(name), objc.String(ext), objc.String(bundlePath))
	return rv
}

// Returns an array containing the pathnames for all bundle resources having the specified extension and residing in the bundle directory at the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/paths(forResourcesOfType:inDirectory:)-swift.type.method
func (bc _BundleClass) PathsForResourcesOfTypeInDirectory(ext string, bundlePath string) []string {
	rv := objc.Send[[]string](objc.ID(bc.class), objc.Sel("pathsForResourcesOfType:inDirectory:"), objc.String(ext), objc.String(bundlePath))
	return rv
}

// Returns one or more localizations from the specified list that a bundle object would use to locate resources for the current user.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/preferredLocalizations(from:)
func (bc _BundleClass) PreferredLocalizationsFromArray(localizationsArray unsafe.Pointer) []string {
	rv := objc.Send[[]string](objc.ID(bc.class), objc.Sel("preferredLocalizationsFromArray:"), localizationsArray)
	return rv
}

// Returns locale identifiers for which a bundle would provide localized content, given a specified list of candidates for a user’s language preferences.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/preferredLocalizations(from:forPreferences:)
func (bc _BundleClass) PreferredLocalizationsFromArrayForPreferences(localizationsArray unsafe.Pointer, preferencesArray unsafe.Pointer) []string {
	rv := objc.Send[[]string](objc.ID(bc.class), objc.Sel("preferredLocalizationsFromArray:forPreferences:"), localizationsArray, preferencesArray)
	return rv
}

// Creates and returns a file URL for the resource with the specified name and extension in the specified bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/url(forResource:withExtension:subdirectory:in:)
func (bc _BundleClass) URLForResourceWithExtensionSubdirectoryInBundleWithURL(name string, ext string, subpath string, bundleURL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("URLForResource:withExtension:subdirectory:inBundleWithURL:"), objc.String(name), objc.String(ext), objc.String(subpath), bundleURL)
	return rv
}

// Returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, within the specified bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/urls(forResourcesWithExtension:subdirectory:in:)
func (bc _BundleClass) URLsForResourcesWithExtensionSubdirectoryInBundleWithURL(ext string, subpath string, bundleURL unsafe.Pointer) []URL {
	rv := objc.Send[[]URL](objc.ID(bc.class), objc.Sel("URLsForResourcesWithExtension:subdirectory:inBundleWithURL:"), objc.String(ext), objc.String(subpath), bundleURL)
	return rv
}

// Returns an object that corresponds to the specified directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundle/bundleWithPath:
func (bc _BundleClass) BundleWithPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("bundleWithPath:"), objc.String(path))
	return rv
}

// Returns an object that corresponds to the specified file URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundle/bundleWithURL:
func (bc _BundleClass) BundleWithURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("bundleWithURL:"), url)
	return rv
}

// Unarchives the contents of the nib file and links them to objects in your program.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundle/loadNibFile:externalNameTable:withZone:-c.type.method
func (bc _BundleClass) LoadNibFileExternalNameTableWithZone(fileName string, context objc.ID, zone unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(bc.class), objc.Sel("loadNibFile:externalNameTable:withZone:"), objc.String(fileName), context, zone)
	return rv
}

// Unarchives the contents of the nib file and links them to a specific owner object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundle/loadNibNamed:owner:
func (bc _BundleClass) LoadNibNamedOwner(nibName string, owner objc.ID) bool {
	rv := objc.Send[bool](objc.ID(bc.class), objc.Sel("loadNibNamed:owner:"), objc.String(nibName), owner)
	return rv
}

// Returns an array of all the application’s non-framework bundles.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/allBundles
func (bc _BundleClass) AllBundles() []Bundle {
	rv := objc.Send[[]Bundle](objc.ID(bc.class), objc.Sel("allBundles"))
	return rv
}
// Returns an array of all of the application’s bundles that represent frameworks.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/allFrameworks
func (bc _BundleClass) AllFrameworks() []Bundle {
	rv := objc.Send[[]Bundle](objc.ID(bc.class), objc.Sel("allFrameworks"))
	return rv
}
// Returns the bundle object that contains the current executable.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/main
func (bc _BundleClass) MainBundle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("mainBundle"))
	return rv
}
// Returns the object for the specified name.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/classNamed(_:)
func (b_ Bundle) ClassNamed(className string) objc.Class {
	rv := objc.Send[objc.Class](b_.ID, objc.Sel("classNamed:"), objc.String(className))
	return rv
}

// Returns the context-sensitive help for the specified key from the bundle’s help file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/contextHelp(forKey:)
func (b_ Bundle) ContextHelpForKey(key unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("contextHelpForKey:"), key)
	return rv
}

// Returns an instance associated with the specified name, which can be backed by multiple files representing different resolution versions of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/image(forResource:)
func (b_ Bundle) ImageForResource(name unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("imageForResource:"), name)
	return rv
}

// Dynamically loads the bundle’s executable code into a running program, if the code has not already been loaded.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/load()
func (b_ Bundle) Load() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("load"))
	return rv
}

// Loads the bundle’s executable code and returns any errors.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/loadAndReturnError()
func (b_ Bundle) LoadAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("loadAndReturnError:"), error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/loadAppleScriptObjectiveCScripts()
func (b_ Bundle) LoadAppleScriptObjectiveCScripts() {
	objc.Send[objc.ID](b_.ID, objc.Sel("loadAppleScriptObjectiveCScripts"))
}

// Unarchives the contents of a nib file located in the receiver’s bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/loadNibNamed(_:owner:options:)
func (b_ Bundle) LoadNibNamedOwnerOptions(name string, owner objc.ID, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("loadNibNamed:owner:options:"), objc.String(name), owner, options)
	return rv
}

// Loads a nib from the bundle with the specified file name and owner.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/loadNibNamed(_:owner:topLevelObjects:)
func (b_ Bundle) LoadNibNamedOwnerTopLevelObjects(nibName unsafe.Pointer, owner objc.ID, topLevelObjects objc.ID) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("loadNibNamed:owner:topLevelObjects:"), nibName, owner, topLevelObjects)
	return rv
}

// Returns a localized version of the string designated by the specified key and residing in the specified table.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/localizedString(forKey:value:table:)
func (b_ Bundle) LocalizedStringForKeyValueTable(key string, value string, tableName string) string {
	rv := objc.Send[string](b_.ID, objc.Sel("localizedStringForKey:value:table:"), objc.String(key), objc.String(value), objc.String(tableName))
	return rv
}

// Returns the value associated with the specified key in the receiver’s information property list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/object(forInfoDictionaryKey:)
func (b_ Bundle) ObjectForInfoDictionaryKey(key string) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("objectForInfoDictionaryKey:"), objc.String(key))
	return rv
}

// Returns the full pathname of the executable with the specified name in the receiver’s bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/path(forAuxiliaryExecutable:)
func (b_ Bundle) PathForAuxiliaryExecutable(executableName string) string {
	rv := objc.Send[string](b_.ID, objc.Sel("pathForAuxiliaryExecutable:"), objc.String(executableName))
	return rv
}

// Returns the full pathname for the resource identified by the specified name and file extension.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/path(forResource:ofType:)
func (b_ Bundle) PathForResourceOfType(name string, ext string) string {
	rv := objc.Send[string](b_.ID, objc.Sel("pathForResource:ofType:"), objc.String(name), objc.String(ext))
	return rv
}

// Returns the full pathname for the resource identified by the specified name and file extension and located in the specified bundle subdirectory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/path(forResource:ofType:inDirectory:)-swift.method
func (b_ Bundle) PathForResourceOfTypeInDirectory(name string, ext string, subpath string) string {
	rv := objc.Send[string](b_.ID, objc.Sel("pathForResource:ofType:inDirectory:"), objc.String(name), objc.String(ext), objc.String(subpath))
	return rv
}

// Returns the full pathname for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/path(forResource:ofType:inDirectory:forLocalization:)
func (b_ Bundle) PathForResourceOfTypeInDirectoryForLocalization(name string, ext string, subpath string, localizationName string) string {
	rv := objc.Send[string](b_.ID, objc.Sel("pathForResource:ofType:inDirectory:forLocalization:"), objc.String(name), objc.String(ext), objc.String(subpath), objc.String(localizationName))
	return rv
}

// Returns the location of the specified sound resource file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/path(forSoundResource:)
func (b_ Bundle) PathForSoundResource(name unsafe.Pointer) string {
	rv := objc.Send[string](b_.ID, objc.Sel("pathForSoundResource:"), name)
	return rv
}

// Returns the location of the specified image resource file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/pathForImageResource(_:)
func (b_ Bundle) PathForImageResource(name unsafe.Pointer) string {
	rv := objc.Send[string](b_.ID, objc.Sel("pathForImageResource:"), name)
	return rv
}

// Returns an array containing the pathnames for all bundle resources having the specified filename extension and residing in the resource subdirectory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/paths(forResourcesOfType:inDirectory:)-swift.method
func (b_ Bundle) PathsForResourcesOfTypeInDirectory(ext string, subpath string) []string {
	rv := objc.Send[[]string](b_.ID, objc.Sel("pathsForResourcesOfType:inDirectory:"), objc.String(ext), objc.String(subpath))
	return rv
}

// Returns an array containing the file for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/paths(forResourcesOfType:inDirectory:forLocalization:)
func (b_ Bundle) PathsForResourcesOfTypeInDirectoryForLocalization(ext string, subpath string, localizationName string) []string {
	rv := objc.Send[[]string](b_.ID, objc.Sel("pathsForResourcesOfType:inDirectory:forLocalization:"), objc.String(ext), objc.String(subpath), objc.String(localizationName))
	return rv
}

// Returns a Boolean value indicating whether the bundle’s executable code could be loaded successfully.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/preflight()
func (b_ Bundle) PreflightAndReturnError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("preflightAndReturnError:"), error_)
	return rv
}

// Returns the current preservation priority for the specified tag.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/preservationPriority(forTag:)
func (b_ Bundle) PreservationPriorityForTag(tag string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("preservationPriorityForTag:"), objc.String(tag))
	return rv
}

// A hint to the system of the relative order for purging tagged sets of resources in the bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/setPreservationPriority(_:forTags:)
func (b_ Bundle) SetPreservationPriorityForTags(priority unsafe.Pointer, tags unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPreservationPriority:forTags:"), priority, tags)
}

// Unloads the code associated with the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/unload()
func (b_ Bundle) Unload() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("unload"))
	return rv
}

// Returns the file URL of the executable with the specified name in the receiver’s bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/url(forAuxiliaryExecutable:)
func (b_ Bundle) URLForAuxiliaryExecutable(executableName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("URLForAuxiliaryExecutable:"), objc.String(executableName))
	return rv
}

// Returns the file URL for the resource identified by the specified name and file extension.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/url(forResource:withExtension:)
func (b_ Bundle) URLForResourceWithExtension(name string, ext string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("URLForResource:withExtension:"), objc.String(name), objc.String(ext))
	return rv
}

// Returns the file URL for the resource file identified by the specified name and extension and residing in a given bundle directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/url(forResource:withExtension:subdirectory:)
func (b_ Bundle) URLForResourceWithExtensionSubdirectory(name string, ext string, subpath string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("URLForResource:withExtension:subdirectory:"), objc.String(name), objc.String(ext), objc.String(subpath))
	return rv
}

// Returns the file URL for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/url(forResource:withExtension:subdirectory:localization:)
func (b_ Bundle) URLForResourceWithExtensionSubdirectoryLocalization(name string, ext string, subpath string, localizationName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("URLForResource:withExtension:subdirectory:localization:"), objc.String(name), objc.String(ext), objc.String(subpath), objc.String(localizationName))
	return rv
}

// Returns the location of the specified image resource as an NSURL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/urlForImageResource(_:)
func (b_ Bundle) URLForImageResource(name unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("URLForImageResource:"), name)
	return rv
}

// Returns an array of file URLs for all resources identified by the specified file extension and located in the specified bundle subdirectory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/urls(forResourcesWithExtension:subdirectory:)
func (b_ Bundle) URLsForResourcesWithExtensionSubdirectory(ext string, subpath string) []URL {
	rv := objc.Send[[]URL](b_.ID, objc.Sel("URLsForResourcesWithExtension:subdirectory:"), objc.String(ext), objc.String(subpath))
	return rv
}

// Returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/urls(forResourcesWithExtension:subdirectory:localization:)
func (b_ Bundle) URLsForResourcesWithExtensionSubdirectoryLocalization(ext string, subpath string, localizationName string) []URL {
	rv := objc.Send[[]URL](b_.ID, objc.Sel("URLsForResourcesWithExtension:subdirectory:localization:"), objc.String(ext), objc.String(subpath), objc.String(localizationName))
	return rv
}

// Unarchives the contents of a nib file located in the receiver’s bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundle/loadNibFile:externalNameTable:withZone:-c.method
func (b_ Bundle) LoadNibFileExternalNameTableWithZone(fileName string, context objc.ID, zone unsafe.Pointer) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("loadNibFile:externalNameTable:withZone:"), objc.String(fileName), context, zone)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundle/localizedAttributedStringForKey:value:table:
func (b_ Bundle) LocalizedAttributedStringForKeyValueTable(key string, value string, tableName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("localizedAttributedStringForKey:value:table:"), objc.String(key), objc.String(value), objc.String(tableName))
	return rv
}

// Look up a localized string given a list of available localizations.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundle/localizedStringForKey:value:table:localizations:
func (b_ Bundle) LocalizedStringForKeyValueTableLocalizations(key string, value string, tableName string, localizations unsafe.Pointer) string {
	rv := objc.Send[string](b_.ID, objc.Sel("localizedStringForKey:value:table:localizations:"), objc.String(key), objc.String(value), objc.String(tableName), localizations)
	return rv
}

// Returns an array of all the application’s non-framework bundles.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/allBundles
func (b_ Bundle) AllBundles() []Bundle {
	rv := objc.Send[[]Bundle](b_.ID, objc.Sel("allBundles"))
	return rv
}

// Returns an array of all of the application’s bundles that represent frameworks.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/allFrameworks
func (b_ Bundle) AllFrameworks() []Bundle {
	rv := objc.Send[[]Bundle](b_.ID, objc.Sel("allFrameworks"))
	return rv
}

// The file URL for the bundle’s App Store receipt.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/appStoreReceiptURL
func (b_ Bundle) AppStoreReceiptURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("appStoreReceiptURL"))
	return rv
}

// The full pathname of the receiver’s subdirectory containing plug-ins.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/builtInPlugInsPath
func (b_ Bundle) BuiltInPlugInsPath() string {
	rv := objc.Send[string](b_.ID, objc.Sel("builtInPlugInsPath"))
	return rv
}

// The file URL of the receiver’s subdirectory containing plug-ins.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/builtInPlugInsURL
func (b_ Bundle) BuiltInPlugInsURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("builtInPlugInsURL"))
	return rv
}

// The receiver’s bundle identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/bundleIdentifier
func (b_ Bundle) BundleIdentifier() string {
	rv := objc.Send[string](b_.ID, objc.Sel("bundleIdentifier"))
	return rv
}

// The full pathname of the receiver’s bundle directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/bundlePath
func (b_ Bundle) BundlePath() string {
	rv := objc.Send[string](b_.ID, objc.Sel("bundlePath"))
	return rv
}

// The full URL of the receiver’s bundle directory.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/bundleURL
func (b_ Bundle) BundleURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("bundleURL"))
	return rv
}

// The localization for the development language.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/developmentLocalization
func (b_ Bundle) DevelopmentLocalization() string {
	rv := objc.Send[string](b_.ID, objc.Sel("developmentLocalization"))
	return rv
}

// An array of numbers indicating the architecture types supported by the bundle’s executable.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/executableArchitectures
func (b_ Bundle) ExecutableArchitectures() []Number {
	rv := objc.Send[[]Number](b_.ID, objc.Sel("executableArchitectures"))
	return rv
}

// The full pathname of the receiver’s executable file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/executablePath
func (b_ Bundle) ExecutablePath() string {
	rv := objc.Send[string](b_.ID, objc.Sel("executablePath"))
	return rv
}

// The file URL of the receiver’s executable file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/executableURL
func (b_ Bundle) ExecutableURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("executableURL"))
	return rv
}

// A dictionary, constructed from the bundle’s file, that contains information about the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/infoDictionary
func (b_ Bundle) InfoDictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("infoDictionary"))
	return rv
}

// The load status of a bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/isLoaded
func (b_ Bundle) Loaded() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("loaded"))
	return rv
}

// A list of all the localizations contained in the bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/localizations
func (b_ Bundle) Localizations() []string {
	rv := objc.Send[[]string](b_.ID, objc.Sel("localizations"))
	return rv
}

// A dictionary with the keys from the bundle’s localized property list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/localizedInfoDictionary
func (b_ Bundle) LocalizedInfoDictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("localizedInfoDictionary"))
	return rv
}

// Returns the bundle object that contains the current executable.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/main
func (b_ Bundle) MainBundle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("mainBundle"))
	return rv
}

// An ordered list of preferred localizations contained in the bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/preferredLocalizations
func (b_ Bundle) PreferredLocalizations() []string {
	rv := objc.Send[[]string](b_.ID, objc.Sel("preferredLocalizations"))
	return rv
}

// The bundle’s principal class.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/principalClass
func (b_ Bundle) PrincipalClass() objc.Class {
	rv := objc.Send[objc.Class](b_.ID, objc.Sel("principalClass"))
	return rv
}

// The full pathname of the bundle’s subdirectory containing private frameworks.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/privateFrameworksPath
func (b_ Bundle) PrivateFrameworksPath() string {
	rv := objc.Send[string](b_.ID, objc.Sel("privateFrameworksPath"))
	return rv
}

// The file URL of the bundle’s subdirectory containing private frameworks.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/privateFrameworksURL
func (b_ Bundle) PrivateFrameworksURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("privateFrameworksURL"))
	return rv
}

// The full pathname of the bundle’s subdirectory containing resources.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/resourcePath
func (b_ Bundle) ResourcePath() string {
	rv := objc.Send[string](b_.ID, objc.Sel("resourcePath"))
	return rv
}

// The file URL of the bundle’s subdirectory containing resource files.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/resourceURL
func (b_ Bundle) ResourceURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("resourceURL"))
	return rv
}

// The full pathname of the bundle’s subdirectory containing shared frameworks.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/sharedFrameworksPath
func (b_ Bundle) SharedFrameworksPath() string {
	rv := objc.Send[string](b_.ID, objc.Sel("sharedFrameworksPath"))
	return rv
}

// The file URL of the receiver’s subdirectory containing shared frameworks.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/sharedFrameworksURL
func (b_ Bundle) SharedFrameworksURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("sharedFrameworksURL"))
	return rv
}

// The full pathname of the bundle’s subdirectory containing shared support files.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/sharedSupportPath
func (b_ Bundle) SharedSupportPath() string {
	rv := objc.Send[string](b_.ID, objc.Sel("sharedSupportPath"))
	return rv
}

// The file URL of the bundle’s subdirectory containing shared support files.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/sharedSupportURL
func (b_ Bundle) SharedSupportURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("sharedSupportURL"))
	return rv
}


