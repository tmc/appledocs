// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHContentEditingOutput] class.
var (
	PHContentEditingOutputClass     _PHContentEditingOutputClass
	PHContentEditingOutputClassOnce sync.Once
)

func getPHContentEditingOutputClass() _PHContentEditingOutputClass {
	PHContentEditingOutputClassOnce.Do(func() {
		PHContentEditingOutputClass = _PHContentEditingOutputClass{objc.GetClass("PHContentEditingOutput")}
	})
	return PHContentEditingOutputClass
}

type _PHContentEditingOutputClass struct {
	class objc.Class
}

// An interface definition for the [PHContentEditingOutput] class.
type IPHContentEditingOutput interface {
	objectivec.IObject
}

// A container to which you provide the results of editing the photo, video, or Live Photo content of a Photos asset.
//
// To edit an asset’s photo or video content: Fetch a object that represents the photo or video to be edited. Call the asset’s method to retrieve a object. This object provides information about the asset, the asset data to be edited, and a preview image for display. Apply your edits to the asset. To allow a user to continue working with the edit later (for example, to adjust the parameters of a photo filter), create a object describing the changes. Initialize a object. For photo- or video-only assets, provide the edited content with the property. For Live Photo assets, create a object to edit the Live Photo content and pass your content editing output to the method. For all asset types, provide your adjustment data with the property of the content editing output. 5. Use a photo library change block to commit the edit. (For details, see .) In the block, create a object and set its property to the editing output that you created. Each call prompts the user for permission to edit the contents of the photo library—to edit multiple assets in one batch, create multiple objects within the same change block, each with its own corresponding object. You can also edit assets from photo editing extensions. In this case, instead of working with a object, you implement methods in the protocol. Photos provides a object when your extension begins editing. When editing is complete, Photos requests a object that contains the edited asset content.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingOutput
type PHContentEditingOutput struct {
	objectivec.Object
}

// PHContentEditingOutputFrom constructs a [PHContentEditingOutput] from an unsafe.Pointer.
//
// A container to which you provide the results of editing the photo, video, or Live Photo content of a Photos asset.
func PHContentEditingOutputFrom(ptr unsafe.Pointer) PHContentEditingOutput {
	return PHContentEditingOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHContentEditingOutputClass) Alloc() PHContentEditingOutput {
	rv := objc.Send[PHContentEditingOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHContentEditingOutputClass) New() PHContentEditingOutput {
	rv := objc.Send[PHContentEditingOutput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHContentEditingOutput) Init() PHContentEditingOutput {
	rv := objc.Send[PHContentEditingOutput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHContentEditingOutput) Autorelease() PHContentEditingOutput {
	rv := objc.Send[PHContentEditingOutput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHContentEditingOutput creates a new PHContentEditingOutput instance.
func NewPHContentEditingOutput() PHContentEditingOutput {
	return getPHContentEditingOutputClass().New()
}


// An object describing the changes made to the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingOutput/adjustmentData
func (p_ PHContentEditingOutput) AdjustmentData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("adjustmentData"))
	return rv
}


// SetAdjustmentData sets the value of the adjustmentData property.
// An object describing the changes made to the asset.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingOutput/adjustmentData
func (p_ PHContentEditingOutput) SetAdjustmentData(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAdjustmentData:"), value)
}

// The URL at which to write a file containing edited asset content.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingOutput/renderedContentURL
func (p_ PHContentEditingOutput) RenderedContentURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("renderedContentURL"))
	return rv
}



