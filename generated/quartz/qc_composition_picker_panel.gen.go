// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class QCCompositionPickerPanel */


/* debug [class_header]: Header for QCCompositionPickerPanel */
// The class instance for the [QCCompositionPickerPanel] class.
var (
	QCCompositionPickerPanelClass     _QCCompositionPickerPanelClass
	QCCompositionPickerPanelClassOnce sync.Once
)

func getQCCompositionPickerPanelClass() _QCCompositionPickerPanelClass {
	QCCompositionPickerPanelClassOnce.Do(func() {
		QCCompositionPickerPanelClass = _QCCompositionPickerPanelClass{objc.GetClass("QCCompositionPickerPanel")}
	})
	return QCCompositionPickerPanelClass
}

type _QCCompositionPickerPanelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QCCompositionPickerPanel */
// An interface definition for the [QCCompositionPickerPanel] class.
type IQCCompositionPickerPanel interface {
	appkit.IPanel
	
/* debug [class_interface_properties]: Properties for QCCompositionPickerPanel */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QCCompositionPickerPanel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QCCompositionPickerPanel */
// Alloc allocates a new instance without initialization.
func (qc _QCCompositionPickerPanelClass) Alloc() QCCompositionPickerPanel {
	rv := objc.Send[QCCompositionPickerPanel](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QCCompositionPickerPanelClass) New() QCCompositionPickerPanel {
	rv := objc.Send[QCCompositionPickerPanel](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCCompositionPickerPanel) Init() QCCompositionPickerPanel {
	rv := objc.Send[QCCompositionPickerPanel](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCCompositionPickerPanel) Autorelease() QCCompositionPickerPanel {
	rv := objc.Send[QCCompositionPickerPanel](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCCompositionPickerPanel creates a new QCCompositionPickerPanel instance.
func NewQCCompositionPickerPanel() QCCompositionPickerPanel {
	return getQCCompositionPickerPanelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QCCompositionPickerPanel */
// The class represents a utility window that allows users to browse compositions that are in the Quartz Composer composition repository and, if supported, preview the composition. The class cannot be subclassed.


// The class represents a utility window that allows users to browse compositions that are in the Quartz Composer composition repository and, if supported, preview the composition. The class cannot be subclassed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionPickerPanel
type QCCompositionPickerPanel struct {
	appkit.Panel
}

// QCCompositionPickerPanelFrom constructs a [QCCompositionPickerPanel] from an unsafe.Pointer.
//
// The class represents a utility window that allows users to browse compositions that are in the Quartz Composer composition repository and, if supported, preview the composition. The class cannot be subclassed.
func QCCompositionPickerPanelFrom(ptr unsafe.Pointer) QCCompositionPickerPanel {
	return QCCompositionPickerPanel{
		Panel: appkit.PanelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QCCompositionPickerPanel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QCCompositionPickerPanel */

// Returns the shared instance of the composition picker panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionPickerPanel/shared()
func (qc _QCCompositionPickerPanelClass) SharedCompositionPickerPanel() QCCompositionPickerPanel {
	rv := objc.Send[QCCompositionPickerPanel](objc.ID(qc.class), objc.Sel("sharedCompositionPickerPanel"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedCompositionPickerPanel) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QCCompositionPickerPanel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QCCompositionPickerPanel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QCCompositionPickerPanel */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QCCompositionPickerPanel */



