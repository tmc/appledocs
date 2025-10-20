// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AssetDownloadStorageManagementPolicy] class.
var (
	AssetDownloadStorageManagementPolicyClass     _AssetDownloadStorageManagementPolicyClass
	AssetDownloadStorageManagementPolicyClassOnce sync.Once
)

func getAssetDownloadStorageManagementPolicyClass() _AssetDownloadStorageManagementPolicyClass {
	AssetDownloadStorageManagementPolicyClassOnce.Do(func() {
		AssetDownloadStorageManagementPolicyClass = _AssetDownloadStorageManagementPolicyClass{objc.GetClass("AVAssetDownloadStorageManagementPolicy")}
	})
	return AssetDownloadStorageManagementPolicyClass
}

type _AssetDownloadStorageManagementPolicyClass struct {
	class objc.Class
}

// An interface definition for the [AssetDownloadStorageManagementPolicy] class.
type IAssetDownloadStorageManagementPolicy interface {
	objectivec.IObject
}

// An object that defines a policy to automatically manage the storage of downloaded assets.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManagementPolicy
type AssetDownloadStorageManagementPolicy struct {
	objectivec.Object
}

// AssetDownloadStorageManagementPolicyFrom constructs a [AssetDownloadStorageManagementPolicy] from an unsafe.Pointer.
//
// An object that defines a policy to automatically manage the storage of downloaded assets.
func AssetDownloadStorageManagementPolicyFrom(ptr unsafe.Pointer) AssetDownloadStorageManagementPolicy {
	return AssetDownloadStorageManagementPolicy{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetDownloadStorageManagementPolicyClass) Alloc() AssetDownloadStorageManagementPolicy {
	rv := objc.Send[AssetDownloadStorageManagementPolicy](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetDownloadStorageManagementPolicyClass) New() AssetDownloadStorageManagementPolicy {
	rv := objc.Send[AssetDownloadStorageManagementPolicy](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetDownloadStorageManagementPolicy) Init() AssetDownloadStorageManagementPolicy {
	rv := objc.Send[AssetDownloadStorageManagementPolicy](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetDownloadStorageManagementPolicy) Autorelease() AssetDownloadStorageManagementPolicy {
	rv := objc.Send[AssetDownloadStorageManagementPolicy](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetDownloadStorageManagementPolicy creates a new AssetDownloadStorageManagementPolicy instance.
func NewAssetDownloadStorageManagementPolicy() AssetDownloadStorageManagementPolicy {
	return getAssetDownloadStorageManagementPolicyClass().New()
}


// The expiration date for an asset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManagementPolicy/expirationDate
func (a_ AssetDownloadStorageManagementPolicy) ExpirationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("expirationDate"))
	return rv
}

// The eviction priority for an asset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManagementPolicy/priority
func (a_ AssetDownloadStorageManagementPolicy) Priority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("priority"))
	return rv
}



