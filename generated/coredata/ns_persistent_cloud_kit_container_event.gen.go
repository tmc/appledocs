// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentCloudKitContainerEvent] class.
var (
	PersistentCloudKitContainerEventClass     _PersistentCloudKitContainerEventClass
	PersistentCloudKitContainerEventClassOnce sync.Once
)

func getPersistentCloudKitContainerEventClass() _PersistentCloudKitContainerEventClass {
	PersistentCloudKitContainerEventClassOnce.Do(func() {
		PersistentCloudKitContainerEventClass = _PersistentCloudKitContainerEventClass{objc.GetClass("NSPersistentCloudKitContainerEvent")}
	})
	return PersistentCloudKitContainerEventClass
}

type _PersistentCloudKitContainerEventClass struct {
	class objc.Class
}

// An interface definition for the [PersistentCloudKitContainerEvent] class.
type IPersistentCloudKitContainerEvent interface {
	objectivec.IObject
}

// An object that represents activity in a persistent CloudKit container.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/Event
type PersistentCloudKitContainerEvent struct {
	objectivec.Object
}

// PersistentCloudKitContainerEventFrom constructs a [PersistentCloudKitContainerEvent] from an unsafe.Pointer.
//
// An object that represents activity in a persistent CloudKit container.
func PersistentCloudKitContainerEventFrom(ptr unsafe.Pointer) PersistentCloudKitContainerEvent {
	return PersistentCloudKitContainerEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentCloudKitContainerEventClass) Alloc() PersistentCloudKitContainerEvent {
	rv := objc.Send[PersistentCloudKitContainerEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentCloudKitContainerEventClass) New() PersistentCloudKitContainerEvent {
	rv := objc.Send[PersistentCloudKitContainerEvent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentCloudKitContainerEvent) Init() PersistentCloudKitContainerEvent {
	rv := objc.Send[PersistentCloudKitContainerEvent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentCloudKitContainerEvent) Autorelease() PersistentCloudKitContainerEvent {
	rv := objc.Send[PersistentCloudKitContainerEvent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentCloudKitContainerEvent creates a new PersistentCloudKitContainerEvent instance.
func NewPersistentCloudKitContainerEvent() PersistentCloudKitContainerEvent {
	return getPersistentCloudKitContainerEventClass().New()
}


// An error that indicates why an operation fails.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/Event/error
func (p_ PersistentCloudKitContainerEvent) Error() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("error"))
	return rv
}

// A Boolean value that indicates whether the operation the event represents is successful.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/Event/succeeded
func (p_ PersistentCloudKitContainerEvent) Succeeded() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("succeeded"))
	return rv
}



