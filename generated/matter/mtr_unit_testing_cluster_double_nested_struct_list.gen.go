// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterDoubleNestedStructList] class.
var (
	MTRUnitTestingClusterDoubleNestedStructListClass     _MTRUnitTestingClusterDoubleNestedStructListClass
	MTRUnitTestingClusterDoubleNestedStructListClassOnce sync.Once
)

func getMTRUnitTestingClusterDoubleNestedStructListClass() _MTRUnitTestingClusterDoubleNestedStructListClass {
	MTRUnitTestingClusterDoubleNestedStructListClassOnce.Do(func() {
		MTRUnitTestingClusterDoubleNestedStructListClass = _MTRUnitTestingClusterDoubleNestedStructListClass{objc.GetClass("MTRUnitTestingClusterDoubleNestedStructList")}
	})
	return MTRUnitTestingClusterDoubleNestedStructListClass
}

type _MTRUnitTestingClusterDoubleNestedStructListClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterDoubleNestedStructList] class.
type IMTRUnitTestingClusterDoubleNestedStructList interface {
	objectivec.IObject
	// properties:
	A() unsafe.Pointer
	SetA(value unsafe.Pointer)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterDoubleNestedStructList
type MTRUnitTestingClusterDoubleNestedStructList struct {
	objectivec.Object
}

// MTRUnitTestingClusterDoubleNestedStructListFrom constructs a [MTRUnitTestingClusterDoubleNestedStructList] from an unsafe.Pointer.
func MTRUnitTestingClusterDoubleNestedStructListFrom(ptr unsafe.Pointer) MTRUnitTestingClusterDoubleNestedStructList {
	return MTRUnitTestingClusterDoubleNestedStructList{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterDoubleNestedStructListClass) Alloc() MTRUnitTestingClusterDoubleNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterDoubleNestedStructList](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterDoubleNestedStructListClass) New() MTRUnitTestingClusterDoubleNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterDoubleNestedStructList](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterDoubleNestedStructList) Init() MTRUnitTestingClusterDoubleNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterDoubleNestedStructList](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterDoubleNestedStructList) Autorelease() MTRUnitTestingClusterDoubleNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterDoubleNestedStructList](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterDoubleNestedStructList creates a new MTRUnitTestingClusterDoubleNestedStructList instance.
func NewMTRUnitTestingClusterDoubleNestedStructList() MTRUnitTestingClusterDoubleNestedStructList {
	return getMTRUnitTestingClusterDoubleNestedStructListClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterdoublenestedstructlist/a
func (m_ MTRUnitTestingClusterDoubleNestedStructList) A() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("a"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusterdoublenestedstructlist/a
func (m_ MTRUnitTestingClusterDoubleNestedStructList) SetA(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}



