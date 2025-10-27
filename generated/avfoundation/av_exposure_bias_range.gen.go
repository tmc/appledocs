// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ExposureBiasRange] class.
var (
	ExposureBiasRangeClass     _ExposureBiasRangeClass
	ExposureBiasRangeClassOnce sync.Once
)

func getExposureBiasRangeClass() _ExposureBiasRangeClass {
	ExposureBiasRangeClassOnce.Do(func() {
		ExposureBiasRangeClass = _ExposureBiasRangeClass{objc.GetClass("AVExposureBiasRange")}
	})
	return ExposureBiasRangeClass
}

type _ExposureBiasRangeClass struct {
	class objc.Class
}





// An interface definition for the [ExposureBiasRange] class.
type IExposureBiasRange interface {
	objectivec.IObject
	

	// properties:
	MaxExposureBias() float32
	MinExposureBias() float32
	MaxExposureDuration() objectivec.IObject
	SetMaxExposureDuration(value objectivec.IObject)
	MaxISO() float32
	SetMaxISO(value float32)
	MinExposureDuration() objectivec.IObject
	SetMinExposureDuration(value objectivec.IObject)
	MinISO() float32
	SetMinISO(value float32)


	

	// methods:
	ContainsExposureBias(exposureBias float32) bool


}





// Alloc allocates a new instance without initialization.
func (ec _ExposureBiasRangeClass) Alloc() ExposureBiasRange {
	rv := objc.Send[ExposureBiasRange](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _ExposureBiasRangeClass) New() ExposureBiasRange {
	rv := objc.Send[ExposureBiasRange](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExposureBiasRange) Init() ExposureBiasRange {
	rv := objc.Send[ExposureBiasRange](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExposureBiasRange) Autorelease() ExposureBiasRange {
	rv := objc.Send[ExposureBiasRange](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExposureBiasRange creates a new ExposureBiasRange instance.
func NewExposureBiasRange() ExposureBiasRange {
	return getExposureBiasRangeClass().New()
}





// An object that expresses an inclusive range of supported exposure bias values, in EV units.
//
// A defines its range using this type.


// An object that expresses an inclusive range of supported exposure bias values, in EV units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExposureBiasRange
type ExposureBiasRange struct {
	objectivec.Object
}

// ExposureBiasRangeFrom constructs a [ExposureBiasRange] from an unsafe.Pointer.
//
// An object that expresses an inclusive range of supported exposure bias values, in EV units.
func ExposureBiasRangeFrom(ptr unsafe.Pointer) ExposureBiasRange {
	return ExposureBiasRange{objectivec.Object{objc.ID(ptr)}}
}




















// Determines whether the range contains the specified exposure bias.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExposureBiasRange/containsExposureBias:
func (e_ ExposureBiasRange) ContainsExposureBias(exposureBias float32) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("containsExposureBias:"), exposureBias)
	return rv
}







// The maximum exposure bias in EV units that this range supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExposureBiasRange/maxExposureBias
func (e_ ExposureBiasRange) MaxExposureBias() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("maxExposureBias"))
	return rv
}


// The minimum exposure bias in EV units that this range supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExposureBiasRange/minExposureBias
func (e_ ExposureBiasRange) MinExposureBias() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("minExposureBias"))
	return rv
}


// A time value that indicates the maximum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxexposureduration
func (e_ ExposureBiasRange) MaxExposureDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("maxExposureDuration"))
	return rv
}


// A time value that indicates the maximum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxexposureduration
func (e_ ExposureBiasRange) SetMaxExposureDuration(value objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMaxExposureDuration:"), value)
}


// A floating-point number that indicates the maximum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxiso
func (e_ ExposureBiasRange) MaxISO() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("maxISO"))
	return rv
}


// A floating-point number that indicates the maximum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxiso
func (e_ ExposureBiasRange) SetMaxISO(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMaxISO:"), value)
}


// A time value that indicates the minimum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/minexposureduration
func (e_ ExposureBiasRange) MinExposureDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("minExposureDuration"))
	return rv
}


// A time value that indicates the minimum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/minexposureduration
func (e_ ExposureBiasRange) SetMinExposureDuration(value objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMinExposureDuration:"), value)
}


// A floating-point number that indicates the minimum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/miniso
func (e_ ExposureBiasRange) MinISO() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("minISO"))
	return rv
}


// A floating-point number that indicates the minimum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/miniso
func (e_ ExposureBiasRange) SetMinISO(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMinISO:"), value)
}








