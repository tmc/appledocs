// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class QLPreviewController */


/* debug [class_header]: Header for QLPreviewController */
// The class instance for the [PreviewController] class.
var (
	PreviewControllerClass     _PreviewControllerClass
	PreviewControllerClassOnce sync.Once
)

func getPreviewControllerClass() _PreviewControllerClass {
	PreviewControllerClassOnce.Do(func() {
		PreviewControllerClass = _PreviewControllerClass{objc.GetClass("QLPreviewController")}
	})
	return PreviewControllerClass
}

type _PreviewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PreviewController */
// An interface definition for the [PreviewController] class.
type IPreviewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for PreviewController */
	// properties:
	PreviewItemTitle() objc.IObject /* cross-framework: NSString */
	SetPreviewItemTitle(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PreviewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PreviewController */
// Alloc allocates a new instance without initialization.
func (pc _PreviewControllerClass) Alloc() PreviewController {
	rv := objc.Send[PreviewController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PreviewControllerClass) New() PreviewController {
	rv := objc.Send[PreviewController](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewController) Init() PreviewController {
	rv := objc.Send[PreviewController](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewController) Autorelease() PreviewController {
	rv := objc.Send[PreviewController](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewController creates a new PreviewController instance.
func NewPreviewController() PreviewController {
	return getPreviewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PreviewController */
// A specialized view controller for previewing an item.
//
// A can display previews for many common file types, including the following: iWork documents Microsoft Office documents Rich text format, or RTF, documents PDF files Images Text files with a uniform type identifier that conforms to the type. To learn more, see . Comma-separated values, or CSV, files 3D models in the USDZ format with both standalone and AR views for viewing the model


// A specialized view controller for previewing an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewController
type PreviewController struct {
	ViewController
}

// PreviewControllerFrom constructs a [PreviewController] from an unsafe.Pointer.
//
// A specialized view controller for previewing an item.
func PreviewControllerFrom(ptr unsafe.Pointer) PreviewController {
	return PreviewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PreviewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PreviewController */

// Returns a Boolean value that indicates whether the preview controller can display an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewController/canPreview(_:)
func (pc _PreviewControllerClass) CanPreviewItem(item unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("canPreviewItem:"), item)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CanPreviewItem) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PreviewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PreviewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PreviewController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewitem/previewitemtitle
func (p_ PreviewController) PreviewItemTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("previewItemTitle"))
	return rv
}/* debug [instance_properties/getter]: previewItemTitle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quicklook/qlpreviewitem/previewitemtitle
func (p_ PreviewController) SetPreviewItemTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreviewItemTitle:"), value)
}/* debug [instance_properties/setter]: previewItemTitle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLPreviewController */


