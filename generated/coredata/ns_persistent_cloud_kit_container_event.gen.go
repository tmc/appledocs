// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	Error() foundation.Error
	Succeeded() bool
	EndDate() foundation.Date
	SetEndDate(value foundation.IDate)
	Identifier() foundation.UUID
	SetIdentifier(value foundation.IUUID)
	StartDate() foundation.Date
	SetStartDate(value foundation.IDate)
	StoreIdentifier() string
	SetStoreIdentifier(value string)
	Type() unsafe.Pointer
	SetType(value unsafe.Pointer)
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
func (p_ PersistentCloudKitContainerEvent) Error() foundation.Error {
	rv := objc.Send[foundation.Error](p_.ID, objc.Sel("error"))
	return rv
}

// A Boolean value that indicates whether the operation the event represents is successful.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainer/Event/succeeded
func (p_ PersistentCloudKitContainerEvent) Succeeded() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("succeeded"))
	return rv
}

// The end date of the operation that the event represents.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/event/enddate
func (p_ PersistentCloudKitContainerEvent) EndDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("endDate"))
	return rv
}


// SetEndDate sets the value of the endDate property.
// The end date of the operation that the event represents.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/event/enddate
func (p_ PersistentCloudKitContainerEvent) SetEndDate(value foundation.IDate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEndDate:"), value)
}

// A unique identifier for the event in a container.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/event/identifier
func (p_ PersistentCloudKitContainerEvent) Identifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](p_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// A unique identifier for the event in a container.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/event/identifier
func (p_ PersistentCloudKitContainerEvent) SetIdentifier(value foundation.IUUID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIdentifier:"), value)
}

// The start date of the operation that the event represents.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/event/startdate
func (p_ PersistentCloudKitContainerEvent) StartDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("startDate"))
	return rv
}


// SetStartDate sets the value of the startDate property.
// The start date of the operation that the event represents.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/event/startdate
func (p_ PersistentCloudKitContainerEvent) SetStartDate(value foundation.IDate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStartDate:"), value)
}

// The associated store identifier in the container for the event.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/event/storeidentifier
func (p_ PersistentCloudKitContainerEvent) StoreIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("storeIdentifier"))
	return rv
}


// SetStoreIdentifier sets the value of the storeIdentifier property.
// The associated store identifier in the container for the event.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/event/storeidentifier
func (p_ PersistentCloudKitContainerEvent) SetStoreIdentifier(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStoreIdentifier:"), objc.String(value))
}

// The type of event, either setup, import, or export.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/event/type
func (p_ PersistentCloudKitContainerEvent) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The type of event, either setup, import, or export.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentcloudkitcontainer/event/type
func (p_ PersistentCloudKitContainerEvent) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setType:"), value)
}



