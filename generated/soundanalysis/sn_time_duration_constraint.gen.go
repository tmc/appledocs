// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SNTimeDurationConstraint */


/* debug [class_header]: Header for SNTimeDurationConstraint */
// The class instance for the [SNTimeDurationConstraint] class.
var (
	SNTimeDurationConstraintClass     _SNTimeDurationConstraintClass
	SNTimeDurationConstraintClassOnce sync.Once
)

func getSNTimeDurationConstraintClass() _SNTimeDurationConstraintClass {
	SNTimeDurationConstraintClassOnce.Do(func() {
		SNTimeDurationConstraintClass = _SNTimeDurationConstraintClass{objc.GetClass("SNTimeDurationConstraint")}
	})
	return SNTimeDurationConstraintClass
}

type _SNTimeDurationConstraintClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SNTimeDurationConstraint */
// An interface definition for the [SNTimeDurationConstraint] class.
type ISNTimeDurationConstraint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SNTimeDurationConstraint */
	// properties:
	DurationRange() TimeRange /* not a class type */
	EnumeratedDurations() []foundation.Value
	Type() SNTimeDurationConstraintType
	KnownClassifications() objc.IObject /* cross-framework: NSString */
	SetKnownClassifications(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SNTimeDurationConstraint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SNTimeDurationConstraint */
// Alloc allocates a new instance without initialization.
func (sc _SNTimeDurationConstraintClass) Alloc() SNTimeDurationConstraint {
	rv := objc.Send[SNTimeDurationConstraint](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SNTimeDurationConstraintClass) New() SNTimeDurationConstraint {
	rv := objc.Send[SNTimeDurationConstraint](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SNTimeDurationConstraint) Init() SNTimeDurationConstraint {
	rv := objc.Send[SNTimeDurationConstraint](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SNTimeDurationConstraint) Autorelease() SNTimeDurationConstraint {
	rv := objc.Send[SNTimeDurationConstraint](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNTimeDurationConstraint creates a new SNTimeDurationConstraint instance.
func NewSNTimeDurationConstraint() SNTimeDurationConstraint {
	return getSNTimeDurationConstraintClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SNTimeDurationConstraint */
// Defines the time duration windows the request’s underlying sound classifier accepts with a range, or an array, of durations.
//
// Inspect the constraint’s property first to determine whether to check or next.


// Defines the time duration windows the request’s underlying sound classifier accepts with a range, or an array, of durations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class
type SNTimeDurationConstraint struct {
	objectivec.Object
}

// SNTimeDurationConstraintFrom constructs a [SNTimeDurationConstraint] from an unsafe.Pointer.
//
// Defines the time duration windows the request’s underlying sound classifier accepts with a range, or an array, of durations.
func SNTimeDurationConstraintFrom(ptr unsafe.Pointer) SNTimeDurationConstraint {
	return SNTimeDurationConstraint{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SNTimeDurationConstraint */

// Creates a constraint with a time duration range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class/initWithDurationRange:
func NewSNTimeDurationConstraintWithDurationRange(durationRange TimeRange /* not a class type */) SNTimeDurationConstraint {
	instance := getSNTimeDurationConstraintClass().Alloc()
	rv := objc.Send[SNTimeDurationConstraint](instance.ID, objc.Sel("initWithDurationRange:"), durationRange)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSNTimeDurationConstraintWithDurationRange */


// Creates a constraint with discrete time durations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class/initWithEnumeratedDurations:
func NewSNTimeDurationConstraintWithEnumeratedDurations(enumeratedDurations []foundation.Value) SNTimeDurationConstraint {
	instance := getSNTimeDurationConstraintClass().Alloc()
	rv := objc.Send[SNTimeDurationConstraint](instance.ID, objc.Sel("initWithEnumeratedDurations:"), enumeratedDurations)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSNTimeDurationConstraintWithEnumeratedDurations */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SNTimeDurationConstraint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SNTimeDurationConstraint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SNTimeDurationConstraint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SNTimeDurationConstraint */

// A time duration range the request’s underlying sound classifier accepts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class/durationRange
func (s_ SNTimeDurationConstraint) DurationRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](s_.ID, objc.Sel("durationRange"))
	return rv
}/* debug [instance_properties/getter]: durationRange */


// An array of time durations the request’s underlying sound classifier accepts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class/enumeratedDurations
func (s_ SNTimeDurationConstraint) EnumeratedDurations() []foundation.Value {
	rv := objc.Send[[]foundation.Value](s_.ID, objc.Sel("enumeratedDurations"))
	return rv
}/* debug [instance_properties/getter]: enumeratedDurations */


// An enumeration that tells you which constraint property to inspect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class/type
func (s_ SNTimeDurationConstraint) Type() SNTimeDurationConstraintType {
	rv := objc.Send[SNTimeDurationConstraintType](s_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// A string array that contains every prediction label in the request’s underlying sound classifier model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassifysoundrequest/knownclassifications
func (s_ SNTimeDurationConstraint) KnownClassifications() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("knownClassifications"))
	return rv
}/* debug [instance_properties/getter]: knownClassifications */


// A string array that contains every prediction label in the request’s underlying sound classifier model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassifysoundrequest/knownclassifications
func (s_ SNTimeDurationConstraint) SetKnownClassifications(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKnownClassifications:"), value)
}/* debug [instance_properties/setter]: knownClassifications */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SNTimeDurationConstraint */


