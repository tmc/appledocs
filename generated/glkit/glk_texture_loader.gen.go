// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [GLKTextureLoader] class.
var (
	GLKTextureLoaderClass     _GLKTextureLoaderClass
	GLKTextureLoaderClassOnce sync.Once
)

func getGLKTextureLoaderClass() _GLKTextureLoaderClass {
	GLKTextureLoaderClassOnce.Do(func() {
		GLKTextureLoaderClass = _GLKTextureLoaderClass{objc.GetClass("GLKTextureLoader")}
	})
	return GLKTextureLoaderClass
}

type _GLKTextureLoaderClass struct {
	class objc.Class
}

// An interface definition for the [GLKTextureLoader] class.
type IGLKTextureLoader interface {
	objectivec.IObject
	CubeMapWithContentsOfURLOptionsQueueCompletionHandler(url unsafe.Pointer, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer)
	CubeMapWithContentsOfFileOptionsQueueCompletionHandler(path string, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer)
	CubeMapWithContentsOfFilesOptionsQueueCompletionHandler(paths unsafe.Pointer, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer)
	TextureWithCGImageOptionsQueueCompletionHandler(cgImage CGImageRef, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer)
	TextureWithContentsOfURLOptionsQueueCompletionHandler(url unsafe.Pointer, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer)
	TextureWithContentsOfDataOptionsQueueCompletionHandler(data unsafe.Pointer, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer)
	TextureWithContentsOfFileOptionsQueueCompletionHandler(path string, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer)
	TextureWithNameScaleFactorBundleOptionsQueueCompletionHandler(name string, scaleFactor float64, bundle unsafe.Pointer, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer)
}

// A utility class that simplifies loading OpenGL or OpenGL ES texture datas from a variety of image file formats.
//
// The class can load two-dimensional or cubemap textures in most image formats supported by the Image I/O framework. In iOS, it can also load textures compressed in the PVRTC format. It can load the data synchronously or asynchronously. To load textures synchronously, make a context with the desired sharegroup the current context, and then call one or more of the class methods. The returned texture info object includes details about the loaded texture. To load textures asynchronously, your initialization code allocates and initializes a new object using the sharegroup object that should be the destination for new textures. Then, to load a texture, your app calls one of the texture loader’s instance methods, passing in a completion handler block to be called when the texture has been loaded. The following OpenGL properties are set for a newly created, non-mipmapped texture: : : : : The following OpenGL properties are set for a newly created, mipmapped texture: : : : : The and classes do not manage the OpenGL texture for you. Once the texture is returned to your app, you are responsible for it. This means that after your app is finished using an OpenGL texture, it must explicitly deallocate it by calling the function.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader
type GLKTextureLoader struct {
	objectivec.Object
}

// GLKTextureLoaderFrom constructs a [GLKTextureLoader] from an unsafe.Pointer.
//
// A utility class that simplifies loading OpenGL or OpenGL ES texture datas from a variety of image file formats.
func GLKTextureLoaderFrom(ptr unsafe.Pointer) GLKTextureLoader {
	return GLKTextureLoader{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GLKTextureLoaderClass) Alloc() GLKTextureLoader {
	rv := objc.Send[GLKTextureLoader](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GLKTextureLoaderClass) New() GLKTextureLoader {
	rv := objc.Send[GLKTextureLoader](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GLKTextureLoader) Init() GLKTextureLoader {
	rv := objc.Send[GLKTextureLoader](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GLKTextureLoader) Autorelease() GLKTextureLoader {
	rv := objc.Send[GLKTextureLoader](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGLKTextureLoader creates a new GLKTextureLoader instance.
func NewGLKTextureLoader() GLKTextureLoader {
	return getGLKTextureLoaderClass().New()
}


// Initializes a new texture loader object.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/init(share:)
func NewGLKTextureLoaderWithShareContext(context unsafe.Pointer) GLKTextureLoader {
	instance := getGLKTextureLoaderClass().Alloc()
	rv := objc.Send[GLKTextureLoader](instance.ID, objc.Sel("initWithShareContext:"), context)
	rv.Autorelease()
	return rv
}

// Initializes a new texture loader object.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/init(sharegroup:)
func NewGLKTextureLoaderWithSharegroup(sharegroup unsafe.Pointer) GLKTextureLoader {
	instance := getGLKTextureLoaderClass().Alloc()
	rv := objc.Send[GLKTextureLoader](instance.ID, objc.Sel("initWithSharegroup:"), sharegroup)
	rv.Autorelease()
	return rv
}


// Loads a cube map texture image from a single URL and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/cubeMap(withContentsOf:options:)
func (gc _GLKTextureLoaderClass) CubeMapWithContentsOfURLOptionsError(url unsafe.Pointer, options unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("cubeMapWithContentsOfURL:options:error:"), url, options, outError)
	return rv
}

// Loads a cube map texture image from a single file and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/cubeMap(withContentsOfFile:options:)
func (gc _GLKTextureLoaderClass) CubeMapWithContentsOfFileOptionsError(path string, options unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("cubeMapWithContentsOfFile:options:error:"), objc.String(path), options, outError)
	return rv
}

// Loads a cube map texture image from a series of files and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/cubeMap(withContentsOfFiles:options:)
func (gc _GLKTextureLoaderClass) CubeMapWithContentsOfFilesOptionsError(paths unsafe.Pointer, options unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("cubeMapWithContentsOfFiles:options:error:"), paths, options, outError)
	return rv
}

// Loads a 2D texture image from a Quartz image and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/texture(with:options:)
func (gc _GLKTextureLoaderClass) TextureWithCGImageOptionsError(cgImage CGImageRef, options unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("textureWithCGImage:options:error:"), cgImage, options, outError)
	return rv
}

// Loads a 2D texture image from a memory range and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/texture(withContentsOf:options:)-2ljxb
func (gc _GLKTextureLoaderClass) TextureWithContentsOfDataOptionsError(data unsafe.Pointer, options unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("textureWithContentsOfData:options:error:"), data, options, outError)
	return rv
}

// Loads a 2D texture image from a URL and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/texture(withContentsOf:options:)-708ft
func (gc _GLKTextureLoaderClass) TextureWithContentsOfURLOptionsError(url unsafe.Pointer, options unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("textureWithContentsOfURL:options:error:"), url, options, outError)
	return rv
}

// Loads a 2D texture image from a file and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/texture(withContentsOfFile:options:)
func (gc _GLKTextureLoaderClass) TextureWithContentsOfFileOptionsError(path string, options unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("textureWithContentsOfFile:options:error:"), objc.String(path), options, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/texture(withName:scaleFactor:bundle:options:)
func (gc _GLKTextureLoaderClass) TextureWithNameScaleFactorBundleOptionsError(name string, scaleFactor float64, bundle unsafe.Pointer, options unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("textureWithName:scaleFactor:bundle:options:error:"), objc.String(name), scaleFactor, bundle, options, outError)
	return rv
}

// Asynchronously loads a cube map texture image from a single URL and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/cubeMap(withContentsOf:options:queue:completionHandler:)
func (g_ GLKTextureLoader) CubeMapWithContentsOfURLOptionsQueueCompletionHandler(url unsafe.Pointer, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("cubeMapWithContentsOfURL:options:queue:completionHandler:"), url, options, queue, block)
}

// Asynchronously loads a cube map texture image from a single file and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/cubeMap(withContentsOfFile:options:queue:completionHandler:)
func (g_ GLKTextureLoader) CubeMapWithContentsOfFileOptionsQueueCompletionHandler(path string, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("cubeMapWithContentsOfFile:options:queue:completionHandler:"), objc.String(path), options, queue, block)
}

// Asynchronously loads a cube map texture image from a series of files and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/cubeMap(withContentsOfFiles:options:queue:completionHandler:)
func (g_ GLKTextureLoader) CubeMapWithContentsOfFilesOptionsQueueCompletionHandler(paths unsafe.Pointer, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("cubeMapWithContentsOfFiles:options:queue:completionHandler:"), paths, options, queue, block)
}

// Asynchronously loads a 2D texture image from a Quartz image and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/texture(with:options:queue:completionHandler:)
func (g_ GLKTextureLoader) TextureWithCGImageOptionsQueueCompletionHandler(cgImage CGImageRef, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("textureWithCGImage:options:queue:completionHandler:"), cgImage, options, queue, block)
}

// Asynchronously loads a 2D texture image from a URL and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/texture(withContentsOf:options:queue:completionHandler:)-55187
func (g_ GLKTextureLoader) TextureWithContentsOfURLOptionsQueueCompletionHandler(url unsafe.Pointer, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("textureWithContentsOfURL:options:queue:completionHandler:"), url, options, queue, block)
}

// Asynchronously loads a 2D texture image from a memory range and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/texture(withContentsOf:options:queue:completionHandler:)-6n0cf
func (g_ GLKTextureLoader) TextureWithContentsOfDataOptionsQueueCompletionHandler(data unsafe.Pointer, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("textureWithContentsOfData:options:queue:completionHandler:"), data, options, queue, block)
}

// Asynchronously loads a 2D texture image from a file and creates a new texture from the data.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/texture(withContentsOfFile:options:queue:completionHandler:)
func (g_ GLKTextureLoader) TextureWithContentsOfFileOptionsQueueCompletionHandler(path string, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("textureWithContentsOfFile:options:queue:completionHandler:"), objc.String(path), options, queue, block)
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureLoader/texture(withName:scaleFactor:bundle:options:queue:completionHandler:)
func (g_ GLKTextureLoader) TextureWithNameScaleFactorBundleOptionsQueueCompletionHandler(name string, scaleFactor float64, bundle unsafe.Pointer, options unsafe.Pointer, queue unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("textureWithName:scaleFactor:bundle:options:queue:completionHandler:"), objc.String(name), scaleFactor, bundle, options, queue, block)
}


