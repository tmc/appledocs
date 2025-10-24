// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class LARight */


/* debug [class_header]: Header for LARight */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Right */
// An interface definition for the [Right] class.
type IRight interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Right */
	// properties:
	State() RightState
	Tag() int
	SetTag(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Right */
	// methods:
	AuthorizeWithLocalizedReasonCompletion(localizedReason objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
	AuthorizeWithLocalizedReasonInPresentationContextCompletion(localizedReason objc.IObject /* cross-framework: NSString */, presentationContext PresentationContext /* not a class type */, handler unsafe.Pointer)
	CheckCanAuthorizeWithCompletion(handler unsafe.Pointer)
	DeauthorizeWithCompletion(handler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Right */
// Alloc allocates a new instance without initialization.
func (rc _RightClass) Alloc() Right {
	rv := objc.Send[Right](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Right */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Right */

// Creates a right with the authentication requirements you supply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/init(requirement:)
func NewRightWithRequirement(requirement ILAAuthenticationRequirement) Right {
	instance := getRightClass().Alloc()
	rv := objc.Send[Right](instance.ID, objc.Sel("initWithRequirement:"), requirement)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRightWithRequirement */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Right */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Right */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Right */

// Performs an authorization on the right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/authorize(localizedReason:completion:)
func (r_ Right) AuthorizeWithLocalizedReasonCompletion(localizedReason objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("authorizeWithLocalizedReason:completion:"), localizedReason, handler)
}/* debug [instance_methods/method]: AuthorizeWithLocalizedReasonCompletion */


// Performs an authorization on the right with a window context you supply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/authorize(localizedReason:in:completion:)
func (r_ Right) AuthorizeWithLocalizedReasonInPresentationContextCompletion(localizedReason objc.IObject /* cross-framework: NSString */, presentationContext PresentationContext /* not a class type */, handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("authorizeWithLocalizedReason:inPresentationContext:completion:"), localizedReason, presentationContext, handler)
}/* debug [instance_methods/method]: AuthorizeWithLocalizedReasonInPresentationContextCompletion */


// Checks whether the right has permission to perform authorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/checkCanAuthorize(completion:)
func (r_ Right) CheckCanAuthorizeWithCompletion(handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("checkCanAuthorizeWithCompletion:"), handler)
}/* debug [instance_methods/method]: CheckCanAuthorizeWithCompletion */


// Invalidates a previously authorized right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/deauthorize(completion:)
func (r_ Right) DeauthorizeWithCompletion(handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("deauthorizeWithCompletion:"), handler)
}/* debug [instance_methods/method]: DeauthorizeWithCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Right */

// The current authorization state for a right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/state-swift.property
func (r_ Right) State() RightState {
	rv := objc.Send[RightState](r_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// An integer you use to identify a right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/tag
func (r_ Right) Tag() int {
	rv := objc.Send[int](r_.ID, objc.Sel("tag"))
	return rv
}/* debug [instance_properties/getter]: tag */


// An integer you use to identify a right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/tag
func (r_ Right) SetTag(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTag:"), value)
}/* debug [instance_properties/setter]: tag */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LARight */


