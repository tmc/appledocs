// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit
import (
	"unsafe"
)


// C struct types
// GKBox - The definition of an axis-aligned rectangular bounding volume addressed by the tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKBox
type GKBox struct {
	BoxMax unsafe.Pointer // The corner of the box with the highest coordinate values (in most coordinate systems, the near-upper-right corner).
	BoxMin unsafe.Pointer // The corner of the box with the lowest coordinate values (in most coordinate systems, the far-lower-left corner).
}/* debug [types.gen.go/struct]: GKBox */

// GKQuad - The definition of an axis-aligned rectangle addressed by the tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuad
type GKQuad struct {
	QuadMax unsafe.Pointer // The corner of the rectangle with the highest coordinate values (in most coordinate systems, the upper-right corner).
	QuadMin unsafe.Pointer // The corner of the rectangle with the lowest coordinate values (in most coordinate systems, the lower-left corner).
}/* debug [types.gen.go/struct]: GKQuad */

// GKTriangle - The definition of a triangle in the mesh, available with the 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKTriangle
type GKTriangle struct {
	Points unsafe.Pointer // A set of three points describing the triangle.
}/* debug [types.gen.go/struct]: GKTriangle */





