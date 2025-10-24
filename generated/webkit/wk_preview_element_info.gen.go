// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKPreviewElementInfo */


/* debug [class_header]: Header for WKPreviewElementInfo */
// The class instance for the [PreviewElementInfo] class.
var (
	PreviewElementInfoClass     _PreviewElementInfoClass
	PreviewElementInfoClassOnce sync.Once
)

func getPreviewElementInfoClass() _PreviewElementInfoClass {
	PreviewElementInfoClassOnce.Do(func() {
		PreviewElementInfoClass = _PreviewElementInfoClass{objc.GetClass("WKPreviewElementInfo")}
	})
	return PreviewElementInfoClass
}

type _PreviewElementInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PreviewElementInfo */
// An interface definition for the [PreviewElementInfo] class.
type IPreviewElementInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PreviewElementInfo */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PreviewElementInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PreviewElementInfo */
// Alloc allocates a new instance without initialization.
func (pc _PreviewElementInfoClass) Alloc() PreviewElementInfo {
	rv := objc.Send[PreviewElementInfo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PreviewElementInfoClass) New() PreviewElementInfo {
	rv := objc.Send[PreviewElementInfo](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewElementInfo) Init() PreviewElementInfo {
	rv := objc.Send[PreviewElementInfo](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewElementInfo) Autorelease() PreviewElementInfo {
	rv := objc.Send[PreviewElementInfo](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewElementInfo creates a new PreviewElementInfo instance.
func NewPreviewElementInfo() PreviewElementInfo {
	return getPreviewElementInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PreviewElementInfo */
// The object contains information for previewing a webpage.


// The object contains information for previewing a webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreviewElementInfo
type PreviewElementInfo struct {
	objectivec.Object
}

// PreviewElementInfoFrom constructs a [PreviewElementInfo] from an unsafe.Pointer.
//
// The object contains information for previewing a webpage.
func PreviewElementInfoFrom(ptr unsafe.Pointer) PreviewElementInfo {
	return PreviewElementInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PreviewElementInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PreviewElementInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PreviewElementInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PreviewElementInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PreviewElementInfo */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKPreviewElementInfo */


