// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AEAssessmentApplication] class.
var (
	AEAssessmentApplicationClass     _AEAssessmentApplicationClass
	AEAssessmentApplicationClassOnce sync.Once
)

func getAEAssessmentApplicationClass() _AEAssessmentApplicationClass {
	AEAssessmentApplicationClassOnce.Do(func() {
		AEAssessmentApplicationClass = _AEAssessmentApplicationClass{objc.GetClass("AEAssessmentApplication")}
	})
	return AEAssessmentApplicationClass
}

type _AEAssessmentApplicationClass struct {
	class objc.Class
}

// An interface definition for the [AEAssessmentApplication] class.
type IAEAssessmentApplication interface {
	objectivec.IObject
}

// A representation of an app that users can access during an assessment.
//
// Use an instance of this class when you want to make an app besides yours, like a calculator or a dictionary, available during an assessment. Create a representation of the app that you want to allow using the app’s bundle identifier and optionally the identifier of the team that distributes the app. You can get both identifiers for an app that you have installed using the command line utility: By default, the system requires that the app’s code signature is valid, and that either Apple distributes the app, or the developer notarizes the app or distributes it through the App Store. You can relax these requirements by setting the property to , but that creates a potential security risk. In that case, the only requirement is that the app has the specified bundle and team identifiers. Prefer to keep the signature requirement. Add the app to a session configuration by calling the method, and then apply the configuration to either a new session that you create, or an existing session with the method.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentApplication
type AEAssessmentApplication struct {
	objectivec.Object
}

// AEAssessmentApplicationFrom constructs a [AEAssessmentApplication] from an unsafe.Pointer.
//
// A representation of an app that users can access during an assessment.
func AEAssessmentApplicationFrom(ptr unsafe.Pointer) AEAssessmentApplication {
	return AEAssessmentApplication{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AEAssessmentApplicationClass) Alloc() AEAssessmentApplication {
	rv := objc.Send[AEAssessmentApplication](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AEAssessmentApplicationClass) New() AEAssessmentApplication {
	rv := objc.Send[AEAssessmentApplication](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AEAssessmentApplication) Init() AEAssessmentApplication {
	rv := objc.Send[AEAssessmentApplication](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AEAssessmentApplication) Autorelease() AEAssessmentApplication {
	rv := objc.Send[AEAssessmentApplication](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAEAssessmentApplication creates a new AEAssessmentApplication instance.
func NewAEAssessmentApplication() AEAssessmentApplication {
	return getAEAssessmentApplicationClass().New()
}


// Creates a representation of an app using its bundle identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentApplication/init(bundleIdentifier:)
func NewAEAssessmentApplicationWithBundleIdentifier(bundleIdentifier string) AEAssessmentApplication {
	instance := getAEAssessmentApplicationClass().Alloc()
	rv := objc.Send[AEAssessmentApplication](instance.ID, objc.Sel("initWithBundleIdentifier:"), objc.String(bundleIdentifier))
	rv.Autorelease()
	return rv
}

// Creates a representation of an app using its bundle and team identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentApplication/init(bundleIdentifier:teamIdentifier:)
func NewAEAssessmentApplicationWithBundleIdentifierTeamIdentifier(bundleIdentifier string, teamIdentifier string) AEAssessmentApplication {
	instance := getAEAssessmentApplicationClass().Alloc()
	rv := objc.Send[AEAssessmentApplication](instance.ID, objc.Sel("initWithBundleIdentifier:teamIdentifier:"), objc.String(bundleIdentifier), objc.String(teamIdentifier))
	rv.Autorelease()
	return rv
}


// The bundle identifier of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentApplication/bundleIdentifier
func (a_ AEAssessmentApplication) BundleIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("bundleIdentifier"))
	return rv
}

// A Boolean that indicates whether the session requires the app to have a valid code signature to run.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentApplication/requiresSignatureValidation
func (a_ AEAssessmentApplication) RequiresSignatureValidation() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("requiresSignatureValidation"))
	return rv
}


// SetRequiresSignatureValidation sets the value of the requiresSignatureValidation property.
// A Boolean that indicates whether the session requires the app to have a valid code signature to run.

//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentApplication/requiresSignatureValidation
func (a_ AEAssessmentApplication) SetRequiresSignatureValidation(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequiresSignatureValidation:"), value)
}
// The team identifier of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentApplication/teamIdentifier
func (a_ AEAssessmentApplication) TeamIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("teamIdentifier"))
	return rv
}


