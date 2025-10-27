// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PImageDelegate is the NSImageDelegate protocol interface.
//
// A set of optional methods that you can use to respond to drawing failures and manage incremental loads.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSImageDelegate
type PImageDelegate interface {
	// Required methods
	ImageDidLoadPartOfRepresentationWithValidRows(image IImage, rep IImageRep, rows int)
	ImageDidLoadRepresentationWithStatus(image IImage, rep IImageRep, status ImageLoadStatus)
	ImageDidLoadRepresentationHeader(image IImage, rep IImageRep)
	ImageWillLoadRepresentation(image IImage, rep IImageRep)
	// Optional methods
	ImageDidNotDrawInRect(sender IImage, rect corefoundation.CGRect) IImage
	HasImageDidNotDrawInRect() bool
}

// ImageDelegate is a delegate implementation builder for the PImageDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ImageDelegate struct {
	_ImageDidNotDrawInRect func(sender IImage, rect corefoundation.CGRect) IImage
	_ImageDidLoadPartOfRepresentationWithValidRows func(image IImage, rep IImageRep, rows int)
	_ImageDidLoadRepresentationWithStatus func(image IImage, rep IImageRep, status ImageLoadStatus)
	_ImageDidLoadRepresentationHeader func(image IImage, rep IImageRep)
	_ImageWillLoadRepresentation func(image IImage, rep IImageRep)
}

// SetImageDidNotDrawInRect sets the handler for the ImageDidNotDrawInRect delegate method.
//
// Tells the delegate that the image object is unable, for whatever reason, to lock focus on its image or draw in the specified rectangle.
func (d *ImageDelegate) SetImageDidNotDrawInRect(f func(sender IImage, rect corefoundation.CGRect) IImage) {
	d._ImageDidNotDrawInRect = f
}

// SetImageDidLoadPartOfRepresentationWithValidRows sets the handler for the ImageDidLoadPartOfRepresentationWithValidRows delegate method.
//
// Tells the delegate that part of the incrementally loaded image data is available.
func (d *ImageDelegate) SetImageDidLoadPartOfRepresentationWithValidRows(f func(image IImage, rep IImageRep, rows int)) {
	d._ImageDidLoadPartOfRepresentationWithValidRows = f
}

// SetImageDidLoadRepresentationWithStatus sets the handler for the ImageDidLoadRepresentationWithStatus delegate method.
//
// Tells the delegate that an incrementally loaded image has been loaded and decompressed as completely as is possible.
func (d *ImageDelegate) SetImageDidLoadRepresentationWithStatus(f func(image IImage, rep IImageRep, status ImageLoadStatus)) {
	d._ImageDidLoadRepresentationWithStatus = f
}

// SetImageDidLoadRepresentationHeader sets the handler for the ImageDidLoadRepresentationHeader delegate method.
//
// Tells the delegate that enough data has been read to determine the size of the image.
func (d *ImageDelegate) SetImageDidLoadRepresentationHeader(f func(image IImage, rep IImageRep)) {
	d._ImageDidLoadRepresentationHeader = f
}

// SetImageWillLoadRepresentation sets the handler for the ImageWillLoadRepresentation delegate method.
//
// Tells the delegate that the image object is about to access its underlying bitmap data.
func (d *ImageDelegate) SetImageWillLoadRepresentation(f func(image IImage, rep IImageRep)) {
	d._ImageWillLoadRepresentation = f
}

// ImageDidNotDrawInRect implements the PImageDelegate interface.
func (d *ImageDelegate) ImageDidNotDrawInRect(sender IImage, rect corefoundation.CGRect) IImage {
	if d._ImageDidNotDrawInRect != nil {
		return d._ImageDidNotDrawInRect(sender, rect)
	}
	var zero IImage
	return zero
}

// HasImageDidNotDrawInRect returns true if a handler for ImageDidNotDrawInRect has been set.
func (d *ImageDelegate) HasImageDidNotDrawInRect() bool {
	return d._ImageDidNotDrawInRect != nil
}

// ImageDidLoadPartOfRepresentationWithValidRows implements the PImageDelegate interface.
func (d *ImageDelegate) ImageDidLoadPartOfRepresentationWithValidRows(image IImage, rep IImageRep, rows int) {
	if d._ImageDidLoadPartOfRepresentationWithValidRows != nil {
		d._ImageDidLoadPartOfRepresentationWithValidRows(image, rep, rows)
	}
}

// HasImageDidLoadPartOfRepresentationWithValidRows returns true if a handler for ImageDidLoadPartOfRepresentationWithValidRows has been set.
func (d *ImageDelegate) HasImageDidLoadPartOfRepresentationWithValidRows() bool {
	return d._ImageDidLoadPartOfRepresentationWithValidRows != nil
}

// ImageDidLoadRepresentationWithStatus implements the PImageDelegate interface.
func (d *ImageDelegate) ImageDidLoadRepresentationWithStatus(image IImage, rep IImageRep, status ImageLoadStatus) {
	if d._ImageDidLoadRepresentationWithStatus != nil {
		d._ImageDidLoadRepresentationWithStatus(image, rep, status)
	}
}

// HasImageDidLoadRepresentationWithStatus returns true if a handler for ImageDidLoadRepresentationWithStatus has been set.
func (d *ImageDelegate) HasImageDidLoadRepresentationWithStatus() bool {
	return d._ImageDidLoadRepresentationWithStatus != nil
}

// ImageDidLoadRepresentationHeader implements the PImageDelegate interface.
func (d *ImageDelegate) ImageDidLoadRepresentationHeader(image IImage, rep IImageRep) {
	if d._ImageDidLoadRepresentationHeader != nil {
		d._ImageDidLoadRepresentationHeader(image, rep)
	}
}

// HasImageDidLoadRepresentationHeader returns true if a handler for ImageDidLoadRepresentationHeader has been set.
func (d *ImageDelegate) HasImageDidLoadRepresentationHeader() bool {
	return d._ImageDidLoadRepresentationHeader != nil
}

// ImageWillLoadRepresentation implements the PImageDelegate interface.
func (d *ImageDelegate) ImageWillLoadRepresentation(image IImage, rep IImageRep) {
	if d._ImageWillLoadRepresentation != nil {
		d._ImageWillLoadRepresentation(image, rep)
	}
}

// HasImageWillLoadRepresentation returns true if a handler for ImageWillLoadRepresentation has been set.
func (d *ImageDelegate) HasImageWillLoadRepresentation() bool {
	return d._ImageWillLoadRepresentation != nil
}

// ImageDelegateObject wraps an existing Objective-C object that conforms to the PImageDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type ImageDelegateObject struct {
	objectivec.Object
}

// NewImageDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSImageDelegate protocol.
func NewImageDelegateObject(obj objectivec.Object) *ImageDelegateObject {
	return &ImageDelegateObject{obj}
}

// Make sure ImageDelegateObject implements PImageDelegate.
var _ PImageDelegate = (*ImageDelegateObject)(nil)

// ImageDidLoadPartOfRepresentationWithValidRows implements the PImageDelegate interface.
// This required method is always available on objects conforming to ImageDidLoadPartOfRepresentationWithValidRows.
func (o *ImageDelegateObject) ImageDidLoadPartOfRepresentationWithValidRows(image IImage, rep IImageRep, rows int) {
	objc.Send[objc.ID](o.ID, objc.Sel("image:didLoadPartOfRepresentation:withValidRows:"), image, rep, rows)
}

// ImageDidLoadRepresentationWithStatus implements the PImageDelegate interface.
// This required method is always available on objects conforming to ImageDidLoadRepresentationWithStatus.
func (o *ImageDelegateObject) ImageDidLoadRepresentationWithStatus(image IImage, rep IImageRep, status ImageLoadStatus) {
	objc.Send[objc.ID](o.ID, objc.Sel("image:didLoadRepresentation:withStatus:"), image, rep, status)
}

// ImageDidLoadRepresentationHeader implements the PImageDelegate interface.
// This required method is always available on objects conforming to ImageDidLoadRepresentationHeader.
func (o *ImageDelegateObject) ImageDidLoadRepresentationHeader(image IImage, rep IImageRep) {
	objc.Send[objc.ID](o.ID, objc.Sel("image:didLoadRepresentationHeader:"), image, rep)
}

// ImageWillLoadRepresentation implements the PImageDelegate interface.
// This required method is always available on objects conforming to ImageWillLoadRepresentation.
func (o *ImageDelegateObject) ImageWillLoadRepresentation(image IImage, rep IImageRep) {
	objc.Send[objc.ID](o.ID, objc.Sel("image:willLoadRepresentation:"), image, rep)
}

// ImageDidNotDrawInRect implements the PImageDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *ImageDelegateObject) ImageDidNotDrawInRect(sender IImage, rect corefoundation.CGRect) IImage {
	return objc.Send[IImage](o.ID, objc.Sel("imageDidNotDraw:inRect:"), sender, rect)
}

// HasImageDidNotDrawInRect returns true; this is a placeholder for optional method checks.
func (o *ImageDelegateObject) HasImageDidNotDrawInRect() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
