// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterNestedStruct] class.
var (
	MTRUnitTestingClusterNestedStructClass     _MTRUnitTestingClusterNestedStructClass
	MTRUnitTestingClusterNestedStructClassOnce sync.Once
)

func getMTRUnitTestingClusterNestedStructClass() _MTRUnitTestingClusterNestedStructClass {
	MTRUnitTestingClusterNestedStructClassOnce.Do(func() {
		MTRUnitTestingClusterNestedStructClass = _MTRUnitTestingClusterNestedStructClass{objc.GetClass("MTRUnitTestingClusterNestedStruct")}
	})
	return MTRUnitTestingClusterNestedStructClass
}

type _MTRUnitTestingClusterNestedStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterNestedStruct] class.
type IMTRUnitTestingClusterNestedStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStruct
type MTRUnitTestingClusterNestedStruct struct {
	objectivec.Object
}

// MTRUnitTestingClusterNestedStructFrom constructs a [MTRUnitTestingClusterNestedStruct] from an unsafe.Pointer.
func MTRUnitTestingClusterNestedStructFrom(ptr unsafe.Pointer) MTRUnitTestingClusterNestedStruct {
	return MTRUnitTestingClusterNestedStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterNestedStructClass) Alloc() MTRUnitTestingClusterNestedStruct {
	rv := objc.Send[MTRUnitTestingClusterNestedStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterNestedStructClass) New() MTRUnitTestingClusterNestedStruct {
	rv := objc.Send[MTRUnitTestingClusterNestedStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterNestedStruct) Init() MTRUnitTestingClusterNestedStruct {
	rv := objc.Send[MTRUnitTestingClusterNestedStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterNestedStruct) Autorelease() MTRUnitTestingClusterNestedStruct {
	rv := objc.Send[MTRUnitTestingClusterNestedStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterNestedStruct creates a new MTRUnitTestingClusterNestedStruct instance.
func NewMTRUnitTestingClusterNestedStruct() MTRUnitTestingClusterNestedStruct {
	return getMTRUnitTestingClusterNestedStructClass().New()
}




