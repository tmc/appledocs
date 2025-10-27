// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Submesh] class.
var (
	SubmeshClass     _SubmeshClass
	SubmeshClassOnce sync.Once
)

func getSubmeshClass() _SubmeshClass {
	SubmeshClassOnce.Do(func() {
		SubmeshClass = _SubmeshClass{objc.GetClass("MTKSubmesh")}
	})
	return SubmeshClass
}

type _SubmeshClass struct {
	class objc.Class
}





// An interface definition for the [Submesh] class.
type ISubmesh interface {
	objectivec.IObject
	

	// properties:
	IndexBuffer() IMTKMeshBuffer
	IndexCount() uint
	IndexType() IndexType /* not a class type */
	Mesh() IMTKMesh
	Name() foundation.foundation.INSString
	SetName(value foundation.foundation.INSString)
	PrimitiveType() PrimitiveType /* not a class type */


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _SubmeshClass) Alloc() Submesh {
	rv := objc.Send[Submesh](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SubmeshClass) New() Submesh {
	rv := objc.Send[Submesh](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Submesh) Init() Submesh {
	rv := objc.Send[Submesh](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Submesh) Autorelease() Submesh {
	rv := objc.Send[Submesh](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSubmesh creates a new Submesh instance.
func NewSubmesh() Submesh {
	return getSubmeshClass().New()
}





// A container for the index data of a Model I/O submesh, suitable for use in a Metal app.
//
// The class provides a container for a segment of mesh data that can be rendered in a single draw call. A submesh can only be initialized as part of a object. Each submesh contains an index buffer with which the parent’s mesh data can be rendered. Actual submesh vertex data resides in the submesh’s parent mesh. For more information on Model I/O submeshes, see .


// A container for the index data of a Model I/O submesh, suitable for use in a Metal app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKSubmesh
type Submesh struct {
	objectivec.Object
}

// SubmeshFrom constructs a [Submesh] from an unsafe.Pointer.
//
// A container for the index data of a Model I/O submesh, suitable for use in a Metal app.
func SubmeshFrom(ptr unsafe.Pointer) Submesh {
	return Submesh{objectivec.Object{objc.ID(ptr)}}
}

























// The index buffer used to render the submesh object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKSubmesh/indexBuffer
func (s_ Submesh) IndexBuffer() IMTKMeshBuffer {
	rv := objc.Send[MeshBuffer](s_.ID, objc.Sel("indexBuffer"))
	return rv
}


// The number of indices in the index buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKSubmesh/indexCount
func (s_ Submesh) IndexCount() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("indexCount"))
	return rv
}


// The type of index data in the index buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKSubmesh/indexType
func (s_ Submesh) IndexType() IndexType /* not a class type */ {
	rv := objc.Send[IndexType](s_.ID, objc.Sel("indexType"))
	return rv
}


// The parent mesh containing the vertex data of this submesh.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKSubmesh/mesh
func (s_ Submesh) Mesh() IMTKMesh {
	rv := objc.Send[Mesh](s_.ID, objc.Sel("mesh"))
	return rv
}


// The name of the submesh.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKSubmesh/name
func (s_ Submesh) Name() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("name"))
	return rv
}


// The name of the submesh.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKSubmesh/name
func (s_ Submesh) SetName(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setName:"), value)
}


// The primitive type with which to draw the submesh object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKSubmesh/primitiveType
func (s_ Submesh) PrimitiveType() PrimitiveType /* not a class type */ {
	rv := objc.Send[PrimitiveType](s_.ID, objc.Sel("primitiveType"))
	return rv
}








