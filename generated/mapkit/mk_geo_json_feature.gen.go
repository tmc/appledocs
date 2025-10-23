// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKGeoJSONFeature] class.
var (
	MKGeoJSONFeatureClass     _MKGeoJSONFeatureClass
	MKGeoJSONFeatureClassOnce sync.Once
)

func getMKGeoJSONFeatureClass() _MKGeoJSONFeatureClass {
	MKGeoJSONFeatureClassOnce.Do(func() {
		MKGeoJSONFeatureClass = _MKGeoJSONFeatureClass{objc.GetClass("MKGeoJSONFeature")}
	})
	return MKGeoJSONFeatureClass
}

type _MKGeoJSONFeatureClass struct {
	class objc.Class
}

// An interface definition for the [MKGeoJSONFeature] class.
type IMKGeoJSONFeature interface {
	objectivec.IObject
	Identifier() string
	Geometry() unsafe.Pointer
	SetGeometry(value unsafe.Pointer)
	Properties() foundation.Data
	SetProperties(value foundation.IData)
}

// The decoded representation of a GeoJSON feature.
//
// A feature is an object with associated geometry and optional properties in JSON that you define. MapKit exposes these optional properties, but treats them as opaque. is one of the classes that the GeoJSON decoder ( ) can return. See the GeoJSON standards specification for more information about objects.


// The decoded representation of a GeoJSON feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeoJSONFeature
type MKGeoJSONFeature struct {
	objectivec.Object
}

// MKGeoJSONFeatureFrom constructs a [MKGeoJSONFeature] from an unsafe.Pointer.
//
// The decoded representation of a GeoJSON feature.
func MKGeoJSONFeatureFrom(ptr unsafe.Pointer) MKGeoJSONFeature {
	return MKGeoJSONFeature{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKGeoJSONFeatureClass) Alloc() MKGeoJSONFeature {
	rv := objc.Send[MKGeoJSONFeature](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKGeoJSONFeatureClass) New() MKGeoJSONFeature {
	rv := objc.Send[MKGeoJSONFeature](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKGeoJSONFeature) Init() MKGeoJSONFeature {
	rv := objc.Send[MKGeoJSONFeature](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKGeoJSONFeature) Autorelease() MKGeoJSONFeature {
	rv := objc.Send[MKGeoJSONFeature](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKGeoJSONFeature creates a new MKGeoJSONFeature instance.
func NewMKGeoJSONFeature() MKGeoJSONFeature {
	return getMKGeoJSONFeatureClass().New()
}



// An optional identifier the class returns as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeoJSONFeature/identifier
func (m_ MKGeoJSONFeature) Identifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("identifier"))
	return rv
}


// The shape or shapes associated with the GeoJSON feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeojsonfeature/geometry
func (m_ MKGeoJSONFeature) Geometry() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("geometry"))
	return rv
}


// The shape or shapes associated with the GeoJSON feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeojsonfeature/geometry
func (m_ MKGeoJSONFeature) SetGeometry(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGeometry:"), value)
}


// Optional serialized JSON data that corresponds to the properties key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeojsonfeature/properties
func (m_ MKGeoJSONFeature) Properties() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("properties"))
	return rv
}


// Optional serialized JSON data that corresponds to the properties key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkgeojsonfeature/properties
func (m_ MKGeoJSONFeature) SetProperties(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProperties:"), value)
}



