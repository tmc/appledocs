// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVContentProposalViewController */


/* debug [class_header]: Header for AVContentProposalViewController */
// The class instance for the [ContentProposalViewController] class.
var (
	ContentProposalViewControllerClass     _ContentProposalViewControllerClass
	ContentProposalViewControllerClassOnce sync.Once
)

func getContentProposalViewControllerClass() _ContentProposalViewControllerClass {
	ContentProposalViewControllerClassOnce.Do(func() {
		ContentProposalViewControllerClass = _ContentProposalViewControllerClass{objc.GetClass("AVContentProposalViewController")}
	})
	return ContentProposalViewControllerClass
}

type _ContentProposalViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ContentProposalViewController */
// An interface definition for the [ContentProposalViewController] class.
type IContentProposalViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for ContentProposalViewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ContentProposalViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ContentProposalViewController */
// Alloc allocates a new instance without initialization.
func (cc _ContentProposalViewControllerClass) Alloc() ContentProposalViewController {
	rv := objc.Send[ContentProposalViewController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContentProposalViewControllerClass) New() ContentProposalViewController {
	rv := objc.Send[ContentProposalViewController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContentProposalViewController) Init() ContentProposalViewController {
	rv := objc.Send[ContentProposalViewController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContentProposalViewController) Autorelease() ContentProposalViewController {
	rv := objc.Send[ContentProposalViewController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContentProposalViewController creates a new ContentProposalViewController instance.
func NewContentProposalViewController() ContentProposalViewController {
	return getContentProposalViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ContentProposalViewController */
// A view controller that proposes content to watch next.
//
// Subclass this class to define the user interface for your content proposal.


// A view controller that proposes content to watch next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController
type ContentProposalViewController struct {
	ViewController
}

// ContentProposalViewControllerFrom constructs a [ContentProposalViewController] from an unsafe.Pointer.
//
// A view controller that proposes content to watch next.
func ContentProposalViewControllerFrom(ptr unsafe.Pointer) ContentProposalViewController {
	return ContentProposalViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ContentProposalViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ContentProposalViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ContentProposalViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ContentProposalViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ContentProposalViewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVContentProposalViewController */


