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
	// properties:
	PrincipalClass() objc.Class
	AppStoreReceiptURL() IURL
	SetAppStoreReceiptURL(value IURL)
	BuiltInPlugInsPath() string /* primitive/slice/pointer */
	SetBuiltInPlugInsPath(value string /* primitive/slice/pointer */)
	BuiltInPlugInsURL() IURL
	SetBuiltInPlugInsURL(value IURL)
	BundleIdentifier() string /* primitive/slice/pointer */
	SetBundleIdentifier(value string /* primitive/slice/pointer */)
	BundlePath() string /* primitive/slice/pointer */
	SetBundlePath(value string /* primitive/slice/pointer */)
	BundleURL() IURL
	SetBundleURL(value IURL)
	DevelopmentLocalization() string /* primitive/slice/pointer */
	SetDevelopmentLocalization(value string /* primitive/slice/pointer */)
	ExecutableArchitectures() Number /* foo */
	SetExecutableArchitectures(value Number /* foo */)
	ExecutablePath() string /* primitive/slice/pointer */
	SetExecutablePath(value string /* primitive/slice/pointer */)
	ExecutableURL() IURL
	SetExecutableURL(value IURL)
	InfoDictionary() string /* primitive/slice/pointer */
	SetInfoDictionary(value string /* primitive/slice/pointer */)
	IsLoaded() bool /* primitive/slice/pointer */
	SetIsLoaded(value bool /* primitive/slice/pointer */)
	Localizations() string /* primitive/slice/pointer */
	SetLocalizations(value string /* primitive/slice/pointer */)
	LocalizedInfoDictionary() string /* primitive/slice/pointer */
	SetLocalizedInfoDictionary(value string /* primitive/slice/pointer */)
	PreferredLocalizations() string /* primitive/slice/pointer */
	SetPreferredLocalizations(value string /* primitive/slice/pointer */)
	PrivateFrameworksPath() string /* primitive/slice/pointer */
	SetPrivateFrameworksPath(value string /* primitive/slice/pointer */)
	PrivateFrameworksURL() IURL
	SetPrivateFrameworksURL(value IURL)
	ResourcePath() string /* primitive/slice/pointer */
	SetResourcePath(value string /* primitive/slice/pointer */)
	ResourceURL() IURL
	SetResourceURL(value IURL)
	SharedFrameworksPath() string /* primitive/slice/pointer */
	SetSharedFrameworksPath(value string /* primitive/slice/pointer */)
	SharedFrameworksURL() IURL
	SetSharedFrameworksURL(value IURL)
	SharedSupportPath() string /* primitive/slice/pointer */
	SetSharedSupportPath(value string /* primitive/slice/pointer */)
	SharedSupportURL() IURL
	SetSharedSupportURL(value IURL)
	NSExecutableArchitectureMismatchError() int /* primitive/slice/pointer */
	SetNSExecutableArchitectureMismatchError(value int /* primitive/slice/pointer */)
	NSExecutableErrorMaximum() int /* primitive/slice/pointer */
	SetNSExecutableErrorMaximum(value int /* primitive/slice/pointer */)
	NSExecutableErrorMinimum() int /* primitive/slice/pointer */
	SetNSExecutableErrorMinimum(value int /* primitive/slice/pointer */)
	NSExecutableLinkError() int /* primitive/slice/pointer */
	SetNSExecutableLinkError(value int /* primitive/slice/pointer */)
	NSExecutableLoadError() int /* primitive/slice/pointer */
	SetNSExecutableLoadError(value int /* primitive/slice/pointer */)
	NSExecutableNotLoadableError() int /* primitive/slice/pointer */
	SetNSExecutableNotLoadableError(value int /* primitive/slice/pointer */)
	NSExecutableRuntimeMismatchError() int /* primitive/slice/pointer */
	SetNSExecutableRuntimeMismatchError(value int /* primitive/slice/pointer */)
	NSLoadedClasses() string /* primitive/slice/pointer */
	// methods:
	LoadAndReturnError(error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	LocalizedStringForKeyValueTable(key string /* primitive/slice/pointer */, value string /* primitive/slice/pointer */, tableName string /* primitive/slice/pointer */) String /* foo */
	PathsForResourcesOfTypeInDirectory(ext string /* primitive/slice/pointer */, subpath string /* primitive/slice/pointer */) []string /* primitive/slice/pointer */
}

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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(for:)
func NewBundleForClass(aClass objc.Class) Bundle {
	rv := objc.Send[Bundle](objc.ID(getBundleClass().class), objc.Sel("bundleForClass:"), aClass)
	return rv
}



// Returns the object with which the specified class is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/init(for:)
func (bc _BundleClass) BundleForClass(aClass objc.Class) IBundle {
	rv := objc.Send[Bundle](objc.ID(bc.class), objc.Sel("bundleForClass:"), aClass)
	return rv
}


// Returns the bundle object that contains the current executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/main
func (bc _BundleClass) MainBundle() Bundle {
	rv := objc.Send[Bundle](objc.ID(bc.class), objc.Sel("mainBundle"))
	return rv
}

// Loads the bundle’s executable code and returns any errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/loadAndReturnError()
func (b_ Bundle) LoadAndReturnError(error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("loadAndReturnError:"), error_)
	return rv
}


// Returns a localized version of the string designated by the specified key and residing in the specified table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/localizedString(forKey:value:table:)
func (b_ Bundle) LocalizedStringForKeyValueTable(key string /* primitive/slice/pointer */, value string /* primitive/slice/pointer */, tableName string /* primitive/slice/pointer */) String /* foo */ {
	rv := objc.Send[String](b_.ID, objc.Sel("localizedStringForKey:value:table:"), objc.String(key), objc.String(value), objc.String(tableName))
	return rv
}


// Returns an array containing the pathnames for all bundle resources having the specified filename extension and residing in the resource subdirectory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/paths(forResourcesOfType:inDirectory:)-swift.method
func (b_ Bundle) PathsForResourcesOfTypeInDirectory(ext string /* primitive/slice/pointer */, subpath string /* primitive/slice/pointer */) []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](b_.ID, objc.Sel("pathsForResourcesOfType:inDirectory:"), objc.String(ext), objc.String(subpath))
	return rv
}


// Returns the bundle object that contains the current executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/main
func (b_ Bundle) MainBundle() IBundle {
	rv := objc.Send[Bundle](b_.ID, objc.Sel("mainBundle"))
	return rv
}


// The bundle’s principal class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/principalClass
func (b_ Bundle) PrincipalClass() objc.Class {
	rv := objc.Send[objc.Class](b_.ID, objc.Sel("principalClass"))
	return rv
}


// The file URL for the bundle’s App Store receipt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/appstorereceipturl
func (b_ Bundle) AppStoreReceiptURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("appStoreReceiptURL"))
	return rv
}


// The file URL for the bundle’s App Store receipt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/appstorereceipturl
func (b_ Bundle) SetAppStoreReceiptURL(value IURL) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAppStoreReceiptURL:"), value)
}


// The full pathname of the receiver’s subdirectory containing plug-ins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/builtinpluginspath
func (b_ Bundle) BuiltInPlugInsPath() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("builtInPlugInsPath"))
	return rv
}


// The full pathname of the receiver’s subdirectory containing plug-ins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/builtinpluginspath
func (b_ Bundle) SetBuiltInPlugInsPath(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBuiltInPlugInsPath:"), objc.String(value))
}


// The file URL of the receiver’s subdirectory containing plug-ins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/builtinpluginsurl
func (b_ Bundle) BuiltInPlugInsURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("builtInPlugInsURL"))
	return rv
}


// The file URL of the receiver’s subdirectory containing plug-ins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/builtinpluginsurl
func (b_ Bundle) SetBuiltInPlugInsURL(value IURL) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBuiltInPlugInsURL:"), value)
}


// The receiver’s bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/bundleidentifier
func (b_ Bundle) BundleIdentifier() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("bundleIdentifier"))
	return rv
}


// The receiver’s bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/bundleidentifier
func (b_ Bundle) SetBundleIdentifier(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBundleIdentifier:"), objc.String(value))
}


// The full pathname of the receiver’s bundle directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/bundlepath
func (b_ Bundle) BundlePath() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("bundlePath"))
	return rv
}


// The full pathname of the receiver’s bundle directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/bundlepath
func (b_ Bundle) SetBundlePath(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBundlePath:"), objc.String(value))
}


// The full URL of the receiver’s bundle directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/bundleurl
func (b_ Bundle) BundleURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("bundleURL"))
	return rv
}


// The full URL of the receiver’s bundle directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/bundleurl
func (b_ Bundle) SetBundleURL(value IURL) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBundleURL:"), value)
}


// The localization for the development language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/developmentlocalization
func (b_ Bundle) DevelopmentLocalization() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("developmentLocalization"))
	return rv
}


// The localization for the development language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/developmentlocalization
func (b_ Bundle) SetDevelopmentLocalization(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDevelopmentLocalization:"), objc.String(value))
}


// An array of numbers indicating the architecture types supported by the bundle’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/executablearchitectures
func (b_ Bundle) ExecutableArchitectures() Number /* foo */ {
	rv := objc.Send[Number](b_.ID, objc.Sel("executableArchitectures"))
	return rv
}


// An array of numbers indicating the architecture types supported by the bundle’s executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/executablearchitectures
func (b_ Bundle) SetExecutableArchitectures(value Number /* foo */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setExecutableArchitectures:"), value)
}


// The full pathname of the receiver’s executable file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/executablepath
func (b_ Bundle) ExecutablePath() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("executablePath"))
	return rv
}


// The full pathname of the receiver’s executable file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/executablepath
func (b_ Bundle) SetExecutablePath(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setExecutablePath:"), objc.String(value))
}


// The file URL of the receiver’s executable file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/executableurl
func (b_ Bundle) ExecutableURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("executableURL"))
	return rv
}


// The file URL of the receiver’s executable file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/executableurl
func (b_ Bundle) SetExecutableURL(value IURL) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setExecutableURL:"), value)
}


// A dictionary, constructed from the bundle’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/infodictionary
func (b_ Bundle) InfoDictionary() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("infoDictionary"))
	return rv
}


// A dictionary, constructed from the bundle’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/infodictionary
func (b_ Bundle) SetInfoDictionary(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setInfoDictionary:"), objc.String(value))
}


// The load status of a bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/isloaded
func (b_ Bundle) IsLoaded() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isLoaded"))
	return rv
}


// The load status of a bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/isloaded
func (b_ Bundle) SetIsLoaded(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsLoaded:"), value)
}


// A list of all the localizations contained in the bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/localizations
func (b_ Bundle) Localizations() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("localizations"))
	return rv
}


// A list of all the localizations contained in the bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/localizations
func (b_ Bundle) SetLocalizations(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLocalizations:"), objc.String(value))
}


// A dictionary with the keys from the bundle’s localized property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/localizedinfodictionary
func (b_ Bundle) LocalizedInfoDictionary() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("localizedInfoDictionary"))
	return rv
}


// A dictionary with the keys from the bundle’s localized property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/localizedinfodictionary
func (b_ Bundle) SetLocalizedInfoDictionary(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setLocalizedInfoDictionary:"), objc.String(value))
}


// An ordered list of preferred localizations contained in the bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/preferredlocalizations
func (b_ Bundle) PreferredLocalizations() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("preferredLocalizations"))
	return rv
}


// An ordered list of preferred localizations contained in the bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/preferredlocalizations
func (b_ Bundle) SetPreferredLocalizations(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPreferredLocalizations:"), objc.String(value))
}


// The full pathname of the bundle’s subdirectory containing private frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/privateframeworkspath
func (b_ Bundle) PrivateFrameworksPath() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("privateFrameworksPath"))
	return rv
}


// The full pathname of the bundle’s subdirectory containing private frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/privateframeworkspath
func (b_ Bundle) SetPrivateFrameworksPath(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrivateFrameworksPath:"), objc.String(value))
}


// The file URL of the bundle’s subdirectory containing private frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/privateframeworksurl
func (b_ Bundle) PrivateFrameworksURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("privateFrameworksURL"))
	return rv
}


// The file URL of the bundle’s subdirectory containing private frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/privateframeworksurl
func (b_ Bundle) SetPrivateFrameworksURL(value IURL) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setPrivateFrameworksURL:"), value)
}


// The full pathname of the bundle’s subdirectory containing resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/resourcepath
func (b_ Bundle) ResourcePath() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("resourcePath"))
	return rv
}


// The full pathname of the bundle’s subdirectory containing resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/resourcepath
func (b_ Bundle) SetResourcePath(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setResourcePath:"), objc.String(value))
}


// The file URL of the bundle’s subdirectory containing resource files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/resourceurl
func (b_ Bundle) ResourceURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("resourceURL"))
	return rv
}


// The file URL of the bundle’s subdirectory containing resource files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/resourceurl
func (b_ Bundle) SetResourceURL(value IURL) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setResourceURL:"), value)
}


// The full pathname of the bundle’s subdirectory containing shared frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/sharedframeworkspath
func (b_ Bundle) SharedFrameworksPath() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("sharedFrameworksPath"))
	return rv
}


// The full pathname of the bundle’s subdirectory containing shared frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/sharedframeworkspath
func (b_ Bundle) SetSharedFrameworksPath(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSharedFrameworksPath:"), objc.String(value))
}


// The file URL of the receiver’s subdirectory containing shared frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/sharedframeworksurl
func (b_ Bundle) SharedFrameworksURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("sharedFrameworksURL"))
	return rv
}


// The file URL of the receiver’s subdirectory containing shared frameworks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/sharedframeworksurl
func (b_ Bundle) SetSharedFrameworksURL(value IURL) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSharedFrameworksURL:"), value)
}


// The full pathname of the bundle’s subdirectory containing shared support files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/sharedsupportpath
func (b_ Bundle) SharedSupportPath() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("sharedSupportPath"))
	return rv
}


// The full pathname of the bundle’s subdirectory containing shared support files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/sharedsupportpath
func (b_ Bundle) SetSharedSupportPath(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSharedSupportPath:"), objc.String(value))
}


// The file URL of the bundle’s subdirectory containing shared support files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/sharedsupporturl
func (b_ Bundle) SharedSupportURL() IURL {
	rv := objc.Send[URL](b_.ID, objc.Sel("sharedSupportURL"))
	return rv
}


// The file URL of the bundle’s subdirectory containing shared support files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/bundle/sharedsupporturl
func (b_ Bundle) SetSharedSupportURL(value IURL) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSharedSupportURL:"), value)
}


// The executable doesn’t provide an architecture compatible with the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutablearchitecturemismatcherror-swift.var
func (b_ Bundle) NSExecutableArchitectureMismatchError() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableArchitectureMismatchError"))
	return rv
}


// The executable doesn’t provide an architecture compatible with the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutablearchitecturemismatcherror-swift.var
func (b_ Bundle) SetNSExecutableArchitectureMismatchError(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableArchitectureMismatchError:"), value)
}


// The end of the range of error codes reserved for errors related to executable files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableerrormaximum-swift.var
func (b_ Bundle) NSExecutableErrorMaximum() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableErrorMaximum"))
	return rv
}


// The end of the range of error codes reserved for errors related to executable files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableerrormaximum-swift.var
func (b_ Bundle) SetNSExecutableErrorMaximum(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableErrorMaximum:"), value)
}


// The beginning of the range of error codes reserved for errors related to executable files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableerrorminimum-swift.var
func (b_ Bundle) NSExecutableErrorMinimum() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableErrorMinimum"))
	return rv
}


// The beginning of the range of error codes reserved for errors related to executable files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableerrorminimum-swift.var
func (b_ Bundle) SetNSExecutableErrorMinimum(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableErrorMinimum:"), value)
}


// The executable failed due to linking issues.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutablelinkerror-swift.var
func (b_ Bundle) NSExecutableLinkError() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableLinkError"))
	return rv
}


// The executable failed due to linking issues.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutablelinkerror-swift.var
func (b_ Bundle) SetNSExecutableLinkError(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableLinkError:"), value)
}


// Executable cannot be loaded for an otherwise-unspecified reason.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableloaderror-swift.var
func (b_ Bundle) NSExecutableLoadError() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableLoadError"))
	return rv
}


// Executable cannot be loaded for an otherwise-unspecified reason.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableloaderror-swift.var
func (b_ Bundle) SetNSExecutableLoadError(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableLoadError:"), value)
}


// The executable type isn’t loadable in the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutablenotloadableerror-swift.var
func (b_ Bundle) NSExecutableNotLoadableError() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableNotLoadableError"))
	return rv
}


// The executable type isn’t loadable in the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutablenotloadableerror-swift.var
func (b_ Bundle) SetNSExecutableNotLoadableError(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableNotLoadableError:"), value)
}


// The executable has Objective-C runtime information that’s incompatible with the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableruntimemismatcherror-swift.var
func (b_ Bundle) NSExecutableRuntimeMismatchError() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](b_.ID, objc.Sel("NSExecutableRuntimeMismatchError"))
	return rv
}


// The executable has Objective-C runtime information that’s incompatible with the current process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsexecutableruntimemismatcherror-swift.var
func (b_ Bundle) SetNSExecutableRuntimeMismatchError(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setNSExecutableRuntimeMismatchError:"), value)
}


// A constant used as a key for the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsloadedclasses
func (b_ Bundle) NSLoadedClasses() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](b_.ID, objc.Sel("NSLoadedClasses"))
	return rv
}


