// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MapTable] class.
var mapTableClass = _MapTableClass{objc.GetClass("NSMapTable")}

type _MapTableClass struct {
	class objc.Class
}

// An interface definition for the [MapTable] class.
type IMapTable interface {
	objectivec.IObject
}

// A collection similar to a dictionary, but with a broader range of available memory semantics. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return mapTableClass.New()
}




