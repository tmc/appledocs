// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QCCompositionParameterView */


/* debug [class_header]: Header for QCCompositionParameterView */
// The class instance for the [QCCompositionParameterView] class.
var (
	QCCompositionParameterViewClass     _QCCompositionParameterViewClass
	QCCompositionParameterViewClassOnce sync.Once
)

func getQCCompositionParameterViewClass() _QCCompositionParameterViewClass {
	QCCompositionParameterViewClassOnce.Do(func() {
		QCCompositionParameterViewClass = _QCCompositionParameterViewClass{objc.GetClass("QCCompositionParameterView")}
	})
	return QCCompositionParameterViewClass
}

type _QCCompositionParameterViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QCCompositionParameterView */
// An interface definition for the [QCCompositionParameterView] class.
type IQCCompositionParameterView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for QCCompositionParameterView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QCCompositionParameterView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QCCompositionParameterView */
// Alloc allocates a new instance without initialization.
func (qc _QCCompositionParameterViewClass) Alloc() QCCompositionParameterView {
	rv := objc.Send[QCCompositionParameterView](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QCCompositionParameterViewClass) New() QCCompositionParameterView {
	rv := objc.Send[QCCompositionParameterView](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCCompositionParameterView) Init() QCCompositionParameterView {
	rv := objc.Send[QCCompositionParameterView](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCCompositionParameterView) Autorelease() QCCompositionParameterView {
	rv := objc.Send[QCCompositionParameterView](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCCompositionParameterView creates a new QCCompositionParameterView instance.
func NewQCCompositionParameterView() QCCompositionParameterView {
	return getQCCompositionParameterViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QCCompositionParameterView */
// A class that allows users to edit the input parameters of a composition in real time. The composition can be rendering in any of the following objects: , , or .


// A class that allows users to edit the input parameters of a composition in real time. The composition can be rendering in any of the following objects: , , or .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionParameterView
type QCCompositionParameterView struct {
	appkit.View
}

// QCCompositionParameterViewFrom constructs a [QCCompositionParameterView] from an unsafe.Pointer.
//
// A class that allows users to edit the input parameters of a composition in real time. The composition can be rendering in any of the following objects: , , or .
func QCCompositionParameterViewFrom(ptr unsafe.Pointer) QCCompositionParameterView {
	return QCCompositionParameterView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QCCompositionParameterView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QCCompositionParameterView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QCCompositionParameterView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QCCompositionParameterView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QCCompositionParameterView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QCCompositionParameterView */



