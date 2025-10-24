// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [ContentProposalViewController] class.
type IContentProposalViewController interface {
	appkit.IViewController
	// properties:
	// methods:
}

// A view controller that proposes content to watch next.
//
// Subclass this class to define the user interface for your content proposal.


// A view controller that proposes content to watch next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVContentProposalViewController
type ContentProposalViewController struct {
	appkit.ViewController
}

// ContentProposalViewControllerFrom constructs a [ContentProposalViewController] from an unsafe.Pointer.
//
// A view controller that proposes content to watch next.
func ContentProposalViewControllerFrom(ptr unsafe.Pointer) ContentProposalViewController {
	return ContentProposalViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ContentProposalViewControllerClass) Alloc() ContentProposalViewController {
	rv := objc.Send[ContentProposalViewController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



