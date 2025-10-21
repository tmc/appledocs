// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKMapItemIdentifier] class.
var (
	MKMapItemIdentifierClass     _MKMapItemIdentifierClass
	MKMapItemIdentifierClassOnce sync.Once
)

func getMKMapItemIdentifierClass() _MKMapItemIdentifierClass {
	MKMapItemIdentifierClassOnce.Do(func() {
		MKMapItemIdentifierClass = _MKMapItemIdentifierClass{objc.GetClass("MKMapItemIdentifier")}
	})
	return MKMapItemIdentifierClass
}

type _MKMapItemIdentifierClass struct {
	class objc.Class
}

// An interface definition for the [MKMapItemIdentifier] class.
type IMKMapItemIdentifier interface {
	objectivec.IObject
}

// A unique identifier for a place.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapItem/Identifier-swift.class
type MKMapItemIdentifier struct {
	objectivec.Object
}

// MKMapItemIdentifierFrom constructs a [MKMapItemIdentifier] from an unsafe.Pointer.
//
// A unique identifier for a place.
func MKMapItemIdentifierFrom(ptr unsafe.Pointer) MKMapItemIdentifier {
	return MKMapItemIdentifier{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMapItemIdentifierClass) Alloc() MKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMapItemIdentifierClass) New() MKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapItemIdentifier) Init() MKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapItemIdentifier) Autorelease() MKMapItemIdentifier {
	rv := objc.Send[MKMapItemIdentifier](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapItemIdentifier creates a new MKMapItemIdentifier instance.
func NewMKMapItemIdentifier() MKMapItemIdentifier {
	return getMKMapItemIdentifierClass().New()
}


// The point-of-interest category for the map item.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/pointofinterestcategory
func (m_ MKMapItemIdentifier) PointOfInterestCategory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pointOfInterestCategory"))
	return rv
}


// SetPointOfInterestCategory sets the value of the pointOfInterestCategory property.
// The point-of-interest category for the map item.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/pointofinterestcategory
func (m_ MKMapItemIdentifier) SetPointOfInterestCategory(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestCategory:"), value)
}

// The descriptive name associated with the map item.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/name
func (m_ MKMapItemIdentifier) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The descriptive name associated with the map item.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/name
func (m_ MKMapItemIdentifier) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}

// The time zone of the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/timezone
func (m_ MKMapItemIdentifier) TimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timeZone"))
	return rv
}


// SetTimeZone sets the value of the timeZone property.
// The time zone of the specified location.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/timezone
func (m_ MKMapItemIdentifier) SetTimeZone(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeZone:"), value)
}

// The phone number associated with a business at the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/phonenumber
func (m_ MKMapItemIdentifier) PhoneNumber() string {
	rv := objc.Send[string](m_.ID, objc.Sel("phoneNumber"))
	return rv
}


// SetPhoneNumber sets the value of the phoneNumber property.
// The phone number associated with a business at the specified location.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/phonenumber
func (m_ MKMapItemIdentifier) SetPhoneNumber(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPhoneNumber:"), objc.String(value))
}

// A set of alternative identifiers for a place.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/alternateidentifiers
func (m_ MKMapItemIdentifier) AlternateIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("alternateIdentifiers"))
	return rv
}


// SetAlternateIdentifiers sets the value of the alternateIdentifiers property.
// A set of alternative identifiers for a place.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/alternateidentifiers
func (m_ MKMapItemIdentifier) SetAlternateIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlternateIdentifiers:"), value)
}

// The URL associated with the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/url
func (m_ MKMapItemIdentifier) Url() foundation.URL {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The URL associated with the specified location.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/url
func (m_ MKMapItemIdentifier) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUrl:"), value)
}

// A unique identifier for a place.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/identifier-swift.property
func (m_ MKMapItemIdentifier) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// A unique identifier for a place.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/identifier-swift.property
func (m_ MKMapItemIdentifier) SetIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), value)
}

// A Boolean value that indicates whether the map item represents the user’s location.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/iscurrentlocation
func (m_ MKMapItemIdentifier) IsCurrentLocation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCurrentLocation"))
	return rv
}


// SetIsCurrentLocation sets the value of the isCurrentLocation property.
// A Boolean value that indicates whether the map item represents the user’s location.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/iscurrentlocation
func (m_ MKMapItemIdentifier) SetIsCurrentLocation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsCurrentLocation:"), value)
}

// The placemark object containing the location information.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/placemark
func (m_ MKMapItemIdentifier) Placemark() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("placemark"))
	return rv
}


// SetPlacemark sets the value of the placemark property.
// The placemark object containing the location information.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapitem/placemark
func (m_ MKMapItemIdentifier) SetPlacemark(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlacemark:"), value)
}



