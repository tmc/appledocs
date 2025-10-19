// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BatchDeleteRequest] class.
var (
	batchDeleteRequestClass     _BatchDeleteRequestClass
	batchDeleteRequestClassOnce sync.Once
)

func getBatchDeleteRequestClass() _BatchDeleteRequestClass {
	batchDeleteRequestClassOnce.Do(func() {
		batchDeleteRequestClass = _BatchDeleteRequestClass{objc.GetClass("NSBatchDeleteRequest")}
	})
	return batchDeleteRequestClass
}

type _BatchDeleteRequestClass struct {
	class objc.Class
}

// An interface definition for the [BatchDeleteRequest] class.
type IBatchDeleteRequest interface {
	IPersistentStoreRequest
}

// A request that deletes objects in the SQLite persistent store without loading them into memory.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequest
type BatchDeleteRequest struct {
	PersistentStoreRequest
}

// BatchDeleteRequestFrom constructs a [BatchDeleteRequest] from an unsafe.Pointer.
//
// A request that deletes objects in the SQLite persistent store without loading them into memory.
func BatchDeleteRequestFrom(ptr unsafe.Pointer) BatchDeleteRequest {
	return BatchDeleteRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BatchDeleteRequestClass) Alloc() BatchDeleteRequest {
	rv := objc.Send[BatchDeleteRequest](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BatchDeleteRequestClass) New() BatchDeleteRequest {
	rv := objc.Send[BatchDeleteRequest](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BatchDeleteRequest) Init() BatchDeleteRequest {
	rv := objc.Send[BatchDeleteRequest](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BatchDeleteRequest) Autorelease() BatchDeleteRequest {
	rv := objc.Send[BatchDeleteRequest](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBatchDeleteRequest creates a new BatchDeleteRequest instance.
func NewBatchDeleteRequest() BatchDeleteRequest {
	return getBatchDeleteRequestClass().New()
}




