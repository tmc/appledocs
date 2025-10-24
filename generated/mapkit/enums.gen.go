// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

/* debug [enums.gen.go]: Generating 25 enums for MapKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum MKAddressFilterOption (6 cases) */
// MKAddressFilterOption - A structure that contains options for filtering results in a search.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/Options
type MKAddressFilterOption uint

const (
	// MKAddressFilterOptionAdministrativeArea - The primary administrative divisions of countries or regions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/Options/administrativeArea
	MKAddressFilterOptionAdministrativeArea MKAddressFilterOption = 0
	// MKAddressFilterOptionCountry - Countries and regions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/Options/country
	MKAddressFilterOptionCountry MKAddressFilterOption = 0
	// MKAddressFilterOptionLocality - Local administrative divisions, postal cities, and populated places.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/Options/locality
	MKAddressFilterOptionLocality MKAddressFilterOption = 0
	// MKAddressFilterOptionPostalCode - An address code for mail sorting and delivery.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/Options/postalCode
	MKAddressFilterOptionPostalCode MKAddressFilterOption = 0
	// MKAddressFilterOptionSubAdministrativeArea - The secondary administrative divisions of countries or regions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/Options/subAdministrativeArea
	MKAddressFilterOptionSubAdministrativeArea MKAddressFilterOption = 0
	// MKAddressFilterOptionSubLocality - Local administrative subdivisions, postal city subdistricts, and neighborhoods.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressFilter/Options/subLocality
	MKAddressFilterOptionSubLocality MKAddressFilterOption = 0
)

/* debug [enums.gen.go]: Processing enum MKAddressRepresentationsContextStyle (3 cases) */
// MKAddressRepresentationsContextStyle - Values that describe the degree of disambiguation context to include in an address representation.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressRepresentations/ContextStyle
type MKAddressRepresentationsContextStyle uint

const (
	// MKAddressRepresentationsContextStyleAutomatic - The value that represents the automatic context style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressRepresentations/ContextStyle/automatic
	MKAddressRepresentationsContextStyleAutomatic MKAddressRepresentationsContextStyle = 0
	// MKAddressRepresentationsContextStyleFull - The value that represents the full context style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressRepresentations/ContextStyle/full
	MKAddressRepresentationsContextStyleFull MKAddressRepresentationsContextStyle = 0
	// MKAddressRepresentationsContextStyleShort - The value that represents the short context style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAddressRepresentations/ContextStyle/short
	MKAddressRepresentationsContextStyleShort MKAddressRepresentationsContextStyle = 0
)

/* debug [enums.gen.go]: Processing enum MKAnnotationViewCollisionMode (3 cases) */
// MKAnnotationViewCollisionMode - Constants that indicates how to interpret the collision frame rectangle of an annotation view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/CollisionMode-swift.enum
type MKAnnotationViewCollisionMode uint

const (
	// MKAnnotationViewCollisionModeCircle - A constant that indicates that the annotation view uses an inscribed circle in the collision frame rectangle to determine collisions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/CollisionMode-swift.enum/circle
	MKAnnotationViewCollisionModeCircle MKAnnotationViewCollisionMode = 0
	// MKAnnotationViewCollisionModeNone - A constant indicating that collisions can’t occur.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/CollisionMode-swift.enum/none
	MKAnnotationViewCollisionModeNone MKAnnotationViewCollisionMode = 0
	// MKAnnotationViewCollisionModeRectangle - A constant that indicates that the annotation view uses the full collision frame rectangle for detecting collisions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/CollisionMode-swift.enum/rectangle
	MKAnnotationViewCollisionModeRectangle MKAnnotationViewCollisionMode = 0
)

/* debug [enums.gen.go]: Processing enum MKAnnotationViewDragState (5 cases) */
// MKAnnotationViewDragState - Constants that indicate the drag state of an annotation view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/DragState-swift.enum
type MKAnnotationViewDragState uint

const (
	// MKAnnotationViewDragStateCanceling - An annotation view cancels drag operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/DragState-swift.enum/canceling
	MKAnnotationViewDragStateCanceling MKAnnotationViewDragState = 0
	// MKAnnotationViewDragStateDragging - An annotation view is actively dragging.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/DragState-swift.enum/dragging
	MKAnnotationViewDragStateDragging MKAnnotationViewDragState = 0
	// MKAnnotationViewDragStateEnding - An annotation view ends dragging.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/DragState-swift.enum/ending
	MKAnnotationViewDragStateEnding MKAnnotationViewDragState = 0
	// MKAnnotationViewDragStateNone - An annotation view that doesn’t have a drag operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/DragState-swift.enum/none
	MKAnnotationViewDragStateNone MKAnnotationViewDragState = 0
	// MKAnnotationViewDragStateStarting - An annotation view begins dragging.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKAnnotationView/DragState-swift.enum/starting
	MKAnnotationViewDragStateStarting MKAnnotationViewDragState = 0
)

/* debug [enums.gen.go]: Processing enum MKDirectionsRoutePreference (2 cases) */
// MKDirectionsRoutePreference - Options that modify how the framework selects routes when calculating directions.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/RoutePreference
type MKDirectionsRoutePreference uint

const (
	// MKDirectionsRoutePreferenceAny - The option that specifies any available route.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/RoutePreference/any
	MKDirectionsRoutePreferenceAny MKDirectionsRoutePreference = 0
	// MKDirectionsRoutePreferenceAvoid - The option that requests the framework avoid certain routes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirections/RoutePreference/avoid
	MKDirectionsRoutePreferenceAvoid MKDirectionsRoutePreference = 0
)

/* debug [enums.gen.go]: Processing enum MKDirectionsTransportType (5 cases) */
// MKDirectionsTransportType - Constants that specify the type of conveyance to use.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirectionsTransportType
type MKDirectionsTransportType uint

const (
	// MKDirectionsTransportTypeAny - Directions suitable for any transportation option.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirectionsTransportType/any
	MKDirectionsTransportTypeAny MKDirectionsTransportType = 0
	// MKDirectionsTransportTypeAutomobile - Directions suitable for use while driving.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirectionsTransportType/automobile
	MKDirectionsTransportTypeAutomobile MKDirectionsTransportType = 0
	// MKDirectionsTransportTypeCycling - Directions suitable for use while cycling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirectionsTransportType/cycling
	MKDirectionsTransportTypeCycling MKDirectionsTransportType = 0
	// MKDirectionsTransportTypeTransit - Directions suitable for public transportation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirectionsTransportType/transit
	MKDirectionsTransportTypeTransit MKDirectionsTransportType = 0
	// MKDirectionsTransportTypeWalking - Directions suitable for a pedestrian.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDirectionsTransportType/walking
	MKDirectionsTransportTypeWalking MKDirectionsTransportType = 0
)

/* debug [enums.gen.go]: Processing enum MKDistanceFormatterUnitStyle (3 cases) */
// MKDistanceFormatterUnitStyle - Constants that indicate the format style to use for strings.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/DistanceUnitStyle
type MKDistanceFormatterUnitStyle uint

const (
	// MKDistanceFormatterUnitStyleAbbreviated - Abbreviates units.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/DistanceUnitStyle/abbreviated
	MKDistanceFormatterUnitStyleAbbreviated MKDistanceFormatterUnitStyle = 0
	// MKDistanceFormatterUnitStyleDefault - Bases the determination to abbreviate on the current locale and user language settings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/DistanceUnitStyle/default
	MKDistanceFormatterUnitStyleDefault MKDistanceFormatterUnitStyle = 0
	// MKDistanceFormatterUnitStyleFull - Spells out units in full.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/DistanceUnitStyle/full
	MKDistanceFormatterUnitStyleFull MKDistanceFormatterUnitStyle = 0
)

/* debug [enums.gen.go]: Processing enum MKDistanceFormatterUnits (4 cases) */
// MKDistanceFormatterUnits - Constants that reflect the type of units to use in the string.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/Units-swift.enum
type MKDistanceFormatterUnits uint

const (
	// MKDistanceFormatterUnitsDefault - The format uses the locale information to determine which units to use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/Units-swift.enum/default
	MKDistanceFormatterUnitsDefault MKDistanceFormatterUnits = 0
	// MKDistanceFormatterUnitsImperial - The format uses imperial units.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/Units-swift.enum/imperial
	MKDistanceFormatterUnitsImperial MKDistanceFormatterUnits = 0
	// MKDistanceFormatterUnitsImperialWithYards - The format uses imperial units that include measurements in yards.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/Units-swift.enum/imperialWithYards
	MKDistanceFormatterUnitsImperialWithYards MKDistanceFormatterUnits = 0
	// MKDistanceFormatterUnitsMetric - The format uses metric units.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKDistanceFormatter/Units-swift.enum/metric
	MKDistanceFormatterUnitsMetric MKDistanceFormatterUnits = 0
)

/* debug [enums.gen.go]: Processing enum MKErrorCode (6 cases) */
// MKErrorCode - Error constants for the MapKit framework.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKError/Code
type MKErrorCode uint

const (
	// MKErrorDecodingFailed - GeoJSON decoding failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKError/Code/decodingFailed
	MKErrorDecodingFailed MKErrorCode = 0
	// MKErrorDirectionsNotFound - The framework couldn’t find the specified directions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKError/Code/directionsNotFound
	MKErrorDirectionsNotFound MKErrorCode = 0
	// MKErrorLoadingThrottled - The data didn’t load because data throttling is in effect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKError/Code/loadingThrottled
	MKErrorLoadingThrottled MKErrorCode = 0
	// MKErrorPlacemarkNotFound - The specified placemark could not be found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKError/Code/placemarkNotFound
	MKErrorPlacemarkNotFound MKErrorCode = 0
	// MKErrorServerFailure - The map server was unable to return the desired information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKError/Code/serverFailure
	MKErrorServerFailure MKErrorCode = 0
	// MKErrorUnknown - An unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKError/Code/unknown
	MKErrorUnknown MKErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum MKFeatureVisibility (3 cases) */
// MKFeatureVisibility - Constants that indicate the visibility of different map features.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKFeatureVisibility
type MKFeatureVisibility uint

const (
	// MKFeatureVisibilityAdaptive - A constant indicating that the feature adapts to the current map state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKFeatureVisibility/adaptive
	MKFeatureVisibilityAdaptive MKFeatureVisibility = 0
	// MKFeatureVisibilityHidden - A constant indicating that the feature is hidden.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKFeatureVisibility/hidden
	MKFeatureVisibilityHidden MKFeatureVisibility = 0
	// MKFeatureVisibilityVisible - A constant indicating that the feature is visible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKFeatureVisibility/visible
	MKFeatureVisibilityVisible MKFeatureVisibility = 0
)

/* debug [enums.gen.go]: Processing enum MKLocalSearchResultType (3 cases) */
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

/* debug [enums.gen.go]: Processing enum MKSearchCompletionFilterType (2 cases) */
// MKSearchCompletionFilterType - Constants indicating the types of search completions to return.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/FilterType-swift.enum
type MKSearchCompletionFilterType uint

const (
	// MKSearchCompletionFilterTypeLocationsAndQueries - Points of interest and query suggestions. Specify this value when you want both map-based points of interest and common query terms used to find locations. For example, the search string   yields a completion for  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/FilterType-swift.enum/locationsAndQueries
	MKSearchCompletionFilterTypeLocationsAndQueries MKSearchCompletionFilterType = 0
	// MKSearchCompletionFilterTypeLocationsOnly - Points of interest only. Specify this value when you want the search string to yield completions that correspond to a specific point-of-interest on the map.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/FilterType-swift.enum/locationsOnly
	MKSearchCompletionFilterTypeLocationsOnly MKSearchCompletionFilterType = 0
)

/* debug [enums.gen.go]: Processing enum MKLocalSearchCompleterResultType (4 cases) */
// MKLocalSearchCompleterResultType - Options that indicate types of search completions.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/ResultType
type MKLocalSearchCompleterResultType uint

const (
	// MKLocalSearchCompleterResultTypeAddress - A value that indicates that the search completer includes address completions in the result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/ResultType/address
	MKLocalSearchCompleterResultTypeAddress MKLocalSearchCompleterResultType = 0
	// MKLocalSearchCompleterResultTypePhysicalFeature - A value that indicates that the search completer includes physical feature completions in the result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/ResultType/physicalFeature
	MKLocalSearchCompleterResultTypePhysicalFeature MKLocalSearchCompleterResultType = 0
	// MKLocalSearchCompleterResultTypePointOfInterest - A value that indicates that the search completer includes point-of-interest completions in the result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/ResultType/pointOfInterest
	MKLocalSearchCompleterResultTypePointOfInterest MKLocalSearchCompleterResultType = 0
	// MKLocalSearchCompleterResultTypeQuery - A value that indicates that the search completer includes query completions in the result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchCompleter/ResultType/query
	MKLocalSearchCompleterResultTypeQuery MKLocalSearchCompleterResultType = 0
)

/* debug [enums.gen.go]: Processing enum MKLocalSearchRegionPriority (2 cases) */
// MKLocalSearchRegionPriority - A value that indicates the importance of the configured region.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchRegionPriority
type MKLocalSearchRegionPriority uint

const (
	// MKLocalSearchRegionPriorityDefault - A value indicating that the results can originate from outside the specified region.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchRegionPriority/default
	MKLocalSearchRegionPriorityDefault MKLocalSearchRegionPriority = 0
	// MKLocalSearchRegionPriorityRequired - A value indicating that no results can originate from outside the specified region.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLocalSearchRegionPriority/required
	MKLocalSearchRegionPriorityRequired MKLocalSearchRegionPriority = 0
)

/* debug [enums.gen.go]: Processing enum MKLookAroundBadgePosition (3 cases) */
// MKLookAroundBadgePosition - Constants that control the position of badges on LookAround views.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundBadgePosition
type MKLookAroundBadgePosition uint

const (
	// MKLookAroundBadgePositionBottomTrailing - The value that indicates the bottom-right badge position.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundBadgePosition/bottomTrailing
	MKLookAroundBadgePositionBottomTrailing MKLookAroundBadgePosition = 0
	// MKLookAroundBadgePositionTopLeading - The value that indicates the top-left badge position.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundBadgePosition/topLeading
	MKLookAroundBadgePositionTopLeading MKLookAroundBadgePosition = 0
	// MKLookAroundBadgePositionTopTrailing - The value that indicates the top-right badge position.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundBadgePosition/topTrailing
	MKLookAroundBadgePositionTopTrailing MKLookAroundBadgePosition = 0
)

/* debug [enums.gen.go]: Processing enum MKMapElevationStyle (2 cases) */
// MKMapElevationStyle - Values that control the map’s elevation style.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapConfiguration/ElevationStyle-swift.enum
type MKMapElevationStyle uint

const (
	// MKMapElevationStyleFlat - The value that represents the flat map elevation style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapConfiguration/ElevationStyle-swift.enum/flat
	MKMapElevationStyleFlat MKMapElevationStyle = 0
	// MKMapElevationStyleRealistic - The value that represents a map elevation style with realistic ground contours.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapConfiguration/ElevationStyle-swift.enum/realistic
	MKMapElevationStyleRealistic MKMapElevationStyle = 0
)

/* debug [enums.gen.go]: Processing enum MKMapFeatureType (3 cases) */
// MKMapFeatureType - Values that describe the kinds of features visible on the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureAnnotation/FeatureType-swift.enum
type MKMapFeatureType uint

const (
	// MKMapFeatureTypePhysicalFeature - A physical feature on the Earth such as a mountain range, river, or ocean basin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureAnnotation/FeatureType-swift.enum/physicalFeature
	MKMapFeatureTypePhysicalFeature MKMapFeatureType = 0
	// MKMapFeatureTypePointOfInterest - A point of interest, such as a museum, cafe, park, or school.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureAnnotation/FeatureType-swift.enum/pointOfInterest
	MKMapFeatureTypePointOfInterest MKMapFeatureType = 0
	// MKMapFeatureTypeTerritory - A territorial or regional boundary, such as a national border, a state boundary, or a neighborhood.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureAnnotation/FeatureType-swift.enum/territory
	MKMapFeatureTypeTerritory MKMapFeatureType = 0
)

/* debug [enums.gen.go]: Processing enum MKMapFeatureOptions (3 cases) */
// MKMapFeatureOptions - A structure you use to tell the map which kinds of features users can interact with.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureOptions
type MKMapFeatureOptions uint

const (
	// MKMapFeatureOptionPhysicalFeatures - The option that represents physical map features such as mountain ranges, rivers, and ocean basins.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureOptions/physicalFeatures
	MKMapFeatureOptionPhysicalFeatures MKMapFeatureOptions = 0
	// MKMapFeatureOptionPointsOfInterest - The option that represents points of interest such as museums, cafes, parks, or schools.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureOptions/pointsOfInterest
	MKMapFeatureOptionPointsOfInterest MKMapFeatureOptions = 0
	// MKMapFeatureOptionTerritories - The option that represents territorial boundaries such as a national border, a state boundary, or a neighborhood.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapFeatureOptions/territories
	MKMapFeatureOptionTerritories MKMapFeatureOptions = 0
)

/* debug [enums.gen.go]: Processing enum MKMapType (6 cases) */
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

/* debug [enums.gen.go]: Processing enum MKOverlayLevel (2 cases) */
// MKOverlayLevel - Constants that indicate the position of overlays relative to other content.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayLevel
type MKOverlayLevel uint

const (
	// MKOverlayLevelAboveLabels - Place the overlay above map labels, shields, or point-of-interest icons but below annotations and 3D projections of buildings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayLevel/aboveLabels
	MKOverlayLevelAboveLabels MKOverlayLevel = 0
	// MKOverlayLevelAboveRoads - Place the overlay above roadways but below map labels, shields, or point-of-interest icons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayLevel/aboveRoads
	MKOverlayLevelAboveRoads MKOverlayLevel = 0
)

/* debug [enums.gen.go]: Processing enum MKPinAnnotationColor (3 cases) */
// MKPinAnnotationColor - The supported colors for pin annotations.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationColor
type MKPinAnnotationColor uint

const (
	// MKPinAnnotationColorGreen - The head of the pin is green. Green pins indicate starting points on the map.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationColor/green
	MKPinAnnotationColorGreen MKPinAnnotationColor = 0
	// MKPinAnnotationColorPurple - The head of the pin is purple. Purple pins indicate user-specified points on the map.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationColor/purple
	MKPinAnnotationColorPurple MKPinAnnotationColor = 0
	// MKPinAnnotationColorRed - The head of the pin is red. Red pins indicate destination points on the map.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPinAnnotationColor/red
	MKPinAnnotationColorRed MKPinAnnotationColor = 0
)

/* debug [enums.gen.go]: Processing enum MKScaleViewAlignment (3 cases) */
// MKScaleViewAlignment - Constants that indicate how the framework should align measurements.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKScaleView/Alignment
type MKScaleViewAlignment uint

const (
	// MKScaleViewAlignmentCenter - Scale measurements appear horizontally centered within the view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKScaleView/Alignment/center
	MKScaleViewAlignmentCenter MKScaleViewAlignment = 0
	// MKScaleViewAlignmentLeading - Scale measurements begin at the leading edge of the view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKScaleView/Alignment/leading
	MKScaleViewAlignmentLeading MKScaleViewAlignment = 0
	// MKScaleViewAlignmentTrailing - Scale measurements begin at the trailing edge of the view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKScaleView/Alignment/trailing
	MKScaleViewAlignmentTrailing MKScaleViewAlignment = 0
)

/* debug [enums.gen.go]: Processing enum MKMapItemDetailSelectionAccessoryCalloutStyle (3 cases) */
// MKMapItemDetailSelectionAccessoryCalloutStyle - The style to use for a map item detail callout presentation.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKSelectionAccessory/MapItemDetailPresentationStyle/CalloutStyle
type MKMapItemDetailSelectionAccessoryCalloutStyle uint

const (
	// MKMapItemDetailSelectionAccessoryCalloutStyleAutomatic - A value that allows the framework to choose an appropriate callout style automatically.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKSelectionAccessory/MapItemDetailPresentationStyle/CalloutStyle/automatic
	MKMapItemDetailSelectionAccessoryCalloutStyleAutomatic MKMapItemDetailSelectionAccessoryCalloutStyle = 0
	// MKMapItemDetailSelectionAccessoryCalloutStyleCompact - A compact, space-saving callout style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKSelectionAccessory/MapItemDetailPresentationStyle/CalloutStyle/compact
	MKMapItemDetailSelectionAccessoryCalloutStyleCompact MKMapItemDetailSelectionAccessoryCalloutStyle = 0
	// MKMapItemDetailSelectionAccessoryCalloutStyleFull - A rich, detailed callout style that is suitable for large map views.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKSelectionAccessory/MapItemDetailPresentationStyle/CalloutStyle/full
	MKMapItemDetailSelectionAccessoryCalloutStyleFull MKMapItemDetailSelectionAccessoryCalloutStyle = 0
)

/* debug [enums.gen.go]: Processing enum MKStandardMapEmphasisStyle (2 cases) */
// MKStandardMapEmphasisStyle - Values that control how the framework emphasizes map features.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/EmphasisStyle-swift.enum
type MKStandardMapEmphasisStyle uint

const (
	// MKStandardMapEmphasisStyleDefault - The default level of emphasis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/EmphasisStyle-swift.enum/default
	MKStandardMapEmphasisStyleDefault MKStandardMapEmphasisStyle = 0
	// MKStandardMapEmphasisStyleMuted - The muted level of emphasis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/EmphasisStyle-swift.enum/muted
	MKStandardMapEmphasisStyleMuted MKStandardMapEmphasisStyle = 0
)

/* debug [enums.gen.go]: Processing enum MKUserTrackingMode (3 cases) */
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


