// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EKObject] class.
var (
	EKObjectClass     _EKObjectClass
	EKObjectClassOnce sync.Once
)

func getEKObjectClass() _EKObjectClass {
	EKObjectClassOnce.Do(func() {
		EKObjectClass = _EKObjectClass{objc.GetClass("EKObject")}
	})
	return EKObjectClass
}

type _EKObjectClass struct {
	class objc.Class
}

// An interface definition for the [EKObject] class.
type IEKObject interface {
	objectivec.IObject
	Refresh() bool
	Reset()
	Rollback()
	HasChanges() bool
	New() bool
	IsNew() bool
	SetIsNew(value bool)
}

// An abstract superclass for all EventKit classes that have persistent instances.
//
// provides fine control when saving and restoring property settings. For example, you can find out if a persistent object was modified locally and whether it needs to be saved. If the object has changed in the event store since it was fetched, you can refresh the local copy by keeping local changes or by removing local changes. You can also roll back the object to the state when it was first fetched.


// An abstract superclass for all EventKit classes that have persistent instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKObject

type EKObject struct {
	objectivec.Object
}

// EKObjectFrom constructs a [EKObject] from an unsafe.Pointer.
//
// An abstract superclass for all EventKit classes that have persistent instances.
func EKObjectFrom(ptr unsafe.Pointer) EKObject {
	return EKObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EKObjectClass) Alloc() EKObject {
	rv := objc.Send[EKObject](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EKObjectClass) New() EKObject {
	rv := objc.Send[EKObject](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKObject) Init() EKObject {
	rv := objc.Send[EKObject](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKObject) Autorelease() EKObject {
	rv := objc.Send[EKObject](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKObject creates a new EKObject instance.
func NewEKObject() EKObject {
	return getEKObjectClass().New()
}




// Merges changes to this object with the latest saved values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKObject/refresh()

func (e_ EKObject) Refresh() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("refresh"))
	return rv
}



// Returns this object to its saved state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKObject/reset()

func (e_ EKObject) Reset() {
	objc.Send[objc.ID](e_.ID, objc.Sel("reset"))
}



// Rolls back the property values of this object to its original state when it was first fetched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKObject/rollback()

func (e_ EKObject) Rollback() {
	objc.Send[objc.ID](e_.ID, objc.Sel("rollback"))
}


// Returns whether this object or any of the objects it contains has uncommitted changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKObject/hasChanges

func (e_ EKObject) HasChanges() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasChanges"))
	return rv
}


// A Boolean value that indicates whether this object has ever been saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKObject/isNew

func (e_ EKObject) New() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("new"))
	return rv
}


// A Boolean value that indicates whether this object has ever been saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekobject/isnew

func (e_ EKObject) IsNew() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isNew"))
	return rv
}


// A Boolean value that indicates whether this object has ever been saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekobject/isnew

func (e_ EKObject) SetIsNew(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsNew:"), value)
}



