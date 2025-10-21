// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Mesh] class.
var (
	MeshClass     _MeshClass
	MeshClassOnce sync.Once
)

func getMeshClass() _MeshClass {
	MeshClassOnce.Do(func() {
		MeshClass = _MeshClass{objc.GetClass("MTKMesh")}
	})
	return MeshClass
}

type _MeshClass struct {
	class objc.Class
}

// An interface definition for the [Mesh] class.
type IMesh interface {
	objectivec.IObject
}

// A container for the vertex data of a Model I/O mesh, suitable for use in a Metal app.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMesh
type Mesh struct {
	objectivec.Object
}

// MeshFrom constructs a [Mesh] from an unsafe.Pointer.
//
// A container for the vertex data of a Model I/O mesh, suitable for use in a Metal app.
func MeshFrom(ptr unsafe.Pointer) Mesh {
	return Mesh{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MeshClass) Alloc() Mesh {
	rv := objc.Send[Mesh](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MeshClass) New() Mesh {
	rv := objc.Send[Mesh](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Mesh) Init() Mesh {
	rv := objc.Send[Mesh](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Mesh) Autorelease() Mesh {
	rv := objc.Send[Mesh](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMesh creates a new Mesh instance.
func NewMesh() Mesh {
	return getMeshClass().New()
}




// Initializes a MetalKit mesh and its submeshes from a Model I/O mesh.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMesh/init(mesh:device:)
func NewMeshWithMeshDeviceError(mesh unsafe.Pointer, device objectivec.IObject, error_ unsafe.Pointer) Mesh {
	instance := getMeshClass().Alloc()
	rv := objc.Send[Mesh](instance.ID, objc.Sel("initWithMesh:device:error:"), mesh, device, error_)
	rv.Autorelease()
	return rv
}


// Creates and initializes MetalKit meshes from all Model I/O meshes in a Model I/O asset.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMesh/newMeshesFromAsset:device:sourceMeshes:error:
func (mc _MeshClass) NewMeshesFromAssetDeviceSourceMeshesError(asset unsafe.Pointer, device objectivec.IObject, sourceMeshes []MDLMesh, error_ unsafe.Pointer) []Mesh {
	rv := objc.Send[[]Mesh](objc.ID(mc.class), objc.Sel("newMeshesFromAsset:device:sourceMeshes:error:"), asset, device, sourceMeshes, error_)
	return rv
}

// The name of the mesh.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMesh/name
func (m_ Mesh) Name() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the mesh.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMesh/name
func (m_ Mesh) SetName(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}

// An array of submeshes containing index buffers referencing the mesh vertices.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMesh/submeshes
func (m_ Mesh) Submeshes() []Submesh {
	rv := objc.Send[[]Submesh](m_.ID, objc.Sel("submeshes"))
	return rv
}

// An array of buffers in which mesh vertex data resides.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMesh/vertexBuffers
func (m_ Mesh) VertexBuffers() []MeshBuffer {
	rv := objc.Send[[]MeshBuffer](m_.ID, objc.Sel("vertexBuffers"))
	return rv
}

// The number of vertices in the vertex buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMesh/vertexCount
func (m_ Mesh) VertexCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("vertexCount"))
	return rv
}

// A Model I/O vertex descriptor specifying the data layout in the vertex buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMesh/vertexDescriptor
func (m_ Mesh) VertexDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("vertexDescriptor"))
	return rv
}


