// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [QCView] class.
type IQCView interface {
	appkit.IView
	OpenGLPixelFormat() unsafe.Pointer
}

// The class is a custom class that loads, plays, and controls Quartz Composer compositions. It is an autonomous view that is driven by an internal timer running on the main thread.
//
// The view can be set to render a composition automatically when it is placed onscreen. The view stops rendering when it is placed offscreen. When not rendering, the view is filled with the current erase color. The rendered composition automatically synchronizes to the vertical retrace of the monitor. When you archive a object, it saves the composition that’s loaded at the time the view is archived. If you want to perform custom operations while a composition is rendering such as setting input parameters or drawing OpenGL content, you need to subclass and implement the method.
//
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

// Alloc allocates a new instance without initialization.
func (qc _QCViewClass) Alloc() QCView {
	rv := objc.Send[QCView](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the OpenGL pixel format used by the view.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCView/openGLPixelFormat()
func (q_ QCView) OpenGLPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](q_.ID, objc.Sel("openGLPixelFormat"))
	return rv
}



