// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSScrollEdgeEffectStyle */


/* debug [class_header]: Header for NSScrollEdgeEffectStyle */
// The class instance for the [ScrollEdgeEffectStyle] class.
var (
	ScrollEdgeEffectStyleClass     _ScrollEdgeEffectStyleClass
	ScrollEdgeEffectStyleClassOnce sync.Once
)

func getScrollEdgeEffectStyleClass() _ScrollEdgeEffectStyleClass {
	ScrollEdgeEffectStyleClassOnce.Do(func() {
		ScrollEdgeEffectStyleClass = _ScrollEdgeEffectStyleClass{objc.GetClass("NSScrollEdgeEffectStyle")}
	})
	return ScrollEdgeEffectStyleClass
}

type _ScrollEdgeEffectStyleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScrollEdgeEffectStyle */
// An interface definition for the [ScrollEdgeEffectStyle] class.
type IScrollEdgeEffectStyle interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScrollEdgeEffectStyle */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScrollEdgeEffectStyle */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScrollEdgeEffectStyle */
// Alloc allocates a new instance without initialization.
func (sc _ScrollEdgeEffectStyleClass) Alloc() ScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrollEdgeEffectStyleClass) New() ScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrollEdgeEffectStyle) Init() ScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrollEdgeEffectStyle) Autorelease() ScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrollEdgeEffectStyle creates a new ScrollEdgeEffectStyle instance.
func NewScrollEdgeEffectStyle() ScrollEdgeEffectStyle {
	return getScrollEdgeEffectStyleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScrollEdgeEffectStyle */
// Styles for a scroll view’s edge effect.


// Styles for a scroll view’s edge effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollEdgeEffectStyle
type ScrollEdgeEffectStyle struct {
	objectivec.Object
}

// ScrollEdgeEffectStyleFrom constructs a [ScrollEdgeEffectStyle] from an unsafe.Pointer.
//
// Styles for a scroll view’s edge effect.
func ScrollEdgeEffectStyleFrom(ptr unsafe.Pointer) ScrollEdgeEffectStyle {
	return ScrollEdgeEffectStyle{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScrollEdgeEffectStyle *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScrollEdgeEffectStyle */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScrollEdgeEffectStyle */

// The automatic scroll edge effect style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollEdgeEffectStyle/automatic
func (sc _ScrollEdgeEffectStyleClass) AutomaticStyle() ScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](objc.ID(sc.class), objc.Sel("automaticStyle"))
	return rv
}/* debug [class_properties_class/property]: automaticStyle */

// A scroll edge effect with a hard cutoff.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollEdgeEffectStyle/hard
func (sc _ScrollEdgeEffectStyleClass) HardStyle() ScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](objc.ID(sc.class), objc.Sel("hardStyle"))
	return rv
}/* debug [class_properties_class/property]: hardStyle */

// A scroll edge effect with a soft edge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollEdgeEffectStyle/soft
func (sc _ScrollEdgeEffectStyleClass) SoftStyle() ScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](objc.ID(sc.class), objc.Sel("softStyle"))
	return rv
}/* debug [class_properties_class/property]: softStyle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScrollEdgeEffectStyle */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScrollEdgeEffectStyle */

// The automatic scroll edge effect style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollEdgeEffectStyle/automatic
func (s_ ScrollEdgeEffectStyle) AutomaticStyle() IScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](s_.ID, objc.Sel("automaticStyle"))
	return rv
}/* debug [instance_properties/getter]: automaticStyle */


// A scroll edge effect with a hard cutoff.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollEdgeEffectStyle/hard
func (s_ ScrollEdgeEffectStyle) HardStyle() IScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](s_.ID, objc.Sel("hardStyle"))
	return rv
}/* debug [instance_properties/getter]: hardStyle */


// A scroll edge effect with a soft edge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollEdgeEffectStyle/soft
func (s_ ScrollEdgeEffectStyle) SoftStyle() IScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](s_.ID, objc.Sel("softStyle"))
	return rv
}/* debug [instance_properties/getter]: softStyle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScrollEdgeEffectStyle */



