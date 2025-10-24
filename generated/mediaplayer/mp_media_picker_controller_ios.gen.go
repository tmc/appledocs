//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for MediaPickerController


// iOS-only properties

// A Boolean value specifying the default selection behavior for a media item picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/allowsPickingMultipleItems
func (m_ MediaPickerController) AllowsPickingMultipleItems() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsPickingMultipleItems"))
	return rv
}
func (m_ MediaPickerController) SetAllowsPickingMultipleItems(value bool) {
	m_.ID.Send(objc.RegisterName("setAllowsPickingMultipleItems:"), value)
}

// The delegate for a media item picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/delegate
func (m_ MediaPickerController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("delegate"))
	return rv
}
func (m_ MediaPickerController) SetDelegate(value objc.ID) {
	m_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The media types that media item picker presents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/mediaTypes
func (m_ MediaPickerController) MediaTypes() MediaType {
	rv := objc.Send[MediaType](m_.ID, objc.Sel("mediaTypes"))
	return rv
}

// A prompt, for the user, that appears above the navigation bar buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/prompt
func (m_ MediaPickerController) Prompt() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("prompt"))
	return rv
}
func (m_ MediaPickerController) SetPrompt(value objc.IObject /* cross-framework: NSString */) {
	m_.ID.Send(objc.RegisterName("setPrompt:"), value)
}

// A Boolean value specifying whether to display iCloud Media Library items for a media picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/showsCloudItems
func (m_ MediaPickerController) ShowsCloudItems() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsCloudItems"))
	return rv
}
func (m_ MediaPickerController) SetShowsCloudItems(value bool) {
	m_.ID.Send(objc.RegisterName("setShowsCloudItems:"), value)
}

// A Boolean value that specifies whether the media item picker displays protected assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/showsItemsWithProtectedAssets
func (m_ MediaPickerController) ShowsItemsWithProtectedAssets() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsItemsWithProtectedAssets"))
	return rv
}
func (m_ MediaPickerController) SetShowsItemsWithProtectedAssets(value bool) {
	m_.ID.Send(objc.RegisterName("setShowsItemsWithProtectedAssets:"), value)
}




