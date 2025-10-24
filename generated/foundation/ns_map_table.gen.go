// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMapTable */


/* debug [class_header]: Header for NSMapTable */
// The class instance for the [MapTable] class.
var (
	MapTableClass     _MapTableClass
	MapTableClassOnce sync.Once
)

func getMapTableClass() _MapTableClass {
	MapTableClassOnce.Do(func() {
		MapTableClass = _MapTableClass{objc.GetClass("NSMapTable")}
	})
	return MapTableClass
}

type _MapTableClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MapTable */
// An interface definition for the [MapTable] class.
type IMapTable interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MapTable */
	// properties:
	Count() uint
	KeyPointerFunctions() IPointerFunctions
	ValuePointerFunctions() IPointerFunctions
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MapTable */
	// methods:
	DictionaryRepresentation() IDictionary
	KeyEnumerator() unsafe.Pointer
	ObjectForKey(aKey objectivec.IObject) objectivec.IObject
	ObjectEnumerator() unsafe.Pointer
	RemoveAllObjects()
	RemoveObjectForKey(aKey objectivec.IObject)
	SetObjectForKey(anObject objectivec.IObject, aKey objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MapTable */
// Alloc allocates a new instance without initialization.
func (mc _MapTableClass) Alloc() MapTable {
	rv := objc.Send[MapTable](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MapTableClass) New() MapTable {
	rv := objc.Send[MapTable](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MapTable) Init() MapTable {
	rv := objc.Send[MapTable](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MapTable) Autorelease() MapTable {
	rv := objc.Send[MapTable](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMapTable creates a new MapTable instance.
func NewMapTable() MapTable {
	return getMapTableClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MapTable */
// A collection similar to a dictionary, but with a broader range of available memory semantics.
//
// The map table is modeled after with the following differences: Keys and/or values are optionally held “weakly” such that entries are removed when one of the objects is reclaimed. Its keys or values may be copied on input or may use pointer identity for equality and hashing. It can contain arbitrary pointers (its contents are not constrained to being objects). You can configure an instance to operate on arbitrary pointers and not just objects, although typically you are encouraged to use the C function API for void * pointers. The object-based API (such as ) will not work for non-object pointers without type-casting. When configuring map tables, note that only the options listed in guarantee that the rest of the API will work correctly—including copying, archiving, and fast enumeration. While other options are used for certain configurations, such as to hold arbitrary pointers, not all combinations of the options are valid. With some combinations the map table may not work correctly, or may not even be initialized correctly.


// A collection similar to a dictionary, but with a broader range of available memory semantics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable
type MapTable struct {
	objectivec.Object
}

// MapTableFrom constructs a [MapTable] from an unsafe.Pointer.
//
// A collection similar to a dictionary, but with a broader range of available memory semantics.
func MapTableFrom(ptr unsafe.Pointer) MapTable {
	return MapTable{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MapTable */

// Returns a new map table, initialized with the given options
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/init(keyOptions:valueOptions:)
func NewMapTableWithKeyOptionsValueOptions(keyOptions PointerFunctionsOptions, valueOptions PointerFunctionsOptions) MapTable {
	rv := objc.Send[MapTable](objc.ID(getMapTableClass().class), objc.Sel("mapTableWithKeyOptions:valueOptions:"), keyOptions, valueOptions)
	return rv
}/* debug [class_init_methods/constructor]: NewMapTableWithKeyOptionsValueOptions */


// Returns a map table, initialized with the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/init(keyOptions:valueOptions:capacity:)
func NewMapTableWithKeyOptionsValueOptionsCapacity(keyOptions PointerFunctionsOptions, valueOptions PointerFunctionsOptions, initialCapacity uint) MapTable {
	instance := getMapTableClass().Alloc()
	rv := objc.Send[MapTable](instance.ID, objc.Sel("initWithKeyOptions:valueOptions:capacity:"), keyOptions, valueOptions, initialCapacity)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMapTableWithKeyOptionsValueOptionsCapacity */


// Returns a map table, initialized with the given functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/init(keyPointerFunctions:valuePointerFunctions:capacity:)
func NewMapTableWithKeyPointerFunctionsValuePointerFunctionsCapacity(keyFunctions IPointerFunctions, valueFunctions IPointerFunctions, initialCapacity uint) MapTable {
	instance := getMapTableClass().Alloc()
	rv := objc.Send[MapTable](instance.ID, objc.Sel("initWithKeyPointerFunctions:valuePointerFunctions:capacity:"), keyFunctions, valueFunctions, initialCapacity)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMapTableWithKeyPointerFunctionsValuePointerFunctionsCapacity */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MapTable */

// Returns a new map table, initialized with the given options
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/init(keyOptions:valueOptions:)
func (mc _MapTableClass) MapTableWithKeyOptionsValueOptions(keyOptions PointerFunctionsOptions, valueOptions PointerFunctionsOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("mapTableWithKeyOptions:valueOptions:"), keyOptions, valueOptions)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MapTableWithKeyOptionsValueOptions) */


// Returns a new map table object which has strong references to the keys and values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/mapTableWithStrongToStrongObjects
func (mc _MapTableClass) MapTableWithStrongToStrongObjects() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("mapTableWithStrongToStrongObjects"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MapTableWithStrongToStrongObjects) */


// Returns a new map table object which has strong references to the keys and weak references to the values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/mapTableWithStrongToWeakObjects
func (mc _MapTableClass) MapTableWithStrongToWeakObjects() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("mapTableWithStrongToWeakObjects"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MapTableWithStrongToWeakObjects) */


// Returns a new map table object which has weak references to the keys and strong references to the values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/mapTableWithWeakToStrongObjects
func (mc _MapTableClass) MapTableWithWeakToStrongObjects() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("mapTableWithWeakToStrongObjects"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MapTableWithWeakToStrongObjects) */


// Returns a new map table object which has weak references to the keys and values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/mapTableWithWeakToWeakObjects
func (mc _MapTableClass) MapTableWithWeakToWeakObjects() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("mapTableWithWeakToWeakObjects"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MapTableWithWeakToWeakObjects) */


// Returns a new map table object which has strong references to the keys and values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/strongToStrongObjects()
func (mc _MapTableClass) StrongToStrongObjectsMapTable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("strongToStrongObjectsMapTable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StrongToStrongObjectsMapTable) */


// Returns a new map table object which has strong references to the keys and weak references to the values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/strongToWeakObjects()
func (mc _MapTableClass) StrongToWeakObjectsMapTable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("strongToWeakObjectsMapTable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StrongToWeakObjectsMapTable) */


// Returns a new map table object which has weak references to the keys and strong references to the values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/weakToStrongObjects()
func (mc _MapTableClass) WeakToStrongObjectsMapTable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("weakToStrongObjectsMapTable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WeakToStrongObjectsMapTable) */


// Returns a new map table object which has weak references to the keys and values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/weakToWeakObjects()
func (mc _MapTableClass) WeakToWeakObjectsMapTable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("weakToWeakObjectsMapTable"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WeakToWeakObjectsMapTable) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MapTable */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MapTable */

// Returns a dictionary representation of the map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/dictionaryRepresentation()
func (m_ MapTable) DictionaryRepresentation() IDictionary {
	rv := objc.Send[Dictionary](m_.ID, objc.Sel("dictionaryRepresentation"))
	return rv
}/* debug [instance_methods/method]: DictionaryRepresentation */


// Returns an enumerator object that lets you access each key in the map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/keyEnumerator()
func (m_ MapTable) KeyEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("keyEnumerator"))
	return rv
}/* debug [instance_methods/method]: KeyEnumerator */


// Returns a the value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/object(forKey:)
func (m_ MapTable) ObjectForKey(aKey objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("objectForKey:"), aKey)
	return rv
}/* debug [instance_methods/method]: ObjectForKey */


// Returns an enumerator object that lets you access each value in the map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/objectEnumerator()
func (m_ MapTable) ObjectEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("objectEnumerator"))
	return rv
}/* debug [instance_methods/method]: ObjectEnumerator */


// Empties the map table of its entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/removeAllObjects()
func (m_ MapTable) RemoveAllObjects() {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeAllObjects"))
}/* debug [instance_methods/method]: RemoveAllObjects */


// Removes a given key and its associated value from the map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/removeObject(forKey:)
func (m_ MapTable) RemoveObjectForKey(aKey objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectForKey:"), aKey)
}/* debug [instance_methods/method]: RemoveObjectForKey */


// Adds a given key-value pair to the map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/setObject(_:forKey:)
func (m_ MapTable) SetObjectForKey(anObject objectivec.IObject, aKey objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:forKey:"), anObject, aKey)
}/* debug [instance_methods/method]: SetObjectForKey */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MapTable */

// The number of key-value pairs in the map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/count
func (m_ MapTable) Count() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("count"))
	return rv
}/* debug [instance_properties/getter]: count */


// The pointer functions the map table uses to manage keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/keyPointerFunctions
func (m_ MapTable) KeyPointerFunctions() IPointerFunctions {
	rv := objc.Send[PointerFunctions](m_.ID, objc.Sel("keyPointerFunctions"))
	return rv
}/* debug [instance_properties/getter]: keyPointerFunctions */


// The pointer functions the map table uses to manage values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/valuePointerFunctions
func (m_ MapTable) ValuePointerFunctions() IPointerFunctions {
	rv := objc.Send[PointerFunctions](m_.ID, objc.Sel("valuePointerFunctions"))
	return rv
}/* debug [instance_properties/getter]: valuePointerFunctions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMapTable */


