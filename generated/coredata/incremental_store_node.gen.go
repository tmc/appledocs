// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [IncrementalStoreNode] class.
var incrementalStoreNodeClass = _IncrementalStoreNodeClass{objc.GetClass("NSIncrementalStoreNode")}

type _IncrementalStoreNodeClass struct {
	class objc.Class
}

// An interface definition for the [IncrementalStoreNode] class.
type IIncrementalStoreNode interface {
	objectivec.IObject
}

// A concrete class used to represent basic nodes in a Core Data incremental store. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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
	return incrementalStoreNodeClass.New()
}




