// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKOverlay */

/* debug [class_header]: Header for SKOverlay */
// The class instance for the [Overlay] class.
var (
	OverlayClass     _OverlayClass
	OverlayClassOnce sync.Once
)

func getOverlayClass() _OverlayClass {
	OverlayClassOnce.Do(func() {
		OverlayClass = _OverlayClass{objc.GetClass("SKOverlay")}
	})
	return OverlayClass
}

type _OverlayClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for Overlay */
// An interface definition for the [Overlay] class.
type IOverlay interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for Overlay */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for Overlay */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for Overlay */
// Alloc allocates a new instance without initialization.
func (oc _OverlayClass) Alloc() Overlay {
	rv := objc.Send[Overlay](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OverlayClass) New() Overlay {
	rv := objc.Send[Overlay](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ Overlay) Init() Overlay {
	rv := objc.Send[Overlay](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ Overlay) Autorelease() Overlay {
	rv := objc.Send[Overlay](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOverlay creates a new Overlay instance.
func NewOverlay() Overlay {
	return getOverlayClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for Overlay */
// A class that displays an overlay you can use to recommend another app or an App Clip’s corresponding full app.
//
// By displaying an overlay, you can recommend another app to users and enable them to download it immediately. To recommend media that’s not an app, or to display a product page within your app, use . If you’re using SwiftUI, make use of the modifier. For example usage, see . To display an App Store overlay in an app that uses : Create an with the iTunes identifier of the app you want to recommend. Initialize with the configuration object. Present the overlay. The following code displays an overlay at the bottom of the visible scene: To respond to the overlay’s appearance, dismissal, or failure to load, set the and implement the methods defined in .

// A class that displays an overlay you can use to recommend another app or an App Clip’s corresponding full app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay
type Overlay struct {
	objectivec.Object
}

// OverlayFrom constructs a [Overlay] from an unsafe.Pointer.
//
// A class that displays an overlay you can use to recommend another app or an App Clip’s corresponding full app.
func OverlayFrom(ptr unsafe.Pointer) Overlay {
	return Overlay{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for Overlay */

// Creates an overlay you use to recommend another app on the App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/init(configuration:)
func NewOverlayWithConfiguration(configuration ISKOverlayConfiguration) Overlay {
	instance := getOverlayClass().Alloc()
	rv := objc.Send[Overlay](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewOverlayWithConfiguration */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for Overlay */

// Dismisses an App Store overlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/dismiss(in:)
func (oc _OverlayClass) DismissOverlayInScene(scene WindowScene /* not a class type */) {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("dismissOverlayInScene:"), scene)
} /* debug [class_methods/method]: Class method for%!(EXTRA string=DismissOverlayInScene) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for Overlay */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for Overlay */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for Overlay */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SKOverlay */
