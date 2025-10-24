// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PIKImageEditPanelDataSource is the IKImageEditPanelDataSource protocol interface.
//
// The   protocol describes the methods that an   object uses to access the contents of its data source object.
//
// Availability:
//   - macOS 10.4+
//
// See: doc://com.apple.quartz/documentation/Quartz/IKImageEditPanelDataSource
type PIKImageEditPanelDataSource interface {
	// Required methods
	SetImageImageProperties(image ImageRef /* not a class type */, metaData objc.IObject /* cross-framework: NSDictionary */)/* debug [protocol_interface/required_method]: SetImageImageProperties */
	// Optional methods
	ThumbnailWithMaximumSize(size Size /* not a class type */) ImageRef
	HasThumbnailWithMaximumSize() bool
}

// IKImageEditPanelDataSource is a delegate implementation builder for the PIKImageEditPanelDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type IKImageEditPanelDataSource struct {
	_ThumbnailWithMaximumSize func(size Size /* not a class type */) ImageRef
	_SetImageImageProperties func(image ImageRef /* not a class type */, metaData objc.IObject /* cross-framework: NSDictionary */)
}

// SetThumbnailWithMaximumSize sets the handler for the ThumbnailWithMaximumSize delegate method.
//
// Returns a thumbnail image whose size is no larger than the specified size.
func (d *IKImageEditPanelDataSource) SetThumbnailWithMaximumSize(f func(size Size /* not a class type */) ImageRef) {
	d._ThumbnailWithMaximumSize = f
}

// SetSetImageImageProperties sets the handler for the SetImageImageProperties delegate method.
//
// Sets an image with the specified properties.
func (d *IKImageEditPanelDataSource) SetSetImageImageProperties(f func(image ImageRef /* not a class type */, metaData objc.IObject /* cross-framework: NSDictionary */)) {
	d._SetImageImageProperties = f
}

// ThumbnailWithMaximumSize implements the PIKImageEditPanelDataSource interface.
func (d *IKImageEditPanelDataSource) ThumbnailWithMaximumSize(size Size /* not a class type */) ImageRef {
	if d._ThumbnailWithMaximumSize != nil {
		return d._ThumbnailWithMaximumSize(size)
	}
	var zero ImageRef
	return zero
}

// HasThumbnailWithMaximumSize returns true if a handler for ThumbnailWithMaximumSize has been set.
func (d *IKImageEditPanelDataSource) HasThumbnailWithMaximumSize() bool {
	return d._ThumbnailWithMaximumSize != nil
}

// SetImageImageProperties implements the PIKImageEditPanelDataSource interface.
func (d *IKImageEditPanelDataSource) SetImageImageProperties(image ImageRef /* not a class type */, metaData objc.IObject /* cross-framework: NSDictionary */) {
	if d._SetImageImageProperties != nil {
		d._SetImageImageProperties(image, metaData)
	}
}

// HasSetImageImageProperties returns true if a handler for SetImageImageProperties has been set.
func (d *IKImageEditPanelDataSource) HasSetImageImageProperties() bool {
	return d._SetImageImageProperties != nil
}
