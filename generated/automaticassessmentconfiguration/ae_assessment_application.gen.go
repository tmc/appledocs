// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AEAssessmentApplication */


/* debug [class_header]: Header for AEAssessmentApplication */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AEAssessmentApplication */
// An interface definition for the [AEAssessmentApplication] class.
type IAEAssessmentApplication interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AEAssessmentApplication */
	// properties:
	BundleIdentifier() objc.IObject /* cross-framework: NSString */
	RequiresSignatureValidation() bool
	SetRequiresSignatureValidation(value bool)
	TeamIdentifier() objc.IObject /* cross-framework: NSString */
	ConfigurationsByApplication() IAEAssessmentParticipantConfiguration
	SetConfigurationsByApplication(value IAEAssessmentParticipantConfiguration)
	MainParticipantConfiguration() IAEAssessmentParticipantConfiguration
	SetMainParticipantConfiguration(value IAEAssessmentParticipantConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AEAssessmentApplication */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AEAssessmentApplication */
// Alloc allocates a new instance without initialization.
func (ac _AEAssessmentApplicationClass) Alloc() AEAssessmentApplication {
	rv := objc.Send[AEAssessmentApplication](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AEAssessmentApplication */
// A representation of an app that users can access during an assessment.
//
// Use an instance of this class when you want to make an app besides yours, like a calculator or a dictionary, available during an assessment. Create a representation of the app that you want to allow using the app’s bundle identifier and optionally the identifier of the team that distributes the app. You can get both identifiers for an app that you have installed using the command line utility: By default, the system requires that the app’s code signature is valid, and that either Apple distributes the app, or the developer notarizes the app or distributes it through the App Store. You can relax these requirements by setting the property to , but that creates a potential security risk. In that case, the only requirement is that the app has the specified bundle and team identifiers. Prefer to keep the signature requirement. Add the app to a session configuration by calling the method, and then apply the configuration to either a new session that you create, or an existing session with the method.


// A representation of an app that users can access during an assessment.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AEAssessmentApplication */

// Creates a representation of an app using its bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentApplication/init(bundleIdentifier:)
func NewAEAssessmentApplicationWithBundleIdentifier(bundleIdentifier objc.IObject /* cross-framework: NSString */) AEAssessmentApplication {
	instance := getAEAssessmentApplicationClass().Alloc()
	rv := objc.Send[AEAssessmentApplication](instance.ID, objc.Sel("initWithBundleIdentifier:"), bundleIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAEAssessmentApplicationWithBundleIdentifier */


// Creates a representation of an app using its bundle and team identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentApplication/init(bundleIdentifier:teamIdentifier:)
func NewAEAssessmentApplicationWithBundleIdentifierTeamIdentifier(bundleIdentifier objc.IObject /* cross-framework: NSString */, teamIdentifier objc.IObject /* cross-framework: NSString */) AEAssessmentApplication {
	instance := getAEAssessmentApplicationClass().Alloc()
	rv := objc.Send[AEAssessmentApplication](instance.ID, objc.Sel("initWithBundleIdentifier:teamIdentifier:"), bundleIdentifier, teamIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAEAssessmentApplicationWithBundleIdentifierTeamIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AEAssessmentApplication */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AEAssessmentApplication */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AEAssessmentApplication */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AEAssessmentApplication */

// The bundle identifier of the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentApplication/bundleIdentifier
func (a_ AEAssessmentApplication) BundleIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("bundleIdentifier"))
	return rv
}/* debug [instance_properties/getter]: bundleIdentifier */


// A Boolean that indicates whether the session requires the app to have a valid code signature to run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentApplication/requiresSignatureValidation
func (a_ AEAssessmentApplication) RequiresSignatureValidation() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("requiresSignatureValidation"))
	return rv
}/* debug [instance_properties/getter]: requiresSignatureValidation */


// A Boolean that indicates whether the session requires the app to have a valid code signature to run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentApplication/requiresSignatureValidation
func (a_ AEAssessmentApplication) SetRequiresSignatureValidation(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRequiresSignatureValidation:"), value)
}/* debug [instance_properties/setter]: requiresSignatureValidation */


// The team identifier of the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AutomaticAssessmentConfiguration/AEAssessmentApplication/teamIdentifier
func (a_ AEAssessmentApplication) TeamIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("teamIdentifier"))
	return rv
}/* debug [instance_properties/getter]: teamIdentifier */


// The collection of apps available during an assessment, along with their associated configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/configurationsbyapplication
func (a_ AEAssessmentApplication) ConfigurationsByApplication() IAEAssessmentParticipantConfiguration {
	rv := objc.Send[AEAssessmentParticipantConfiguration](a_.ID, objc.Sel("configurationsByApplication"))
	return rv
}/* debug [instance_properties/getter]: configurationsByApplication */


// The collection of apps available during an assessment, along with their associated configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/configurationsbyapplication
func (a_ AEAssessmentApplication) SetConfigurationsByApplication(value IAEAssessmentParticipantConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setConfigurationsByApplication:"), value)
}/* debug [instance_properties/setter]: configurationsByApplication */


// The app-specific configuration for the app that invokes the assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/mainparticipantconfiguration
func (a_ AEAssessmentApplication) MainParticipantConfiguration() IAEAssessmentParticipantConfiguration {
	rv := objc.Send[AEAssessmentParticipantConfiguration](a_.ID, objc.Sel("mainParticipantConfiguration"))
	return rv
}/* debug [instance_properties/getter]: mainParticipantConfiguration */


// The app-specific configuration for the app that invokes the assessment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/automaticassessmentconfiguration/aeassessmentconfiguration/mainparticipantconfiguration
func (a_ AEAssessmentApplication) SetMainParticipantConfiguration(value IAEAssessmentParticipantConfiguration) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMainParticipantConfiguration:"), value)
}/* debug [instance_properties/setter]: mainParticipantConfiguration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AEAssessmentApplication */


