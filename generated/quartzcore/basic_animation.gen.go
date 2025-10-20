// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BasicAnimation] class.
var (
	basicAnimationClass     _BasicAnimationClass
	basicAnimationClassOnce sync.Once
)

func getBasicAnimationClass() _BasicAnimationClass {
	basicAnimationClassOnce.Do(func() {
		basicAnimationClass = _BasicAnimationClass{objc.GetClass("CABasicAnimation")}
	})
	return basicAnimationClass
}

type _BasicAnimationClass struct {
	class objc.Class
}

// An interface definition for the [BasicAnimation] class.
type IBasicAnimation interface {
	IPropertyAnimation
}

// An object that provides basic, single-keyframe animation capabilities for a layer property.
//
// You create an instance of using the inherited method, specifying the key path of the property to be animated in the render tree. For example, you can animate a layer’s scalar (i.e. containing a single value) properties such as its . The following code fades in a layer by animating its opacity from to . Non-scalar properties, such as , can also be animated. Core Animation will interpolate between the color and the color. The animation created in the following code fades a layer’s background color from red to blue. If you want to animate the individual components of a non-scalar property with different values, you pass the values to and as arrays. The following animation moves a layer from to . The can access the individual components of a property. For example, the following animation stretches a layer by animating its object’s from to .
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation
type BasicAnimation struct {
	PropertyAnimation
}

// BasicAnimationFrom constructs a [BasicAnimation] from an unsafe.Pointer.
//
// An object that provides basic, single-keyframe animation capabilities for a layer property.
func BasicAnimationFrom(ptr unsafe.Pointer) BasicAnimation {
	return BasicAnimation{
		PropertyAnimation: PropertyAnimationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BasicAnimationClass) Alloc() BasicAnimation {
	rv := objc.Send[BasicAnimation](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




