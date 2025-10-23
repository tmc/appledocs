// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [ChallengesViewController] class.
var (
	ChallengesViewControllerClass     _ChallengesViewControllerClass
	ChallengesViewControllerClassOnce sync.Once
)

func getChallengesViewControllerClass() _ChallengesViewControllerClass {
	ChallengesViewControllerClassOnce.Do(func() {
		ChallengesViewControllerClass = _ChallengesViewControllerClass{objc.GetClass("GKChallengesViewController")}
	})
	return ChallengesViewControllerClass
}

type _ChallengesViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [ChallengesViewController] class.
type IChallengesViewController interface {
	appkit.IViewController
	ChallengeDelegate() objc.ID
	SetChallengeDelegate(value objc.ID)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengesViewController
type ChallengesViewController struct {
	appkit.ViewController
}

// ChallengesViewControllerFrom constructs a [ChallengesViewController] from an unsafe.Pointer.
func ChallengesViewControllerFrom(ptr unsafe.Pointer) ChallengesViewController {
	return ChallengesViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ChallengesViewControllerClass) Alloc() ChallengesViewController {
	rv := objc.Send[ChallengesViewController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ChallengesViewControllerClass) New() ChallengesViewController {
	rv := objc.Send[ChallengesViewController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChallengesViewController) Init() ChallengesViewController {
	rv := objc.Send[ChallengesViewController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChallengesViewController) Autorelease() ChallengesViewController {
	rv := objc.Send[ChallengesViewController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChallengesViewController creates a new ChallengesViewController instance.
func NewChallengesViewController() ChallengesViewController {
	return getChallengesViewControllerClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengesViewController/challengeDelegate
func (c_ ChallengesViewController) ChallengeDelegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("challengeDelegate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengesViewController/challengeDelegate
func (c_ ChallengesViewController) SetChallengeDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setChallengeDelegate:"), value)
}


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (c_ ChallengesViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (c_ ChallengesViewController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}



