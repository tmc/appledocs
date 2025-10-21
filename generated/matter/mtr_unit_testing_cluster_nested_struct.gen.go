// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstruct/c
func (m_ MTRUnitTestingClusterNestedStruct) C() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("c"))
	return rv
}


// SetC sets the value of the c property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstruct/c
func (m_ MTRUnitTestingClusterNestedStruct) SetC(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setC:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstruct/b
func (m_ MTRUnitTestingClusterNestedStruct) B() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("b"))
	return rv
}


// SetB sets the value of the b property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstruct/b
func (m_ MTRUnitTestingClusterNestedStruct) SetB(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setB:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstruct/a
func (m_ MTRUnitTestingClusterNestedStruct) A() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("a"))
	return rv
}


// SetA sets the value of the a property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstruct/a
func (m_ MTRUnitTestingClusterNestedStruct) SetA(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}



