// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKOpenPanelParameters */

/* debug [class_header]: Header for WKOpenPanelParameters */
// The class instance for the [OpenPanelParameters] class.
var (
	OpenPanelParametersClass     _OpenPanelParametersClass
	OpenPanelParametersClassOnce sync.Once
)

func getOpenPanelParametersClass() _OpenPanelParametersClass {
	OpenPanelParametersClassOnce.Do(func() {
		OpenPanelParametersClass = _OpenPanelParametersClass{objc.GetClass("WKOpenPanelParameters")}
	})
	return OpenPanelParametersClass
}

type _OpenPanelParametersClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for OpenPanelParameters */
// An interface definition for the [OpenPanelParameters] class.
type IOpenPanelParameters interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for OpenPanelParameters */
	// properties:
	AllowsDirectories() bool
	AllowsMultipleSelection() bool
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for OpenPanelParameters */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for OpenPanelParameters */
// Alloc allocates a new instance without initialization.
func (oc _OpenPanelParametersClass) Alloc() OpenPanelParameters {
	rv := objc.Send[OpenPanelParameters](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OpenPanelParametersClass) New() OpenPanelParameters {
	rv := objc.Send[OpenPanelParameters](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OpenPanelParameters) Init() OpenPanelParameters {
	rv := objc.Send[OpenPanelParameters](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OpenPanelParameters) Autorelease() OpenPanelParameters {
	rv := objc.Send[OpenPanelParameters](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOpenPanelParameters creates a new OpenPanelParameters instance.
func NewOpenPanelParameters() OpenPanelParameters {
	return getOpenPanelParametersClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for OpenPanelParameters */
// The configuration details of a file upload control in your web content.
//
// Use a to determine the configuration of a file upload control. You don’t create this object directly. Instead, a web view creates one and passes it to the method of its UI delegate object when it displays a file upload control.

// The configuration details of a file upload control in your web content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKOpenPanelParameters
type OpenPanelParameters struct {
	objectivec.Object
}

// OpenPanelParametersFrom constructs a [OpenPanelParameters] from an unsafe.Pointer.
//
// The configuration details of a file upload control in your web content.
func OpenPanelParametersFrom(ptr unsafe.Pointer) OpenPanelParameters {
	return OpenPanelParameters{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for OpenPanelParameters */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for OpenPanelParameters */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for OpenPanelParameters */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for OpenPanelParameters */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for OpenPanelParameters */

// A Boolean value that indicates whether the file upload control supports the selection of directories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKOpenPanelParameters/allowsDirectories
func (o_ OpenPanelParameters) AllowsDirectories() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("allowsDirectories"))
	return rv
} /* debug [instance_properties/getter]: allowsDirectories */

// A Boolean value that indicates whether the file upload control supports multiple files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKOpenPanelParameters/allowsMultipleSelection
func (o_ OpenPanelParameters) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
} /* debug [instance_properties/getter]: allowsMultipleSelection */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKOpenPanelParameters */
