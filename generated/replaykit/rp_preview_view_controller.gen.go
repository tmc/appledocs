// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [RPPreviewViewController] class.
var (
	RPPreviewViewControllerClass     _RPPreviewViewControllerClass
	RPPreviewViewControllerClassOnce sync.Once
)

func getRPPreviewViewControllerClass() _RPPreviewViewControllerClass {
	RPPreviewViewControllerClassOnce.Do(func() {
		RPPreviewViewControllerClass = _RPPreviewViewControllerClass{objc.GetClass("RPPreviewViewController")}
	})
	return RPPreviewViewControllerClass
}

type _RPPreviewViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [RPPreviewViewController] class.
type IRPPreviewViewController interface {
	appkit.IViewController
}

// An object that displays a user interface where users preview and edit a screen recording that you create with ReplayKit.
//
// Upon completion of a successful recording, the preview view controller is passed into the completion handler for .
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPPreviewViewController
type RPPreviewViewController struct {
	appkit.ViewController
}

// RPPreviewViewControllerFrom constructs a [RPPreviewViewController] from an unsafe.Pointer.
//
// An object that displays a user interface where users preview and edit a screen recording that you create with ReplayKit.
func RPPreviewViewControllerFrom(ptr unsafe.Pointer) RPPreviewViewController {
	return RPPreviewViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RPPreviewViewControllerClass) Alloc() RPPreviewViewController {
	rv := objc.Send[RPPreviewViewController](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RPPreviewViewControllerClass) New() RPPreviewViewController {
	rv := objc.Send[RPPreviewViewController](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPPreviewViewController) Init() RPPreviewViewController {
	rv := objc.Send[RPPreviewViewController](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPPreviewViewController) Autorelease() RPPreviewViewController {
	rv := objc.Send[RPPreviewViewController](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPPreviewViewController creates a new RPPreviewViewController instance.
func NewRPPreviewViewController() RPPreviewViewController {
	return getRPPreviewViewControllerClass().New()
}


// The type of screen that appears when the view is presented.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPPreviewViewController/mode
func (r_ RPPreviewViewController) Mode() RPPreviewViewControllerMode {
	rv := objc.Send[RPPreviewViewControllerMode](r_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
// The type of screen that appears when the view is presented.

//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPPreviewViewController/mode
func (r_ RPPreviewViewController) SetMode(value RPPreviewViewControllerMode) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMode:"), value)
}

// The preview view controller’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPPreviewViewController/previewControllerDelegate
func (r_ RPPreviewViewController) PreviewControllerDelegate() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("previewControllerDelegate"))
	return rv
}


// SetPreviewControllerDelegate sets the value of the previewControllerDelegate property.
// The preview view controller’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPPreviewViewController/previewControllerDelegate
func (r_ RPPreviewViewController) SetPreviewControllerDelegate(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPreviewControllerDelegate:"), value)
}



