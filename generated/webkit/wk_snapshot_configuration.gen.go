// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKSnapshotConfiguration */

/* debug [class_header]: Header for WKSnapshotConfiguration */
// The class instance for the [SnapshotConfiguration] class.
var (
	SnapshotConfigurationClass     _SnapshotConfigurationClass
	SnapshotConfigurationClassOnce sync.Once
)

func getSnapshotConfigurationClass() _SnapshotConfigurationClass {
	SnapshotConfigurationClassOnce.Do(func() {
		SnapshotConfigurationClass = _SnapshotConfigurationClass{objc.GetClass("WKSnapshotConfiguration")}
	})
	return SnapshotConfigurationClass
}

type _SnapshotConfigurationClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for SnapshotConfiguration */
// An interface definition for the [SnapshotConfiguration] class.
type ISnapshotConfiguration interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for SnapshotConfiguration */
	// properties:
	AfterScreenUpdates() bool
	SetAfterScreenUpdates(value bool)
	Rect() corefoundation.CGRect
	SetRect(value corefoundation.CGRect)
	SnapshotWidth() objc.IObject /* cross-framework: NSNumber */
	SetSnapshotWidth(value objc.IObject /* cross-framework: NSNumber */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for SnapshotConfiguration */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for SnapshotConfiguration */
// Alloc allocates a new instance without initialization.
func (sc _SnapshotConfigurationClass) Alloc() SnapshotConfiguration {
	rv := objc.Send[SnapshotConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SnapshotConfigurationClass) New() SnapshotConfiguration {
	rv := objc.Send[SnapshotConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SnapshotConfiguration) Init() SnapshotConfiguration {
	rv := objc.Send[SnapshotConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SnapshotConfiguration) Autorelease() SnapshotConfiguration {
	rv := objc.Send[SnapshotConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSnapshotConfiguration creates a new SnapshotConfiguration instance.
func NewSnapshotConfiguration() SnapshotConfiguration {
	return getSnapshotConfigurationClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for SnapshotConfiguration */
// The configuration data to use when generating an image from a web view’s contents.
//
// Create a object when you want to generate an image based on your web view’s content. Use this object to specify the portion of the web view to capture and the capture behavior. To generate the snapshot, pass the configuration object to the method of , which returns a platform-native image for you to use.

// The configuration data to use when generating an image from a web view’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration
type SnapshotConfiguration struct {
	objectivec.Object
}

// SnapshotConfigurationFrom constructs a [SnapshotConfiguration] from an unsafe.Pointer.
//
// The configuration data to use when generating an image from a web view’s contents.
func SnapshotConfigurationFrom(ptr unsafe.Pointer) SnapshotConfiguration {
	return SnapshotConfiguration{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for SnapshotConfiguration */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for SnapshotConfiguration */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for SnapshotConfiguration */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for SnapshotConfiguration */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for SnapshotConfiguration */

// A Boolean value that indicates whether to take the snapshot after incorporating any pending screen updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration/afterScreenUpdates
func (s_ SnapshotConfiguration) AfterScreenUpdates() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("afterScreenUpdates"))
	return rv
} /* debug [instance_properties/getter]: afterScreenUpdates */

// A Boolean value that indicates whether to take the snapshot after incorporating any pending screen updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration/afterScreenUpdates
func (s_ SnapshotConfiguration) SetAfterScreenUpdates(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAfterScreenUpdates:"), value)
} /* debug [instance_properties/setter]: afterScreenUpdates */

// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration/rect
func (s_ SnapshotConfiguration) Rect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s_.ID, objc.Sel("rect"))
	return rv
} /* debug [instance_properties/getter]: rect */

// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration/rect
func (s_ SnapshotConfiguration) SetRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRect:"), value)
} /* debug [instance_properties/setter]: rect */

// The width of the captured image, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration/snapshotWidth
func (s_ SnapshotConfiguration) SnapshotWidth() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](s_.ID, objc.Sel("snapshotWidth"))
	return rv
} /* debug [instance_properties/getter]: snapshotWidth */

// The width of the captured image, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration/snapshotWidth
func (s_ SnapshotConfiguration) SetSnapshotWidth(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSnapshotWidth:"), value)
} /* debug [instance_properties/setter]: snapshotWidth */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKSnapshotConfiguration */
