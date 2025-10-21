// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CKSyncEngineFetchChangesOptions] class.
var (
	CKSyncEngineFetchChangesOptionsClass     _CKSyncEngineFetchChangesOptionsClass
	CKSyncEngineFetchChangesOptionsClassOnce sync.Once
)

func getCKSyncEngineFetchChangesOptionsClass() _CKSyncEngineFetchChangesOptionsClass {
	CKSyncEngineFetchChangesOptionsClassOnce.Do(func() {
		CKSyncEngineFetchChangesOptionsClass = _CKSyncEngineFetchChangesOptionsClass{objc.GetClass("CKSyncEngineFetchChangesOptions")}
	})
	return CKSyncEngineFetchChangesOptionsClass
}

type _CKSyncEngineFetchChangesOptionsClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineFetchChangesOptions] class.
type ICKSyncEngineFetchChangesOptions interface {
	objectivec.IObject
}

// A set of options to use with a fetch operation.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesOptions
type CKSyncEngineFetchChangesOptions struct {
	objectivec.Object
}

// CKSyncEngineFetchChangesOptionsFrom constructs a [CKSyncEngineFetchChangesOptions] from an unsafe.Pointer.
//
// A set of options to use with a fetch operation.
func CKSyncEngineFetchChangesOptionsFrom(ptr unsafe.Pointer) CKSyncEngineFetchChangesOptions {
	return CKSyncEngineFetchChangesOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFetchChangesOptionsClass) Alloc() CKSyncEngineFetchChangesOptions {
	rv := objc.Send[CKSyncEngineFetchChangesOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineFetchChangesOptionsClass) New() CKSyncEngineFetchChangesOptions {
	rv := objc.Send[CKSyncEngineFetchChangesOptions](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFetchChangesOptions) Init() CKSyncEngineFetchChangesOptions {
	rv := objc.Send[CKSyncEngineFetchChangesOptions](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFetchChangesOptions) Autorelease() CKSyncEngineFetchChangesOptions {
	rv := objc.Send[CKSyncEngineFetchChangesOptions](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchChangesOptions creates a new CKSyncEngineFetchChangesOptions instance.
func NewCKSyncEngineFetchChangesOptions() CKSyncEngineFetchChangesOptions {
	return getCKSyncEngineFetchChangesOptionsClass().New()
}




