// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHProjectSection] class.
var (
	PHProjectSectionClass     _PHProjectSectionClass
	PHProjectSectionClassOnce sync.Once
)

func getPHProjectSectionClass() _PHProjectSectionClass {
	PHProjectSectionClassOnce.Do(func() {
		PHProjectSectionClass = _PHProjectSectionClass{objc.GetClass("PHProjectSection")}
	})
	return PHProjectSectionClass
}

type _PHProjectSectionClass struct {
	class objc.Class
}

// An interface definition for the [PHProjectSection] class.
type IPHProjectSection interface {
	objectivec.IObject
}

// A collection of content representing curated asset and text elements.
//
// Each project section contains at least one object, which represents a suggested curation of the content. The number of sections included in varies depending on the source used to initialize the project: There will be one cover section with a key asset element and title, as well as a section containing multiple levels of curation, mirroring the Show Summary and Show More options of the Memory in Photos. The number of sections depends on the Album size. A small Album yields a single section, but an Album with a large quantity of photos is broken down into sections based on Moments in the user’s Photo Library. The sections will match the pagination in that project; for example, a book will break down into one section per page.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSection
type PHProjectSection struct {
	objectivec.Object
}

// PHProjectSectionFrom constructs a [PHProjectSection] from an unsafe.Pointer.
//
// A collection of content representing curated asset and text elements.
func PHProjectSectionFrom(ptr unsafe.Pointer) PHProjectSection {
	return PHProjectSection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHProjectSectionClass) Alloc() PHProjectSection {
	rv := objc.Send[PHProjectSection](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHProjectSectionClass) New() PHProjectSection {
	rv := objc.Send[PHProjectSection](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHProjectSection) Init() PHProjectSection {
	rv := objc.Send[PHProjectSection](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHProjectSection) Autorelease() PHProjectSection {
	rv := objc.Send[PHProjectSection](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHProjectSection creates a new PHProjectSection instance.
func NewPHProjectSection() PHProjectSection {
	return getPHProjectSectionClass().New()
}


// An array containing PHProjectionSessionContent objects.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSection/sectionContents
func (p_ PHProjectSection) SectionContents() []PHProjectSectionContent {
	rv := objc.Send[[]PHProjectSectionContent](p_.ID, objc.Sel("sectionContents"))
	return rv
}

// The intended usage of the section: cover, content, or auxiliary.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSection/sectionType-swift.property
func (p_ PHProjectSection) SectionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("sectionType"))
	return rv
}

// The optional section title.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectSection/title
func (p_ PHProjectSection) Title() string {
	rv := objc.Send[string](p_.ID, objc.Sel("title"))
	return rv
}



