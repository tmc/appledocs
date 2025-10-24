// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SNClassification */

/* debug [class_header]: Header for SNClassification */
// The class instance for the [SNClassification] class.
var (
	SNClassificationClass     _SNClassificationClass
	SNClassificationClassOnce sync.Once
)

func getSNClassificationClass() _SNClassificationClass {
	SNClassificationClassOnce.Do(func() {
		SNClassificationClass = _SNClassificationClass{objc.GetClass("SNClassification")}
	})
	return SNClassificationClass
}

type _SNClassificationClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for SNClassification */
// An interface definition for the [SNClassification] class.
type ISNClassification interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for SNClassification */
	// properties:
	Confidence() float64
	Identifier() objc.IObject /* cross-framework: NSString */
	Classifications() ISNClassification
	SetClassifications(value ISNClassification)
	TimeRange() TimeRange /* not a class type */
	SetTimeRange(value TimeRange /* not a class type */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for SNClassification */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for SNClassification */
// Alloc allocates a new instance without initialization.
func (sc _SNClassificationClass) Alloc() SNClassification {
	rv := objc.Send[SNClassification](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SNClassificationClass) New() SNClassification {
	rv := objc.Send[SNClassification](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SNClassification) Init() SNClassification {
	rv := objc.Send[SNClassification](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SNClassification) Autorelease() SNClassification {
	rv := objc.Send[SNClassification](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNClassification creates a new SNClassification instance.
func NewSNClassification() SNClassification {
	return getSNClassificationClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for SNClassification */
// A type that pairs a sound classifier’s prediction with its confidence in that prediction.
//
// An represents a single sound classification prediction, and the sound classifier model’s confidence in that prediction.

// A type that pairs a sound classifier’s prediction with its confidence in that prediction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassification
type SNClassification struct {
	objectivec.Object
}

// SNClassificationFrom constructs a [SNClassification] from an unsafe.Pointer.
//
// A type that pairs a sound classifier’s prediction with its confidence in that prediction.
func SNClassificationFrom(ptr unsafe.Pointer) SNClassification {
	return SNClassification{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for SNClassification */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for SNClassification */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for SNClassification */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for SNClassification */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for SNClassification */

// The confidence value the model has in its prediction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassification/confidence
func (s_ SNClassification) Confidence() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("confidence"))
	return rv
} /* debug [instance_properties/getter]: confidence */

// A prediction label that’s one of the classifications a sound classifier’s underlying model defines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassification/identifier
func (s_ SNClassification) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("identifier"))
	return rv
} /* debug [instance_properties/getter]: identifier */

// A sorted array of the request’s top classification candidates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassificationresult/classifications
func (s_ SNClassification) Classifications() ISNClassification {
	rv := objc.Send[SNClassification](s_.ID, objc.Sel("classifications"))
	return rv
} /* debug [instance_properties/getter]: classifications */

// A sorted array of the request’s top classification candidates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassificationresult/classifications
func (s_ SNClassification) SetClassifications(value ISNClassification) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setClassifications:"), value)
} /* debug [instance_properties/setter]: classifications */

// The time span that corresponds to the result’s classifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassificationresult/timerange
func (s_ SNClassification) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](s_.ID, objc.Sel("timeRange"))
	return rv
} /* debug [instance_properties/getter]: timeRange */

// The time span that corresponds to the result’s classifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/soundanalysis/snclassificationresult/timerange
func (s_ SNClassification) SetTimeRange(value TimeRange /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTimeRange:"), value)
} /* debug [instance_properties/setter]: timeRange */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SNClassification */
