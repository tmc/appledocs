// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Cache] class.
var (
	CacheClass     _CacheClass
	CacheClassOnce sync.Once
)

func getCacheClass() _CacheClass {
	CacheClassOnce.Do(func() {
		CacheClass = _CacheClass{objc.GetClass("NSCache")}
	})
	return CacheClass
}

type _CacheClass struct {
	class objc.Class
}

// An interface definition for the [Cache] class.
type ICache interface {
	objectivec.IObject
	// properties:
	CountLimit() uint /* primitive/slice/pointer. */
	SetCountLimit(value uint /* primitive/slice/pointer. */)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	EvictsObjectsWithDiscardedContent() bool /* primitive/slice/pointer. */
	SetEvictsObjectsWithDiscardedContent(value bool /* primitive/slice/pointer. */)
	Name() string /* primitive/slice/pointer. */
	SetName(value string /* primitive/slice/pointer. */)
	TotalCostLimit() uint /* primitive/slice/pointer. */
	SetTotalCostLimit(value uint /* primitive/slice/pointer. */)
	// methods:
	ObjectForKey(key unsafe.Pointer) unsafe.Pointer
	RemoveAllObjects()
	RemoveObjectForKey(key unsafe.Pointer)
	SetObjectForKey(obj unsafe.Pointer, key unsafe.Pointer)
	SetObjectForKeyCost(obj unsafe.Pointer, key unsafe.Pointer, g uint /* primitive/slice/pointer. */)
}

// A mutable collection you use to temporarily store transient key-value pairs that are subject to eviction when resources are low.
//
// Cache objects differ from other mutable collections in a few ways: The class incorporates various auto-eviction policies, which ensure that a cache doesn’t use too much of the system’s memory. If memory is needed by other applications, these policies remove some items from the cache, minimizing its memory footprint. You can add, remove, and query items in the cache from different threads without having to lock the cache yourself. Unlike an object, a cache does not copy the key objects that are put into it. You typically use objects to temporarily store objects with transient data that are expensive to create. Reusing these objects can provide performance benefits, because their values do not have to be recalculated. However, the objects are not critical to the application and can be discarded if memory is tight. If discarded, their values will have to be recomputed again when needed. Objects that have subcomponents that can be discarded when not being used can adopt the protocol to improve cache eviction behavior. By default, objects in a cache are automatically removed if their content is discarded, although this automatic removal policy can be changed. If an object is put into the cache, the cache calls on it upon its removal.


// A mutable collection you use to temporarily store transient key-value pairs that are subject to eviction when resources are low.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache
type Cache struct {
	objectivec.Object
}

// CacheFrom constructs a [Cache] from an unsafe.Pointer.
//
// A mutable collection you use to temporarily store transient key-value pairs that are subject to eviction when resources are low.
func CacheFrom(ptr unsafe.Pointer) Cache {
	return Cache{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CacheClass) Alloc() Cache {
	rv := objc.Send[Cache](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CacheClass) New() Cache {
	rv := objc.Send[Cache](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Cache) Init() Cache {
	rv := objc.Send[Cache](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Cache) Autorelease() Cache {
	rv := objc.Send[Cache](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCache creates a new Cache instance.
func NewCache() Cache {
	return getCacheClass().New()
}



// Returns the value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/object(forKey:)
func (c_ Cache) ObjectForKey(key unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("objectForKey:"), key)
	return rv
}


// Empties the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/removeAllObjects()
func (c_ Cache) RemoveAllObjects() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllObjects"))
}


// Removes the value of the specified key in the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/removeObject(forKey:)
func (c_ Cache) RemoveObjectForKey(key unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeObjectForKey:"), key)
}


// Sets the value of the specified key in the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/setObject(_:forKey:)
func (c_ Cache) SetObjectForKey(obj unsafe.Pointer, key unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObject:forKey:"), obj, key)
}


// Sets the value of the specified key in the cache, and associates the key-value pair with the specified cost.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/setObject(_:forKey:cost:)
func (c_ Cache) SetObjectForKeyCost(obj unsafe.Pointer, key unsafe.Pointer, g uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObject:forKey:cost:"), obj, key, g)
}


// The maximum number of objects the cache should hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/countLimit
func (c_ Cache) CountLimit() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](c_.ID, objc.Sel("countLimit"))
	return rv
}


// The maximum number of objects the cache should hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/countLimit
func (c_ Cache) SetCountLimit(value uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCountLimit:"), value)
}


// The cache’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/delegate
func (c_ Cache) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// The cache’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/delegate
func (c_ Cache) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}


// Whether the cache will automatically evict discardable-content objects whose content has been discarded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/evictsObjectsWithDiscardedContent
func (c_ Cache) EvictsObjectsWithDiscardedContent() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("evictsObjectsWithDiscardedContent"))
	return rv
}


// Whether the cache will automatically evict discardable-content objects whose content has been discarded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/evictsObjectsWithDiscardedContent
func (c_ Cache) SetEvictsObjectsWithDiscardedContent(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEvictsObjectsWithDiscardedContent:"), value)
}


// The name of the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/name
func (c_ Cache) Name() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("name"))
	return rv
}


// The name of the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/name
func (c_ Cache) SetName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), objc.String(value))
}


// The maximum total cost that the cache can hold before it starts evicting objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/totalCostLimit
func (c_ Cache) TotalCostLimit() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](c_.ID, objc.Sel("totalCostLimit"))
	return rv
}


// The maximum total cost that the cache can hold before it starts evicting objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/totalCostLimit
func (c_ Cache) SetTotalCostLimit(value uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTotalCostLimit:"), value)
}



