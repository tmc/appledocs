// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [GenericViewController] class.
var (
	GenericViewControllerClass     _GenericViewControllerClass
	GenericViewControllerClassOnce sync.Once
)

func getGenericViewControllerClass() _GenericViewControllerClass {
	GenericViewControllerClassOnce.Do(func() {
		GenericViewControllerClass = _GenericViewControllerClass{objc.GetClass("AUGenericViewController")}
	})
	return GenericViewControllerClass
}

type _GenericViewControllerClass struct {
	class objc.Class
}





// An interface definition for the [GenericViewController] class.
type IGenericViewController interface {
	IViewController
	

	// properties:
	AuAudioUnit() AudioUnit /* not a class type */
	SetAuAudioUnit(value AudioUnit /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (gc _GenericViewControllerClass) Alloc() GenericViewController {
	rv := objc.Send[GenericViewController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GenericViewControllerClass) New() GenericViewController {
	rv := objc.Send[GenericViewController](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GenericViewController) Init() GenericViewController {
	rv := objc.Send[GenericViewController](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GenericViewController) Autorelease() GenericViewController {
	rv := objc.Send[GenericViewController](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGenericViewController creates a new GenericViewController instance.
func NewGenericViewController() GenericViewController {
	return getGenericViewControllerClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericViewController
type GenericViewController struct {
	ViewController
}

// GenericViewControllerFrom constructs a [GenericViewController] from an unsafe.Pointer.
func GenericViewControllerFrom(ptr unsafe.Pointer) GenericViewController {
	return GenericViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericViewController/auAudioUnit
func (g_ GenericViewController) AuAudioUnit() AudioUnit /* not a class type */ {
	rv := objc.Send[AudioUnit](g_.ID, objc.Sel("auAudioUnit"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUGenericViewController/auAudioUnit
func (g_ GenericViewController) SetAuAudioUnit(value AudioUnit /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAuAudioUnit:"), value)
}








