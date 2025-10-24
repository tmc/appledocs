// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEEnvelope */


/* debug [class_header]: Header for PHASEEnvelope */
// The class instance for the [PHASEEnvelope] class.
var (
	PHASEEnvelopeClass     _PHASEEnvelopeClass
	PHASEEnvelopeClassOnce sync.Once
)

func getPHASEEnvelopeClass() _PHASEEnvelopeClass {
	PHASEEnvelopeClassOnce.Do(func() {
		PHASEEnvelopeClass = _PHASEEnvelopeClass{objc.GetClass("PHASEEnvelope")}
	})
	return PHASEEnvelopeClass
}

type _PHASEEnvelopeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEEnvelope */
// An interface definition for the [PHASEEnvelope] class.
type IPHASEEnvelope interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEEnvelope */
	// properties:
	Domain() IPHASENumericPair
	Range() IPHASENumericPair
	Segments() []PHASEEnvelopeSegment
	StartPoint() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEEnvelope */
	// methods:
	EvaluateForValue(x float64) float64
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEEnvelope */
// Alloc allocates a new instance without initialization.
func (pc _PHASEEnvelopeClass) Alloc() PHASEEnvelope {
	rv := objc.Send[PHASEEnvelope](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEEnvelopeClass) New() PHASEEnvelope {
	rv := objc.Send[PHASEEnvelope](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEEnvelope) Init() PHASEEnvelope {
	rv := objc.Send[PHASEEnvelope](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEEnvelope) Autorelease() PHASEEnvelope {
	rv := objc.Send[PHASEEnvelope](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEEnvelope creates a new PHASEEnvelope instance.
func NewPHASEEnvelope() PHASEEnvelope {
	return getPHASEEnvelopeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEEnvelope */
// A collection of segments that connect to graph a complex curve over a linear input.
//
// In traditional audio uses, an defines a complex graph that determines the volume of audio data over an input duration. PHASE uses envelopes in a similar way. Given a value on the envelope’s input axis, the function plots and returns the result on the output axis. The following are possible uses of this class: Sound event nodes, such as , can shape their volume using an envelope; see . Distance models shape sounds with a 3D position using an envelope; see . An envelope can do more than shape audio. To gradually change an envelope’s input value over time, use the class, which creates a function with a metaparameter value as input. An app can use the numeric result for any purpose. For example, the x-axis can be distance and the y-axis can be playback rate. At runtime, PHASE determines whether a particular member of the array slopes up or down along the domain depending on the envelope’s particular use case.


// A collection of segments that connect to graph a complex curve over a linear input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelope
type PHASEEnvelope struct {
	objectivec.Object
}

// PHASEEnvelopeFrom constructs a [PHASEEnvelope] from an unsafe.Pointer.
//
// A collection of segments that connect to graph a complex curve over a linear input.
func PHASEEnvelopeFrom(ptr unsafe.Pointer) PHASEEnvelope {
	return PHASEEnvelope{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEEnvelope */

// Creates an envelope with a start point and segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelope/init(startPoint:segments:)
func NewPHASEEnvelopeWithStartPointSegments(startPoint unsafe.Pointer, segments []PHASEEnvelopeSegment) PHASEEnvelope {
	instance := getPHASEEnvelopeClass().Alloc()
	rv := objc.Send[PHASEEnvelope](instance.ID, objc.Sel("initWithStartPoint:segments:"), startPoint, segments)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEEnvelopeWithStartPointSegments */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEEnvelope */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEEnvelope */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEEnvelope */

// Provides the height of the envelope for an input value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelope/evaluate(x:)
func (p_ PHASEEnvelope) EvaluateForValue(x float64) float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("evaluateForValue:"), x)
	return rv
}/* debug [instance_methods/method]: EvaluateForValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEEnvelope */

// The range of the envelope’s possible input values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelope/domain
func (p_ PHASEEnvelope) Domain() IPHASENumericPair {
	rv := objc.Send[PHASENumericPair](p_.ID, objc.Sel("domain"))
	return rv
}/* debug [instance_properties/getter]: domain */


// The bounds of the output value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelope/range
func (p_ PHASEEnvelope) Range() IPHASENumericPair {
	rv := objc.Send[PHASENumericPair](p_.ID, objc.Sel("range"))
	return rv
}/* debug [instance_properties/getter]: range */


// An array of the envelope’s segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelope/segments
func (p_ PHASEEnvelope) Segments() []PHASEEnvelopeSegment {
	rv := objc.Send[[]PHASEEnvelopeSegment](p_.ID, objc.Sel("segments"))
	return rv
}/* debug [instance_properties/getter]: segments */


// The starting point along the envelope’s duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelope/startPoint
func (p_ PHASEEnvelope) StartPoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("startPoint"))
	return rv
}/* debug [instance_properties/getter]: startPoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEEnvelope */


