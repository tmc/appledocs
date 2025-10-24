// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [SNTimeDurationConstraint] class.
type ISNTimeDurationConstraint interface {
	objectivec.IObject
	// properties:
	KnownClassifications() objc.IObject /* cross-framework: NSString */
	SetKnownClassifications(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (sc _SNTimeDurationConstraintClass) Alloc() SNTimeDurationConstraint {
	rv := objc.Send[SNTimeDurationConstraint](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A string array that contains every prediction label in the request’s underlying sound classifier model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassifysoundrequest/knownclassifications
func (s_ SNTimeDurationConstraint) KnownClassifications() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("knownClassifications"))
	return rv
}


// A string array that contains every prediction label in the request’s underlying sound classifier model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassifysoundrequest/knownclassifications
func (s_ SNTimeDurationConstraint) SetKnownClassifications(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKnownClassifications:"), value)
}




