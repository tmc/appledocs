// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSScrubberSelectionStyle */


/* debug [class_header]: Header for NSScrubberSelectionStyle */
// The class instance for the [ScrubberSelectionStyle] class.
var (
	ScrubberSelectionStyleClass     _ScrubberSelectionStyleClass
	ScrubberSelectionStyleClassOnce sync.Once
)

func getScrubberSelectionStyleClass() _ScrubberSelectionStyleClass {
	ScrubberSelectionStyleClassOnce.Do(func() {
		ScrubberSelectionStyleClass = _ScrubberSelectionStyleClass{objc.GetClass("NSScrubberSelectionStyle")}
	})
	return ScrubberSelectionStyleClass
}

type _ScrubberSelectionStyleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScrubberSelectionStyle */
// An interface definition for the [ScrubberSelectionStyle] class.
type IScrubberSelectionStyle interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ScrubberSelectionStyle */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScrubberSelectionStyle */
	// methods:
	MakeSelectionView() IScrubberSelectionView
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScrubberSelectionStyle */
// Alloc allocates a new instance without initialization.
func (sc _ScrubberSelectionStyleClass) Alloc() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberSelectionStyleClass) New() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberSelectionStyle) Init() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberSelectionStyle) Autorelease() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberSelectionStyle creates a new ScrubberSelectionStyle instance.
func NewScrubberSelectionStyle() ScrubberSelectionStyle {
	return getScrubberSelectionStyleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScrubberSelectionStyle */
// An abstract class that provides decorative accessory views for selected and highlighted items within a scrubber control.
//
// Choose a selection style ( or ), or create a custom selection style by subclassing and overriding .


// An abstract class that provides decorative accessory views for selected and highlighted items within a scrubber control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle
type ScrubberSelectionStyle struct {
	objectivec.Object
}

// ScrubberSelectionStyleFrom constructs a [ScrubberSelectionStyle] from an unsafe.Pointer.
//
// An abstract class that provides decorative accessory views for selected and highlighted items within a scrubber control.
func ScrubberSelectionStyleFrom(ptr unsafe.Pointer) ScrubberSelectionStyle {
	return ScrubberSelectionStyle{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScrubberSelectionStyle */

// Initializes a scrubber selection style when included from a nib or Storyboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/init(coder:)
func NewScrubberSelectionStyleWithCoder(coder foundation.Coder) ScrubberSelectionStyle {
	instance := getScrubberSelectionStyleClass().Alloc()
	rv := objc.Send[ScrubberSelectionStyle](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScrubberSelectionStyleWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScrubberSelectionStyle */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScrubberSelectionStyle */

// A built-in selection style that draws the outline of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/outlineOverlay
func (sc _ScrubberSelectionStyleClass) OutlineOverlayStyle() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](objc.ID(sc.class), objc.Sel("outlineOverlayStyle"))
	return rv
}/* debug [class_properties_class/property]: outlineOverlayStyle */

// A built-in selection style that draws a rounded rectangle as the background of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/roundedBackground
func (sc _ScrubberSelectionStyleClass) RoundedBackgroundStyle() ScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](objc.ID(sc.class), objc.Sel("roundedBackgroundStyle"))
	return rv
}/* debug [class_properties_class/property]: roundedBackgroundStyle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScrubberSelectionStyle */

// Provides an opportunity to create a customized scrubber selection style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/makeSelectionView()
func (s_ ScrubberSelectionStyle) MakeSelectionView() IScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](s_.ID, objc.Sel("makeSelectionView"))
	return rv
}/* debug [instance_methods/method]: MakeSelectionView */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScrubberSelectionStyle */

// A built-in selection style that draws the outline of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/outlineOverlay
func (s_ ScrubberSelectionStyle) OutlineOverlayStyle() IScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](s_.ID, objc.Sel("outlineOverlayStyle"))
	return rv
}/* debug [instance_properties/getter]: outlineOverlayStyle */


// A built-in selection style that draws a rounded rectangle as the background of the scrubber item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionStyle/roundedBackground
func (s_ ScrubberSelectionStyle) RoundedBackgroundStyle() IScrubberSelectionStyle {
	rv := objc.Send[ScrubberSelectionStyle](s_.ID, objc.Sel("roundedBackgroundStyle"))
	return rv
}/* debug [instance_properties/getter]: roundedBackgroundStyle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScrubberSelectionStyle */


