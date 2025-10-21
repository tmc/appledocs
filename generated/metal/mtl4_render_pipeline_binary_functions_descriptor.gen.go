// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTL4RenderPipelineBinaryFunctionsDescriptor] class.
var (
	MTL4RenderPipelineBinaryFunctionsDescriptorClass     _MTL4RenderPipelineBinaryFunctionsDescriptorClass
	MTL4RenderPipelineBinaryFunctionsDescriptorClassOnce sync.Once
)

func getMTL4RenderPipelineBinaryFunctionsDescriptorClass() _MTL4RenderPipelineBinaryFunctionsDescriptorClass {
	MTL4RenderPipelineBinaryFunctionsDescriptorClassOnce.Do(func() {
		MTL4RenderPipelineBinaryFunctionsDescriptorClass = _MTL4RenderPipelineBinaryFunctionsDescriptorClass{objc.GetClass("MTL4RenderPipelineBinaryFunctionsDescriptor")}
	})
	return MTL4RenderPipelineBinaryFunctionsDescriptorClass
}

type _MTL4RenderPipelineBinaryFunctionsDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4RenderPipelineBinaryFunctionsDescriptor] class.
type IMTL4RenderPipelineBinaryFunctionsDescriptor interface {
	objectivec.IObject
	Reset()
}

// Allows you to specify additional binary functions to link to each stage of a render pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor
type MTL4RenderPipelineBinaryFunctionsDescriptor struct {
	objectivec.Object
}

// MTL4RenderPipelineBinaryFunctionsDescriptorFrom constructs a [MTL4RenderPipelineBinaryFunctionsDescriptor] from an unsafe.Pointer.
//
// Allows you to specify additional binary functions to link to each stage of a render pipeline.
func MTL4RenderPipelineBinaryFunctionsDescriptorFrom(ptr unsafe.Pointer) MTL4RenderPipelineBinaryFunctionsDescriptor {
	return MTL4RenderPipelineBinaryFunctionsDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4RenderPipelineBinaryFunctionsDescriptorClass) Alloc() MTL4RenderPipelineBinaryFunctionsDescriptor {
	rv := objc.Send[MTL4RenderPipelineBinaryFunctionsDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4RenderPipelineBinaryFunctionsDescriptorClass) New() MTL4RenderPipelineBinaryFunctionsDescriptor {
	rv := objc.Send[MTL4RenderPipelineBinaryFunctionsDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) Init() MTL4RenderPipelineBinaryFunctionsDescriptor {
	rv := objc.Send[MTL4RenderPipelineBinaryFunctionsDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) Autorelease() MTL4RenderPipelineBinaryFunctionsDescriptor {
	rv := objc.Send[MTL4RenderPipelineBinaryFunctionsDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPipelineBinaryFunctionsDescriptor creates a new MTL4RenderPipelineBinaryFunctionsDescriptor instance.
func NewMTL4RenderPipelineBinaryFunctionsDescriptor() MTL4RenderPipelineBinaryFunctionsDescriptor {
	return getMTL4RenderPipelineBinaryFunctionsDescriptorClass().New()
}


// Resets this descriptor to its default state.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/reset()
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) Reset() {
	objc.Send[objc.ID](m_.ID, objc.Sel("reset"))
}

// Provides an array of binary functions representing additional binary fragment shader functions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/fragmentAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) FragmentAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("fragmentAdditionalBinaryFunctions"))
	return rv
}


// SetFragmentAdditionalBinaryFunctions sets the value of the fragmentAdditionalBinaryFunctions property.
// Provides an array of binary functions representing additional binary fragment shader functions.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/fragmentAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) SetFragmentAdditionalBinaryFunctions(value []objc.ID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setFragmentAdditionalBinaryFunctions:"), nsArray)
}

// Provides an array of binary functions representing additional binary mesh shader functions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/meshAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) MeshAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("meshAdditionalBinaryFunctions"))
	return rv
}


// SetMeshAdditionalBinaryFunctions sets the value of the meshAdditionalBinaryFunctions property.
// Provides an array of binary functions representing additional binary mesh shader functions.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/meshAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) SetMeshAdditionalBinaryFunctions(value []objc.ID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeshAdditionalBinaryFunctions:"), nsArray)
}

// Provides an array of binary functions representing additional binary object shader functions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/objectAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) ObjectAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("objectAdditionalBinaryFunctions"))
	return rv
}


// SetObjectAdditionalBinaryFunctions sets the value of the objectAdditionalBinaryFunctions property.
// Provides an array of binary functions representing additional binary object shader functions.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/objectAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) SetObjectAdditionalBinaryFunctions(value []objc.ID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectAdditionalBinaryFunctions:"), nsArray)
}

// Provides an array of binary functions representing additional binary tile shader functions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/tileAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) TileAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("tileAdditionalBinaryFunctions"))
	return rv
}


// SetTileAdditionalBinaryFunctions sets the value of the tileAdditionalBinaryFunctions property.
// Provides an array of binary functions representing additional binary tile shader functions.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/tileAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) SetTileAdditionalBinaryFunctions(value []objc.ID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setTileAdditionalBinaryFunctions:"), nsArray)
}

// Provides an array of binary functions representing additional binary vertex shader functions.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/vertexAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) VertexAdditionalBinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("vertexAdditionalBinaryFunctions"))
	return rv
}


// SetVertexAdditionalBinaryFunctions sets the value of the vertexAdditionalBinaryFunctions property.
// Provides an array of binary functions representing additional binary vertex shader functions.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineBinaryFunctionsDescriptor/vertexAdditionalBinaryFunctions
func (m_ MTL4RenderPipelineBinaryFunctionsDescriptor) SetVertexAdditionalBinaryFunctions(value []objc.ID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setVertexAdditionalBinaryFunctions:"), nsArray)
}



