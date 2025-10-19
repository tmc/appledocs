// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTKMesh] class.
var (
	mTKMeshClass     _MTKMeshClass
	mTKMeshClassOnce sync.Once
)

func getMTKMeshClass() _MTKMeshClass {
	mTKMeshClassOnce.Do(func() {
		mTKMeshClass = _MTKMeshClass{objc.GetClass("MTKMesh")}
	})
	return mTKMeshClass
}

type _MTKMeshClass struct {
	class objc.Class
}

// An interface definition for the [MTKMesh] class.
type IMTKMesh interface {
	objectivec.IObject
}

// A container for the vertex data of a Model I/O mesh, suitable for use in a Metal app.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMesh
type MTKMesh struct {
	objectivec.Object
}

// MTKMeshFrom constructs a [MTKMesh] from an unsafe.Pointer.
//
// A container for the vertex data of a Model I/O mesh, suitable for use in a Metal app.
func MTKMeshFrom(ptr unsafe.Pointer) MTKMesh {
	return MTKMesh{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTKMeshClass) Alloc() MTKMesh {
	rv := objc.Send[MTKMesh](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTKMeshClass) New() MTKMesh {
	rv := objc.Send[MTKMesh](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTKMesh) Init() MTKMesh {
	rv := objc.Send[MTKMesh](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTKMesh) Autorelease() MTKMesh {
	rv := objc.Send[MTKMesh](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTKMesh creates a new MTKMesh instance.
func NewMTKMesh() MTKMesh {
	return getMTKMeshClass().New()
}


// Initializes a MetalKit mesh and its submeshes from a Model I/O mesh.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMesh/init(mesh:device:)
func NewMTKMeshWithMeshDeviceError(mesh unsafe.Pointer, device unsafe.Pointer, error unsafe.Pointer) MTKMesh {
	instance := getMTKMeshClass().Alloc()
	rv := objc.Send[MTKMesh](instance.ID, objc.Sel("initWithMesh:device:error:"), mesh, device, error)
	rv.Autorelease()
	return rv
}


// Creates and initializes MetalKit meshes from all Model I/O meshes in a Model I/O asset.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKMesh/newMeshesFromAsset:device:sourceMeshes:error:
func (mc _MTKMeshClass) NewMeshesFromAssetDeviceSourceMeshesError(asset unsafe.Pointer, device unsafe.Pointer, sourceMeshes unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("newMeshesFromAsset:device:sourceMeshes:error:"), asset, device, sourceMeshes, error)
	return rv
}

