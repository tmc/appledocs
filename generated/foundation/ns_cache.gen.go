// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCache */


/* debug [class_header]: Header for NSCache */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Cache */
// An interface definition for the [Cache] class.
type ICache interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Cache */
	// properties:
	CountLimit() uint
	SetCountLimit(value uint)
	EvictsObjectsWithDiscardedContent() bool
	SetEvictsObjectsWithDiscardedContent(value bool)
	Name() IString
	SetName(value IString)
	TotalCostLimit() uint
	SetTotalCostLimit(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Cache */
	// methods:
	ObjectForKey(key objectivec.IObject) objectivec.IObject
	RemoveAllObjects()
	RemoveObjectForKey(key objectivec.IObject)
	SetObjectForKey(obj objectivec.IObject, key objectivec.IObject)
	SetObjectForKeyCost(obj objectivec.IObject, key objectivec.IObject, g uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Cache */
// Alloc allocates a new instance without initialization.
func (cc _CacheClass) Alloc() Cache {
	rv := objc.Send[Cache](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Cache */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Cache *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Cache */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Cache */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Cache */

// Returns the value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/object(forKey:)
func (c_ Cache) ObjectForKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("objectForKey:"), key)
	return rv
}/* debug [instance_methods/method]: ObjectForKey */


// Empties the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/removeAllObjects()
func (c_ Cache) RemoveAllObjects() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllObjects"))
}/* debug [instance_methods/method]: RemoveAllObjects */


// Removes the value of the specified key in the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/removeObject(forKey:)
func (c_ Cache) RemoveObjectForKey(key objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeObjectForKey:"), key)
}/* debug [instance_methods/method]: RemoveObjectForKey */


// Sets the value of the specified key in the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/setObject(_:forKey:)
func (c_ Cache) SetObjectForKey(obj objectivec.IObject, key objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObject:forKey:"), obj, key)
}/* debug [instance_methods/method]: SetObjectForKey */


// Sets the value of the specified key in the cache, and associates the key-value pair with the specified cost.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/setObject(_:forKey:cost:)
func (c_ Cache) SetObjectForKeyCost(obj objectivec.IObject, key objectivec.IObject, g uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setObject:forKey:cost:"), obj, key, g)
}/* debug [instance_methods/method]: SetObjectForKeyCost */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Cache */

// The maximum number of objects the cache should hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/countLimit
func (c_ Cache) CountLimit() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("countLimit"))
	return rv
}/* debug [instance_properties/getter]: countLimit */


// The maximum number of objects the cache should hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/countLimit
func (c_ Cache) SetCountLimit(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCountLimit:"), value)
}/* debug [instance_properties/setter]: countLimit */


// Whether the cache will automatically evict discardable-content objects whose content has been discarded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/evictsObjectsWithDiscardedContent
func (c_ Cache) EvictsObjectsWithDiscardedContent() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("evictsObjectsWithDiscardedContent"))
	return rv
}/* debug [instance_properties/getter]: evictsObjectsWithDiscardedContent */


// Whether the cache will automatically evict discardable-content objects whose content has been discarded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/evictsObjectsWithDiscardedContent
func (c_ Cache) SetEvictsObjectsWithDiscardedContent(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEvictsObjectsWithDiscardedContent:"), value)
}/* debug [instance_properties/setter]: evictsObjectsWithDiscardedContent */


// The name of the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/name
func (c_ Cache) Name() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/name
func (c_ Cache) SetName(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// The maximum total cost that the cache can hold before it starts evicting objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/totalCostLimit
func (c_ Cache) TotalCostLimit() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("totalCostLimit"))
	return rv
}/* debug [instance_properties/getter]: totalCostLimit */


// The maximum total cost that the cache can hold before it starts evicting objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCache/totalCostLimit
func (c_ Cache) SetTotalCostLimit(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTotalCostLimit:"), value)
}/* debug [instance_properties/setter]: totalCostLimit */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCache */



