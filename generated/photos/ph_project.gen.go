// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHProject] class.
var (
	PHProjectClass     _PHProjectClass
	PHProjectClassOnce sync.Once
)

func getPHProjectClass() _PHProjectClass {
	PHProjectClassOnce.Do(func() {
		PHProjectClass = _PHProjectClass{objc.GetClass("PHProject")}
	})
	return PHProjectClass
}

type _PHProjectClass struct {
	class objc.Class
}

// An interface definition for the [PHProject] class.
type IPHProject interface {
	IPHAssetCollection
	// properties:
	HasProjectPreview() bool
	SetHasProjectPreview(value bool)
	ProjectExtensionData() objc.IObject /* cross-framework: Data */
	SetProjectExtensionData(value objc.IObject /* cross-framework: Data */)
	// methods:
}

// A representation of a Photos app project extension.
//
// This class represents the project when extended from macOS Photos. Projects can have the following types: Book Calendar Card Prints Slideshow Wall decor Users create projects by selecting one or more assets, right-clicking the selection, and grouping the assets, much like an album collection. Your app treats the project as a separate entity, represented as a .


// A representation of a Photos app project extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHProject
type PHProject struct {
	PHAssetCollection
}

// PHProjectFrom constructs a [PHProject] from an unsafe.Pointer.
//
// A representation of a Photos app project extension.
func PHProjectFrom(ptr unsafe.Pointer) PHProject {
	return PHProject{
		PHAssetCollection: PHAssetCollectionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHProjectClass) Alloc() PHProject {
	rv := objc.Send[PHProject](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHProjectClass) New() PHProject {
	rv := objc.Send[PHProject](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHProject) Init() PHProject {
	rv := objc.Send[PHProject](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHProject) Autorelease() PHProject {
	rv := objc.Send[PHProject](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHProject creates a new PHProject instance.
func NewPHProject() PHProject {
	return getPHProjectClass().New()
}



// A property that indicates whether a project preview was previously set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phproject/hasprojectpreview
func (p_ PHProject) HasProjectPreview() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasProjectPreview"))
	return rv
}


// A property that indicates whether a project preview was previously set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phproject/hasprojectpreview
func (p_ PHProject) SetHasProjectPreview(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHasProjectPreview:"), value)
}


// Data associated with the project extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phproject/projectextensiondata
func (p_ PHProject) ProjectExtensionData() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("projectExtensionData"))
	return rv
}


// Data associated with the project extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phproject/projectextensiondata
func (p_ PHProject) SetProjectExtensionData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setProjectExtensionData:"), value)
}



