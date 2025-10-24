// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSBundle */


/* debug [class_header]: Header for NSBundle */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Bundle */
// An interface definition for the [Bundle] class.
type IBundle interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Bundle */
	// properties:
	AppStoreReceiptURL() IURL
	BuiltInPlugInsPath() IString
	BuiltInPlugInsURL() IURL
	BundleIdentifier() IString
	BundlePath() IString
	BundleURL() IURL
	DevelopmentLocalization() IString
	ExecutableArchitectures() []Number
	ExecutablePath() IString
	ExecutableURL() IURL
	InfoDictionary() IDictionary
	Loaded() bool
	Localizations() []string
	LocalizedInfoDictionary() IDictionary
	PreferredLocalizations() []string
	PrincipalClass() objc.Class
	PrivateFrameworksPath() IString
	PrivateFrameworksURL() IURL
	ResourcePath() IString
	ResourceURL() IURL
	SharedFrameworksPath() IString
	SharedFrameworksURL() IURL
	SharedSupportPath() IString
	SharedSupportURL() IURL
	IsLoaded() bool
	SetIsLoaded(value bool)
	NSExecutableArchitectureMismatchError() int
	SetNSExecutableArchitectureMismatchError(value int)
	NSExecutableErrorMaximum() int
	SetNSExecutableErrorMaximum(value int)
	NSExecutableErrorMinimum() int
	SetNSExecutableErrorMinimum(value int)
	NSExecutableLinkError() int
	SetNSExecutableLinkError(value int)
	NSExecutableLoadError() int
	SetNSExecutableLoadError(value int)
	NSExecutableNotLoadableError() int
	SetNSExecutableNotLoadableError(value int)
	NSExecutableRuntimeMismatchError() int
	SetNSExecutableRuntimeMismatchError(value int)
	NSLoadedClasses() IString
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Bundle */
	// methods:
	ClassNamed(className IString) objc.Class
	ContextHelpForKey(key HelpManagerContextHelpKey /* not a class type */) IAttributedString
	ImageForResource(name ImageName /* not a class type */) objectivec.IObject
	LoadAndReturnError(error_ IError) bool
	LoadAppleScriptObjectiveCScripts()
	LoadNibNamedOwnerTopLevelObjects(nibName NibName /* not a class type */, owner objc.IObject, topLevelObjects IArray) bool
	LocalizedStringForKeyValueTable(key IString, value IString, tableName IString) IString
	ObjectForInfoDictionaryKey(key IString) objc.ID
	PathForAuxiliaryExecutable(executableName IString) IString
	PathForResourceOfType(name IString, ext IString) IString
	PathForResourceOfTypeInDirectory(name IString, ext IString, subpath IString) IString
	PathForResourceOfTypeInDirectoryForLocalization(name IString, ext IString, subpath IString, localizationName IString) IString
	PathForSoundResource(name SoundName /* not a class type */) IString
	PathForImageResource(name ImageName /* not a class type */) IString
	PathsForResourcesOfTypeInDirectory(ext IString, subpath IString) []string
	PathsForResourcesOfTypeInDirectoryForLocalization(ext IString, subpath IString, localizationName IString) []string
	PreflightAndReturnError(error_ IError) bool
	Unload() bool
	URLForAuxiliaryExecutable(executableName IString) IURL
	URLForResourceWithExtension(name IString, ext IString) IURL
	URLForResourceWithExtensionSubdirectory(name IString, ext IString, subpath IString) IURL
	URLForResourceWithExtensionSubdirectoryLocalization(name IString, ext IString, subpath IString, localizationName IString) IURL
	URLForImageResource(name ImageName /* not a class type */) IURL
	URLsForResourcesWithExtensionSubdirectory(ext IString, subpath IString) []URL
	URLsForResourcesWithExtensionSubdirectoryLocalization(ext IString, subpath IString, localizationName IString) []URL
	LocalizedAttributedStringForKeyValueTable(key IString, value IString, tableName IString) IAttributedString
	LocalizedStringForKeyValueTableLocalizations(key IString, value IString, tableName IString, localizations []string) IString
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Bundle */
// Alloc allocates a new instance without initialization.
func (bc _BundleClass) Alloc() Bundle {
	rv := objc.Send[Bundle](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Bundle */
// A representation of the code and resources stored in a bundle directory on disk.
//
// Apple uses bundles to represent apps, frameworks, plug-ins, and many other specific types of content. Bundles organize their contained resources into well-defined subdirectories, and bundle structures vary depending on the platform and the type of the bundle. By using a bundle object, you can access a bundle’s resources without knowing the structure of the bundle. The bundle object provides a single interface for locating items, taking into account the bundle structure, user preferences, available localizations, and other relevant factors. Any executable can use a bundle object to locate resources, either inside an app’s bundle or in a known bundle located elsewhere. You don’t use a bundle object to locate files in a container directory or in other parts of the file system. The general pattern for using a bundle object is as follows: Create a bundle object for the intended bundle directory. Use the methods of the bundle object to locate or load the needed resource. Use other system APIs to interact with the resource. Some types of frequently used resources can be located and opened without a bundle. For example, when loading images, you store images in asset catalogs and load them using the methods of or . Similarly, for string resources, you use to load individual strings instead of loading the entire file yourself.


// A representation of the code and resources stored in a bundle directory on disk.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Bundle */

// Returns the object with which the specified class is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(for:)
func NewBundleForClass(aClass objc.Class) Bundle {
	rv := objc.Send[Bundle](objc.ID(getBundleClass().class), objc.Sel("bundleForClass:"), aClass)
	return rv
}/* debug [class_init_methods/constructor]: NewBundleForClass */


// Returns the instance that has the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(identifier:)
func NewBundleWithIdentifier(identifier IString) Bundle {
	rv := objc.Send[Bundle](objc.ID(getBundleClass().class), objc.Sel("bundleWithIdentifier:"), identifier)
	return rv
}/* debug [class_init_methods/constructor]: NewBundleWithIdentifier */


// Returns an object initialized to correspond to the specified directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(path:)
func NewBundleWithPath(path IString) Bundle {
	instance := getBundleClass().Alloc()
	rv := objc.Send[Bundle](instance.ID, objc.Sel("initWithPath:"), path)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBundleWithPath */


// Returns an object initialized to correspond to the specified file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(url:)
func NewBundleWithURL(url IURL) Bundle {
	instance := getBundleClass().Alloc()
	rv := objc.Send[Bundle](instance.ID, objc.Sel("initWithURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBundleWithURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Bundle */

// Returns the object with which the specified class is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(for:)
func (bc _BundleClass) BundleForClass(aClass objc.Class) IBundle {
	rv := objc.Send[Bundle](objc.ID(bc.class), objc.Sel("bundleForClass:"), aClass)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BundleForClass) */


// Returns the instance that has the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(identifier:)
func (bc _BundleClass) BundleWithIdentifier(identifier IString) IBundle {
	rv := objc.Send[Bundle](objc.ID(bc.class), objc.Sel("bundleWithIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BundleWithIdentifier) */


// Returns the full pathname for the resource file identified by the specified name and extension and residing in a given bundle directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/path(forResource:ofType:inDirectory:)-swift.type.method
func (bc _BundleClass) PathForResourceOfTypeInDirectory(name IString, ext IString, bundlePath IString) IString {
	rv := objc.Send[String](objc.ID(bc.class), objc.Sel("pathForResource:ofType:inDirectory:"), name, ext, bundlePath)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PathForResourceOfTypeInDirectory) */


// Returns an array containing the pathnames for all bundle resources having the specified extension and residing in the bundle directory at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/paths(forResourcesOfType:inDirectory:)-swift.type.method
func (bc _BundleClass) PathsForResourcesOfTypeInDirectory(ext IString, bundlePath IString) []string {
	rv := objc.Send[[]string](objc.ID(bc.class), objc.Sel("pathsForResourcesOfType:inDirectory:"), ext, bundlePath)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PathsForResourcesOfTypeInDirectory) */


// Returns one or more localizations from the specified list that a bundle object would use to locate resources for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/preferredLocalizations(from:)
func (bc _BundleClass) PreferredLocalizationsFromArray(localizationsArray []string) []string {
	rv := objc.Send[[]string](objc.ID(bc.class), objc.Sel("preferredLocalizationsFromArray:"), localizationsArray)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PreferredLocalizationsFromArray) */


// Returns locale identifiers for which a bundle would provide localized content, given a specified list of candidates for a user’s language preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/preferredLocalizations(from:forPreferences:)
func (bc _BundleClass) PreferredLocalizationsFromArrayForPreferences(localizationsArray []string, preferencesArray []string) []string {
	rv := objc.Send[[]string](objc.ID(bc.class), objc.Sel("preferredLocalizationsFromArray:forPreferences:"), localizationsArray, preferencesArray)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PreferredLocalizationsFromArrayForPreferences) */


// Creates and returns a file URL for the resource with the specified name and extension in the specified bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/url(forResource:withExtension:subdirectory:in:)
func (bc _BundleClass) URLForResourceWithExtensionSubdirectoryInBundleWithURL(name IString, ext IString, subpath IString, bundleURL IURL) IURL {
	rv := objc.Send[URL](objc.ID(bc.class), objc.Sel("URLForResource:withExtension:subdirectory:inBundleWithURL:"), name, ext, subpath, bundleURL)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=URLForResourceWithExtensionSubdirectoryInBundleWithURL) */


// Returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, within the specified bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/urls(forResourcesWithExtension:subdirectory:in:)
func (bc _BundleClass) URLsForResourcesWithExtensionSubdirectoryInBundleWithURL(ext IString, subpath IString, bundleURL IURL) []URL {
	rv := objc.Send[[]URL](objc.ID(bc.class), objc.Sel("URLsForResourcesWithExtension:subdirectory:inBundleWithURL:"), ext, subpath, bundleURL)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=URLsForResourcesWithExtensionSubdirectoryInBundleWithURL) */


// Returns an object that corresponds to the specified directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundle/bundleWithPath:
func (bc _BundleClass) BundleWithPath(path IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("bundleWithPath:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BundleWithPath) */


// Returns an object that corresponds to the specified file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundle/bundleWithURL:
func (bc _BundleClass) BundleWithURL(url IURL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("bundleWithURL:"), url)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BundleWithURL) */


// Unarchives the contents of the nib file and links them to objects in your program.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundle/loadNibFile:externalNameTable:withZone:-c.type.method
func (bc _BundleClass) LoadNibFileExternalNameTableWithZone(fileName IString, context IDictionary, zone Zone /* not a class type */) bool {
	rv := objc.Send[bool](objc.ID(bc.class), objc.Sel("loadNibFile:externalNameTable:withZone:"), fileName, context, zone)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadNibFileExternalNameTableWithZone) */


// Unarchives the contents of the nib file and links them to a specific owner object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundle/loadNibNamed:owner:
func (bc _BundleClass) LoadNibNamedOwner(nibName IString, owner objc.IObject) bool {
	rv := objc.Send[bool](objc.ID(bc.class), objc.Sel("loadNibNamed:owner:"), nibName, owner)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadNibNamedOwner) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Bundle */

// Returns an array of all the application’s non-framework bundles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/allBundles
func (bc _BundleClass) AllBundles() []Bundle {
	rv := objc.Send[[]Bundle](objc.ID(bc.class), objc.Sel("allBundles"))
	return rv
}/* debug [class_properties_class/property]: allBundles */

// Returns an array of all of the application’s bundles that represent frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/allFrameworks
func (bc _BundleClass) AllFrameworks() []Bundle {
	rv := objc.Send[[]Bundle](objc.ID(bc.class), objc.Sel("allFrameworks"))
	return rv
}/* debug [class_properties_class/property]: allFrameworks */

// Returns the bundle object that contains the current executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/main
func (bc _BundleClass) MainBundle() Bundle {
	rv := objc.Send[Bundle](objc.ID(bc.class), objc.Sel("mainBundle"))
	return rv
}/* debug [class_properties_class/property]: mainBundle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Bundle */

// Returns the object for the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/classNamed(_:)
func (b_ Bundle) ClassNamed(className IString) objc.Class {
	rv := objc.Send[objc.Class](b_.ID, objc.Sel("classNamed:"), className)
	return rv
}/* debug [instance_methods/method]: ClassNamed */


// Returns the context-sensitive help for the specified key from the bundle’s help file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/contextHelp(forKey:)
func (b_ Bundle) ContextHelpForKey(key HelpManagerContextHelpKey /* not a class type */) IAttributedString {
	rv := objc.Send[AttributedString](b_.ID, objc.Sel("contextHelpForKey:"), key)
	return rv
}/* debug [instance_methods/method]: ContextHelpForKey */


// Returns an instance associated with the specified name, which can be backed by multiple files representing different resolution versions of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/image(forResource:)
func (b_ Bundle) ImageForResource(name ImageName /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](b_.ID, objc.Sel("imageForResource:"), name)
	return rv
}/* debug [instance_methods/method]: ImageForResource */


// Dynamically loads the bundle’s executable code into a running program, if the code has not already been loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/load()
func (b_ Bundle) Load() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("load"))
	return rv
}/* debug [instance_methods/method]: Load */


// Loads the bundle’s executable code and returns any errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/loadAndReturnError()
func (b_ Bundle) LoadAndReturnError(error_ IError) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("loadAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: LoadAndReturnError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/loadAppleScriptObjectiveCScripts()
func (b_ Bundle) LoadAppleScriptObjectiveCScripts() {
	objc.Send[objc.ID](b_.ID, objc.Sel("loadAppleScriptObjectiveCScripts"))
}/* debug [instance_methods/method]: LoadAppleScriptObjectiveCScripts */


// Loads a nib from the bundle with the specified file name and owner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/loadNibNamed(_:owner:topLevelObjects:)
func (b_ Bundle) LoadNibNamedOwnerTopLevelObjects(nibName NibName /* not a class type */, owner objc.IObject, topLevelObjects IArray) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("loadNibNamed:owner:topLevelObjects:"), nibName, owner, topLevelObjects)
	return rv
}/* debug [instance_methods/method]: LoadNibNamedOwnerTopLevelObjects */


// Returns a localized version of the string designated by the specified key and residing in the specified table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/localizedString(forKey:value:table:)
func (b_ Bundle) LocalizedStringForKeyValueTable(key IString, value IString, tableName IString) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("localizedStringForKey:value:table:"), key, value, tableName)
	return rv
}/* debug [instance_methods/method]: LocalizedStringForKeyValueTable */


// Returns the value associated with the specified key in the receiver’s information property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/object(forInfoDictionaryKey:)
func (b_ Bundle) ObjectForInfoDictionaryKey(key IString) objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("objectForInfoDictionaryKey:"), key)
	return rv
}/* debug [instance_methods/method]: ObjectForInfoDictionaryKey */


// Returns the full pathname of the executable with the specified name in the receiver’s bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/path(forAuxiliaryExecutable:)
func (b_ Bundle) PathForAuxiliaryExecutable(executableName IString) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("pathForAuxiliaryExecutable:"), executableName)
	return rv
}/* debug [instance_methods/method]: PathForAuxiliaryExecutable */


// Returns the full pathname for the resource identified by the specified name and file extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/path(forResource:ofType:)
func (b_ Bundle) PathForResourceOfType(name IString, ext IString) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("pathForResource:ofType:"), name, ext)
	return rv
}/* debug [instance_methods/method]: PathForResourceOfType */


// Returns the full pathname for the resource identified by the specified name and file extension and located in the specified bundle subdirectory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/path(forResource:ofType:inDirectory:)-swift.method
func (b_ Bundle) PathForResourceOfTypeInDirectory(name IString, ext IString, subpath IString) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("pathForResource:ofType:inDirectory:"), name, ext, subpath)
	return rv
}/* debug [instance_methods/method]: PathForResourceOfTypeInDirectory */


// Returns the full pathname for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/path(forResource:ofType:inDirectory:forLocalization:)
func (b_ Bundle) PathForResourceOfTypeInDirectoryForLocalization(name IString, ext IString, subpath IString, localizationName IString) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("pathForResource:ofType:inDirectory:forLocalization:"), name, ext, subpath, localizationName)
	return rv
}/* debug [instance_methods/method]: PathForResourceOfTypeInDirectoryForLocalization */


// Returns the location of the specified sound resource file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/path(forSoundResource:)
func (b_ Bundle) PathForSoundResource(name SoundName /* not a class type */) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("pathForSoundResource:"), name)
	return rv
}/* debug [instance_methods/method]: PathForSoundResource */


// Returns the location of the specified image resource file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/pathForImageResource(_:)
func (b_ Bundle) PathForImageResource(name ImageName /* not a class type */) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("pathForImageResource:"), name)
	return rv
}/* debug [instance_methods/method]: PathForImageResource */


// Returns an array containing the pathnames for all bundle resources having the specified filename extension and residing in the resource subdirectory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/paths(forResourcesOfType:inDirectory:)-swift.method
func (b_ Bundle) PathsForResourcesOfTypeInDirectory(ext IString, subpath IString) []string {
	rv := objc.Send[[]string](b_.ID, objc.Sel("pathsForResourcesOfType:inDirectory:"), ext, subpath)
	return rv
}/* debug [instance_methods/method]: PathsForResourcesOfTypeInDirectory */


// Returns an array containing the file for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/paths(forResourcesOfType:inDirectory:forLocalization:)
func (b_ Bundle) PathsForResourcesOfTypeInDirectoryForLocalization(ext IString, subpath IString, localizationName IString) []string {
	rv := objc.Send[[]string](b_.ID, objc.Sel("pathsForResourcesOfType:inDirectory:forLocalization:"), ext, subpath, localizationName)
	return rv
}/* debug [instance_methods/method]: PathsForResourcesOfTypeInDirectoryForLocalization */


// Returns a Boolean value indicating whether the bundle’s executable code could be loaded successfully.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/preflight()
func (b_ Bundle) PreflightAndReturnError(error_ IError) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("preflightAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: PreflightAndReturnError */


// Unloads the code associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/unload()
func (b_ Bundle) Unload() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("unload"))
	return rv
}/* debug [instance_methods/method]: Unload */


// Returns the file URL of the executable with the specified name in the receiver’s bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/url(forAuxiliaryExecutable:)
func (b_ Bundle) URLForAuxiliaryExecutable(executableName IString) IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("URLForAuxiliaryExecutable:"), executableName)
	return rv
}/* debug [instance_methods/method]: URLForAuxiliaryExecutable */


// Returns the file URL for the resource identified by the specified name and file extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/url(forResource:withExtension:)
func (b_ Bundle) URLForResourceWithExtension(name IString, ext IString) IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("URLForResource:withExtension:"), name, ext)
	return rv
}/* debug [instance_methods/method]: URLForResourceWithExtension */


// Returns the file URL for the resource file identified by the specified name and extension and residing in a given bundle directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/url(forResource:withExtension:subdirectory:)
func (b_ Bundle) URLForResourceWithExtensionSubdirectory(name IString, ext IString, subpath IString) IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("URLForResource:withExtension:subdirectory:"), name, ext, subpath)
	return rv
}/* debug [instance_methods/method]: URLForResourceWithExtensionSubdirectory */


// Returns the file URL for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/url(forResource:withExtension:subdirectory:localization:)
func (b_ Bundle) URLForResourceWithExtensionSubdirectoryLocalization(name IString, ext IString, subpath IString, localizationName IString) IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("URLForResource:withExtension:subdirectory:localization:"), name, ext, subpath, localizationName)
	return rv
}/* debug [instance_methods/method]: URLForResourceWithExtensionSubdirectoryLocalization */


// Returns the location of the specified image resource as an NSURL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/urlForImageResource(_:)
func (b_ Bundle) URLForImageResource(name ImageName /* not a class type */) IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("URLForImageResource:"), name)
	return rv
}/* debug [instance_methods/method]: URLForImageResource */


// Returns an array of file URLs for all resources identified by the specified file extension and located in the specified bundle subdirectory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/urls(forResourcesWithExtension:subdirectory:)
func (b_ Bundle) URLsForResourcesWithExtensionSubdirectory(ext IString, subpath IString) []URL {
	rv := objc.Send[[]URL](b_.ID, objc.Sel("URLsForResourcesWithExtension:subdirectory:"), ext, subpath)
	return rv
}/* debug [instance_methods/method]: URLsForResourcesWithExtensionSubdirectory */


// Returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/urls(forResourcesWithExtension:subdirectory:localization:)
func (b_ Bundle) URLsForResourcesWithExtensionSubdirectoryLocalization(ext IString, subpath IString, localizationName IString) []URL {
	rv := objc.Send[[]URL](b_.ID, objc.Sel("URLsForResourcesWithExtension:subdirectory:localization:"), ext, subpath, localizationName)
	return rv
}/* debug [instance_methods/method]: URLsForResourcesWithExtensionSubdirectoryLocalization */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundle/localizedAttributedStringForKey:value:table:
func (b_ Bundle) LocalizedAttributedStringForKeyValueTable(key IString, value IString, tableName IString) IAttributedString {
	rv := objc.Send[AttributedString](b_.ID, objc.Sel("localizedAttributedStringForKey:value:table:"), key, value, tableName)
	return rv
}/* debug [instance_methods/method]: LocalizedAttributedStringForKeyValueTable */


// Look up a localized string given a list of available localizations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundle/localizedStringForKey:value:table:localizations:
func (b_ Bundle) LocalizedStringForKeyValueTableLocalizations(key IString, value IString, tableName IString, localizations []string) IString {
	rv := objc.Send[String](b_.ID, objc.Sel("localizedStringForKey:value:table:localizations:"), key, value, tableName, localizations)
	return rv
}/* debug [instance_methods/method]: LocalizedStringForKeyValueTableLocalizations */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Bundle */

// Returns an array of all the application’s non-framework bundles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/allBundles
func (b_ Bundle) AllBundles() []Bundle {
	rv := objc.Send[[]Bundle](b_.ID, objc.Sel("allBundles"))
	return rv
}/* debug [instance_properties/getter]: allBundles */


// Returns an array of all of the application’s bundles that represent frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/allFrameworks
func (b_ Bundle) AllFrameworks() []Bundle {
	rv := objc.Send[[]Bundle](b_.ID, objc.Sel("allFrameworks"))
	return rv
}/* debug [instance_properties/getter]: allFrameworks */


// The file URL for the bundle’s App Store receipt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/appStoreReceiptURL
func (b_ Bundle) AppStoreReceiptURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("appStoreReceiptURL"))
	return rv
}/* debug [instance_properties/getter]: appStoreReceiptURL */


// The full pathname of the receiver’s subdirectory containing plug-ins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/builtInPlugInsPath
func (b_ Bundle) BuiltInPlugInsPath() IString {
	rv := objc.Send[String](b_.ID, objc.Sel("builtInPlugInsPath"))
	return rv
}/* debug [instance_properties/getter]: builtInPlugInsPath */


// The file URL of the receiver’s subdirectory containing plug-ins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/builtInPlugInsURL
func (b_ Bundle) BuiltInPlugInsURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("builtInPlugInsURL"))
	return rv
}/* debug [instance_properties/getter]: builtInPlugInsURL */


// The receiver’s bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/bundleIdentifier
func (b_ Bundle) BundleIdentifier() IString {
	rv := objc.Send[String](b_.ID, objc.Sel("bundleIdentifier"))
	return rv
}/* debug [instance_properties/getter]: bundleIdentifier */


// The full pathname of the receiver’s bundle directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/bundlePath
func (b_ Bundle) BundlePath() IString {
	rv := objc.Send[String](b_.ID, objc.Sel("bundlePath"))
	return rv
}/* debug [instance_properties/getter]: bundlePath */


// The full URL of the receiver’s bundle directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/bundleURL
func (b_ Bundle) BundleURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("bundleURL"))
	return rv
}/* debug [instance_properties/getter]: bundleURL */


// The localization for the development language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/developmentLocalization
func (b_ Bundle) DevelopmentLocalization() IString {
	rv := objc.Send[String](b_.ID, objc.Sel("developmentLocalization"))
	return rv
}/* debug [instance_properties/getter]: developmentLocalization */


// An array of numbers indicating the architecture types supported by the bundle’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/executableArchitectures
func (b_ Bundle) ExecutableArchitectures() []Number {
	rv := objc.Send[[]Number](b_.ID, objc.Sel("executableArchitectures"))
	return rv
}/* debug [instance_properties/getter]: executableArchitectures */


// The full pathname of the receiver’s executable file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/executablePath
func (b_ Bundle) ExecutablePath() IString {
	rv := objc.Send[String](b_.ID, objc.Sel("executablePath"))
	return rv
}/* debug [instance_properties/getter]: executablePath */


// The file URL of the receiver’s executable file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/executableURL
func (b_ Bundle) ExecutableURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("executableURL"))
	return rv
}/* debug [instance_properties/getter]: executableURL */


// A dictionary, constructed from the bundle’s file, that contains information about the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/infoDictionary
func (b_ Bundle) InfoDictionary() IDictionary {
	rv := objc.Send[Dictionary](b_.ID, objc.Sel("infoDictionary"))
	return rv
}/* debug [instance_properties/getter]: infoDictionary */


// The load status of a bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/isLoaded
func (b_ Bundle) Loaded() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("loaded"))
	return rv
}/* debug [instance_properties/getter]: loaded */


// A list of all the localizations contained in the bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/localizations
func (b_ Bundle) Localizations() []string {
	rv := objc.Send[[]string](b_.ID, objc.Sel("localizations"))
	return rv
}/* debug [instance_properties/getter]: localizations */


// A dictionary with the keys from the bundle’s localized property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/localizedInfoDictionary
func (b_ Bundle) LocalizedInfoDictionary() IDictionary {
	rv := objc.Send[Dictionary](b_.ID, objc.Sel("localizedInfoDictionary"))
	return rv
}/* debug [instance_properties/getter]: localizedInfoDictionary */


// Returns the bundle object that contains the current executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/main
func (b_ Bundle) MainBundle() IBundle {
	rv := objc.Send[Bundle](b_.ID, objc.Sel("mainBundle"))
	return rv
}/* debug [instance_properties/getter]: mainBundle */


// An ordered list of preferred localizations contained in the bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/preferredLocalizations
func (b_ Bundle) PreferredLocalizations() []string {
	rv := objc.Send[[]string](b_.ID, objc.Sel("preferredLocalizations"))
	return rv
}/* debug [instance_properties/getter]: preferredLocalizations */


// The bundle’s principal class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/principalClass
func (b_ Bundle) PrincipalClass() objc.Class {
	rv := objc.Send[objc.Class](b_.ID, objc.Sel("principalClass"))
	return rv
}/* debug [instance_properties/getter]: principalClass */


// The full pathname of the bundle’s subdirectory containing private frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/privateFrameworksPath
func (b_ Bundle) PrivateFrameworksPath() IString {
	rv := objc.Send[String](b_.ID, objc.Sel("privateFrameworksPath"))
	return rv
}/* debug [instance_properties/getter]: privateFrameworksPath */


// The file URL of the bundle’s subdirectory containing private frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/privateFrameworksURL
func (b_ Bundle) PrivateFrameworksURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("privateFrameworksURL"))
	return rv
}/* debug [instance_properties/getter]: privateFrameworksURL */


// The full pathname of the bundle’s subdirectory containing resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/resourcePath
func (b_ Bundle) ResourcePath() IString {
	rv := objc.Send[String](b_.ID, objc.Sel("resourcePath"))
	return rv
}/* debug [instance_properties/getter]: resourcePath */


// The file URL of the bundle’s subdirectory containing resource files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/resourceURL
func (b_ Bundle) ResourceURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("resourceURL"))
	return rv
}/* debug [instance_properties/getter]: resourceURL */


// The full pathname of the bundle’s subdirectory containing shared frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/sharedFrameworksPath
func (b_ Bundle) SharedFrameworksPath() IString {
	rv := objc.Send[String](b_.ID, objc.Sel("sharedFrameworksPath"))
	return rv
}/* debug [instance_properties/getter]: sharedFrameworksPath */


// The file URL of the receiver’s subdirectory containing shared frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/sharedFrameworksURL
func (b_ Bundle) SharedFrameworksURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("sharedFrameworksURL"))
	return rv
}/* debug [instance_properties/getter]: sharedFrameworksURL */


// The full pathname of the bundle’s subdirectory containing shared support files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/sharedSupportPath
func (b_ Bundle) SharedSupportPath() IString {
	rv := objc.Send[String](b_.ID, objc.Sel("sharedSupportPath"))
	return rv
}/* debug [instance_properties/getter]: sharedSupportPath */


// The file URL of the bundle’s subdirectory containing shared support files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/sharedSupportURL
func (b_ Bundle) SharedSupportURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("sharedSupportURL"))
	return rv
}/* debug [instance_properties/getter]: sharedSupportURL */


// The load status of a bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/isloaded
func (b_ Bundle) IsLoaded() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isLoaded"))
	return rv
}/* debug [instance_properties/getter]: isLoaded */


// The load status of a bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/isloaded
func (b_ Bundle) SetIsLoaded(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsLoaded:"), value)
}/* debug [instance_properties/setter]: isLoaded */


// The executable doesn’t provide an architecture compatible with the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutablearchitecturemismatcherror-swift.var
func (b_ Bundle) NSExecutableArchitectureMismatchError() int {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableArchitectureMismatchError"))
	return rv
}/* debug [instance_properties/getter]: NSExecutableArchitectureMismatchError */


// The executable doesn’t provide an architecture compatible with the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutablearchitecturemismatcherror-swift.var
func (b_ Bundle) SetNSExecutableArchitectureMismatchError(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableArchitectureMismatchError:"), value)
}/* debug [instance_properties/setter]: NSExecutableArchitectureMismatchError */


// The end of the range of error codes reserved for errors related to executable files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableerrormaximum-swift.var
func (b_ Bundle) NSExecutableErrorMaximum() int {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableErrorMaximum"))
	return rv
}/* debug [instance_properties/getter]: NSExecutableErrorMaximum */


// The end of the range of error codes reserved for errors related to executable files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableerrormaximum-swift.var
func (b_ Bundle) SetNSExecutableErrorMaximum(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableErrorMaximum:"), value)
}/* debug [instance_properties/setter]: NSExecutableErrorMaximum */


// The beginning of the range of error codes reserved for errors related to executable files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableerrorminimum-swift.var
func (b_ Bundle) NSExecutableErrorMinimum() int {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableErrorMinimum"))
	return rv
}/* debug [instance_properties/getter]: NSExecutableErrorMinimum */


// The beginning of the range of error codes reserved for errors related to executable files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableerrorminimum-swift.var
func (b_ Bundle) SetNSExecutableErrorMinimum(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableErrorMinimum:"), value)
}/* debug [instance_properties/setter]: NSExecutableErrorMinimum */


// The executable failed due to linking issues.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutablelinkerror-swift.var
func (b_ Bundle) NSExecutableLinkError() int {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableLinkError"))
	return rv
}/* debug [instance_properties/getter]: NSExecutableLinkError */


// The executable failed due to linking issues.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutablelinkerror-swift.var
func (b_ Bundle) SetNSExecutableLinkError(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableLinkError:"), value)
}/* debug [instance_properties/setter]: NSExecutableLinkError */


// Executable cannot be loaded for an otherwise-unspecified reason.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableloaderror-swift.var
func (b_ Bundle) NSExecutableLoadError() int {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableLoadError"))
	return rv
}/* debug [instance_properties/getter]: NSExecutableLoadError */


// Executable cannot be loaded for an otherwise-unspecified reason.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableloaderror-swift.var
func (b_ Bundle) SetNSExecutableLoadError(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableLoadError:"), value)
}/* debug [instance_properties/setter]: NSExecutableLoadError */


// The executable type isn’t loadable in the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutablenotloadableerror-swift.var
func (b_ Bundle) NSExecutableNotLoadableError() int {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableNotLoadableError"))
	return rv
}/* debug [instance_properties/getter]: NSExecutableNotLoadableError */


// The executable type isn’t loadable in the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutablenotloadableerror-swift.var
func (b_ Bundle) SetNSExecutableNotLoadableError(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableNotLoadableError:"), value)
}/* debug [instance_properties/setter]: NSExecutableNotLoadableError */


// The executable has Objective-C runtime information that’s incompatible with the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableruntimemismatcherror-swift.var
func (b_ Bundle) NSExecutableRuntimeMismatchError() int {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableRuntimeMismatchError"))
	return rv
}/* debug [instance_properties/getter]: NSExecutableRuntimeMismatchError */


// The executable has Objective-C runtime information that’s incompatible with the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableruntimemismatcherror-swift.var
func (b_ Bundle) SetNSExecutableRuntimeMismatchError(value int) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableRuntimeMismatchError:"), value)
}/* debug [instance_properties/setter]: NSExecutableRuntimeMismatchError */


// A constant used as a key for the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsloadedclasses
func (b_ Bundle) NSLoadedClasses() IString {
	rv := objc.Send[String](b_.ID, objc.Sel("NSLoadedClasses"))
	return rv
}/* debug [instance_properties/getter]: NSLoadedClasses */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSBundle */


