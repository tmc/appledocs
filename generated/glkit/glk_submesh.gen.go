// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GLKSubmesh */


/* debug [class_header]: Header for GLKSubmesh */
// The class instance for the [GLKSubmesh] class.
var (
	GLKSubmeshClass     _GLKSubmeshClass
	GLKSubmeshClassOnce sync.Once
)

func getGLKSubmeshClass() _GLKSubmeshClass {
	GLKSubmeshClassOnce.Do(func() {
		GLKSubmeshClass = _GLKSubmeshClass{objc.GetClass("GLKSubmesh")}
	})
	return GLKSubmeshClass
}

type _GLKSubmeshClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GLKSubmesh */
// An interface definition for the [GLKSubmesh] class.
type IGLKSubmesh interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GLKSubmesh */
	// properties:
	ElementBuffer() IGLKMeshBuffer
	ElementCount() unsafe.Pointer
	Mesh() IGLKMesh
	Mode() unsafe.Pointer
	Name() objc.IObject /* cross-framework: NSString */
	Type() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GLKSubmesh */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GLKSubmesh */
// Alloc allocates a new instance without initialization.
func (gc _GLKSubmeshClass) Alloc() GLKSubmesh {
	rv := objc.Send[GLKSubmesh](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GLKSubmeshClass) New() GLKSubmesh {
	rv := objc.Send[GLKSubmesh](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKSubmesh) Init() GLKSubmesh {
	rv := objc.Send[GLKSubmesh](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKSubmesh) Autorelease() GLKSubmesh {
	rv := objc.Send[GLKSubmesh](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKSubmesh creates a new GLKSubmesh instance.
func NewGLKSubmesh() GLKSubmesh {
	return getGLKSubmeshClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GLKSubmesh */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSubmesh
type GLKSubmesh struct {
	objectivec.Object
}

// GLKSubmeshFrom constructs a [GLKSubmesh] from an unsafe.Pointer.
func GLKSubmeshFrom(ptr unsafe.Pointer) GLKSubmesh {
	return GLKSubmesh{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GLKSubmesh *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GLKSubmesh */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GLKSubmesh */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GLKSubmesh */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GLKSubmesh */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSubmesh/elementBuffer
func (g_ GLKSubmesh) ElementBuffer() IGLKMeshBuffer {
	rv := objc.Send[GLKMeshBuffer](g_.ID, objc.Sel("elementBuffer"))
	return rv
}/* debug [instance_properties/getter]: elementBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSubmesh/elementCount
func (g_ GLKSubmesh) ElementCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("elementCount"))
	return rv
}/* debug [instance_properties/getter]: elementCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSubmesh/mesh
func (g_ GLKSubmesh) Mesh() IGLKMesh {
	rv := objc.Send[GLKMesh](g_.ID, objc.Sel("mesh"))
	return rv
}/* debug [instance_properties/getter]: mesh */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSubmesh/mode
func (g_ GLKSubmesh) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSubmesh/name
func (g_ GLKSubmesh) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSubmesh/type
func (g_ GLKSubmesh) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GLKSubmesh */



