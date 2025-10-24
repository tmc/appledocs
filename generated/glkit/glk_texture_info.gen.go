// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GLKTextureInfo */


/* debug [class_header]: Header for GLKTextureInfo */
// The class instance for the [GLKTextureInfo] class.
var (
	GLKTextureInfoClass     _GLKTextureInfoClass
	GLKTextureInfoClassOnce sync.Once
)

func getGLKTextureInfoClass() _GLKTextureInfoClass {
	GLKTextureInfoClassOnce.Do(func() {
		GLKTextureInfoClass = _GLKTextureInfoClass{objc.GetClass("GLKTextureInfo")}
	})
	return GLKTextureInfoClass
}

type _GLKTextureInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GLKTextureInfo */
// An interface definition for the [GLKTextureInfo] class.
type IGLKTextureInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GLKTextureInfo */
	// properties:
	AlphaState() GLKTextureInfoAlphaState
	ArrayLength() unsafe.Pointer
	ContainsMipmaps() bool
	Depth() unsafe.Pointer
	Height() unsafe.Pointer
	MimapLevelCount() unsafe.Pointer
	Name() unsafe.Pointer
	Target() unsafe.Pointer
	TextureOrigin() GLKTextureInfoOrigin
	Width() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GLKTextureInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GLKTextureInfo */
// Alloc allocates a new instance without initialization.
func (gc _GLKTextureInfoClass) Alloc() GLKTextureInfo {
	rv := objc.Send[GLKTextureInfo](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GLKTextureInfoClass) New() GLKTextureInfo {
	rv := objc.Send[GLKTextureInfo](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKTextureInfo) Init() GLKTextureInfo {
	rv := objc.Send[GLKTextureInfo](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKTextureInfo) Autorelease() GLKTextureInfo {
	rv := objc.Send[GLKTextureInfo](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKTextureInfo creates a new GLKTextureInfo instance.
func NewGLKTextureInfo() GLKTextureInfo {
	return getGLKTextureInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GLKTextureInfo */
// Information about OpenGL textures created by the class.
//
// When your app loads textures using the class, the texture loader returns information about the textures using objects. Your app never creates objects directly.


// Information about OpenGL textures created by the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo
type GLKTextureInfo struct {
	objectivec.Object
}

// GLKTextureInfoFrom constructs a [GLKTextureInfo] from an unsafe.Pointer.
//
// Information about OpenGL textures created by the class.
func GLKTextureInfoFrom(ptr unsafe.Pointer) GLKTextureInfo {
	return GLKTextureInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GLKTextureInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GLKTextureInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GLKTextureInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GLKTextureInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GLKTextureInfo */

// The state of the alpha component in the loaded texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/alphaState-swift.property
func (g_ GLKTextureInfo) AlphaState() GLKTextureInfoAlphaState {
	rv := objc.Send[GLKTextureInfoAlphaState](g_.ID, objc.Sel("alphaState"))
	return rv
}/* debug [instance_properties/getter]: alphaState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/arrayLength-swift.property
func (g_ GLKTextureInfo) ArrayLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("arrayLength"))
	return rv
}/* debug [instance_properties/getter]: arrayLength */


// A Boolean value that states whether the loaded texture contains mip maps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/containsMipmaps-swift.property
func (g_ GLKTextureInfo) ContainsMipmaps() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("containsMipmaps"))
	return rv
}/* debug [instance_properties/getter]: containsMipmaps */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/depth-swift.property
func (g_ GLKTextureInfo) Depth() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("depth"))
	return rv
}/* debug [instance_properties/getter]: depth */


// The height of the loaded texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/height-swift.property
func (g_ GLKTextureInfo) Height() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/mimapLevelCount-swift.property
func (g_ GLKTextureInfo) MimapLevelCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("mimapLevelCount"))
	return rv
}/* debug [instance_properties/getter]: mimapLevelCount */


// The OpenGL context’s name for the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/name-swift.property
func (g_ GLKTextureInfo) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The OpenGL binding target for the texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/target-swift.property
func (g_ GLKTextureInfo) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// The location of the origin in the loaded texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/textureOrigin-swift.property
func (g_ GLKTextureInfo) TextureOrigin() GLKTextureInfoOrigin {
	rv := objc.Send[GLKTextureInfoOrigin](g_.ID, objc.Sel("textureOrigin"))
	return rv
}/* debug [instance_properties/getter]: textureOrigin */


// The width of the loaded texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/width-swift.property
func (g_ GLKTextureInfo) Width() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GLKTextureInfo */



