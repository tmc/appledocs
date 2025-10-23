// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [HKGlassesLensSpecification] class.
var (
	HKGlassesLensSpecificationClass     _HKGlassesLensSpecificationClass
	HKGlassesLensSpecificationClassOnce sync.Once
)

func getHKGlassesLensSpecificationClass() _HKGlassesLensSpecificationClass {
	HKGlassesLensSpecificationClassOnce.Do(func() {
		HKGlassesLensSpecificationClass = _HKGlassesLensSpecificationClass{objc.GetClass("HKGlassesLensSpecification")}
	})
	return HKGlassesLensSpecificationClass
}

type _HKGlassesLensSpecificationClass struct {
	class objc.Class
}

// An interface definition for the [HKGlassesLensSpecification] class.
type IHKGlassesLensSpecification interface {
	IHKLensSpecification
	// properties:
	FarPupillaryDistance() IHKQuantity
	SetFarPupillaryDistance(value IHKQuantity)
	NearPupillaryDistance() IHKQuantity
	SetNearPupillaryDistance(value IHKQuantity)
	Prism() IHKVisionPrism
	SetPrism(value IHKVisionPrism)
	VertexDistance() IHKQuantity
	SetVertexDistance(value IHKQuantity)
	// methods:
}

// An object that contains the glasses prescription data for one eye.
//
// To create a sample that stores a glasses prescription, start by defining a specification for each eye. Each lens specification object requires a parameter. This measures the lens’s strength for correcting either nearsightedness or farsightedness (measured in units). Next, create values for any of the prescription’s optional parameters. For example, if the prescription corrects for astigmatism, create the and values. The value uses units, while the uses . To add a multifocal correction for reading, create an value using units. To add a correction for eye alignment, create an object. To add information about the distance between the eye and the back of the lens, or the pupil and the center of the nose, create , , and values. All of these use millimeters. Then you can create the lens specification. After you create your lens specifications, you can create an sample. Then save the sample to the HealthKit store. Finally, add an image or PDF of the prescription to the sample as an attachment.


// An object that contains the glasses prescription data for one eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGlassesLensSpecification
type HKGlassesLensSpecification struct {
	HKLensSpecification
}

// HKGlassesLensSpecificationFrom constructs a [HKGlassesLensSpecification] from an unsafe.Pointer.
//
// An object that contains the glasses prescription data for one eye.
func HKGlassesLensSpecificationFrom(ptr unsafe.Pointer) HKGlassesLensSpecification {
	return HKGlassesLensSpecification{
		HKLensSpecification: HKLensSpecificationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (hc _HKGlassesLensSpecificationClass) Alloc() HKGlassesLensSpecification {
	rv := objc.Send[HKGlassesLensSpecification](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKGlassesLensSpecificationClass) New() HKGlassesLensSpecification {
	rv := objc.Send[HKGlassesLensSpecification](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKGlassesLensSpecification) Init() HKGlassesLensSpecification {
	rv := objc.Send[HKGlassesLensSpecification](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKGlassesLensSpecification) Autorelease() HKGlassesLensSpecification {
	rv := objc.Send[HKGlassesLensSpecification](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKGlassesLensSpecification creates a new HKGlassesLensSpecification instance.
func NewHKGlassesLensSpecification() HKGlassesLensSpecification {
	return getHKGlassesLensSpecificationClass().New()
}



// The distance between the pupil and the center of the nose when looking at an object far away, measured in mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkglasseslensspecification/farpupillarydistance
func (h_ HKGlassesLensSpecification) FarPupillaryDistance() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("farPupillaryDistance"))
	return rv
}


// The distance between the pupil and the center of the nose when looking at an object far away, measured in mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkglasseslensspecification/farpupillarydistance
func (h_ HKGlassesLensSpecification) SetFarPupillaryDistance(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setFarPupillaryDistance:"), value)
}


// The distance between the pupil and the center of the nose when looking at a nearby object, measured in mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkglasseslensspecification/nearpupillarydistance
func (h_ HKGlassesLensSpecification) NearPupillaryDistance() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("nearPupillaryDistance"))
	return rv
}


// The distance between the pupil and the center of the nose when looking at a nearby object, measured in mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkglasseslensspecification/nearpupillarydistance
func (h_ HKGlassesLensSpecification) SetNearPupillaryDistance(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setNearPupillaryDistance:"), value)
}


// An object that contains information about the eye alignment correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkglasseslensspecification/prism
func (h_ HKGlassesLensSpecification) Prism() IHKVisionPrism {
	rv := objc.Send[HKVisionPrism](h_.ID, objc.Sel("prism"))
	return rv
}


// An object that contains information about the eye alignment correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkglasseslensspecification/prism
func (h_ HKGlassesLensSpecification) SetPrism(value IHKVisionPrism) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPrism:"), value)
}


// The distance between the back of the lens and the eye, measured in mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkglasseslensspecification/vertexdistance
func (h_ HKGlassesLensSpecification) VertexDistance() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("vertexDistance"))
	return rv
}


// The distance between the back of the lens and the eye, measured in mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkglasseslensspecification/vertexdistance
func (h_ HKGlassesLensSpecification) SetVertexDistance(value IHKQuantity) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setVertexDistance:"), value)
}



