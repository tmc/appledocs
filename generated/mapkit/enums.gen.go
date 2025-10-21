// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

// Enum types and constants
// MKLocalSearchResultType - Options that indicate types of search results.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/ResultType
type MKLocalSearchResultType uint

const (
// MKLocalSearchResultTypeAddress - A value that indicates that search results include addresses.
//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/ResultType/address
MKLocalSearchResultTypeAddress MKLocalSearchResultType = 0
// MKLocalSearchResultTypePhysicalFeature - A value that indicates that search results include physical features.
//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/ResultType/physicalFeature
MKLocalSearchResultTypePhysicalFeature MKLocalSearchResultType = 0
// MKLocalSearchResultTypePointOfInterest - A value that indicates that search results include points of interest.
//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearch/ResultType/pointOfInterest
MKLocalSearchResultTypePointOfInterest MKLocalSearchResultType = 0
)

// MKMapElevationStyle - Values that control the map’s elevation style.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapConfiguration/ElevationStyle-swift.enum
type MKMapElevationStyle uint

const (
// MKMapElevationStyleFlat - The value that represents the flat map elevation style.
//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapConfiguration/ElevationStyle-swift.enum/flat
MKMapElevationStyleFlat MKMapElevationStyle = 0
)

// MKMapFeatureType - Values that describe the kinds of features visible on the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureAnnotation/FeatureType-swift.enum
type MKMapFeatureType uint

// MKMapFeatureOptions - A structure you use to tell the map which kinds of features users can interact with.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureOptions
type MKMapFeatureOptions uint

// MKMapType - The type of map to display.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapType
type MKMapType uint

const (
// MKMapTypeHybrid - A satellite image of the area with road and road name information layered on top.
//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapType/hybrid
MKMapTypeHybrid MKMapType = 0
// MKMapTypeHybridFlyover - A hybrid satellite image with flyover data where available.
//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapType/hybridFlyover
MKMapTypeHybridFlyover MKMapType = 0
// MKMapTypeMutedStandard - A street map where MapKit emphasizes your data over the underlying map details.
//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapType/mutedStandard
MKMapTypeMutedStandard MKMapType = 0
// MKMapTypeSatellite - Satellite imagery of the area.
//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapType/satellite
MKMapTypeSatellite MKMapType = 0
// MKMapTypeSatelliteFlyover - A satellite image of the area with flyover data where available.
//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapType/satelliteFlyover
MKMapTypeSatelliteFlyover MKMapType = 0
// MKMapTypeStandard - A street map that shows the position of all roads and some road names.
//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapType/standard
MKMapTypeStandard MKMapType = 0
)

// MKOverlayLevel - Constants that indicate the position of overlays relative to other content.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayLevel
type MKOverlayLevel uint

// MKUserTrackingMode - The mode to use for tracking the user’s location on the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingMode
type MKUserTrackingMode uint

const (
// MKUserTrackingModeFollow - The map follows the user location.
//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingMode/follow
MKUserTrackingModeFollow MKUserTrackingMode = 0
// MKUserTrackingModeFollowWithHeading - The map follows the user’s location and rotates when the heading changes.
//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingMode/followWithHeading
MKUserTrackingModeFollowWithHeading MKUserTrackingMode = 0
// MKUserTrackingModeNone - The map doesn’t follow the user’s location.
//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingMode/none
MKUserTrackingModeNone MKUserTrackingMode = 0
)


