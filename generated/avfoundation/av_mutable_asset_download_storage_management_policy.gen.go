// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableAssetDownloadStorageManagementPolicy] class.
var (
	MutableAssetDownloadStorageManagementPolicyClass     _MutableAssetDownloadStorageManagementPolicyClass
	MutableAssetDownloadStorageManagementPolicyClassOnce sync.Once
)

func getMutableAssetDownloadStorageManagementPolicyClass() _MutableAssetDownloadStorageManagementPolicyClass {
	MutableAssetDownloadStorageManagementPolicyClassOnce.Do(func() {
		MutableAssetDownloadStorageManagementPolicyClass = _MutableAssetDownloadStorageManagementPolicyClass{objc.GetClass("AVMutableAssetDownloadStorageManagementPolicy")}
	})
	return MutableAssetDownloadStorageManagementPolicyClass
}

type _MutableAssetDownloadStorageManagementPolicyClass struct {
	class objc.Class
}

// An interface definition for the [MutableAssetDownloadStorageManagementPolicy] class.
type IMutableAssetDownloadStorageManagementPolicy interface {
	IAssetDownloadStorageManagementPolicy
}

// A mutable object that you use to create a new storage management policy.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAssetDownloadStorageManagementPolicy
type MutableAssetDownloadStorageManagementPolicy struct {
	AssetDownloadStorageManagementPolicy
}

// MutableAssetDownloadStorageManagementPolicyFrom constructs a [MutableAssetDownloadStorageManagementPolicy] from an unsafe.Pointer.
//
// A mutable object that you use to create a new storage management policy.
func MutableAssetDownloadStorageManagementPolicyFrom(ptr unsafe.Pointer) MutableAssetDownloadStorageManagementPolicy {
	return MutableAssetDownloadStorageManagementPolicy{
		AssetDownloadStorageManagementPolicy: AssetDownloadStorageManagementPolicyFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableAssetDownloadStorageManagementPolicyClass) Alloc() MutableAssetDownloadStorageManagementPolicy {
	rv := objc.Send[MutableAssetDownloadStorageManagementPolicy](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableAssetDownloadStorageManagementPolicyClass) New() MutableAssetDownloadStorageManagementPolicy {
	rv := objc.Send[MutableAssetDownloadStorageManagementPolicy](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableAssetDownloadStorageManagementPolicy) Init() MutableAssetDownloadStorageManagementPolicy {
	rv := objc.Send[MutableAssetDownloadStorageManagementPolicy](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableAssetDownloadStorageManagementPolicy) Autorelease() MutableAssetDownloadStorageManagementPolicy {
	rv := objc.Send[MutableAssetDownloadStorageManagementPolicy](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableAssetDownloadStorageManagementPolicy creates a new MutableAssetDownloadStorageManagementPolicy instance.
func NewMutableAssetDownloadStorageManagementPolicy() MutableAssetDownloadStorageManagementPolicy {
	return getMutableAssetDownloadStorageManagementPolicyClass().New()
}




