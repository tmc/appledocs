// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHLivePhoto] class.
var (
	PHLivePhotoClass     _PHLivePhotoClass
	PHLivePhotoClassOnce sync.Once
)

func getPHLivePhotoClass() _PHLivePhotoClass {
	PHLivePhotoClassOnce.Do(func() {
		PHLivePhotoClass = _PHLivePhotoClass{objc.GetClass("PHLivePhoto")}
	})
	return PHLivePhotoClass
}

type _PHLivePhotoClass struct {
	class objc.Class
}

// An interface definition for the [PHLivePhoto] class.
type IPHLivePhoto interface {
	objectivec.IObject
	// properties:
	Size() objc.IObject /* cross-framework: Size */
	SetSize(value objc.IObject /* cross-framework: Size */)
	// methods:
}

// A displayable representation of a Live Photo—a picture that includes motion and sound from the moments just before and after its capture.
//
// In iOS and tvOS, you can use this class to reference Live Photos from the user’s library (fetched with the and classes), to load displayable Live Photo objects from data obtained elsewhere (such as pictures shared through a social network), and to assign Live Photos to objects for display. In iOS, tvOS, and macOS, you can use this class to display edits in progress for Live Photo content in a photo editing extension. The class serves in much the same role for Live Photos as the or class serves for static images. A or object represents not the data file an image is loaded from, but instead a ready-to-use image that can be displayed in a view—similarly, a object represents a Live Photo ready to display with motion and sound using a object, not an entry in the Photos library or the data resources that constitute a Live Photo. (To work with Live Photos as elements of the Photos library, use the class. To work with the data files that constitute a Live Photo, use the class.)

// A displayable representation of a Live Photo—a picture that includes motion and sound from the moments just before and after its capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhoto
type PHLivePhoto struct {
	objectivec.Object
}

// PHLivePhotoFrom constructs a [PHLivePhoto] from an unsafe.Pointer.
//
// A displayable representation of a Live Photo—a picture that includes motion and sound from the moments just before and after its capture.
func PHLivePhotoFrom(ptr unsafe.Pointer) PHLivePhoto {
	return PHLivePhoto{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHLivePhotoClass) Alloc() PHLivePhoto {
	rv := objc.Send[PHLivePhoto](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHLivePhotoClass) New() PHLivePhoto {
	rv := objc.Send[PHLivePhoto](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHLivePhoto) Init() PHLivePhoto {
	rv := objc.Send[PHLivePhoto](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHLivePhoto) Autorelease() PHLivePhoto {
	rv := objc.Send[PHLivePhoto](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHLivePhoto creates a new PHLivePhoto instance.
func NewPHLivePhoto() PHLivePhoto {
	return getPHLivePhotoClass().New()
}

// The size, in pixels, of the Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephoto/size
func (p_ PHLivePhoto) Size() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](p_.ID, objc.Sel("size"))
	return rv
}

// The size, in pixels, of the Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephoto/size
func (p_ PHLivePhoto) SetSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSize:"), value)
}
