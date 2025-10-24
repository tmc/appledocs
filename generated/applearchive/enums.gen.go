// Code generated from Apple documentation for AppleArchive. DO NOT EDIT.

package applearchive

/* debug [enums.gen.go]: Generating 11 enums for AppleArchive */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum AAACEQualifierTypes (4 cases) */
// AAACEQualifierTypes enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAACEQualifierTypes
type AAACEQualifierTypes uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAACEQualifierTypes/AA_ACE_QUALIFIER_TYPE_GROUP
	AA_ACE_QUALIFIER_TYPE_GROUP AAACEQualifierTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAACEQualifierTypes/AA_ACE_QUALIFIER_TYPE_SID
	AA_ACE_QUALIFIER_TYPE_SID AAACEQualifierTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAACEQualifierTypes/AA_ACE_QUALIFIER_TYPE_USER
	AA_ACE_QUALIFIER_TYPE_USER AAACEQualifierTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAACEQualifierTypes/AA_ACE_QUALIFIER_TYPE_UUID
	AA_ACE_QUALIFIER_TYPE_UUID AAACEQualifierTypes = 0
)

/* debug [enums.gen.go]: Processing enum AACompressionAlgorithms (6 cases) */
// AACompressionAlgorithms enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACompressionAlgorithms
type AACompressionAlgorithms uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACompressionAlgorithms/AA_COMPRESSION_ALGORITHM_LZ4
	AA_COMPRESSION_ALGORITHM_LZ4 AACompressionAlgorithms = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACompressionAlgorithms/AA_COMPRESSION_ALGORITHM_LZBITMAP
	AA_COMPRESSION_ALGORITHM_LZBITMAP AACompressionAlgorithms = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACompressionAlgorithms/AA_COMPRESSION_ALGORITHM_LZFSE
	AA_COMPRESSION_ALGORITHM_LZFSE AACompressionAlgorithms = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACompressionAlgorithms/AA_COMPRESSION_ALGORITHM_LZMA
	AA_COMPRESSION_ALGORITHM_LZMA AACompressionAlgorithms = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACompressionAlgorithms/AA_COMPRESSION_ALGORITHM_NONE
	AA_COMPRESSION_ALGORITHM_NONE AACompressionAlgorithms = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AACompressionAlgorithms/AA_COMPRESSION_ALGORITHM_ZLIB
	AA_COMPRESSION_ALGORITHM_ZLIB AACompressionAlgorithms = 0
)

/* debug [enums.gen.go]: Processing enum AAEntryMessages (14 cases) */
// AAEntryMessages enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages
type AAEntryMessages uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_CONVERT_EXCLUDE
	AA_ENTRY_MESSAGE_CONVERT_EXCLUDE AAEntryMessages = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_DECODE_READING
	AA_ENTRY_MESSAGE_DECODE_READING AAEntryMessages = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_ENCODE_SCANNING
	AA_ENTRY_MESSAGE_ENCODE_SCANNING AAEntryMessages = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_ENCODE_WRITING
	AA_ENTRY_MESSAGE_ENCODE_WRITING AAEntryMessages = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_EXTRACT_ACL
	AA_ENTRY_MESSAGE_EXTRACT_ACL AAEntryMessages = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_EXTRACT_ATTRIBUTES
	AA_ENTRY_MESSAGE_EXTRACT_ATTRIBUTES AAEntryMessages = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_EXTRACT_BEGIN
	AA_ENTRY_MESSAGE_EXTRACT_BEGIN AAEntryMessages = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_EXTRACT_END
	AA_ENTRY_MESSAGE_EXTRACT_END AAEntryMessages = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_EXTRACT_FAIL
	AA_ENTRY_MESSAGE_EXTRACT_FAIL AAEntryMessages = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_EXTRACT_XAT
	AA_ENTRY_MESSAGE_EXTRACT_XAT AAEntryMessages = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_PROCESS_EXCLUDE
	AA_ENTRY_MESSAGE_PROCESS_EXCLUDE AAEntryMessages = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_SEARCH_EXCLUDE
	AA_ENTRY_MESSAGE_SEARCH_EXCLUDE AAEntryMessages = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_SEARCH_FAIL
	AA_ENTRY_MESSAGE_SEARCH_FAIL AAEntryMessages = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryMessages/AA_ENTRY_MESSAGE_SEARCH_PRUNE_DIR
	AA_ENTRY_MESSAGE_SEARCH_PRUNE_DIR AAEntryMessages = 0
)

/* debug [enums.gen.go]: Processing enum AAEntryTypes (11 cases) */
// AAEntryTypes enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryTypes
type AAEntryTypes uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryTypes/AA_ENTRY_TYPE_BLK
	AA_ENTRY_TYPE_BLK AAEntryTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryTypes/AA_ENTRY_TYPE_CHR
	AA_ENTRY_TYPE_CHR AAEntryTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryTypes/AA_ENTRY_TYPE_DIR
	AA_ENTRY_TYPE_DIR AAEntryTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryTypes/AA_ENTRY_TYPE_DOOR
	AA_ENTRY_TYPE_DOOR AAEntryTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryTypes/AA_ENTRY_TYPE_FIFO
	AA_ENTRY_TYPE_FIFO AAEntryTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryTypes/AA_ENTRY_TYPE_LNK
	AA_ENTRY_TYPE_LNK AAEntryTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryTypes/AA_ENTRY_TYPE_METADATA
	AA_ENTRY_TYPE_METADATA AAEntryTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryTypes/AA_ENTRY_TYPE_PORT
	AA_ENTRY_TYPE_PORT AAEntryTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryTypes/AA_ENTRY_TYPE_REG
	AA_ENTRY_TYPE_REG AAEntryTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryTypes/AA_ENTRY_TYPE_SOCK
	AA_ENTRY_TYPE_SOCK AAEntryTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAEntryTypes/AA_ENTRY_TYPE_WHT
	AA_ENTRY_TYPE_WHT AAEntryTypes = 0
)

/* debug [enums.gen.go]: Processing enum AAFieldTypes (6 cases) */
// AAFieldTypes enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldTypes
type AAFieldTypes uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldTypes/AA_FIELD_TYPE_BLOB
	AA_FIELD_TYPE_BLOB AAFieldTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldTypes/AA_FIELD_TYPE_FLAG
	AA_FIELD_TYPE_FLAG AAFieldTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldTypes/AA_FIELD_TYPE_HASH
	AA_FIELD_TYPE_HASH AAFieldTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldTypes/AA_FIELD_TYPE_STRING
	AA_FIELD_TYPE_STRING AAFieldTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldTypes/AA_FIELD_TYPE_TIMESPEC
	AA_FIELD_TYPE_TIMESPEC AAFieldTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFieldTypes/AA_FIELD_TYPE_UINT
	AA_FIELD_TYPE_UINT AAFieldTypes = 0
)

/* debug [enums.gen.go]: Processing enum AAFlags (15 cases) */
// AAFlags enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags
type AAFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_ARCHIVE_DEDUPLICATE_DAT
	AA_FLAG_ARCHIVE_DEDUPLICATE_DAT AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_ARCHIVE_NO_RESOLVE_ACL_QUALIFIERS
	AA_FLAG_ARCHIVE_NO_RESOLVE_ACL_QUALIFIERS AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_CROSS_VOLUME_BOUNDARIES
	AA_FLAG_CROSS_VOLUME_BOUNDARIES AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_DECODE_INSERT_IDX
	AA_FLAG_DECODE_INSERT_IDX AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_EXCLUDE_METADATA_ENTRIES
	AA_FLAG_EXCLUDE_METADATA_ENTRIES AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_EXTRACT_AUTO_DEDUP_AS_HARD_LINKS
	AA_FLAG_EXTRACT_AUTO_DEDUP_AS_HARD_LINKS AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_EXTRACT_NO_AUTO_DEDUP
	AA_FLAG_EXTRACT_NO_AUTO_DEDUP AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_EXTRACT_NO_AUTO_SPARSE
	AA_FLAG_EXTRACT_NO_AUTO_SPARSE AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_IGNORE_EPERM
	AA_FLAG_IGNORE_EPERM AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_PROCESS_RANDOM_ACCESS_OUTPUT
	AA_FLAG_PROCESS_RANDOM_ACCESS_OUTPUT AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_REPLACE_ATTRIBUTES
	AA_FLAG_REPLACE_ATTRIBUTES AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_VERBOSITY_0
	AA_FLAG_VERBOSITY_0 AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_VERBOSITY_1
	AA_FLAG_VERBOSITY_1 AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_VERBOSITY_2
	AA_FLAG_VERBOSITY_2 AAFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAFlags/AA_FLAG_VERBOSITY_3
	AA_FLAG_VERBOSITY_3 AAFlags = 0
)

/* debug [enums.gen.go]: Processing enum AAHashFunctions (5 cases) */
// AAHashFunctions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHashFunctions
type AAHashFunctions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHashFunctions/AA_HASH_FUNCTION_CRC32
	AA_HASH_FUNCTION_CRC32 AAHashFunctions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHashFunctions/AA_HASH_FUNCTION_SHA1
	AA_HASH_FUNCTION_SHA1 AAHashFunctions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHashFunctions/AA_HASH_FUNCTION_SHA256
	AA_HASH_FUNCTION_SHA256 AAHashFunctions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHashFunctions/AA_HASH_FUNCTION_SHA384
	AA_HASH_FUNCTION_SHA384 AAHashFunctions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AAHashFunctions/AA_HASH_FUNCTION_SHA512
	AA_HASH_FUNCTION_SHA512 AAHashFunctions = 0
)

/* debug [enums.gen.go]: Processing enum AEAContextFieldRepresentations (3 cases) */
// AEAContextFieldRepresentations enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldRepresentations
type AEAContextFieldRepresentations uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldRepresentations/AEA_CONTEXT_FIELD_REPRESENTATION_GENERATE
	AEA_CONTEXT_FIELD_REPRESENTATION_GENERATE AEAContextFieldRepresentations = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldRepresentations/AEA_CONTEXT_FIELD_REPRESENTATION_RAW
	AEA_CONTEXT_FIELD_REPRESENTATION_RAW AEAContextFieldRepresentations = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldRepresentations/AEA_CONTEXT_FIELD_REPRESENTATION_X963
	AEA_CONTEXT_FIELD_REPRESENTATION_X963 AEAContextFieldRepresentations = 0
)

/* debug [enums.gen.go]: Processing enum AEAContextFields (18 cases) */
// AEAContextFields enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields
type AEAContextFields uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_ARCHIVE_IDENTIFIER
	AEA_CONTEXT_FIELD_ARCHIVE_IDENTIFIER AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_AUTH_DATA
	AEA_CONTEXT_FIELD_AUTH_DATA AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_BLOCKS_PER_CLUSTER
	AEA_CONTEXT_FIELD_BLOCKS_PER_CLUSTER AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_CHECKSUM_MODE
	AEA_CONTEXT_FIELD_CHECKSUM_MODE AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_COMPRESSION_ALGORITHM
	AEA_CONTEXT_FIELD_COMPRESSION_ALGORITHM AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_COMPRESSION_BLOCK_SIZE
	AEA_CONTEXT_FIELD_COMPRESSION_BLOCK_SIZE AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_CONTAINER_SIZE
	AEA_CONTEXT_FIELD_CONTAINER_SIZE AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_MAIN_KEY
	AEA_CONTEXT_FIELD_MAIN_KEY AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_PADDING_SIZE
	AEA_CONTEXT_FIELD_PADDING_SIZE AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_PASSWORD
	AEA_CONTEXT_FIELD_PASSWORD AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_PROFILE
	AEA_CONTEXT_FIELD_PROFILE AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_RAW_SIZE
	AEA_CONTEXT_FIELD_RAW_SIZE AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_RECIPIENT_PRIVATE_KEY
	AEA_CONTEXT_FIELD_RECIPIENT_PRIVATE_KEY AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_RECIPIENT_PUBLIC_KEY
	AEA_CONTEXT_FIELD_RECIPIENT_PUBLIC_KEY AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_SIGNATURE_ENCRYPTION_KEY
	AEA_CONTEXT_FIELD_SIGNATURE_ENCRYPTION_KEY AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_SIGNING_PRIVATE_KEY
	AEA_CONTEXT_FIELD_SIGNING_PRIVATE_KEY AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_SIGNING_PUBLIC_KEY
	AEA_CONTEXT_FIELD_SIGNING_PUBLIC_KEY AEAContextFields = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFields/AEA_CONTEXT_FIELD_SYMMETRIC_KEY
	AEA_CONTEXT_FIELD_SYMMETRIC_KEY AEAContextFields = 0
)

/* debug [enums.gen.go]: Processing enum AEAContextFieldValues (14 cases) */
// AEAContextFieldValues enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues
type AEAContextFieldValues uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_CHECKSUM_MURMURHASH64
	AEA_CONTEXT_CHECKSUM_MURMURHASH64 AEAContextFieldValues = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_CHECKSUM_NONE
	AEA_CONTEXT_CHECKSUM_NONE AEAContextFieldValues = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_CHECKSUM_SHA256
	AEA_CONTEXT_CHECKSUM_SHA256 AEAContextFieldValues = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_CIPHERSUITE_HKDF_SHA256_AESCTR_HMAC
	AEA_CONTEXT_CIPHERSUITE_HKDF_SHA256_AESCTR_HMAC AEAContextFieldValues = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_CIPHERSUITE_HKDF_SHA256_HMAC
	AEA_CONTEXT_CIPHERSUITE_HKDF_SHA256_HMAC AEAContextFieldValues = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_ENCRYPTION_ECDHE_P256
	AEA_CONTEXT_ENCRYPTION_ECDHE_P256 AEAContextFieldValues = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_ENCRYPTION_NONE
	AEA_CONTEXT_ENCRYPTION_NONE AEAContextFieldValues = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_ENCRYPTION_SCRYPT
	AEA_CONTEXT_ENCRYPTION_SCRYPT AEAContextFieldValues = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_ENCRYPTION_SYMMETRIC
	AEA_CONTEXT_ENCRYPTION_SYMMETRIC AEAContextFieldValues = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_PADDING_ADAPTIVE
	AEA_CONTEXT_PADDING_ADAPTIVE AEAContextFieldValues = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_PADDING_MIN_SIZE
	AEA_CONTEXT_PADDING_MIN_SIZE AEAContextFieldValues = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_PADDING_NONE
	AEA_CONTEXT_PADDING_NONE AEAContextFieldValues = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_SIGNATURE_ECDSA_P256
	AEA_CONTEXT_SIGNATURE_ECDSA_P256 AEAContextFieldValues = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAContextFieldValues/AEA_CONTEXT_SIGNATURE_NONE
	AEA_CONTEXT_SIGNATURE_NONE AEAContextFieldValues = 0
)

/* debug [enums.gen.go]: Processing enum AEAProfiles (6 cases) */
// AEAProfiles enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAProfiles
type AEAProfiles uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAProfiles/AEA_PROFILE__HKDF_SHA256_AESCTR_HMAC__ECDHE_P256__ECDSA_P256
	AEA_PROFILE__HKDF_SHA256_AESCTR_HMAC__ECDHE_P256__ECDSA_P256 AEAProfiles = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAProfiles/AEA_PROFILE__HKDF_SHA256_AESCTR_HMAC__ECDHE_P256__NONE
	AEA_PROFILE__HKDF_SHA256_AESCTR_HMAC__ECDHE_P256__NONE AEAProfiles = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAProfiles/AEA_PROFILE__HKDF_SHA256_AESCTR_HMAC__SCRYPT__NONE
	AEA_PROFILE__HKDF_SHA256_AESCTR_HMAC__SCRYPT__NONE AEAProfiles = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAProfiles/AEA_PROFILE__HKDF_SHA256_AESCTR_HMAC__SYMMETRIC__ECDSA_P256
	AEA_PROFILE__HKDF_SHA256_AESCTR_HMAC__SYMMETRIC__ECDSA_P256 AEAProfiles = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAProfiles/AEA_PROFILE__HKDF_SHA256_AESCTR_HMAC__SYMMETRIC__NONE
	AEA_PROFILE__HKDF_SHA256_AESCTR_HMAC__SYMMETRIC__NONE AEAProfiles = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppleArchive/AEAProfiles/AEA_PROFILE__HKDF_SHA256_HMAC__NONE__ECDSA_P256
	AEA_PROFILE__HKDF_SHA256_HMAC__NONE__ECDSA_P256 AEAProfiles = 0
)


