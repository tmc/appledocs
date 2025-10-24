// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class IKImageEditPanel */


/* debug [class_header]: Header for IKImageEditPanel */
// The class instance for the [IKImageEditPanel] class.
var (
	IKImageEditPanelClass     _IKImageEditPanelClass
	IKImageEditPanelClassOnce sync.Once
)

func getIKImageEditPanelClass() _IKImageEditPanelClass {
	IKImageEditPanelClassOnce.Do(func() {
		IKImageEditPanelClass = _IKImageEditPanelClass{objc.GetClass("IKImageEditPanel")}
	})
	return IKImageEditPanelClass
}

type _IKImageEditPanelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IKImageEditPanel */
// An interface definition for the [IKImageEditPanel] class.
type IIKImageEditPanel interface {
	appkit.IPanel
	
/* debug [class_interface_properties]: Properties for IKImageEditPanel */
	// properties:
	DataSource() unsafe.Pointer
	SetDataSource(value unsafe.Pointer)
	FilterArray() objc.IObject /* cross-framework: NSArray */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IKImageEditPanel */
	// methods:
	ReloadData()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IKImageEditPanel */
// Alloc allocates a new instance without initialization.
func (ic _IKImageEditPanelClass) Alloc() IKImageEditPanel {
	rv := objc.Send[IKImageEditPanel](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IKImageEditPanelClass) New() IKImageEditPanel {
	rv := objc.Send[IKImageEditPanel](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKImageEditPanel) Init() IKImageEditPanel {
	rv := objc.Send[IKImageEditPanel](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKImageEditPanel) Autorelease() IKImageEditPanel {
	rv := objc.Send[IKImageEditPanel](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKImageEditPanel creates a new IKImageEditPanel instance.
func NewIKImageEditPanel() IKImageEditPanel {
	return getIKImageEditPanelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IKImageEditPanel */
// The class provides a panel, that is, a utility window that floats on top of document windows, optimized for image editing.


// The class provides a panel, that is, a utility window that floats on top of document windows, optimized for image editing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageEditPanel
type IKImageEditPanel struct {
	appkit.Panel
}

// IKImageEditPanelFrom constructs a [IKImageEditPanel] from an unsafe.Pointer.
//
// The class provides a panel, that is, a utility window that floats on top of document windows, optimized for image editing.
func IKImageEditPanelFrom(ptr unsafe.Pointer) IKImageEditPanel {
	return IKImageEditPanel{
		Panel: appkit.PanelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IKImageEditPanel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IKImageEditPanel */

// Creates a shared instance of an image editing panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageEditPanel/shared()
func (ic _IKImageEditPanelClass) SharedImageEditPanel() IKImageEditPanel {
	rv := objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("sharedImageEditPanel"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedImageEditPanel) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IKImageEditPanel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IKImageEditPanel */

// Reloads the data from the data associated with an image editing panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageEditPanel/reloadData()
func (i_ IKImageEditPanel) ReloadData() {
	objc.Send[objc.ID](i_.ID, objc.Sel("reloadData"))
}/* debug [instance_methods/method]: ReloadData */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IKImageEditPanel */

// Specifies the edit panel’s dataSource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageEditPanel/dataSource
func (i_ IKImageEditPanel) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("dataSource"))
	return rv
}/* debug [instance_properties/getter]: dataSource */


// Specifies the edit panel’s dataSource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageEditPanel/dataSource
func (i_ IKImageEditPanel) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDataSource:"), value)
}/* debug [instance_properties/setter]: dataSource */


// Returns the current array of user adjustments to effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKImageEditPanel/filterArray
func (i_ IKImageEditPanel) FilterArray() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](i_.ID, objc.Sel("filterArray"))
	return rv
}/* debug [instance_properties/getter]: filterArray */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IKImageEditPanel */



