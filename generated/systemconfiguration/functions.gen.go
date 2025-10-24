// Code generated from Apple documentation for SystemConfiguration. DO NOT EDIT.

package systemconfiguration

/* debug [functions.gen.go]: Generating 178 functions for SystemConfiguration */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// SystemConfiguration Functions (178 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CNCopyCurrentNetworkInfo func(StringRef) DictionaryRef
	_CNCopySupportedInterfaces func() ArrayRef
	_CNMarkPortalOffline func(StringRef) unsafe.Pointer
	_CNMarkPortalOnline func(StringRef) unsafe.Pointer
	_CNSetSupportedSSIDs func(ArrayRef) unsafe.Pointer
	_DHCPClientPreferencesCopyApplicationOptions func(StringRef, unsafe.Pointer) unsafe.Pointer
	_DHCPClientPreferencesSetApplicationOptions func(StringRef, unsafe.Pointer, Index) unsafe.Pointer
	_DHCPInfoGetLeaseExpirationTime func(DictionaryRef) DateRef
	_DHCPInfoGetLeaseStartTime func(DictionaryRef) DateRef
	_DHCPInfoGetOptionData func(DictionaryRef, unsafe.Pointer) DataRef
	_SCBondInterfaceCopyAll func(PreferencesRef) ArrayRef
	_SCBondInterfaceCopyAvailableMemberInterfaces func(PreferencesRef) ArrayRef
	_SCBondInterfaceCopyStatus func(BondInterfaceRef) BondStatusRef
	_SCBondInterfaceCreate func(PreferencesRef) BondInterfaceRef
	_SCBondInterfaceGetMemberInterfaces func(BondInterfaceRef) ArrayRef
	_SCBondInterfaceGetOptions func(BondInterfaceRef) DictionaryRef
	_SCBondInterfaceRemove func(BondInterfaceRef) unsafe.Pointer
	_SCBondInterfaceSetLocalizedDisplayName func(BondInterfaceRef, StringRef) unsafe.Pointer
	_SCBondInterfaceSetMemberInterfaces func(BondInterfaceRef, ArrayRef) unsafe.Pointer
	_SCBondInterfaceSetOptions func(BondInterfaceRef, DictionaryRef) unsafe.Pointer
	_SCBondStatusGetInterfaceStatus func(BondStatusRef, NetworkInterfaceRef) DictionaryRef
	_SCBondStatusGetMemberInterfaces func(BondStatusRef) ArrayRef
	_SCBondStatusGetTypeID func() TypeID
	_SCCopyLastError func() ErrorRef
	_SCDynamicStoreAddTemporaryValue func(DynamicStoreRef, StringRef, PropertyListRef) unsafe.Pointer
	_SCDynamicStoreAddValue func(DynamicStoreRef, StringRef, PropertyListRef) unsafe.Pointer
	_SCDynamicStoreCopyComputerName func(DynamicStoreRef, unsafe.Pointer) StringRef
	_SCDynamicStoreCopyConsoleUser func(DynamicStoreRef, unsafe.Pointer, unsafe.Pointer) StringRef
	_SCDynamicStoreCopyDHCPInfo func(DynamicStoreRef, StringRef) DictionaryRef
	_SCDynamicStoreCopyKeyList func(DynamicStoreRef, StringRef) ArrayRef
	_SCDynamicStoreCopyLocalHostName func(DynamicStoreRef) StringRef
	_SCDynamicStoreCopyLocation func(DynamicStoreRef) StringRef
	_SCDynamicStoreCopyMultiple func(DynamicStoreRef, ArrayRef, ArrayRef) DictionaryRef
	_SCDynamicStoreCopyNotifiedKeys func(DynamicStoreRef) ArrayRef
	_SCDynamicStoreCopyProxies func(DynamicStoreRef) DictionaryRef
	_SCDynamicStoreCopyValue func(DynamicStoreRef, StringRef) PropertyListRef
	_SCDynamicStoreCreate func(AllocatorRef, StringRef, DynamicStoreCallBack, unsafe.Pointer) DynamicStoreRef
	_SCDynamicStoreCreateRunLoopSource func(AllocatorRef, DynamicStoreRef, Index) RunLoopSourceRef
	_SCDynamicStoreCreateWithOptions func(AllocatorRef, StringRef, DictionaryRef, DynamicStoreCallBack, unsafe.Pointer) DynamicStoreRef
	_SCDynamicStoreGetTypeID func() TypeID
	_SCDynamicStoreKeyCreate func(AllocatorRef, StringRef) StringRef
	_SCDynamicStoreKeyCreateComputerName func(AllocatorRef) StringRef
	_SCDynamicStoreKeyCreateConsoleUser func(AllocatorRef) StringRef
	_SCDynamicStoreKeyCreateHostNames func(AllocatorRef) StringRef
	_SCDynamicStoreKeyCreateLocation func(AllocatorRef) StringRef
	_SCDynamicStoreKeyCreateNetworkGlobalEntity func(AllocatorRef, StringRef, StringRef) StringRef
	_SCDynamicStoreKeyCreateNetworkInterface func(AllocatorRef, StringRef) StringRef
	_SCDynamicStoreKeyCreateNetworkInterfaceEntity func(AllocatorRef, StringRef, StringRef, StringRef) StringRef
	_SCDynamicStoreKeyCreateNetworkServiceEntity func(AllocatorRef, StringRef, StringRef, StringRef) StringRef
	_SCDynamicStoreKeyCreateProxies func(AllocatorRef) StringRef
	_SCDynamicStoreNotifyValue func(DynamicStoreRef, StringRef) unsafe.Pointer
	_SCDynamicStoreRemoveValue func(DynamicStoreRef, StringRef) unsafe.Pointer
	_SCDynamicStoreSetDispatchQueue func(DynamicStoreRef, unsafe.Pointer) unsafe.Pointer
	_SCDynamicStoreSetMultiple func(DynamicStoreRef, DictionaryRef, ArrayRef, ArrayRef) unsafe.Pointer
	_SCDynamicStoreSetNotificationKeys func(DynamicStoreRef, ArrayRef, ArrayRef) unsafe.Pointer
	_SCDynamicStoreSetValue func(DynamicStoreRef, StringRef, PropertyListRef) unsafe.Pointer
	_SCError func() int
	_SCErrorString func(int) unsafe.Pointer
	_SCNetworkCheckReachabilityByAddress func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkCheckReachabilityByName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionCopyExtendedStatus func(NetworkConnectionRef) DictionaryRef
	_SCNetworkConnectionCopyServiceID func(NetworkConnectionRef) StringRef
	_SCNetworkConnectionCopyStatistics func(NetworkConnectionRef) DictionaryRef
	_SCNetworkConnectionCopyUserOptions func(NetworkConnectionRef) DictionaryRef
	_SCNetworkConnectionCopyUserPreferences func(DictionaryRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionCreateWithServiceID func(AllocatorRef, StringRef, NetworkConnectionCallBack, unsafe.Pointer) NetworkConnectionRef
	_SCNetworkConnectionGetStatus func(NetworkConnectionRef) NetworkConnectionStatus
	_SCNetworkConnectionGetTypeID func() TypeID
	_SCNetworkConnectionScheduleWithRunLoop func(NetworkConnectionRef, RunLoopRef, StringRef) unsafe.Pointer
	_SCNetworkConnectionSetDispatchQueue func(NetworkConnectionRef, unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionStart func(NetworkConnectionRef, DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionStop func(NetworkConnectionRef, unsafe.Pointer) unsafe.Pointer
	_SCNetworkConnectionUnscheduleFromRunLoop func(NetworkConnectionRef, RunLoopRef, StringRef) unsafe.Pointer
	_SCNetworkInterfaceCopyAll func() ArrayRef
	_SCNetworkInterfaceCopyMediaOptions func(NetworkInterfaceRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SCNetworkInterfaceCopyMediaSubTypeOptions func(ArrayRef, StringRef) ArrayRef
	_SCNetworkInterfaceCopyMediaSubTypes func(ArrayRef) ArrayRef
	_SCNetworkInterfaceCopyMTU func(NetworkInterfaceRef, []int, []int, []int) unsafe.Pointer
	_SCNetworkInterfaceCreateWithInterface func(NetworkInterfaceRef, StringRef) NetworkInterfaceRef
	_SCNetworkInterfaceForceConfigurationRefresh func(NetworkInterfaceRef) unsafe.Pointer
	_SCNetworkInterfaceGetBSDName func(NetworkInterfaceRef) StringRef
	_SCNetworkInterfaceGetConfiguration func(NetworkInterfaceRef) DictionaryRef
	_SCNetworkInterfaceGetExtendedConfiguration func(NetworkInterfaceRef, StringRef) DictionaryRef
	_SCNetworkInterfaceGetHardwareAddressString func(NetworkInterfaceRef) StringRef
	_SCNetworkInterfaceGetInterface func(NetworkInterfaceRef) NetworkInterfaceRef
	_SCNetworkInterfaceGetInterfaceType func(NetworkInterfaceRef) StringRef
	_SCNetworkInterfaceGetLocalizedDisplayName func(NetworkInterfaceRef) StringRef
	_SCNetworkInterfaceGetSupportedInterfaceTypes func(NetworkInterfaceRef) ArrayRef
	_SCNetworkInterfaceGetSupportedProtocolTypes func(NetworkInterfaceRef) ArrayRef
	_SCNetworkInterfaceGetTypeID func() TypeID
	_SCNetworkInterfaceRefreshConfiguration func(StringRef) unsafe.Pointer
	_SCNetworkInterfaceSetConfiguration func(NetworkInterfaceRef, DictionaryRef) unsafe.Pointer
	_SCNetworkInterfaceSetExtendedConfiguration func(NetworkInterfaceRef, StringRef, DictionaryRef) unsafe.Pointer
	_SCNetworkInterfaceSetMediaOptions func(NetworkInterfaceRef, StringRef, ArrayRef) unsafe.Pointer
	_SCNetworkInterfaceSetMTU func(NetworkInterfaceRef, int) unsafe.Pointer
	_SCNetworkProtocolGetConfiguration func(NetworkProtocolRef) DictionaryRef
	_SCNetworkProtocolGetEnabled func(NetworkProtocolRef) unsafe.Pointer
	_SCNetworkProtocolGetProtocolType func(NetworkProtocolRef) StringRef
	_SCNetworkProtocolGetTypeID func() TypeID
	_SCNetworkProtocolSetConfiguration func(NetworkProtocolRef, DictionaryRef) unsafe.Pointer
	_SCNetworkProtocolSetEnabled func(NetworkProtocolRef, unsafe.Pointer) unsafe.Pointer
	_SCNetworkReachabilityCreateWithAddress func(AllocatorRef, unsafe.Pointer) NetworkReachabilityRef
	_SCNetworkReachabilityCreateWithAddressPair func(AllocatorRef, unsafe.Pointer, unsafe.Pointer) NetworkReachabilityRef
	_SCNetworkReachabilityCreateWithName func(AllocatorRef, unsafe.Pointer) NetworkReachabilityRef
	_SCNetworkReachabilityGetFlags func(NetworkReachabilityRef, unsafe.Pointer) unsafe.Pointer
	_SCNetworkReachabilityGetTypeID func() TypeID
	_SCNetworkReachabilityScheduleWithRunLoop func(NetworkReachabilityRef, RunLoopRef, StringRef) unsafe.Pointer
	_SCNetworkReachabilitySetCallback func(NetworkReachabilityRef, NetworkReachabilityCallBack, unsafe.Pointer) unsafe.Pointer
	_SCNetworkReachabilitySetDispatchQueue func(NetworkReachabilityRef, unsafe.Pointer) unsafe.Pointer
	_SCNetworkReachabilityUnscheduleFromRunLoop func(NetworkReachabilityRef, RunLoopRef, StringRef) unsafe.Pointer
	_SCNetworkServiceAddProtocolType func(NetworkServiceRef, StringRef) unsafe.Pointer
	_SCNetworkServiceCopy func(PreferencesRef, StringRef) NetworkServiceRef
	_SCNetworkServiceCopyAll func(PreferencesRef) ArrayRef
	_SCNetworkServiceCopyProtocol func(NetworkServiceRef, StringRef) NetworkProtocolRef
	_SCNetworkServiceCopyProtocols func(NetworkServiceRef) ArrayRef
	_SCNetworkServiceCreate func(PreferencesRef, NetworkInterfaceRef) NetworkServiceRef
	_SCNetworkServiceEstablishDefaultConfiguration func(NetworkServiceRef) unsafe.Pointer
	_SCNetworkServiceGetEnabled func(NetworkServiceRef) unsafe.Pointer
	_SCNetworkServiceGetInterface func(NetworkServiceRef) NetworkInterfaceRef
	_SCNetworkServiceGetName func(NetworkServiceRef) StringRef
	_SCNetworkServiceGetServiceID func(NetworkServiceRef) StringRef
	_SCNetworkServiceGetTypeID func() TypeID
	_SCNetworkServiceRemove func(NetworkServiceRef) unsafe.Pointer
	_SCNetworkServiceRemoveProtocolType func(NetworkServiceRef, StringRef) unsafe.Pointer
	_SCNetworkServiceSetEnabled func(NetworkServiceRef, unsafe.Pointer) unsafe.Pointer
	_SCNetworkServiceSetName func(NetworkServiceRef, StringRef) unsafe.Pointer
	_SCNetworkSetAddService func(NetworkSetRef, NetworkServiceRef) unsafe.Pointer
	_SCNetworkSetContainsInterface func(NetworkSetRef, NetworkInterfaceRef) unsafe.Pointer
	_SCNetworkSetCopy func(PreferencesRef, StringRef) NetworkSetRef
	_SCNetworkSetCopyAll func(PreferencesRef) ArrayRef
	_SCNetworkSetCopyCurrent func(PreferencesRef) NetworkSetRef
	_SCNetworkSetCopyServices func(NetworkSetRef) ArrayRef
	_SCNetworkSetCreate func(PreferencesRef) NetworkSetRef
	_SCNetworkSetGetName func(NetworkSetRef) StringRef
	_SCNetworkSetGetServiceOrder func(NetworkSetRef) ArrayRef
	_SCNetworkSetGetSetID func(NetworkSetRef) StringRef
	_SCNetworkSetGetTypeID func() TypeID
	_SCNetworkSetRemove func(NetworkSetRef) unsafe.Pointer
	_SCNetworkSetRemoveService func(NetworkSetRef, NetworkServiceRef) unsafe.Pointer
	_SCNetworkSetSetCurrent func(NetworkSetRef) unsafe.Pointer
	_SCNetworkSetSetName func(NetworkSetRef, StringRef) unsafe.Pointer
	_SCNetworkSetSetServiceOrder func(NetworkSetRef, ArrayRef) unsafe.Pointer
	_SCPreferencesAddValue func(PreferencesRef, StringRef, PropertyListRef) unsafe.Pointer
	_SCPreferencesApplyChanges func(PreferencesRef) unsafe.Pointer
	_SCPreferencesCommitChanges func(PreferencesRef) unsafe.Pointer
	_SCPreferencesCopyKeyList func(PreferencesRef) ArrayRef
	_SCPreferencesCreate func(AllocatorRef, StringRef, StringRef) PreferencesRef
	_SCPreferencesCreateWithAuthorization func(AllocatorRef, StringRef, StringRef, AuthorizationRef) PreferencesRef
	_SCPreferencesGetSignature func(PreferencesRef) DataRef
	_SCPreferencesGetTypeID func() TypeID
	_SCPreferencesGetValue func(PreferencesRef, StringRef) PropertyListRef
	_SCPreferencesLock func(PreferencesRef, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesPathCreateUniqueChild func(PreferencesRef, StringRef) StringRef
	_SCPreferencesPathGetLink func(PreferencesRef, StringRef) StringRef
	_SCPreferencesPathGetValue func(PreferencesRef, StringRef) DictionaryRef
	_SCPreferencesPathRemoveValue func(PreferencesRef, StringRef) unsafe.Pointer
	_SCPreferencesPathSetLink func(PreferencesRef, StringRef, StringRef) unsafe.Pointer
	_SCPreferencesPathSetValue func(PreferencesRef, StringRef, DictionaryRef) unsafe.Pointer
	_SCPreferencesRemoveValue func(PreferencesRef, StringRef) unsafe.Pointer
	_SCPreferencesScheduleWithRunLoop func(PreferencesRef, RunLoopRef, StringRef) unsafe.Pointer
	_SCPreferencesSetCallback func(PreferencesRef, PreferencesCallBack, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesSetComputerName func(PreferencesRef, StringRef, StringEncoding) unsafe.Pointer
	_SCPreferencesSetDispatchQueue func(PreferencesRef, unsafe.Pointer) unsafe.Pointer
	_SCPreferencesSetLocalHostName func(PreferencesRef, StringRef) unsafe.Pointer
	_SCPreferencesSetValue func(PreferencesRef, StringRef, PropertyListRef) unsafe.Pointer
	_SCPreferencesSynchronize func(PreferencesRef)
	_SCPreferencesUnlock func(PreferencesRef) unsafe.Pointer
	_SCPreferencesUnscheduleFromRunLoop func(PreferencesRef, RunLoopRef, StringRef) unsafe.Pointer
	_SCVLANInterfaceCopyAll func(PreferencesRef) ArrayRef
	_SCVLANInterfaceCopyAvailablePhysicalInterfaces func() ArrayRef
	_SCVLANInterfaceCreate func(PreferencesRef, NetworkInterfaceRef, NumberRef) VLANInterfaceRef
	_SCVLANInterfaceGetOptions func(VLANInterfaceRef) DictionaryRef
	_SCVLANInterfaceGetPhysicalInterface func(VLANInterfaceRef) NetworkInterfaceRef
	_SCVLANInterfaceGetTag func(VLANInterfaceRef) NumberRef
	_SCVLANInterfaceRemove func(VLANInterfaceRef) unsafe.Pointer
	_SCVLANInterfaceSetLocalizedDisplayName func(VLANInterfaceRef, StringRef) unsafe.Pointer
	_SCVLANInterfaceSetOptions func(VLANInterfaceRef, DictionaryRef) unsafe.Pointer
	_SCVLANInterfaceSetPhysicalInterfaceAndTag func(VLANInterfaceRef, NetworkInterfaceRef, NumberRef) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CNCopyCurrentNetworkInfo, lib, "CNCopyCurrentNetworkInfo")
	tryRegister(&_CNCopySupportedInterfaces, lib, "CNCopySupportedInterfaces")
	tryRegister(&_CNMarkPortalOffline, lib, "CNMarkPortalOffline")
	tryRegister(&_CNMarkPortalOnline, lib, "CNMarkPortalOnline")
	tryRegister(&_CNSetSupportedSSIDs, lib, "CNSetSupportedSSIDs")
	tryRegister(&_DHCPClientPreferencesCopyApplicationOptions, lib, "DHCPClientPreferencesCopyApplicationOptions")
	tryRegister(&_DHCPClientPreferencesSetApplicationOptions, lib, "DHCPClientPreferencesSetApplicationOptions")
	tryRegister(&_DHCPInfoGetLeaseExpirationTime, lib, "DHCPInfoGetLeaseExpirationTime")
	tryRegister(&_DHCPInfoGetLeaseStartTime, lib, "DHCPInfoGetLeaseStartTime")
	tryRegister(&_DHCPInfoGetOptionData, lib, "DHCPInfoGetOptionData")
	tryRegister(&_SCBondInterfaceCopyAll, lib, "SCBondInterfaceCopyAll")
	tryRegister(&_SCBondInterfaceCopyAvailableMemberInterfaces, lib, "SCBondInterfaceCopyAvailableMemberInterfaces")
	tryRegister(&_SCBondInterfaceCopyStatus, lib, "SCBondInterfaceCopyStatus")
	tryRegister(&_SCBondInterfaceCreate, lib, "SCBondInterfaceCreate")
	tryRegister(&_SCBondInterfaceGetMemberInterfaces, lib, "SCBondInterfaceGetMemberInterfaces")
	tryRegister(&_SCBondInterfaceGetOptions, lib, "SCBondInterfaceGetOptions")
	tryRegister(&_SCBondInterfaceRemove, lib, "SCBondInterfaceRemove")
	tryRegister(&_SCBondInterfaceSetLocalizedDisplayName, lib, "SCBondInterfaceSetLocalizedDisplayName")
	tryRegister(&_SCBondInterfaceSetMemberInterfaces, lib, "SCBondInterfaceSetMemberInterfaces")
	tryRegister(&_SCBondInterfaceSetOptions, lib, "SCBondInterfaceSetOptions")
	tryRegister(&_SCBondStatusGetInterfaceStatus, lib, "SCBondStatusGetInterfaceStatus")
	tryRegister(&_SCBondStatusGetMemberInterfaces, lib, "SCBondStatusGetMemberInterfaces")
	tryRegister(&_SCBondStatusGetTypeID, lib, "SCBondStatusGetTypeID")
	tryRegister(&_SCCopyLastError, lib, "SCCopyLastError")
	tryRegister(&_SCDynamicStoreAddTemporaryValue, lib, "SCDynamicStoreAddTemporaryValue")
	tryRegister(&_SCDynamicStoreAddValue, lib, "SCDynamicStoreAddValue")
	tryRegister(&_SCDynamicStoreCopyComputerName, lib, "SCDynamicStoreCopyComputerName")
	tryRegister(&_SCDynamicStoreCopyConsoleUser, lib, "SCDynamicStoreCopyConsoleUser")
	tryRegister(&_SCDynamicStoreCopyDHCPInfo, lib, "SCDynamicStoreCopyDHCPInfo")
	tryRegister(&_SCDynamicStoreCopyKeyList, lib, "SCDynamicStoreCopyKeyList")
	tryRegister(&_SCDynamicStoreCopyLocalHostName, lib, "SCDynamicStoreCopyLocalHostName")
	tryRegister(&_SCDynamicStoreCopyLocation, lib, "SCDynamicStoreCopyLocation")
	tryRegister(&_SCDynamicStoreCopyMultiple, lib, "SCDynamicStoreCopyMultiple")
	tryRegister(&_SCDynamicStoreCopyNotifiedKeys, lib, "SCDynamicStoreCopyNotifiedKeys")
	tryRegister(&_SCDynamicStoreCopyProxies, lib, "SCDynamicStoreCopyProxies")
	tryRegister(&_SCDynamicStoreCopyValue, lib, "SCDynamicStoreCopyValue")
	tryRegister(&_SCDynamicStoreCreate, lib, "SCDynamicStoreCreate")
	tryRegister(&_SCDynamicStoreCreateRunLoopSource, lib, "SCDynamicStoreCreateRunLoopSource")
	tryRegister(&_SCDynamicStoreCreateWithOptions, lib, "SCDynamicStoreCreateWithOptions")
	tryRegister(&_SCDynamicStoreGetTypeID, lib, "SCDynamicStoreGetTypeID")
	tryRegister(&_SCDynamicStoreKeyCreate, lib, "SCDynamicStoreKeyCreate")
	tryRegister(&_SCDynamicStoreKeyCreateComputerName, lib, "SCDynamicStoreKeyCreateComputerName")
	tryRegister(&_SCDynamicStoreKeyCreateConsoleUser, lib, "SCDynamicStoreKeyCreateConsoleUser")
	tryRegister(&_SCDynamicStoreKeyCreateHostNames, lib, "SCDynamicStoreKeyCreateHostNames")
	tryRegister(&_SCDynamicStoreKeyCreateLocation, lib, "SCDynamicStoreKeyCreateLocation")
	tryRegister(&_SCDynamicStoreKeyCreateNetworkGlobalEntity, lib, "SCDynamicStoreKeyCreateNetworkGlobalEntity")
	tryRegister(&_SCDynamicStoreKeyCreateNetworkInterface, lib, "SCDynamicStoreKeyCreateNetworkInterface")
	tryRegister(&_SCDynamicStoreKeyCreateNetworkInterfaceEntity, lib, "SCDynamicStoreKeyCreateNetworkInterfaceEntity")
	tryRegister(&_SCDynamicStoreKeyCreateNetworkServiceEntity, lib, "SCDynamicStoreKeyCreateNetworkServiceEntity")
	tryRegister(&_SCDynamicStoreKeyCreateProxies, lib, "SCDynamicStoreKeyCreateProxies")
	tryRegister(&_SCDynamicStoreNotifyValue, lib, "SCDynamicStoreNotifyValue")
	tryRegister(&_SCDynamicStoreRemoveValue, lib, "SCDynamicStoreRemoveValue")
	tryRegister(&_SCDynamicStoreSetDispatchQueue, lib, "SCDynamicStoreSetDispatchQueue")
	tryRegister(&_SCDynamicStoreSetMultiple, lib, "SCDynamicStoreSetMultiple")
	tryRegister(&_SCDynamicStoreSetNotificationKeys, lib, "SCDynamicStoreSetNotificationKeys")
	tryRegister(&_SCDynamicStoreSetValue, lib, "SCDynamicStoreSetValue")
	tryRegister(&_SCError, lib, "SCError")
	tryRegister(&_SCErrorString, lib, "SCErrorString")
	tryRegister(&_SCNetworkCheckReachabilityByAddress, lib, "SCNetworkCheckReachabilityByAddress")
	tryRegister(&_SCNetworkCheckReachabilityByName, lib, "SCNetworkCheckReachabilityByName")
	tryRegister(&_SCNetworkConnectionCopyExtendedStatus, lib, "SCNetworkConnectionCopyExtendedStatus")
	tryRegister(&_SCNetworkConnectionCopyServiceID, lib, "SCNetworkConnectionCopyServiceID")
	tryRegister(&_SCNetworkConnectionCopyStatistics, lib, "SCNetworkConnectionCopyStatistics")
	tryRegister(&_SCNetworkConnectionCopyUserOptions, lib, "SCNetworkConnectionCopyUserOptions")
	tryRegister(&_SCNetworkConnectionCopyUserPreferences, lib, "SCNetworkConnectionCopyUserPreferences")
	tryRegister(&_SCNetworkConnectionCreateWithServiceID, lib, "SCNetworkConnectionCreateWithServiceID")
	tryRegister(&_SCNetworkConnectionGetStatus, lib, "SCNetworkConnectionGetStatus")
	tryRegister(&_SCNetworkConnectionGetTypeID, lib, "SCNetworkConnectionGetTypeID")
	tryRegister(&_SCNetworkConnectionScheduleWithRunLoop, lib, "SCNetworkConnectionScheduleWithRunLoop")
	tryRegister(&_SCNetworkConnectionSetDispatchQueue, lib, "SCNetworkConnectionSetDispatchQueue")
	tryRegister(&_SCNetworkConnectionStart, lib, "SCNetworkConnectionStart")
	tryRegister(&_SCNetworkConnectionStop, lib, "SCNetworkConnectionStop")
	tryRegister(&_SCNetworkConnectionUnscheduleFromRunLoop, lib, "SCNetworkConnectionUnscheduleFromRunLoop")
	tryRegister(&_SCNetworkInterfaceCopyAll, lib, "SCNetworkInterfaceCopyAll")
	tryRegister(&_SCNetworkInterfaceCopyMediaOptions, lib, "SCNetworkInterfaceCopyMediaOptions")
	tryRegister(&_SCNetworkInterfaceCopyMediaSubTypeOptions, lib, "SCNetworkInterfaceCopyMediaSubTypeOptions")
	tryRegister(&_SCNetworkInterfaceCopyMediaSubTypes, lib, "SCNetworkInterfaceCopyMediaSubTypes")
	tryRegister(&_SCNetworkInterfaceCopyMTU, lib, "SCNetworkInterfaceCopyMTU")
	tryRegister(&_SCNetworkInterfaceCreateWithInterface, lib, "SCNetworkInterfaceCreateWithInterface")
	tryRegister(&_SCNetworkInterfaceForceConfigurationRefresh, lib, "SCNetworkInterfaceForceConfigurationRefresh")
	tryRegister(&_SCNetworkInterfaceGetBSDName, lib, "SCNetworkInterfaceGetBSDName")
	tryRegister(&_SCNetworkInterfaceGetConfiguration, lib, "SCNetworkInterfaceGetConfiguration")
	tryRegister(&_SCNetworkInterfaceGetExtendedConfiguration, lib, "SCNetworkInterfaceGetExtendedConfiguration")
	tryRegister(&_SCNetworkInterfaceGetHardwareAddressString, lib, "SCNetworkInterfaceGetHardwareAddressString")
	tryRegister(&_SCNetworkInterfaceGetInterface, lib, "SCNetworkInterfaceGetInterface")
	tryRegister(&_SCNetworkInterfaceGetInterfaceType, lib, "SCNetworkInterfaceGetInterfaceType")
	tryRegister(&_SCNetworkInterfaceGetLocalizedDisplayName, lib, "SCNetworkInterfaceGetLocalizedDisplayName")
	tryRegister(&_SCNetworkInterfaceGetSupportedInterfaceTypes, lib, "SCNetworkInterfaceGetSupportedInterfaceTypes")
	tryRegister(&_SCNetworkInterfaceGetSupportedProtocolTypes, lib, "SCNetworkInterfaceGetSupportedProtocolTypes")
	tryRegister(&_SCNetworkInterfaceGetTypeID, lib, "SCNetworkInterfaceGetTypeID")
	tryRegister(&_SCNetworkInterfaceRefreshConfiguration, lib, "SCNetworkInterfaceRefreshConfiguration")
	tryRegister(&_SCNetworkInterfaceSetConfiguration, lib, "SCNetworkInterfaceSetConfiguration")
	tryRegister(&_SCNetworkInterfaceSetExtendedConfiguration, lib, "SCNetworkInterfaceSetExtendedConfiguration")
	tryRegister(&_SCNetworkInterfaceSetMediaOptions, lib, "SCNetworkInterfaceSetMediaOptions")
	tryRegister(&_SCNetworkInterfaceSetMTU, lib, "SCNetworkInterfaceSetMTU")
	tryRegister(&_SCNetworkProtocolGetConfiguration, lib, "SCNetworkProtocolGetConfiguration")
	tryRegister(&_SCNetworkProtocolGetEnabled, lib, "SCNetworkProtocolGetEnabled")
	tryRegister(&_SCNetworkProtocolGetProtocolType, lib, "SCNetworkProtocolGetProtocolType")
	tryRegister(&_SCNetworkProtocolGetTypeID, lib, "SCNetworkProtocolGetTypeID")
	tryRegister(&_SCNetworkProtocolSetConfiguration, lib, "SCNetworkProtocolSetConfiguration")
	tryRegister(&_SCNetworkProtocolSetEnabled, lib, "SCNetworkProtocolSetEnabled")
	tryRegister(&_SCNetworkReachabilityCreateWithAddress, lib, "SCNetworkReachabilityCreateWithAddress")
	tryRegister(&_SCNetworkReachabilityCreateWithAddressPair, lib, "SCNetworkReachabilityCreateWithAddressPair")
	tryRegister(&_SCNetworkReachabilityCreateWithName, lib, "SCNetworkReachabilityCreateWithName")
	tryRegister(&_SCNetworkReachabilityGetFlags, lib, "SCNetworkReachabilityGetFlags")
	tryRegister(&_SCNetworkReachabilityGetTypeID, lib, "SCNetworkReachabilityGetTypeID")
	tryRegister(&_SCNetworkReachabilityScheduleWithRunLoop, lib, "SCNetworkReachabilityScheduleWithRunLoop")
	tryRegister(&_SCNetworkReachabilitySetCallback, lib, "SCNetworkReachabilitySetCallback")
	tryRegister(&_SCNetworkReachabilitySetDispatchQueue, lib, "SCNetworkReachabilitySetDispatchQueue")
	tryRegister(&_SCNetworkReachabilityUnscheduleFromRunLoop, lib, "SCNetworkReachabilityUnscheduleFromRunLoop")
	tryRegister(&_SCNetworkServiceAddProtocolType, lib, "SCNetworkServiceAddProtocolType")
	tryRegister(&_SCNetworkServiceCopy, lib, "SCNetworkServiceCopy")
	tryRegister(&_SCNetworkServiceCopyAll, lib, "SCNetworkServiceCopyAll")
	tryRegister(&_SCNetworkServiceCopyProtocol, lib, "SCNetworkServiceCopyProtocol")
	tryRegister(&_SCNetworkServiceCopyProtocols, lib, "SCNetworkServiceCopyProtocols")
	tryRegister(&_SCNetworkServiceCreate, lib, "SCNetworkServiceCreate")
	tryRegister(&_SCNetworkServiceEstablishDefaultConfiguration, lib, "SCNetworkServiceEstablishDefaultConfiguration")
	tryRegister(&_SCNetworkServiceGetEnabled, lib, "SCNetworkServiceGetEnabled")
	tryRegister(&_SCNetworkServiceGetInterface, lib, "SCNetworkServiceGetInterface")
	tryRegister(&_SCNetworkServiceGetName, lib, "SCNetworkServiceGetName")
	tryRegister(&_SCNetworkServiceGetServiceID, lib, "SCNetworkServiceGetServiceID")
	tryRegister(&_SCNetworkServiceGetTypeID, lib, "SCNetworkServiceGetTypeID")
	tryRegister(&_SCNetworkServiceRemove, lib, "SCNetworkServiceRemove")
	tryRegister(&_SCNetworkServiceRemoveProtocolType, lib, "SCNetworkServiceRemoveProtocolType")
	tryRegister(&_SCNetworkServiceSetEnabled, lib, "SCNetworkServiceSetEnabled")
	tryRegister(&_SCNetworkServiceSetName, lib, "SCNetworkServiceSetName")
	tryRegister(&_SCNetworkSetAddService, lib, "SCNetworkSetAddService")
	tryRegister(&_SCNetworkSetContainsInterface, lib, "SCNetworkSetContainsInterface")
	tryRegister(&_SCNetworkSetCopy, lib, "SCNetworkSetCopy")
	tryRegister(&_SCNetworkSetCopyAll, lib, "SCNetworkSetCopyAll")
	tryRegister(&_SCNetworkSetCopyCurrent, lib, "SCNetworkSetCopyCurrent")
	tryRegister(&_SCNetworkSetCopyServices, lib, "SCNetworkSetCopyServices")
	tryRegister(&_SCNetworkSetCreate, lib, "SCNetworkSetCreate")
	tryRegister(&_SCNetworkSetGetName, lib, "SCNetworkSetGetName")
	tryRegister(&_SCNetworkSetGetServiceOrder, lib, "SCNetworkSetGetServiceOrder")
	tryRegister(&_SCNetworkSetGetSetID, lib, "SCNetworkSetGetSetID")
	tryRegister(&_SCNetworkSetGetTypeID, lib, "SCNetworkSetGetTypeID")
	tryRegister(&_SCNetworkSetRemove, lib, "SCNetworkSetRemove")
	tryRegister(&_SCNetworkSetRemoveService, lib, "SCNetworkSetRemoveService")
	tryRegister(&_SCNetworkSetSetCurrent, lib, "SCNetworkSetSetCurrent")
	tryRegister(&_SCNetworkSetSetName, lib, "SCNetworkSetSetName")
	tryRegister(&_SCNetworkSetSetServiceOrder, lib, "SCNetworkSetSetServiceOrder")
	tryRegister(&_SCPreferencesAddValue, lib, "SCPreferencesAddValue")
	tryRegister(&_SCPreferencesApplyChanges, lib, "SCPreferencesApplyChanges")
	tryRegister(&_SCPreferencesCommitChanges, lib, "SCPreferencesCommitChanges")
	tryRegister(&_SCPreferencesCopyKeyList, lib, "SCPreferencesCopyKeyList")
	tryRegister(&_SCPreferencesCreate, lib, "SCPreferencesCreate")
	tryRegister(&_SCPreferencesCreateWithAuthorization, lib, "SCPreferencesCreateWithAuthorization")
	tryRegister(&_SCPreferencesGetSignature, lib, "SCPreferencesGetSignature")
	tryRegister(&_SCPreferencesGetTypeID, lib, "SCPreferencesGetTypeID")
	tryRegister(&_SCPreferencesGetValue, lib, "SCPreferencesGetValue")
	tryRegister(&_SCPreferencesLock, lib, "SCPreferencesLock")
	tryRegister(&_SCPreferencesPathCreateUniqueChild, lib, "SCPreferencesPathCreateUniqueChild")
	tryRegister(&_SCPreferencesPathGetLink, lib, "SCPreferencesPathGetLink")
	tryRegister(&_SCPreferencesPathGetValue, lib, "SCPreferencesPathGetValue")
	tryRegister(&_SCPreferencesPathRemoveValue, lib, "SCPreferencesPathRemoveValue")
	tryRegister(&_SCPreferencesPathSetLink, lib, "SCPreferencesPathSetLink")
	tryRegister(&_SCPreferencesPathSetValue, lib, "SCPreferencesPathSetValue")
	tryRegister(&_SCPreferencesRemoveValue, lib, "SCPreferencesRemoveValue")
	tryRegister(&_SCPreferencesScheduleWithRunLoop, lib, "SCPreferencesScheduleWithRunLoop")
	tryRegister(&_SCPreferencesSetCallback, lib, "SCPreferencesSetCallback")
	tryRegister(&_SCPreferencesSetComputerName, lib, "SCPreferencesSetComputerName")
	tryRegister(&_SCPreferencesSetDispatchQueue, lib, "SCPreferencesSetDispatchQueue")
	tryRegister(&_SCPreferencesSetLocalHostName, lib, "SCPreferencesSetLocalHostName")
	tryRegister(&_SCPreferencesSetValue, lib, "SCPreferencesSetValue")
	tryRegister(&_SCPreferencesSynchronize, lib, "SCPreferencesSynchronize")
	tryRegister(&_SCPreferencesUnlock, lib, "SCPreferencesUnlock")
	tryRegister(&_SCPreferencesUnscheduleFromRunLoop, lib, "SCPreferencesUnscheduleFromRunLoop")
	tryRegister(&_SCVLANInterfaceCopyAll, lib, "SCVLANInterfaceCopyAll")
	tryRegister(&_SCVLANInterfaceCopyAvailablePhysicalInterfaces, lib, "SCVLANInterfaceCopyAvailablePhysicalInterfaces")
	tryRegister(&_SCVLANInterfaceCreate, lib, "SCVLANInterfaceCreate")
	tryRegister(&_SCVLANInterfaceGetOptions, lib, "SCVLANInterfaceGetOptions")
	tryRegister(&_SCVLANInterfaceGetPhysicalInterface, lib, "SCVLANInterfaceGetPhysicalInterface")
	tryRegister(&_SCVLANInterfaceGetTag, lib, "SCVLANInterfaceGetTag")
	tryRegister(&_SCVLANInterfaceRemove, lib, "SCVLANInterfaceRemove")
	tryRegister(&_SCVLANInterfaceSetLocalizedDisplayName, lib, "SCVLANInterfaceSetLocalizedDisplayName")
	tryRegister(&_SCVLANInterfaceSetOptions, lib, "SCVLANInterfaceSetOptions")
	tryRegister(&_SCVLANInterfaceSetPhysicalInterfaceAndTag, lib, "SCVLANInterfaceSetPhysicalInterfaceAndTag")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Returns the current network information for a given network interface.

// Returns the current network information for a given network interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/CNCopyCurrentNetworkInfo
func CNCopyCurrentNetworkInfo(interfaceName StringRef) DictionaryRef {
	return _CNCopyCurrentNetworkInfo(interfaceName)
}/* debug [functions.gen.go/function]: CNCopyCurrentNetworkInfo */

// Returns the names of all network interfaces Captive Network Support is monitoring.
//
// Added in macOS 10.8.
// Returns the names of all network interfaces Captive Network Support is monitoring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/CNCopySupportedInterfaces
func CNCopySupportedInterfaces() ArrayRef {
	return _CNCopySupportedInterfaces()
}/* debug [functions.gen.go/function]: CNCopySupportedInterfaces */

// Informs Captive Network Support that the device is not authenticated on a captive network.
//
// Added in macOS 10.8.
// Informs Captive Network Support that the device is not authenticated on a captive network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/CNMarkPortalOffline
func CNMarkPortalOffline(interfaceName StringRef) unsafe.Pointer {
	return _CNMarkPortalOffline(interfaceName)
}/* debug [functions.gen.go/function]: CNMarkPortalOffline */

// Informs Captive Network Support that the application has successfully authenticated the device to a captive network. Captive Network Support notifies the rest of the system that WiFi is a viable interface.
//
// Added in macOS 10.8.
// Informs Captive Network Support that the application has successfully authenticated the device to a captive network. Captive Network Support notifies the rest of the system that WiFi is a viable interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/CNMarkPortalOnline
func CNMarkPortalOnline(interfaceName StringRef) unsafe.Pointer {
	return _CNMarkPortalOnline(interfaceName)
}/* debug [functions.gen.go/function]: CNMarkPortalOnline */

// Specifies an updated list of captive network SSIDs that the application performs authentication on.
//
// Added in macOS 10.8.
// Specifies an updated list of captive network SSIDs that the application performs authentication on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/CNSetSupportedSSIDs
func CNSetSupportedSSIDs(ssidArray ArrayRef) unsafe.Pointer {
	return _CNSetSupportedSSIDs(ssidArray)
}/* debug [functions.gen.go/function]: CNSetSupportedSSIDs */

// Returns the list of options for the specified application ID.
//
// Added in macOS 10.1.
// Returns the list of options for the specified application ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/DHCPClientPreferencesCopyApplicationOptions
func DHCPClientPreferencesCopyApplicationOptions(applicationID StringRef, count unsafe.Pointer) unsafe.Pointer {
	return _DHCPClientPreferencesCopyApplicationOptions(applicationID, count)
}/* debug [functions.gen.go/function]: DHCPClientPreferencesCopyApplicationOptions */

// Updates the DHCP client preferences to include the specified list of options for the specified application ID.
//
// Added in macOS 10.1.
// Updates the DHCP client preferences to include the specified list of options for the specified application ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/DHCPClientPreferencesSetApplicationOptions
func DHCPClientPreferencesSetApplicationOptions(applicationID StringRef, options unsafe.Pointer, count Index) unsafe.Pointer {
	return _DHCPClientPreferencesSetApplicationOptions(applicationID, options, count)
}/* debug [functions.gen.go/function]: DHCPClientPreferencesSetApplicationOptions */

// Returns the lease expiration time data.
//
// Added in macOS 10.8.
// Returns the lease expiration time data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/DHCPInfoGetLeaseExpirationTime
func DHCPInfoGetLeaseExpirationTime(info DictionaryRef) DateRef {
	return _DHCPInfoGetLeaseExpirationTime(info)
}/* debug [functions.gen.go/function]: DHCPInfoGetLeaseExpirationTime */

// Returns the lease start time data.
//
// Added in macOS 10.1.
// Returns the lease start time data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/DHCPInfoGetLeaseStartTime
func DHCPInfoGetLeaseStartTime(info DictionaryRef) DateRef {
	return _DHCPInfoGetLeaseStartTime(info)
}/* debug [functions.gen.go/function]: DHCPInfoGetLeaseStartTime */

// Returns DHCP option data, if present.
//
// Added in macOS 10.1.
// Returns DHCP option data, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/DHCPInfoGetOptionData
func DHCPInfoGetOptionData(info DictionaryRef, code unsafe.Pointer) DataRef {
	return _DHCPInfoGetOptionData(info, code)
}/* debug [functions.gen.go/function]: DHCPInfoGetOptionData */

// Returns all Ethernet bond interfaces on the system.
//
// Added in macOS 10.5.
// Returns all Ethernet bond interfaces on the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceCopyAll(_:)
func SCBondInterfaceCopyAll(prefs PreferencesRef) ArrayRef {
	return _SCBondInterfaceCopyAll(prefs)
}/* debug [functions.gen.go/function]: SCBondInterfaceCopyAll */

// Returns all network capable devices on the system that can be added to an Ethernet bond interface.
//
// Added in macOS 10.5.
// Returns all network capable devices on the system that can be added to an Ethernet bond interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceCopyAvailableMemberInterfaces(_:)
func SCBondInterfaceCopyAvailableMemberInterfaces(prefs PreferencesRef) ArrayRef {
	return _SCBondInterfaceCopyAvailableMemberInterfaces(prefs)
}/* debug [functions.gen.go/function]: SCBondInterfaceCopyAvailableMemberInterfaces */

// Returns the status of the specified Ethernet bond interface.
//
// Added in macOS 10.5.
// Returns the status of the specified Ethernet bond interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceCopyStatus(_:)
func SCBondInterfaceCopyStatus(bond BondInterfaceRef) BondStatusRef {
	return _SCBondInterfaceCopyStatus(bond)
}/* debug [functions.gen.go/function]: SCBondInterfaceCopyStatus */

// Creates a new Ethernet bond interface.
//
// Added in macOS 10.5.
// Creates a new Ethernet bond interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceCreate(_:)
func SCBondInterfaceCreate(prefs PreferencesRef) BondInterfaceRef {
	return _SCBondInterfaceCreate(prefs)
}/* debug [functions.gen.go/function]: SCBondInterfaceCreate */

// Returns the member interfaces for the specified Ethernet bond interface.
//
// Added in macOS 10.5.
// Returns the member interfaces for the specified Ethernet bond interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceGetMemberInterfaces(_:)
func SCBondInterfaceGetMemberInterfaces(bond BondInterfaceRef) ArrayRef {
	return _SCBondInterfaceGetMemberInterfaces(bond)
}/* debug [functions.gen.go/function]: SCBondInterfaceGetMemberInterfaces */

// Returns the configuration settings associated with the specified Ethernet bond interface.
//
// Added in macOS 10.5.
// Returns the configuration settings associated with the specified Ethernet bond interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceGetOptions(_:)
func SCBondInterfaceGetOptions(bond BondInterfaceRef) DictionaryRef {
	return _SCBondInterfaceGetOptions(bond)
}/* debug [functions.gen.go/function]: SCBondInterfaceGetOptions */

// Removes the Ethernet bond interface from the configuration.
//
// Added in macOS 10.5.
// Removes the Ethernet bond interface from the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceRemove(_:)
func SCBondInterfaceRemove(bond BondInterfaceRef) unsafe.Pointer {
	return _SCBondInterfaceRemove(bond)
}/* debug [functions.gen.go/function]: SCBondInterfaceRemove */

// Sets the localized display name for the specified Ethernet bond interface.
//
// Added in macOS 10.5.
// Sets the localized display name for the specified Ethernet bond interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceSetLocalizedDisplayName(_:_:)
func SCBondInterfaceSetLocalizedDisplayName(bond BondInterfaceRef, newName StringRef) unsafe.Pointer {
	return _SCBondInterfaceSetLocalizedDisplayName(bond, newName)
}/* debug [functions.gen.go/function]: SCBondInterfaceSetLocalizedDisplayName */

// Sets the member interfaces for the specified Ethernet bond interface.
//
// Added in macOS 10.5.
// Sets the member interfaces for the specified Ethernet bond interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceSetMemberInterfaces(_:_:)
func SCBondInterfaceSetMemberInterfaces(bond BondInterfaceRef, members ArrayRef) unsafe.Pointer {
	return _SCBondInterfaceSetMemberInterfaces(bond, members)
}/* debug [functions.gen.go/function]: SCBondInterfaceSetMemberInterfaces */

// Sets the configuration settings for the specified Ethernet bond interface.
//
// Added in macOS 10.5.
// Sets the configuration settings for the specified Ethernet bond interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondInterfaceSetOptions(_:_:)
func SCBondInterfaceSetOptions(bond BondInterfaceRef, newOptions DictionaryRef) unsafe.Pointer {
	return _SCBondInterfaceSetOptions(bond, newOptions)
}/* debug [functions.gen.go/function]: SCBondInterfaceSetOptions */

// Returns the status of the specified member interface of an Ethernet bond or the status of the bond as a whole.
//
// Added in macOS 10.5.
// Returns the status of the specified member interface of an Ethernet bond or the status of the bond as a whole.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondStatusGetInterfaceStatus(_:_:)
func SCBondStatusGetInterfaceStatus(bondStatus BondStatusRef, interface_ NetworkInterfaceRef) DictionaryRef {
	return _SCBondStatusGetInterfaceStatus(bondStatus, interface_)
}/* debug [functions.gen.go/function]: SCBondStatusGetInterfaceStatus */

// Returns the member interfaces that are represented with the Ethernet bond interface.
//
// Added in macOS 10.5.
// Returns the member interfaces that are represented with the Ethernet bond interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondStatusGetMemberInterfaces(_:)
func SCBondStatusGetMemberInterfaces(bondStatus BondStatusRef) ArrayRef {
	return _SCBondStatusGetMemberInterfaces(bondStatus)
}/* debug [functions.gen.go/function]: SCBondStatusGetMemberInterfaces */

// Returns the type identifier of all instances.
//
// Added in macOS 10.5.
// Returns the type identifier of all instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCBondStatusGetTypeID()
func SCBondStatusGetTypeID() TypeID {
	return _SCBondStatusGetTypeID()
}/* debug [functions.gen.go/function]: SCBondStatusGetTypeID */

// Returns an error or status code associated with the most recent function call.
//
// Added in macOS 10.5.
// Returns an error or status code associated with the most recent function call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCCopyLastError()
func SCCopyLastError() ErrorRef {
	return _SCCopyLastError()
}/* debug [functions.gen.go/function]: SCCopyLastError */

// Temporarily adds the specified key-value pair to the dynamic store, if no such key already exists.
//
// Added in macOS 10.1.
// Temporarily adds the specified key-value pair to the dynamic store, if no such key already exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreAddTemporaryValue(_:_:_:)
func SCDynamicStoreAddTemporaryValue(store DynamicStoreRef, key StringRef, value PropertyListRef) unsafe.Pointer {
	return _SCDynamicStoreAddTemporaryValue(store, key, value)
}/* debug [functions.gen.go/function]: SCDynamicStoreAddTemporaryValue */

// Adds the specified key-value pair to the dynamic store, if no such key already exists.
//
// Added in macOS 10.1.
// Adds the specified key-value pair to the dynamic store, if no such key already exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreAddValue(_:_:_:)
func SCDynamicStoreAddValue(store DynamicStoreRef, key StringRef, value PropertyListRef) unsafe.Pointer {
	return _SCDynamicStoreAddValue(store, key, value)
}/* debug [functions.gen.go/function]: SCDynamicStoreAddValue */

// Returns the current computer name.
//
// Added in macOS 10.1.
// Returns the current computer name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyComputerName(_:_:)
func SCDynamicStoreCopyComputerName(store DynamicStoreRef, nameEncoding unsafe.Pointer) StringRef {
	return _SCDynamicStoreCopyComputerName(store, nameEncoding)
}/* debug [functions.gen.go/function]: SCDynamicStoreCopyComputerName */

// Returns information about the user currently logged into the system.
//
// Added in macOS 10.1.
// Returns information about the user currently logged into the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyConsoleUser(_:_:_:)
func SCDynamicStoreCopyConsoleUser(store DynamicStoreRef, uid unsafe.Pointer, gid unsafe.Pointer) StringRef {
	return _SCDynamicStoreCopyConsoleUser(store, uid, gid)
}/* debug [functions.gen.go/function]: SCDynamicStoreCopyConsoleUser */

// Returns the DHCP information for the specified service.
//
// Added in macOS 10.1.
// Returns the DHCP information for the specified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyDHCPInfo
func SCDynamicStoreCopyDHCPInfo(store DynamicStoreRef, serviceID StringRef) DictionaryRef {
	return _SCDynamicStoreCopyDHCPInfo(store, serviceID)
}/* debug [functions.gen.go/function]: SCDynamicStoreCopyDHCPInfo */

// Returns the keys that represent the current dynamic store entries that match the specified pattern.
//
// Added in macOS 10.1.
// Returns the keys that represent the current dynamic store entries that match the specified pattern.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyKeyList(_:_:)
func SCDynamicStoreCopyKeyList(store DynamicStoreRef, pattern StringRef) ArrayRef {
	return _SCDynamicStoreCopyKeyList(store, pattern)
}/* debug [functions.gen.go/function]: SCDynamicStoreCopyKeyList */

// Returns the current local host name.
//
// Added in macOS 10.1.
// Returns the current local host name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyLocalHostName(_:)
func SCDynamicStoreCopyLocalHostName(store DynamicStoreRef) StringRef {
	return _SCDynamicStoreCopyLocalHostName(store)
}/* debug [functions.gen.go/function]: SCDynamicStoreCopyLocalHostName */

// Returns the current location identifier.
//
// Added in macOS 10.1.
// Returns the current location identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyLocation(_:)
func SCDynamicStoreCopyLocation(store DynamicStoreRef) StringRef {
	return _SCDynamicStoreCopyLocation(store)
}/* debug [functions.gen.go/function]: SCDynamicStoreCopyLocation */

// Returns the key-value pairs that match the specified keys and key patterns.
//
// Added in macOS 10.1.
// Returns the key-value pairs that match the specified keys and key patterns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyMultiple(_:_:_:)
func SCDynamicStoreCopyMultiple(store DynamicStoreRef, keys ArrayRef, patterns ArrayRef) DictionaryRef {
	return _SCDynamicStoreCopyMultiple(store, keys, patterns)
}/* debug [functions.gen.go/function]: SCDynamicStoreCopyMultiple */

// Returns the keys that have changed since the last call to this function.
//
// Added in macOS 10.1.
// Returns the keys that have changed since the last call to this function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyNotifiedKeys(_:)
func SCDynamicStoreCopyNotifiedKeys(store DynamicStoreRef) ArrayRef {
	return _SCDynamicStoreCopyNotifiedKeys(store)
}/* debug [functions.gen.go/function]: SCDynamicStoreCopyNotifiedKeys */

// Returns the key-value pairs that represent the current internet proxy settings.
//
// Added in macOS 10.1.
// Returns the key-value pairs that represent the current internet proxy settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyProxies(_:)
func SCDynamicStoreCopyProxies(store DynamicStoreRef) DictionaryRef {
	return _SCDynamicStoreCopyProxies(store)
}/* debug [functions.gen.go/function]: SCDynamicStoreCopyProxies */

// Returns the value associated with the specified key.
//
// Added in macOS 10.1.
// Returns the value associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCopyValue(_:_:)
func SCDynamicStoreCopyValue(store DynamicStoreRef, key StringRef) PropertyListRef {
	return _SCDynamicStoreCopyValue(store, key)
}/* debug [functions.gen.go/function]: SCDynamicStoreCopyValue */

// Creates a new session used to interact with the dynamic store maintained by the System Configuration server.
//
// Added in macOS 10.1.
// Creates a new session used to interact with the dynamic store maintained by the System Configuration server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCreate(_:_:_:_:)
func SCDynamicStoreCreate(allocator AllocatorRef, name StringRef, callout DynamicStoreCallBack, context unsafe.Pointer) DynamicStoreRef {
	return _SCDynamicStoreCreate(allocator, name, callout, context)
}/* debug [functions.gen.go/function]: SCDynamicStoreCreate */

// Creates a run loop source object that can be added to the application’s run loop.
//
// Added in macOS 10.1.
// Creates a run loop source object that can be added to the application’s run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCreateRunLoopSource(_:_:_:)
func SCDynamicStoreCreateRunLoopSource(allocator AllocatorRef, store DynamicStoreRef, order Index) RunLoopSourceRef {
	return _SCDynamicStoreCreateRunLoopSource(allocator, store, order)
}/* debug [functions.gen.go/function]: SCDynamicStoreCreateRunLoopSource */

// Creates a new session used to interact with the dynamic store maintained by the System Configuration server.
//
// Added in macOS 10.4.
// Creates a new session used to interact with the dynamic store maintained by the System Configuration server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreCreateWithOptions(_:_:_:_:_:)
func SCDynamicStoreCreateWithOptions(allocator AllocatorRef, name StringRef, storeOptions DictionaryRef, callout DynamicStoreCallBack, context unsafe.Pointer) DynamicStoreRef {
	return _SCDynamicStoreCreateWithOptions(allocator, name, storeOptions, callout, context)
}/* debug [functions.gen.go/function]: SCDynamicStoreCreateWithOptions */

// Returns the type identifier of all instances.
//
// Added in macOS 10.1.
// Returns the type identifier of all instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreGetTypeID()
func SCDynamicStoreGetTypeID() TypeID {
	return _SCDynamicStoreGetTypeID()
}/* debug [functions.gen.go/function]: SCDynamicStoreGetTypeID */

// Creates a dynamic store key using the specified format.
//
// Added in macOS 10.1.
// Creates a dynamic store key using the specified format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreate
func SCDynamicStoreKeyCreate(allocator AllocatorRef, fmt StringRef) StringRef {
	return _SCDynamicStoreKeyCreate(allocator, fmt)
}/* debug [functions.gen.go/function]: SCDynamicStoreKeyCreate */

// Creates a key that can be used to receive notifications when the current computer name changes.
//
// Added in macOS 10.1.
// Creates a key that can be used to receive notifications when the current computer name changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateComputerName(_:)
func SCDynamicStoreKeyCreateComputerName(allocator AllocatorRef) StringRef {
	return _SCDynamicStoreKeyCreateComputerName(allocator)
}/* debug [functions.gen.go/function]: SCDynamicStoreKeyCreateComputerName */

// Creates a key that can be used to receive notifications when the current console user changes.
//
// Added in macOS 10.1.
// Creates a key that can be used to receive notifications when the current console user changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateConsoleUser(_:)
func SCDynamicStoreKeyCreateConsoleUser(allocator AllocatorRef) StringRef {
	return _SCDynamicStoreKeyCreateConsoleUser(allocator)
}/* debug [functions.gen.go/function]: SCDynamicStoreKeyCreateConsoleUser */

// Creates a key that can be used to receive notifications when the entity changes.
//
// Added in macOS 10.2.
// Creates a key that can be used to receive notifications when the entity changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateHostNames(_:)
func SCDynamicStoreKeyCreateHostNames(allocator AllocatorRef) StringRef {
	return _SCDynamicStoreKeyCreateHostNames(allocator)
}/* debug [functions.gen.go/function]: SCDynamicStoreKeyCreateHostNames */

// Creates a key that can be used to receive notifications when the location identifier changes.
//
// Added in macOS 10.2.
// Creates a key that can be used to receive notifications when the location identifier changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateLocation(_:)
func SCDynamicStoreKeyCreateLocation(allocator AllocatorRef) StringRef {
	return _SCDynamicStoreKeyCreateLocation(allocator)
}/* debug [functions.gen.go/function]: SCDynamicStoreKeyCreateLocation */

// Creates a dynamic store key that can be used to access a specific global (as opposed to a per-service or per-interface) network configuration entity.
//
// Added in macOS 10.1.
// Creates a dynamic store key that can be used to access a specific global (as opposed to a per-service or per-interface) network configuration entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateNetworkGlobalEntity(_:_:_:)
func SCDynamicStoreKeyCreateNetworkGlobalEntity(allocator AllocatorRef, domain StringRef, entity StringRef) StringRef {
	return _SCDynamicStoreKeyCreateNetworkGlobalEntity(allocator, domain, entity)
}/* debug [functions.gen.go/function]: SCDynamicStoreKeyCreateNetworkGlobalEntity */

// Creates a dynamic store key that can be used to access the network interface configuration information in the dynamic store.
//
// Added in macOS 10.1.
// Creates a dynamic store key that can be used to access the network interface configuration information in the dynamic store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateNetworkInterface(_:_:)
func SCDynamicStoreKeyCreateNetworkInterface(allocator AllocatorRef, domain StringRef) StringRef {
	return _SCDynamicStoreKeyCreateNetworkInterface(allocator, domain)
}/* debug [functions.gen.go/function]: SCDynamicStoreKeyCreateNetworkInterface */

// Creates a dynamic store key that can be used to access the per-interface network configuration information in the dynamic store.
//
// Added in macOS 10.1.
// Creates a dynamic store key that can be used to access the per-interface network configuration information in the dynamic store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateNetworkInterfaceEntity(_:_:_:_:)
func SCDynamicStoreKeyCreateNetworkInterfaceEntity(allocator AllocatorRef, domain StringRef, ifname StringRef, entity StringRef) StringRef {
	return _SCDynamicStoreKeyCreateNetworkInterfaceEntity(allocator, domain, ifname, entity)
}/* debug [functions.gen.go/function]: SCDynamicStoreKeyCreateNetworkInterfaceEntity */

// Creates a dynamic store key that can be used to access the per-service network configuration information.
//
// Added in macOS 10.1.
// Creates a dynamic store key that can be used to access the per-service network configuration information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateNetworkServiceEntity(_:_:_:_:)
func SCDynamicStoreKeyCreateNetworkServiceEntity(allocator AllocatorRef, domain StringRef, serviceID StringRef, entity StringRef) StringRef {
	return _SCDynamicStoreKeyCreateNetworkServiceEntity(allocator, domain, serviceID, entity)
}/* debug [functions.gen.go/function]: SCDynamicStoreKeyCreateNetworkServiceEntity */

// Creates a key that can be used to receive notifications when the current network proxy settings are changed.
//
// Added in macOS 10.1.
// Creates a key that can be used to receive notifications when the current network proxy settings are changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreKeyCreateProxies(_:)
func SCDynamicStoreKeyCreateProxies(allocator AllocatorRef) StringRef {
	return _SCDynamicStoreKeyCreateProxies(allocator)
}/* debug [functions.gen.go/function]: SCDynamicStoreKeyCreateProxies */

// Causes a notification to be delivered for the specified key in the dynamic store.
//
// Added in macOS 10.1.
// Causes a notification to be delivered for the specified key in the dynamic store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreNotifyValue(_:_:)
func SCDynamicStoreNotifyValue(store DynamicStoreRef, key StringRef) unsafe.Pointer {
	return _SCDynamicStoreNotifyValue(store, key)
}/* debug [functions.gen.go/function]: SCDynamicStoreNotifyValue */

// Removes the value of the specified key from the dynamic store.
//
// Added in macOS 10.1.
// Removes the value of the specified key from the dynamic store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreRemoveValue(_:_:)
func SCDynamicStoreRemoveValue(store DynamicStoreRef, key StringRef) unsafe.Pointer {
	return _SCDynamicStoreRemoveValue(store, key)
}/* debug [functions.gen.go/function]: SCDynamicStoreRemoveValue */

// Initiates notifications for the notification keys, using the specified dispatch queue for the callback.
//
// Added in macOS 10.6.
// Initiates notifications for the notification keys, using the specified dispatch queue for the callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreSetDispatchQueue(_:_:)
func SCDynamicStoreSetDispatchQueue(store DynamicStoreRef, queue unsafe.Pointer) unsafe.Pointer {
	return _SCDynamicStoreSetDispatchQueue(store, queue)
}/* debug [functions.gen.go/function]: SCDynamicStoreSetDispatchQueue */

// Updates multiple values in the dynamic store.
//
// Added in macOS 10.1.
// Updates multiple values in the dynamic store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreSetMultiple(_:_:_:_:)
func SCDynamicStoreSetMultiple(store DynamicStoreRef, keysToSet DictionaryRef, keysToRemove ArrayRef, keysToNotify ArrayRef) unsafe.Pointer {
	return _SCDynamicStoreSetMultiple(store, keysToSet, keysToRemove, keysToNotify)
}/* debug [functions.gen.go/function]: SCDynamicStoreSetMultiple */

// Specifies a set of keys and key patterns that should be monitored for changes.
//
// Added in macOS 10.1.
// Specifies a set of keys and key patterns that should be monitored for changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreSetNotificationKeys(_:_:_:)
func SCDynamicStoreSetNotificationKeys(store DynamicStoreRef, keys ArrayRef, patterns ArrayRef) unsafe.Pointer {
	return _SCDynamicStoreSetNotificationKeys(store, keys, patterns)
}/* debug [functions.gen.go/function]: SCDynamicStoreSetNotificationKeys */

// Adds or replaces a value in the dynamic store for the specified key.
//
// Added in macOS 10.1.
// Adds or replaces a value in the dynamic store for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCDynamicStoreSetValue(_:_:_:)
func SCDynamicStoreSetValue(store DynamicStoreRef, key StringRef, value PropertyListRef) unsafe.Pointer {
	return _SCDynamicStoreSetValue(store, key, value)
}/* debug [functions.gen.go/function]: SCDynamicStoreSetValue */

// Returns an error or status code associated with the most recent function call.
//
// Added in macOS 10.1.
// Returns an error or status code associated with the most recent function call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCError()
func SCError() int {
	return _SCError()
}/* debug [functions.gen.go/function]: SCError */

// Returns a string describing the specified status code or error code.
//
// Added in macOS 10.1.
// Returns a string describing the specified status code or error code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCErrorString(_:)
func SCErrorString(status int) unsafe.Pointer {
	return _SCErrorString(status)
}/* debug [functions.gen.go/function]: SCErrorString */

// Determines whether the specified network address is reachable using the current network configuration.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
// Determines whether the specified network address is reachable using the current network configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkCheckReachabilityByAddress
func SCNetworkCheckReachabilityByAddress(address unsafe.Pointer, addrlen unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkCheckReachabilityByAddress(address, addrlen, flags)
}/* debug [functions.gen.go/function]: SCNetworkCheckReachabilityByAddress */

// Determines whether the specified network host or node name is reachable using the current network configuration.
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.1.
// Determines whether the specified network host or node name is reachable using the current network configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkCheckReachabilityByName
func SCNetworkCheckReachabilityByName(nodename unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkCheckReachabilityByName(nodename, flags)
}/* debug [functions.gen.go/function]: SCNetworkCheckReachabilityByName */

// Returns the extended status of the connection.
//
// Added in macOS 10.3.
// Returns the extended status of the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyExtendedStatus(_:)
func SCNetworkConnectionCopyExtendedStatus(connection NetworkConnectionRef) DictionaryRef {
	return _SCNetworkConnectionCopyExtendedStatus(connection)
}/* debug [functions.gen.go/function]: SCNetworkConnectionCopyExtendedStatus */

// Returns the service ID associated with the specified network connection.
//
// Added in macOS 10.3.
// Returns the service ID associated with the specified network connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyServiceID(_:)
func SCNetworkConnectionCopyServiceID(connection NetworkConnectionRef) StringRef {
	return _SCNetworkConnectionCopyServiceID(connection)
}/* debug [functions.gen.go/function]: SCNetworkConnectionCopyServiceID */

// Returns the statistics of the specified connection.
//
// Added in macOS 10.3.
// Returns the statistics of the specified connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyStatistics(_:)
func SCNetworkConnectionCopyStatistics(connection NetworkConnectionRef) DictionaryRef {
	return _SCNetworkConnectionCopyStatistics(connection)
}/* debug [functions.gen.go/function]: SCNetworkConnectionCopyStatistics */

// Gets the user options used to start the specified connection.
//
// Added in macOS 10.3.
// Gets the user options used to start the specified connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyUserOptions(_:)
func SCNetworkConnectionCopyUserOptions(connection NetworkConnectionRef) DictionaryRef {
	return _SCNetworkConnectionCopyUserOptions(connection)
}/* debug [functions.gen.go/function]: SCNetworkConnectionCopyUserOptions */

// Provides the default service ID and a dictionary of user options for the specified connection.
//
// Added in macOS 10.3.
// Provides the default service ID and a dictionary of user options for the specified connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCopyUserPreferences(_:_:_:)
func SCNetworkConnectionCopyUserPreferences(selectionOptions DictionaryRef, serviceID unsafe.Pointer, userOptions unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionCopyUserPreferences(selectionOptions, serviceID, userOptions)
}/* debug [functions.gen.go/function]: SCNetworkConnectionCopyUserPreferences */

// Creates a new connection reference to use for getting the status or for connecting or disconnecting the associated service.
//
// Added in macOS 10.3.
// Creates a new connection reference to use for getting the status or for connecting or disconnecting the associated service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionCreateWithServiceID(_:_:_:_:)
func SCNetworkConnectionCreateWithServiceID(allocator AllocatorRef, serviceID StringRef, callout NetworkConnectionCallBack, context unsafe.Pointer) NetworkConnectionRef {
	return _SCNetworkConnectionCreateWithServiceID(allocator, serviceID, callout, context)
}/* debug [functions.gen.go/function]: SCNetworkConnectionCreateWithServiceID */

// Returns the status of the specified network connection.
//
// Added in macOS 10.3.
// Returns the status of the specified network connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionGetStatus(_:)
func SCNetworkConnectionGetStatus(connection NetworkConnectionRef) NetworkConnectionStatus {
	return _SCNetworkConnectionGetStatus(connection)
}/* debug [functions.gen.go/function]: SCNetworkConnectionGetStatus */

// Returns the type identifier of all instances.
//
// Added in macOS 10.3.
// Returns the type identifier of all instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionGetTypeID()
func SCNetworkConnectionGetTypeID() TypeID {
	return _SCNetworkConnectionGetTypeID()
}/* debug [functions.gen.go/function]: SCNetworkConnectionGetTypeID */

// Schedules the specified connection with the specified run loop.
//
// Added in macOS 10.3.
// Schedules the specified connection with the specified run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionScheduleWithRunLoop(_:_:_:)
func SCNetworkConnectionScheduleWithRunLoop(connection NetworkConnectionRef, runLoop RunLoopRef, runLoopMode StringRef) unsafe.Pointer {
	return _SCNetworkConnectionScheduleWithRunLoop(connection, runLoop, runLoopMode)
}/* debug [functions.gen.go/function]: SCNetworkConnectionScheduleWithRunLoop */

// Specifies a dispatch queue to use for the connection’s callback function and enables notifications.
//
// Added in macOS 10.6.
// Specifies a dispatch queue to use for the connection’s callback function and enables notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionSetDispatchQueue(_:_:)
func SCNetworkConnectionSetDispatchQueue(connection NetworkConnectionRef, queue unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionSetDispatchQueue(connection, queue)
}/* debug [functions.gen.go/function]: SCNetworkConnectionSetDispatchQueue */

// Starts the connection process for the specified network connection.
//
// Added in macOS 10.3.
// Starts the connection process for the specified network connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionStart(_:_:_:)
func SCNetworkConnectionStart(connection NetworkConnectionRef, userOptions DictionaryRef, linger unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionStart(connection, userOptions, linger)
}/* debug [functions.gen.go/function]: SCNetworkConnectionStart */

// Stops the connection process for the specified network connection.
//
// Added in macOS 10.3.
// Stops the connection process for the specified network connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionStop(_:_:)
func SCNetworkConnectionStop(connection NetworkConnectionRef, forceDisconnect unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkConnectionStop(connection, forceDisconnect)
}/* debug [functions.gen.go/function]: SCNetworkConnectionStop */

// Unschedules the specified connection from the specified run loop.
//
// Added in macOS 10.3.
// Unschedules the specified connection from the specified run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkConnectionUnscheduleFromRunLoop(_:_:_:)
func SCNetworkConnectionUnscheduleFromRunLoop(connection NetworkConnectionRef, runLoop RunLoopRef, runLoopMode StringRef) unsafe.Pointer {
	return _SCNetworkConnectionUnscheduleFromRunLoop(connection, runLoop, runLoopMode)
}/* debug [functions.gen.go/function]: SCNetworkConnectionUnscheduleFromRunLoop */

// Returns all network-capable interfaces on the system.
//
// Added in macOS 10.4.
// Returns all network-capable interfaces on the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyAll()
func SCNetworkInterfaceCopyAll() ArrayRef {
	return _SCNetworkInterfaceCopyAll()
}/* debug [functions.gen.go/function]: SCNetworkInterfaceCopyAll */

// Returns information media options for the specified network interface.
//
// Added in macOS 10.5.
// Returns information media options for the specified network interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyMediaOptions(_:_:_:_:_:)
func SCNetworkInterfaceCopyMediaOptions(interface_ NetworkInterfaceRef, current unsafe.Pointer, active unsafe.Pointer, available unsafe.Pointer, filter unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkInterfaceCopyMediaOptions(interface_, current, active, available, filter)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceCopyMediaOptions */

// Returns a list of available media options for the specified interface configuration options and subtype.
//
// Added in macOS 10.5.
// Returns a list of available media options for the specified interface configuration options and subtype.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyMediaSubTypeOptions(_:_:)
func SCNetworkInterfaceCopyMediaSubTypeOptions(available ArrayRef, subType StringRef) ArrayRef {
	return _SCNetworkInterfaceCopyMediaSubTypeOptions(available, subType)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceCopyMediaSubTypeOptions */

// Returns a list of available media subtypes for the specified interface configuration options.
//
// Added in macOS 10.5.
// Returns a list of available media subtypes for the specified interface configuration options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyMediaSubTypes(_:)
func SCNetworkInterfaceCopyMediaSubTypes(available ArrayRef) ArrayRef {
	return _SCNetworkInterfaceCopyMediaSubTypes(available)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceCopyMediaSubTypes */

// Returns the current MTU setting and the range of allowable values for the specified network interface.
//
// Added in macOS 10.5.
// Returns the current MTU setting and the range of allowable values for the specified network interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCopyMTU(_:_:_:_:)
func SCNetworkInterfaceCopyMTU(interface_ NetworkInterfaceRef, mtu_cur []int, mtu_min []int, mtu_max []int) unsafe.Pointer {
	return _SCNetworkInterfaceCopyMTU(interface_, mtu_cur, mtu_min, mtu_max)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceCopyMTU */

// Creates a new network interface layered on top of the specified interface.
//
// Added in macOS 10.4.
// Creates a new network interface layered on top of the specified interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceCreateWithInterface(_:_:)
func SCNetworkInterfaceCreateWithInterface(interface_ NetworkInterfaceRef, interfaceType StringRef) NetworkInterfaceRef {
	return _SCNetworkInterfaceCreateWithInterface(interface_, interfaceType)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceCreateWithInterface */

// Sends a notification to interested network configuration agents to immediately retry their configuration.
//
// Added in macOS 10.5.
// Sends a notification to interested network configuration agents to immediately retry their configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceForceConfigurationRefresh(_:)
func SCNetworkInterfaceForceConfigurationRefresh(interface_ NetworkInterfaceRef) unsafe.Pointer {
	return _SCNetworkInterfaceForceConfigurationRefresh(interface_)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceForceConfigurationRefresh */

// Returns the BSD interface or device name for the specified interface.
//
// Added in macOS 10.4.
// Returns the BSD interface or device name for the specified interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetBSDName(_:)
func SCNetworkInterfaceGetBSDName(interface_ NetworkInterfaceRef) StringRef {
	return _SCNetworkInterfaceGetBSDName(interface_)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceGetBSDName */

// Returns the configuration settings associated with the specified interface.
//
// Added in macOS 10.4.
// Returns the configuration settings associated with the specified interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetConfiguration(_:)
func SCNetworkInterfaceGetConfiguration(interface_ NetworkInterfaceRef) DictionaryRef {
	return _SCNetworkInterfaceGetConfiguration(interface_)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceGetConfiguration */

// Returns the extended configuration settings associated with the specified interface.
//
// Added in macOS 10.5.
// Returns the extended configuration settings associated with the specified interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetExtendedConfiguration(_:_:)
func SCNetworkInterfaceGetExtendedConfiguration(interface_ NetworkInterfaceRef, extendedType StringRef) DictionaryRef {
	return _SCNetworkInterfaceGetExtendedConfiguration(interface_, extendedType)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceGetExtendedConfiguration */

// Returns a displayable link layer address for the specified interface.
//
// Added in macOS 10.4.
// Returns a displayable link layer address for the specified interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetHardwareAddressString(_:)
func SCNetworkInterfaceGetHardwareAddressString(interface_ NetworkInterfaceRef) StringRef {
	return _SCNetworkInterfaceGetHardwareAddressString(interface_)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceGetHardwareAddressString */

// Returns the underlying interface, for layered network interfaces.
//
// Added in macOS 10.4.
// Returns the underlying interface, for layered network interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetInterface(_:)
func SCNetworkInterfaceGetInterface(interface_ NetworkInterfaceRef) NetworkInterfaceRef {
	return _SCNetworkInterfaceGetInterface(interface_)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceGetInterface */

// Returns the network interface type of the specified interface.
//
// Added in macOS 10.4.
// Returns the network interface type of the specified interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetInterfaceType(_:)
func SCNetworkInterfaceGetInterfaceType(interface_ NetworkInterfaceRef) StringRef {
	return _SCNetworkInterfaceGetInterfaceType(interface_)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceGetInterfaceType */

// Returns the localized display name, such as “Ethernet” or “FireWire”, for the specified interface.
//
// Added in macOS 10.4.
// Returns the localized display name, such as “Ethernet” or “FireWire”, for the specified interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetLocalizedDisplayName(_:)
func SCNetworkInterfaceGetLocalizedDisplayName(interface_ NetworkInterfaceRef) StringRef {
	return _SCNetworkInterfaceGetLocalizedDisplayName(interface_)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceGetLocalizedDisplayName */

// Identifies all of the network interface types, such as PPP, that can be layered on top of the specified interface.
//
// Added in macOS 10.4.
// Identifies all of the network interface types, such as PPP, that can be layered on top of the specified interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetSupportedInterfaceTypes(_:)
func SCNetworkInterfaceGetSupportedInterfaceTypes(interface_ NetworkInterfaceRef) ArrayRef {
	return _SCNetworkInterfaceGetSupportedInterfaceTypes(interface_)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceGetSupportedInterfaceTypes */

// Identifies all of the network protocol types, such as IPv4 and IPv6, that can be layered on top of the specified interface.
//
// Added in macOS 10.4.
// Identifies all of the network protocol types, such as IPv4 and IPv6, that can be layered on top of the specified interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetSupportedProtocolTypes(_:)
func SCNetworkInterfaceGetSupportedProtocolTypes(interface_ NetworkInterfaceRef) ArrayRef {
	return _SCNetworkInterfaceGetSupportedProtocolTypes(interface_)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceGetSupportedProtocolTypes */

// Returns the type identifier of all instances.
//
// Added in macOS 10.4.
// Returns the type identifier of all instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceGetTypeID()
func SCNetworkInterfaceGetTypeID() TypeID {
	return _SCNetworkInterfaceGetTypeID()
}/* debug [functions.gen.go/function]: SCNetworkInterfaceGetTypeID */

// Sends a notification to interested configuration agents to have them immediately retry their configuration over a particular network interface.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.1.
// Sends a notification to interested configuration agents to have them immediately retry their configuration over a particular network interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceRefreshConfiguration
func SCNetworkInterfaceRefreshConfiguration(ifName StringRef) unsafe.Pointer {
	return _SCNetworkInterfaceRefreshConfiguration(ifName)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceRefreshConfiguration */

// Stores the configuration settings for the specified interface.
//
// Added in macOS 10.4.
// Stores the configuration settings for the specified interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceSetConfiguration(_:_:)
func SCNetworkInterfaceSetConfiguration(interface_ NetworkInterfaceRef, config DictionaryRef) unsafe.Pointer {
	return _SCNetworkInterfaceSetConfiguration(interface_, config)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceSetConfiguration */

// Stores the extended configuration settings for the specified interface.
//
// Added in macOS 10.5.
// Stores the extended configuration settings for the specified interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceSetExtendedConfiguration(_:_:_:)
func SCNetworkInterfaceSetExtendedConfiguration(interface_ NetworkInterfaceRef, extendedType StringRef, config DictionaryRef) unsafe.Pointer {
	return _SCNetworkInterfaceSetExtendedConfiguration(interface_, extendedType, config)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceSetExtendedConfiguration */

// Sets the requested media subtype and options for the specified network interface.
//
// Added in macOS 10.5.
// Sets the requested media subtype and options for the specified network interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceSetMediaOptions(_:_:_:)
func SCNetworkInterfaceSetMediaOptions(interface_ NetworkInterfaceRef, subtype StringRef, options ArrayRef) unsafe.Pointer {
	return _SCNetworkInterfaceSetMediaOptions(interface_, subtype, options)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceSetMediaOptions */

// Sets the requested MTU setting for the specified network interface.
//
// Added in macOS 10.5.
// Sets the requested MTU setting for the specified network interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkInterfaceSetMTU(_:_:)
func SCNetworkInterfaceSetMTU(interface_ NetworkInterfaceRef, mtu int) unsafe.Pointer {
	return _SCNetworkInterfaceSetMTU(interface_, mtu)
}/* debug [functions.gen.go/function]: SCNetworkInterfaceSetMTU */

// Returns the configuration settings associated with the specified protocol.
//
// Added in macOS 10.4.
// Returns the configuration settings associated with the specified protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolGetConfiguration(_:)
func SCNetworkProtocolGetConfiguration(protocol_ NetworkProtocolRef) DictionaryRef {
	return _SCNetworkProtocolGetConfiguration(protocol_)
}/* debug [functions.gen.go/function]: SCNetworkProtocolGetConfiguration */

// Returns a Boolean value indicating whether the specified protocol is enabled.
//
// Added in macOS 10.4.
// Returns a Boolean value indicating whether the specified protocol is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolGetEnabled(_:)
func SCNetworkProtocolGetEnabled(protocol_ NetworkProtocolRef) unsafe.Pointer {
	return _SCNetworkProtocolGetEnabled(protocol_)
}/* debug [functions.gen.go/function]: SCNetworkProtocolGetEnabled */

// Returns the type of the specified network protocol.
//
// Added in macOS 10.4.
// Returns the type of the specified network protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolGetProtocolType(_:)
func SCNetworkProtocolGetProtocolType(protocol_ NetworkProtocolRef) StringRef {
	return _SCNetworkProtocolGetProtocolType(protocol_)
}/* debug [functions.gen.go/function]: SCNetworkProtocolGetProtocolType */

// Returns the type identifier of all instances.
//
// Added in macOS 10.4.
// Returns the type identifier of all instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolGetTypeID()
func SCNetworkProtocolGetTypeID() TypeID {
	return _SCNetworkProtocolGetTypeID()
}/* debug [functions.gen.go/function]: SCNetworkProtocolGetTypeID */

// Stores the configuration settings for the specified network protocol.
//
// Added in macOS 10.4.
// Stores the configuration settings for the specified network protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolSetConfiguration(_:_:)
func SCNetworkProtocolSetConfiguration(protocol_ NetworkProtocolRef, config DictionaryRef) unsafe.Pointer {
	return _SCNetworkProtocolSetConfiguration(protocol_, config)
}/* debug [functions.gen.go/function]: SCNetworkProtocolSetConfiguration */

// Enables or disables the specified protocol.
//
// Added in macOS 10.4.
// Enables or disables the specified protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkProtocolSetEnabled(_:_:)
func SCNetworkProtocolSetEnabled(protocol_ NetworkProtocolRef, enabled unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkProtocolSetEnabled(protocol_, enabled)
}/* debug [functions.gen.go/function]: SCNetworkProtocolSetEnabled */

// Creates a reachability reference to the specified network address.
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
// Creates a reachability reference to the specified network address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityCreateWithAddress(_:_:)
func SCNetworkReachabilityCreateWithAddress(allocator AllocatorRef, address unsafe.Pointer) NetworkReachabilityRef {
	return _SCNetworkReachabilityCreateWithAddress(allocator, address)
}/* debug [functions.gen.go/function]: SCNetworkReachabilityCreateWithAddress */

// Creates a reachability reference to the specified network address.
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
// Creates a reachability reference to the specified network address.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityCreateWithAddressPair(_:_:_:)
func SCNetworkReachabilityCreateWithAddressPair(allocator AllocatorRef, localAddress unsafe.Pointer, remoteAddress unsafe.Pointer) NetworkReachabilityRef {
	return _SCNetworkReachabilityCreateWithAddressPair(allocator, localAddress, remoteAddress)
}/* debug [functions.gen.go/function]: SCNetworkReachabilityCreateWithAddressPair */

// Creates a reachability reference to the specified network host or node name.
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
// Creates a reachability reference to the specified network host or node name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityCreateWithName(_:_:)
func SCNetworkReachabilityCreateWithName(allocator AllocatorRef, nodename unsafe.Pointer) NetworkReachabilityRef {
	return _SCNetworkReachabilityCreateWithName(allocator, nodename)
}/* debug [functions.gen.go/function]: SCNetworkReachabilityCreateWithName */

// Determines if the specified network target is reachable using the current network configuration.
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
// Determines if the specified network target is reachable using the current network configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityGetFlags(_:_:)
func SCNetworkReachabilityGetFlags(target NetworkReachabilityRef, flags unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkReachabilityGetFlags(target, flags)
}/* debug [functions.gen.go/function]: SCNetworkReachabilityGetFlags */

// Returns the type identifier of all instances.
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
// Returns the type identifier of all instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityGetTypeID()
func SCNetworkReachabilityGetTypeID() TypeID {
	return _SCNetworkReachabilityGetTypeID()
}/* debug [functions.gen.go/function]: SCNetworkReachabilityGetTypeID */

// Schedules the specified network target with the specified run loop and mode.
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
// Schedules the specified network target with the specified run loop and mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityScheduleWithRunLoop(_:_:_:)
func SCNetworkReachabilityScheduleWithRunLoop(target NetworkReachabilityRef, runLoop RunLoopRef, runLoopMode StringRef) unsafe.Pointer {
	return _SCNetworkReachabilityScheduleWithRunLoop(target, runLoop, runLoopMode)
}/* debug [functions.gen.go/function]: SCNetworkReachabilityScheduleWithRunLoop */

// Assigns a client to the specified target, which receives callbacks when the reachability of the target changes.
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
// Assigns a client to the specified target, which receives callbacks when the reachability of the target changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilitySetCallback(_:_:_:)
func SCNetworkReachabilitySetCallback(target NetworkReachabilityRef, callout NetworkReachabilityCallBack, context unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkReachabilitySetCallback(target, callout, context)
}/* debug [functions.gen.go/function]: SCNetworkReachabilitySetCallback */

// Schedules callbacks for the specified target on the specified dispatch queue.
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.6.
// Schedules callbacks for the specified target on the specified dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilitySetDispatchQueue(_:_:)
func SCNetworkReachabilitySetDispatchQueue(target NetworkReachabilityRef, queue unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkReachabilitySetDispatchQueue(target, queue)
}/* debug [functions.gen.go/function]: SCNetworkReachabilitySetDispatchQueue */

// Unschedules the specified target from the specified run loop and mode.
//
// Deprecated: This function was deprecated in macOS 14.4.
//
// Added in macOS 10.3.
// Unschedules the specified target from the specified run loop and mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachabilityUnscheduleFromRunLoop(_:_:_:)
func SCNetworkReachabilityUnscheduleFromRunLoop(target NetworkReachabilityRef, runLoop RunLoopRef, runLoopMode StringRef) unsafe.Pointer {
	return _SCNetworkReachabilityUnscheduleFromRunLoop(target, runLoop, runLoopMode)
}/* debug [functions.gen.go/function]: SCNetworkReachabilityUnscheduleFromRunLoop */

// Adds the network protocol of the specified type to the specified service.
//
// Added in macOS 10.4.
// Adds the network protocol of the specified type to the specified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceAddProtocolType(_:_:)
func SCNetworkServiceAddProtocolType(service NetworkServiceRef, protocolType StringRef) unsafe.Pointer {
	return _SCNetworkServiceAddProtocolType(service, protocolType)
}/* debug [functions.gen.go/function]: SCNetworkServiceAddProtocolType */

// Returns the network service with the specified identifier.
//
// Added in macOS 10.4.
// Returns the network service with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCopy(_:_:)
func SCNetworkServiceCopy(prefs PreferencesRef, serviceID StringRef) NetworkServiceRef {
	return _SCNetworkServiceCopy(prefs, serviceID)
}/* debug [functions.gen.go/function]: SCNetworkServiceCopy */

// Returns all available network services for the specified preferences.
//
// Added in macOS 10.4.
// Returns all available network services for the specified preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCopyAll(_:)
func SCNetworkServiceCopyAll(prefs PreferencesRef) ArrayRef {
	return _SCNetworkServiceCopyAll(prefs)
}/* debug [functions.gen.go/function]: SCNetworkServiceCopyAll */

// Returns the network protocol of the specified type for the specified service.
//
// Added in macOS 10.4.
// Returns the network protocol of the specified type for the specified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCopyProtocol(_:_:)
func SCNetworkServiceCopyProtocol(service NetworkServiceRef, protocolType StringRef) NetworkProtocolRef {
	return _SCNetworkServiceCopyProtocol(service, protocolType)
}/* debug [functions.gen.go/function]: SCNetworkServiceCopyProtocol */

// Returns all network protocols associated with the specified service.
//
// Added in macOS 10.4.
// Returns all network protocols associated with the specified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCopyProtocols(_:)
func SCNetworkServiceCopyProtocols(service NetworkServiceRef) ArrayRef {
	return _SCNetworkServiceCopyProtocols(service)
}/* debug [functions.gen.go/function]: SCNetworkServiceCopyProtocols */

// Creates a new network service for the specified interface in the configuration.
//
// Added in macOS 10.4.
// Creates a new network service for the specified interface in the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceCreate(_:_:)
func SCNetworkServiceCreate(prefs PreferencesRef, interface_ NetworkInterfaceRef) NetworkServiceRef {
	return _SCNetworkServiceCreate(prefs, interface_)
}/* debug [functions.gen.go/function]: SCNetworkServiceCreate */

// Establishes the default configuration for the specified network service.
//
// Added in macOS 10.5.
// Establishes the default configuration for the specified network service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceEstablishDefaultConfiguration(_:)
func SCNetworkServiceEstablishDefaultConfiguration(service NetworkServiceRef) unsafe.Pointer {
	return _SCNetworkServiceEstablishDefaultConfiguration(service)
}/* debug [functions.gen.go/function]: SCNetworkServiceEstablishDefaultConfiguration */

// Returns a Boolean value indicating whether the specified service is enabled.
//
// Added in macOS 10.4.
// Returns a Boolean value indicating whether the specified service is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetEnabled(_:)
func SCNetworkServiceGetEnabled(service NetworkServiceRef) unsafe.Pointer {
	return _SCNetworkServiceGetEnabled(service)
}/* debug [functions.gen.go/function]: SCNetworkServiceGetEnabled */

// Returns the network interface associated with the specified service.
//
// Added in macOS 10.4.
// Returns the network interface associated with the specified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetInterface(_:)
func SCNetworkServiceGetInterface(service NetworkServiceRef) NetworkInterfaceRef {
	return _SCNetworkServiceGetInterface(service)
}/* debug [functions.gen.go/function]: SCNetworkServiceGetInterface */

// Returns the user-specified name associated with the specified service.
//
// Added in macOS 10.4.
// Returns the user-specified name associated with the specified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetName(_:)
func SCNetworkServiceGetName(service NetworkServiceRef) StringRef {
	return _SCNetworkServiceGetName(service)
}/* debug [functions.gen.go/function]: SCNetworkServiceGetName */

// Returns the identifier for the specified service.
//
// Added in macOS 10.4.
// Returns the identifier for the specified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetServiceID(_:)
func SCNetworkServiceGetServiceID(service NetworkServiceRef) StringRef {
	return _SCNetworkServiceGetServiceID(service)
}/* debug [functions.gen.go/function]: SCNetworkServiceGetServiceID */

// Returns the type identifier of all instances.
//
// Added in macOS 10.4.
// Returns the type identifier of all instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceGetTypeID()
func SCNetworkServiceGetTypeID() TypeID {
	return _SCNetworkServiceGetTypeID()
}/* debug [functions.gen.go/function]: SCNetworkServiceGetTypeID */

// Removes the specified network service from the configuration.
//
// Added in macOS 10.4.
// Removes the specified network service from the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceRemove(_:)
func SCNetworkServiceRemove(service NetworkServiceRef) unsafe.Pointer {
	return _SCNetworkServiceRemove(service)
}/* debug [functions.gen.go/function]: SCNetworkServiceRemove */

// Removes the network protocol of the specified type from the specified service.
//
// Added in macOS 10.4.
// Removes the network protocol of the specified type from the specified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceRemoveProtocolType(_:_:)
func SCNetworkServiceRemoveProtocolType(service NetworkServiceRef, protocolType StringRef) unsafe.Pointer {
	return _SCNetworkServiceRemoveProtocolType(service, protocolType)
}/* debug [functions.gen.go/function]: SCNetworkServiceRemoveProtocolType */

// Enables or disables the specified service.
//
// Added in macOS 10.4.
// Enables or disables the specified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceSetEnabled(_:_:)
func SCNetworkServiceSetEnabled(service NetworkServiceRef, enabled unsafe.Pointer) unsafe.Pointer {
	return _SCNetworkServiceSetEnabled(service, enabled)
}/* debug [functions.gen.go/function]: SCNetworkServiceSetEnabled */

// Stores the user-specified name for the specified service.
//
// Added in macOS 10.4.
// Stores the user-specified name for the specified service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkServiceSetName(_:_:)
func SCNetworkServiceSetName(service NetworkServiceRef, name StringRef) unsafe.Pointer {
	return _SCNetworkServiceSetName(service, name)
}/* debug [functions.gen.go/function]: SCNetworkServiceSetName */

// Adds the specified network service to the specified set.
//
// Added in macOS 10.4.
// Adds the specified network service to the specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetAddService(_:_:)
func SCNetworkSetAddService(set NetworkSetRef, service NetworkServiceRef) unsafe.Pointer {
	return _SCNetworkSetAddService(set, service)
}/* debug [functions.gen.go/function]: SCNetworkSetAddService */

// Returns a Boolean value indicating whether the specified interface is represented by at least one network service in the specified set.
//
// Added in macOS 10.5.
// Returns a Boolean value indicating whether the specified interface is represented by at least one network service in the specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetContainsInterface(_:_:)
func SCNetworkSetContainsInterface(set NetworkSetRef, interface_ NetworkInterfaceRef) unsafe.Pointer {
	return _SCNetworkSetContainsInterface(set, interface_)
}/* debug [functions.gen.go/function]: SCNetworkSetContainsInterface */

// Returns the set with the specified identifier.
//
// Added in macOS 10.4.
// Returns the set with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCopy(_:_:)
func SCNetworkSetCopy(prefs PreferencesRef, setID StringRef) NetworkSetRef {
	return _SCNetworkSetCopy(prefs, setID)
}/* debug [functions.gen.go/function]: SCNetworkSetCopy */

// Returns all available sets for the specified preferences session.
//
// Added in macOS 10.4.
// Returns all available sets for the specified preferences session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCopyAll(_:)
func SCNetworkSetCopyAll(prefs PreferencesRef) ArrayRef {
	return _SCNetworkSetCopyAll(prefs)
}/* debug [functions.gen.go/function]: SCNetworkSetCopyAll */

// Returns the current set.
//
// Added in macOS 10.4.
// Returns the current set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCopyCurrent(_:)
func SCNetworkSetCopyCurrent(prefs PreferencesRef) NetworkSetRef {
	return _SCNetworkSetCopyCurrent(prefs)
}/* debug [functions.gen.go/function]: SCNetworkSetCopyCurrent */

// Returns all network services associated with the specified set.
//
// Added in macOS 10.4.
// Returns all network services associated with the specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCopyServices(_:)
func SCNetworkSetCopyServices(set NetworkSetRef) ArrayRef {
	return _SCNetworkSetCopyServices(set)
}/* debug [functions.gen.go/function]: SCNetworkSetCopyServices */

// Creates a new set in the configuration.
//
// Added in macOS 10.4.
// Creates a new set in the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetCreate(_:)
func SCNetworkSetCreate(prefs PreferencesRef) NetworkSetRef {
	return _SCNetworkSetCreate(prefs)
}/* debug [functions.gen.go/function]: SCNetworkSetCreate */

// Returns the user-specified name associated with the specified set.
//
// Added in macOS 10.4.
// Returns the user-specified name associated with the specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetGetName(_:)
func SCNetworkSetGetName(set NetworkSetRef) StringRef {
	return _SCNetworkSetGetName(set)
}/* debug [functions.gen.go/function]: SCNetworkSetGetName */

// Returns the user-specified ordering of network services within the specified set.
//
// Added in macOS 10.4.
// Returns the user-specified ordering of network services within the specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetGetServiceOrder(_:)
func SCNetworkSetGetServiceOrder(set NetworkSetRef) ArrayRef {
	return _SCNetworkSetGetServiceOrder(set)
}/* debug [functions.gen.go/function]: SCNetworkSetGetServiceOrder */

// Returns the identifier for the specified set.
//
// Added in macOS 10.4.
// Returns the identifier for the specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetGetSetID(_:)
func SCNetworkSetGetSetID(set NetworkSetRef) StringRef {
	return _SCNetworkSetGetSetID(set)
}/* debug [functions.gen.go/function]: SCNetworkSetGetSetID */

// Returns the type identifier of all instances.
//
// Added in macOS 10.4.
// Returns the type identifier of all instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetGetTypeID()
func SCNetworkSetGetTypeID() TypeID {
	return _SCNetworkSetGetTypeID()
}/* debug [functions.gen.go/function]: SCNetworkSetGetTypeID */

// Removes the specified set from the configuration.
//
// Added in macOS 10.4.
// Removes the specified set from the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetRemove(_:)
func SCNetworkSetRemove(set NetworkSetRef) unsafe.Pointer {
	return _SCNetworkSetRemove(set)
}/* debug [functions.gen.go/function]: SCNetworkSetRemove */

// Removes the specified network service from the specified set.
//
// Added in macOS 10.4.
// Removes the specified network service from the specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetRemoveService(_:_:)
func SCNetworkSetRemoveService(set NetworkSetRef, service NetworkServiceRef) unsafe.Pointer {
	return _SCNetworkSetRemoveService(set, service)
}/* debug [functions.gen.go/function]: SCNetworkSetRemoveService */

// Specifies the set that should be the current set.
//
// Added in macOS 10.4.
// Specifies the set that should be the current set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetSetCurrent(_:)
func SCNetworkSetSetCurrent(set NetworkSetRef) unsafe.Pointer {
	return _SCNetworkSetSetCurrent(set)
}/* debug [functions.gen.go/function]: SCNetworkSetSetCurrent */

// Stores the user-specified name for the specified set.
//
// Added in macOS 10.4.
// Stores the user-specified name for the specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetSetName(_:_:)
func SCNetworkSetSetName(set NetworkSetRef, name StringRef) unsafe.Pointer {
	return _SCNetworkSetSetName(set, name)
}/* debug [functions.gen.go/function]: SCNetworkSetSetName */

// Stores the user-specified ordering of network services for the specified set.
//
// Added in macOS 10.4.
// Stores the user-specified ordering of network services for the specified set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkSetSetServiceOrder(_:_:)
func SCNetworkSetSetServiceOrder(set NetworkSetRef, newOrder ArrayRef) unsafe.Pointer {
	return _SCNetworkSetSetServiceOrder(set, newOrder)
}/* debug [functions.gen.go/function]: SCNetworkSetSetServiceOrder */

// Associates the specified value with the specified preference key.
//
// Added in macOS 10.1.
// Associates the specified value with the specified preference key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesAddValue(_:_:_:)
func SCPreferencesAddValue(prefs PreferencesRef, key StringRef, value PropertyListRef) unsafe.Pointer {
	return _SCPreferencesAddValue(prefs, key, value)
}/* debug [functions.gen.go/function]: SCPreferencesAddValue */

// Requests that the currently stored configuration preferences be applied to the active configuration.
//
// Added in macOS 10.1.
// Requests that the currently stored configuration preferences be applied to the active configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesApplyChanges(_:)
func SCPreferencesApplyChanges(prefs PreferencesRef) unsafe.Pointer {
	return _SCPreferencesApplyChanges(prefs)
}/* debug [functions.gen.go/function]: SCPreferencesApplyChanges */

// Commits changes made to the configuration preferences to persistent storage.
//
// Added in macOS 10.1.
// Commits changes made to the configuration preferences to persistent storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCommitChanges(_:)
func SCPreferencesCommitChanges(prefs PreferencesRef) unsafe.Pointer {
	return _SCPreferencesCommitChanges(prefs)
}/* debug [functions.gen.go/function]: SCPreferencesCommitChanges */

// Returns the currently defined preference keys.
//
// Added in macOS 10.1.
// Returns the currently defined preference keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCopyKeyList(_:)
func SCPreferencesCopyKeyList(prefs PreferencesRef) ArrayRef {
	return _SCPreferencesCopyKeyList(prefs)
}/* debug [functions.gen.go/function]: SCPreferencesCopyKeyList */

// Initiates access to the per-system set of configuration preferences.
//
// Added in macOS 10.1.
// Initiates access to the per-system set of configuration preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCreate(_:_:_:)
func SCPreferencesCreate(allocator AllocatorRef, name StringRef, prefsID StringRef) PreferencesRef {
	return _SCPreferencesCreate(allocator, name, prefsID)
}/* debug [functions.gen.go/function]: SCPreferencesCreate */

// Initiates access to the per-system set of configuration preferences with the specified authorization.
//
// Added in macOS 10.5.
// Initiates access to the per-system set of configuration preferences with the specified authorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesCreateWithAuthorization(_:_:_:_:)
func SCPreferencesCreateWithAuthorization(allocator AllocatorRef, name StringRef, prefsID StringRef, authorization AuthorizationRef) PreferencesRef {
	return _SCPreferencesCreateWithAuthorization(allocator, name, prefsID, authorization)
}/* debug [functions.gen.go/function]: SCPreferencesCreateWithAuthorization */

// Returns a value that can be used to determine if the saved configuration preferences have changed.
//
// Added in macOS 10.1.
// Returns a value that can be used to determine if the saved configuration preferences have changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesGetSignature(_:)
func SCPreferencesGetSignature(prefs PreferencesRef) DataRef {
	return _SCPreferencesGetSignature(prefs)
}/* debug [functions.gen.go/function]: SCPreferencesGetSignature */

// Returns the type identifier of all instances.
//
// Added in macOS 10.1.
// Returns the type identifier of all instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesGetTypeID()
func SCPreferencesGetTypeID() TypeID {
	return _SCPreferencesGetTypeID()
}/* debug [functions.gen.go/function]: SCPreferencesGetTypeID */

// Retrieves the value associated with the specified preference key.
//
// Added in macOS 10.1.
// Retrieves the value associated with the specified preference key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesGetValue(_:_:)
func SCPreferencesGetValue(prefs PreferencesRef, key StringRef) PropertyListRef {
	return _SCPreferencesGetValue(prefs, key)
}/* debug [functions.gen.go/function]: SCPreferencesGetValue */

// Locks access to the configuration preferences.
//
// Added in macOS 10.1.
// Locks access to the configuration preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesLock(_:_:)
func SCPreferencesLock(prefs PreferencesRef, wait unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesLock(prefs, wait)
}/* debug [functions.gen.go/function]: SCPreferencesLock */

// Creates a new path component rooted at the specified path in the dictionary hierarchy.
//
// Added in macOS 10.1.
// Creates a new path component rooted at the specified path in the dictionary hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathCreateUniqueChild(_:_:)
func SCPreferencesPathCreateUniqueChild(prefs PreferencesRef, prefix StringRef) StringRef {
	return _SCPreferencesPathCreateUniqueChild(prefs, prefix)
}/* debug [functions.gen.go/function]: SCPreferencesPathCreateUniqueChild */

// Returns the link associated with the specified path.
//
// Added in macOS 10.1.
// Returns the link associated with the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathGetLink(_:_:)
func SCPreferencesPathGetLink(prefs PreferencesRef, path StringRef) StringRef {
	return _SCPreferencesPathGetLink(prefs, path)
}/* debug [functions.gen.go/function]: SCPreferencesPathGetLink */

// Returns the dictionary associated with the specified path.
//
// Added in macOS 10.1.
// Returns the dictionary associated with the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathGetValue(_:_:)
func SCPreferencesPathGetValue(prefs PreferencesRef, path StringRef) DictionaryRef {
	return _SCPreferencesPathGetValue(prefs, path)
}/* debug [functions.gen.go/function]: SCPreferencesPathGetValue */

// Removes the data associated with the specified path.
//
// Added in macOS 10.1.
// Removes the data associated with the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathRemoveValue(_:_:)
func SCPreferencesPathRemoveValue(prefs PreferencesRef, path StringRef) unsafe.Pointer {
	return _SCPreferencesPathRemoveValue(prefs, path)
}/* debug [functions.gen.go/function]: SCPreferencesPathRemoveValue */

// Associates a link to a second dictionary at the specified path.
//
// Added in macOS 10.1.
// Associates a link to a second dictionary at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathSetLink(_:_:_:)
func SCPreferencesPathSetLink(prefs PreferencesRef, path StringRef, link StringRef) unsafe.Pointer {
	return _SCPreferencesPathSetLink(prefs, path, link)
}/* debug [functions.gen.go/function]: SCPreferencesPathSetLink */

// Associates the specified dictionary with the specified path.
//
// Added in macOS 10.1.
// Associates the specified dictionary with the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesPathSetValue(_:_:_:)
func SCPreferencesPathSetValue(prefs PreferencesRef, path StringRef, value DictionaryRef) unsafe.Pointer {
	return _SCPreferencesPathSetValue(prefs, path, value)
}/* debug [functions.gen.go/function]: SCPreferencesPathSetValue */

// Removes the data associated with the specified preference key.
//
// Added in macOS 10.1.
// Removes the data associated with the specified preference key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesRemoveValue(_:_:)
func SCPreferencesRemoveValue(prefs PreferencesRef, key StringRef) unsafe.Pointer {
	return _SCPreferencesRemoveValue(prefs, key)
}/* debug [functions.gen.go/function]: SCPreferencesRemoveValue */

// Schedules commit and apply notifications for the specified preferences session using the specified run loop and mode.
//
// Added in macOS 10.4.
// Schedules commit and apply notifications for the specified preferences session using the specified run loop and mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesScheduleWithRunLoop(_:_:_:)
func SCPreferencesScheduleWithRunLoop(prefs PreferencesRef, runLoop RunLoopRef, runLoopMode StringRef) unsafe.Pointer {
	return _SCPreferencesScheduleWithRunLoop(prefs, runLoop, runLoopMode)
}/* debug [functions.gen.go/function]: SCPreferencesScheduleWithRunLoop */

// Assigns the specified callback to the specified preferences session.
//
// Added in macOS 10.4.
// Assigns the specified callback to the specified preferences session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetCallback(_:_:_:)
func SCPreferencesSetCallback(prefs PreferencesRef, callout PreferencesCallBack, context unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesSetCallback(prefs, callout, context)
}/* debug [functions.gen.go/function]: SCPreferencesSetCallback */

// Sets the computer name preference to the specified name.
//
// Added in macOS 10.1.
// Sets the computer name preference to the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetComputerName(_:_:_:)
func SCPreferencesSetComputerName(prefs PreferencesRef, name StringRef, nameEncoding StringEncoding) unsafe.Pointer {
	return _SCPreferencesSetComputerName(prefs, name, nameEncoding)
}/* debug [functions.gen.go/function]: SCPreferencesSetComputerName */

// Schedules commit and apply notifications for the specified preferences session using the specified dispatch queue.
//
// Added in macOS 10.6.
// Schedules commit and apply notifications for the specified preferences session using the specified dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetDispatchQueue(_:_:)
func SCPreferencesSetDispatchQueue(prefs PreferencesRef, queue unsafe.Pointer) unsafe.Pointer {
	return _SCPreferencesSetDispatchQueue(prefs, queue)
}/* debug [functions.gen.go/function]: SCPreferencesSetDispatchQueue */

// Sets the local host name to the specified name.
//
// Added in macOS 10.2.
// Sets the local host name to the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetLocalHostName(_:_:)
func SCPreferencesSetLocalHostName(prefs PreferencesRef, name StringRef) unsafe.Pointer {
	return _SCPreferencesSetLocalHostName(prefs, name)
}/* debug [functions.gen.go/function]: SCPreferencesSetLocalHostName */

// Updates the data associated with the specified preference key with the specified value.
//
// Added in macOS 10.1.
// Updates the data associated with the specified preference key with the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSetValue(_:_:_:)
func SCPreferencesSetValue(prefs PreferencesRef, key StringRef, value PropertyListRef) unsafe.Pointer {
	return _SCPreferencesSetValue(prefs, key, value)
}/* debug [functions.gen.go/function]: SCPreferencesSetValue */

// Synchronizes accessed preferences with committed changes.
//
// Added in macOS 10.4.
// Synchronizes accessed preferences with committed changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesSynchronize(_:)
func SCPreferencesSynchronize(prefs PreferencesRef) {
	_SCPreferencesSynchronize(prefs)
}/* debug [functions.gen.go/function]: SCPreferencesSynchronize */

// Releases exclusive access to the configuration preferences.
//
// Added in macOS 10.1.
// Releases exclusive access to the configuration preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesUnlock(_:)
func SCPreferencesUnlock(prefs PreferencesRef) unsafe.Pointer {
	return _SCPreferencesUnlock(prefs)
}/* debug [functions.gen.go/function]: SCPreferencesUnlock */

// Unschedules commit and apply notifications for the specified preferences session from the specified run loop and mode.
//
// Added in macOS 10.4.
// Unschedules commit and apply notifications for the specified preferences session from the specified run loop and mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCPreferencesUnscheduleFromRunLoop(_:_:_:)
func SCPreferencesUnscheduleFromRunLoop(prefs PreferencesRef, runLoop RunLoopRef, runLoopMode StringRef) unsafe.Pointer {
	return _SCPreferencesUnscheduleFromRunLoop(prefs, runLoop, runLoopMode)
}/* debug [functions.gen.go/function]: SCPreferencesUnscheduleFromRunLoop */

// Returns all virtual LAN (VLAN) interfaces on the system.
//
// Added in macOS 10.5.
// Returns all virtual LAN (VLAN) interfaces on the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceCopyAll(_:)
func SCVLANInterfaceCopyAll(prefs PreferencesRef) ArrayRef {
	return _SCVLANInterfaceCopyAll(prefs)
}/* debug [functions.gen.go/function]: SCVLANInterfaceCopyAll */

// Returns the network capable devices on the system that can be associated with a virtual LAN (VLAN) interface.
//
// Added in macOS 10.5.
// Returns the network capable devices on the system that can be associated with a virtual LAN (VLAN) interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceCopyAvailablePhysicalInterfaces()
func SCVLANInterfaceCopyAvailablePhysicalInterfaces() ArrayRef {
	return _SCVLANInterfaceCopyAvailablePhysicalInterfaces()
}/* debug [functions.gen.go/function]: SCVLANInterfaceCopyAvailablePhysicalInterfaces */

// Creates a new virtual LAN (VLAN) interface.
//
// Added in macOS 10.5.
// Creates a new virtual LAN (VLAN) interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceCreate(_:_:_:)
func SCVLANInterfaceCreate(prefs PreferencesRef, physical NetworkInterfaceRef, tag NumberRef) VLANInterfaceRef {
	return _SCVLANInterfaceCreate(prefs, physical, tag)
}/* debug [functions.gen.go/function]: SCVLANInterfaceCreate */

// Returns the configuration settings associated with the virtual LAN (VLAN) interface.
//
// Added in macOS 10.5.
// Returns the configuration settings associated with the virtual LAN (VLAN) interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceGetOptions(_:)
func SCVLANInterfaceGetOptions(vlan VLANInterfaceRef) DictionaryRef {
	return _SCVLANInterfaceGetOptions(vlan)
}/* debug [functions.gen.go/function]: SCVLANInterfaceGetOptions */

// Returns the physical interface for the specified virtual LAN (VLAN) interface.
//
// Added in macOS 10.5.
// Returns the physical interface for the specified virtual LAN (VLAN) interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceGetPhysicalInterface(_:)
func SCVLANInterfaceGetPhysicalInterface(vlan VLANInterfaceRef) NetworkInterfaceRef {
	return _SCVLANInterfaceGetPhysicalInterface(vlan)
}/* debug [functions.gen.go/function]: SCVLANInterfaceGetPhysicalInterface */

// Returns the tag for the specified virtual LAN (VLAN) interface.
//
// Added in macOS 10.5.
// Returns the tag for the specified virtual LAN (VLAN) interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceGetTag(_:)
func SCVLANInterfaceGetTag(vlan VLANInterfaceRef) NumberRef {
	return _SCVLANInterfaceGetTag(vlan)
}/* debug [functions.gen.go/function]: SCVLANInterfaceGetTag */

// Removes the virtual LAN (VLAN) interface from the configuration.
//
// Added in macOS 10.5.
// Removes the virtual LAN (VLAN) interface from the configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceRemove(_:)
func SCVLANInterfaceRemove(vlan VLANInterfaceRef) unsafe.Pointer {
	return _SCVLANInterfaceRemove(vlan)
}/* debug [functions.gen.go/function]: SCVLANInterfaceRemove */

// Sets the localized display name for the specified virtual LAN (VLAN) interface.
//
// Added in macOS 10.5.
// Sets the localized display name for the specified virtual LAN (VLAN) interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceSetLocalizedDisplayName(_:_:)
func SCVLANInterfaceSetLocalizedDisplayName(vlan VLANInterfaceRef, newName StringRef) unsafe.Pointer {
	return _SCVLANInterfaceSetLocalizedDisplayName(vlan, newName)
}/* debug [functions.gen.go/function]: SCVLANInterfaceSetLocalizedDisplayName */

// Sets the specified configuration settings for the specified virtual LAN (VLAN) interface.
//
// Added in macOS 10.5.
// Sets the specified configuration settings for the specified virtual LAN (VLAN) interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceSetOptions(_:_:)
func SCVLANInterfaceSetOptions(vlan VLANInterfaceRef, newOptions DictionaryRef) unsafe.Pointer {
	return _SCVLANInterfaceSetOptions(vlan, newOptions)
}/* debug [functions.gen.go/function]: SCVLANInterfaceSetOptions */

// Updates the specified virtual LAN (VLAN) interface with the specified information.
//
// Added in macOS 10.5.
// Updates the specified virtual LAN (VLAN) interface with the specified information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCVLANInterfaceSetPhysicalInterfaceAndTag(_:_:_:)
func SCVLANInterfaceSetPhysicalInterfaceAndTag(vlan VLANInterfaceRef, physical NetworkInterfaceRef, tag NumberRef) unsafe.Pointer {
	return _SCVLANInterfaceSetPhysicalInterfaceAndTag(vlan, physical, tag)
}/* debug [functions.gen.go/function]: SCVLANInterfaceSetPhysicalInterfaceAndTag */




