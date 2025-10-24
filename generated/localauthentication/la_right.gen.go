// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Right] class.
var (
	RightClass     _RightClass
	RightClassOnce sync.Once
)

func getRightClass() _RightClass {
	RightClassOnce.Do(func() {
		RightClass = _RightClass{objc.GetClass("LARight")}
	})
	return RightClass
}

type _RightClass struct {
	class objc.Class
}

// An interface definition for the [Right] class.
type IRight interface {
	objectivec.IObject
	// properties:
	State() RightState
	Tag() int
	SetTag(value int)
	// methods:
	AuthorizeWithLocalizedReasonCompletion(localizedReason objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
	AuthorizeWithLocalizedReasonInPresentationContextCompletion(localizedReason objc.IObject /* cross-framework: NSString */, presentationContext PresentationContext /* not a class type */, handler unsafe.Pointer)
	CheckCanAuthorizeWithCompletion(handler unsafe.Pointer)
	DeauthorizeWithCompletion(handler unsafe.Pointer)
}

// A grouped set of requirements that gate access to a resource or operation.
//
// Use instances to protect access to portions of your app that may contain sensitive information. By default, instances require people to authenticate with Face ID, Touch ID, Apple Watch, or the device passcode. The following creates an with the default authentication requirements:


// A grouped set of requirements that gate access to a resource or operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight
type Right struct {
	objectivec.Object
}

// RightFrom constructs a [Right] from an unsafe.Pointer.
//
// A grouped set of requirements that gate access to a resource or operation.
func RightFrom(ptr unsafe.Pointer) Right {
	return Right{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RightClass) Alloc() Right {
	rv := objc.Send[Right](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RightClass) New() Right {
	rv := objc.Send[Right](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ Right) Init() Right {
	rv := objc.Send[Right](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ Right) Autorelease() Right {
	rv := objc.Send[Right](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRight creates a new Right instance.
func NewRight() Right {
	return getRightClass().New()
}



// Creates a right with the authentication requirements you supply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/init(requirement:)
func NewRightWithRequirement(requirement ILAAuthenticationRequirement) Right {
	instance := getRightClass().Alloc()
	rv := objc.Send[Right](instance.ID, objc.Sel("initWithRequirement:"), requirement)
	rv.Autorelease()
	return rv
}



// Performs an authorization on the right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/authorize(localizedReason:completion:)
func (r_ Right) AuthorizeWithLocalizedReasonCompletion(localizedReason objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("authorizeWithLocalizedReason:completion:"), localizedReason, handler)
}


// Performs an authorization on the right with a window context you supply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/authorize(localizedReason:in:completion:)
func (r_ Right) AuthorizeWithLocalizedReasonInPresentationContextCompletion(localizedReason objc.IObject /* cross-framework: NSString */, presentationContext PresentationContext /* not a class type */, handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("authorizeWithLocalizedReason:inPresentationContext:completion:"), localizedReason, presentationContext, handler)
}


// Checks whether the right has permission to perform authorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/checkCanAuthorize(completion:)
func (r_ Right) CheckCanAuthorizeWithCompletion(handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("checkCanAuthorizeWithCompletion:"), handler)
}


// Invalidates a previously authorized right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/deauthorize(completion:)
func (r_ Right) DeauthorizeWithCompletion(handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("deauthorizeWithCompletion:"), handler)
}


// The current authorization state for a right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/state-swift.property
func (r_ Right) State() RightState {
	rv := objc.Send[RightState](r_.ID, objc.Sel("state"))
	return rv
}


// An integer you use to identify a right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/tag
func (r_ Right) Tag() int {
	rv := objc.Send[int](r_.ID, objc.Sel("tag"))
	return rv
}


// An integer you use to identify a right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/tag
func (r_ Right) SetTag(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTag:"), value)
}


