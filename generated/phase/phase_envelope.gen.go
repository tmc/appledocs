// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PHASEEnvelope] class.
type IPHASEEnvelope interface {
	objectivec.IObject
	EvaluateForValue(x unsafe.Pointer) unsafe.Pointer
}

// A collection of segments that connect to graph a complex curve over a linear input.
//
// In traditional audio uses, an defines a complex graph that determines the volume of audio data over an input duration. PHASE uses envelopes in a similar way. Given a value on the envelope’s input axis, the function plots and returns the result on the output axis. The following are possible uses of this class: Sound event nodes, such as , can shape their volume using an envelope; see . Distance models shape sounds with a 3D position using an envelope; see . An envelope can do more than shape audio. To gradually change an envelope’s input value over time, use the class, which creates a function with a metaparameter value as input. An app can use the numeric result for any purpose. For example, the x-axis can be distance and the y-axis can be playback rate. At runtime, PHASE determines whether a particular member of the array slopes up or down along the domain depending on the envelope’s particular use case.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PHASEEnvelopeClass) Alloc() PHASEEnvelope {
	rv := objc.Send[PHASEEnvelope](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates an envelope with a start point and segments.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelope/init(startPoint:segments:)
func NewPHASEEnvelopeWithStartPointSegments(startPoint unsafe.Pointer, segments []PHASEEnvelopeSegment) PHASEEnvelope {
	instance := getPHASEEnvelopeClass().Alloc()
	rv := objc.Send[PHASEEnvelope](instance.ID, objc.Sel("initWithStartPoint:segments:"), startPoint, segments)
	rv.Autorelease()
	return rv
}


// Provides the height of the envelope for an input value.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelope/evaluate(x:)
func (p_ PHASEEnvelope) EvaluateForValue(x unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("evaluateForValue:"), x)
	return rv
}

// The range of the envelope’s possible input values.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelope/domain
func (p_ PHASEEnvelope) Domain() PHASENumericPair {
	rv := objc.Send[PHASENumericPair](p_.ID, objc.Sel("domain"))
	return rv
}

// The bounds of the output value.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelope/range
func (p_ PHASEEnvelope) Range() PHASENumericPair {
	rv := objc.Send[PHASENumericPair](p_.ID, objc.Sel("range"))
	return rv
}

// An array of the envelope’s segments.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelope/segments
func (p_ PHASEEnvelope) Segments() []PHASEEnvelopeSegment {
	rv := objc.Send[[]PHASEEnvelopeSegment](p_.ID, objc.Sel("segments"))
	return rv
}

// The starting point along the envelope’s duration.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelope/startPoint
func (p_ PHASEEnvelope) StartPoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("startPoint"))
	return rv
}


