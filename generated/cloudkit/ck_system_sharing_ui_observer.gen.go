// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSystemSharingUIObserver */


/* debug [class_header]: Header for CKSystemSharingUIObserver */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSystemSharingUIObserver */
// An interface definition for the [CKSystemSharingUIObserver] class.
type ICKSystemSharingUIObserver interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSystemSharingUIObserver */
	// properties:
	SystemSharingUIDidSaveShareBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	SetSystemSharingUIDidSaveShareBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer))
	SystemSharingUIDidStopSharingBlock() func(unsafe.Pointer, unsafe.Pointer)
	SetSystemSharingUIDidStopSharingBlock(value func(unsafe.Pointer, unsafe.Pointer))
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSystemSharingUIObserver */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSystemSharingUIObserver */
// Alloc allocates a new instance without initialization.
func (cc _CKSystemSharingUIObserverClass) Alloc() CKSystemSharingUIObserver {
	rv := objc.Send[CKSystemSharingUIObserver](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSystemSharingUIObserver */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSystemSharingUIObserver */

// Creates and initializes an observer using the provided container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver/init(container:)
func NewCKSystemSharingUIObserverWithContainer(container ICKContainer) CKSystemSharingUIObserver {
	instance := getCKSystemSharingUIObserverClass().Alloc()
	rv := objc.Send[CKSystemSharingUIObserver](instance.ID, objc.Sel("initWithContainer:"), container)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSystemSharingUIObserverWithContainer */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSystemSharingUIObserver */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSystemSharingUIObserver */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSystemSharingUIObserver */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSystemSharingUIObserver */

// A callback block the system invokes after the success or failure of a system sharing UI save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver/systemSharingUIDidSaveShareBlock-39zlv
func (c_ CKSystemSharingUIObserver) SystemSharingUIDidSaveShareBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("systemSharingUIDidSaveShareBlock"))
	return rv
}/* debug [instance_properties/getter]: systemSharingUIDidSaveShareBlock */


// A callback block the system invokes after the success or failure of a system sharing UI save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver/systemSharingUIDidSaveShareBlock-39zlv
func (c_ CKSystemSharingUIObserver) SetSystemSharingUIDidSaveShareBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemSharingUIDidSaveShareBlock:"), value)
}/* debug [instance_properties/setter]: systemSharingUIDidSaveShareBlock */


// A callback block the system invokes after the success or failure of a system sharing UI delete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver/systemSharingUIDidStopSharingBlock-4g5bn
func (c_ CKSystemSharingUIObserver) SystemSharingUIDidStopSharingBlock() func(unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("systemSharingUIDidStopSharingBlock"))
	return rv
}/* debug [instance_properties/getter]: systemSharingUIDidStopSharingBlock */


// A callback block the system invokes after the success or failure of a system sharing UI delete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSystemSharingUIObserver/systemSharingUIDidStopSharingBlock-4g5bn
func (c_ CKSystemSharingUIObserver) SetSystemSharingUIDidStopSharingBlock(value func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemSharingUIDidStopSharingBlock:"), value)
}/* debug [instance_properties/setter]: systemSharingUIDidStopSharingBlock */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSystemSharingUIObserver */


