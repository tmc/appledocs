// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [QCPatchController] class.
type IQCPatchController interface {
	appkit.IController
	// properties:
	// methods:
}



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

// Alloc allocates a new instance without initialization.
func (qc _QCPatchControllerClass) Alloc() QCPatchController {
	rv := objc.Send[QCPatchController](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




