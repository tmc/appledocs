// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

// Security Functions
//
// This file contains function declarations discovered from Apple's documentation.
// To use these functions, you need to:
//   1. Map C types to Go types
//   2. Create function variables
//   3. Register them with purego.RegisterLibFunc
//
// Example:
//   var CGContextSetRGBFillColor func(c CGContextRef, red, green, blue, alpha CGFloat)
//   purego.RegisterLibFunc(&CGContextSetRGBFillColor, lib, "CGContextSetRGBFillColor")

// Discovered functions (111 total):

// AuthorizationExecuteWithPrivileges(AuthorizationRef authorization,  const  char *pathToTool,  AuthorizationFlags options,  char * const   _Nonnull *arguments,  FILE *  _Nullable *communicationsPipe);) OSStatus
//
// Availability:
//   - macOS 10.1+ (Deprecated in 10.7)
//
// Deprecated: This function is deprecated.

// AuthorizationPluginCreate(const  AuthorizationCallbacks *callbacks,  AuthorizationPluginRef   _Nullable *outPlugin,  const  AuthorizationPluginInterface *  _Nullable *outPluginInterface);) OSStatus
//
// Availability:
//   - macOS 10.4+

// AuthorizationCopyInfo(authorization _, tag :  AuthorizationRef,  _, info :  AuthorizationString?,  _, :  UnsafeMutablePointer< UnsafeMutablePointer< AuthorizationItemSet>?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationCopyPrivilegedReference(authorization AuthorizationRef *, flags ,  AuthorizationFlags, );) OSStatus
//
// Availability:
//   - macOS 10.1+ (Deprecated in 10.7)
//
// Deprecated: This function is deprecated.

// AuthorizationCopyRights(authorization _, rights :  AuthorizationRef,  _, environment :  UnsafePointer< AuthorizationRights>,  _, flags :  UnsafePointer< AuthorizationEnvironment>?,  _, authorizedRights :  AuthorizationFlags,  _, :  UnsafeMutablePointer< UnsafeMutablePointer< AuthorizationRights>?>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationCopyRightsAsync(authorization _, rights :  AuthorizationRef,  _, environment :  UnsafePointer< AuthorizationRights>,  _, flags :  UnsafePointer< AuthorizationEnvironment>?,  _, callbackBlock :  AuthorizationFlags,  _, :  @escaping  AuthorizationAsyncCallback) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.7+

// AuthorizationCreate(rights _, environment :  UnsafePointer< AuthorizationRights>?,  _, flags :  UnsafePointer< AuthorizationEnvironment>?,  _, authorization :  AuthorizationFlags,  _, :  UnsafeMutablePointer< AuthorizationRef?>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationCreateFromExternalForm(extForm _, authorization :  UnsafePointer< AuthorizationExternalForm>,  _, :  UnsafeMutablePointer< AuthorizationRef?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationExecuteWithPrivileges(authorization AuthorizationRef, pathToTool ,  const  char *, options ,  AuthorizationFlags, arguments ,  char *  const *, communicationsPipe ,  FILE * *, );) OSStatus
//
// Availability:
//   - macOS 10.1+ (Deprecated in 10.7)
//
// Deprecated: This function is deprecated.

// AuthorizationFree(authorization _, flags :  AuthorizationRef,  _, :  AuthorizationFlags) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationFreeItemSet(set _, :  UnsafeMutablePointer< AuthorizationItemSet>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationMakeExternalForm(authorization _, extForm :  AuthorizationRef,  _, :  UnsafeMutablePointer< AuthorizationExternalForm>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationRightGet(rightName _, rightDefinition :  UnsafePointer< CChar>,  _, :  UnsafeMutablePointer< CFDictionary?>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationRightRemove(authRef _, rightName :  AuthorizationRef,  _, :  UnsafePointer< CChar>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationRightSet(authRef _, rightName :  AuthorizationRef,  _, rightDefinition :  UnsafePointer< CChar>,  _, descriptionKey :  CFTypeRef,  _, bundle :  CFString?,  _, localeTableName :  CFBundle?,  _, :  CFString?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// CMSDecoderCopyAllCerts(cmsDecoder _, certsOut :  CMSDecoder,  _, :  UnsafeMutablePointer< CFArray?>) ->  OSStatus) func
//
// Availability:
//   - macOS 10.5+

// CMSDecoderCopySignerCert(cmsDecoder _, signerIndex :  CMSDecoder,  _, signerCertOut :  Int,  _, :  UnsafeMutablePointer< SecCertificate?>) ->  OSStatus) func
//
// Availability:
//   - macOS 10.5+

// SSLSetCertificate(context _, certRefs :  SSLContext,  _, :  CFArray?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 5.0+ (Deprecated in 13.0)
//   - iPadOS 5.0+ (Deprecated in 13.0)
//   - macOS 10.2+ (Deprecated in 10.15)
//
// Deprecated: This function is deprecated.

// SSLSetDiffieHellmanParams(context _, dhParams :  SSLContext,  _, dhParamsLen :  UnsafeRawPointer?,  _, :  Int) ->  OSStatus) func
//
// Availability:
//   - macOS 10.2+ (Deprecated in 10.15)
//
// Deprecated: This function is deprecated.

// SecCertificateCopyNotValidAfterDate(certificate _, :  SecCertificate) ->  CFDate?) func
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//   - watchOS 11.0+

// SecCodeCopyPath(staticCode _, flags :  SecStaticCode,  _, path :  SecCSFlags,  _, :  UnsafeMutablePointer< CFURL?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// SecCodeCopyStaticCode(code _, flags :  SecCode,  _, staticCode :  SecCSFlags,  _, :  UnsafeMutablePointer< SecStaticCode?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// SecCodeValidateFileResource(code _, relativePath :  SecStaticCode,  _, fileData :  CFString,  _, flags :  CFData,  _, :  SecCSFlags) ->  OSStatus) func
//
// Availability:
//   - macOS 10.13+

// SecCopyErrorMessageString(status _, reserved :  OSStatus,  _, :  UnsafeMutableRawPointer?) ->  CFString?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.3+
//   - iPadOS 11.3+
//   - macOS 10.3+
//   - tvOS 11.3+
//   - visionOS 1.0+
//   - watchOS 4.3+

// SecCreateSharedWebCredentialPassword() func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 11.0+
//   - visionOS 1.0+

// SecEncryptTransformCreate(keyRef _, error :  SecKey,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>?) ->  SecTransform) func
//
// Availability:
//   - macOS 10.7+ (Deprecated in 13.0)
//
// Deprecated: This function is deprecated.

// SecIdentityCreate(allocator _, certificate :  CFAllocator?,  _, privateKey :  SecCertificate,  _, :  SecKey) ->  SecIdentity?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.2+
//   - iPadOS 11.2+
//   - macOS 10.12+
//   - tvOS 11.2+
//   - visionOS 1.0+
//   - watchOS 4.2+

// SecItemAdd(attributes _, result :  CFDictionary,  _, :  UnsafeMutablePointer< CFTypeRef?>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// SecItemCopyMatching(query _, result :  CFDictionary,  _, :  UnsafeMutablePointer< CFTypeRef?>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// SecKeyCopyExternalRepresentation(key _, error :  SecKey,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>?) ->  CFData?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.0+

// SecKeyCreateWithData(keyData _, attributes :  CFData,  _, error :  CFDictionary,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>?) ->  SecKey?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.0+

// SecKeyIsAlgorithmSupported(key _, operation :  SecKey,  _, algorithm :  SecKeyOperationType,  _, :  SecKeyAlgorithm) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.0+

// SecPolicyCreateBasicX509() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// SecRandomCopyBytes(rnd _, count :  SecRandomRef?,  _, bytes :  Int,  _, :  UnsafeMutableRawPointer) ->  Int32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// SecRequestSharedWebCredential(fqdn _, account :  CFString?,  _, completionHandler :  CFString?,  _, :  @escaping  CFArray?,  CFError?) ->  Void) func
//
// Availability:
//   - Mac Catalyst 14.0+ (Deprecated in 14.0)
//   - iOS 8.0+ (Deprecated in 14.0)
//   - iPadOS 8.0+ (Deprecated in 14.0)
//   - macOS 11.0+ (Deprecated in 11.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Deprecated: This function is deprecated.

// SecStaticCodeCheckValidityWithErrors(staticCode _, flags :  SecStaticCode,  _, requirement :  SecCSFlags,  _, errors :  SecRequirement?,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// SecStaticCodeCreateWithPath(path _, flags :  CFURL,  _, staticCode :  SecCSFlags,  _, :  UnsafeMutablePointer< SecStaticCode?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// SecTaskCreateFromSelf(allocator _, :  CFAllocator?) ->  SecTask?) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// SecTaskGetCodeSignStatus(task _, :  SecTask) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 11.0+
//   - iOS 10.0+
//   - iPadOS 10.0+

// SecTransformConnectTransforms(sourceTransformRef _, sourceAttributeName :  SecTransform,  _, destinationTransformRef :  CFString,  _, destinationAttributeName :  SecTransform,  _, group :  CFString,  _, error :  SecGroupTransform,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>?) ->  SecGroupTransform?) func
//
// Availability:
//   - macOS 10.7+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// SecTransformCreateGroupTransform() func
//
// Availability:
//   - macOS 10.7+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// SecTransformCustomSetAttribute(ref _, attribute :  SecTransformImplementationRef,  _, type :  SecTransformStringOrAttribute,  _, value :  SecTransformMetaAttributeType,  _, :  CFTypeRef?) ->  CFTypeRef?) func
//
// Availability:
//   - macOS 10.7+ (Deprecated in 13.0)
//
// Deprecated: This function is deprecated.

// SecTransformExecute(transformRef _, errorRef :  SecTransform,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>?) ->  CFTypeRef) func
//
// Availability:
//   - macOS 10.7+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// SecureDownloadCopyName(downloadRef SecureDownloadRef, name ,  CFStringRef *, );) OSStatus
//
// Availability:
//   - macOS 10.5+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// SecureDownloadCopyTicketLocation(url CFURLRef, ticketLocation ,  CFURLRef *, );) OSStatus
//
// Availability:
//   - macOS 10.5+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// SecureDownloadFinished(downloadRef SecureDownloadRef, );) OSStatus
//
// Availability:
//   - macOS 10.5+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// SecureDownloadRelease(downloadRef SecureDownloadRef, );) OSStatus
//
// Availability:
//   - macOS 10.5+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// SecureDownloadUpdateWithData(downloadRef SecureDownloadRef, data ,  CFDataRef, );) OSStatus
//
// Availability:
//   - macOS 10.5+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// SessionCreate(flags _, attributes :  SessionCreationFlags,  _, :  SessionAttributeBits) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// SessionGetInfo(session _, sessionId :  SecuritySessionId,  _, attributes :  UnsafeMutablePointer< SecuritySessionId>?,  _, :  UnsafeMutablePointer< SessionAttributeBits>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// sec_certificate_copy_ref(certificate _, :  sec_certificate_t) ->  Unmanaged< SecCertificate>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_certificate_create(certificate _, :  SecCertificate) ->  sec_certificate_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_identity_access_certificates(identity _, handler :  sec_identity_t,  _, :  @escaping  sec_certificate_t) ->  Void) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_identity_copy_certificates_ref(identity _, :  sec_identity_t) ->  Unmanaged< CFArray>?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_identity_copy_ref(identity _, :  sec_identity_t) ->  Unmanaged< SecIdentity>?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_identity_create(identity _, :  SecIdentity) ->  sec_identity_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_identity_create_with_certificates(identity _, certificates :  SecIdentity,  _, :  CFArray) ->  sec_identity_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_access_distinguished_names(metadata _, handler :  sec_protocol_metadata_t,  _, :  @escaping  dispatch_data_t) ->  Void) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_access_ocsp_response(metadata _, handler :  sec_protocol_metadata_t,  _, :  @escaping  dispatch_data_t) ->  Void) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_access_peer_certificate_chain(metadata _, handler :  sec_protocol_metadata_t,  _, :  @escaping  sec_certificate_t) ->  Void) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_access_pre_shared_keys(metadata _, handler :  sec_protocol_metadata_t,  _, :  @escaping  dispatch_data_t,  dispatch_data_t) ->  Void) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_metadata_access_supported_signature_algorithms(metadata _, handler :  sec_protocol_metadata_t,  _, :  @escaping  UInt16) ->  Void) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_challenge_parameters_are_equal(metadataA _, metadataB :  sec_protocol_metadata_t,  _, :  sec_protocol_metadata_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_copy_negotiated_protocol(metadata _, :  sec_protocol_metadata_t) ->  UnsafePointer< CChar>?) func
//
// Availability:
//   - Mac Catalyst 18.5+
//   - iOS 18.5+
//   - iPadOS 18.5+
//   - macOS 15.5+
//   - tvOS 18.5+
//   - visionOS 2.5+
//   - watchOS 11.5+

// sec_protocol_metadata_copy_peer_public_key(metadata _, :  sec_protocol_metadata_t) ->  dispatch_data_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_copy_server_name(metadata _, :  sec_protocol_metadata_t) ->  UnsafePointer< CChar>?) func
//
// Availability:
//   - Mac Catalyst 18.5+
//   - iOS 18.5+
//   - iPadOS 18.5+
//   - macOS 15.5+
//   - tvOS 18.5+
//   - visionOS 2.5+
//   - watchOS 11.5+

// sec_protocol_metadata_create_secret(metadata _, label_len :  sec_protocol_metadata_t,  _, label :  Int,  _, exporter_length :  UnsafePointer< CChar>,  _, :  Int) ->  dispatch_data_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_create_secret_with_context(metadata _, label_len :  sec_protocol_metadata_t,  _, label :  Int,  _, context_len :  UnsafePointer< CChar>,  _, context :  Int,  _, exporter_length :  UnsafePointer< UInt8>,  _, :  Int) ->  dispatch_data_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_get_early_data_accepted(metadata _, :  sec_protocol_metadata_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_get_negotiated_ciphersuite(metadata _, :  sec_protocol_metadata_t) ->  SSLCipherSuite) func
//
// Availability:
//   - Mac Catalyst 12.0+ (Deprecated in 13.0)
//   - iOS 12.0+ (Deprecated in 13.0)
//   - iPadOS 12.0+ (Deprecated in 13.0)
//   - macOS 10.14+ (Deprecated in 10.15)
//   - tvOS 12.0+ (Deprecated in 13.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 5.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.

// sec_protocol_metadata_get_negotiated_protocol(metadata _, :  sec_protocol_metadata_t) ->  UnsafePointer< CChar>?) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.5)
//   - iOS 12.0+ (Deprecated in 18.5)
//   - iPadOS 12.0+ (Deprecated in 18.5)
//   - macOS 10.14+ (Deprecated in 15.5)
//   - tvOS 12.0+ (Deprecated in 18.5)
//   - visionOS 1.0+ (Deprecated in 2.5)
//   - watchOS 5.0+ (Deprecated in 11.5)
//
// Deprecated: This function is deprecated.

// sec_protocol_metadata_get_negotiated_protocol_version(metadata _, :  sec_protocol_metadata_t) ->  SSLProtocol) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 12.0+ (Deprecated in 13.0)
//   - iPadOS 12.0+ (Deprecated in 13.0)
//   - macOS 10.14+ (Deprecated in 10.15)
//   - tvOS 12.0+ (Deprecated in 13.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 5.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.

// sec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata _, :  sec_protocol_metadata_t) ->  tls_ciphersuite_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_metadata_get_negotiated_tls_protocol_version(metadata _, :  sec_protocol_metadata_t) ->  tls_protocol_version_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_metadata_get_server_name(metadata _, :  sec_protocol_metadata_t) ->  UnsafePointer< CChar>?) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.5)
//   - iOS 12.0+ (Deprecated in 18.5)
//   - iPadOS 12.0+ (Deprecated in 18.5)
//   - macOS 10.14+ (Deprecated in 15.5)
//   - tvOS 12.0+ (Deprecated in 18.5)
//   - visionOS 1.0+ (Deprecated in 2.5)
//   - watchOS 5.0+ (Deprecated in 11.5)
//
// Deprecated: This function is deprecated.

// sec_protocol_metadata_peers_are_equal(metadataA _, metadataB :  sec_protocol_metadata_t,  _, :  sec_protocol_metadata_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_add_pre_shared_key(options _, psk :  sec_protocol_options_t,  _, psk_identity :  dispatch_data_t,  _, :  dispatch_data_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_add_tls_application_protocol(options _, application_protocol :  sec_protocol_options_t,  _, :  UnsafePointer< CChar>)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_add_tls_ciphersuite(options _, ciphersuite :  sec_protocol_options_t,  _, :  SSLCipherSuite) func
//
// Availability:
//   - Mac Catalyst 12.0+ (Deprecated in 13.0)
//   - iOS 12.0+ (Deprecated in 13.0)
//   - iPadOS 12.0+ (Deprecated in 13.0)
//   - macOS 10.14+ (Deprecated in 10.15)
//   - tvOS 12.0+ (Deprecated in 13.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 5.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.

// sec_protocol_options_add_tls_ciphersuite_group(options _, group :  sec_protocol_options_t,  _, :  SSLCiphersuiteGroup) func
//
// Availability:
//   - Mac Catalyst 12.0+ (Deprecated in 13.0)
//   - iOS 12.0+ (Deprecated in 13.0)
//   - iPadOS 12.0+ (Deprecated in 13.0)
//   - macOS 10.14+ (Deprecated in 10.15)
//   - tvOS 12.0+ (Deprecated in 13.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 5.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.

// sec_protocol_options_append_tls_ciphersuite(options _, ciphersuite :  sec_protocol_options_t,  _, :  tls_ciphersuite_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_append_tls_ciphersuite_group(options _, group :  sec_protocol_options_t,  _, :  tls_ciphersuite_group_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_are_equal(optionsA _, optionsB :  sec_protocol_options_t,  _, :  sec_protocol_options_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_get_default_max_dtls_protocol_version() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_get_default_max_tls_protocol_version() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_get_default_min_dtls_protocol_version() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_get_default_min_tls_protocol_version() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_set_challenge_block(options _, challenge_block :  sec_protocol_options_t,  _, challenge_queue :  @escaping  sec_protocol_challenge_t,  _, :  dispatch_queue_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_key_update_block(options _, key_update_block :  sec_protocol_options_t,  _, key_update_queue :  @escaping  sec_protocol_key_update_t,  _, :  dispatch_queue_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_local_identity(options _, identity :  sec_protocol_options_t,  _, :  sec_identity_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_max_tls_protocol_version(options _, version :  sec_protocol_options_t,  _, :  tls_protocol_version_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_set_min_tls_protocol_version(options _, version :  sec_protocol_options_t,  _, :  tls_protocol_version_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_set_peer_authentication_required(options _, peer_authentication_required :  sec_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_pre_shared_key_selection_block(options _, psk_selection_block :  sec_protocol_options_t,  _, psk_selection_queue :  @escaping  sec_protocol_pre_shared_key_selection_t,  _, :  dispatch_queue_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_set_tls_diffie_hellman_parameters(options _, params :  sec_protocol_options_t,  _, :  dispatch_data_t) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 12.0+ (Deprecated in 13.0)
//   - iPadOS 12.0+ (Deprecated in 13.0)
//   - macOS 10.14+ (Deprecated in 10.15)
//   - tvOS 12.0+ (Deprecated in 13.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 5.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.

// sec_protocol_options_set_tls_false_start_enabled(options _, false_start_enabled :  sec_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_tls_is_fallback_attempt(options _, is_fallback_attempt :  sec_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_tls_max_version(options _, version :  sec_protocol_options_t,  _, :  SSLProtocol) func
//
// Availability:
//   - Mac Catalyst 12.0+ (Deprecated in 13.0)
//   - iOS 12.0+ (Deprecated in 13.0)
//   - iPadOS 12.0+ (Deprecated in 13.0)
//   - macOS 10.14+ (Deprecated in 10.15)
//   - tvOS 12.0+ (Deprecated in 13.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 5.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.

// sec_protocol_options_set_tls_min_version(options _, version :  sec_protocol_options_t,  _, :  SSLProtocol) func
//
// Availability:
//   - Mac Catalyst 12.0+ (Deprecated in 13.0)
//   - iOS 12.0+ (Deprecated in 13.0)
//   - iPadOS 12.0+ (Deprecated in 13.0)
//   - macOS 10.14+ (Deprecated in 10.15)
//   - tvOS 12.0+ (Deprecated in 13.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 5.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.

// sec_protocol_options_set_tls_ocsp_enabled(options _, ocsp_enabled :  sec_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_tls_pre_shared_key_identity_hint(options _, psk_identity_hint :  sec_protocol_options_t,  _, :  dispatch_data_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_set_tls_renegotiation_enabled(options _, renegotiation_enabled :  sec_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_tls_resumption_enabled(options _, resumption_enabled :  sec_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_tls_sct_enabled(options _, sct_enabled :  sec_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_tls_server_name(options _, server_name :  sec_protocol_options_t,  _, :  UnsafePointer< CChar>)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_tls_tickets_enabled(options _, tickets_enabled :  sec_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_verify_block(options _, verify_block :  sec_protocol_options_t,  _, verify_block_queue :  @escaping  sec_protocol_verify_t,  _, :  dispatch_queue_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_release(obj _, :  UnsafeMutableRawPointer!)) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// sec_retain(obj _, :  UnsafeMutableRawPointer!) ->  UnsafeMutableRawPointer!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// sec_trust_copy_ref(trust _, :  sec_trust_t) ->  Unmanaged< SecTrust>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_trust_create(trust _, :  SecTrust) ->  sec_trust_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+
