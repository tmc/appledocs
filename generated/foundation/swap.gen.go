// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [swap] class.
var (
	SwapClass     _swapClass
	SwapClassOnce sync.Once
)

func getswapClass() _swapClass {
	SwapClassOnce.Do(func() {
		SwapClass = _swapClass{objc.GetClass("swap")}
	})
	return SwapClass
}

type _swapClass struct {
	class objc.Class
}





// An interface definition for the [swap] class.
type Iswap interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _swapClass) Alloc() swap {
	rv := objc.Send[swap](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _swapClass) New() swap {
	rv := objc.Send[swap](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ swap) Init() swap {
	rv := objc.Send[swap](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ swap) Autorelease() swap {
	rv := objc.Send[swap](s_.ID, objc.Sel("autorelease"))
	return rv
}

// Newswap creates a new swap instance.
func Newswap() swap {
	return getswapClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/swap
type swap struct {
	objectivec.Object
}

// swapFrom constructs a [swap] from an unsafe.Pointer.
func swapFrom(ptr unsafe.Pointer) swap {
	return swap{objectivec.Object{objc.ID(ptr)}}
}































