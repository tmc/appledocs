// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestEventEvent] class.
var (
	MTRUnitTestingClusterTestEventEventClass     _MTRUnitTestingClusterTestEventEventClass
	MTRUnitTestingClusterTestEventEventClassOnce sync.Once
)

func getMTRUnitTestingClusterTestEventEventClass() _MTRUnitTestingClusterTestEventEventClass {
	MTRUnitTestingClusterTestEventEventClassOnce.Do(func() {
		MTRUnitTestingClusterTestEventEventClass = _MTRUnitTestingClusterTestEventEventClass{objc.GetClass("MTRUnitTestingClusterTestEventEvent")}
	})
	return MTRUnitTestingClusterTestEventEventClass
}

type _MTRUnitTestingClusterTestEventEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestEventEvent] class.
type IMTRUnitTestingClusterTestEventEvent interface {
	objectivec.IObject
	Arg1() foundation.Number
	SetArg1(value foundation.INumber)
	Arg2() foundation.Number
	SetArg2(value foundation.INumber)
	Arg3() foundation.Number
	SetArg3(value foundation.INumber)
	Arg4() MTRUnitTestingClusterSimpleStruct
	SetArg4(value IMTRUnitTestingClusterSimpleStruct)
	Arg5() unsafe.Pointer
	SetArg5(value unsafe.Pointer)
	Arg6() unsafe.Pointer
	SetArg6(value unsafe.Pointer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestEventEvent
type MTRUnitTestingClusterTestEventEvent struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestEventEventFrom constructs a [MTRUnitTestingClusterTestEventEvent] from an unsafe.Pointer.
func MTRUnitTestingClusterTestEventEventFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestEventEvent {
	return MTRUnitTestingClusterTestEventEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestEventEventClass) Alloc() MTRUnitTestingClusterTestEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestEventEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestEventEventClass) New() MTRUnitTestingClusterTestEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestEventEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestEventEvent) Init() MTRUnitTestingClusterTestEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestEventEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestEventEvent) Autorelease() MTRUnitTestingClusterTestEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestEventEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestEventEvent creates a new MTRUnitTestingClusterTestEventEvent instance.
func NewMTRUnitTestingClusterTestEventEvent() MTRUnitTestingClusterTestEventEvent {
	return getMTRUnitTestingClusterTestEventEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertesteventevent/arg1
func (m_ MTRUnitTestingClusterTestEventEvent) Arg1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertesteventevent/arg1
func (m_ MTRUnitTestingClusterTestEventEvent) SetArg1(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertesteventevent/arg2
func (m_ MTRUnitTestingClusterTestEventEvent) Arg2() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg2"))
	return rv
}


// SetArg2 sets the value of the arg2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertesteventevent/arg2
func (m_ MTRUnitTestingClusterTestEventEvent) SetArg2(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertesteventevent/arg3
func (m_ MTRUnitTestingClusterTestEventEvent) Arg3() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg3"))
	return rv
}


// SetArg3 sets the value of the arg3 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertesteventevent/arg3
func (m_ MTRUnitTestingClusterTestEventEvent) SetArg3(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg3:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertesteventevent/arg4
func (m_ MTRUnitTestingClusterTestEventEvent) Arg4() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("arg4"))
	return rv
}


// SetArg4 sets the value of the arg4 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertesteventevent/arg4
func (m_ MTRUnitTestingClusterTestEventEvent) SetArg4(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg4:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertesteventevent/arg5
func (m_ MTRUnitTestingClusterTestEventEvent) Arg5() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg5"))
	return rv
}


// SetArg5 sets the value of the arg5 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertesteventevent/arg5
func (m_ MTRUnitTestingClusterTestEventEvent) SetArg5(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg5:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertesteventevent/arg6
func (m_ MTRUnitTestingClusterTestEventEvent) Arg6() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg6"))
	return rv
}


// SetArg6 sets the value of the arg6 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertesteventevent/arg6
func (m_ MTRUnitTestingClusterTestEventEvent) SetArg6(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg6:"), value)
}



