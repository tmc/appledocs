// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

// Enum types and constants
// MDLGeometryType - Types of geometric primitives for rendering a submesh, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLGeometryType
type MDLGeometryType uint

const (
	// MDLGeometryTypeLines - Each pair of consecutive indices in the submesh refers to two vertices to be rendered as a line segment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLGeometryType/lines
	MDLGeometryTypeLines MDLGeometryType = 0
	// MDLGeometryTypePoints - Each index in the submesh refers to a vertex to be rendered as a single point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLGeometryType/points
	MDLGeometryTypePoints MDLGeometryType = 0
	// MDLGeometryTypeQuads - Each set of four consecutive indices in the submesh refers to four vertices to be rendered as a quadrilateral.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLGeometryType/quads
	MDLGeometryTypeQuads MDLGeometryType = 0
	// MDLGeometryTypeTriangleStrips - The first three consecutive indices in the submesh refer to three vertices to be rendered as a triangle. Each subsequent index refers to another vertex that completes a triangle formed by connecting it to the previous two vertices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLGeometryType/triangleStrips
	MDLGeometryTypeTriangleStrips MDLGeometryType = 0
	// MDLGeometryTypeTriangles - Each set of three consecutive indices in the submesh refers to three vertices to be rendered as a triangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLGeometryType/triangles
	MDLGeometryTypeTriangles MDLGeometryType = 0
	// MDLGeometryTypeVariableTopology - The submesh’s index buffer does not contain a uniform set of primitives.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLGeometryType/variableTopology
	MDLGeometryTypeVariableTopology MDLGeometryType = 0
)

// MDLIndexBitDepth - Options for the size of integer data in a submesh’s index buffer, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLIndexBitDepth
type MDLIndexBitDepth uint

// MDLMaterialFace enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialFace
type MDLMaterialFace uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialFace/back
	MDLMaterialFaceBack MDLMaterialFace = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialFace/doubleSided
	MDLMaterialFaceDoubleSided MDLMaterialFace = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialFace/front
	MDLMaterialFaceFront MDLMaterialFace = 0
)

// MDLMaterialMipMapFilterMode - Modes for sampling textures at sizes between mipmap levels, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialMipMapFilterMode
type MDLMaterialMipMapFilterMode uint

const (
	// MDLMaterialMipMapFilterModeLinear - Sampling a texture at a size between mipmap levels should linearly interpolate between mipmap levels.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialMipMapFilterMode/linear
	MDLMaterialMipMapFilterModeLinear MDLMaterialMipMapFilterMode = 0
	// MDLMaterialMipMapFilterModeNearest - Sampling a texture at a size between mipmap levels should return a texel value from the nearest mipmap level.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialMipMapFilterMode/nearest
	MDLMaterialMipMapFilterModeNearest MDLMaterialMipMapFilterMode = 0
)

// MDLTransformOpRotationOrder enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTransformOpRotationOrder
type MDLTransformOpRotationOrder uint

// MDLVertexFormat - Descriptions of the data size and layout for a vertex attribute, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat
type MDLVertexFormat uint

const (
	// MDLVertexFormatInvalid - The vertex attribute has just been initialized or its format is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/invalid
	MDLVertexFormatInvalid MDLVertexFormat = 0
)


