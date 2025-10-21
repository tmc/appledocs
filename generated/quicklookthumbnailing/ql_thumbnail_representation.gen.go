// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ThumbnailRepresentation] class.
var (
	ThumbnailRepresentationClass     _ThumbnailRepresentationClass
	ThumbnailRepresentationClassOnce sync.Once
)

func getThumbnailRepresentationClass() _ThumbnailRepresentationClass {
	ThumbnailRepresentationClassOnce.Do(func() {
		ThumbnailRepresentationClass = _ThumbnailRepresentationClass{objc.GetClass("QLThumbnailRepresentation")}
	})
	return ThumbnailRepresentationClass
}

type _ThumbnailRepresentationClass struct {
	class objc.Class
}

// An interface definition for the [ThumbnailRepresentation] class.
type IThumbnailRepresentation interface {
	objectivec.IObject
}

// Information about the thumbnail that the thumbnail generator returns.
//
// QuickLook Thumbnailing is a non-UI framework, so your app doesn’t have to link to either or . Quicklook Thumbnailing generates a thumbnail as a Core Graphics image object and makes the thumbnail available as the property. If an app links to AppKit or UIKit, the thumbnail is available through the or properties. For more information on the different types of thumbnails that can create, see .
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailRepresentation
type ThumbnailRepresentation struct {
	objectivec.Object
}

// ThumbnailRepresentationFrom constructs a [ThumbnailRepresentation] from an unsafe.Pointer.
//
// Information about the thumbnail that the thumbnail generator returns.
func ThumbnailRepresentationFrom(ptr unsafe.Pointer) ThumbnailRepresentation {
	return ThumbnailRepresentation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _ThumbnailRepresentationClass) Alloc() ThumbnailRepresentation {
	rv := objc.Send[ThumbnailRepresentation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ThumbnailRepresentationClass) New() ThumbnailRepresentation {
	rv := objc.Send[ThumbnailRepresentation](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ThumbnailRepresentation) Init() ThumbnailRepresentation {
	rv := objc.Send[ThumbnailRepresentation](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ThumbnailRepresentation) Autorelease() ThumbnailRepresentation {
	rv := objc.Send[ThumbnailRepresentation](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewThumbnailRepresentation creates a new ThumbnailRepresentation instance.
func NewThumbnailRepresentation() ThumbnailRepresentation {
	return getThumbnailRepresentationClass().New()
}


// A thumbnail in the form of a Core Graphics image object.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailRepresentation/cgImage
func (t_ ThumbnailRepresentation) CGImage() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](t_.ID, objc.Sel("CGImage"))
	return rv
}

// The rectangle within the thumbnail image of the document that represents its contents.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailrepresentation/contentrect
func (t_ ThumbnailRepresentation) ContentRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("contentRect"))
	return rv
}


// SetContentRect sets the value of the contentRect property.
// The rectangle within the thumbnail image of the document that represents its contents.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailrepresentation/contentrect
func (t_ ThumbnailRepresentation) SetContentRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContentRect:"), value)
}

// A thumbnail in the form of an AppKit image object.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailrepresentation/nsimage
func (t_ ThumbnailRepresentation) NsImage() appkit.Image {
	rv := objc.Send[appkit.Image](t_.ID, objc.Sel("nsImage"))
	return rv
}


// SetNsImage sets the value of the nsImage property.
// A thumbnail in the form of an AppKit image object.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailrepresentation/nsimage
func (t_ ThumbnailRepresentation) SetNsImage(value appkit.IImage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNsImage:"), value)
}

// The type of thumbnail.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailrepresentation/type
func (t_ ThumbnailRepresentation) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The type of thumbnail.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailrepresentation/type
func (t_ ThumbnailRepresentation) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setType:"), value)
}

// A thumbnail in the form of a UIKit image object.
//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailrepresentation/uiimage
func (t_ ThumbnailRepresentation) UiImage() appkit.Image {
	rv := objc.Send[appkit.Image](t_.ID, objc.Sel("uiImage"))
	return rv
}


// SetUiImage sets the value of the uiImage property.
// A thumbnail in the form of a UIKit image object.

//
// [Full Topic]: https://developer.apple.com/documentation/quicklookthumbnailing/qlthumbnailrepresentation/uiimage
func (t_ ThumbnailRepresentation) SetUiImage(value appkit.IImage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUiImage:"), value)
}




