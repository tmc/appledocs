// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [MTKTextureLoader] class.
var (
	mTKTextureLoaderClass     _MTKTextureLoaderClass
	mTKTextureLoaderClassOnce sync.Once
)

func getMTKTextureLoaderClass() _MTKTextureLoaderClass {
	mTKTextureLoaderClassOnce.Do(func() {
		mTKTextureLoaderClass = _MTKTextureLoaderClass{objc.GetClass("MTKTextureLoader")}
	})
	return mTKTextureLoaderClass
}

type _MTKTextureLoaderClass struct {
	class objc.Class
}

// An interface definition for the [MTKTextureLoader] class.
type IMTKTextureLoader interface {
	objectivec.IObject
	NewTextureWithContentsOfURLOptionsError(URL unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	NewTextureWithContentsOfURLOptionsCompletionHandler(URL unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer)
	NewTextureWithCGImageOptionsError(cgImage coregraphics.CGImageRef, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	NewTextureWithCGImageOptionsCompletionHandler(cgImage coregraphics.CGImageRef, options unsafe.Pointer, completionHandler unsafe.Pointer)
	NewTextureWithDataOptionsError(data unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	NewTextureWithDataOptionsCompletionHandler(data unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer)
	NewTextureWithNameScaleFactorBundleOptionsError(name string, scaleFactor float64, bundle unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	NewTextureWithNameScaleFactorBundleOptionsCompletionHandler(name string, scaleFactor float64, bundle unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer)
	NewTextureWithNameScaleFactorDisplayGamutBundleOptionsError(name string, scaleFactor float64, displayGamut unsafe.Pointer, bundle unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler(name string, scaleFactor float64, displayGamut unsafe.Pointer, bundle unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer)
	NewTextureWithMDLTextureOptionsError(texture unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	NewTextureWithMDLTextureOptionsCompletionHandler(texture unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer)
	NewTexturesWithContentsOfURLsOptionsCompletionHandler(URLs unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer)
	NewTexturesWithContentsOfURLsOptionsError(URLs unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer
	NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler(names unsafe.Pointer, scaleFactor float64, bundle unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer)
	NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler(names unsafe.Pointer, scaleFactor float64, displayGamut unsafe.Pointer, bundle unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer)
}

// An object that creates textures from existing data in common image formats.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader
type MTKTextureLoader struct {
	objectivec.Object
}

// MTKTextureLoaderFrom constructs a [MTKTextureLoader] from an unsafe.Pointer.
//
// An object that creates textures from existing data in common image formats.
func MTKTextureLoaderFrom(ptr unsafe.Pointer) MTKTextureLoader {
	return MTKTextureLoader{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTKTextureLoaderClass) Alloc() MTKTextureLoader {
	rv := objc.Send[MTKTextureLoader](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTKTextureLoaderClass) New() MTKTextureLoader {
	rv := objc.Send[MTKTextureLoader](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTKTextureLoader) Init() MTKTextureLoader {
	rv := objc.Send[MTKTextureLoader](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTKTextureLoader) Autorelease() MTKTextureLoader {
	rv := objc.Send[MTKTextureLoader](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTKTextureLoader creates a new MTKTextureLoader instance.
func NewMTKTextureLoader() MTKTextureLoader {
	return getMTKTextureLoaderClass().New()
}


// Initializes a new texture loader object.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/init(device:)
func NewMTKTextureLoaderWithDevice(device unsafe.Pointer) MTKTextureLoader {
	instance := getMTKTextureLoaderClass().Alloc()
	rv := objc.Send[MTKTextureLoader](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// Synchronously loads image data and creates a new Metal texture from a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(URL:options:)
func (m_ MTKTextureLoader) NewTextureWithContentsOfURLOptionsError(URL unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("newTextureWithContentsOfURL:options:error:"), URL, options, error)
	return rv
}
// Asynchronously loads image data and creates a new Metal texture from a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(URL:options:completionHandler:)
func (m_ MTKTextureLoader) NewTextureWithContentsOfURLOptionsCompletionHandler(URL unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("newTextureWithContentsOfURL:options:completionHandler:"), URL, options, completionHandler)
}
// Synchronously loads image data and creates a new Metal texture from a given bitmap image.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(cgImage:options:)
func (m_ MTKTextureLoader) NewTextureWithCGImageOptionsError(cgImage coregraphics.CGImageRef, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("newTextureWithCGImage:options:error:"), cgImage, options, error)
	return rv
}
// Asynchronously loads image data and creates a new Metal texture from a given bitmap image.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(cgImage:options:completionHandler:)
func (m_ MTKTextureLoader) NewTextureWithCGImageOptionsCompletionHandler(cgImage coregraphics.CGImageRef, options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("newTextureWithCGImage:options:completionHandler:"), cgImage, options, completionHandler)
}
// Synchronously creates a new Metal texture from an in-memory representation of the texture’s data.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(data:options:)
func (m_ MTKTextureLoader) NewTextureWithDataOptionsError(data unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("newTextureWithData:options:error:"), data, options, error)
	return rv
}
// Asynchronously creates a new Metal texture from an in-memory representation of the texture’s data.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(data:options:completionHandler:)
func (m_ MTKTextureLoader) NewTextureWithDataOptionsCompletionHandler(data unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("newTextureWithData:options:completionHandler:"), data, options, completionHandler)
}
// Synchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:bundle:options:)
func (m_ MTKTextureLoader) NewTextureWithNameScaleFactorBundleOptionsError(name string, scaleFactor float64, bundle unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("newTextureWithName:scaleFactor:bundle:options:error:"), objc.String(name), scaleFactor, bundle, options, error)
	return rv
}
// Asynchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:bundle:options:completionHandler:)
func (m_ MTKTextureLoader) NewTextureWithNameScaleFactorBundleOptionsCompletionHandler(name string, scaleFactor float64, bundle unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("newTextureWithName:scaleFactor:bundle:options:completionHandler:"), objc.String(name), scaleFactor, bundle, options, completionHandler)
}
// Synchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog, using a specified display gamut.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:displayGamut:bundle:options:)
func (m_ MTKTextureLoader) NewTextureWithNameScaleFactorDisplayGamutBundleOptionsError(name string, scaleFactor float64, displayGamut unsafe.Pointer, bundle unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("newTextureWithName:scaleFactor:displayGamut:bundle:options:error:"), objc.String(name), scaleFactor, displayGamut, bundle, options, error)
	return rv
}
// Asynchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:displayGamut:bundle:options:completionHandler:)
func (m_ MTKTextureLoader) NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler(name string, scaleFactor float64, displayGamut unsafe.Pointer, bundle unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("newTextureWithName:scaleFactor:displayGamut:bundle:options:completionHandler:"), objc.String(name), scaleFactor, displayGamut, bundle, options, completionHandler)
}
// Synchronously loads image data and creates a Metal texture from the specified Model I/O texture.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(texture:options:)
func (m_ MTKTextureLoader) NewTextureWithMDLTextureOptionsError(texture unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("newTextureWithMDLTexture:options:error:"), texture, options, error)
	return rv
}
// Asynchronously loads image data and creates a Metal texture from the specified Model I/O texture.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(texture:options:completionHandler:)
func (m_ MTKTextureLoader) NewTextureWithMDLTextureOptionsCompletionHandler(texture unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("newTextureWithMDLTexture:options:completionHandler:"), texture, options, completionHandler)
}
// Asynchronously loads image data and creates new Metal textures from the specified list of URLs.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(URLs:options:completionHandler:)
func (m_ MTKTextureLoader) NewTexturesWithContentsOfURLsOptionsCompletionHandler(URLs unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("newTexturesWithContentsOfURLs:options:completionHandler:"), URLs, options, completionHandler)
}
// Synchronously loads image data and creates new Metal textures from the specified list of URLs.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(URLs:options:error:)
func (m_ MTKTextureLoader) NewTexturesWithContentsOfURLsOptionsError(URLs unsafe.Pointer, options unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("newTexturesWithContentsOfURLs:options:error:"), URLs, options, error)
	return rv
}
// Asynchronously loads image data and creates Metal textures from the specified list of named texture assets in an asset catalog.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(names:scaleFactor:bundle:options:completionHandler:)
func (m_ MTKTextureLoader) NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler(names unsafe.Pointer, scaleFactor float64, bundle unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("newTexturesWithNames:scaleFactor:bundle:options:completionHandler:"), names, scaleFactor, bundle, options, completionHandler)
}
// Asynchronously loads image data and creates Metal textures from the specified list of named texture assets in an asset catalog.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(names:scaleFactor:displayGamut:bundle:options:completionHandler:)
func (m_ MTKTextureLoader) NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler(names unsafe.Pointer, scaleFactor float64, displayGamut unsafe.Pointer, bundle unsafe.Pointer, options unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("newTexturesWithNames:scaleFactor:displayGamut:bundle:options:completionHandler:"), names, scaleFactor, displayGamut, bundle, options, completionHandler)
}

