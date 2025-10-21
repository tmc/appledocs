// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRUnitTestingClusterSimpleStruct] class.
var (
	MTRUnitTestingClusterSimpleStructClass     _MTRUnitTestingClusterSimpleStructClass
	MTRUnitTestingClusterSimpleStructClassOnce sync.Once
)

func getMTRUnitTestingClusterSimpleStructClass() _MTRUnitTestingClusterSimpleStructClass {
	MTRUnitTestingClusterSimpleStructClassOnce.Do(func() {
		MTRUnitTestingClusterSimpleStructClass = _MTRUnitTestingClusterSimpleStructClass{objc.GetClass("MTRUnitTestingClusterSimpleStruct")}
	})
	return MTRUnitTestingClusterSimpleStructClass
}

type _MTRUnitTestingClusterSimpleStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterSimpleStruct] class.
type IMTRUnitTestingClusterSimpleStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct
type MTRUnitTestingClusterSimpleStruct struct {
	objectivec.Object
}

// MTRUnitTestingClusterSimpleStructFrom constructs a [MTRUnitTestingClusterSimpleStruct] from an unsafe.Pointer.
func MTRUnitTestingClusterSimpleStructFrom(ptr unsafe.Pointer) MTRUnitTestingClusterSimpleStruct {
	return MTRUnitTestingClusterSimpleStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterSimpleStructClass) Alloc() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterSimpleStructClass) New() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterSimpleStruct) Init() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterSimpleStruct) Autorelease() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterSimpleStruct creates a new MTRUnitTestingClusterSimpleStruct instance.
func NewMTRUnitTestingClusterSimpleStruct() MTRUnitTestingClusterSimpleStruct {
	return getMTRUnitTestingClusterSimpleStructClass().New()
}




