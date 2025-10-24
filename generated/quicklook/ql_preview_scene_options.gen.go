// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QLPreviewSceneOptions */


/* debug [class_header]: Header for QLPreviewSceneOptions */
// The class instance for the [PreviewSceneOptions] class.
var (
	PreviewSceneOptionsClass     _PreviewSceneOptionsClass
	PreviewSceneOptionsClassOnce sync.Once
)

func getPreviewSceneOptionsClass() _PreviewSceneOptionsClass {
	PreviewSceneOptionsClassOnce.Do(func() {
		PreviewSceneOptionsClass = _PreviewSceneOptionsClass{objc.GetClass("QLPreviewSceneOptions")}
	})
	return PreviewSceneOptionsClass
}

type _PreviewSceneOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PreviewSceneOptions */
// An interface definition for the [PreviewSceneOptions] class.
type IPreviewSceneOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PreviewSceneOptions */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PreviewSceneOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PreviewSceneOptions */
// Alloc allocates a new instance without initialization.
func (pc _PreviewSceneOptionsClass) Alloc() PreviewSceneOptions {
	rv := objc.Send[PreviewSceneOptions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PreviewSceneOptionsClass) New() PreviewSceneOptions {
	rv := objc.Send[PreviewSceneOptions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewSceneOptions) Init() PreviewSceneOptions {
	rv := objc.Send[PreviewSceneOptions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewSceneOptions) Autorelease() PreviewSceneOptions {
	rv := objc.Send[PreviewSceneOptions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewSceneOptions creates a new PreviewSceneOptions instance.
func NewPreviewSceneOptions() PreviewSceneOptions {
	return getPreviewSceneOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PreviewSceneOptions */
// A class that represents the configuration for a preview scene activation.


// A class that represents the configuration for a preview scene activation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewSceneActivationConfiguration/Options
type PreviewSceneOptions struct {
	objectivec.Object
}

// PreviewSceneOptionsFrom constructs a [PreviewSceneOptions] from an unsafe.Pointer.
//
// A class that represents the configuration for a preview scene activation.
func PreviewSceneOptionsFrom(ptr unsafe.Pointer) PreviewSceneOptions {
	return PreviewSceneOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PreviewSceneOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PreviewSceneOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PreviewSceneOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PreviewSceneOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PreviewSceneOptions */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLPreviewSceneOptions */


