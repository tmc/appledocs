// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterDoubleNestedStructList] class.
var (
	MTRTestClusterClusterDoubleNestedStructListClass     _MTRTestClusterClusterDoubleNestedStructListClass
	MTRTestClusterClusterDoubleNestedStructListClassOnce sync.Once
)

func getMTRTestClusterClusterDoubleNestedStructListClass() _MTRTestClusterClusterDoubleNestedStructListClass {
	MTRTestClusterClusterDoubleNestedStructListClassOnce.Do(func() {
		MTRTestClusterClusterDoubleNestedStructListClass = _MTRTestClusterClusterDoubleNestedStructListClass{objc.GetClass("MTRTestClusterClusterDoubleNestedStructList")}
	})
	return MTRTestClusterClusterDoubleNestedStructListClass
}

type _MTRTestClusterClusterDoubleNestedStructListClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterDoubleNestedStructList] class.
type IMTRTestClusterClusterDoubleNestedStructList interface {
	IMTRUnitTestingClusterDoubleNestedStructList
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterDoubleNestedStructList
type MTRTestClusterClusterDoubleNestedStructList struct {
	MTRUnitTestingClusterDoubleNestedStructList
}

// MTRTestClusterClusterDoubleNestedStructListFrom constructs a [MTRTestClusterClusterDoubleNestedStructList] from an unsafe.Pointer.
func MTRTestClusterClusterDoubleNestedStructListFrom(ptr unsafe.Pointer) MTRTestClusterClusterDoubleNestedStructList {
	return MTRTestClusterClusterDoubleNestedStructList{
		MTRUnitTestingClusterDoubleNestedStructList: MTRUnitTestingClusterDoubleNestedStructListFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterDoubleNestedStructListClass) Alloc() MTRTestClusterClusterDoubleNestedStructList {
	rv := objc.Send[MTRTestClusterClusterDoubleNestedStructList](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterDoubleNestedStructListClass) New() MTRTestClusterClusterDoubleNestedStructList {
	rv := objc.Send[MTRTestClusterClusterDoubleNestedStructList](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterDoubleNestedStructList) Init() MTRTestClusterClusterDoubleNestedStructList {
	rv := objc.Send[MTRTestClusterClusterDoubleNestedStructList](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterDoubleNestedStructList) Autorelease() MTRTestClusterClusterDoubleNestedStructList {
	rv := objc.Send[MTRTestClusterClusterDoubleNestedStructList](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterDoubleNestedStructList creates a new MTRTestClusterClusterDoubleNestedStructList instance.
func NewMTRTestClusterClusterDoubleNestedStructList() MTRTestClusterClusterDoubleNestedStructList {
	return getMTRTestClusterClusterDoubleNestedStructListClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterdoublenestedstructlist/a
func (m_ MTRTestClusterClusterDoubleNestedStructList) A() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("a"))
	return rv
}


// SetA sets the value of the a property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusterdoublenestedstructlist/a
func (m_ MTRTestClusterClusterDoubleNestedStructList) SetA(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}



