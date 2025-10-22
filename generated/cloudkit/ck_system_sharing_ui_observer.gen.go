// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSystemSharingUIObserver] class.
var (
	CKSystemSharingUIObserverClass     _CKSystemSharingUIObserverClass
	CKSystemSharingUIObserverClassOnce sync.Once
)

func getCKSystemSharingUIObserverClass() _CKSystemSharingUIObserverClass {
	CKSystemSharingUIObserverClassOnce.Do(func() {
		CKSystemSharingUIObserverClass = _CKSystemSharingUIObserverClass{objc.GetClass("CKSystemSharingUIObserver")}
	})
	return CKSystemSharingUIObserverClass
}

type _CKSystemSharingUIObserverClass struct {
	class objc.Class
}

// An interface definition for the [CKSystemSharingUIObserver] class.
type ICKSystemSharingUIObserver interface {
	objectivec.IObject
	SystemSharingUIDidSaveShareBlock() unsafe.Pointer
	SetSystemSharingUIDidSaveShareBlock(value unsafe.Pointer)
	SystemSharingUIDidStopSharingBlock() unsafe.Pointer
	SetSystemSharingUIDidStopSharingBlock(value unsafe.Pointer)
}

// An object the system uses to monitor changes in sharing.
//
// Initialize a instance with your when preparing to share an item using the share sheet. Use your implementation to update the local state of a shared item when your app receives a , or to delete a locally cached share when the system notifies your app about a share deletion. The system only propagates changes on the local device using  . The system doesn’t notify your app about any remote changes on the server. For more information about how to keep your local cache in sync with remote changes, see .


// An object the system uses to monitor changes in sharing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver

type CKSystemSharingUIObserver struct {
	objectivec.Object
}

// CKSystemSharingUIObserverFrom constructs a [CKSystemSharingUIObserver] from an unsafe.Pointer.
//
// An object the system uses to monitor changes in sharing.
func CKSystemSharingUIObserverFrom(ptr unsafe.Pointer) CKSystemSharingUIObserver {
	return CKSystemSharingUIObserver{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSystemSharingUIObserverClass) Alloc() CKSystemSharingUIObserver {
	rv := objc.Send[CKSystemSharingUIObserver](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSystemSharingUIObserverClass) New() CKSystemSharingUIObserver {
	rv := objc.Send[CKSystemSharingUIObserver](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSystemSharingUIObserver) Init() CKSystemSharingUIObserver {
	rv := objc.Send[CKSystemSharingUIObserver](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSystemSharingUIObserver) Autorelease() CKSystemSharingUIObserver {
	rv := objc.Send[CKSystemSharingUIObserver](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSystemSharingUIObserver creates a new CKSystemSharingUIObserver instance.
func NewCKSystemSharingUIObserver() CKSystemSharingUIObserver {
	return getCKSystemSharingUIObserverClass().New()
}




// Creates and initializes an observer using the provided container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver/init(container:)

func NewCKSystemSharingUIObserverWithContainer(container ICKContainer) CKSystemSharingUIObserver {
	instance := getCKSystemSharingUIObserverClass().Alloc()
	rv := objc.Send[CKSystemSharingUIObserver](instance.ID, objc.Sel("initWithContainer:"), container)
	rv.Autorelease()
	return rv
}



// A callback block the system invokes after the success or failure of a system sharing UI save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver/systemSharingUIDidSaveShareBlock-39zlv

func (c_ CKSystemSharingUIObserver) SystemSharingUIDidSaveShareBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("systemSharingUIDidSaveShareBlock"))
	return rv
}


// A callback block the system invokes after the success or failure of a system sharing UI save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver/systemSharingUIDidSaveShareBlock-39zlv

func (c_ CKSystemSharingUIObserver) SetSystemSharingUIDidSaveShareBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemSharingUIDidSaveShareBlock:"), value)
}


// A callback block the system invokes after the success or failure of a system sharing UI delete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver/systemSharingUIDidStopSharingBlock-4g5bn

func (c_ CKSystemSharingUIObserver) SystemSharingUIDidStopSharingBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("systemSharingUIDidStopSharingBlock"))
	return rv
}


// A callback block the system invokes after the success or failure of a system sharing UI delete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver/systemSharingUIDidStopSharingBlock-4g5bn

func (c_ CKSystemSharingUIObserver) SetSystemSharingUIDidStopSharingBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemSharingUIDidStopSharingBlock:"), value)
}


