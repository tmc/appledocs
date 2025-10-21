// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MapTable] class.
type IMapTable interface {
	objectivec.IObject
	SetObjectForKey(anObject unsafe.Pointer, aKey unsafe.Pointer)
}

// A collection similar to a dictionary, but with a broader range of available memory semantics.
//
// The map table is modeled after with the following differences: Keys and/or values are optionally held “weakly” such that entries are removed when one of the objects is reclaimed. Its keys or values may be copied on input or may use pointer identity for equality and hashing. It can contain arbitrary pointers (its contents are not constrained to being objects). You can configure an instance to operate on arbitrary pointers and not just objects, although typically you are encouraged to use the C function API for void * pointers. The object-based API (such as ) will not work for non-object pointers without type-casting. When configuring map tables, note that only the options listed in guarantee that the rest of the API will work correctly—including copying, archiving, and fast enumeration. While other options are used for certain configurations, such as to hold arbitrary pointers, not all combinations of the options are valid. With some combinations the map table may not work correctly, or may not even be initialized correctly.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MapTableClass) Alloc() MapTable {
	rv := objc.Send[MapTable](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Adds a given key-value pair to the map table.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTable/setObject(_:forKey:)
func (m_ MapTable) SetObjectForKey(anObject unsafe.Pointer, aKey unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObject:forKey:"), anObject, aKey)
}

// The number of key-value pairs in the map table.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/count
func (m_ MapTable) Count() int {
	rv := objc.Send[int](m_.ID, objc.Sel("count"))
	return rv
}


// SetCount sets the value of the count property.
// The number of key-value pairs in the map table.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/count
func (m_ MapTable) SetCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCount:"), value)
}

// The pointer functions the map table uses to manage keys.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/keypointerfunctions
func (m_ MapTable) KeyPointerFunctions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("keyPointerFunctions"))
	return rv
}


// SetKeyPointerFunctions sets the value of the keyPointerFunctions property.
// The pointer functions the map table uses to manage keys.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/keypointerfunctions
func (m_ MapTable) SetKeyPointerFunctions(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setKeyPointerFunctions:"), value)
}

// The pointer functions the map table uses to manage values.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/valuepointerfunctions
func (m_ MapTable) ValuePointerFunctions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("valuePointerFunctions"))
	return rv
}


// SetValuePointerFunctions sets the value of the valuePointerFunctions property.
// The pointer functions the map table uses to manage values.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmaptable/valuepointerfunctions
func (m_ MapTable) SetValuePointerFunctions(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValuePointerFunctions:"), value)
}



