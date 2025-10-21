// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GLKMesh] class.
var (
	GLKMeshClass     _GLKMeshClass
	GLKMeshClassOnce sync.Once
)

func getGLKMeshClass() _GLKMeshClass {
	GLKMeshClassOnce.Do(func() {
		GLKMeshClass = _GLKMeshClass{objc.GetClass("GLKMesh")}
	})
	return GLKMeshClass
}

type _GLKMeshClass struct {
	class objc.Class
}

// An interface definition for the [GLKMesh] class.
type IGLKMesh interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh
type GLKMesh struct {
	objectivec.Object
}

// GLKMeshFrom constructs a [GLKMesh] from an unsafe.Pointer.
func GLKMeshFrom(ptr unsafe.Pointer) GLKMesh {
	return GLKMesh{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GLKMeshClass) Alloc() GLKMesh {
	rv := objc.Send[GLKMesh](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GLKMeshClass) New() GLKMesh {
	rv := objc.Send[GLKMesh](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKMesh) Init() GLKMesh {
	rv := objc.Send[GLKMesh](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKMesh) Autorelease() GLKMesh {
	rv := objc.Send[GLKMesh](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKMesh creates a new GLKMesh instance.
func NewGLKMesh() GLKMesh {
	return getGLKMeshClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/init(mesh:)
func NewGLKMeshWithMeshError(mesh unsafe.Pointer, error_ unsafe.Pointer) GLKMesh {
	instance := getGLKMeshClass().Alloc()
	rv := objc.Send[GLKMesh](instance.ID, objc.Sel("initWithMesh:error:"), mesh, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/newMeshes(from:sourceMeshes:)
func (gc _GLKMeshClass) NewMeshesFromAssetSourceMeshesError(asset unsafe.Pointer, sourceMeshes []MDLMesh, error_ unsafe.Pointer) []GLKMesh {
	rv := objc.Send[[]GLKMesh](objc.ID(gc.class), objc.Sel("newMeshesFromAsset:sourceMeshes:error:"), asset, sourceMeshes, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/name
func (g_ GLKMesh) Name() appkit.string {
	rv := objc.Send[appkit.string](g_.ID, objc.Sel("name"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/submeshes
func (g_ GLKMesh) Submeshes() []GLKSubmesh {
	rv := objc.Send[[]GLKSubmesh](g_.ID, objc.Sel("submeshes"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/vertexBuffers
func (g_ GLKMesh) VertexBuffers() []GLKMeshBuffer {
	rv := objc.Send[[]GLKMeshBuffer](g_.ID, objc.Sel("vertexBuffers"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/vertexCount
func (g_ GLKMesh) VertexCount() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("vertexCount"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/vertexDescriptor
func (g_ GLKMesh) VertexDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("vertexDescriptor"))
	return rv
}


