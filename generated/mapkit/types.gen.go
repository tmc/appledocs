// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit


// C struct types
// MKCoordinateRegion - A rectangular geographic region that centers around a specific latitude and longitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCoordinateRegion
type MKCoordinateRegion struct {
	Center LocationCoordinate2D // The center point of the region.
	Span MKCoordinateSpan // The horizontal and vertical span representing the amount of map to display.
}/* debug [types.gen.go/struct]: MKCoordinateRegion */

// MKCoordinateSpan - The width and height of a map region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCoordinateSpan
type MKCoordinateSpan struct {
	LatitudeDelta LocationDegrees // The amount of north-to-south distance (measured in degrees) to display on the map.
	LongitudeDelta LocationDegrees // The amount of east-to-west distance (measured in degrees) to display for the map region.
}/* debug [types.gen.go/struct]: MKCoordinateSpan */

// MKMapPoint - A point on a two-dimensional map projection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapPoint
type MKMapPoint struct {
	X float64 // The location of the point along the x-axis of the map.
	Y float64 // The location of the point along the y-axis of the map.
}/* debug [types.gen.go/struct]: MKMapPoint */

// MKMapRect - A rectangular area on a two-dimensional map projection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapRect
type MKMapRect struct {
	Origin MKMapPoint // The origin point of the rectangle.
	Size MKMapSize // The width and height of the rectangle, starting from the origin point.
}/* debug [types.gen.go/struct]: MKMapRect */

// MKMapSize - Width and height information on a two-dimensional map projection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapSize
type MKMapSize struct {
	Height float64 // The height of the specified area, measured in map points.
	Width float64 // The width of the specified area, measured in map points.
}/* debug [types.gen.go/struct]: MKMapSize */

// MKTileOverlayPath - Values that specify the path indexes for a single overlay tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKTileOverlayPath
type MKTileOverlayPath struct {
	ContentScaleFactor float64 // The tile’s intended screen scale factor.
	X int // The index of the tile along the x-axis of the map.
	Y int // The index of the tile along the y-axis of the map.
	Z int // The index of the tile along the z-axis of the map.
}/* debug [types.gen.go/struct]: MKTileOverlayPath */





