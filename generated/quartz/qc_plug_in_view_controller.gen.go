// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [QCPlugInViewController] class.
var (
	QCPlugInViewControllerClass     _QCPlugInViewControllerClass
	QCPlugInViewControllerClassOnce sync.Once
)

func getQCPlugInViewControllerClass() _QCPlugInViewControllerClass {
	QCPlugInViewControllerClassOnce.Do(func() {
		QCPlugInViewControllerClass = _QCPlugInViewControllerClass{objc.GetClass("QCPlugInViewController")}
	})
	return QCPlugInViewControllerClass
}

type _QCPlugInViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [QCPlugInViewController] class.
type IQCPlugInViewController interface {
	appkit.IViewController
	PlugIn() QCPlugIn
}

// The class communicates (through Cocoa bindings) between a custom patch and the view used for the internal settings of the custom patch. Only custom patches that use internal settings exposed to the user need to use the class.
//
// You access the internal settings of a custom patch through key-value coding (KVC). All the KVC keys that represent the internal settings of the custom patch must be listed in its method. The view controller for a custom patch expects the nib file class set to the class the view outlet connected to the view that contains the editing controls The controls are bound to the as the target and as the model key path, where is the KVC key for a given internal setting of the custom patch instance.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugInViewController
type QCPlugInViewController struct {
	appkit.ViewController
}

// QCPlugInViewControllerFrom constructs a [QCPlugInViewController] from an unsafe.Pointer.
//
// The class communicates (through Cocoa bindings) between a custom patch and the view used for the internal settings of the custom patch. Only custom patches that use internal settings exposed to the user need to use the class.
func QCPlugInViewControllerFrom(ptr unsafe.Pointer) QCPlugInViewController {
	return QCPlugInViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (qc _QCPlugInViewControllerClass) Alloc() QCPlugInViewController {
	rv := objc.Send[QCPlugInViewController](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QCPlugInViewControllerClass) New() QCPlugInViewController {
	rv := objc.Send[QCPlugInViewController](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QCPlugInViewController) Init() QCPlugInViewController {
	rv := objc.Send[QCPlugInViewController](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QCPlugInViewController) Autorelease() QCPlugInViewController {
	rv := objc.Send[QCPlugInViewController](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQCPlugInViewController creates a new QCPlugInViewController instance.
func NewQCPlugInViewController() QCPlugInViewController {
	return getQCPlugInViewControllerClass().New()
}




// Creates and initializes a controller for the specified object and nib file.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugInViewController/init(plugIn:viewNibName:)
func NewQCPlugInViewControllerWithPlugInViewNibName(plugIn IQCPlugIn, name appkit.string) QCPlugInViewController {
	instance := getQCPlugInViewControllerClass().Alloc()
	rv := objc.Send[QCPlugInViewController](instance.ID, objc.Sel("initWithPlugIn:viewNibName:"), plugIn, name)
	rv.Autorelease()
	return rv
}


// Returns the object associated with the view controller for the custom patch.
//
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCPlugInViewController/plugIn()
func (q_ QCPlugInViewController) PlugIn() QCPlugIn {
	rv := objc.Send[QCPlugIn](q_.ID, objc.Sel("plugIn"))
	return rv
}


