// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [MKLookAroundViewController] class.
var (
	MKLookAroundViewControllerClass     _MKLookAroundViewControllerClass
	MKLookAroundViewControllerClassOnce sync.Once
)

func getMKLookAroundViewControllerClass() _MKLookAroundViewControllerClass {
	MKLookAroundViewControllerClassOnce.Do(func() {
		MKLookAroundViewControllerClass = _MKLookAroundViewControllerClass{objc.GetClass("MKLookAroundViewController")}
	})
	return MKLookAroundViewControllerClass
}

type _MKLookAroundViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [MKLookAroundViewController] class.
type IMKLookAroundViewController interface {
	appkit.IViewController
}

// A class that manages the presentation and display of a LookAround view.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKLookAroundViewController
type MKLookAroundViewController struct {
	appkit.ViewController
}

// MKLookAroundViewControllerFrom constructs a [MKLookAroundViewController] from an unsafe.Pointer.
//
// A class that manages the presentation and display of a LookAround view.
func MKLookAroundViewControllerFrom(ptr unsafe.Pointer) MKLookAroundViewController {
	return MKLookAroundViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKLookAroundViewControllerClass) Alloc() MKLookAroundViewController {
	rv := objc.Send[MKLookAroundViewController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKLookAroundViewControllerClass) New() MKLookAroundViewController {
	rv := objc.Send[MKLookAroundViewController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKLookAroundViewController) Init() MKLookAroundViewController {
	rv := objc.Send[MKLookAroundViewController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKLookAroundViewController) Autorelease() MKLookAroundViewController {
	rv := objc.Send[MKLookAroundViewController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKLookAroundViewController creates a new MKLookAroundViewController instance.
func NewMKLookAroundViewController() MKLookAroundViewController {
	return getMKLookAroundViewControllerClass().New()
}




