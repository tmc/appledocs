// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class ICCameraFolder */


/* debug [class_header]: Header for ICCameraFolder */
// The class instance for the [ICCameraFolder] class.
var (
	ICCameraFolderClass     _ICCameraFolderClass
	ICCameraFolderClassOnce sync.Once
)

func getICCameraFolderClass() _ICCameraFolderClass {
	ICCameraFolderClassOnce.Do(func() {
		ICCameraFolderClass = _ICCameraFolderClass{objc.GetClass("ICCameraFolder")}
	})
	return ICCameraFolderClass
}

type _ICCameraFolderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICCameraFolder */
// An interface definition for the [ICCameraFolder] class.
type IICCameraFolder interface {
	IICCameraItem
	
/* debug [class_interface_properties]: Properties for ICCameraFolder */
	// properties:
	Contents() ICCameraItem
	SetContents(value ICCameraItem)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICCameraFolder */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICCameraFolder */
// Alloc allocates a new instance without initialization.
func (ic _ICCameraFolderClass) Alloc() ICCameraFolder {
	rv := objc.Send[ICCameraFolder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICCameraFolderClass) New() ICCameraFolder {
	rv := objc.Send[ICCameraFolder](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICCameraFolder) Init() ICCameraFolder {
	rv := objc.Send[ICCameraFolder](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICCameraFolder) Autorelease() ICCameraFolder {
	rv := objc.Send[ICCameraFolder](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICCameraFolder creates a new ICCameraFolder instance.
func NewICCameraFolder() ICCameraFolder {
	return getICCameraFolderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICCameraFolder */
// An object that represents a folder on a camera.


// An object that represents a folder on a camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraFolder
type ICCameraFolder struct {
	ICCameraItem
}

// ICCameraFolderFrom constructs a [ICCameraFolder] from an unsafe.Pointer.
//
// An object that represents a folder on a camera.
func ICCameraFolderFrom(ptr unsafe.Pointer) ICCameraFolder {
	return ICCameraFolder{
		ICCameraItem: ICCameraItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICCameraFolder *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICCameraFolder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICCameraFolder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICCameraFolder */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICCameraFolder */

// A list of items that this folder contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafolder/1389005-contents
func (i_ ICCameraFolder) Contents() ICCameraItem {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("contents"))
	return rv
}/* debug [instance_properties/getter]: contents */


// A list of items that this folder contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/iccamerafolder/1389005-contents
func (i_ ICCameraFolder) SetContents(value ICCameraItem) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setContents:"), value)
}/* debug [instance_properties/setter]: contents */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICCameraFolder */



