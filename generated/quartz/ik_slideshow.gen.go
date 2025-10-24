// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IKSlideshow */


/* debug [class_header]: Header for IKSlideshow */
// The class instance for the [IKSlideshow] class.
var (
	IKSlideshowClass     _IKSlideshowClass
	IKSlideshowClassOnce sync.Once
)

func getIKSlideshowClass() _IKSlideshowClass {
	IKSlideshowClassOnce.Do(func() {
		IKSlideshowClass = _IKSlideshowClass{objc.GetClass("IKSlideshow")}
	})
	return IKSlideshowClass
}

type _IKSlideshowClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IKSlideshow */
// An interface definition for the [IKSlideshow] class.
type IIKSlideshow interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for IKSlideshow */
	// properties:
	AutoPlayDelay() float64
	SetAutoPlayDelay(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IKSlideshow */
	// methods:
	IndexOfCurrentSlideshowItem() uint
	ReloadData()
	ReloadSlideshowItemAtIndex(index uint)
	RunSlideshowWithDataSourceInModeOptions(dataSource unsafe.Pointer, slideshowMode objc.IObject /* cross-framework: NSString */, slideshowOptions objc.IObject /* cross-framework: NSDictionary */)
	StopSlideshow(sender objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IKSlideshow */
// Alloc allocates a new instance without initialization.
func (ic _IKSlideshowClass) Alloc() IKSlideshow {
	rv := objc.Send[IKSlideshow](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IKSlideshowClass) New() IKSlideshow {
	rv := objc.Send[IKSlideshow](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IKSlideshow) Init() IKSlideshow {
	rv := objc.Send[IKSlideshow](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IKSlideshow) Autorelease() IKSlideshow {
	rv := objc.Send[IKSlideshow](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIKSlideshow creates a new IKSlideshow instance.
func NewIKSlideshow() IKSlideshow {
	return getIKSlideshowClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IKSlideshow */
// The class encapsulates a data source and options for a slideshow.


// The class encapsulates a data source and options for a slideshow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSlideshow
type IKSlideshow struct {
	objectivec.Object
}

// IKSlideshowFrom constructs a [IKSlideshow] from an unsafe.Pointer.
//
// The class encapsulates a data source and options for a slideshow.
func IKSlideshowFrom(ptr unsafe.Pointer) IKSlideshow {
	return IKSlideshow{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IKSlideshow *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IKSlideshow */

// Finds out whether the slideshow can export its contents to an application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSlideshow/canExport(toApplication:)
func (ic _IKSlideshowClass) CanExportToApplication(applicationBundleIdentifier objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("canExportToApplication:"), applicationBundleIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CanExportToApplication) */


// Exports a slideshow item to the application that has the provided bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSlideshow/exportItem(_:toApplication:)
func (ic _IKSlideshowClass) ExportSlideshowItemToApplication(item objc.IObject, applicationBundleIdentifier objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("exportSlideshowItem:toApplication:"), item, applicationBundleIdentifier)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ExportSlideshowItemToApplication) */


// Returns a shared instance of a slideshow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSlideshow/shared()
func (ic _IKSlideshowClass) SharedSlideshow() IKSlideshow {
	rv := objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("sharedSlideshow"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedSlideshow) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IKSlideshow */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IKSlideshow */

// Returns the index of the current slideshow item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSlideshow/indexOfCurrentSlideshowItem()
func (i_ IKSlideshow) IndexOfCurrentSlideshowItem() uint {
	rv := objc.Send[uint](i_.ID, objc.Sel("indexOfCurrentSlideshowItem"))
	return rv
}/* debug [instance_methods/method]: IndexOfCurrentSlideshowItem */


// Reloads the data for a slideshow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSlideshow/reloadData()
func (i_ IKSlideshow) ReloadData() {
	objc.Send[objc.ID](i_.ID, objc.Sel("reloadData"))
}/* debug [instance_methods/method]: ReloadData */


// Reloads the data for a slideshow, starting at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSlideshow/reloadItem(at:)
func (i_ IKSlideshow) ReloadSlideshowItemAtIndex(index uint) {
	objc.Send[objc.ID](i_.ID, objc.Sel("reloadSlideshowItemAtIndex:"), index)
}/* debug [instance_methods/method]: ReloadSlideshowItemAtIndex */


// Runs a slideshow that contains the specified kind of items, provided from a data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSlideshow/run(with:inMode:options:)
func (i_ IKSlideshow) RunSlideshowWithDataSourceInModeOptions(dataSource unsafe.Pointer, slideshowMode objc.IObject /* cross-framework: NSString */, slideshowOptions objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("runSlideshowWithDataSource:inMode:options:"), dataSource, slideshowMode, slideshowOptions)
}/* debug [instance_methods/method]: RunSlideshowWithDataSourceInModeOptions */


// Stops a slideshow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSlideshow/stop(_:)
func (i_ IKSlideshow) StopSlideshow(sender objc.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("stopSlideshow:"), sender)
}/* debug [instance_methods/method]: StopSlideshow */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IKSlideshow */

// Controls the interval of time before a slideshow starts to play automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSlideshow/autoPlayDelay
func (i_ IKSlideshow) AutoPlayDelay() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("autoPlayDelay"))
	return rv
}/* debug [instance_properties/getter]: autoPlayDelay */


// Controls the interval of time before a slideshow starts to play automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/IKSlideshow/autoPlayDelay
func (i_ IKSlideshow) SetAutoPlayDelay(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAutoPlayDelay:"), value)
}/* debug [instance_properties/setter]: autoPlayDelay */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IKSlideshow */



