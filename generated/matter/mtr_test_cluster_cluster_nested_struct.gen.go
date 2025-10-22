// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterNestedStruct] class.
var (
	MTRTestClusterClusterNestedStructClass     _MTRTestClusterClusterNestedStructClass
	MTRTestClusterClusterNestedStructClassOnce sync.Once
)

func getMTRTestClusterClusterNestedStructClass() _MTRTestClusterClusterNestedStructClass {
	MTRTestClusterClusterNestedStructClassOnce.Do(func() {
		MTRTestClusterClusterNestedStructClass = _MTRTestClusterClusterNestedStructClass{objc.GetClass("MTRTestClusterClusterNestedStruct")}
	})
	return MTRTestClusterClusterNestedStructClass
}

type _MTRTestClusterClusterNestedStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterNestedStruct] class.
type IMTRTestClusterClusterNestedStruct interface {
	IMTRUnitTestingClusterNestedStruct
	A() foundation.Number
	SetA(value foundation.INumber)
	B() foundation.Number
	SetB(value foundation.INumber)
	C() MTRTestClusterClusterSimpleStruct
	SetC(value IMTRTestClusterClusterSimpleStruct)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterNestedStruct
type MTRTestClusterClusterNestedStruct struct {
	MTRUnitTestingClusterNestedStruct
}

// MTRTestClusterClusterNestedStructFrom constructs a [MTRTestClusterClusterNestedStruct] from an unsafe.Pointer.
func MTRTestClusterClusterNestedStructFrom(ptr unsafe.Pointer) MTRTestClusterClusterNestedStruct {
	return MTRTestClusterClusterNestedStruct{
		MTRUnitTestingClusterNestedStruct: MTRUnitTestingClusterNestedStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterNestedStructClass) Alloc() MTRTestClusterClusterNestedStruct {
	rv := objc.Send[MTRTestClusterClusterNestedStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterNestedStructClass) New() MTRTestClusterClusterNestedStruct {
	rv := objc.Send[MTRTestClusterClusterNestedStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterNestedStruct) Init() MTRTestClusterClusterNestedStruct {
	rv := objc.Send[MTRTestClusterClusterNestedStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterNestedStruct) Autorelease() MTRTestClusterClusterNestedStruct {
	rv := objc.Send[MTRTestClusterClusterNestedStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterNestedStruct creates a new MTRTestClusterClusterNestedStruct instance.
func NewMTRTestClusterClusterNestedStruct() MTRTestClusterClusterNestedStruct {
	return getMTRTestClusterClusterNestedStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstruct/a
func (m_ MTRTestClusterClusterNestedStruct) A() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("a"))
	return rv
}


// SetA sets the value of the a property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstruct/a
func (m_ MTRTestClusterClusterNestedStruct) SetA(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstruct/b
func (m_ MTRTestClusterClusterNestedStruct) B() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("b"))
	return rv
}


// SetB sets the value of the b property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstruct/b
func (m_ MTRTestClusterClusterNestedStruct) SetB(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setB:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstruct/c
func (m_ MTRTestClusterClusterNestedStruct) C() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](m_.ID, objc.Sel("c"))
	return rv
}


// SetC sets the value of the c property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstruct/c
func (m_ MTRTestClusterClusterNestedStruct) SetC(value IMTRTestClusterClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setC:"), value)
}



