// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextureLoader] class.
var (
	TextureLoaderClass     _TextureLoaderClass
	TextureLoaderClassOnce sync.Once
)

func getTextureLoaderClass() _TextureLoaderClass {
	TextureLoaderClassOnce.Do(func() {
		TextureLoaderClass = _TextureLoaderClass{objc.GetClass("MTKTextureLoader")}
	})
	return TextureLoaderClass
}

type _TextureLoaderClass struct {
	class objc.Class
}

// An interface definition for the [TextureLoader] class.
type ITextureLoader interface {
	objectivec.IObject
	// properties:
	Device() objc.ID
	// methods:
	NewTextureWithContentsOfURLOptionsError(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary, error_ unsafe.Pointer) objc.ID
	NewTextureWithContentsOfURLOptionsCompletionHandler(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */)
	NewTextureWithCGImageOptionsError(cgImage ImageRef /* not a class type */, options foundation.IDictionary, error_ unsafe.Pointer) objc.ID
	NewTextureWithCGImageOptionsCompletionHandler(cgImage ImageRef /* not a class type */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */)
	NewTextureWithDataOptionsError(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, error_ unsafe.Pointer) objc.ID
	NewTextureWithDataOptionsCompletionHandler(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */)
	NewTextureWithNameScaleFactorBundleOptionsError(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, bundle objc.IObject /* cross-framework: Bundle */, options foundation.IDictionary, error_ unsafe.Pointer) objc.ID
	NewTextureWithNameScaleFactorBundleOptionsCompletionHandler(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, bundle objc.IObject /* cross-framework: Bundle */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */)
	NewTextureWithNameScaleFactorDisplayGamutBundleOptionsError(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, displayGamut DisplayGamut /* not a class type */, bundle objc.IObject /* cross-framework: Bundle */, options foundation.IDictionary, error_ unsafe.Pointer) objc.ID
	NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, displayGamut DisplayGamut /* not a class type */, bundle objc.IObject /* cross-framework: Bundle */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */)
	NewTextureWithMDLTextureOptionsError(texture unsafe.Pointer, options foundation.IDictionary, error_ unsafe.Pointer) objc.ID
	NewTextureWithMDLTextureOptionsCompletionHandler(texture unsafe.Pointer, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */)
	NewTexturesWithContentsOfURLsOptionsCompletionHandler(URLs []objc.IObject /* cross-framework: URL */, options foundation.IDictionary, completionHandler TextureLoaderArrayCallback /* not a class type */)
	NewTexturesWithContentsOfURLsOptionsError(URLs []objc.IObject /* cross-framework: URL */, options foundation.IDictionary, error_ unsafe.Pointer) []objc.ID
	NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler(names []string, scaleFactor float64, bundle objc.IObject /* cross-framework: Bundle */, options foundation.IDictionary, completionHandler TextureLoaderArrayCallback /* not a class type */)
	NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler(names []string, scaleFactor float64, displayGamut DisplayGamut /* not a class type */, bundle objc.IObject /* cross-framework: Bundle */, options foundation.IDictionary, completionHandler TextureLoaderArrayCallback /* not a class type */)
}

// An object that creates textures from existing data in common image formats.
//
// Use the class to create a Metal texture from existing image data. This class supports common file formats, like PNG, JPEG, and TIFF. It also loads image data from KTX and PVR files, asset catalogs, Core Graphics images, and other sources. It infers the output texture format and pixel format from the image data. You create textures synchronously or asynchronously using methods that return instances. Pass options to these methods that customize the image-loading and texture-creation process. First create an instance, passing the device that it uses to create textures. Then use one of the texture loader’s methods to create a texture. The code example below synchronously creates a texture from data at a URL, using the default options: If you use custom data formats, or change the image data at runtime, use methods instead. For more information, see .


// An object that creates textures from existing data in common image formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader
type TextureLoader struct {
	objectivec.Object
}

// TextureLoaderFrom constructs a [TextureLoader] from an unsafe.Pointer.
//
// An object that creates textures from existing data in common image formats.
func TextureLoaderFrom(ptr unsafe.Pointer) TextureLoader {
	return TextureLoader{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextureLoaderClass) Alloc() TextureLoader {
	rv := objc.Send[TextureLoader](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextureLoaderClass) New() TextureLoader {
	rv := objc.Send[TextureLoader](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextureLoader) Init() TextureLoader {
	rv := objc.Send[TextureLoader](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextureLoader) Autorelease() TextureLoader {
	rv := objc.Send[TextureLoader](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextureLoader creates a new TextureLoader instance.
func NewTextureLoader() TextureLoader {
	return getTextureLoaderClass().New()
}



// Initializes a new texture loader object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/init(device:)
func NewTextureLoaderWithDevice(device objectivec.IObject) TextureLoader {
	instance := getTextureLoaderClass().Alloc()
	rv := objc.Send[TextureLoader](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



// Synchronously loads image data and creates a new Metal texture from a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(URL:options:)
func (t_ TextureLoader) NewTextureWithContentsOfURLOptionsError(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithContentsOfURL:options:error:"), URL, options, error_)
	return rv
}


// Asynchronously loads image data and creates a new Metal texture from a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(URL:options:completionHandler:)
func (t_ TextureLoader) NewTextureWithContentsOfURLOptionsCompletionHandler(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithContentsOfURL:options:completionHandler:"), URL, options, completionHandler)
}


// Synchronously loads image data and creates a new Metal texture from a given bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(cgImage:options:)
func (t_ TextureLoader) NewTextureWithCGImageOptionsError(cgImage ImageRef /* not a class type */, options foundation.IDictionary, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithCGImage:options:error:"), cgImage, options, error_)
	return rv
}


// Asynchronously loads image data and creates a new Metal texture from a given bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(cgImage:options:completionHandler:)
func (t_ TextureLoader) NewTextureWithCGImageOptionsCompletionHandler(cgImage ImageRef /* not a class type */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithCGImage:options:completionHandler:"), cgImage, options, completionHandler)
}


// Synchronously creates a new Metal texture from an in-memory representation of the texture’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(data:options:)
func (t_ TextureLoader) NewTextureWithDataOptionsError(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithData:options:error:"), data, options, error_)
	return rv
}


// Asynchronously creates a new Metal texture from an in-memory representation of the texture’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(data:options:completionHandler:)
func (t_ TextureLoader) NewTextureWithDataOptionsCompletionHandler(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithData:options:completionHandler:"), data, options, completionHandler)
}


// Synchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:bundle:options:)
func (t_ TextureLoader) NewTextureWithNameScaleFactorBundleOptionsError(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, bundle objc.IObject /* cross-framework: Bundle */, options foundation.IDictionary, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithName:scaleFactor:bundle:options:error:"), name, scaleFactor, bundle, options, error_)
	return rv
}


// Asynchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:bundle:options:completionHandler:)
func (t_ TextureLoader) NewTextureWithNameScaleFactorBundleOptionsCompletionHandler(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, bundle objc.IObject /* cross-framework: Bundle */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithName:scaleFactor:bundle:options:completionHandler:"), name, scaleFactor, bundle, options, completionHandler)
}


// Synchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog, using a specified display gamut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:displayGamut:bundle:options:)
func (t_ TextureLoader) NewTextureWithNameScaleFactorDisplayGamutBundleOptionsError(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, displayGamut DisplayGamut /* not a class type */, bundle objc.IObject /* cross-framework: Bundle */, options foundation.IDictionary, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithName:scaleFactor:displayGamut:bundle:options:error:"), name, scaleFactor, displayGamut, bundle, options, error_)
	return rv
}


// Asynchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:displayGamut:bundle:options:completionHandler:)
func (t_ TextureLoader) NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, displayGamut DisplayGamut /* not a class type */, bundle objc.IObject /* cross-framework: Bundle */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithName:scaleFactor:displayGamut:bundle:options:completionHandler:"), name, scaleFactor, displayGamut, bundle, options, completionHandler)
}


// Synchronously loads image data and creates a Metal texture from the specified Model I/O texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(texture:options:)
func (t_ TextureLoader) NewTextureWithMDLTextureOptionsError(texture unsafe.Pointer, options foundation.IDictionary, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithMDLTexture:options:error:"), texture, options, error_)
	return rv
}


// Asynchronously loads image data and creates a Metal texture from the specified Model I/O texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(texture:options:completionHandler:)
func (t_ TextureLoader) NewTextureWithMDLTextureOptionsCompletionHandler(texture unsafe.Pointer, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithMDLTexture:options:completionHandler:"), texture, options, completionHandler)
}


// Asynchronously loads image data and creates new Metal textures from the specified list of URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(URLs:options:completionHandler:)
func (t_ TextureLoader) NewTexturesWithContentsOfURLsOptionsCompletionHandler(URLs []objc.IObject /* cross-framework: URL */, options foundation.IDictionary, completionHandler TextureLoaderArrayCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTexturesWithContentsOfURLs:options:completionHandler:"), URLs, options, completionHandler)
}


// Synchronously loads image data and creates new Metal textures from the specified list of URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(URLs:options:error:)
func (t_ TextureLoader) NewTexturesWithContentsOfURLsOptionsError(URLs []objc.IObject /* cross-framework: URL */, options foundation.IDictionary, error_ unsafe.Pointer) []objc.ID {
	rv := objc.Send[[]objc.ID](t_.ID, objc.Sel("newTexturesWithContentsOfURLs:options:error:"), URLs, options, error_)
	return rv
}


// Asynchronously loads image data and creates Metal textures from the specified list of named texture assets in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(names:scaleFactor:bundle:options:completionHandler:)
func (t_ TextureLoader) NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler(names []string, scaleFactor float64, bundle objc.IObject /* cross-framework: Bundle */, options foundation.IDictionary, completionHandler TextureLoaderArrayCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTexturesWithNames:scaleFactor:bundle:options:completionHandler:"), names, scaleFactor, bundle, options, completionHandler)
}


// Asynchronously loads image data and creates Metal textures from the specified list of named texture assets in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(names:scaleFactor:displayGamut:bundle:options:completionHandler:)
func (t_ TextureLoader) NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler(names []string, scaleFactor float64, displayGamut DisplayGamut /* not a class type */, bundle objc.IObject /* cross-framework: Bundle */, options foundation.IDictionary, completionHandler TextureLoaderArrayCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTexturesWithNames:scaleFactor:displayGamut:bundle:options:completionHandler:"), names, scaleFactor, displayGamut, bundle, options, completionHandler)
}


// The device object that the texture loader uses to create textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/device
func (t_ TextureLoader) Device() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("device"))
	return rv
}


