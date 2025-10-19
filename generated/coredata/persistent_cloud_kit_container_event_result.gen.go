// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentCloudKitContainerEventResult] class.
var persistentCloudKitContainerEventResultClass = _PersistentCloudKitContainerEventResultClass{objc.GetClass("NSPersistentCloudKitContainerEventResult")}

type _PersistentCloudKitContainerEventResultClass struct {
	class objc.Class
}

// The result of a request to fetch persistent CloudKit container events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventResult

type PersistentCloudKitContainerEventResult struct {
	PersistentStoreResult
}

// PersistentCloudKitContainerEventResultFrom constructs a [PersistentCloudKitContainerEventResult] from an unsafe.Pointer.
//
// The result of a request to fetch persistent CloudKit container events.
func PersistentCloudKitContainerEventResultFrom(ptr unsafe.Pointer) PersistentCloudKitContainerEventResult {
	return PersistentCloudKitContainerEventResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}



