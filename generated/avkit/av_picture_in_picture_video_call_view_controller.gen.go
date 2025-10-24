// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVPictureInPictureVideoCallViewController */


/* debug [class_header]: Header for AVPictureInPictureVideoCallViewController */
// The class instance for the [PictureInPictureVideoCallViewController] class.
var (
	PictureInPictureVideoCallViewControllerClass     _PictureInPictureVideoCallViewControllerClass
	PictureInPictureVideoCallViewControllerClassOnce sync.Once
)

func getPictureInPictureVideoCallViewControllerClass() _PictureInPictureVideoCallViewControllerClass {
	PictureInPictureVideoCallViewControllerClassOnce.Do(func() {
		PictureInPictureVideoCallViewControllerClass = _PictureInPictureVideoCallViewControllerClass{objc.GetClass("AVPictureInPictureVideoCallViewController")}
	})
	return PictureInPictureVideoCallViewControllerClass
}

type _PictureInPictureVideoCallViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PictureInPictureVideoCallViewController */
// An interface definition for the [PictureInPictureVideoCallViewController] class.
type IPictureInPictureVideoCallViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for PictureInPictureVideoCallViewController */
	// properties:
	ActiveVideoCallContentViewController() IAVPictureInPictureVideoCallViewController
	SetActiveVideoCallContentViewController(value IAVPictureInPictureVideoCallViewController)
	ActiveVideoCallSourceView() appkit.View
	SetActiveVideoCallSourceView(value appkit.View)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PictureInPictureVideoCallViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PictureInPictureVideoCallViewController */
// Alloc allocates a new instance without initialization.
func (pc _PictureInPictureVideoCallViewControllerClass) Alloc() PictureInPictureVideoCallViewController {
	rv := objc.Send[PictureInPictureVideoCallViewController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PictureInPictureVideoCallViewControllerClass) New() PictureInPictureVideoCallViewController {
	rv := objc.Send[PictureInPictureVideoCallViewController](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PictureInPictureVideoCallViewController) Init() PictureInPictureVideoCallViewController {
	rv := objc.Send[PictureInPictureVideoCallViewController](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PictureInPictureVideoCallViewController) Autorelease() PictureInPictureVideoCallViewController {
	rv := objc.Send[PictureInPictureVideoCallViewController](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPictureInPictureVideoCallViewController creates a new PictureInPictureVideoCallViewController instance.
func NewPictureInPictureVideoCallViewController() PictureInPictureVideoCallViewController {
	return getPictureInPictureVideoCallViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PictureInPictureVideoCallViewController */
// A view controller that presents content from a video call in Picture in Picture.


// A view controller that presents content from a video call in Picture in Picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPictureInPictureVideoCallViewController
type PictureInPictureVideoCallViewController struct {
	ViewController
}

// PictureInPictureVideoCallViewControllerFrom constructs a [PictureInPictureVideoCallViewController] from an unsafe.Pointer.
//
// A view controller that presents content from a video call in Picture in Picture.
func PictureInPictureVideoCallViewControllerFrom(ptr unsafe.Pointer) PictureInPictureVideoCallViewController {
	return PictureInPictureVideoCallViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PictureInPictureVideoCallViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PictureInPictureVideoCallViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PictureInPictureVideoCallViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PictureInPictureVideoCallViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PictureInPictureVideoCallViewController */

// The view controller that presents the video call content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.class/activevideocallcontentviewcontroller
func (p_ PictureInPictureVideoCallViewController) ActiveVideoCallContentViewController() IAVPictureInPictureVideoCallViewController {
	rv := objc.Send[PictureInPictureVideoCallViewController](p_.ID, objc.Sel("activeVideoCallContentViewController"))
	return rv
}/* debug [instance_properties/getter]: activeVideoCallContentViewController */


// The view controller that presents the video call content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.class/activevideocallcontentviewcontroller
func (p_ PictureInPictureVideoCallViewController) SetActiveVideoCallContentViewController(value IAVPictureInPictureVideoCallViewController) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setActiveVideoCallContentViewController:"), value)
}/* debug [instance_properties/setter]: activeVideoCallContentViewController */


// The view that contains the video content of the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.class/activevideocallsourceview
func (p_ PictureInPictureVideoCallViewController) ActiveVideoCallSourceView() appkit.View {
	rv := objc.Send[appkit.View](p_.ID, objc.Sel("activeVideoCallSourceView"))
	return rv
}/* debug [instance_properties/getter]: activeVideoCallSourceView */


// The view that contains the video content of the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avpictureinpicturecontroller/contentsource-swift.class/activevideocallsourceview
func (p_ PictureInPictureVideoCallViewController) SetActiveVideoCallSourceView(value appkit.View) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setActiveVideoCallSourceView:"), value)
}/* debug [instance_properties/setter]: activeVideoCallSourceView */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPictureInPictureVideoCallViewController */



