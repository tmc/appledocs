// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHProjectSectionContent] class.
var (
	PHProjectSectionContentClass     _PHProjectSectionContentClass
	PHProjectSectionContentClassOnce sync.Once
)

func getPHProjectSectionContentClass() _PHProjectSectionContentClass {
	PHProjectSectionContentClassOnce.Do(func() {
		PHProjectSectionContentClass = _PHProjectSectionContentClass{objc.GetClass("PHProjectSectionContent")}
	})
	return PHProjectSectionContentClass
}

type _PHProjectSectionContentClass struct {
	class objc.Class
}

// An interface definition for the [PHProjectSectionContent] class.
type IPHProjectSectionContent interface {
	objectivec.IObject
}

// An object containing section elements and layout information for a single level of curation.
//
// A section content object contains suggested layout information for every element at a specific level of curation within a . A single section can provide multiple content objects, but only one is used at a time, depending on the level of curation and the amount of content detail.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSectionContent
type PHProjectSectionContent struct {
	objectivec.Object
}

// PHProjectSectionContentFrom constructs a [PHProjectSectionContent] from an unsafe.Pointer.
//
// An object containing section elements and layout information for a single level of curation.
func PHProjectSectionContentFrom(ptr unsafe.Pointer) PHProjectSectionContent {
	return PHProjectSectionContent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHProjectSectionContentClass) Alloc() PHProjectSectionContent {
	rv := objc.Send[PHProjectSectionContent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHProjectSectionContentClass) New() PHProjectSectionContent {
	rv := objc.Send[PHProjectSectionContent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHProjectSectionContent) Init() PHProjectSectionContent {
	rv := objc.Send[PHProjectSectionContent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHProjectSectionContent) Autorelease() PHProjectSectionContent {
	rv := objc.Send[PHProjectSectionContent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHProjectSectionContent creates a new PHProjectSectionContent instance.
func NewPHProjectSectionContent() PHProjectSectionContent {
	return getPHProjectSectionContentClass().New()
}


// The aspect ratio of the full content layout, defined as width over height.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSectionContent/aspectRatio
func (p_ PHProjectSectionContent) AspectRatio() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("aspectRatio"))
	return rv
}

// The background color of the section content when created from an Apple Print Product.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSectionContent/backgroundColor
func (p_ PHProjectSectionContent) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("backgroundColor"))
	return rv
}

// An array containing all cloud asset identifiers referenced in the content.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSectionContent/cloudAssetIdentifiers
func (p_ PHProjectSectionContent) CloudAssetIdentifiers() []PHCloudIdentifier {
	rv := objc.Send[[]PHCloudIdentifier](p_.ID, objc.Sel("cloudAssetIdentifiers"))
	return rv
}

// An array of asset, text, or journal entry elements contained in the content.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSectionContent/elements
func (p_ PHProjectSectionContent) Elements() []PHProjectElement {
	rv := objc.Send[[]PHProjectElement](p_.ID, objc.Sel("elements"))
	return rv
}

// The number of columns if section content is displayed in a grid layout.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSectionContent/numberOfColumns
func (p_ PHProjectSectionContent) NumberOfColumns() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfColumns"))
	return rv
}



