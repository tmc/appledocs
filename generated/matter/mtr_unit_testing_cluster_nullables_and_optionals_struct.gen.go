// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterNullablesAndOptionalsStruct] class.
var (
	MTRUnitTestingClusterNullablesAndOptionalsStructClass     _MTRUnitTestingClusterNullablesAndOptionalsStructClass
	MTRUnitTestingClusterNullablesAndOptionalsStructClassOnce sync.Once
)

func getMTRUnitTestingClusterNullablesAndOptionalsStructClass() _MTRUnitTestingClusterNullablesAndOptionalsStructClass {
	MTRUnitTestingClusterNullablesAndOptionalsStructClassOnce.Do(func() {
		MTRUnitTestingClusterNullablesAndOptionalsStructClass = _MTRUnitTestingClusterNullablesAndOptionalsStructClass{objc.GetClass("MTRUnitTestingClusterNullablesAndOptionalsStruct")}
	})
	return MTRUnitTestingClusterNullablesAndOptionalsStructClass
}

type _MTRUnitTestingClusterNullablesAndOptionalsStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterNullablesAndOptionalsStruct] class.
type IMTRUnitTestingClusterNullablesAndOptionalsStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct
type MTRUnitTestingClusterNullablesAndOptionalsStruct struct {
	objectivec.Object
}

// MTRUnitTestingClusterNullablesAndOptionalsStructFrom constructs a [MTRUnitTestingClusterNullablesAndOptionalsStruct] from an unsafe.Pointer.
func MTRUnitTestingClusterNullablesAndOptionalsStructFrom(ptr unsafe.Pointer) MTRUnitTestingClusterNullablesAndOptionalsStruct {
	return MTRUnitTestingClusterNullablesAndOptionalsStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterNullablesAndOptionalsStructClass) Alloc() MTRUnitTestingClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRUnitTestingClusterNullablesAndOptionalsStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterNullablesAndOptionalsStructClass) New() MTRUnitTestingClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRUnitTestingClusterNullablesAndOptionalsStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) Init() MTRUnitTestingClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRUnitTestingClusterNullablesAndOptionalsStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) Autorelease() MTRUnitTestingClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRUnitTestingClusterNullablesAndOptionalsStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterNullablesAndOptionalsStruct creates a new MTRUnitTestingClusterNullablesAndOptionalsStruct instance.
func NewMTRUnitTestingClusterNullablesAndOptionalsStruct() MTRUnitTestingClusterNullablesAndOptionalsStruct {
	return getMTRUnitTestingClusterNullablesAndOptionalsStructClass().New()
}




