// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders
import (
	"unsafe"
)


// C struct types
// MPSAxisAlignedBoundingBox - An axis-aligned bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAxisAlignedBoundingBox-c.struct
type MPSAxisAlignedBoundingBox struct {
	Max unsafe.Pointer
	Min unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSAxisAlignedBoundingBox */

// MPSOrigin - A position in an image used as the source origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSOrigin
type MPSOrigin struct {
	Y unsafe.Pointer // The y coordinate of the position, in pixels.
	Z unsafe.Pointer // The z coordinate of the position, in pixels.
	X unsafe.Pointer // The x coordinate of the position, in pixels.
}/* debug [types.gen.go/struct]: MPSOrigin */

// MPSPackedFloat3 - A packed three-element vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPackedFloat3-c.struct
type MPSPackedFloat3 struct {
}/* debug [types.gen.go/struct]: MPSPackedFloat3 */

// MPSSize - A size of a region in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSize
type MPSSize struct {
	Height unsafe.Pointer // The height of the region, in pixels.
	Depth unsafe.Pointer // The depth of the region, in pixels.
	Width unsafe.Pointer // The width of the region, in pixels.
}/* debug [types.gen.go/struct]: MPSSize */

// MPSCustomKernelArgumentCount - A structure that contains the number of destination, source, and broadcaset textures used by a custom kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCustomKernelArgumentCount
type MPSCustomKernelArgumentCount struct {
	BroadcastTextureCount unsafe.Pointer
	DestinationTextureCount unsafe.Pointer
	SourceTextureCount unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSCustomKernelArgumentCount */

// MPSCustomKernelInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCustomKernelInfo
type MPSCustomKernelInfo struct {
	SubbatchStride unsafe.Pointer
	Idiv IntegerDivisionParams
	ThreadgroupSize unsafe.Pointer
	SourceImageCount unsafe.Pointer
	ClipOrigin unsafe.Pointer
	ClipSize unsafe.Pointer
	DestImageArraySize unsafe.Pointer
	DestinationFeatureChannels unsafe.Pointer
	SubbatchIndex unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSCustomKernelInfo */

// MPSCustomKernelSourceInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCustomKernelSourceInfo
type MPSCustomKernelSourceInfo struct {
	Stride unsafe.Pointer
	KernelSize unsafe.Pointer
	Offset unsafe.Pointer
	ImageArraySize unsafe.Pointer
	ImageArrayOffset unsafe.Pointer
	DilationRate unsafe.Pointer
	KernelPhase unsafe.Pointer
	KernelOrigin unsafe.Pointer
	FeatureChannels unsafe.Pointer
	FeatureChannelOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSCustomKernelSourceInfo */

// MPSDimensionSlice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSDimensionSlice
type MPSDimensionSlice struct {
	Length unsafe.Pointer
	Start unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSDimensionSlice */

// MPSImageCoordinate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCoordinate
type MPSImageCoordinate struct {
	Channel unsafe.Pointer
	X unsafe.Pointer
	Y unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSImageCoordinate */

// MPSImageHistogramInfo - The information used to compute the histogram channels of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageHistogramInfo
type MPSImageHistogramInfo struct {
	MinPixelValue unsafe.Pointer // Specifies the minimum pixel value. Any pixel value less than this will be clipped to this value (for the purposes of histogram calculation), and assigned to the first histogram entry. This minimum value is applied to each of the four channels separately.
	NumberOfHistogramEntries unsafe.Pointer // Specifies the number of histogram entries ( ) for each channel.
	HistogramForAlpha unsafe.Pointer // Specifies whether the histogram for the alpha channel should be computed or not.
	MaxPixelValue unsafe.Pointer // Specifies the maximum pixel value.  Any pixel value greater than this will be clipped to this value (for the purposes of histogram calculation), and assigned to the first histogram entry. This maximum value is applied to each of the four channels separately.
}/* debug [types.gen.go/struct]: MPSImageHistogramInfo */

// MPSImageKeypointData - A structure that specifies keypoint information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageKeypointData
type MPSImageKeypointData struct {
	KeypointColorValue unsafe.Pointer
	KeypointCoordinate unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSImageKeypointData */

// MPSImageKeypointRangeInfo - A structure that specifies information to find the keypoints in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageKeypointRangeInfo
type MPSImageKeypointRangeInfo struct {
	MinimumThresholdValue unsafe.Pointer
	MaximumKeypoints unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSImageKeypointRangeInfo */

// MPSImageReadWriteParams - Parameters that control reading and writing of a particular set of feature channels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageReadWriteParams
type MPSImageReadWriteParams struct {
	FeatureChannelOffset unsafe.Pointer
	NumberOfFeatureChannelsToReadWrite unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSImageReadWriteParams */

// MPSImageRegion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageRegion
type MPSImageRegion struct {
	Offset ImageCoordinate
	Size ImageCoordinate
}/* debug [types.gen.go/struct]: MPSImageRegion */

// MPSIntegerDivisionParams - Parameters that define the parts of a division operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntegerDivisionParams
type MPSIntegerDivisionParams struct {
	Shift unsafe.Pointer
	Divisor unsafe.Pointer
	Recip unsafe.Pointer
	Addend unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSIntegerDivisionParams */

// MPSIntersectionDistance - An intersection result that contains the distance from the ray origin to the intersection point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistance
type MPSIntersectionDistance struct {
	Distance unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSIntersectionDistance */

// MPSIntersectionDistancePrimitiveIndex - An intersection result that contains the distance from the ray origin to the intersection point, and the index of the intersected primitive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndex
type MPSIntersectionDistancePrimitiveIndex struct {
	Distance unsafe.Pointer
	PrimitiveIndex unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSIntersectionDistancePrimitiveIndex */

// MPSIntersectionDistancePrimitiveIndexBufferIndex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexBufferIndex
type MPSIntersectionDistancePrimitiveIndexBufferIndex struct {
	BufferIndex unsafe.Pointer
	Distance unsafe.Pointer
	PrimitiveIndex unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSIntersectionDistancePrimitiveIndexBufferIndex */

// MPSIntersectionDistancePrimitiveIndexBufferIndexCoordinates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexBufferIndexCoordinates
type MPSIntersectionDistancePrimitiveIndexBufferIndexCoordinates struct {
	BufferIndex unsafe.Pointer
	Coordinates unsafe.Pointer
	Distance unsafe.Pointer
	PrimitiveIndex unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSIntersectionDistancePrimitiveIndexBufferIndexCoordinates */

// MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndex
type MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndex struct {
	BufferIndex unsafe.Pointer
	Distance unsafe.Pointer
	InstanceIndex unsafe.Pointer
	PrimitiveIndex unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndex */

// MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndexCoordinates
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndexCoordinates
type MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndexCoordinates struct {
	BufferIndex unsafe.Pointer
	Coordinates unsafe.Pointer
	Distance unsafe.Pointer
	InstanceIndex unsafe.Pointer
	PrimitiveIndex unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSIntersectionDistancePrimitiveIndexBufferIndexInstanceIndexCoordinates */

// MPSIntersectionDistancePrimitiveIndexCoordinates - An intersection result that contains the origin-intersection distance, intersected primitive index, and intersection point coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexCoordinates
type MPSIntersectionDistancePrimitiveIndexCoordinates struct {
	Coordinates unsafe.Pointer
	Distance unsafe.Pointer
	PrimitiveIndex unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSIntersectionDistancePrimitiveIndexCoordinates */

// MPSIntersectionDistancePrimitiveIndexInstanceIndex - An intersection result that contains the origin-intersection distance, and intersected primitive and instance indices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexInstanceIndex
type MPSIntersectionDistancePrimitiveIndexInstanceIndex struct {
	Distance unsafe.Pointer
	InstanceIndex unsafe.Pointer
	PrimitiveIndex unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSIntersectionDistancePrimitiveIndexInstanceIndex */

// MPSIntersectionDistancePrimitiveIndexInstanceIndexCoordinates - An intersection result that contains the origin-intersection distance, intersected primitive and instance indices, and intersection point coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSIntersectionDistancePrimitiveIndexInstanceIndexCoordinates
type MPSIntersectionDistancePrimitiveIndexInstanceIndexCoordinates struct {
	Coordinates unsafe.Pointer
	Distance unsafe.Pointer
	InstanceIndex unsafe.Pointer
	PrimitiveIndex unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSIntersectionDistancePrimitiveIndexInstanceIndexCoordinates */

// MPSMatrixCopyOffsets - A description of matrix copy operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixCopyOffsets
type MPSMatrixCopyOffsets struct {
	SourceColumnOffset unsafe.Pointer
	DestinationColumnOffset unsafe.Pointer
	SourceRowOffset unsafe.Pointer
	DestinationRowOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSMatrixCopyOffsets */

// MPSMatrixOffset - A description of row and column offsets into a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixOffset
type MPSMatrixOffset struct {
	ColumnOffset unsafe.Pointer
	RowOffset unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSMatrixOffset */

// MPSNDArrayOffsets
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayOffsets
type MPSNDArrayOffsets struct {
	Dimensions unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSNDArrayOffsets */

// MPSNDArraySizes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArraySizes
type MPSNDArraySizes struct {
	Dimensions unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSNDArraySizes */

// MPSOffset - A signed coordinate with x, y, and z components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSOffset
type MPSOffset struct {
	Y unsafe.Pointer // The vertical component of the offset, in pixels.
	X unsafe.Pointer // The horizontal component of the offset, in pixels.
	Z unsafe.Pointer // The depth component of the offset, in pixels.
}/* debug [types.gen.go/struct]: MPSOffset */

// MPSRayOriginDirection - A 3D ray with an origin and a direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayOriginDirection
type MPSRayOriginDirection struct {
	Direction unsafe.Pointer
	Origin unsafe.Pointer
}/* debug [types.gen.go/struct]: MPSRayOriginDirection */

// MPSRayOriginMaskDirectionMaxDistance - A 3D ray with an origin, a direction, and a mask to filter out intersections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayOriginMaskDirectionMaxDistance
type MPSRayOriginMaskDirectionMaxDistance struct {
	Direction PackedFloat3
	Mask unsafe.Pointer
	MaxDistance unsafe.Pointer
	Origin PackedFloat3
}/* debug [types.gen.go/struct]: MPSRayOriginMaskDirectionMaxDistance */

// MPSRayOriginMinDistanceDirectionMaxDistance - A 3D ray with an origin, a direction, and an intersection distance range from the origin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayOriginMinDistanceDirectionMaxDistance
type MPSRayOriginMinDistanceDirectionMaxDistance struct {
	Direction PackedFloat3
	MaxDistance unsafe.Pointer
	MinDistance unsafe.Pointer
	Origin PackedFloat3
}/* debug [types.gen.go/struct]: MPSRayOriginMinDistanceDirectionMaxDistance */

// MPSRayPackedOriginDirection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayPackedOriginDirection
type MPSRayPackedOriginDirection struct {
	Direction PackedFloat3
	Origin PackedFloat3
}/* debug [types.gen.go/struct]: MPSRayPackedOriginDirection */

// MPSRegion - A region of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRegion
type MPSRegion struct {
	Size Size // The size of the region.
	Origin Origin // The top-left corner of the region.
}/* debug [types.gen.go/struct]: MPSRegion */

// MPSScaleTransform - A transform matrix for explicit resampling control with a Lanczos kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSScaleTransform
type MPSScaleTransform struct {
	ScaleX unsafe.Pointer // The horizontal scale factor.
	ScaleY unsafe.Pointer // The vertical scale factor.
	TranslateX unsafe.Pointer // The horizontal translation factor.
	TranslateY unsafe.Pointer // The vertical translation factor.
}/* debug [types.gen.go/struct]: MPSScaleTransform */

// MPSStateTextureInfo - An encapsulation of a texture’s dimensions, format, type, and usage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateTextureInfo
type MPSStateTextureInfo struct {
	Depth unsafe.Pointer
	PixelFormat PixelFormat
	Usage TextureUsage
	Height unsafe.Pointer
	Width unsafe.Pointer
	ArrayLength unsafe.Pointer
	TextureType TextureType
}/* debug [types.gen.go/struct]: MPSStateTextureInfo */





