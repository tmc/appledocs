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
	// properties:
	AuthenticationModuleEntries() objc.IObject /* cross-framework: NSArray */
	SetAuthenticationModuleEntries(value objc.IObject /* cross-framework: NSArray */)
	Comment() objc.IObject /* cross-framework: NSString */
	SetComment(value objc.IObject /* cross-framework: NSString */)
	ConnectionIdleTimeoutInSeconds() int
	SetConnectionIdleTimeoutInSeconds(value int)
	ConnectionSetupTimeoutInSeconds() int
	SetConnectionSetupTimeoutInSeconds(value int)
	DefaultMappings() IODMappings
	SetDefaultMappings(value IODMappings)
	DefaultModuleEntries() objc.IObject /* cross-framework: NSArray */
	SetDefaultModuleEntries(value objc.IObject /* cross-framework: NSArray */)
	DiscoveryModuleEntries() objc.IObject /* cross-framework: NSArray */
	SetDiscoveryModuleEntries(value objc.IObject /* cross-framework: NSArray */)
	GeneralModuleEntries() objc.IObject /* cross-framework: NSArray */
	SetGeneralModuleEntries(value objc.IObject /* cross-framework: NSArray */)
	HideRegistration() bool
	SetHideRegistration(value bool)
	ManInTheMiddleProtection() bool
	SetManInTheMiddleProtection(value bool)
	NodeName() objc.IObject /* cross-framework: NSString */
	SetNodeName(value objc.IObject /* cross-framework: NSString */)
	PacketEncryption() int
	SetPacketEncryption(value int)
	PacketSigning() int
	SetPacketSigning(value int)
	PreferredDestinationHostName() objc.IObject /* cross-framework: NSString */
	SetPreferredDestinationHostName(value objc.IObject /* cross-framework: NSString */)
	PreferredDestinationHostPort() uint16 /* not a class type */
	SetPreferredDestinationHostPort(value uint16 /* not a class type */)
	QueryTimeoutInSeconds() int
	SetQueryTimeoutInSeconds(value int)
	TemplateName() objc.IObject /* cross-framework: NSString */
	SetTemplateName(value objc.IObject /* cross-framework: NSString */)
	TrustAccount() objc.IObject /* cross-framework: NSString */
	TrustKerberosPrincipal() objc.IObject /* cross-framework: NSString */
	TrustMetaAccount() objc.IObject /* cross-framework: NSString */
	TrustType() objc.IObject /* cross-framework: NSString */
	TrustUsesKerberosKeytab() bool
	TrustUsesMutualAuthentication() bool
	TrustUsesSystemKeychain() bool
	VirtualSubnodes() objc.IObject /* cross-framework: NSArray */
	SetVirtualSubnodes(value objc.IObject /* cross-framework: NSArray */)
	// methods:
	AddTrustTypeTrustAccountTrustPasswordUsernamePasswordJoinExistingError(trustType objc.IObject /* cross-framework: NSString */, account objc.IObject /* cross-framework: NSString */, accountPassword objc.IObject /* cross-framework: NSString */, username objc.IObject /* cross-framework: NSString */, password objc.IObject /* cross-framework: NSString */, join bool, error_ unsafe.Pointer) bool
	RemoveTrustUsingUsernamePasswordDeleteTrustAccountError(username objc.IObject /* cross-framework: NSString */, password objc.IObject /* cross-framework: NSString */, deleteAccount bool, error_ unsafe.Pointer) bool
	SaveUsingAuthorizationError(authorization objc.IObject /* cross-framework: SFAuthorization */, error_ unsafe.Pointer) bool
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/configuration
func (oc _ODConfigurationClass) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("configuration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/suggestedTrustAccount(_:)
func (oc _ODConfigurationClass) SuggestedTrustAccount(hostname objc.IObject /* cross-framework: NSString */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[foundation.String](objc.ID(oc.class), objc.Sel("suggestedTrustAccount:"), hostname)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/suggestedTrustPassword(_:)
func (oc _ODConfigurationClass) SuggestedTrustPassword(length uintptr /* not a class type */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[foundation.String](objc.ID(oc.class), objc.Sel("suggestedTrustPassword:"), length)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/addTrustType(_:trustAccount:trustPassword:username:password:joinExisting:)
func (o_ ODConfiguration) AddTrustTypeTrustAccountTrustPasswordUsernamePasswordJoinExistingError(trustType objc.IObject /* cross-framework: NSString */, account objc.IObject /* cross-framework: NSString */, accountPassword objc.IObject /* cross-framework: NSString */, username objc.IObject /* cross-framework: NSString */, password objc.IObject /* cross-framework: NSString */, join bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("addTrustType:trustAccount:trustPassword:username:password:joinExisting:error:"), trustType, account, accountPassword, username, password, join, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/removeTrust(usingUsername:password:deleteTrustAccount:)
func (o_ ODConfiguration) RemoveTrustUsingUsernamePasswordDeleteTrustAccountError(username objc.IObject /* cross-framework: NSString */, password objc.IObject /* cross-framework: NSString */, deleteAccount bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeTrustUsingUsername:password:deleteTrustAccount:error:"), username, password, deleteAccount, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/save(using:)
func (o_ ODConfiguration) SaveUsingAuthorizationError(authorization objc.IObject /* cross-framework: SFAuthorization */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("saveUsingAuthorization:error:"), authorization, error_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/authenticationModuleEntries-swift.property
func (o_ ODConfiguration) AuthenticationModuleEntries() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("authenticationModuleEntries"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/authenticationModuleEntries-swift.property
func (o_ ODConfiguration) SetAuthenticationModuleEntries(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAuthenticationModuleEntries:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/comment-swift.property
func (o_ ODConfiguration) Comment() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("comment"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/comment-swift.property
func (o_ ODConfiguration) SetComment(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setComment:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionIdleTimeoutInSeconds-swift.property
func (o_ ODConfiguration) ConnectionIdleTimeoutInSeconds() int {
	rv := objc.Send[int](o_.ID, objc.Sel("connectionIdleTimeoutInSeconds"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionIdleTimeoutInSeconds-swift.property
func (o_ ODConfiguration) SetConnectionIdleTimeoutInSeconds(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setConnectionIdleTimeoutInSeconds:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionSetupTimeoutInSeconds-swift.property
func (o_ ODConfiguration) ConnectionSetupTimeoutInSeconds() int {
	rv := objc.Send[int](o_.ID, objc.Sel("connectionSetupTimeoutInSeconds"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionSetupTimeoutInSeconds-swift.property
func (o_ ODConfiguration) SetConnectionSetupTimeoutInSeconds(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setConnectionSetupTimeoutInSeconds:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultMappings-swift.property
func (o_ ODConfiguration) DefaultMappings() IODMappings {
	rv := objc.Send[ODMappings](o_.ID, objc.Sel("defaultMappings"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultMappings-swift.property
func (o_ ODConfiguration) SetDefaultMappings(value IODMappings) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDefaultMappings:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultModuleEntries-swift.property
func (o_ ODConfiguration) DefaultModuleEntries() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("defaultModuleEntries"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultModuleEntries-swift.property
func (o_ ODConfiguration) SetDefaultModuleEntries(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDefaultModuleEntries:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/discoveryModuleEntries-swift.property
func (o_ ODConfiguration) DiscoveryModuleEntries() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("discoveryModuleEntries"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/discoveryModuleEntries-swift.property
func (o_ ODConfiguration) SetDiscoveryModuleEntries(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setDiscoveryModuleEntries:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/generalModuleEntries-swift.property
func (o_ ODConfiguration) GeneralModuleEntries() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("generalModuleEntries"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/generalModuleEntries-swift.property
func (o_ ODConfiguration) SetGeneralModuleEntries(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setGeneralModuleEntries:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/hideRegistration-swift.property
func (o_ ODConfiguration) HideRegistration() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("hideRegistration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/hideRegistration-swift.property
func (o_ ODConfiguration) SetHideRegistration(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setHideRegistration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/manInTheMiddleProtection-swift.property
func (o_ ODConfiguration) ManInTheMiddleProtection() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("manInTheMiddleProtection"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/manInTheMiddleProtection-swift.property
func (o_ ODConfiguration) SetManInTheMiddleProtection(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setManInTheMiddleProtection:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/nodeName-swift.property
func (o_ ODConfiguration) NodeName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("nodeName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/nodeName-swift.property
func (o_ ODConfiguration) SetNodeName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setNodeName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetEncryption-swift.property
func (o_ ODConfiguration) PacketEncryption() int {
	rv := objc.Send[int](o_.ID, objc.Sel("packetEncryption"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetEncryption-swift.property
func (o_ ODConfiguration) SetPacketEncryption(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPacketEncryption:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetSigning-swift.property
func (o_ ODConfiguration) PacketSigning() int {
	rv := objc.Send[int](o_.ID, objc.Sel("packetSigning"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetSigning-swift.property
func (o_ ODConfiguration) SetPacketSigning(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPacketSigning:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostName-swift.property
func (o_ ODConfiguration) PreferredDestinationHostName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("preferredDestinationHostName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostName-swift.property
func (o_ ODConfiguration) SetPreferredDestinationHostName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPreferredDestinationHostName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostPort-swift.property
func (o_ ODConfiguration) PreferredDestinationHostPort() uint16 /* not a class type */ {
	rv := objc.Send[uint16](o_.ID, objc.Sel("preferredDestinationHostPort"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostPort-swift.property
func (o_ ODConfiguration) SetPreferredDestinationHostPort(value uint16 /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPreferredDestinationHostPort:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/queryTimeoutInSeconds-swift.property
func (o_ ODConfiguration) QueryTimeoutInSeconds() int {
	rv := objc.Send[int](o_.ID, objc.Sel("queryTimeoutInSeconds"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/queryTimeoutInSeconds-swift.property
func (o_ ODConfiguration) SetQueryTimeoutInSeconds(value int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setQueryTimeoutInSeconds:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/templateName-swift.property
func (o_ ODConfiguration) TemplateName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("templateName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/templateName-swift.property
func (o_ ODConfiguration) SetTemplateName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setTemplateName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustAccount-swift.property
func (o_ ODConfiguration) TrustAccount() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("trustAccount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustKerberosPrincipal-swift.property
func (o_ ODConfiguration) TrustKerberosPrincipal() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("trustKerberosPrincipal"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustMetaAccount-swift.property
func (o_ ODConfiguration) TrustMetaAccount() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("trustMetaAccount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustType-swift.property
func (o_ ODConfiguration) TrustType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("trustType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesKerberosKeytab-swift.property
func (o_ ODConfiguration) TrustUsesKerberosKeytab() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("trustUsesKerberosKeytab"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesMutualAuthentication-swift.property
func (o_ ODConfiguration) TrustUsesMutualAuthentication() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("trustUsesMutualAuthentication"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesSystemKeychain-swift.property
func (o_ ODConfiguration) TrustUsesSystemKeychain() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("trustUsesSystemKeychain"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/virtualSubnodes-swift.property
func (o_ ODConfiguration) VirtualSubnodes() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](o_.ID, objc.Sel("virtualSubnodes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/virtualSubnodes-swift.property
func (o_ ODConfiguration) SetVirtualSubnodes(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setVirtualSubnodes:"), value)
}



