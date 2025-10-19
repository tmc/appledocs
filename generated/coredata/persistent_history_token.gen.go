// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentHistoryToken] class.
var persistentHistoryTokenClass = _PersistentHistoryTokenClass{objc.GetClass("NSPersistentHistoryToken")}

type _PersistentHistoryTokenClass struct {
	class objc.Class
}

// A bookmark for keeping track the most recent history that you’ve processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryToken

type PersistentHistoryToken struct {
	objectivec.Object
}

// PersistentHistoryTokenFrom constructs a [PersistentHistoryToken] from an unsafe.Pointer.
//
// A bookmark for keeping track the most recent history that you’ve processed.
func PersistentHistoryTokenFrom(ptr unsafe.Pointer) PersistentHistoryToken {
	return PersistentHistoryToken{objectivec.Object{objc.ID(ptr)}}
}



