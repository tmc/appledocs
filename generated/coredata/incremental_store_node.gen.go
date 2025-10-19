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



