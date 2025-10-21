// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ODConfiguration] class.
var (
	ODConfigurationClass     _ODConfigurationClass
	ODConfigurationClassOnce sync.Once
)

func getODConfigurationClass() _ODConfigurationClass {
	ODConfigurationClassOnce.Do(func() {
		ODConfigurationClass = _ODConfigurationClass{objc.GetClass("ODConfiguration")}
	})
	return ODConfigurationClass
}

type _ODConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [ODConfiguration] class.
type IODConfiguration interface {
	objectivec.IObject
	AddTrustTypeTrustAccountTrustPasswordUsernamePasswordJoinExistingError(trustType string, account string, accountPassword string, username string, password string, join bool, error_ unsafe.Pointer) bool
	RemoveTrustUsingUsernamePasswordDeleteTrustAccountError(username string, password string, deleteAccount bool, error_ unsafe.Pointer) bool
	SaveUsingAuthorizationError(authorization unsafe.Pointer, error_ unsafe.Pointer) bool
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration
type ODConfiguration struct {
	objectivec.Object
}

// ODConfigurationFrom constructs a [ODConfiguration] from an unsafe.Pointer.
func ODConfigurationFrom(ptr unsafe.Pointer) ODConfiguration {
	return ODConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _ODConfigurationClass) Alloc() ODConfiguration {
	rv := objc.Send[ODConfiguration](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _ODConfigurationClass) New() ODConfiguration {
	rv := objc.Send[ODConfiguration](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ ODConfiguration) Init() ODConfiguration {
	rv := objc.Send[ODConfiguration](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ ODConfiguration) Autorelease() ODConfiguration {
	rv := objc.Send[ODConfiguration](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewODConfiguration creates a new ODConfiguration instance.
func NewODConfiguration() ODConfiguration {
	return getODConfigurationClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/configuration
func (oc _ODConfigurationClass) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("configuration"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/suggestedTrustAccount(_:)
func (oc _ODConfigurationClass) SuggestedTrustAccount(hostname string) string {
	rv := objc.Send[string](objc.ID(oc.class), objc.Sel("suggestedTrustAccount:"), objc.String(hostname))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/suggestedTrustPassword(_:)
func (oc _ODConfigurationClass) SuggestedTrustPassword(length unsafe.Pointer) string {
	rv := objc.Send[string](objc.ID(oc.class), objc.Sel("suggestedTrustPassword:"), length)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/addTrustType(_:trustAccount:trustPassword:username:password:joinExisting:)
func (o_ ODConfiguration) AddTrustTypeTrustAccountTrustPasswordUsernamePasswordJoinExistingError(trustType string, account string, accountPassword string, username string, password string, join bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addTrustType:trustAccount:trustPassword:username:password:joinExisting:error:"), objc.String(trustType), objc.String(account), objc.String(accountPassword), objc.String(username), objc.String(password), join, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/removeTrust(usingUsername:password:deleteTrustAccount:)
func (o_ ODConfiguration) RemoveTrustUsingUsernamePasswordDeleteTrustAccountError(username string, password string, deleteAccount bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeTrustUsingUsername:password:deleteTrustAccount:error:"), objc.String(username), objc.String(password), deleteAccount, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/save(using:)
func (o_ ODConfiguration) SaveUsingAuthorizationError(authorization unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("saveUsingAuthorization:error:"), authorization, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/authenticationModuleEntries-swift.property
func (o_ ODConfiguration) AuthenticationModuleEntries() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("authenticationModuleEntries"))
	return rv
}


// SetAuthenticationModuleEntries sets the value of the authenticationModuleEntries property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/authenticationModuleEntries-swift.property
func (o_ ODConfiguration) SetAuthenticationModuleEntries(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAuthenticationModuleEntries:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/comment-swift.property
func (o_ ODConfiguration) Comment() string {
	rv := objc.Send[string](o_.ID, objc.Sel("comment"))
	return rv
}


// SetComment sets the value of the comment property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/comment-swift.property
func (o_ ODConfiguration) SetComment(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setComment:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionIdleTimeoutInSeconds-swift.property
func (o_ ODConfiguration) ConnectionIdleTimeoutInSeconds() int {
	rv := objc.Send[int](o_.ID, objc.Sel("connectionIdleTimeoutInSeconds"))
	return rv
}


// SetConnectionIdleTimeoutInSeconds sets the value of the connectionIdleTimeoutInSeconds property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionIdleTimeoutInSeconds-swift.property
func (o_ ODConfiguration) SetConnectionIdleTimeoutInSeconds(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setConnectionIdleTimeoutInSeconds:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionSetupTimeoutInSeconds-swift.property
func (o_ ODConfiguration) ConnectionSetupTimeoutInSeconds() int {
	rv := objc.Send[int](o_.ID, objc.Sel("connectionSetupTimeoutInSeconds"))
	return rv
}


// SetConnectionSetupTimeoutInSeconds sets the value of the connectionSetupTimeoutInSeconds property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionSetupTimeoutInSeconds-swift.property
func (o_ ODConfiguration) SetConnectionSetupTimeoutInSeconds(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setConnectionSetupTimeoutInSeconds:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultMappings-swift.property
func (o_ ODConfiguration) DefaultMappings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("defaultMappings"))
	return rv
}


// SetDefaultMappings sets the value of the defaultMappings property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultMappings-swift.property
func (o_ ODConfiguration) SetDefaultMappings(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDefaultMappings:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultModuleEntries-swift.property
func (o_ ODConfiguration) DefaultModuleEntries() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("defaultModuleEntries"))
	return rv
}


// SetDefaultModuleEntries sets the value of the defaultModuleEntries property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultModuleEntries-swift.property
func (o_ ODConfiguration) SetDefaultModuleEntries(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDefaultModuleEntries:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/discoveryModuleEntries-swift.property
func (o_ ODConfiguration) DiscoveryModuleEntries() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("discoveryModuleEntries"))
	return rv
}


// SetDiscoveryModuleEntries sets the value of the discoveryModuleEntries property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/discoveryModuleEntries-swift.property
func (o_ ODConfiguration) SetDiscoveryModuleEntries(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDiscoveryModuleEntries:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/generalModuleEntries-swift.property
func (o_ ODConfiguration) GeneralModuleEntries() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("generalModuleEntries"))
	return rv
}


// SetGeneralModuleEntries sets the value of the generalModuleEntries property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/generalModuleEntries-swift.property
func (o_ ODConfiguration) SetGeneralModuleEntries(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setGeneralModuleEntries:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/hideRegistration-swift.property
func (o_ ODConfiguration) HideRegistration() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("hideRegistration"))
	return rv
}


// SetHideRegistration sets the value of the hideRegistration property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/hideRegistration-swift.property
func (o_ ODConfiguration) SetHideRegistration(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setHideRegistration:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/manInTheMiddleProtection-swift.property
func (o_ ODConfiguration) ManInTheMiddleProtection() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("manInTheMiddleProtection"))
	return rv
}


// SetManInTheMiddleProtection sets the value of the manInTheMiddleProtection property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/manInTheMiddleProtection-swift.property
func (o_ ODConfiguration) SetManInTheMiddleProtection(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setManInTheMiddleProtection:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/nodeName-swift.property
func (o_ ODConfiguration) NodeName() string {
	rv := objc.Send[string](o_.ID, objc.Sel("nodeName"))
	return rv
}


// SetNodeName sets the value of the nodeName property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/nodeName-swift.property
func (o_ ODConfiguration) SetNodeName(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setNodeName:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetEncryption-swift.property
func (o_ ODConfiguration) PacketEncryption() int {
	rv := objc.Send[int](o_.ID, objc.Sel("packetEncryption"))
	return rv
}


// SetPacketEncryption sets the value of the packetEncryption property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetEncryption-swift.property
func (o_ ODConfiguration) SetPacketEncryption(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPacketEncryption:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetSigning-swift.property
func (o_ ODConfiguration) PacketSigning() int {
	rv := objc.Send[int](o_.ID, objc.Sel("packetSigning"))
	return rv
}


// SetPacketSigning sets the value of the packetSigning property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetSigning-swift.property
func (o_ ODConfiguration) SetPacketSigning(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPacketSigning:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostName-swift.property
func (o_ ODConfiguration) PreferredDestinationHostName() string {
	rv := objc.Send[string](o_.ID, objc.Sel("preferredDestinationHostName"))
	return rv
}


// SetPreferredDestinationHostName sets the value of the preferredDestinationHostName property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostName-swift.property
func (o_ ODConfiguration) SetPreferredDestinationHostName(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPreferredDestinationHostName:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostPort-swift.property
func (o_ ODConfiguration) PreferredDestinationHostPort() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("preferredDestinationHostPort"))
	return rv
}


// SetPreferredDestinationHostPort sets the value of the preferredDestinationHostPort property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostPort-swift.property
func (o_ ODConfiguration) SetPreferredDestinationHostPort(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPreferredDestinationHostPort:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/queryTimeoutInSeconds-swift.property
func (o_ ODConfiguration) QueryTimeoutInSeconds() int {
	rv := objc.Send[int](o_.ID, objc.Sel("queryTimeoutInSeconds"))
	return rv
}


// SetQueryTimeoutInSeconds sets the value of the queryTimeoutInSeconds property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/queryTimeoutInSeconds-swift.property
func (o_ ODConfiguration) SetQueryTimeoutInSeconds(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setQueryTimeoutInSeconds:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/templateName-swift.property
func (o_ ODConfiguration) TemplateName() string {
	rv := objc.Send[string](o_.ID, objc.Sel("templateName"))
	return rv
}


// SetTemplateName sets the value of the templateName property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/templateName-swift.property
func (o_ ODConfiguration) SetTemplateName(value string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setTemplateName:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustAccount-swift.property
func (o_ ODConfiguration) TrustAccount() string {
	rv := objc.Send[string](o_.ID, objc.Sel("trustAccount"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustKerberosPrincipal-swift.property
func (o_ ODConfiguration) TrustKerberosPrincipal() string {
	rv := objc.Send[string](o_.ID, objc.Sel("trustKerberosPrincipal"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustMetaAccount-swift.property
func (o_ ODConfiguration) TrustMetaAccount() string {
	rv := objc.Send[string](o_.ID, objc.Sel("trustMetaAccount"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustType-swift.property
func (o_ ODConfiguration) TrustType() string {
	rv := objc.Send[string](o_.ID, objc.Sel("trustType"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesKerberosKeytab-swift.property
func (o_ ODConfiguration) TrustUsesKerberosKeytab() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("trustUsesKerberosKeytab"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesMutualAuthentication-swift.property
func (o_ ODConfiguration) TrustUsesMutualAuthentication() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("trustUsesMutualAuthentication"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesSystemKeychain-swift.property
func (o_ ODConfiguration) TrustUsesSystemKeychain() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("trustUsesSystemKeychain"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/virtualSubnodes-swift.property
func (o_ ODConfiguration) VirtualSubnodes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("virtualSubnodes"))
	return rv
}


// SetVirtualSubnodes sets the value of the virtualSubnodes property.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/virtualSubnodes-swift.property
func (o_ ODConfiguration) SetVirtualSubnodes(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setVirtualSubnodes:"), value)
}


