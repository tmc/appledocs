// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentCloudKitContainerEvent] class.
var persistentCloudKitContainerEventClass = _PersistentCloudKitContainerEventClass{objc.GetClass("NSPersistentCloudKitContainerEvent")}

type _PersistentCloudKitContainerEventClass struct {
	class objc.Class
}

// An object that represents activity in a persistent CloudKit container. [Full Topic]
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



