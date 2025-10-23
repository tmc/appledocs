// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AtomicStoreCacheNode] class.
var (
	AtomicStoreCacheNodeClass     _AtomicStoreCacheNodeClass
	AtomicStoreCacheNodeClassOnce sync.Once
)

func getAtomicStoreCacheNodeClass() _AtomicStoreCacheNodeClass {
	AtomicStoreCacheNodeClassOnce.Do(func() {
		AtomicStoreCacheNodeClass = _AtomicStoreCacheNodeClass{objc.GetClass("NSAtomicStoreCacheNode")}
	})
	return AtomicStoreCacheNodeClass
}

type _AtomicStoreCacheNodeClass struct {
	class objc.Class
}

// An interface definition for the [AtomicStoreCacheNode] class.
type IAtomicStoreCacheNode interface {
	objectivec.IObject
	ObjectID() NSManagedObjectID
	SetObjectID(value IManagedObjectID)
	PropertyCache() foundation.MutableDictionary
	SetPropertyCache(value foundation.IMutableDictionary)
}

// A concrete class that you use to represent basic nodes in a Core Data atomic store.
//
// A node represents a single record in a persistent store. You can subclass to provide custom behavior.


// A concrete class that you use to represent basic nodes in a Core Data atomic store.
//
// [Full Topic]
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

// Alloc allocates a new instance without initialization.
func (ac _AtomicStoreCacheNodeClass) Alloc() AtomicStoreCacheNode {
	rv := objc.Send[AtomicStoreCacheNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AtomicStoreCacheNodeClass) New() AtomicStoreCacheNode {
	rv := objc.Send[AtomicStoreCacheNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AtomicStoreCacheNode) Init() AtomicStoreCacheNode {
	rv := objc.Send[AtomicStoreCacheNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AtomicStoreCacheNode) Autorelease() AtomicStoreCacheNode {
	rv := objc.Send[AtomicStoreCacheNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAtomicStoreCacheNode creates a new AtomicStoreCacheNode instance.
func NewAtomicStoreCacheNode() AtomicStoreCacheNode {
	return getAtomicStoreCacheNodeClass().New()
}



// Sets the value for the given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAtomicStoreCacheNode/setValue(_:forKey:)
func (a_ AtomicStoreCacheNode) SetValueForKey(value objectivec.IObject, key string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setValue:forKey:"), value, objc.String(key))
}


// Returns the value for a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAtomicStoreCacheNode/value(forKey:)
func (a_ AtomicStoreCacheNode) ValueForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("valueForKey:"), objc.String(key))
	return rv
}


// The managed object ID of the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstorecachenode/objectid
func (a_ AtomicStoreCacheNode) ObjectID() NSManagedObjectID {
	rv := objc.Send[NSManagedObjectID](a_.ID, objc.Sel("objectID"))
	return rv
}


// The managed object ID of the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstorecachenode/objectid
func (a_ AtomicStoreCacheNode) SetObjectID(value IManagedObjectID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setObjectID:"), value)
}


// The property cache dictionary of the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstorecachenode/propertycache
func (a_ AtomicStoreCacheNode) PropertyCache() foundation.MutableDictionary {
	rv := objc.Send[foundation.MutableDictionary](a_.ID, objc.Sel("propertyCache"))
	return rv
}


// The property cache dictionary of the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsatomicstorecachenode/propertycache
func (a_ AtomicStoreCacheNode) SetPropertyCache(value foundation.IMutableDictionary) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPropertyCache:"), value)
}



