// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterSimpleStruct] class.
var (
	MTRTestClusterClusterSimpleStructClass     _MTRTestClusterClusterSimpleStructClass
	MTRTestClusterClusterSimpleStructClassOnce sync.Once
)

func getMTRTestClusterClusterSimpleStructClass() _MTRTestClusterClusterSimpleStructClass {
	MTRTestClusterClusterSimpleStructClassOnce.Do(func() {
		MTRTestClusterClusterSimpleStructClass = _MTRTestClusterClusterSimpleStructClass{objc.GetClass("MTRTestClusterClusterSimpleStruct")}
	})
	return MTRTestClusterClusterSimpleStructClass
}

type _MTRTestClusterClusterSimpleStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterSimpleStruct] class.
type IMTRTestClusterClusterSimpleStruct interface {
	IMTRUnitTestingClusterSimpleStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterSimpleStruct
type MTRTestClusterClusterSimpleStruct struct {
	MTRUnitTestingClusterSimpleStruct
}

// MTRTestClusterClusterSimpleStructFrom constructs a [MTRTestClusterClusterSimpleStruct] from an unsafe.Pointer.
func MTRTestClusterClusterSimpleStructFrom(ptr unsafe.Pointer) MTRTestClusterClusterSimpleStruct {
	return MTRTestClusterClusterSimpleStruct{
		MTRUnitTestingClusterSimpleStruct: MTRUnitTestingClusterSimpleStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterSimpleStructClass) Alloc() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterSimpleStructClass) New() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterSimpleStruct) Init() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterSimpleStruct) Autorelease() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterSimpleStruct creates a new MTRTestClusterClusterSimpleStruct instance.
func NewMTRTestClusterClusterSimpleStruct() MTRTestClusterClusterSimpleStruct {
	return getMTRTestClusterClusterSimpleStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/a
func (m_ MTRTestClusterClusterSimpleStruct) A() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("a"))
	return rv
}


// SetA sets the value of the a property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/a
func (m_ MTRTestClusterClusterSimpleStruct) SetA(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/b
func (m_ MTRTestClusterClusterSimpleStruct) B() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("b"))
	return rv
}


// SetB sets the value of the b property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/b
func (m_ MTRTestClusterClusterSimpleStruct) SetB(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setB:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/c
func (m_ MTRTestClusterClusterSimpleStruct) C() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("c"))
	return rv
}


// SetC sets the value of the c property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/c
func (m_ MTRTestClusterClusterSimpleStruct) SetC(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setC:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/d
func (m_ MTRTestClusterClusterSimpleStruct) D() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("d"))
	return rv
}


// SetD sets the value of the d property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/d
func (m_ MTRTestClusterClusterSimpleStruct) SetD(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setD:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/e
func (m_ MTRTestClusterClusterSimpleStruct) E() string {
	rv := objc.Send[string](m_.ID, objc.Sel("e"))
	return rv
}


// SetE sets the value of the e property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/e
func (m_ MTRTestClusterClusterSimpleStruct) SetE(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setE:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/f
func (m_ MTRTestClusterClusterSimpleStruct) F() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("f"))
	return rv
}


// SetF sets the value of the f property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/f
func (m_ MTRTestClusterClusterSimpleStruct) SetF(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setF:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/g
func (m_ MTRTestClusterClusterSimpleStruct) G() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("g"))
	return rv
}


// SetG sets the value of the g property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/g
func (m_ MTRTestClusterClusterSimpleStruct) SetG(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setG:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/h
func (m_ MTRTestClusterClusterSimpleStruct) H() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("h"))
	return rv
}


// SetH sets the value of the h property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/h
func (m_ MTRTestClusterClusterSimpleStruct) SetH(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setH:"), value)
}



