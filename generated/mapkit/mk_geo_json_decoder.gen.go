// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKGeoJSONDecoder] class.
var (
	MKGeoJSONDecoderClass     _MKGeoJSONDecoderClass
	MKGeoJSONDecoderClassOnce sync.Once
)

func getMKGeoJSONDecoderClass() _MKGeoJSONDecoderClass {
	MKGeoJSONDecoderClassOnce.Do(func() {
		MKGeoJSONDecoderClass = _MKGeoJSONDecoderClass{objc.GetClass("MKGeoJSONDecoder")}
	})
	return MKGeoJSONDecoderClass
}

type _MKGeoJSONDecoderClass struct {
	class objc.Class
}

// An interface definition for the [MKGeoJSONDecoder] class.
type IMKGeoJSONDecoder interface {
	objectivec.IObject
	GeoJSONObjectsWithDataError(data foundation.IData, errorPtr unsafe.Pointer) []objc.ID
}

// An object that decodes GeoJSON objects into MapKit types.
//
// The GeoJSON decoder returns objects that conform to the protocol.


// An object that decodes GeoJSON objects into MapKit types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeoJSONDecoder
type MKGeoJSONDecoder struct {
	objectivec.Object
}

// MKGeoJSONDecoderFrom constructs a [MKGeoJSONDecoder] from an unsafe.Pointer.
//
// An object that decodes GeoJSON objects into MapKit types.
func MKGeoJSONDecoderFrom(ptr unsafe.Pointer) MKGeoJSONDecoder {
	return MKGeoJSONDecoder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKGeoJSONDecoderClass) Alloc() MKGeoJSONDecoder {
	rv := objc.Send[MKGeoJSONDecoder](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKGeoJSONDecoderClass) New() MKGeoJSONDecoder {
	rv := objc.Send[MKGeoJSONDecoder](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKGeoJSONDecoder) Init() MKGeoJSONDecoder {
	rv := objc.Send[MKGeoJSONDecoder](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKGeoJSONDecoder) Autorelease() MKGeoJSONDecoder {
	rv := objc.Send[MKGeoJSONDecoder](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKGeoJSONDecoder creates a new MKGeoJSONDecoder instance.
func NewMKGeoJSONDecoder() MKGeoJSONDecoder {
	return getMKGeoJSONDecoderClass().New()
}



// Decodes the provided data into native MapKit types that a map can display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKGeoJSONDecoder/decode(_:)
func (m_ MKGeoJSONDecoder) GeoJSONObjectsWithDataError(data foundation.IData, errorPtr unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("geoJSONObjectsWithData:error:"), data, errorPtr)
	return rv
}



