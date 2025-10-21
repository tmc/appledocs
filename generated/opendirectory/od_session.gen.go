// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [ODSession] class.
type IODSession interface {
	objectivec.IObject
	AddConfigurationAuthorizationError(configuration unsafe.Pointer, authorization unsafe.Pointer, error_ unsafe.Pointer) bool
	ConfigurationForNodename(nodename string) unsafe.Pointer
	ConfigurationAuthorizationAllowingUserInteractionError(allowInteraction bool, error_ unsafe.Pointer) unsafe.Pointer
	DeleteConfigurationAuthorizationError(configuration unsafe.Pointer, authorization unsafe.Pointer, error_ unsafe.Pointer) bool
	DeleteConfigurationWithNodenameAuthorizationError(nodename string, authorization unsafe.Pointer, error_ unsafe.Pointer) bool
	NodeNamesAndReturnError(outError unsafe.Pointer) unsafe.Pointer
}

// An object serves as a Cocoa wrapper for an Open Directory session.
//
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

// Alloc allocates a new instance without initialization.
func (oc _ODSessionClass) Alloc() ODSession {
	rv := objc.Send[ODSession](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a session object directed over proxy to another host.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/init(options:)
func NewODSessionWithOptionsError(inOptions objc.ID, outError unsafe.Pointer) ODSession {
	instance := getODSessionClass().Alloc()
	rv := objc.Send[ODSession](instance.ID, objc.Sel("initWithOptions:error:"), inOptions, outError)
	rv.Autorelease()
	return rv
}


// Returns a shared instance of the local session.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/default()
func (oc _ODSessionClass) DefaultSession() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("defaultSession"))
	return rv
}

// Returns an autoreleased session object directed over proxy to another host.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/sessionWithOptions:error:
func (oc _ODSessionClass) SessionWithOptionsError(inOptions objc.ID, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("sessionWithOptions:error:"), inOptions, outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/add(_:authorization:)
func (o_ ODSession) AddConfigurationAuthorizationError(configuration unsafe.Pointer, authorization unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addConfiguration:authorization:error:"), configuration, authorization, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/configuration(forNodename:)
func (o_ ODSession) ConfigurationForNodename(nodename string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("configurationForNodename:"), objc.String(nodename))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/configurationAuthorizationAllowingUserInteraction(_:)
func (o_ ODSession) ConfigurationAuthorizationAllowingUserInteractionError(allowInteraction bool, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("configurationAuthorizationAllowingUserInteraction:error:"), allowInteraction, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/delete(_:authorization:)
func (o_ ODSession) DeleteConfigurationAuthorizationError(configuration unsafe.Pointer, authorization unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("deleteConfiguration:authorization:error:"), configuration, authorization, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/deleteConfiguration(withNodename:authorization:)
func (o_ ODSession) DeleteConfigurationWithNodenameAuthorizationError(nodename string, authorization unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("deleteConfigurationWithNodename:authorization:error:"), objc.String(nodename), authorization, error_)
	return rv
}

// Returns the node names that are registered with this session.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/nodeNames()
func (o_ ODSession) NodeNamesAndReturnError(outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("nodeNamesAndReturnError:"), outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/configurationTemplateNames
func (o_ ODSession) ConfigurationTemplateNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("configurationTemplateNames"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSession/mappingTemplateNames
func (o_ ODSession) MappingTemplateNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("mappingTemplateNames"))
	return rv
}


