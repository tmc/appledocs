// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [IncrementalStoreNode] class.
var (
	IncrementalStoreNodeClass     _IncrementalStoreNodeClass
	IncrementalStoreNodeClassOnce sync.Once
)

func getIncrementalStoreNodeClass() _IncrementalStoreNodeClass {
	IncrementalStoreNodeClassOnce.Do(func() {
		IncrementalStoreNodeClass = _IncrementalStoreNodeClass{objc.GetClass("NSIncrementalStoreNode")}
	})
	return IncrementalStoreNodeClass
}

type _IncrementalStoreNodeClass struct {
	class objc.Class
}

// An interface definition for the [IncrementalStoreNode] class.
type IIncrementalStoreNode interface {
	objectivec.IObject
	// properties:
	ObjectID() IManagedObjectID
	Version() uint64
	SetVersion(value uint64)
	// methods:
}

// A concrete class used to represent basic nodes in a Core Data incremental store.
//
// A node represents a single record in a persistent store. You can subclass to provide custom behavior.

// A concrete class used to represent basic nodes in a Core Data incremental store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSIncrementalStoreNode
type IncrementalStoreNode struct {
	objectivec.Object
}

// IncrementalStoreNodeFrom constructs a [IncrementalStoreNode] from an unsafe.Pointer.
//
// A concrete class used to represent basic nodes in a Core Data incremental store.
func IncrementalStoreNodeFrom(ptr unsafe.Pointer) IncrementalStoreNode {
	return IncrementalStoreNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _IncrementalStoreNodeClass) Alloc() IncrementalStoreNode {
	rv := objc.Send[IncrementalStoreNode](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IncrementalStoreNodeClass) New() IncrementalStoreNode {
	rv := objc.Send[IncrementalStoreNode](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IncrementalStoreNode) Init() IncrementalStoreNode {
	rv := objc.Send[IncrementalStoreNode](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IncrementalStoreNode) Autorelease() IncrementalStoreNode {
	rv := objc.Send[IncrementalStoreNode](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIncrementalStoreNode creates a new IncrementalStoreNode instance.
func NewIncrementalStoreNode() IncrementalStoreNode {
	return getIncrementalStoreNodeClass().New()
}

// The object ID that identifies the data stored by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSIncrementalStoreNode/objectID
func (i_ IncrementalStoreNode) ObjectID() IManagedObjectID {
	rv := objc.Send[ManagedObjectID](i_.ID, objc.Sel("objectID"))
	return rv
}

// The version of data in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstorenode/version
func (i_ IncrementalStoreNode) Version() uint64 {
	rv := objc.Send[uint64](i_.ID, objc.Sel("version"))
	return rv
}

// The version of data in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsincrementalstorenode/version
func (i_ IncrementalStoreNode) SetVersion(value uint64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setVersion:"), value)
}
