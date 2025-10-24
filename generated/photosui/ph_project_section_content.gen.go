// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/photos"
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
	AspectRatio() float64
	BackgroundColor() appkit.Color
	CloudAssetIdentifiers() []photos.PHCloudIdentifier
	Elements() []PHProjectElement
	NumberOfColumns() int
	SectionContents() PHProjectSectionContent
	SetSectionContents(value IPHProjectSectionContent)
	Title() string
	SetTitle(value string)
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
func (p_ PHProjectSectionContent) AspectRatio() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("aspectRatio"))
	return rv
}

// The background color of the section content when created from an Apple Print Product.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSectionContent/backgroundColor
func (p_ PHProjectSectionContent) BackgroundColor() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("backgroundColor"))
	return rv
}

// An array containing all cloud asset identifiers referenced in the content.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSectionContent/cloudAssetIdentifiers
func (p_ PHProjectSectionContent) CloudAssetIdentifiers() []photos.PHCloudIdentifier {
	rv := objc.Send[[]photos.PHCloudIdentifier](p_.ID, objc.Sel("cloudAssetIdentifiers"))
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

// An array containing PHProjectionSessionContent objects.
//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectsection/sectioncontents
func (p_ PHProjectSectionContent) SectionContents() PHProjectSectionContent {
	rv := objc.Send[PHProjectSectionContent](p_.ID, objc.Sel("sectionContents"))
	return rv
}

// SetSectionContents sets the value of the sectionContents property.
// An array containing PHProjectionSessionContent objects.

//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectsection/sectioncontents
func (p_ PHProjectSectionContent) SetSectionContents(value IPHProjectSectionContent) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSectionContents:"), value)
}

// The optional section title.
//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectsection/title
func (p_ PHProjectSectionContent) Title() string {
	rv := objc.Send[string](p_.ID, objc.Sel("title"))
	return rv
}

// SetTitle sets the value of the title property.
// The optional section title.

//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectsection/title
func (p_ PHProjectSectionContent) SetTitle(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), objc.String(value))
}
