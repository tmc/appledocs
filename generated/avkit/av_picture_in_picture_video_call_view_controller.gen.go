// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [PictureInPictureVideoCallViewController] class.
var (
	PictureInPictureVideoCallViewControllerClass     _PictureInPictureVideoCallViewControllerClass
	PictureInPictureVideoCallViewControllerClassOnce sync.Once
)

func getPictureInPictureVideoCallViewControllerClass() _PictureInPictureVideoCallViewControllerClass {
	PictureInPictureVideoCallViewControllerClassOnce.Do(func() {
		PictureInPictureVideoCallViewControllerClass = _PictureInPictureVideoCallViewControllerClass{objc.GetClass("AVPictureInPictureVideoCallViewController")}
	})
	return PictureInPictureVideoCallViewControllerClass
}

type _PictureInPictureVideoCallViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [PictureInPictureVideoCallViewController] class.
type IPictureInPictureVideoCallViewController interface {
	appkit.IViewController
}

// A view controller that presents content from a video call in Picture in Picture.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureVideoCallViewController
type PictureInPictureVideoCallViewController struct {
	appkit.ViewController
}

// PictureInPictureVideoCallViewControllerFrom constructs a [PictureInPictureVideoCallViewController] from an unsafe.Pointer.
//
// A view controller that presents content from a video call in Picture in Picture.
func PictureInPictureVideoCallViewControllerFrom(ptr unsafe.Pointer) PictureInPictureVideoCallViewController {
	return PictureInPictureVideoCallViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PictureInPictureVideoCallViewControllerClass) Alloc() PictureInPictureVideoCallViewController {
	rv := objc.Send[PictureInPictureVideoCallViewController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PictureInPictureVideoCallViewControllerClass) New() PictureInPictureVideoCallViewController {
	rv := objc.Send[PictureInPictureVideoCallViewController](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PictureInPictureVideoCallViewController) Init() PictureInPictureVideoCallViewController {
	rv := objc.Send[PictureInPictureVideoCallViewController](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PictureInPictureVideoCallViewController) Autorelease() PictureInPictureVideoCallViewController {
	rv := objc.Send[PictureInPictureVideoCallViewController](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPictureInPictureVideoCallViewController creates a new PictureInPictureVideoCallViewController instance.
func NewPictureInPictureVideoCallViewController() PictureInPictureVideoCallViewController {
	return getPictureInPictureVideoCallViewControllerClass().New()
}




