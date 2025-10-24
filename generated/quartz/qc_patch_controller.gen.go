// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class QCPatchController */


/* debug [class_header]: Header for QCPatchController */
// The class instance for the [QCPatchController] class.
var (
	QCPatchControllerClass     _QCPatchControllerClass
	QCPatchControllerClassOnce sync.Once
)

func getQCPatchControllerClass() _QCPatchControllerClass {
	QCPatchControllerClassOnce.Do(func() {
		QCPatchControllerClass = _QCPatchControllerClass{objc.GetClass("QCPatchController")}
	})
	return QCPatchControllerClass
}

type _QCPatchControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QCPatchController */
// An interface definition for the [QCPatchController] class.
type IQCPatchController interface {
	appkit.IController
	
/* debug [class_interface_properties]: Properties for QCPatchController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QCPatchController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QCPatchController */
// Alloc allocates a new instance without initialization.
func (qc _QCPatchControllerClass) Alloc() QCPatchController {
	rv := objc.Send[QCPatchController](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QCPatchControllerClass) New() QCPatchController {
	rv := objc.Send[QCPatchController](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCPatchController) Init() QCPatchController {
	rv := objc.Send[QCPatchController](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCPatchController) Autorelease() QCPatchController {
	rv := objc.Send[QCPatchController](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCPatchController creates a new QCPatchController instance.
func NewQCPatchController() QCPatchController {
	return getQCPatchControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QCPatchController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPatchController
type QCPatchController struct {
	appkit.Controller
}

// QCPatchControllerFrom constructs a [QCPatchController] from an unsafe.Pointer.
func QCPatchControllerFrom(ptr unsafe.Pointer) QCPatchController {
	return QCPatchController{
		Controller: appkit.ControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QCPatchController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QCPatchController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QCPatchController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QCPatchController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QCPatchController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QCPatchController */



