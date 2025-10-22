// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHProjectInfo] class.
var (
	PHProjectInfoClass     _PHProjectInfoClass
	PHProjectInfoClassOnce sync.Once
)

func getPHProjectInfoClass() _PHProjectInfoClass {
	PHProjectInfoClassOnce.Do(func() {
		PHProjectInfoClass = _PHProjectInfoClass{objc.GetClass("PHProjectInfo")}
	})
	return PHProjectInfoClass
}

type _PHProjectInfoClass struct {
	class objc.Class
}

// An interface definition for the [PHProjectInfo] class.
type IPHProjectInfo interface {
	objectivec.IObject
	BrandingEnabled() bool
	SetBrandingEnabled(value bool)
	CreationSource() unsafe.Pointer
	SetCreationSource(value unsafe.Pointer)
	PageNumbersEnabled() bool
	SetPageNumbersEnabled(value bool)
	ProductIdentifier() string
	SetProductIdentifier(value string)
	ProjectType() unsafe.Pointer
	SetProjectType(value unsafe.Pointer)
	Sections() PHProjectSection
	SetSections(value IPHProjectSection)
	ThemeIdentifier() string
	SetThemeIdentifier(value string)
}

// Information about the project extension.
//
// macOS Photos automatically generates a object when creating a new project. Photos passes along the project information with a object. This object contains metadata about the project’s creation source, sections, product type, branding, and page numbers. Your extension leverages project information to influence project layout, autoflow, and theme selection. The properties of this class are immutable, and your extension can’t instantiate the object directly.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectInfo
type PHProjectInfo struct {
	objectivec.Object
}

// PHProjectInfoFrom constructs a [PHProjectInfo] from an unsafe.Pointer.
//
// Information about the project extension.
func PHProjectInfoFrom(ptr unsafe.Pointer) PHProjectInfo {
	return PHProjectInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHProjectInfoClass) Alloc() PHProjectInfo {
	rv := objc.Send[PHProjectInfo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHProjectInfoClass) New() PHProjectInfo {
	rv := objc.Send[PHProjectInfo](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHProjectInfo) Init() PHProjectInfo {
	rv := objc.Send[PHProjectInfo](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHProjectInfo) Autorelease() PHProjectInfo {
	rv := objc.Send[PHProjectInfo](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHProjectInfo creates a new PHProjectInfo instance.
func NewPHProjectInfo() PHProjectInfo {
	return getPHProjectInfoClass().New()
}


// A Boolean value indicating whether branding was enabled in the source project.
//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/brandingenabled
func (p_ PHProjectInfo) BrandingEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("brandingEnabled"))
	return rv
}


// SetBrandingEnabled sets the value of the brandingEnabled property.
// A Boolean value indicating whether branding was enabled in the source project.

//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/brandingenabled
func (p_ PHProjectInfo) SetBrandingEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBrandingEnabled:"), value)
}

// The source from which the project was created.
//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/creationsource-swift.property
func (p_ PHProjectInfo) CreationSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("creationSource"))
	return rv
}


// SetCreationSource sets the value of the creationSource property.
// The source from which the project was created.

//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/creationsource-swift.property
func (p_ PHProjectInfo) SetCreationSource(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCreationSource:"), value)
}

// A Boolean value indicating whether page numbering was enabled in the source project.
//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/pagenumbersenabled
func (p_ PHProjectInfo) PageNumbersEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("pageNumbersEnabled"))
	return rv
}


// SetPageNumbersEnabled sets the value of the pageNumbersEnabled property.
// A Boolean value indicating whether page numbering was enabled in the source project.

//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/pagenumbersenabled
func (p_ PHProjectInfo) SetPageNumbersEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageNumbersEnabled:"), value)
}

// The product identifier of the originating Apple Print Product.
//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/productidentifier
func (p_ PHProjectInfo) ProductIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("productIdentifier"))
	return rv
}


// SetProductIdentifier sets the value of the productIdentifier property.
// The product identifier of the originating Apple Print Product.

//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/productidentifier
func (p_ PHProjectInfo) SetProductIdentifier(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProductIdentifier:"), objc.String(value))
}

// The project type that the user selected from the project extension options.
//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/projecttype
func (p_ PHProjectInfo) ProjectType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("projectType"))
	return rv
}


// SetProjectType sets the value of the projectType property.
// The project type that the user selected from the project extension options.

//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/projecttype
func (p_ PHProjectInfo) SetProjectType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProjectType:"), value)
}

// An array of project sections, each containing one or more section content objects.
//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/sections
func (p_ PHProjectInfo) Sections() PHProjectSection {
	rv := objc.Send[PHProjectSection](p_.ID, objc.Sel("sections"))
	return rv
}


// SetSections sets the value of the sections property.
// An array of project sections, each containing one or more section content objects.

//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/sections
func (p_ PHProjectInfo) SetSections(value IPHProjectSection) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSections:"), value)
}

// The product theme identifier of the originating Apple Print Product.
//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/themeidentifier
func (p_ PHProjectInfo) ThemeIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("themeIdentifier"))
	return rv
}


// SetThemeIdentifier sets the value of the themeIdentifier property.
// The product theme identifier of the originating Apple Print Product.

//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectinfo/themeidentifier
func (p_ PHProjectInfo) SetThemeIdentifier(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setThemeIdentifier:"), objc.String(value))
}



