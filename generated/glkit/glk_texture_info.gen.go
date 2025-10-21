// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [GLKTextureInfo] class.
type IGLKTextureInfo interface {
	objectivec.IObject
}

// Information about OpenGL textures created by the class.
//
// When your app loads textures using the class, the texture loader returns information about the textures using objects. Your app never creates objects directly.
//
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

// Alloc allocates a new instance without initialization.
func (gc _GLKTextureInfoClass) Alloc() GLKTextureInfo {
	rv := objc.Send[GLKTextureInfo](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The state of the alpha component in the loaded texture.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/alphaState-swift.property
func (g_ GLKTextureInfo) AlphaState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("alphaState"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/arrayLength-swift.property
func (g_ GLKTextureInfo) ArrayLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("arrayLength"))
	return rv
}

// A Boolean value that states whether the loaded texture contains mip maps.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/containsMipmaps-swift.property
func (g_ GLKTextureInfo) ContainsMipmaps() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("containsMipmaps"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/depth-swift.property
func (g_ GLKTextureInfo) Depth() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("depth"))
	return rv
}

// The height of the loaded texture.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/height-swift.property
func (g_ GLKTextureInfo) Height() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("height"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/mimapLevelCount-swift.property
func (g_ GLKTextureInfo) MimapLevelCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("mimapLevelCount"))
	return rv
}

// The OpenGL context’s name for the texture.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/name-swift.property
func (g_ GLKTextureInfo) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("name"))
	return rv
}

// The OpenGL binding target for the texture.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/target-swift.property
func (g_ GLKTextureInfo) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("target"))
	return rv
}

// The location of the origin in the loaded texture.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/textureOrigin-swift.property
func (g_ GLKTextureInfo) TextureOrigin() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("textureOrigin"))
	return rv
}

// The width of the loaded texture.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/width-swift.property
func (g_ GLKTextureInfo) Width() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("width"))
	return rv
}



