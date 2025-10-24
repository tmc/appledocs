// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QCView */


/* debug [class_header]: Header for QCView */
// The class instance for the [QCView] class.
var (
	QCViewClass     _QCViewClass
	QCViewClassOnce sync.Once
)

func getQCViewClass() _QCViewClass {
	QCViewClassOnce.Do(func() {
		QCViewClass = _QCViewClass{objc.GetClass("QCView")}
	})
	return QCViewClass
}

type _QCViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QCView */
// An interface definition for the [QCView] class.
type IQCView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for QCView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QCView */
	// methods:
	Play(sender objc.IObject)
	Start(sender objc.IObject)
	Stop(sender objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QCView */
// Alloc allocates a new instance without initialization.
func (qc _QCViewClass) Alloc() QCView {
	rv := objc.Send[QCView](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QCViewClass) New() QCView {
	rv := objc.Send[QCView](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCView) Init() QCView {
	rv := objc.Send[QCView](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCView) Autorelease() QCView {
	rv := objc.Send[QCView](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCView creates a new QCView instance.
func NewQCView() QCView {
	return getQCViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QCView */
// The class is a custom class that loads, plays, and controls Quartz Composer compositions. It is an autonomous view that is driven by an internal timer running on the main thread.
//
// The view can be set to render a composition automatically when it is placed onscreen. The view stops rendering when it is placed offscreen. When not rendering, the view is filled with the current erase color. The rendered composition automatically synchronizes to the vertical retrace of the monitor. When you archive a object, it saves the composition that’s loaded at the time the view is archived. If you want to perform custom operations while a composition is rendering such as setting input parameters or drawing OpenGL content, you need to subclass and implement the method.


// The class is a custom class that loads, plays, and controls Quartz Composer compositions. It is an autonomous view that is driven by an internal timer running on the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCView
type QCView struct {
	appkit.View
}

// QCViewFrom constructs a [QCView] from an unsafe.Pointer.
//
// The class is a custom class that loads, plays, and controls Quartz Composer compositions. It is an autonomous view that is driven by an internal timer running on the main thread.
func QCViewFrom(ptr unsafe.Pointer) QCView {
	return QCView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QCView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QCView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QCView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QCView */

// Plays or pauses a composition in a view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCView/play(_:)
func (q_ QCView) Play(sender objc.IObject) {
	objc.Send[objc.ID](q_.ID, objc.Sel("play:"), sender)
}/* debug [instance_methods/method]: Play */


// Starts rendering a composition in a view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCView/start(_:)
func (q_ QCView) Start(sender objc.IObject) {
	objc.Send[objc.ID](q_.ID, objc.Sel("start:"), sender)
}/* debug [instance_methods/method]: Start */


// Stops rendering a composition in a view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCView/stop(_:)
func (q_ QCView) Stop(sender objc.IObject) {
	objc.Send[objc.ID](q_.ID, objc.Sel("stop:"), sender)
}/* debug [instance_methods/method]: Stop */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QCView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QCView */



