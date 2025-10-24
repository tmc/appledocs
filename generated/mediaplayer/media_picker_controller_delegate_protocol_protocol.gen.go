// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PMediaPickerControllerDelegate is the MPMediaPickerControllerDelegate protocol interface.
//
// The protocol you implement so that a media item picker can respond to a user making media item selections.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//
// See: doc://com.apple.mediaplayer/documentation/MediaPlayer/MPMediaPickerControllerDelegate
type PMediaPickerControllerDelegate interface {
	// Optional methods
	MediaPickerDidPickMediaItems(mediaPicker IMPMediaPickerController, mediaItemCollection IMPMediaItemCollection)
	HasMediaPickerDidPickMediaItems() bool
	MediaPickerDidCancel(mediaPicker IMPMediaPickerController)
	HasMediaPickerDidCancel() bool
}

// MediaPickerControllerDelegate is a delegate implementation builder for the PMediaPickerControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MediaPickerControllerDelegate struct {
	_MediaPickerDidPickMediaItems func(mediaPicker IMPMediaPickerController, mediaItemCollection IMPMediaItemCollection)
	_MediaPickerDidCancel func(mediaPicker IMPMediaPickerController)
}

// SetMediaPickerDidPickMediaItems sets the handler for the MediaPickerDidPickMediaItems delegate method.
//
// A method that the system calls when a user selects a set of media items.
func (d *MediaPickerControllerDelegate) SetMediaPickerDidPickMediaItems(f func(mediaPicker IMPMediaPickerController, mediaItemCollection IMPMediaItemCollection)) {
	d._MediaPickerDidPickMediaItems = f
}

// SetMediaPickerDidCancel sets the handler for the MediaPickerDidCancel delegate method.
//
// A method that the system calls when a user taps Cancel to dismiss a media item picker.
func (d *MediaPickerControllerDelegate) SetMediaPickerDidCancel(f func(mediaPicker IMPMediaPickerController)) {
	d._MediaPickerDidCancel = f
}

// MediaPickerDidPickMediaItems implements the PMediaPickerControllerDelegate interface.
func (d *MediaPickerControllerDelegate) MediaPickerDidPickMediaItems(mediaPicker IMPMediaPickerController, mediaItemCollection IMPMediaItemCollection) {
	if d._MediaPickerDidPickMediaItems != nil {
		d._MediaPickerDidPickMediaItems(mediaPicker, mediaItemCollection)
	}
}

// HasMediaPickerDidPickMediaItems returns true if a handler for MediaPickerDidPickMediaItems has been set.
func (d *MediaPickerControllerDelegate) HasMediaPickerDidPickMediaItems() bool {
	return d._MediaPickerDidPickMediaItems != nil
}

// MediaPickerDidCancel implements the PMediaPickerControllerDelegate interface.
func (d *MediaPickerControllerDelegate) MediaPickerDidCancel(mediaPicker IMPMediaPickerController) {
	if d._MediaPickerDidCancel != nil {
		d._MediaPickerDidCancel(mediaPicker)
	}
}

// HasMediaPickerDidCancel returns true if a handler for MediaPickerDidCancel has been set.
func (d *MediaPickerControllerDelegate) HasMediaPickerDidCancel() bool {
	return d._MediaPickerDidCancel != nil
}
