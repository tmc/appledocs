// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSImageCell */


/* debug [class_header]: Header for NSImageCell */
// The class instance for the [ImageCell] class.
var (
	ImageCellClass     _ImageCellClass
	ImageCellClassOnce sync.Once
)

func getImageCellClass() _ImageCellClass {
	ImageCellClassOnce.Do(func() {
		ImageCellClass = _ImageCellClass{objc.GetClass("NSImageCell")}
	})
	return ImageCellClass
}

type _ImageCellClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageCell */
// An interface definition for the [ImageCell] class.
type IImageCell interface {
	ICell
	
/* debug [class_interface_properties]: Properties for ImageCell */
	// properties:
	ImageAlignment() ImageAlignment
	SetImageAlignment(value ImageAlignment)
	ImageFrameStyle() ImageFrameStyle
	SetImageFrameStyle(value ImageFrameStyle)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageCell */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageCell */
// Alloc allocates a new instance without initialization.
func (ic _ImageCellClass) Alloc() ImageCell {
	rv := objc.Send[ImageCell](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageCellClass) New() ImageCell {
	rv := objc.Send[ImageCell](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageCell) Init() ImageCell {
	rv := objc.Send[ImageCell](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageCell) Autorelease() ImageCell {
	rv := objc.Send[ImageCell](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageCell creates a new ImageCell instance.
func NewImageCell() ImageCell {
	return getImageCellClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageCell */
// An object displays a single image (encapsulated in an object) in a frame. This class provides methods for choosing the frame and for aligning and scaling the image to fit the frame.
//
// The object value of an object must be an object, so if you use the method of , be sure to supply an object as an argument. Because an object does not need to be converted for display, do not use the methods relating to formatters. An object is usually associated with some kind of control object. For example, an or an .


// An object displays a single image (encapsulated in an object) in a frame. This class provides methods for choosing the frame and for aligning and scaling the image to fit the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell
type ImageCell struct {
	Cell
}

// ImageCellFrom constructs a [ImageCell] from an unsafe.Pointer.
//
// An object displays a single image (encapsulated in an object) in a frame. This class provides methods for choosing the frame and for aligning and scaling the image to fit the frame.
func ImageCellFrom(ptr unsafe.Pointer) ImageCell {
	return ImageCell{
		Cell: CellFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageCell */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageCell */

// The alignment of the receiver’s image relative to its frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell/imageAlignment
func (i_ ImageCell) ImageAlignment() ImageAlignment {
	rv := objc.Send[ImageAlignment](i_.ID, objc.Sel("imageAlignment"))
	return rv
}/* debug [instance_properties/getter]: imageAlignment */


// The alignment of the receiver’s image relative to its frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell/imageAlignment
func (i_ ImageCell) SetImageAlignment(value ImageAlignment) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageAlignment:"), value)
}/* debug [instance_properties/setter]: imageAlignment */


// The style of the frame that borders the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell/imageFrameStyle
func (i_ ImageCell) ImageFrameStyle() ImageFrameStyle {
	rv := objc.Send[ImageFrameStyle](i_.ID, objc.Sel("imageFrameStyle"))
	return rv
}/* debug [instance_properties/getter]: imageFrameStyle */


// The style of the frame that borders the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell/imageFrameStyle
func (i_ ImageCell) SetImageFrameStyle(value ImageFrameStyle) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageFrameStyle:"), value)
}/* debug [instance_properties/setter]: imageFrameStyle */


// The scaling mode used to fit the receiver’s image into the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell/imageScaling
func (i_ ImageCell) ImageScaling() ImageScaling {
	rv := objc.Send[ImageScaling](i_.ID, objc.Sel("imageScaling"))
	return rv
}/* debug [instance_properties/getter]: imageScaling */


// The scaling mode used to fit the receiver’s image into the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageCell/imageScaling
func (i_ ImageCell) SetImageScaling(value ImageScaling) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setImageScaling:"), value)
}/* debug [instance_properties/setter]: imageScaling */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSImageCell */



