// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AtomicStoreCacheNode] class.
var atomicStoreCacheNodeClass = _AtomicStoreCacheNodeClass{objc.GetClass("NSAtomicStoreCacheNode")}

type _AtomicStoreCacheNodeClass struct {
	class objc.Class
}

// A concrete class that you use to represent basic nodes in a Core Data atomic store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAtomicStoreCacheNode

type AtomicStoreCacheNode struct {
	objectivec.Object
}

// AtomicStoreCacheNodeFrom constructs a [AtomicStoreCacheNode] from an unsafe.Pointer.
//
// A concrete class that you use to represent basic nodes in a Core Data atomic store.
func AtomicStoreCacheNodeFrom(ptr unsafe.Pointer) AtomicStoreCacheNode {
	return AtomicStoreCacheNode{objectivec.Object{objc.ID(ptr)}}
}

// Returns the value for a given key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAtomicStoreCacheNode/value(forKey:)
func (a_ AtomicStoreCacheNode) ValueForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("valueForKey:"), key)
	return rv
}


