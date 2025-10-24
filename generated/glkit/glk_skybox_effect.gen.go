// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GLKSkyboxEffect] class.
var (
	GLKSkyboxEffectClass     _GLKSkyboxEffectClass
	GLKSkyboxEffectClassOnce sync.Once
)

func getGLKSkyboxEffectClass() _GLKSkyboxEffectClass {
	GLKSkyboxEffectClassOnce.Do(func() {
		GLKSkyboxEffectClass = _GLKSkyboxEffectClass{objc.GetClass("GLKSkyboxEffect")}
	})
	return GLKSkyboxEffectClass
}

type _GLKSkyboxEffectClass struct {
	class objc.Class
}

// An interface definition for the [GLKSkyboxEffect] class.
type IGLKSkyboxEffect interface {
	objectivec.IObject
	// properties:
	Center() GLKVector3 /* typedef */
	SetCenter(value GLKVector3 /* typedef */)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	TextureCubeMap() IGLKEffectPropertyTexture
	Transform() IGLKEffectPropertyTransform
	XSize() unsafe.Pointer
	SetXSize(value unsafe.Pointer)
	YSize() unsafe.Pointer
	SetYSize(value unsafe.Pointer)
	ZSize() unsafe.Pointer
	SetZSize(value unsafe.Pointer)
	// methods:
}

// A simple skybox visual effect for use in shader-based OpenGL rendering.
//
// The provides a standard skybox effect for your application. Unlike the class, the skybox does not require your application to configure and submit vertex data. Instead, it creates its own vertex data based on the configuration data you supply. At initialization time, your application first creates a compatible context and makes it current. Then, it creates new skybox effect, configures its properties, and calls its method. Binding the effect causes a shader to be compiled and bound to the current context. At rendering time, your application calls the effect’s method to prepare the effect and then calls its method to draw the sky box.


// A simple skybox visual effect for use in shader-based OpenGL rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSkyboxEffect
type GLKSkyboxEffect struct {
	objectivec.Object
}

// GLKSkyboxEffectFrom constructs a [GLKSkyboxEffect] from an unsafe.Pointer.
//
// A simple skybox visual effect for use in shader-based OpenGL rendering.
func GLKSkyboxEffectFrom(ptr unsafe.Pointer) GLKSkyboxEffect {
	return GLKSkyboxEffect{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GLKSkyboxEffectClass) Alloc() GLKSkyboxEffect {
	rv := objc.Send[GLKSkyboxEffect](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GLKSkyboxEffectClass) New() GLKSkyboxEffect {
	rv := objc.Send[GLKSkyboxEffect](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKSkyboxEffect) Init() GLKSkyboxEffect {
	rv := objc.Send[GLKSkyboxEffect](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKSkyboxEffect) Autorelease() GLKSkyboxEffect {
	rv := objc.Send[GLKSkyboxEffect](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKSkyboxEffect creates a new GLKSkyboxEffect instance.
func NewGLKSkyboxEffect() GLKSkyboxEffect {
	return getGLKSkyboxEffectClass().New()
}



// The center of the skybox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSkyboxEffect/center
func (g_ GLKSkyboxEffect) Center() GLKVector3 /* typedef */ {
	rv := objc.Send[GLKVector3](g_.ID, objc.Sel("center"))
	return rv
}


// The center of the skybox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSkyboxEffect/center
func (g_ GLKSkyboxEffect) SetCenter(value GLKVector3 /* typedef */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCenter:"), value)
}


// A string used to name your effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSkyboxEffect/label
func (g_ GLKSkyboxEffect) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("label"))
	return rv
}


// A string used to name your effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSkyboxEffect/label
func (g_ GLKSkyboxEffect) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLabel:"), value)
}


// The texture to apply to the skybox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSkyboxEffect/textureCubeMap
func (g_ GLKSkyboxEffect) TextureCubeMap() IGLKEffectPropertyTexture {
	rv := objc.Send[GLKEffectPropertyTexture](g_.ID, objc.Sel("textureCubeMap"))
	return rv
}


// The transform applied before drawing the skybox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSkyboxEffect/transform
func (g_ GLKSkyboxEffect) Transform() IGLKEffectPropertyTransform {
	rv := objc.Send[GLKEffectPropertyTransform](g_.ID, objc.Sel("transform"))
	return rv
}


// The width of the skybox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSkyboxEffect/xSize
func (g_ GLKSkyboxEffect) XSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("xSize"))
	return rv
}


// The width of the skybox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSkyboxEffect/xSize
func (g_ GLKSkyboxEffect) SetXSize(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setXSize:"), value)
}


// The height of the skybox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSkyboxEffect/ySize
func (g_ GLKSkyboxEffect) YSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("ySize"))
	return rv
}


// The height of the skybox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSkyboxEffect/ySize
func (g_ GLKSkyboxEffect) SetYSize(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setYSize:"), value)
}


// The depth of the skybox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSkyboxEffect/zSize
func (g_ GLKSkyboxEffect) ZSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("zSize"))
	return rv
}


// The depth of the skybox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKSkyboxEffect/zSize
func (g_ GLKSkyboxEffect) SetZSize(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setZSize:"), value)
}



