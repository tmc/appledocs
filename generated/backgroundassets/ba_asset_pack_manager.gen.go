// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BAAssetPackManager] class.
var (
	BAAssetPackManagerClass     _BAAssetPackManagerClass
	BAAssetPackManagerClassOnce sync.Once
)

func getBAAssetPackManagerClass() _BAAssetPackManagerClass {
	BAAssetPackManagerClassOnce.Do(func() {
		BAAssetPackManagerClass = _BAAssetPackManagerClass{objc.GetClass("BAAssetPackManager")}
	})
	return BAAssetPackManagerClass
}

type _BAAssetPackManagerClass struct {
	class objc.Class
}

// An interface definition for the [BAAssetPackManager] class.
type IBAAssetPackManager interface {
	objectivec.IObject
	URLForPathError(path string, error_ unsafe.Pointer) unsafe.Pointer
	CheckForUpdatesWithCompletionHandler(completionHandler unsafe.Pointer)
	ContentsAtPathSearchingInAssetPackWithIdentifierOptionsError(path string, assetPackIdentifier string, options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	EnsureLocalAvailabilityOfAssetPackCompletionHandler(assetPack unsafe.Pointer, completionHandler unsafe.Pointer)
	FileDescriptorForPathSearchingInAssetPackWithIdentifierError(path string, assetPackIdentifier string, error_ unsafe.Pointer) unsafe.Pointer
	GetAllAssetPacksWithCompletionHandler(completionHandler unsafe.Pointer)
	GetAssetPackWithIdentifierCompletionHandler(assetPackIdentifier string, completionHandler unsafe.Pointer)
	GetStatusOfAssetPackWithIdentifierCompletionHandler(assetPackIdentifier string, completionHandler unsafe.Pointer)
	RemoveAssetPackWithIdentifierCompletionHandler(assetPackIdentifier string, completionHandler unsafe.Pointer)
}

// A class that manages asset packs.
//
// The first time that your code refers to the shared manager, Background Assets considers that your app is opting into automatic system management of your asset packs.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager
type BAAssetPackManager struct {
	objectivec.Object
}

// BAAssetPackManagerFrom constructs a [BAAssetPackManager] from an unsafe.Pointer.
//
// A class that manages asset packs.
func BAAssetPackManagerFrom(ptr unsafe.Pointer) BAAssetPackManager {
	return BAAssetPackManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BAAssetPackManagerClass) Alloc() BAAssetPackManager {
	rv := objc.Send[BAAssetPackManager](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BAAssetPackManagerClass) New() BAAssetPackManager {
	rv := objc.Send[BAAssetPackManager](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BAAssetPackManager) Init() BAAssetPackManager {
	rv := objc.Send[BAAssetPackManager](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BAAssetPackManager) Autorelease() BAAssetPackManager {
	rv := objc.Send[BAAssetPackManager](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBAAssetPackManager creates a new BAAssetPackManager instance.
func NewBAAssetPackManager() BAAssetPackManager {
	return getBAAssetPackManagerClass().New()
}


// The shared asset-pack manager.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager/sharedManager
func (bc _BAAssetPackManagerClass) SharedManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("sharedManager"))
	return rv
}
// Returns a URL for the specified relative path.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager/URLForPath:error:
func (b_ BAAssetPackManager) URLForPathError(path string, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("URLForPath:error:"), objc.String(path), error_)
	return rv
}

// Gets the latest asset-pack information from the server, updates outdated asset packs, and removes obsolete asset packs.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager/checkForUpdatesWithCompletionHandler:
func (b_ BAAssetPackManager) CheckForUpdatesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("checkForUpdatesWithCompletionHandler:"), completionHandler)
}

// Returns the contents of an asset file at the specified relative path.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager/contentsAtPath:searchingInAssetPackWithIdentifier:options:error:
func (b_ BAAssetPackManager) ContentsAtPathSearchingInAssetPackWithIdentifierOptionsError(path string, assetPackIdentifier string, options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("contentsAtPath:searchingInAssetPackWithIdentifier:options:error:"), objc.String(path), objc.String(assetPackIdentifier), options, error_)
	return rv
}

// Ensures that the specified asset pack be available locally.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager/ensureLocalAvailabilityOfAssetPack:completionHandler:
func (b_ BAAssetPackManager) EnsureLocalAvailabilityOfAssetPackCompletionHandler(assetPack unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("ensureLocalAvailabilityOfAssetPack:completionHandler:"), assetPack, completionHandler)
}

// Opens and returns a file descriptor for the asset file at the specified relative path.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager/fileDescriptorForPath:searchingInAssetPackWithIdentifier:error:
func (b_ BAAssetPackManager) FileDescriptorForPathSearchingInAssetPackWithIdentifierError(path string, assetPackIdentifier string, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("fileDescriptorForPath:searchingInAssetPackWithIdentifier:error:"), objc.String(path), objc.String(assetPackIdentifier), error_)
	return rv
}

// Gets the asset packs that are available to download.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager/getAllAssetPacksWithCompletionHandler:
func (b_ BAAssetPackManager) GetAllAssetPacksWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("getAllAssetPacksWithCompletionHandler:"), completionHandler)
}

// Gets the asset pack with the given identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager/getAssetPackWithIdentifier:completionHandler:
func (b_ BAAssetPackManager) GetAssetPackWithIdentifierCompletionHandler(assetPackIdentifier string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("getAssetPackWithIdentifier:completionHandler:"), objc.String(assetPackIdentifier), completionHandler)
}

// Gets the status of the asset pack with the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager/getStatusOfAssetPackWithIdentifier:completionHandler:
func (b_ BAAssetPackManager) GetStatusOfAssetPackWithIdentifierCompletionHandler(assetPackIdentifier string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("getStatusOfAssetPackWithIdentifier:completionHandler:"), objc.String(assetPackIdentifier), completionHandler)
}

// Removes the specified asset pack from the device.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager/removeAssetPackWithIdentifier:completionHandler:
func (b_ BAAssetPackManager) RemoveAssetPackWithIdentifierCompletionHandler(assetPackIdentifier string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("removeAssetPackWithIdentifier:completionHandler:"), objc.String(assetPackIdentifier), completionHandler)
}

// An object that receives notifications about events that occur as an asset pack is downloaded.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager/delegate
func (b_ BAAssetPackManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// An object that receives notifications about events that occur as an asset pack is downloaded.

//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager/delegate
func (b_ BAAssetPackManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDelegate:"), value)
}
// The shared asset-pack manager.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BAAssetPackManager/sharedManager
func (b_ BAAssetPackManager) SharedManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("sharedManager"))
	return rv
}



