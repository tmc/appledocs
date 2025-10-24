// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/securityfoundation"
)

/* debug [class.gen.go]: Generating class ODSession */


/* debug [class_header]: Header for ODSession */
// The class instance for the [ODSession] class.
var (
	ODSessionClass     _ODSessionClass
	ODSessionClassOnce sync.Once
)

func getODSessionClass() _ODSessionClass {
	ODSessionClassOnce.Do(func() {
		ODSessionClass = _ODSessionClass{objc.GetClass("ODSession")}
	})
	return ODSessionClass
}

type _ODSessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ODSession */
// An interface definition for the [ODSession] class.
type IODSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ODSession */
	// properties:
	ConfigurationTemplateNames() objc.IObject /* cross-framework: NSArray */
	MappingTemplateNames() objc.IObject /* cross-framework: NSArray */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ODSession */
	// methods:
	AddConfigurationAuthorizationError(configuration IODConfiguration, authorization securityfoundation.SFAuthorization, error_ unsafe.Pointer) bool
	ConfigurationForNodename(nodename objc.IObject /* cross-framework: NSString */) IODConfiguration
	ConfigurationAuthorizationAllowingUserInteractionError(allowInteraction bool, error_ unsafe.Pointer) securityfoundation.SFAuthorization
	DeleteConfigurationAuthorizationError(configuration IODConfiguration, authorization securityfoundation.SFAuthorization, error_ unsafe.Pointer) bool
	DeleteConfigurationWithNodenameAuthorizationError(nodename objc.IObject /* cross-framework: NSString */, authorization securityfoundation.SFAuthorization, error_ unsafe.Pointer) bool
	NodeNamesAndReturnError(outError unsafe.Pointer) foundation.Array
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ODSession */
// Alloc allocates a new instance without initialization.
func (oc _ODSessionClass) Alloc() ODSession {
	rv := objc.Send[ODSession](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _ODSessionClass) New() ODSession {
	rv := objc.Send[ODSession](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ODSession) Init() ODSession {
	rv := objc.Send[ODSession](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ODSession) Autorelease() ODSession {
	rv := objc.Send[ODSession](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewODSession creates a new ODSession instance.
func NewODSession() ODSession {
	return getODSessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ODSession */
// An object serves as a Cocoa wrapper for an Open Directory session.


// An object serves as a Cocoa wrapper for an Open Directory session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession
type ODSession struct {
	objectivec.Object
}

// ODSessionFrom constructs a [ODSession] from an unsafe.Pointer.
//
// An object serves as a Cocoa wrapper for an Open Directory session.
func ODSessionFrom(ptr unsafe.Pointer) ODSession {
	return ODSession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ODSession */

// Creates a session object directed over proxy to another host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/init(options:)
func NewODSessionWithOptionsError(inOptions objc.IObject /* cross-framework: NSDictionary */, outError unsafe.Pointer) ODSession {
	instance := getODSessionClass().Alloc()
	rv := objc.Send[ODSession](instance.ID, objc.Sel("initWithOptions:error:"), inOptions, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewODSessionWithOptionsError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ODSession */

// Returns a shared instance of the local session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/default()
func (oc _ODSessionClass) DefaultSession() ODSession {
	rv := objc.Send[ODSession](objc.ID(oc.class), objc.Sel("defaultSession"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultSession) */


// Returns an autoreleased session object directed over proxy to another host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/sessionWithOptions:error:
func (oc _ODSessionClass) SessionWithOptionsError(inOptions objc.IObject /* cross-framework: NSDictionary */, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("sessionWithOptions:error:"), inOptions, outError)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SessionWithOptionsError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ODSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ODSession */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/add(_:authorization:)
func (o_ ODSession) AddConfigurationAuthorizationError(configuration IODConfiguration, authorization securityfoundation.SFAuthorization, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addConfiguration:authorization:error:"), configuration, authorization, error_)
	return rv
}/* debug [instance_methods/method]: AddConfigurationAuthorizationError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/configuration(forNodename:)
func (o_ ODSession) ConfigurationForNodename(nodename objc.IObject /* cross-framework: NSString */) IODConfiguration {
	rv := objc.Send[ODConfiguration](o_.ID, objc.Sel("configurationForNodename:"), nodename)
	return rv
}/* debug [instance_methods/method]: ConfigurationForNodename */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/configurationAuthorizationAllowingUserInteraction(_:)
func (o_ ODSession) ConfigurationAuthorizationAllowingUserInteractionError(allowInteraction bool, error_ unsafe.Pointer) securityfoundation.SFAuthorization {
	rv := objc.Send[securityfoundation.SFAuthorization](o_.ID, objc.Sel("configurationAuthorizationAllowingUserInteraction:error:"), allowInteraction, error_)
	return rv
}/* debug [instance_methods/method]: ConfigurationAuthorizationAllowingUserInteractionError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/delete(_:authorization:)
func (o_ ODSession) DeleteConfigurationAuthorizationError(configuration IODConfiguration, authorization securityfoundation.SFAuthorization, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("deleteConfiguration:authorization:error:"), configuration, authorization, error_)
	return rv
}/* debug [instance_methods/method]: DeleteConfigurationAuthorizationError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/deleteConfiguration(withNodename:authorization:)
func (o_ ODSession) DeleteConfigurationWithNodenameAuthorizationError(nodename objc.IObject /* cross-framework: NSString */, authorization securityfoundation.SFAuthorization, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("deleteConfigurationWithNodename:authorization:error:"), nodename, authorization, error_)
	return rv
}/* debug [instance_methods/method]: DeleteConfigurationWithNodenameAuthorizationError */


// Returns the node names that are registered with this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/nodeNames()
func (o_ ODSession) NodeNamesAndReturnError(outError unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](o_.ID, objc.Sel("nodeNamesAndReturnError:"), outError)
	return rv
}/* debug [instance_methods/method]: NodeNamesAndReturnError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ODSession */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/configurationTemplateNames
func (o_ ODSession) ConfigurationTemplateNames() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("configurationTemplateNames"))
	return rv
}/* debug [instance_properties/getter]: configurationTemplateNames */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/mappingTemplateNames
func (o_ ODSession) MappingTemplateNames() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("mappingTemplateNames"))
	return rv
}/* debug [instance_properties/getter]: mappingTemplateNames */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ODSession */


