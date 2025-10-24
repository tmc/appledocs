// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GLKMesh */


/* debug [class_header]: Header for GLKMesh */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GLKMesh */
// An interface definition for the [GLKMesh] class.
type IGLKMesh interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GLKMesh */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	Submeshes() []GLKSubmesh
	VertexBuffers() []GLKMeshBuffer
	VertexCount() uint
	VertexDescriptor() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GLKMesh */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GLKMesh */
// Alloc allocates a new instance without initialization.
func (gc _GLKMeshClass) Alloc() GLKMesh {
	rv := objc.Send[GLKMesh](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GLKMesh */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh
type GLKMesh struct {
	objectivec.Object
}

// GLKMeshFrom constructs a [GLKMesh] from an unsafe.Pointer.
func GLKMeshFrom(ptr unsafe.Pointer) GLKMesh {
	return GLKMesh{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GLKMesh */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/init(mesh:)
func NewGLKMeshWithMeshError(mesh unsafe.Pointer, error_ unsafe.Pointer) GLKMesh {
	instance := getGLKMeshClass().Alloc()
	rv := objc.Send[GLKMesh](instance.ID, objc.Sel("initWithMesh:error:"), mesh, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGLKMeshWithMeshError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GLKMesh */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/newMeshes(from:sourceMeshes:)
func (gc _GLKMeshClass) NewMeshesFromAssetSourceMeshesError(asset unsafe.Pointer, sourceMeshes []MDLMesh /* not a class type */, error_ unsafe.Pointer) []GLKMesh {
	rv := objc.Send[[]GLKMesh](objc.ID(gc.class), objc.Sel("newMeshesFromAsset:sourceMeshes:error:"), asset, sourceMeshes, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NewMeshesFromAssetSourceMeshesError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GLKMesh */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GLKMesh */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GLKMesh */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/name
func (g_ GLKMesh) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/submeshes
func (g_ GLKMesh) Submeshes() []GLKSubmesh {
	rv := objc.Send[[]GLKSubmesh](g_.ID, objc.Sel("submeshes"))
	return rv
}/* debug [instance_properties/getter]: submeshes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/vertexBuffers
func (g_ GLKMesh) VertexBuffers() []GLKMeshBuffer {
	rv := objc.Send[[]GLKMeshBuffer](g_.ID, objc.Sel("vertexBuffers"))
	return rv
}/* debug [instance_properties/getter]: vertexBuffers */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/vertexCount
func (g_ GLKMesh) VertexCount() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("vertexCount"))
	return rv
}/* debug [instance_properties/getter]: vertexCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMesh/vertexDescriptor
func (g_ GLKMesh) VertexDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("vertexDescriptor"))
	return rv
}/* debug [instance_properties/getter]: vertexDescriptor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GLKMesh */


