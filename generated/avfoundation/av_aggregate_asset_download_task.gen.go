// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVAggregateAssetDownloadTask */


/* debug [class_header]: Header for AVAggregateAssetDownloadTask */
// The class instance for the [AggregateAssetDownloadTask] class.
var (
	AggregateAssetDownloadTaskClass     _AggregateAssetDownloadTaskClass
	AggregateAssetDownloadTaskClassOnce sync.Once
)

func getAggregateAssetDownloadTaskClass() _AggregateAssetDownloadTaskClass {
	AggregateAssetDownloadTaskClassOnce.Do(func() {
		AggregateAssetDownloadTaskClass = _AggregateAssetDownloadTaskClass{objc.GetClass("AVAggregateAssetDownloadTask")}
	})
	return AggregateAssetDownloadTaskClass
}

type _AggregateAssetDownloadTaskClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AggregateAssetDownloadTask */
// An interface definition for the [AggregateAssetDownloadTask] class.
type IAggregateAssetDownloadTask interface {
	IURLSessionTask
	
/* debug [class_interface_properties]: Properties for AggregateAssetDownloadTask */
	// properties:
	URLAsset() IAVURLAsset
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AggregateAssetDownloadTask */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AggregateAssetDownloadTask */
// Alloc allocates a new instance without initialization.
func (ac _AggregateAssetDownloadTaskClass) Alloc() AggregateAssetDownloadTask {
	rv := objc.Send[AggregateAssetDownloadTask](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AggregateAssetDownloadTaskClass) New() AggregateAssetDownloadTask {
	rv := objc.Send[AggregateAssetDownloadTask](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AggregateAssetDownloadTask) Init() AggregateAssetDownloadTask {
	rv := objc.Send[AggregateAssetDownloadTask](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AggregateAssetDownloadTask) Autorelease() AggregateAssetDownloadTask {
	rv := objc.Send[AggregateAssetDownloadTask](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAggregateAssetDownloadTask creates a new AggregateAssetDownloadTask instance.
func NewAggregateAssetDownloadTask() AggregateAssetDownloadTask {
	return getAggregateAssetDownloadTaskClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AggregateAssetDownloadTask */
// A task that downloads multiple media selections for an asset.


// A task that downloads multiple media selections for an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAggregateAssetDownloadTask
type AggregateAssetDownloadTask struct {
	URLSessionTask
}

// AggregateAssetDownloadTaskFrom constructs a [AggregateAssetDownloadTask] from an unsafe.Pointer.
//
// A task that downloads multiple media selections for an asset.
func AggregateAssetDownloadTaskFrom(ptr unsafe.Pointer) AggregateAssetDownloadTask {
	return AggregateAssetDownloadTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AggregateAssetDownloadTask *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AggregateAssetDownloadTask */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AggregateAssetDownloadTask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AggregateAssetDownloadTask */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AggregateAssetDownloadTask */

// The asset the parent task downloads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAggregateAssetDownloadTask/urlAsset
func (a_ AggregateAssetDownloadTask) URLAsset() IAVURLAsset {
	rv := objc.Send[URLAsset](a_.ID, objc.Sel("URLAsset"))
	return rv
}/* debug [instance_properties/getter]: URLAsset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAggregateAssetDownloadTask */



