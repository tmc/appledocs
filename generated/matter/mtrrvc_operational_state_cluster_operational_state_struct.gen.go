// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCOperationalStateClusterOperationalStateStruct] class.
var (
	MTRRVCOperationalStateClusterOperationalStateStructClass     _MTRRVCOperationalStateClusterOperationalStateStructClass
	MTRRVCOperationalStateClusterOperationalStateStructClassOnce sync.Once
)

func getMTRRVCOperationalStateClusterOperationalStateStructClass() _MTRRVCOperationalStateClusterOperationalStateStructClass {
	MTRRVCOperationalStateClusterOperationalStateStructClassOnce.Do(func() {
		MTRRVCOperationalStateClusterOperationalStateStructClass = _MTRRVCOperationalStateClusterOperationalStateStructClass{objc.GetClass("MTRRVCOperationalStateClusterOperationalStateStruct")}
	})
	return MTRRVCOperationalStateClusterOperationalStateStructClass
}

type _MTRRVCOperationalStateClusterOperationalStateStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCOperationalStateClusterOperationalStateStruct] class.
type IMTRRVCOperationalStateClusterOperationalStateStruct interface {
	objectivec.IObject
	OperationalStateID() foundation.Number
	SetOperationalStateID(value foundation.INumber)
	OperationalStateLabel() string
	SetOperationalStateLabel(value string)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterOperationalStateStruct
type MTRRVCOperationalStateClusterOperationalStateStruct struct {
	objectivec.Object
}

// MTRRVCOperationalStateClusterOperationalStateStructFrom constructs a [MTRRVCOperationalStateClusterOperationalStateStruct] from an unsafe.Pointer.
func MTRRVCOperationalStateClusterOperationalStateStructFrom(ptr unsafe.Pointer) MTRRVCOperationalStateClusterOperationalStateStruct {
	return MTRRVCOperationalStateClusterOperationalStateStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCOperationalStateClusterOperationalStateStructClass) Alloc() MTRRVCOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalStateStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCOperationalStateClusterOperationalStateStructClass) New() MTRRVCOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalStateStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCOperationalStateClusterOperationalStateStruct) Init() MTRRVCOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalStateStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCOperationalStateClusterOperationalStateStruct) Autorelease() MTRRVCOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalStateStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCOperationalStateClusterOperationalStateStruct creates a new MTRRVCOperationalStateClusterOperationalStateStruct instance.
func NewMTRRVCOperationalStateClusterOperationalStateStruct() MTRRVCOperationalStateClusterOperationalStateStruct {
	return getMTRRVCOperationalStateClusterOperationalStateStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationalstatestruct/operationalstateid
func (m_ MTRRVCOperationalStateClusterOperationalStateStruct) OperationalStateID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("operationalStateID"))
	return rv
}


// SetOperationalStateID sets the value of the operationalStateID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationalstatestruct/operationalstateid
func (m_ MTRRVCOperationalStateClusterOperationalStateStruct) SetOperationalStateID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalStateID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationalstatestruct/operationalstatelabel
func (m_ MTRRVCOperationalStateClusterOperationalStateStruct) OperationalStateLabel() string {
	rv := objc.Send[string](m_.ID, objc.Sel("operationalStateLabel"))
	return rv
}


// SetOperationalStateLabel sets the value of the operationalStateLabel property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationalstatestruct/operationalstatelabel
func (m_ MTRRVCOperationalStateClusterOperationalStateStruct) SetOperationalStateLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalStateLabel:"), objc.String(value))
}



