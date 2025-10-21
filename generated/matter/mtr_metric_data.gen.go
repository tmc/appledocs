// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMetricData] class.
var (
	MTRMetricDataClass     _MTRMetricDataClass
	MTRMetricDataClassOnce sync.Once
)

func getMTRMetricDataClass() _MTRMetricDataClass {
	MTRMetricDataClassOnce.Do(func() {
		MTRMetricDataClass = _MTRMetricDataClass{objc.GetClass("MTRMetricData")}
	})
	return MTRMetricDataClass
}

type _MTRMetricDataClass struct {
	class objc.Class
}

// An interface definition for the [MTRMetricData] class.
type IMTRMetricData interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMetricData
type MTRMetricData struct {
	objectivec.Object
}

// MTRMetricDataFrom constructs a [MTRMetricData] from an unsafe.Pointer.
func MTRMetricDataFrom(ptr unsafe.Pointer) MTRMetricData {
	return MTRMetricData{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMetricDataClass) Alloc() MTRMetricData {
	rv := objc.Send[MTRMetricData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMetricDataClass) New() MTRMetricData {
	rv := objc.Send[MTRMetricData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMetricData) Init() MTRMetricData {
	rv := objc.Send[MTRMetricData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMetricData) Autorelease() MTRMetricData {
	rv := objc.Send[MTRMetricData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMetricData creates a new MTRMetricData instance.
func NewMTRMetricData() MTRMetricData {
	return getMTRMetricDataClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmetricdata/duration
func (m_ MTRMetricData) Duration() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmetricdata/duration
func (m_ MTRMetricData) SetDuration(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmetricdata/errorcode
func (m_ MTRMetricData) ErrorCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("errorCode"))
	return rv
}


// SetErrorCode sets the value of the errorCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmetricdata/errorcode
func (m_ MTRMetricData) SetErrorCode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmetricdata/value
func (m_ MTRMetricData) Value() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmetricdata/value
func (m_ MTRMetricData) SetValue(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



