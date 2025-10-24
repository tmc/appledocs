// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ARQuickLookPreviewItem */


/* debug [class_header]: Header for ARQuickLookPreviewItem */
// The class instance for the [QuickLookPreviewItem] class.
var (
	QuickLookPreviewItemClass     _QuickLookPreviewItemClass
	QuickLookPreviewItemClassOnce sync.Once
)

func getQuickLookPreviewItemClass() _QuickLookPreviewItemClass {
	QuickLookPreviewItemClassOnce.Do(func() {
		QuickLookPreviewItemClass = _QuickLookPreviewItemClass{objc.GetClass("ARQuickLookPreviewItem")}
	})
	return QuickLookPreviewItemClass
}

type _QuickLookPreviewItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for QuickLookPreviewItem */
// An interface definition for the [QuickLookPreviewItem] class.
type IQuickLookPreviewItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for QuickLookPreviewItem */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for QuickLookPreviewItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for QuickLookPreviewItem */
// Alloc allocates a new instance without initialization.
func (qc _QuickLookPreviewItemClass) Alloc() QuickLookPreviewItem {
	rv := objc.Send[QuickLookPreviewItem](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QuickLookPreviewItemClass) New() QuickLookPreviewItem {
	rv := objc.Send[QuickLookPreviewItem](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QuickLookPreviewItem) Init() QuickLookPreviewItem {
	rv := objc.Send[QuickLookPreviewItem](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QuickLookPreviewItem) Autorelease() QuickLookPreviewItem {
	rv := objc.Send[QuickLookPreviewItem](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuickLookPreviewItem creates a new QuickLookPreviewItem instance.
func NewQuickLookPreviewItem() QuickLookPreviewItem {
	return getQuickLookPreviewItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for QuickLookPreviewItem */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/ARQuickLookPreviewItem
type QuickLookPreviewItem struct {
	objectivec.Object
}

// QuickLookPreviewItemFrom constructs a [QuickLookPreviewItem] from an unsafe.Pointer.
func QuickLookPreviewItemFrom(ptr unsafe.Pointer) QuickLookPreviewItem {
	return QuickLookPreviewItem{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for QuickLookPreviewItem */

// Creates an object representing the 3D content that will be previewed in AR Quick Look.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/ARQuickLookPreviewItem/init(fileAt:)
func NewQuickLookPreviewItemWithFileAtURL(url objc.IObject /* cross-framework: NSURL */) QuickLookPreviewItem {
	instance := getQuickLookPreviewItemClass().Alloc()
	rv := objc.Send[QuickLookPreviewItem](instance.ID, objc.Sel("initWithFileAtURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewQuickLookPreviewItemWithFileAtURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for QuickLookPreviewItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for QuickLookPreviewItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for QuickLookPreviewItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for QuickLookPreviewItem */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ARQuickLookPreviewItem */


