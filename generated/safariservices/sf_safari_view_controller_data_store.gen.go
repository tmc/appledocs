// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFSafariViewControllerDataStore] class.
var (
	SFSafariViewControllerDataStoreClass     _SFSafariViewControllerDataStoreClass
	SFSafariViewControllerDataStoreClassOnce sync.Once
)

func getSFSafariViewControllerDataStoreClass() _SFSafariViewControllerDataStoreClass {
	SFSafariViewControllerDataStoreClassOnce.Do(func() {
		SFSafariViewControllerDataStoreClass = _SFSafariViewControllerDataStoreClass{objc.GetClass("SFSafariViewControllerDataStore")}
	})
	return SFSafariViewControllerDataStoreClass
}

type _SFSafariViewControllerDataStoreClass struct {
	class objc.Class
}

// An interface definition for the [SFSafariViewControllerDataStore] class.
type ISFSafariViewControllerDataStore interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/DataStore
type SFSafariViewControllerDataStore struct {
	objectivec.Object
}

// SFSafariViewControllerDataStoreFrom constructs a [SFSafariViewControllerDataStore] from an unsafe.Pointer.
func SFSafariViewControllerDataStoreFrom(ptr unsafe.Pointer) SFSafariViewControllerDataStore {
	return SFSafariViewControllerDataStore{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSafariViewControllerDataStoreClass) Alloc() SFSafariViewControllerDataStore {
	rv := objc.Send[SFSafariViewControllerDataStore](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSafariViewControllerDataStoreClass) New() SFSafariViewControllerDataStore {
	rv := objc.Send[SFSafariViewControllerDataStore](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariViewControllerDataStore) Init() SFSafariViewControllerDataStore {
	rv := objc.Send[SFSafariViewControllerDataStore](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariViewControllerDataStore) Autorelease() SFSafariViewControllerDataStore {
	rv := objc.Send[SFSafariViewControllerDataStore](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariViewControllerDataStore creates a new SFSafariViewControllerDataStore instance.
func NewSFSafariViewControllerDataStore() SFSafariViewControllerDataStore {
	return getSFSafariViewControllerDataStoreClass().New()
}




