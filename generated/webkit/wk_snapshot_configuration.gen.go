// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [SnapshotConfiguration] class.
type ISnapshotConfiguration interface {
	objectivec.IObject
}

// The configuration data to use when generating an image from a web view’s contents.
//
// Create a object when you want to generate an image based on your web view’s content. Use this object to specify the portion of the web view to capture and the capture behavior. To generate the snapshot, pass the configuration object to the method of , which returns a platform-native image for you to use.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SnapshotConfigurationClass) Alloc() SnapshotConfiguration {
	rv := objc.Send[SnapshotConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration/rect
func (s_ SnapshotConfiguration) Rect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("rect"))
	return rv
}


// SetRect sets the value of the rect property.
// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKSnapshotConfiguration/rect
func (s_ SnapshotConfiguration) SetRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRect:"), value)
}



