// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRServiceAreaClusterMapStruct] class.
var (
	MTRServiceAreaClusterMapStructClass     _MTRServiceAreaClusterMapStructClass
	MTRServiceAreaClusterMapStructClassOnce sync.Once
)

func getMTRServiceAreaClusterMapStructClass() _MTRServiceAreaClusterMapStructClass {
	MTRServiceAreaClusterMapStructClassOnce.Do(func() {
		MTRServiceAreaClusterMapStructClass = _MTRServiceAreaClusterMapStructClass{objc.GetClass("MTRServiceAreaClusterMapStruct")}
	})
	return MTRServiceAreaClusterMapStructClass
}

type _MTRServiceAreaClusterMapStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRServiceAreaClusterMapStruct] class.
type IMTRServiceAreaClusterMapStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterMapStruct
type MTRServiceAreaClusterMapStruct struct {
	objectivec.Object
}

// MTRServiceAreaClusterMapStructFrom constructs a [MTRServiceAreaClusterMapStruct] from an unsafe.Pointer.
func MTRServiceAreaClusterMapStructFrom(ptr unsafe.Pointer) MTRServiceAreaClusterMapStruct {
	return MTRServiceAreaClusterMapStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterMapStructClass) Alloc() MTRServiceAreaClusterMapStruct {
	rv := objc.Send[MTRServiceAreaClusterMapStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServiceAreaClusterMapStructClass) New() MTRServiceAreaClusterMapStruct {
	rv := objc.Send[MTRServiceAreaClusterMapStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterMapStruct) Init() MTRServiceAreaClusterMapStruct {
	rv := objc.Send[MTRServiceAreaClusterMapStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterMapStruct) Autorelease() MTRServiceAreaClusterMapStruct {
	rv := objc.Send[MTRServiceAreaClusterMapStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterMapStruct creates a new MTRServiceAreaClusterMapStruct instance.
func NewMTRServiceAreaClusterMapStruct() MTRServiceAreaClusterMapStruct {
	return getMTRServiceAreaClusterMapStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterMapStruct/mapID
func (m_ MTRServiceAreaClusterMapStruct) MapID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mapID"))
	return rv
}


// SetMapID sets the value of the mapID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterMapStruct/mapID
func (m_ MTRServiceAreaClusterMapStruct) SetMapID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapID:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterMapStruct/name
func (m_ MTRServiceAreaClusterMapStruct) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterMapStruct/name
func (m_ MTRServiceAreaClusterMapStruct) SetName(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}


