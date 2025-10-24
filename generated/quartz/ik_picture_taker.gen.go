// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IKPictureTaker */


/* debug [class_header]: Header for IKPictureTaker */
// The class instance for the [IKPictureTaker] class.
var (
	IKPictureTakerClass     _IKPictureTakerClass
	IKPictureTakerClassOnce sync.Once
)

func getIKPictureTakerClass() _IKPictureTakerClass {
	IKPictureTakerClassOnce.Do(func() {
		IKPictureTakerClass = _IKPictureTakerClass{objc.GetClass("IKPictureTaker")}
	})
	return IKPictureTakerClass
}

type _IKPictureTakerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IKPictureTaker */
// An interface definition for the [IKPictureTaker] class.
type IIKPictureTaker interface {
	appkit.IPanel
	
/* debug [class_interface_properties]: Properties for IKPictureTaker */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IKPictureTaker */
	// methods:
	BeginPictureTakerWithDelegateDidEndSelectorContextInfo(delegate objc.IObject, didEndSelector objc.SEL, contextInfo objectivec.IObject)
	BeginPictureTakerSheetForWindowWithDelegateDidEndSelectorContextInfo(aWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo objectivec.IObject)
	InputImage() appkit.Image
	Mirroring() bool
	OutputImage() appkit.Image
	PopUpRecentsMenuForViewWithDelegateDidEndSelectorContextInfo(aView appkit.View, delegate objc.IObject, didEndSelector objc.SEL, contextInfo objectivec.IObject)
	RunModal() int
	SetInputImage(image appkit.Image)
	SetMirroring(b bool)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IKPictureTaker */
// Alloc allocates a new instance without initialization.
func (ic _IKPictureTakerClass) Alloc() IKPictureTaker {
	rv := objc.Send[IKPictureTaker](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IKPictureTakerClass) New() IKPictureTaker {
	rv := objc.Send[IKPictureTaker](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKPictureTaker) Init() IKPictureTaker {
	rv := objc.Send[IKPictureTaker](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKPictureTaker) Autorelease() IKPictureTaker {
	rv := objc.Send[IKPictureTaker](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKPictureTaker creates a new IKPictureTaker instance.
func NewIKPictureTaker() IKPictureTaker {
	return getIKPictureTakerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IKPictureTaker */
// The class represents a panel that allows users to choose images by browsing the file system. The picture taker panel provides an Open Recent menu, supports image cropping, and supports taking snapshots from an iSight or other digital camera.


// The class represents a panel that allows users to choose images by browsing the file system. The picture taker panel provides an Open Recent menu, supports image cropping, and supports taking snapshots from an iSight or other digital camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKPictureTaker
type IKPictureTaker struct {
	appkit.Panel
}

// IKPictureTakerFrom constructs a [IKPictureTaker] from an unsafe.Pointer.
//
// The class represents a panel that allows users to choose images by browsing the file system. The picture taker panel provides an Open Recent menu, supports image cropping, and supports taking snapshots from an iSight or other digital camera.
func IKPictureTakerFrom(ptr unsafe.Pointer) IKPictureTaker {
	return IKPictureTaker{
		Panel: appkit.PanelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IKPictureTaker *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IKPictureTaker */

// Returns a shared instance, creating it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKPictureTaker/pictureTaker()
func (ic _IKPictureTakerClass) PictureTaker() IKPictureTaker {
	rv := objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("pictureTaker"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PictureTaker) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IKPictureTaker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IKPictureTaker */

// Opens a picture taker pane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKPictureTaker/begin(withDelegate:didEnd:contextInfo:)
func (i_ IKPictureTaker) BeginPictureTakerWithDelegateDidEndSelectorContextInfo(delegate objc.IObject, didEndSelector objc.SEL, contextInfo objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("beginPictureTakerWithDelegate:didEndSelector:contextInfo:"), delegate, didEndSelector, contextInfo)
}/* debug [instance_methods/method]: BeginPictureTakerWithDelegateDidEndSelectorContextInfo */


// Opens a picture taker as a sheet whose parent is the specified window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKPictureTaker/beginSheet(for:withDelegate:didEnd:contextInfo:)
func (i_ IKPictureTaker) BeginPictureTakerSheetForWindowWithDelegateDidEndSelectorContextInfo(aWindow appkit.Window, delegate objc.IObject, didEndSelector objc.SEL, contextInfo objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("beginPictureTakerSheetForWindow:withDelegate:didEndSelector:contextInfo:"), aWindow, delegate, didEndSelector, contextInfo)
}/* debug [instance_methods/method]: BeginPictureTakerSheetForWindowWithDelegateDidEndSelectorContextInfo */


// Returns the input image associated with the picture taker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKPictureTaker/inputImage()
func (i_ IKPictureTaker) InputImage() appkit.Image {
	rv := objc.Send[appkit.Image](i_.ID, objc.Sel("inputImage"))
	return rv
}/* debug [instance_methods/method]: InputImage */


// Returns whether video mirroring is enabled during snapshots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKPictureTaker/mirroring()
func (i_ IKPictureTaker) Mirroring() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("mirroring"))
	return rv
}/* debug [instance_methods/method]: Mirroring */


// Returns the edited image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKPictureTaker/outputImage()
func (i_ IKPictureTaker) OutputImage() appkit.Image {
	rv := objc.Send[appkit.Image](i_.ID, objc.Sel("outputImage"))
	return rv
}/* debug [instance_methods/method]: OutputImage */


// Displays the Open Recent popup menu associated with the picture taker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKPictureTaker/popUpRecentsMenu(for:withDelegate:didEnd:contextInfo:)
func (i_ IKPictureTaker) PopUpRecentsMenuForViewWithDelegateDidEndSelectorContextInfo(aView appkit.View, delegate objc.IObject, didEndSelector objc.SEL, contextInfo objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("popUpRecentsMenuForView:withDelegate:didEndSelector:contextInfo:"), aView, delegate, didEndSelector, contextInfo)
}/* debug [instance_methods/method]: PopUpRecentsMenuForViewWithDelegateDidEndSelectorContextInfo */


// Opens a modal picture taker dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKPictureTaker/runModal()
func (i_ IKPictureTaker) RunModal() int {
	rv := objc.Send[int](i_.ID, objc.Sel("runModal"))
	return rv
}/* debug [instance_methods/method]: RunModal */


// Set the image input for the picture taker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKPictureTaker/setInputImage(_:)
func (i_ IKPictureTaker) SetInputImage(image appkit.Image) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInputImage:"), image)
}/* debug [instance_methods/method]: SetInputImage */


// Controls whether the receiver enables video mirroring during snapshots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKPictureTaker/setMirroring(_:)
func (i_ IKPictureTaker) SetMirroring(b bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMirroring:"), b)
}/* debug [instance_methods/method]: SetMirroring */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IKPictureTaker */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IKPictureTaker */



