// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetDownloadStorageManager] class.
var (
	AssetDownloadStorageManagerClass     _AssetDownloadStorageManagerClass
	AssetDownloadStorageManagerClassOnce sync.Once
)

func getAssetDownloadStorageManagerClass() _AssetDownloadStorageManagerClass {
	AssetDownloadStorageManagerClassOnce.Do(func() {
		AssetDownloadStorageManagerClass = _AssetDownloadStorageManagerClass{objc.GetClass("AVAssetDownloadStorageManager")}
	})
	return AssetDownloadStorageManagerClass
}

type _AssetDownloadStorageManagerClass struct {
	class objc.Class
}





// An interface definition for the [AssetDownloadStorageManager] class.
type IAssetDownloadStorageManager interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	SetStorageManagementPolicyForURL(storageManagementPolicy IAVAssetDownloadStorageManagementPolicy, downloadStorageURL objc.IObject /* cross-framework: NSURL */)
	StorageManagementPolicyForURL(downloadStorageURL objc.IObject /* cross-framework: NSURL */) IAssetDownloadStorageManagementPolicy


}





// Alloc allocates a new instance without initialization.
func (ac _AssetDownloadStorageManagerClass) Alloc() AssetDownloadStorageManager {
	rv := objc.Send[AssetDownloadStorageManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetDownloadStorageManagerClass) New() AssetDownloadStorageManager {
	rv := objc.Send[AssetDownloadStorageManager](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetDownloadStorageManager) Init() AssetDownloadStorageManager {
	rv := objc.Send[AssetDownloadStorageManager](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetDownloadStorageManager) Autorelease() AssetDownloadStorageManager {
	rv := objc.Send[AssetDownloadStorageManager](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetDownloadStorageManager creates a new AssetDownloadStorageManager instance.
func NewAssetDownloadStorageManager() AssetDownloadStorageManager {
	return getAssetDownloadStorageManagerClass().New()
}





// An object that manages policies to automatically purge downloaded assets.


// An object that manages policies to automatically purge downloaded assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManager
type AssetDownloadStorageManager struct {
	objectivec.Object
}

// AssetDownloadStorageManagerFrom constructs a [AssetDownloadStorageManager] from an unsafe.Pointer.
//
// An object that manages policies to automatically purge downloaded assets.
func AssetDownloadStorageManagerFrom(ptr unsafe.Pointer) AssetDownloadStorageManager {
	return AssetDownloadStorageManager{objectivec.Object{objc.ID(ptr)}}
}










// Returns the shared storage manager instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManager/shared()
func (ac _AssetDownloadStorageManagerClass) SharedDownloadStorageManager() IAssetDownloadStorageManager {
	rv := objc.Send[AssetDownloadStorageManager](objc.ID(ac.class), objc.Sel("sharedDownloadStorageManager"))
	return rv
}












// Sets a storage policy for the downloaded asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManager/setStorageManagementPolicy(_:for:)
func (a_ AssetDownloadStorageManager) SetStorageManagementPolicyForURL(storageManagementPolicy IAVAssetDownloadStorageManagementPolicy, downloadStorageURL objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStorageManagementPolicy:forURL:"), storageManagementPolicy, downloadStorageURL)
}


// Returns the storage management policy for a downloaded asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManager/storageManagementPolicy(for:)
func (a_ AssetDownloadStorageManager) StorageManagementPolicyForURL(downloadStorageURL objc.IObject /* cross-framework: NSURL */) IAssetDownloadStorageManagementPolicy {
	rv := objc.Send[AssetDownloadStorageManagementPolicy](a_.ID, objc.Sel("storageManagementPolicyForURL:"), downloadStorageURL)
	return rv
}













