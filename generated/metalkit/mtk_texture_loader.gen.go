// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTKTextureLoader */


/* debug [class_header]: Header for MTKTextureLoader */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextureLoader */
// An interface definition for the [TextureLoader] class.
type ITextureLoader interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextureLoader */
	// properties:
	Device() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextureLoader */
	// methods:
	NewTextureWithCGImageOptionsError(cgImage ImageRef /* not a class type */, options foundation.IDictionary, error_ objectivec.IObject) unsafe.Pointer
	NewTextureWithCGImageOptionsCompletionHandler(cgImage ImageRef /* not a class type */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */)
	NewTextureWithDataOptionsError(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, error_ objectivec.IObject) unsafe.Pointer
	NewTextureWithDataOptionsCompletionHandler(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */)
	NewTextureWithNameScaleFactorBundleOptionsError(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, bundle foundation.Bundle, options foundation.IDictionary, error_ objectivec.IObject) unsafe.Pointer
	NewTextureWithNameScaleFactorBundleOptionsCompletionHandler(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, bundle foundation.Bundle, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */)
	NewTextureWithNameScaleFactorDisplayGamutBundleOptionsError(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, displayGamut DisplayGamut /* not a class type */, bundle foundation.Bundle, options foundation.IDictionary, error_ objectivec.IObject) unsafe.Pointer
	NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, displayGamut DisplayGamut /* not a class type */, bundle foundation.Bundle, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */)
	NewTextureWithMDLTextureOptionsError(texture objectivec.IObject, options foundation.IDictionary, error_ objectivec.IObject) unsafe.Pointer
	NewTextureWithMDLTextureOptionsCompletionHandler(texture objectivec.IObject, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */)
	NewTextureWithContentsOfURLOptionsError(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary, error_ objectivec.IObject) unsafe.Pointer
	NewTextureWithContentsOfURLOptionsCompletionHandler(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */)
	NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler(names []string, scaleFactor float64, bundle foundation.Bundle, options foundation.IDictionary, completionHandler TextureLoaderArrayCallback /* not a class type */)
	NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler(names []string, scaleFactor float64, displayGamut DisplayGamut /* not a class type */, bundle foundation.Bundle, options foundation.IDictionary, completionHandler TextureLoaderArrayCallback /* not a class type */)
	NewTexturesWithContentsOfURLsOptionsCompletionHandler(URLs []foundation.URL, options foundation.IDictionary, completionHandler TextureLoaderArrayCallback /* not a class type */)
	NewTexturesWithContentsOfURLsOptionsError(URLs []foundation.URL, options foundation.IDictionary, error_ objectivec.IObject) []objc.ID
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextureLoader */
// Alloc allocates a new instance without initialization.
func (tc _TextureLoaderClass) Alloc() TextureLoader {
	rv := objc.Send[TextureLoader](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextureLoader */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextureLoader */

// Initializes a new texture loader object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/init(device:)
func NewTextureLoaderWithDevice(device unsafe.Pointer) TextureLoader {
	instance := getTextureLoaderClass().Alloc()
	rv := objc.Send[TextureLoader](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextureLoaderWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextureLoader */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextureLoader */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextureLoader */

// Synchronously loads image data and creates a new Metal texture from a given bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(cgImage:options:)
func (t_ TextureLoader) NewTextureWithCGImageOptionsError(cgImage ImageRef /* not a class type */, options foundation.IDictionary, error_ objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("newTextureWithCGImage:options:error:"), cgImage, options, error_)
	return rv
}/* debug [instance_methods/method]: NewTextureWithCGImageOptionsError */


// Asynchronously loads image data and creates a new Metal texture from a given bitmap image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(cgImage:options:completionHandler:)
func (t_ TextureLoader) NewTextureWithCGImageOptionsCompletionHandler(cgImage ImageRef /* not a class type */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithCGImage:options:completionHandler:"), cgImage, options, completionHandler)
}/* debug [instance_methods/method]: NewTextureWithCGImageOptionsCompletionHandler */


// Synchronously creates a new Metal texture from an in-memory representation of the texture’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(data:options:)
func (t_ TextureLoader) NewTextureWithDataOptionsError(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, error_ objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("newTextureWithData:options:error:"), data, options, error_)
	return rv
}/* debug [instance_methods/method]: NewTextureWithDataOptionsError */


// Asynchronously creates a new Metal texture from an in-memory representation of the texture’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(data:options:completionHandler:)
func (t_ TextureLoader) NewTextureWithDataOptionsCompletionHandler(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithData:options:completionHandler:"), data, options, completionHandler)
}/* debug [instance_methods/method]: NewTextureWithDataOptionsCompletionHandler */


// Synchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:bundle:options:)
func (t_ TextureLoader) NewTextureWithNameScaleFactorBundleOptionsError(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, bundle foundation.Bundle, options foundation.IDictionary, error_ objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("newTextureWithName:scaleFactor:bundle:options:error:"), name, scaleFactor, bundle, options, error_)
	return rv
}/* debug [instance_methods/method]: NewTextureWithNameScaleFactorBundleOptionsError */


// Asynchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:bundle:options:completionHandler:)
func (t_ TextureLoader) NewTextureWithNameScaleFactorBundleOptionsCompletionHandler(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, bundle foundation.Bundle, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithName:scaleFactor:bundle:options:completionHandler:"), name, scaleFactor, bundle, options, completionHandler)
}/* debug [instance_methods/method]: NewTextureWithNameScaleFactorBundleOptionsCompletionHandler */


// Synchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog, using a specified display gamut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:displayGamut:bundle:options:)
func (t_ TextureLoader) NewTextureWithNameScaleFactorDisplayGamutBundleOptionsError(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, displayGamut DisplayGamut /* not a class type */, bundle foundation.Bundle, options foundation.IDictionary, error_ objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("newTextureWithName:scaleFactor:displayGamut:bundle:options:error:"), name, scaleFactor, displayGamut, bundle, options, error_)
	return rv
}/* debug [instance_methods/method]: NewTextureWithNameScaleFactorDisplayGamutBundleOptionsError */


// Asynchronously loads image data and creates a Metal texture from the named texture asset in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(name:scaleFactor:displayGamut:bundle:options:completionHandler:)
func (t_ TextureLoader) NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler(name objc.IObject /* cross-framework: NSString */, scaleFactor float64, displayGamut DisplayGamut /* not a class type */, bundle foundation.Bundle, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithName:scaleFactor:displayGamut:bundle:options:completionHandler:"), name, scaleFactor, displayGamut, bundle, options, completionHandler)
}/* debug [instance_methods/method]: NewTextureWithNameScaleFactorDisplayGamutBundleOptionsCompletionHandler */


// Synchronously loads image data and creates a Metal texture from the specified Model I/O texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(texture:options:)
func (t_ TextureLoader) NewTextureWithMDLTextureOptionsError(texture objectivec.IObject, options foundation.IDictionary, error_ objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("newTextureWithMDLTexture:options:error:"), texture, options, error_)
	return rv
}/* debug [instance_methods/method]: NewTextureWithMDLTextureOptionsError */


// Asynchronously loads image data and creates a Metal texture from the specified Model I/O texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(texture:options:completionHandler:)
func (t_ TextureLoader) NewTextureWithMDLTextureOptionsCompletionHandler(texture objectivec.IObject, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithMDLTexture:options:completionHandler:"), texture, options, completionHandler)
}/* debug [instance_methods/method]: NewTextureWithMDLTextureOptionsCompletionHandler */


// Synchronously loads image data and creates a new Metal texture from a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(URL:options:)
func (t_ TextureLoader) NewTextureWithContentsOfURLOptionsError(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary, error_ objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("newTextureWithContentsOfURL:options:error:"), URL, options, error_)
	return rv
}/* debug [instance_methods/method]: NewTextureWithContentsOfURLOptionsError */


// Asynchronously loads image data and creates a new Metal texture from a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTexture(URL:options:completionHandler:)
func (t_ TextureLoader) NewTextureWithContentsOfURLOptionsCompletionHandler(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary, completionHandler TextureLoaderCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTextureWithContentsOfURL:options:completionHandler:"), URL, options, completionHandler)
}/* debug [instance_methods/method]: NewTextureWithContentsOfURLOptionsCompletionHandler */


// Asynchronously loads image data and creates Metal textures from the specified list of named texture assets in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(names:scaleFactor:bundle:options:completionHandler:)
func (t_ TextureLoader) NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler(names []string, scaleFactor float64, bundle foundation.Bundle, options foundation.IDictionary, completionHandler TextureLoaderArrayCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTexturesWithNames:scaleFactor:bundle:options:completionHandler:"), names, scaleFactor, bundle, options, completionHandler)
}/* debug [instance_methods/method]: NewTexturesWithNamesScaleFactorBundleOptionsCompletionHandler */


// Asynchronously loads image data and creates Metal textures from the specified list of named texture assets in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(names:scaleFactor:displayGamut:bundle:options:completionHandler:)
func (t_ TextureLoader) NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler(names []string, scaleFactor float64, displayGamut DisplayGamut /* not a class type */, bundle foundation.Bundle, options foundation.IDictionary, completionHandler TextureLoaderArrayCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTexturesWithNames:scaleFactor:displayGamut:bundle:options:completionHandler:"), names, scaleFactor, displayGamut, bundle, options, completionHandler)
}/* debug [instance_methods/method]: NewTexturesWithNamesScaleFactorDisplayGamutBundleOptionsCompletionHandler */


// Asynchronously loads image data and creates new Metal textures from the specified list of URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(URLs:options:completionHandler:)
func (t_ TextureLoader) NewTexturesWithContentsOfURLsOptionsCompletionHandler(URLs []foundation.URL, options foundation.IDictionary, completionHandler TextureLoaderArrayCallback /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("newTexturesWithContentsOfURLs:options:completionHandler:"), URLs, options, completionHandler)
}/* debug [instance_methods/method]: NewTexturesWithContentsOfURLsOptionsCompletionHandler */


// Synchronously loads image data and creates new Metal textures from the specified list of URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/newTextures(URLs:options:error:)
func (t_ TextureLoader) NewTexturesWithContentsOfURLsOptionsError(URLs []foundation.URL, options foundation.IDictionary, error_ objectivec.IObject) []objc.ID {
	rv := objc.Send[[]objc.ID](t_.ID, objc.Sel("newTexturesWithContentsOfURLs:options:error:"), URLs, options, error_)
	return rv
}/* debug [instance_methods/method]: NewTexturesWithContentsOfURLsOptionsError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextureLoader */

// The device object that the texture loader uses to create textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/device
func (t_ TextureLoader) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTKTextureLoader */


