// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [lastLabel] class.
var (
	LastLabelClass     _lastLabelClass
	LastLabelClassOnce sync.Once
)

func getlastLabelClass() _lastLabelClass {
	LastLabelClassOnce.Do(func() {
		LastLabelClass = _lastLabelClass{objc.GetClass("lastLabel")}
	})
	return LastLabelClass
}

type _lastLabelClass struct {
	class objc.Class
}





// An interface definition for the [lastLabel] class.
type IlastLabel interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (lc _lastLabelClass) Alloc() lastLabel {
	rv := objc.Send[lastLabel](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _lastLabelClass) New() lastLabel {
	rv := objc.Send[lastLabel](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ lastLabel) Init() lastLabel {
	rv := objc.Send[lastLabel](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ lastLabel) Autorelease() lastLabel {
	rv := objc.Send[lastLabel](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewlastLabel creates a new lastLabel instance.
func NewlastLabel() lastLabel {
	return getlastLabelClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/lastLabel
type lastLabel struct {
	objectivec.Object
}

// lastLabelFrom constructs a [lastLabel] from an unsafe.Pointer.
func lastLabelFrom(ptr unsafe.Pointer) lastLabel {
	return lastLabel{objectivec.Object{objc.ID(ptr)}}
}































