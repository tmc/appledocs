// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetDownloadStorageManager */


/* debug [class_header]: Header for AVAssetDownloadStorageManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetDownloadStorageManager */
// An interface definition for the [AssetDownloadStorageManager] class.
type IAssetDownloadStorageManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetDownloadStorageManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetDownloadStorageManager */
	// methods:
	SetStorageManagementPolicyForURL(storageManagementPolicy IAVAssetDownloadStorageManagementPolicy, downloadStorageURL objc.IObject /* cross-framework: NSURL */)
	StorageManagementPolicyForURL(downloadStorageURL objc.IObject /* cross-framework: NSURL */) IAssetDownloadStorageManagementPolicy
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetDownloadStorageManager */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetDownloadStorageManager */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetDownloadStorageManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetDownloadStorageManager */

// Returns the shared storage manager instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManager/shared()
func (ac _AssetDownloadStorageManagerClass) SharedDownloadStorageManager() IAssetDownloadStorageManager {
	rv := objc.Send[AssetDownloadStorageManager](objc.ID(ac.class), objc.Sel("sharedDownloadStorageManager"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedDownloadStorageManager) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetDownloadStorageManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetDownloadStorageManager */

// Sets a storage policy for the downloaded asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManager/setStorageManagementPolicy(_:for:)
func (a_ AssetDownloadStorageManager) SetStorageManagementPolicyForURL(storageManagementPolicy IAVAssetDownloadStorageManagementPolicy, downloadStorageURL objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStorageManagementPolicy:forURL:"), storageManagementPolicy, downloadStorageURL)
}/* debug [instance_methods/method]: SetStorageManagementPolicyForURL */


// Returns the storage management policy for a downloaded asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManager/storageManagementPolicy(for:)
func (a_ AssetDownloadStorageManager) StorageManagementPolicyForURL(downloadStorageURL objc.IObject /* cross-framework: NSURL */) IAssetDownloadStorageManagementPolicy {
	rv := objc.Send[AssetDownloadStorageManagementPolicy](a_.ID, objc.Sel("storageManagementPolicyForURL:"), downloadStorageURL)
	return rv
}/* debug [instance_methods/method]: StorageManagementPolicyForURL */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetDownloadStorageManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetDownloadStorageManager */



