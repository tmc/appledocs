// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEEnvelopeSegment */


/* debug [class_header]: Header for PHASEEnvelopeSegment */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEEnvelopeSegment */
// An interface definition for the [PHASEEnvelopeSegment] class.
type IPHASEEnvelopeSegment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEEnvelopeSegment */
	// properties:
	CurveType() PHASECurveType
	SetCurveType(value PHASECurveType)
	EndPoint() unsafe.Pointer
	SetEndPoint(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEEnvelopeSegment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEEnvelopeSegment */
// Alloc allocates a new instance without initialization.
func (pc _PHASEEnvelopeSegmentClass) Alloc() PHASEEnvelopeSegment {
	rv := objc.Send[PHASEEnvelopeSegment](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEEnvelopeSegment */
// A curved portion of an envelope.
//
// This class specifies a curve that determines the _y-_value rate of change over a particular portion of an envelope’s graph. For example, the difference between a segment and an segment is that they share opposite rates of change; where the cubed curve’s value changes fastest in the segment’s domain, the inverse-cubed curve changes slowest, and vice versa.


// A curved portion of an envelope.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEEnvelopeSegment */

// Creates a curved portion of an envelope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeSegment/init(endPoint:curveType:)
func NewPHASEEnvelopeSegmentWithEndPointCurveType(endPoint unsafe.Pointer, curveType PHASECurveType) PHASEEnvelopeSegment {
	instance := getPHASEEnvelopeSegmentClass().Alloc()
	rv := objc.Send[PHASEEnvelopeSegment](instance.ID, objc.Sel("initWithEndPoint:curveType:"), endPoint, curveType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEEnvelopeSegmentWithEndPointCurveType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEEnvelopeSegment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEEnvelopeSegment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEEnvelopeSegment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEEnvelopeSegment */

// A curve along the envelope that shapes the segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeSegment/curveType
func (p_ PHASEEnvelopeSegment) CurveType() PHASECurveType {
	rv := objc.Send[PHASECurveType](p_.ID, objc.Sel("curveType"))
	return rv
}/* debug [instance_properties/getter]: curveType */


// A curve along the envelope that shapes the segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeSegment/curveType
func (p_ PHASEEnvelopeSegment) SetCurveType(value PHASECurveType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurveType:"), value)
}/* debug [instance_properties/setter]: curveType */


// A point that identifies the end of the segment along the envelope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeSegment/endPoint
func (p_ PHASEEnvelopeSegment) EndPoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("endPoint"))
	return rv
}/* debug [instance_properties/getter]: endPoint */


// A point that identifies the end of the segment along the envelope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEEnvelopeSegment/endPoint
func (p_ PHASEEnvelopeSegment) SetEndPoint(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEndPoint:"), value)
}/* debug [instance_properties/setter]: endPoint */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEEnvelopeSegment */


