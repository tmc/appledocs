// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	A() foundation.Number
	SetA(value foundation.INumber)
	B() foundation.Number
	SetB(value foundation.INumber)
	C() foundation.Number
	SetC(value foundation.INumber)
	D() foundation.Data
	SetD(value foundation.IData)
	E() string
	SetE(value string)
	F() foundation.Number
	SetF(value foundation.INumber)
	G() foundation.Number
	SetG(value foundation.INumber)
	H() foundation.Number
	SetH(value foundation.INumber)
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/a
func (m_ MTRUnitTestingClusterSimpleStruct) A() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("a"))
	return rv
}


// SetA sets the value of the a property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/a
func (m_ MTRUnitTestingClusterSimpleStruct) SetA(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/b
func (m_ MTRUnitTestingClusterSimpleStruct) B() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("b"))
	return rv
}


// SetB sets the value of the b property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/b
func (m_ MTRUnitTestingClusterSimpleStruct) SetB(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setB:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/c
func (m_ MTRUnitTestingClusterSimpleStruct) C() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("c"))
	return rv
}


// SetC sets the value of the c property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/c
func (m_ MTRUnitTestingClusterSimpleStruct) SetC(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setC:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/d
func (m_ MTRUnitTestingClusterSimpleStruct) D() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("d"))
	return rv
}


// SetD sets the value of the d property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/d
func (m_ MTRUnitTestingClusterSimpleStruct) SetD(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setD:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/e
func (m_ MTRUnitTestingClusterSimpleStruct) E() string {
	rv := objc.Send[string](m_.ID, objc.Sel("e"))
	return rv
}


// SetE sets the value of the e property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/e
func (m_ MTRUnitTestingClusterSimpleStruct) SetE(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setE:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/f
func (m_ MTRUnitTestingClusterSimpleStruct) F() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("f"))
	return rv
}


// SetF sets the value of the f property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/f
func (m_ MTRUnitTestingClusterSimpleStruct) SetF(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setF:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/g
func (m_ MTRUnitTestingClusterSimpleStruct) G() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("g"))
	return rv
}


// SetG sets the value of the g property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/g
func (m_ MTRUnitTestingClusterSimpleStruct) SetG(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setG:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/h
func (m_ MTRUnitTestingClusterSimpleStruct) H() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("h"))
	return rv
}


// SetH sets the value of the h property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/h
func (m_ MTRUnitTestingClusterSimpleStruct) SetH(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setH:"), value)
}



