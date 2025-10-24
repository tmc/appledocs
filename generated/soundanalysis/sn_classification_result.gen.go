// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SNClassificationResult */


/* debug [class_header]: Header for SNClassificationResult */
// The class instance for the [SNClassificationResult] class.
var (
	SNClassificationResultClass     _SNClassificationResultClass
	SNClassificationResultClassOnce sync.Once
)

func getSNClassificationResultClass() _SNClassificationResultClass {
	SNClassificationResultClassOnce.Do(func() {
		SNClassificationResultClass = _SNClassificationResultClass{objc.GetClass("SNClassificationResult")}
	})
	return SNClassificationResultClass
}

type _SNClassificationResultClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SNClassificationResult */
// An interface definition for the [SNClassificationResult] class.
type ISNClassificationResult interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SNClassificationResult */
	// properties:
	Classifications() []SNClassification
	TimeRange() TimeRange /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SNClassificationResult */
	// methods:
	ClassificationForIdentifier(identifier objc.IObject /* cross-framework: NSString */) ISNClassification
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SNClassificationResult */
// Alloc allocates a new instance without initialization.
func (sc _SNClassificationResultClass) Alloc() SNClassificationResult {
	rv := objc.Send[SNClassificationResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SNClassificationResultClass) New() SNClassificationResult {
	rv := objc.Send[SNClassificationResult](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SNClassificationResult) Init() SNClassificationResult {
	rv := objc.Send[SNClassificationResult](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SNClassificationResult) Autorelease() SNClassificationResult {
	rv := objc.Send[SNClassificationResult](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNClassificationResult creates a new SNClassificationResult instance.
func NewSNClassificationResult() SNClassificationResult {
	return getSNClassificationResultClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SNClassificationResult */
// A result that contains the highest-ranking classifications in a time range.
//
// An represents the predictions that a sound classification model made for a time span in an audio file or stream. Each result contains one or more classification predictions and a time range within the audio data. An audio analyzer, such as and , produces an each time it recognizes a sound for any of its instances.


// A result that contains the highest-ranking classifications in a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassificationResult
type SNClassificationResult struct {
	objectivec.Object
}

// SNClassificationResultFrom constructs a [SNClassificationResult] from an unsafe.Pointer.
//
// A result that contains the highest-ranking classifications in a time range.
func SNClassificationResultFrom(ptr unsafe.Pointer) SNClassificationResult {
	return SNClassificationResult{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SNClassificationResult *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SNClassificationResult */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SNClassificationResult */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SNClassificationResult */

// Returns the classification for an identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassificationResult/classification(forIdentifier:)
func (s_ SNClassificationResult) ClassificationForIdentifier(identifier objc.IObject /* cross-framework: NSString */) ISNClassification {
	rv := objc.Send[SNClassification](s_.ID, objc.Sel("classificationForIdentifier:"), identifier)
	return rv
}/* debug [instance_methods/method]: ClassificationForIdentifier */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SNClassificationResult */

// A sorted array of the request’s top classification candidates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassificationResult/classifications
func (s_ SNClassificationResult) Classifications() []SNClassification {
	rv := objc.Send[[]SNClassification](s_.ID, objc.Sel("classifications"))
	return rv
}/* debug [instance_properties/getter]: classifications */


// The time span that corresponds to the result’s classifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassificationResult/timeRange
func (s_ SNClassificationResult) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](s_.ID, objc.Sel("timeRange"))
	return rv
}/* debug [instance_properties/getter]: timeRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SNClassificationResult */



