// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio
import (
	"unsafe"
)


// C struct types
// MDLAxisAlignedBoundingBox - The minimal volume containing an object, used by the 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLAxisAlignedBoundingBox
type MDLAxisAlignedBoundingBox struct {
	MaxBounds unsafe.Pointer // The corner of the bounding box with the highest x-, y-, and z-coordinate values.
	MinBounds unsafe.Pointer // The corner of the bounding box with the lowest x-, y-, and z-coordinate values.
}/* debug [types.gen.go/struct]: MDLAxisAlignedBoundingBox */

// MDLVoxelIndexExtent - The corner voxel indices defining a solid rectangular volume of voxels. Used by the 
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVoxelIndexExtent
type MDLVoxelIndexExtent struct {
	MaximumExtent MDLVoxelIndex // The highest x, y, and z coordinates in the volume.
	MinimumExtent MDLVoxelIndex // The lowest x, y, and z coordinates in the volume.
}/* debug [types.gen.go/struct]: MDLVoxelIndexExtent */





