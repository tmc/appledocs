// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [BasicAnimation] class.
var (
	BasicAnimationClass     _BasicAnimationClass
	BasicAnimationClassOnce sync.Once
)

func getBasicAnimationClass() _BasicAnimationClass {
	BasicAnimationClassOnce.Do(func() {
		BasicAnimationClass = _BasicAnimationClass{objc.GetClass("CABasicAnimation")}
	})
	return BasicAnimationClass
}

type _BasicAnimationClass struct {
	class objc.Class
}





// An interface definition for the [BasicAnimation] class.
type IBasicAnimation interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (bc _BasicAnimationClass) Alloc() BasicAnimation {
	rv := objc.Send[BasicAnimation](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BasicAnimationClass) New() BasicAnimation {
	rv := objc.Send[BasicAnimation](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BasicAnimation) Init() BasicAnimation {
	rv := objc.Send[BasicAnimation](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BasicAnimation) Autorelease() BasicAnimation {
	rv := objc.Send[BasicAnimation](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBasicAnimation creates a new BasicAnimation instance.
func NewBasicAnimation() BasicAnimation {
	return getBasicAnimationClass().New()
}





// A parent class referenced by other QuartzCore classes.


// A parent class referenced by other QuartzCore classes. [Full Topic]
type BasicAnimation struct {
	objectivec.Object
}

// BasicAnimationFrom constructs a [BasicAnimation] from an unsafe.Pointer.
//
// A parent class referenced by other QuartzCore classes.
func BasicAnimationFrom(ptr unsafe.Pointer) BasicAnimation {
	return BasicAnimation{objectivec.Object{objc.ID(ptr)}}
}































