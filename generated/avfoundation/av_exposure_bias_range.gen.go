// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVExposureBiasRange */


/* debug [class_header]: Header for AVExposureBiasRange */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ExposureBiasRange */
// An interface definition for the [ExposureBiasRange] class.
type IExposureBiasRange interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ExposureBiasRange */
	// properties:
	MaxExposureBias() float32
	MinExposureBias() float32
	MaxExposureDuration() objc.IObject /* cross-framework: Time */
	SetMaxExposureDuration(value objc.IObject /* cross-framework: Time */)
	MaxISO() float32
	SetMaxISO(value float32)
	MinExposureDuration() objc.IObject /* cross-framework: Time */
	SetMinExposureDuration(value objc.IObject /* cross-framework: Time */)
	MinISO() float32
	SetMinISO(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ExposureBiasRange */
	// methods:
	ContainsExposureBias(exposureBias float32) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ExposureBiasRange */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ExposureBiasRange */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ExposureBiasRange *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ExposureBiasRange */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ExposureBiasRange */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ExposureBiasRange */

// Determines whether the range contains the specified exposure bias.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExposureBiasRange/containsExposureBias:
func (e_ ExposureBiasRange) ContainsExposureBias(exposureBias float32) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("containsExposureBias:"), exposureBias)
	return rv
}/* debug [instance_methods/method]: ContainsExposureBias */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ExposureBiasRange */

// The maximum exposure bias in EV units that this range supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExposureBiasRange/maxExposureBias
func (e_ ExposureBiasRange) MaxExposureBias() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("maxExposureBias"))
	return rv
}/* debug [instance_properties/getter]: maxExposureBias */


// The minimum exposure bias in EV units that this range supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExposureBiasRange/minExposureBias
func (e_ ExposureBiasRange) MinExposureBias() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("minExposureBias"))
	return rv
}/* debug [instance_properties/getter]: minExposureBias */


// A time value that indicates the maximum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxexposureduration
func (e_ ExposureBiasRange) MaxExposureDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](e_.ID, objc.Sel("maxExposureDuration"))
	return rv
}/* debug [instance_properties/getter]: maxExposureDuration */


// A time value that indicates the maximum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxexposureduration
func (e_ ExposureBiasRange) SetMaxExposureDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMaxExposureDuration:"), value)
}/* debug [instance_properties/setter]: maxExposureDuration */


// A floating-point number that indicates the maximum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxiso
func (e_ ExposureBiasRange) MaxISO() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("maxISO"))
	return rv
}/* debug [instance_properties/getter]: maxISO */


// A floating-point number that indicates the maximum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxiso
func (e_ ExposureBiasRange) SetMaxISO(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMaxISO:"), value)
}/* debug [instance_properties/setter]: maxISO */


// A time value that indicates the minimum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/minexposureduration
func (e_ ExposureBiasRange) MinExposureDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](e_.ID, objc.Sel("minExposureDuration"))
	return rv
}/* debug [instance_properties/getter]: minExposureDuration */


// A time value that indicates the minimum supported exposure duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/minexposureduration
func (e_ ExposureBiasRange) SetMinExposureDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMinExposureDuration:"), value)
}/* debug [instance_properties/setter]: minExposureDuration */


// A floating-point number that indicates the minimum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/miniso
func (e_ ExposureBiasRange) MinISO() float32 {
	rv := objc.Send[float32](e_.ID, objc.Sel("minISO"))
	return rv
}/* debug [instance_properties/getter]: minISO */


// A floating-point number that indicates the minimum supported exposure ISO value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/miniso
func (e_ ExposureBiasRange) SetMinISO(value float32) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setMinISO:"), value)
}/* debug [instance_properties/setter]: minISO */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVExposureBiasRange */



