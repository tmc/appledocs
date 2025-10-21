// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INImage] class.
var (
	INImageClass     _INImageClass
	INImageClassOnce sync.Once
)

func getINImageClass() _INImageClass {
	INImageClassOnce.Do(func() {
		INImageClass = _INImageClass{objc.GetClass("INImage")}
	})
	return INImageClass
}

type _INImageClass struct {
	class objc.Class
}

// An interface definition for the [INImage] class.
type IINImage interface {
	objectivec.IObject
	FetchUIImageWithCompletion(completion unsafe.Pointer)
}

// Image data inside an Intents extension or Intents UI extension.
//
// is a wrapper for image data that you include in a response to an intent. When providing a response to an intent, you must specify any image parameters using instances of this class. supports the same formats as the underlying platform. When confirming or handling an intent, you provide a response object with details about how your app handles that intent. When that response contains an image, use the method to determine the required image size and then create an instance of this class with the corresponding image data.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage
type INImage struct {
	objectivec.Object
}

// INImageFrom constructs a [INImage] from an unsafe.Pointer.
//
// Image data inside an Intents extension or Intents UI extension.
func INImageFrom(ptr unsafe.Pointer) INImage {
	return INImage{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INImageClass) Alloc() INImage {
	rv := objc.Send[INImage](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INImageClass) New() INImage {
	rv := objc.Send[INImage](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INImage) Init() INImage {
	rv := objc.Send[INImage](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INImage) Autorelease() INImage {
	rv := objc.Send[INImage](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINImage creates a new INImage instance.
func NewINImage() INImage {
	return getINImageClass().New()
}




// Creates an image object from an image file in the extension’s bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(named:)
func NewINImageNamed(name string) INImage {
	rv := objc.Send[INImage](objc.ID(getINImageClass().class), objc.Sel("imageNamed:"), objc.String(name))
	return rv
}



// Creates an image object from the specified Core Graphics image.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(CGImage:)
func NewINImageWithCGImage(imageRef CGImageRef) INImage {
	rv := objc.Send[INImage](objc.ID(getINImageClass().class), objc.Sel("imageWithCGImage:"), imageRef)
	return rv
}



// Creates an image object from the specified data.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(imageData:)
func NewINImageWithImageData(imageData unsafe.Pointer) INImage {
	rv := objc.Send[INImage](objc.ID(getINImageClass().class), objc.Sel("imageWithImageData:"), imageData)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(NSImage:)
func NewINImageWithNSImage(image unsafe.Pointer) INImage {
	rv := objc.Send[INImage](objc.ID(getINImageClass().class), objc.Sel("imageWithNSImage:"), image)
	return rv
}



// Creates an image object from the specified UIKit image.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(UIImage:)
func NewINImageWithUIImage(image unsafe.Pointer) INImage {
	rv := objc.Send[INImage](objc.ID(getINImageClass().class), objc.Sel("imageWithUIImage:"), image)
	return rv
}



// Creates an image object from an image file in the local file system.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(url:)
func NewINImageWithURL(URL foundation.URL) INImage {
	rv := objc.Send[INImage](objc.ID(getINImageClass().class), objc.Sel("imageWithURL:"), URL)
	return rv
}



// Creates an image object, of the specified size, from an image file in the local file system.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(url:width:height:)
func NewINImageWithURLWidthHeight(URL foundation.URL, width unsafe.Pointer, height unsafe.Pointer) INImage {
	rv := objc.Send[INImage](objc.ID(getINImageClass().class), objc.Sel("imageWithURL:width:height:"), URL, width, height)
	return rv
}


// Returns the preferred image size for the specified response object.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/imageSize(for:)
func (ic _INImageClass) ImageSizeForIntentResponse(response unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](objc.ID(ic.class), objc.Sel("imageSizeForIntentResponse:"), response)
	return rv
}

// Creates an image object from the specified Core Graphics image.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(CGImage:)
func (ic _INImageClass) ImageWithCGImage(imageRef CGImageRef) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithCGImage:"), imageRef)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(NSImage:)
func (ic _INImageClass) ImageWithNSImage(image unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithNSImage:"), image)
	return rv
}

// Creates an image object from the specified UIKit image.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(UIImage:)
func (ic _INImageClass) ImageWithUIImage(image unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithUIImage:"), image)
	return rv
}

// Creates an image object from the specified data.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(imageData:)
func (ic _INImageClass) ImageWithImageData(imageData unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithImageData:"), imageData)
	return rv
}

// Creates an image object from an image file in the extension’s bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(named:)
func (ic _INImageClass) ImageNamed(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageNamed:"), objc.String(name))
	return rv
}

// Creates an image object from an image file in the local file system.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(url:)
func (ic _INImageClass) ImageWithURL(URL foundation.URL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithURL:"), URL)
	return rv
}

// Creates an image object, of the specified size, from an image file in the local file system.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/init(url:width:height:)
func (ic _INImageClass) ImageWithURLWidthHeight(URL foundation.URL, width unsafe.Pointer, height unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithURL:width:height:"), URL, width, height)
	return rv
}

// Returns an image object that contains the specified system symbol image.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/systemImageNamed(_:)
func (ic _INImageClass) SystemImageNamed(systemImageName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("systemImageNamed:"), objc.String(systemImageName))
	return rv
}

// Fetches the image and provides it to the specified completion handler.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INImage/fetchUIImage(completion:)
func (i_ INImage) FetchUIImageWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("fetchUIImageWithCompletion:"), completion)
}


