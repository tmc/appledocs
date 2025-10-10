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

// Discovered functions (104 total):

// AuthorizationExecuteWithPrivileges(AuthorizationRef  authorization,  const char  *pathToTool,  AuthorizationFlags  options,  char  * const _Nonnull  *arguments,  FILE  *  _Nullable  *communicationsPipe) OSStatus
//
// Availability:
//   - macOS 10.1+ (Deprecated in 10.7)
//
// Deprecated: This function is deprecated.

// AuthorizationPluginCreate(const AuthorizationCallbacks  *callbacks,  AuthorizationPluginRef _Nullable  *outPlugin,  const AuthorizationPluginInterface  *  _Nullable  *outPluginInterface) OSStatus
//
// Availability:
//   - macOS 10.4+

// AuthorizationCopyInfo(authorization AuthorizationRef, tag ,  AuthorizationString, info ,  AuthorizationItemSet  * *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+


// AuthorizationCopyPrivilegedReference(authorization AuthorizationRef  *, flags ,  AuthorizationFlags, ) OSStatus
//
// Availability:
//   - macOS 10.1+ (Deprecated in 10.7)
//
// Deprecated: This function is deprecated.

// AuthorizationCopyRights(authorization AuthorizationRef, rights ,  const AuthorizationRights  *, environment ,  const AuthorizationEnvironment  *, flags ,  AuthorizationFlags, authorizedRights ,  AuthorizationRights  * *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationCopyRightsAsync(authorization AuthorizationRef, rights ,  const AuthorizationRights  *, environment ,  const AuthorizationEnvironment  *, flags ,  AuthorizationFlags, callbackBlock ,  AuthorizationAsyncCallback, )
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.7+


// AuthorizationCreate(rights const AuthorizationRights  *, environment ,  const AuthorizationEnvironment  *, flags ,  AuthorizationFlags, authorization ,  AuthorizationRef  *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationCreateFromExternalForm(extForm const AuthorizationExternalForm  *, authorization ,  AuthorizationRef  *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationExecuteWithPrivileges(authorization AuthorizationRef, pathToTool ,  const char  *, options ,  AuthorizationFlags, arguments ,  char  *  const  *, communicationsPipe ,  FILE  * *, ) OSStatus
//
// Availability:
//   - macOS 10.1+ (Deprecated in 10.7)
//
// Deprecated: This function is deprecated.


// AuthorizationFree(authorization AuthorizationRef, flags ,  AuthorizationFlags, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationFreeItemSet(set AuthorizationItemSet  *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationMakeExternalForm(authorization AuthorizationRef, extForm ,  AuthorizationExternalForm  *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+


// AuthorizationRightGet(rightName const char  *, rightDefinition ,  CFDictionaryRef  *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationRightRemove(authRef AuthorizationRef, rightName ,  const char  *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// AuthorizationRightSet(authRef AuthorizationRef, rightName ,  const char  *, rightDefinition ,  CFTypeRef, descriptionKey ,  CFStringRef, bundle ,  CFBundleRef, localeTableName ,  CFStringRef, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+


// CMSDecoderCopyAllCerts(cmsDecoder CMSDecoderRef, certsOut ,  CFArrayRef  *, ) OSStatus
//
// Availability:
//   - macOS 10.5+

// CMSDecoderCopySignerCert(cmsDecoder CMSDecoderRef, signerIndex ,  size_t, signerCertOut ,  SecCertificateRef  *, ) OSStatus
//
// Availability:
//   - macOS 10.5+

// SSLSetCertificate(context SSLContextRef, certRefs ,  CFArrayRef, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 5.0+ (Deprecated in 13.0)
//   - iPadOS 5.0+ (Deprecated in 13.0)
//   - macOS 10.2+ (Deprecated in 10.15)
//
// Deprecated: This function is deprecated.


// SSLSetDiffieHellmanParams(context SSLContextRef, dhParams ,  const void  *, dhParamsLen ,  size_t, ) OSStatus
//
// Availability:
//   - macOS 10.2+ (Deprecated in 10.15)
//
// Deprecated: This function is deprecated.

// SecCertificateCopyNotValidAfterDate(certificate SecCertificateRef, ) CFDateRef
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//   - watchOS 11.0+

// SecCodeCopyPath(staticCode SecStaticCodeRef, flags ,  SecCSFlags, path ,  CFURLRef  *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+


// SecCodeCopyStaticCode(code SecCodeRef, flags ,  SecCSFlags, staticCode ,  SecStaticCodeRef  *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// SecCodeValidateFileResource(code SecStaticCodeRef, relativePath ,  CFStringRef, fileData ,  CFDataRef, flags ,  SecCSFlags, ) OSStatus
//
// Availability:
//   - macOS 10.13+

// SecCopyErrorMessageString(status OSStatus, reserved ,  void  *, ) CFStringRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.3+
//   - iPadOS 11.3+
//   - macOS 10.3+
//   - tvOS 11.3+
//   - visionOS 1.0+
//   - watchOS 4.3+


// SecEncryptTransformCreate(keyRef SecKeyRef, error ,  CFErrorRef  *, ) SecTransformRef
//
// Availability:
//   - macOS 10.7+ (Deprecated in 13.0)
//
// Deprecated: This function is deprecated.

// SecIdentityCreate(allocator CFAllocatorRef, certificate ,  SecCertificateRef, privateKey ,  SecKeyRef, ) SecIdentityRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.2+
//   - iPadOS 11.2+
//   - macOS 10.12+
//   - tvOS 11.2+
//   - visionOS 1.0+
//   - watchOS 4.2+

// SecItemAdd(attributes CFDictionaryRef, result ,  CFTypeRef  *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// SecItemCopyMatching(query CFDictionaryRef, result ,  CFTypeRef  *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// SecKeyCopyExternalRepresentation(key SecKeyRef, error ,  CFErrorRef  *, ) CFDataRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.0+

// SecKeyCreateWithData(keyData CFDataRef, attributes ,  CFDictionaryRef, error ,  CFErrorRef  *, ) SecKeyRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.0+


// SecKeyIsAlgorithmSupported(key SecKeyRef, operation ,  SecKeyOperationType, algorithm ,  SecKeyAlgorithm, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.0+

// SecRandomCopyBytes(rnd SecRandomRef, count ,  size_t, bytes ,  void  *, ) int
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// SecRequestSharedWebCredential(fqdn CFStringRef, account ,  CFStringRef, completionHandler ,  void  (^, credentials )( CFArrayRef, error ,  CFErrorRef, )
//
// Availability:
//   - Mac Catalyst 14.0+ (Deprecated in 14.0)
//   - iOS 8.0+ (Deprecated in 14.0)
//   - iPadOS 8.0+ (Deprecated in 14.0)
//   - macOS 11.0+ (Deprecated in 11.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Deprecated: This function is deprecated.


// SecStaticCodeCheckValidityWithErrors(staticCode SecStaticCodeRef, flags ,  SecCSFlags, requirement ,  SecRequirementRef, errors ,  CFErrorRef  *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// SecStaticCodeCreateWithPath(path CFURLRef, flags ,  SecCSFlags, staticCode ,  SecStaticCodeRef  *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// SecTaskCreateFromSelf(allocator CFAllocatorRef, ) SecTaskRef
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+


// SecTaskGetCodeSignStatus(task SecTaskRef, ) uint32_t
//
// Availability:
//   - Mac Catalyst 11.0+
//   - iOS 10.0+
//   - iPadOS 10.0+

// SecTransformConnectTransforms(sourceTransformRef SecTransformRef, sourceAttributeName ,  CFStringRef, destinationTransformRef ,  SecTransformRef, destinationAttributeName ,  CFStringRef, group ,  SecGroupTransformRef, error ,  CFErrorRef  *, ) SecGroupTransformRef
//
// Availability:
//   - macOS 10.7+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// SecTransformCustomSetAttribute(ref SecTransformImplementationRef, attribute ,  SecTransformStringOrAttributeRef, type ,  SecTransformMetaAttributeType, value ,  CFTypeRef, ) CFTypeRef
//
// Availability:
//   - macOS 10.7+ (Deprecated in 13.0)
//
// Deprecated: This function is deprecated.


// SecTransformExecute(transformRef SecTransformRef, errorRef ,  CFErrorRef  *, ) CFTypeRef
//
// Availability:
//   - macOS 10.7+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// SecureDownloadCopyName(downloadRef SecureDownloadRef, name ,  CFStringRef  *, ) OSStatus
//
// Availability:
//   - macOS 10.5+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// SecureDownloadCopyTicketLocation(url CFURLRef, ticketLocation ,  CFURLRef  *, ) OSStatus
//
// Availability:
//   - macOS 10.5+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.


// SecureDownloadFinished(downloadRef SecureDownloadRef, ) OSStatus
//
// Availability:
//   - macOS 10.5+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// SecureDownloadRelease(downloadRef SecureDownloadRef, ) OSStatus
//
// Availability:
//   - macOS 10.5+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// SecureDownloadUpdateWithData(downloadRef SecureDownloadRef, data ,  CFDataRef, ) OSStatus
//
// Availability:
//   - macOS 10.5+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.


// SessionCreate(flags SessionCreationFlags, attributes ,  SessionAttributeBits, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// SessionGetInfo(session SecuritySessionId, sessionId ,  SecuritySessionId  *, attributes ,  SessionAttributeBits  *, ) OSStatus
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// sec_certificate_copy_ref(certificate sec_certificate_t, ) SecCertificateRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// sec_certificate_create(certificate SecCertificateRef, ) sec_certificate_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_identity_access_certificates(identity sec_identity_t, handler ,  void  (^, certificate )( sec_certificate_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_identity_copy_certificates_ref(identity sec_identity_t, ) CFArrayRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// sec_identity_copy_ref(identity sec_identity_t, ) SecIdentityRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_identity_create(identity SecIdentityRef, ) sec_identity_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_identity_create_with_certificates(identity SecIdentityRef, certificates ,  CFArrayRef, ) sec_identity_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// sec_protocol_metadata_access_distinguished_names(metadata sec_protocol_metadata_t, handler ,  void  (^, distinguished_name )( dispatch_data_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_access_ocsp_response(metadata sec_protocol_metadata_t, handler ,  void  (^, ocsp_data )( dispatch_data_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_access_peer_certificate_chain(metadata sec_protocol_metadata_t, handler ,  void  (^, certificate )( sec_certificate_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// sec_protocol_metadata_access_pre_shared_keys(metadata sec_protocol_metadata_t, handler ,  void  (^, psk )( dispatch_data_t, psk_identity ,  dispatch_data_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_metadata_access_supported_signature_algorithms(metadata sec_protocol_metadata_t, handler ,  void  (^, signature_algorithm )( uint16_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_challenge_parameters_are_equal(metadataA sec_protocol_metadata_t, metadataB ,  sec_protocol_metadata_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// sec_protocol_metadata_copy_negotiated_protocol(metadata sec_protocol_metadata_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 18.5+
//   - iOS 18.5+
//   - iPadOS 18.5+
//   - macOS 15.5+
//   - tvOS 18.5+
//   - visionOS 2.5+
//   - watchOS 11.5+

// sec_protocol_metadata_copy_peer_public_key(metadata sec_protocol_metadata_t, ) dispatch_data_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_copy_server_name(metadata sec_protocol_metadata_t, ) const char  *
//
// Availability:
//   - Mac Catalyst 18.5+
//   - iOS 18.5+
//   - iPadOS 18.5+
//   - macOS 15.5+
//   - tvOS 18.5+
//   - visionOS 2.5+
//   - watchOS 11.5+


// sec_protocol_metadata_create_secret(metadata sec_protocol_metadata_t, label_len ,  size_t, label ,  const char  *, exporter_length ,  size_t, ) dispatch_data_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_create_secret_with_context(metadata sec_protocol_metadata_t, label_len ,  size_t, label ,  const char  *, context_len ,  size_t, context ,  const uint8_t  *, exporter_length ,  size_t, ) dispatch_data_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_metadata_get_early_data_accepted(metadata sec_protocol_metadata_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// sec_protocol_metadata_get_negotiated_ciphersuite(metadata sec_protocol_metadata_t, ) SSLCipherSuite
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

// sec_protocol_metadata_get_negotiated_protocol(metadata sec_protocol_metadata_t, ) const char  *
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

// sec_protocol_metadata_get_negotiated_protocol_version(metadata sec_protocol_metadata_t, ) SSLProtocol
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


// sec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata sec_protocol_metadata_t, ) tls_ciphersuite_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_metadata_get_negotiated_tls_protocol_version(metadata sec_protocol_metadata_t, ) tls_protocol_version_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_metadata_get_server_name(metadata sec_protocol_metadata_t, ) const char  *
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


// sec_protocol_metadata_peers_are_equal(metadataA sec_protocol_metadata_t, metadataB ,  sec_protocol_metadata_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_add_pre_shared_key(options sec_protocol_options_t, psk ,  dispatch_data_t, psk_identity ,  dispatch_data_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_add_tls_application_protocol(options sec_protocol_options_t, application_protocol ,  const char  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// sec_protocol_options_add_tls_ciphersuite(options sec_protocol_options_t, ciphersuite ,  SSLCipherSuite, )
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

// sec_protocol_options_add_tls_ciphersuite_group(options sec_protocol_options_t, group ,  SSLCiphersuiteGroup, )
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

// sec_protocol_options_append_tls_ciphersuite(options sec_protocol_options_t, ciphersuite ,  tls_ciphersuite_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// sec_protocol_options_append_tls_ciphersuite_group(options sec_protocol_options_t, group ,  tls_ciphersuite_group_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_are_equal(optionsA sec_protocol_options_t, optionsB ,  sec_protocol_options_t, ) bool
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_set_challenge_block(options sec_protocol_options_t, challenge_block ,  sec_protocol_challenge_t, challenge_queue ,  dispatch_queue_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// sec_protocol_options_set_key_update_block(options sec_protocol_options_t, key_update_block ,  sec_protocol_key_update_t, key_update_queue ,  dispatch_queue_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_local_identity(options sec_protocol_options_t, identity ,  sec_identity_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_max_tls_protocol_version(options sec_protocol_options_t, version ,  tls_protocol_version_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// sec_protocol_options_set_min_tls_protocol_version(options sec_protocol_options_t, version ,  tls_protocol_version_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_set_peer_authentication_required(options sec_protocol_options_t, peer_authentication_required ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_pre_shared_key_selection_block(options sec_protocol_options_t, psk_selection_block ,  sec_protocol_pre_shared_key_selection_t, psk_selection_queue ,  dispatch_queue_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// sec_protocol_options_set_tls_diffie_hellman_parameters(options sec_protocol_options_t, params ,  dispatch_data_t, )
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

// sec_protocol_options_set_tls_false_start_enabled(options sec_protocol_options_t, false_start_enabled ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_tls_is_fallback_attempt(options sec_protocol_options_t, is_fallback_attempt ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// sec_protocol_options_set_tls_max_version(options sec_protocol_options_t, version ,  SSLProtocol, )
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

// sec_protocol_options_set_tls_min_version(options sec_protocol_options_t, version ,  SSLProtocol, )
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

// sec_protocol_options_set_tls_ocsp_enabled(options sec_protocol_options_t, ocsp_enabled ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// sec_protocol_options_set_tls_pre_shared_key_identity_hint(options sec_protocol_options_t, psk_identity_hint ,  dispatch_data_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// sec_protocol_options_set_tls_renegotiation_enabled(options sec_protocol_options_t, renegotiation_enabled ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_tls_resumption_enabled(options sec_protocol_options_t, resumption_enabled ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// sec_protocol_options_set_tls_sct_enabled(options sec_protocol_options_t, sct_enabled ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_tls_server_name(options sec_protocol_options_t, server_name ,  const char  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_protocol_options_set_tls_tickets_enabled(options sec_protocol_options_t, tickets_enabled ,  bool, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+


// sec_protocol_options_set_verify_block(options sec_protocol_options_t, verify_block ,  sec_protocol_verify_t, verify_block_queue ,  dispatch_queue_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_release(obj void  *, )
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// sec_retain(obj void  *, ) void  *
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// sec_trust_copy_ref(trust sec_trust_t, ) SecTrustRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// sec_trust_create(trust SecTrustRef, ) sec_trust_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

