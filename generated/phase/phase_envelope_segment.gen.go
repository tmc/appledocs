// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHASEEnvelopeSegment] class.
var (
	PHASEEnvelopeSegmentClass     _PHASEEnvelopeSegmentClass
	PHASEEnvelopeSegmentClassOnce sync.Once
)

func getPHASEEnvelopeSegmentClass() _PHASEEnvelopeSegmentClass {
	PHASEEnvelopeSegmentClassOnce.Do(func() {
		PHASEEnvelopeSegmentClass = _PHASEEnvelopeSegmentClass{objc.GetClass("PHASEEnvelopeSegment")}
	})
	return PHASEEnvelopeSegmentClass
}

type _PHASEEnvelopeSegmentClass struct {
	class objc.Class
}

// An interface definition for the [PHASEEnvelopeSegment] class.
type IPHASEEnvelopeSegment interface {
	objectivec.IObject
}

// A curved portion of an envelope.
//
// This class specifies a curve that determines the _y-_value rate of change over a particular portion of an envelope’s graph. For example, the difference between a segment and an segment is that they share opposite rates of change; where the cubed curve’s value changes fastest in the segment’s domain, the inverse-cubed curve changes slowest, and vice versa.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeSegment
type PHASEEnvelopeSegment struct {
	objectivec.Object
}

// PHASEEnvelopeSegmentFrom constructs a [PHASEEnvelopeSegment] from an unsafe.Pointer.
//
// A curved portion of an envelope.
func PHASEEnvelopeSegmentFrom(ptr unsafe.Pointer) PHASEEnvelopeSegment {
	return PHASEEnvelopeSegment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEEnvelopeSegmentClass) Alloc() PHASEEnvelopeSegment {
	rv := objc.Send[PHASEEnvelopeSegment](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEEnvelopeSegmentClass) New() PHASEEnvelopeSegment {
	rv := objc.Send[PHASEEnvelopeSegment](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEEnvelopeSegment) Init() PHASEEnvelopeSegment {
	rv := objc.Send[PHASEEnvelopeSegment](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEEnvelopeSegment) Autorelease() PHASEEnvelopeSegment {
	rv := objc.Send[PHASEEnvelopeSegment](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEEnvelopeSegment creates a new PHASEEnvelopeSegment instance.
func NewPHASEEnvelopeSegment() PHASEEnvelopeSegment {
	return getPHASEEnvelopeSegmentClass().New()
}




// Creates a curved portion of an envelope.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeSegment/init(endPoint:curveType:)
func NewPHASEEnvelopeSegmentWithEndPointCurveType(endPoint unsafe.Pointer, curveType unsafe.Pointer) PHASEEnvelopeSegment {
	instance := getPHASEEnvelopeSegmentClass().Alloc()
	rv := objc.Send[PHASEEnvelopeSegment](instance.ID, objc.Sel("initWithEndPoint:curveType:"), endPoint, curveType)
	rv.Autorelease()
	return rv
}


// A curve along the envelope that shapes the segment.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeSegment/curveType
func (p_ PHASEEnvelopeSegment) CurveType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("curveType"))
	return rv
}


// SetCurveType sets the value of the curveType property.
// A curve along the envelope that shapes the segment.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeSegment/curveType
func (p_ PHASEEnvelopeSegment) SetCurveType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurveType:"), value)
}

// A point that identifies the end of the segment along the envelope.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeSegment/endPoint
func (p_ PHASEEnvelopeSegment) EndPoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("endPoint"))
	return rv
}


// SetEndPoint sets the value of the endPoint property.
// A point that identifies the end of the segment along the envelope.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeSegment/endPoint
func (p_ PHASEEnvelopeSegment) SetEndPoint(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEndPoint:"), value)
}


