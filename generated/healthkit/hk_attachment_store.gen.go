// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HKAttachmentStore] class.
var (
	HKAttachmentStoreClass     _HKAttachmentStoreClass
	HKAttachmentStoreClassOnce sync.Once
)

func getHKAttachmentStoreClass() _HKAttachmentStoreClass {
	HKAttachmentStoreClassOnce.Do(func() {
		HKAttachmentStoreClass = _HKAttachmentStoreClass{objc.GetClass("HKAttachmentStore")}
	})
	return HKAttachmentStoreClass
}

type _HKAttachmentStoreClass struct {
	class objc.Class
}

// An interface definition for the [HKAttachmentStore] class.
type IHKAttachmentStore interface {
	objectivec.IObject
	// properties:
	// methods:
}

// The access point for attachments associated with samples in the HealthKit store.
//
// Use an object to manage attachments for samples that your app has saved to the HealthKit store.


// The access point for attachments associated with samples in the HealthKit store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAttachmentStore
type HKAttachmentStore struct {
	objectivec.Object
}

// HKAttachmentStoreFrom constructs a [HKAttachmentStore] from an unsafe.Pointer.
//
// The access point for attachments associated with samples in the HealthKit store.
func HKAttachmentStoreFrom(ptr unsafe.Pointer) HKAttachmentStore {
	return HKAttachmentStore{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKAttachmentStoreClass) Alloc() HKAttachmentStore {
	rv := objc.Send[HKAttachmentStore](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKAttachmentStoreClass) New() HKAttachmentStore {
	rv := objc.Send[HKAttachmentStore](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKAttachmentStore) Init() HKAttachmentStore {
	rv := objc.Send[HKAttachmentStore](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKAttachmentStore) Autorelease() HKAttachmentStore {
	rv := objc.Send[HKAttachmentStore](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKAttachmentStore creates a new HKAttachmentStore instance.
func NewHKAttachmentStore() HKAttachmentStore {
	return getHKAttachmentStoreClass().New()
}




