// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class QLPreviewProvider */


/* debug [class_header]: Header for QLPreviewProvider */
// The class instance for the [PreviewProvider] class.
var (
	PreviewProviderClass     _PreviewProviderClass
	PreviewProviderClassOnce sync.Once
)

func getPreviewProviderClass() _PreviewProviderClass {
	PreviewProviderClassOnce.Do(func() {
		PreviewProviderClass = _PreviewProviderClass{objc.GetClass("QLPreviewProvider")}
	})
	return PreviewProviderClass
}

type _PreviewProviderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PreviewProvider */
// An interface definition for the [PreviewProvider] class.
type IPreviewProvider interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PreviewProvider */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PreviewProvider */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PreviewProvider */
// Alloc allocates a new instance without initialization.
func (pc _PreviewProviderClass) Alloc() PreviewProvider {
	rv := objc.Send[PreviewProvider](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PreviewProviderClass) New() PreviewProvider {
	rv := objc.Send[PreviewProvider](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewProvider) Init() PreviewProvider {
	rv := objc.Send[PreviewProvider](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewProvider) Autorelease() PreviewProvider {
	rv := objc.Send[PreviewProvider](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewProvider creates a new PreviewProvider instance.
func NewPreviewProvider() PreviewProvider {
	return getPreviewProviderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PreviewProvider */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewProvider
type PreviewProvider struct {
	objectivec.Object
}

// PreviewProviderFrom constructs a [PreviewProvider] from an unsafe.Pointer.
func PreviewProviderFrom(ptr unsafe.Pointer) PreviewProvider {
	return PreviewProvider{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PreviewProvider *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PreviewProvider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PreviewProvider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PreviewProvider */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PreviewProvider */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLPreviewProvider */



