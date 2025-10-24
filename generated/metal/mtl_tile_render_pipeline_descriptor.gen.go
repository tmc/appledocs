// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLTileRenderPipelineDescriptor */


/* debug [class_header]: Header for MTLTileRenderPipelineDescriptor */
// The class instance for the [TileRenderPipelineDescriptor] class.
var (
	TileRenderPipelineDescriptorClass     _TileRenderPipelineDescriptorClass
	TileRenderPipelineDescriptorClassOnce sync.Once
)

func getTileRenderPipelineDescriptorClass() _TileRenderPipelineDescriptorClass {
	TileRenderPipelineDescriptorClassOnce.Do(func() {
		TileRenderPipelineDescriptorClass = _TileRenderPipelineDescriptorClass{objc.GetClass("MTLTileRenderPipelineDescriptor")}
	})
	return TileRenderPipelineDescriptorClass
}

type _TileRenderPipelineDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TileRenderPipelineDescriptor */
// An interface definition for the [TileRenderPipelineDescriptor] class.
type ITileRenderPipelineDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TileRenderPipelineDescriptor */
	// properties:
	BinaryArchives() []objc.ID
	SetBinaryArchives(value []objc.ID)
	ColorAttachments() IMTLTileRenderPipelineColorAttachmentDescriptorArray
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	LinkedFunctions() IMTLLinkedFunctions
	SetLinkedFunctions(value IMTLLinkedFunctions)
	MaxCallStackDepth() uint
	SetMaxCallStackDepth(value uint)
	MaxTotalThreadsPerThreadgroup() uint
	SetMaxTotalThreadsPerThreadgroup(value uint)
	PreloadedLibraries() []objc.ID
	SetPreloadedLibraries(value []objc.ID)
	RasterSampleCount() uint
	SetRasterSampleCount(value uint)
	RequiredThreadsPerThreadgroup() objc.IObject /* cross-framework: MTLSize */
	SetRequiredThreadsPerThreadgroup(value objc.IObject /* cross-framework: MTLSize */)
	ShaderValidation() ShaderValidation
	SetShaderValidation(value ShaderValidation)
	SupportAddingBinaryFunctions() bool
	SetSupportAddingBinaryFunctions(value bool)
	ThreadgroupSizeMatchesTileSize() bool
	SetThreadgroupSizeMatchesTileSize(value bool)
	TileBuffers() IMTLPipelineBufferDescriptorArray
	TileFunction() unsafe.Pointer
	SetTileFunction(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TileRenderPipelineDescriptor */
	// methods:
	Reset()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TileRenderPipelineDescriptor */
// Alloc allocates a new instance without initialization.
func (tc _TileRenderPipelineDescriptorClass) Alloc() TileRenderPipelineDescriptor {
	rv := objc.Send[TileRenderPipelineDescriptor](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TileRenderPipelineDescriptorClass) New() TileRenderPipelineDescriptor {
	rv := objc.Send[TileRenderPipelineDescriptor](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TileRenderPipelineDescriptor) Init() TileRenderPipelineDescriptor {
	rv := objc.Send[TileRenderPipelineDescriptor](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TileRenderPipelineDescriptor) Autorelease() TileRenderPipelineDescriptor {
	rv := objc.Send[TileRenderPipelineDescriptor](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTileRenderPipelineDescriptor creates a new TileRenderPipelineDescriptor instance.
func NewTileRenderPipelineDescriptor() TileRenderPipelineDescriptor {
	return getTileRenderPipelineDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TileRenderPipelineDescriptor */
// An object that configures new render pipeline state objects for tile shading.


// An object that configures new render pipeline state objects for tile shading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor
type TileRenderPipelineDescriptor struct {
	objectivec.Object
}

// TileRenderPipelineDescriptorFrom constructs a [TileRenderPipelineDescriptor] from an unsafe.Pointer.
//
// An object that configures new render pipeline state objects for tile shading.
func TileRenderPipelineDescriptorFrom(ptr unsafe.Pointer) TileRenderPipelineDescriptor {
	return TileRenderPipelineDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TileRenderPipelineDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TileRenderPipelineDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TileRenderPipelineDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TileRenderPipelineDescriptor */

// Specifies the default rendering pipeline state values for the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/reset()
func (t_ TileRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](t_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TileRenderPipelineDescriptor */

// An array of binary archives to search for precompiled versions of the shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/binaryArchives
func (t_ TileRenderPipelineDescriptor) BinaryArchives() []objc.ID {
	rv := objc.Send[[]objc.ID](t_.ID, objc.Sel("binaryArchives"))
	return rv
}/* debug [instance_properties/getter]: binaryArchives */


// An array of binary archives to search for precompiled versions of the shader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/binaryArchives
func (t_ TileRenderPipelineDescriptor) SetBinaryArchives(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setBinaryArchives:"), nsArray)
}/* debug [instance_properties/setter]: binaryArchives */


// An array of attachments that store color data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/colorAttachments
func (t_ TileRenderPipelineDescriptor) ColorAttachments() IMTLTileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[TileRenderPipelineColorAttachmentDescriptorArray](t_.ID, objc.Sel("colorAttachments"))
	return rv
}/* debug [instance_properties/getter]: colorAttachments */


// A string that identifies the tile pipeline descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/label
func (t_ TileRenderPipelineDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// A string that identifies the tile pipeline descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/label
func (t_ TileRenderPipelineDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// Functions that you can specify as function arguments for the tile shader when encoding commands that use the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/linkedFunctions
func (t_ TileRenderPipelineDescriptor) LinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[LinkedFunctions](t_.ID, objc.Sel("linkedFunctions"))
	return rv
}/* debug [instance_properties/getter]: linkedFunctions */


// Functions that you can specify as function arguments for the tile shader when encoding commands that use the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/linkedFunctions
func (t_ TileRenderPipelineDescriptor) SetLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLinkedFunctions:"), value)
}/* debug [instance_properties/setter]: linkedFunctions */


// The maximum function call depth from the top-most shader function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/maxCallStackDepth
func (t_ TileRenderPipelineDescriptor) MaxCallStackDepth() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("maxCallStackDepth"))
	return rv
}/* debug [instance_properties/getter]: maxCallStackDepth */


// The maximum function call depth from the top-most shader function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/maxCallStackDepth
func (t_ TileRenderPipelineDescriptor) SetMaxCallStackDepth(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxCallStackDepth:"), value)
}/* debug [instance_properties/setter]: maxCallStackDepth */


// The maximum number of threads in a threadgroup when dispatching a command using the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/maxTotalThreadsPerThreadgroup
func (t_ TileRenderPipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}/* debug [instance_properties/getter]: maxTotalThreadsPerThreadgroup */


// The maximum number of threads in a threadgroup when dispatching a command using the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/maxTotalThreadsPerThreadgroup
func (t_ TileRenderPipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}/* debug [instance_properties/setter]: maxTotalThreadsPerThreadgroup */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/preloadedLibraries
func (t_ TileRenderPipelineDescriptor) PreloadedLibraries() []objc.ID {
	rv := objc.Send[[]objc.ID](t_.ID, objc.Sel("preloadedLibraries"))
	return rv
}/* debug [instance_properties/getter]: preloadedLibraries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/preloadedLibraries
func (t_ TileRenderPipelineDescriptor) SetPreloadedLibraries(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreloadedLibraries:"), nsArray)
}/* debug [instance_properties/setter]: preloadedLibraries */


// The number of samples in each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/rasterSampleCount
func (t_ TileRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("rasterSampleCount"))
	return rv
}/* debug [instance_properties/getter]: rasterSampleCount */


// The number of samples in each fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/rasterSampleCount
func (t_ TileRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRasterSampleCount:"), value)
}/* debug [instance_properties/setter]: rasterSampleCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/requiredThreadsPerThreadgroup
func (t_ TileRenderPipelineDescriptor) RequiredThreadsPerThreadgroup() objc.IObject /* cross-framework: MTLSize */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return rv
}/* debug [instance_properties/getter]: requiredThreadsPerThreadgroup */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/requiredThreadsPerThreadgroup
func (t_ TileRenderPipelineDescriptor) SetRequiredThreadsPerThreadgroup(value objc.IObject /* cross-framework: MTLSize */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}/* debug [instance_properties/setter]: requiredThreadsPerThreadgroup */


// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/shaderValidation
func (t_ TileRenderPipelineDescriptor) ShaderValidation() ShaderValidation {
	rv := objc.Send[ShaderValidation](t_.ID, objc.Sel("shaderValidation"))
	return rv
}/* debug [instance_properties/getter]: shaderValidation */


// A value that enables or disables shader validation for the pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/shaderValidation
func (t_ TileRenderPipelineDescriptor) SetShaderValidation(value ShaderValidation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setShaderValidation:"), value)
}/* debug [instance_properties/setter]: shaderValidation */


// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/supportAddingBinaryFunctions
func (t_ TileRenderPipelineDescriptor) SupportAddingBinaryFunctions() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("supportAddingBinaryFunctions"))
	return rv
}/* debug [instance_properties/getter]: supportAddingBinaryFunctions */


// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/supportAddingBinaryFunctions
func (t_ TileRenderPipelineDescriptor) SetSupportAddingBinaryFunctions(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSupportAddingBinaryFunctions:"), value)
}/* debug [instance_properties/setter]: supportAddingBinaryFunctions */


// A Boolean value that indicates whether all threadgroups for this pipeline completely cover tiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/threadgroupSizeMatchesTileSize
func (t_ TileRenderPipelineDescriptor) ThreadgroupSizeMatchesTileSize() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("threadgroupSizeMatchesTileSize"))
	return rv
}/* debug [instance_properties/getter]: threadgroupSizeMatchesTileSize */


// A Boolean value that indicates whether all threadgroups for this pipeline completely cover tiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/threadgroupSizeMatchesTileSize
func (t_ TileRenderPipelineDescriptor) SetThreadgroupSizeMatchesTileSize(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setThreadgroupSizeMatchesTileSize:"), value)
}/* debug [instance_properties/setter]: threadgroupSizeMatchesTileSize */


// An array that contains the buffer mutability options for a render pipeline’s tile function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/tileBuffers
func (t_ TileRenderPipelineDescriptor) TileBuffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[PipelineBufferDescriptorArray](t_.ID, objc.Sel("tileBuffers"))
	return rv
}/* debug [instance_properties/getter]: tileBuffers */


// The compute kernel or fragment function the pipeline calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/tileFunction
func (t_ TileRenderPipelineDescriptor) TileFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tileFunction"))
	return rv
}/* debug [instance_properties/getter]: tileFunction */


// The compute kernel or fragment function the pipeline calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/tileFunction
func (t_ TileRenderPipelineDescriptor) SetTileFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTileFunction:"), value)
}/* debug [instance_properties/setter]: tileFunction */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLTileRenderPipelineDescriptor */



