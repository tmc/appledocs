// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

/* debug [enums.gen.go]: Generating 17 enums for ModelIO */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum MDLAnimatedValueInterpolation (2 cases) */
// MDLAnimatedValueInterpolation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLAnimatedValueInterpolation
type MDLAnimatedValueInterpolation uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLAnimatedValueInterpolation/constant
	MDLAnimatedValueInterpolationConstant MDLAnimatedValueInterpolation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLAnimatedValueInterpolation/linear
	MDLAnimatedValueInterpolationLinear MDLAnimatedValueInterpolation = 0
)

/* debug [enums.gen.go]: Processing enum MDLCameraProjection (2 cases) */
// MDLCameraProjection - Options for camera projection styles, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLCameraProjection
type MDLCameraProjection uint

const (
	// MDLCameraProjectionOrthographic - An orthographic projection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLCameraProjection/orthographic
	MDLCameraProjectionOrthographic MDLCameraProjection = 0
	// MDLCameraProjectionPerspective - A perspective projection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLCameraProjection/perspective
	MDLCameraProjectionPerspective MDLCameraProjection = 0
)

/* debug [enums.gen.go]: Processing enum MDLDataPrecision (3 cases) */
// MDLDataPrecision enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLDataPrecision
type MDLDataPrecision uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLDataPrecision/double
	MDLDataPrecisionDouble MDLDataPrecision = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLDataPrecision/float
	MDLDataPrecisionFloat MDLDataPrecision = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLDataPrecision/undefined
	MDLDataPrecisionUndefined MDLDataPrecision = 0
)

/* debug [enums.gen.go]: Processing enum MDLGeometryType (6 cases) */
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
	// MDLGeometryTypeTriangles - Each set of three consecutive indices in the submesh refers to three vertices to be rendered as a triangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLGeometryType/triangles
	MDLGeometryTypeTriangles MDLGeometryType = 0
	// MDLGeometryTypeTriangleStrips - The first three consecutive indices in the submesh refer to three vertices to be rendered as a triangle. Each subsequent index refers to another vertex that completes a triangle formed by connecting it to the previous two vertices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLGeometryType/triangleStrips
	MDLGeometryTypeTriangleStrips MDLGeometryType = 0
	// MDLGeometryTypeVariableTopology - The submesh’s index buffer does not contain a uniform set of primitives.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLGeometryType/variableTopology
	MDLGeometryTypeVariableTopology MDLGeometryType = 0
)

/* debug [enums.gen.go]: Processing enum MDLIndexBitDepth (7 cases) */
// MDLIndexBitDepth - Options for the size of integer data in a submesh’s index buffer, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLIndexBitDepth
type MDLIndexBitDepth uint

const (
	// MDLIndexBitDepthInvalid - The submesh has not been initialized or its data type is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLIndexBitDepth/invalid
	MDLIndexBitDepthInvalid MDLIndexBitDepth = 0
	// MDLIndexBitDepthUInt16 - Each index in the submesh’s index buffer is a 16-bit integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLIndexBitDepth/uInt16-swift.enum.case
	MDLIndexBitDepthUInt16 MDLIndexBitDepth = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLIndexBitDepth/uint16-swift.type.property
	MDLIndexBitDepthUint16 MDLIndexBitDepth = 0
	// MDLIndexBitDepthUInt32 - Each index in the submesh’s index buffer is a 32-bit integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLIndexBitDepth/uInt32-swift.enum.case
	MDLIndexBitDepthUInt32 MDLIndexBitDepth = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLIndexBitDepth/uint32-swift.type.property
	MDLIndexBitDepthUint32 MDLIndexBitDepth = 0
	// MDLIndexBitDepthUInt8 - Each index in the submesh’s index buffer is an 8-bit integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLIndexBitDepth/uInt8-swift.enum.case
	MDLIndexBitDepthUInt8 MDLIndexBitDepth = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLIndexBitDepth/uint8-swift.type.property
	MDLIndexBitDepthUint8 MDLIndexBitDepth = 0
)

/* debug [enums.gen.go]: Processing enum MDLLightType (12 cases) */
// MDLLightType - Options for the shape and style of illumination provided by a light, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLLightType
type MDLLightType uint

const (
	// MDLLightTypeAmbient - The light source should illuminate a scene evenly regardless of position or direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLLightType/ambient
	MDLLightTypeAmbient MDLLightType = 0
	// MDLLightTypeDirectional - The light source illuminates a scene from a uniform direction regardless of its position.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLLightType/directional
	MDLLightTypeDirectional MDLLightType = 0
	// MDLLightTypeDiscArea - The light source illuminates a scene in all directions from an area in the shape of a disc.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLLightType/discArea
	MDLLightTypeDiscArea MDLLightType = 0
	// MDLLightTypeEnvironment - The illumination from the light is determined by texture images representing a sample of the surrounding environment for a scene.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLLightType/environment
	MDLLightTypeEnvironment MDLLightType = 0
	// MDLLightTypeLinear - The light source illuminates a scene in all directions from an area in the shape of a line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLLightType/linear
	MDLLightTypeLinear MDLLightType = 0
	// MDLLightTypePhotometric - The illumination from the light is determined by a photometric profile.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLLightType/photometric
	MDLLightTypePhotometric MDLLightType = 0
	// MDLLightTypePoint - The light source illuminates a scene in all directions from a specific position.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLLightType/point
	MDLLightTypePoint MDLLightType = 0
	// MDLLightTypeProbe - The illumination from the light is determined by texture images representing a sample of a scene at a specific point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLLightType/probe
	MDLLightTypeProbe MDLLightType = 0
	// MDLLightTypeRectangularArea - The light source illuminates a scene in all directions from an area in the shape of a rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLLightType/rectangularArea
	MDLLightTypeRectangularArea MDLLightType = 0
	// MDLLightTypeSpot - The light source illuminates a scene from a specific position and direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLLightType/spot
	MDLLightTypeSpot MDLLightType = 0
	// MDLLightTypeSuperElliptical - The light source illuminates a scene in all directions from an area in the shape of a superellipse.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLLightType/superElliptical
	MDLLightTypeSuperElliptical MDLLightType = 0
	// MDLLightTypeUnknown - The type of the light is unknown or has not been initialized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLLightType/unknown
	MDLLightTypeUnknown MDLLightType = 0
)

/* debug [enums.gen.go]: Processing enum MDLMaterialFace (3 cases) */
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

/* debug [enums.gen.go]: Processing enum MDLMaterialMipMapFilterMode (2 cases) */
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

/* debug [enums.gen.go]: Processing enum MDLMaterialPropertyType (11 cases) */
// MDLMaterialPropertyType - Options for the data type of a material property, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialPropertyType
type MDLMaterialPropertyType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialPropertyType/buffer
	MDLMaterialPropertyTypeBuffer MDLMaterialPropertyType = 0
	// MDLMaterialPropertyTypeColor - The material property’s value is a uniform color.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialPropertyType/color
	MDLMaterialPropertyTypeColor MDLMaterialPropertyType = 0
	// MDLMaterialPropertyTypeFloat - The material property’s value is a floating-point scalar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialPropertyType/float
	MDLMaterialPropertyTypeFloat MDLMaterialPropertyType = 0
	// MDLMaterialPropertyTypeFloat2 - The material property’s value is a 2-component floating-point vector.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialPropertyType/float2
	MDLMaterialPropertyTypeFloat2 MDLMaterialPropertyType = 0
	// MDLMaterialPropertyTypeFloat3 - The material property’s value is a 3-component floating-point vector.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialPropertyType/float3
	MDLMaterialPropertyTypeFloat3 MDLMaterialPropertyType = 0
	// MDLMaterialPropertyTypeFloat4 - The material property’s value is a 4-component floating-point vector.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialPropertyType/float4
	MDLMaterialPropertyTypeFloat4 MDLMaterialPropertyType = 0
	// MDLMaterialPropertyTypeMatrix44 - The material property’s value is a 4 x 4 floating-point matrix.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialPropertyType/matrix44
	MDLMaterialPropertyTypeMatrix44 MDLMaterialPropertyType = 0
	// MDLMaterialPropertyTypeNone - The material property’s value has not been initialized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialPropertyType/none
	MDLMaterialPropertyTypeNone MDLMaterialPropertyType = 0
	// MDLMaterialPropertyTypeString - The material’s value is a string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialPropertyType/string
	MDLMaterialPropertyTypeString MDLMaterialPropertyType = 0
	// MDLMaterialPropertyTypeTexture - The material property’s value is a   object that provides both a texture image and texture rendering parameters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialPropertyType/texture
	MDLMaterialPropertyTypeTexture MDLMaterialPropertyType = 0
	// MDLMaterialPropertyTypeURL - The material property’s value is a URL—typically, a URL referencing a texture image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialPropertyType/URL
	MDLMaterialPropertyTypeURL MDLMaterialPropertyType = 0
)

/* debug [enums.gen.go]: Processing enum MDLMaterialSemantic (26 cases) */
// MDLMaterialSemantic - Options for the semantic use of a material property’s value in rendering a particular surface appearance; used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic
type MDLMaterialSemantic uint

const (
	// MDLMaterialSemanticAmbientOcclusion - The attenuation of ambient light due to local geometry variations on a surface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/ambientOcclusion
	MDLMaterialSemanticAmbientOcclusion MDLMaterialSemantic = 0
	// MDLMaterialSemanticAmbientOcclusionScale - The scaling factor for ambient occlusion shading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/ambientOcclusionScale
	MDLMaterialSemanticAmbientOcclusionScale MDLMaterialSemantic = 0
	// MDLMaterialSemanticAnisotropic - The degree to which specular highlights elongate in the direction of the local tangent basis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/anisotropic
	MDLMaterialSemanticAnisotropic MDLMaterialSemantic = 0
	// MDLMaterialSemanticAnisotropicRotation - The angle at which anisotropic effects are rotated relative to the local tangent basis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/anisotropicRotation
	MDLMaterialSemanticAnisotropicRotation MDLMaterialSemantic = 0
	// MDLMaterialSemanticBaseColor - The inherent color of a surface, to be used as a modulator during shading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/baseColor
	MDLMaterialSemanticBaseColor MDLMaterialSemantic = 0
	// MDLMaterialSemanticBump - The degree of perturbation in a material’s surface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/bump
	MDLMaterialSemanticBump MDLMaterialSemantic = 0
	// MDLMaterialSemanticClearcoat - The intensity of a second specular highlight, similar to the gloss that results from a clear coat on an automotive finish.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/clearcoat
	MDLMaterialSemanticClearcoat MDLMaterialSemantic = 0
	// MDLMaterialSemanticClearcoatGloss - The spread of a second specular highlight, similar to the gloss that results from a clear coat on an automotive finish.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/clearcoatGloss
	MDLMaterialSemanticClearcoatGloss MDLMaterialSemantic = 0
	// MDLMaterialSemanticDisplacement - The displacement of a material’s surface relative to the surface normal.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/displacement
	MDLMaterialSemanticDisplacement MDLMaterialSemantic = 0
	// MDLMaterialSemanticDisplacementScale - The scaling factor for displacement of a material’s surface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/displacementScale
	MDLMaterialSemanticDisplacementScale MDLMaterialSemantic = 0
	// MDLMaterialSemanticEmission - The color emitted as radiance from a material’s surface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/emission
	MDLMaterialSemanticEmission MDLMaterialSemantic = 0
	// MDLMaterialSemanticInterfaceIndexOfRefraction - The index of refraction for the medium surrounding a material.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/interfaceIndexOfRefraction
	MDLMaterialSemanticInterfaceIndexOfRefraction MDLMaterialSemantic = 0
	// MDLMaterialSemanticMaterialIndexOfRefraction - The index of refraction for a material itself.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/materialIndexOfRefraction
	MDLMaterialSemanticMaterialIndexOfRefraction MDLMaterialSemantic = 0
	// MDLMaterialSemanticMetallic - The degree to which a material appears as a dielectric surface (lower values) or as a metal (higher values).
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/metallic
	MDLMaterialSemanticMetallic MDLMaterialSemantic = 0
	// MDLMaterialSemanticNone - The material property’s   property has not been initialized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/none
	MDLMaterialSemanticNone MDLMaterialSemantic = 0
	// MDLMaterialSemanticObjectSpaceNormal - The variation in the surface normal vectors in a material, relative to model coordinate space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/objectSpaceNormal
	MDLMaterialSemanticObjectSpaceNormal MDLMaterialSemantic = 0
	// MDLMaterialSemanticOpacity - The opacity of a material’s surface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/opacity
	MDLMaterialSemanticOpacity MDLMaterialSemantic = 0
	// MDLMaterialSemanticRoughness - The degree to which a material appears smooth, affecting both diffuse and specular response.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/roughness
	MDLMaterialSemanticRoughness MDLMaterialSemantic = 0
	// MDLMaterialSemanticSheen - The intensity of highlights that appear only at glancing angles on a material’s surface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/sheen
	MDLMaterialSemanticSheen MDLMaterialSemantic = 0
	// MDLMaterialSemanticSheenTint - The balance of color for highlights that appear only at glancing angles, between the light color (lower values) and the material’s base color (at higher values).
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/sheenTint
	MDLMaterialSemanticSheenTint MDLMaterialSemantic = 0
	// MDLMaterialSemanticSpecular - The intensity of specular highlights that appear on the material’s surface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/specular
	MDLMaterialSemanticSpecular MDLMaterialSemantic = 0
	// MDLMaterialSemanticSpecularExponent - The exponent to be used in Blinn-Phong approximation of the material’s specular response.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/specularExponent
	MDLMaterialSemanticSpecularExponent MDLMaterialSemantic = 0
	// MDLMaterialSemanticSpecularTint - The balance of color for specular highlights, between the light color (lower values) and the material’s base color (at higher values).
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/specularTint
	MDLMaterialSemanticSpecularTint MDLMaterialSemantic = 0
	// MDLMaterialSemanticSubsurface - The degree to which light scatters under the surface of a material.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/subsurface
	MDLMaterialSemanticSubsurface MDLMaterialSemantic = 0
	// MDLMaterialSemanticTangentSpaceNormal - The variation in the surface normal vectors in a material, relative to surface tangent coordinate space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/tangentSpaceNormal
	MDLMaterialSemanticTangentSpaceNormal MDLMaterialSemantic = 0
	// MDLMaterialSemanticUserDefined - The meaning of the material property’s value is not one of the standard semantic uses recognized by Model I/O.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic/userDefined
	MDLMaterialSemanticUserDefined MDLMaterialSemantic = 0
)

/* debug [enums.gen.go]: Processing enum MDLMaterialTextureFilterMode (2 cases) */
// MDLMaterialTextureFilterMode - Modes for sampling textures at coordinates between texels, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialTextureFilterMode
type MDLMaterialTextureFilterMode uint

const (
	// MDLMaterialTextureFilterModeLinear - Sampling at texture coordinates between texels should linearly interpolate between texel values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialTextureFilterMode/linear
	MDLMaterialTextureFilterModeLinear MDLMaterialTextureFilterMode = 0
	// MDLMaterialTextureFilterModeNearest - Sampling at texture coordinates between texels should return the value of the nearest texel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialTextureFilterMode/nearest
	MDLMaterialTextureFilterModeNearest MDLMaterialTextureFilterMode = 0
)

/* debug [enums.gen.go]: Processing enum MDLMaterialTextureWrapMode (3 cases) */
// MDLMaterialTextureWrapMode - Modes for sampling textures at coordinates outside the texture bounds, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialTextureWrapMode
type MDLMaterialTextureWrapMode uint

const (
	// MDLMaterialTextureWrapModeClamp - Sampling at any texture coordinate outside the   to   range returns the texel color from the nearest edge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialTextureWrapMode/clamp
	MDLMaterialTextureWrapModeClamp MDLMaterialTextureWrapMode = 0
	// MDLMaterialTextureWrapModeMirror - Sampling at texture coordinates outside the   to   range results in a mirrored tiling effect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialTextureWrapMode/mirror
	MDLMaterialTextureWrapModeMirror MDLMaterialTextureWrapMode = 0
	// MDLMaterialTextureWrapModeRepeat - Sampling at texture coordinates outside the   to   range results in a repeated tiling effect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMaterialTextureWrapMode/repeat
	MDLMaterialTextureWrapModeRepeat MDLMaterialTextureWrapMode = 0
)

/* debug [enums.gen.go]: Processing enum MDLMeshBufferType (3 cases) */
// MDLMeshBufferType - Options for the content of a mesh buffer, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMeshBufferType
type MDLMeshBufferType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMeshBufferType/custom
	MDLMeshBufferTypeCustom MDLMeshBufferType = 0
	// MDLMeshBufferTypeIndex - The buffer contains index data for a   object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMeshBufferType/index
	MDLMeshBufferTypeIndex MDLMeshBufferType = 0
	// MDLMeshBufferTypeVertex - The buffer contains per-vertex data for one or more vertex attributes of a   object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLMeshBufferType/vertex
	MDLMeshBufferTypeVertex MDLMeshBufferType = 0
)

/* debug [enums.gen.go]: Processing enum MDLProbePlacement (2 cases) */
// MDLProbePlacement - Options affecting automatic placement of light probes in a scene, used with the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLProbePlacement
type MDLProbePlacement uint

const (
	// MDLProbePlacementIrradianceDistribution - An option to examine the lighting conditions at various positions in the scene being evaluated, then place light probes only at the locations where each contributes optimally to scene lighting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLProbePlacement/irradianceDistribution
	MDLProbePlacementIrradianceDistribution MDLProbePlacement = 0
	// MDLProbePlacementUniformGrid - An option to place light probes at each unit coordinate in a three-dimensional grid that evenly divides the region being evaluated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLProbePlacement/uniformGrid
	MDLProbePlacementUniformGrid MDLProbePlacement = 0
)

/* debug [enums.gen.go]: Processing enum MDLTextureChannelEncoding (11 cases) */
// MDLTextureChannelEncoding - Options for the data size and type of texel channel values, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTextureChannelEncoding
type MDLTextureChannelEncoding uint

const (
	// MDLTextureChannelEncodingFloat16 - Each channel value per texel is a 16-bit floating-point value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTextureChannelEncoding/float16
	MDLTextureChannelEncodingFloat16 MDLTextureChannelEncoding = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTextureChannelEncoding/float16SR
	MDLTextureChannelEncodingFloat16SR MDLTextureChannelEncoding = 0
	// MDLTextureChannelEncodingFloat32 - Each channel value per texel is a 32-bit floating-point value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTextureChannelEncoding/float32
	MDLTextureChannelEncodingFloat32 MDLTextureChannelEncoding = 0
	// MDLTextureChannelEncodingUInt16 - Each channel value per texel is a 16-bit unsigned integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTextureChannelEncoding/uInt16-swift.enum.case
	MDLTextureChannelEncodingUInt16 MDLTextureChannelEncoding = 0
	// MDLTextureChannelEncodingUint16 - Each channel value per texel is a 16-bit unsigned integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTextureChannelEncoding/uint16-swift.type.property
	MDLTextureChannelEncodingUint16 MDLTextureChannelEncoding = 0
	// MDLTextureChannelEncodingUInt24 - Each channel value per texel is a 24-bit unsigned integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTextureChannelEncoding/uInt24-swift.enum.case
	MDLTextureChannelEncodingUInt24 MDLTextureChannelEncoding = 0
	// MDLTextureChannelEncodingUint24 - Each channel value per texel is a 24-bit unsigned integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTextureChannelEncoding/uint24-swift.type.property
	MDLTextureChannelEncodingUint24 MDLTextureChannelEncoding = 0
	// MDLTextureChannelEncodingUInt32 - Each channel value per texel is a 32-bit unsigned integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTextureChannelEncoding/uInt32-swift.enum.case
	MDLTextureChannelEncodingUInt32 MDLTextureChannelEncoding = 0
	// MDLTextureChannelEncodingUint32 - Each channel value per texel is a 32-bit unsigned integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTextureChannelEncoding/uint32-swift.type.property
	MDLTextureChannelEncodingUint32 MDLTextureChannelEncoding = 0
	// MDLTextureChannelEncodingUInt8 - Each channel value per texel is an 8-bit unsigned integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTextureChannelEncoding/uInt8-swift.enum.case
	MDLTextureChannelEncodingUInt8 MDLTextureChannelEncoding = 0
	// MDLTextureChannelEncodingUint8 - Each channel value per texel is an 8-bit unsigned integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTextureChannelEncoding/uint8-swift.type.property
	MDLTextureChannelEncodingUint8 MDLTextureChannelEncoding = 0
)

/* debug [enums.gen.go]: Processing enum MDLTransformOpRotationOrder (6 cases) */
// MDLTransformOpRotationOrder enum type
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTransformOpRotationOrder
type MDLTransformOpRotationOrder uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTransformOpRotationOrder/XYZ
	MDLTransformOpRotationOrderXYZ MDLTransformOpRotationOrder = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTransformOpRotationOrder/XZY
	MDLTransformOpRotationOrderXZY MDLTransformOpRotationOrder = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTransformOpRotationOrder/YXZ
	MDLTransformOpRotationOrderYXZ MDLTransformOpRotationOrder = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTransformOpRotationOrder/YZX
	MDLTransformOpRotationOrderYZX MDLTransformOpRotationOrder = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTransformOpRotationOrder/ZXY
	MDLTransformOpRotationOrderZXY MDLTransformOpRotationOrder = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLTransformOpRotationOrder/ZYX
	MDLTransformOpRotationOrderZYX MDLTransformOpRotationOrder = 0
)

/* debug [enums.gen.go]: Processing enum MDLVertexFormat (64 cases) */
// MDLVertexFormat - Descriptions of the data size and layout for a vertex attribute, used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat
type MDLVertexFormat uint

const (
	// MDLVertexFormatChar - The attribute value for each vertex is a scalar of signed 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/char
	MDLVertexFormatChar MDLVertexFormat = 0
	// MDLVertexFormatChar2 - The attribute value for each vertex is a vector with 2 components, each of signed 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/char2
	MDLVertexFormatChar2 MDLVertexFormat = 0
	// MDLVertexFormatChar2Normalized - The attribute value for each vertex is a vector with 2 components, each with a normalized value of signed 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/char2Normalized
	MDLVertexFormatChar2Normalized MDLVertexFormat = 0
	// MDLVertexFormatChar3 - The attribute value for each vertex is a vector with 3 components, each of signed 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/char3
	MDLVertexFormatChar3 MDLVertexFormat = 0
	// MDLVertexFormatChar3Normalized - The attribute value for each vertex is a vector with 3 components, each with a normalized value of signed 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/char3Normalized
	MDLVertexFormatChar3Normalized MDLVertexFormat = 0
	// MDLVertexFormatChar4 - The attribute value for each vertex is a vector with 4 components, each of signed 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/char4
	MDLVertexFormatChar4 MDLVertexFormat = 0
	// MDLVertexFormatChar4Normalized - The attribute value for each vertex is a vector with 4 components, each with a normalized value of signed 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/char4Normalized
	MDLVertexFormatChar4Normalized MDLVertexFormat = 0
	// MDLVertexFormatCharBits - A bit mask for vertex attributes whose components are in 8-bit signed integer format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/charBits
	MDLVertexFormatCharBits MDLVertexFormat = 0
	// MDLVertexFormatCharNormalized - The attribute value for each vertex is a normalized scalar of signed 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/charNormalized
	MDLVertexFormatCharNormalized MDLVertexFormat = 0
	// MDLVertexFormatCharNormalizedBits - A bit mask for vertex attributes whose components are in 8-bit signed normalized integer format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/charNormalizedBits
	MDLVertexFormatCharNormalizedBits MDLVertexFormat = 0
	// MDLVertexFormatFloat - The attribute value for each vertex is a scalar of 32-bit floating-point type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/float
	MDLVertexFormatFloat MDLVertexFormat = 0
	// MDLVertexFormatFloat2 - The attribute value for each vertex is a vector with 2 components, each of 32-bit floating-point type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/float2
	MDLVertexFormatFloat2 MDLVertexFormat = 0
	// MDLVertexFormatFloat3 - The attribute value for each vertex is a vector with 3 components, each of 32-bit floating-point type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/float3
	MDLVertexFormatFloat3 MDLVertexFormat = 0
	// MDLVertexFormatFloat4 - The attribute value for each vertex is a vector with 4 components, each of 32-bit floating-point type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/float4
	MDLVertexFormatFloat4 MDLVertexFormat = 0
	// MDLVertexFormatFloatBits - A bit mask for vertex attributes whose components are in 32-bit floating-point format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/floatBits
	MDLVertexFormatFloatBits MDLVertexFormat = 0
	// MDLVertexFormatHalf - The attribute value for each vertex is a scalar of 16-bit floating-point type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/half
	MDLVertexFormatHalf MDLVertexFormat = 0
	// MDLVertexFormatHalf2 - The attribute value for each vertex is a vector with 2 components, each of 16-bit floating-point type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/half2
	MDLVertexFormatHalf2 MDLVertexFormat = 0
	// MDLVertexFormatHalf3 - The attribute value for each vertex is a vector with 3 components, each of 16-bit floating-point type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/half3
	MDLVertexFormatHalf3 MDLVertexFormat = 0
	// MDLVertexFormatHalf4 - The attribute value for each vertex is a vector with 4 components, each of 16-bit floating-point type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/half4
	MDLVertexFormatHalf4 MDLVertexFormat = 0
	// MDLVertexFormatHalfBits - A bit mask for vertex attributes whose components are in 16-bit floating-point format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/halfBits
	MDLVertexFormatHalfBits MDLVertexFormat = 0
	// MDLVertexFormatInt - The attribute value for each vertex is a scalar of signed 32-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/int
	MDLVertexFormatInt MDLVertexFormat = 0
	// MDLVertexFormatInt1010102Normalized - The attribute value for each vertex is a packed vector with 4 components of signed integer type. The first three components are 10 bits each, and the fourth component is 2 bits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/int1010102Normalized
	MDLVertexFormatInt1010102Normalized MDLVertexFormat = 0
	// MDLVertexFormatInt2 - The attribute value for each vertex is a vector with 2 components, each of signed 32-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/int2
	MDLVertexFormatInt2 MDLVertexFormat = 0
	// MDLVertexFormatInt3 - The attribute value for each vertex is a vector with 3 components, each of signed 32-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/int3
	MDLVertexFormatInt3 MDLVertexFormat = 0
	// MDLVertexFormatInt4 - The attribute value for each vertex is a vector with 4 components, each of signed 32-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/int4
	MDLVertexFormatInt4 MDLVertexFormat = 0
	// MDLVertexFormatIntBits - A bit mask for vertex attributes whose components are in 32-bit signed integer format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/intBits
	MDLVertexFormatIntBits MDLVertexFormat = 0
	// MDLVertexFormatInvalid - The vertex attribute has just been initialized or its format is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/invalid
	MDLVertexFormatInvalid MDLVertexFormat = 0
	// MDLVertexFormatPackedBit - A bit mask for vertex attributes in packed vector formats.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/packedBit
	MDLVertexFormatPackedBit MDLVertexFormat = 0
	// MDLVertexFormatShort - The attribute value for each vertex is a scalar of signed 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/short
	MDLVertexFormatShort MDLVertexFormat = 0
	// MDLVertexFormatShort2 - The attribute value for each vertex is a vector with 2 components, each of signed 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/short2
	MDLVertexFormatShort2 MDLVertexFormat = 0
	// MDLVertexFormatShort2Normalized - The attribute value for each vertex is a vector with 2 components, each with a normalized value of signed 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/short2Normalized
	MDLVertexFormatShort2Normalized MDLVertexFormat = 0
	// MDLVertexFormatShort3 - The attribute value for each vertex is a vector with 3 components, each of signed 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/short3
	MDLVertexFormatShort3 MDLVertexFormat = 0
	// MDLVertexFormatShort3Normalized - The attribute value for each vertex is a vector with 3 components, each with a normalized value of signed 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/short3Normalized
	MDLVertexFormatShort3Normalized MDLVertexFormat = 0
	// MDLVertexFormatShort4 - The attribute value for each vertex is a vector with 4 components, each of signed 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/short4
	MDLVertexFormatShort4 MDLVertexFormat = 0
	// MDLVertexFormatShort4Normalized - The attribute value for each vertex is a vector with 4 components, each with a normalized value of signed 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/short4Normalized
	MDLVertexFormatShort4Normalized MDLVertexFormat = 0
	// MDLVertexFormatShortBits - A bit mask for vertex attributes whose components are in 16-bit signed integer format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/shortBits
	MDLVertexFormatShortBits MDLVertexFormat = 0
	// MDLVertexFormatShortNormalized - The attribute value for each vertex is a normalized scalar of signed 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/shortNormalized
	MDLVertexFormatShortNormalized MDLVertexFormat = 0
	// MDLVertexFormatShortNormalizedBits - A bit mask for vertex attributes whose components are in 16-bit signed normalized integer format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/shortNormalizedBits
	MDLVertexFormatShortNormalizedBits MDLVertexFormat = 0
	// MDLVertexFormatUChar - The attribute value for each vertex is a scalar of unsigned 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uChar
	MDLVertexFormatUChar MDLVertexFormat = 0
	// MDLVertexFormatUChar2 - The attribute value for each vertex is a vector with 2 components, each of unsigned 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uChar2
	MDLVertexFormatUChar2 MDLVertexFormat = 0
	// MDLVertexFormatUChar2Normalized - The attribute value for each vertex is a vector with 2 components, each with a normalized value of unsigned 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uChar2Normalized
	MDLVertexFormatUChar2Normalized MDLVertexFormat = 0
	// MDLVertexFormatUChar3 - The attribute value for each vertex is a vector with 3 components, each of unsigned 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uChar3
	MDLVertexFormatUChar3 MDLVertexFormat = 0
	// MDLVertexFormatUChar3Normalized - The attribute value for each vertex is a vector with 3 components, each with a normalized value of unsigned 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uChar3Normalized
	MDLVertexFormatUChar3Normalized MDLVertexFormat = 0
	// MDLVertexFormatUChar4 - The attribute value for each vertex is a vector with 4 components, each of unsigned 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uChar4
	MDLVertexFormatUChar4 MDLVertexFormat = 0
	// MDLVertexFormatUChar4Normalized - The attribute value for each vertex is a vector with 4 components, each with a normalized value of unsigned 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uChar4Normalized
	MDLVertexFormatUChar4Normalized MDLVertexFormat = 0
	// MDLVertexFormatUCharBits - A bit mask for vertex attributes whose components are in 8-bit unsigned integer format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uCharBits
	MDLVertexFormatUCharBits MDLVertexFormat = 0
	// MDLVertexFormatUCharNormalized - The attribute value for each vertex is a normalized scalar of unsigned 8-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uCharNormalized
	MDLVertexFormatUCharNormalized MDLVertexFormat = 0
	// MDLVertexFormatUCharNormalizedBits - A bit mask for vertex attributes whose components are in 8-bit unsigned normalized integer format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uCharNormalizedBits
	MDLVertexFormatUCharNormalizedBits MDLVertexFormat = 0
	// MDLVertexFormatUInt - The attribute value for each vertex is a scalar of unsigned 32-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uInt
	MDLVertexFormatUInt MDLVertexFormat = 0
	// MDLVertexFormatUInt1010102Normalized - The attribute value for each vertex is a packed vector with 4 components of unsigned integer type. The first three components are 10 bits each, and the fourth component is 2 bits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uInt1010102Normalized
	MDLVertexFormatUInt1010102Normalized MDLVertexFormat = 0
	// MDLVertexFormatUInt2 - The attribute value for each vertex is a vector with 2 components, each of unsigned 32-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uInt2
	MDLVertexFormatUInt2 MDLVertexFormat = 0
	// MDLVertexFormatUInt3 - The attribute value for each vertex is a vector with 3 components, each of unsigned 32-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uInt3
	MDLVertexFormatUInt3 MDLVertexFormat = 0
	// MDLVertexFormatUInt4 - The attribute value for each vertex is a vector with 4 components, each of unsigned 32-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uInt4
	MDLVertexFormatUInt4 MDLVertexFormat = 0
	// MDLVertexFormatUIntBits - A bit mask for vertex attributes whose components are in 32-bit unsigned integer format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uIntBits
	MDLVertexFormatUIntBits MDLVertexFormat = 0
	// MDLVertexFormatUShort - The attribute value for each vertex is a scalar of unsigned 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uShort
	MDLVertexFormatUShort MDLVertexFormat = 0
	// MDLVertexFormatUShort2 - The attribute value for each vertex is a vector with 2 components, each of unsigned 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uShort2
	MDLVertexFormatUShort2 MDLVertexFormat = 0
	// MDLVertexFormatUShort2Normalized - The attribute value for each vertex is a vector with 2 components, each with a normalized value of unsigned 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uShort2Normalized
	MDLVertexFormatUShort2Normalized MDLVertexFormat = 0
	// MDLVertexFormatUShort3 - The attribute value for each vertex is a vector with 3 components, each of unsigned 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uShort3
	MDLVertexFormatUShort3 MDLVertexFormat = 0
	// MDLVertexFormatUShort3Normalized - The attribute value for each vertex is a vector with 3 components, each with a normalized value of unsigned 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uShort3Normalized
	MDLVertexFormatUShort3Normalized MDLVertexFormat = 0
	// MDLVertexFormatUShort4 - The attribute value for each vertex is a vector with 4 components, each of unsigned 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uShort4
	MDLVertexFormatUShort4 MDLVertexFormat = 0
	// MDLVertexFormatUShort4Normalized - The attribute value for each vertex is a vector with 4 components, each with a normalized value of unsigned 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uShort4Normalized
	MDLVertexFormatUShort4Normalized MDLVertexFormat = 0
	// MDLVertexFormatUShortBits - A bit mask for vertex attributes whose components are in 16-bit unsigned integer format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uShortBits
	MDLVertexFormatUShortBits MDLVertexFormat = 0
	// MDLVertexFormatUShortNormalized - The attribute value for each vertex is a normalized scalar of unsigned 16-bit integer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uShortNormalized
	MDLVertexFormatUShortNormalized MDLVertexFormat = 0
	// MDLVertexFormatUShortNormalizedBits - A bit mask for vertex attributes whose components are in 16-bit unsigned normalized integer format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat/uShortNormalizedBits
	MDLVertexFormatUShortNormalizedBits MDLVertexFormat = 0
)


