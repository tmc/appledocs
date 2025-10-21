// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestEventEvent] class.
var (
	MTRTestClusterClusterTestEventEventClass     _MTRTestClusterClusterTestEventEventClass
	MTRTestClusterClusterTestEventEventClassOnce sync.Once
)

func getMTRTestClusterClusterTestEventEventClass() _MTRTestClusterClusterTestEventEventClass {
	MTRTestClusterClusterTestEventEventClassOnce.Do(func() {
		MTRTestClusterClusterTestEventEventClass = _MTRTestClusterClusterTestEventEventClass{objc.GetClass("MTRTestClusterClusterTestEventEvent")}
	})
	return MTRTestClusterClusterTestEventEventClass
}

type _MTRTestClusterClusterTestEventEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestEventEvent] class.
type IMTRTestClusterClusterTestEventEvent interface {
	IMTRUnitTestingClusterTestEventEvent
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestEventEvent
type MTRTestClusterClusterTestEventEvent struct {
	MTRUnitTestingClusterTestEventEvent
}

// MTRTestClusterClusterTestEventEventFrom constructs a [MTRTestClusterClusterTestEventEvent] from an unsafe.Pointer.
func MTRTestClusterClusterTestEventEventFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestEventEvent {
	return MTRTestClusterClusterTestEventEvent{
		MTRUnitTestingClusterTestEventEvent: MTRUnitTestingClusterTestEventEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestEventEventClass) Alloc() MTRTestClusterClusterTestEventEvent {
	rv := objc.Send[MTRTestClusterClusterTestEventEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestEventEventClass) New() MTRTestClusterClusterTestEventEvent {
	rv := objc.Send[MTRTestClusterClusterTestEventEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestEventEvent) Init() MTRTestClusterClusterTestEventEvent {
	rv := objc.Send[MTRTestClusterClusterTestEventEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestEventEvent) Autorelease() MTRTestClusterClusterTestEventEvent {
	rv := objc.Send[MTRTestClusterClusterTestEventEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestEventEvent creates a new MTRTestClusterClusterTestEventEvent instance.
func NewMTRTestClusterClusterTestEventEvent() MTRTestClusterClusterTestEventEvent {
	return getMTRTestClusterClusterTestEventEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertesteventevent/arg3
func (m_ MTRTestClusterClusterTestEventEvent) Arg3() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg3"))
	return rv
}


// SetArg3 sets the value of the arg3 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertesteventevent/arg3
func (m_ MTRTestClusterClusterTestEventEvent) SetArg3(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg3:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertesteventevent/arg6
func (m_ MTRTestClusterClusterTestEventEvent) Arg6() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg6"))
	return rv
}


// SetArg6 sets the value of the arg6 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertesteventevent/arg6
func (m_ MTRTestClusterClusterTestEventEvent) SetArg6(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg6:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertesteventevent/arg2
func (m_ MTRTestClusterClusterTestEventEvent) Arg2() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg2"))
	return rv
}


// SetArg2 sets the value of the arg2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertesteventevent/arg2
func (m_ MTRTestClusterClusterTestEventEvent) SetArg2(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg2:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertesteventevent/arg4
func (m_ MTRTestClusterClusterTestEventEvent) Arg4() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg4"))
	return rv
}


// SetArg4 sets the value of the arg4 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertesteventevent/arg4
func (m_ MTRTestClusterClusterTestEventEvent) SetArg4(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg4:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertesteventevent/arg1
func (m_ MTRTestClusterClusterTestEventEvent) Arg1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("arg1"))
	return rv
}


// SetArg1 sets the value of the arg1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertesteventevent/arg1
func (m_ MTRTestClusterClusterTestEventEvent) SetArg1(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertesteventevent/arg5
func (m_ MTRTestClusterClusterTestEventEvent) Arg5() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arg5"))
	return rv
}


// SetArg5 sets the value of the arg5 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertesteventevent/arg5
func (m_ MTRTestClusterClusterTestEventEvent) SetArg5(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArg5:"), value)
}



