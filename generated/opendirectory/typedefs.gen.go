// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory
import (
"unsafe"
)

// Type aliases and typedefs
// ODAuthenticationType - An Open Directory authentication type.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAuthenticationType
// ODAuthenticationType is a string typedef
type ODAuthenticationType = string
// ODContextRef - An Open Directory context type.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODContext
// ODContextRef has base type: const struct __ODContext *
type ODContextRef uintptr
// ODErrorUserInfoKeyType type alias
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODErrorUserInfoKeyType
// ODErrorUserInfoKeyType is a string typedef
type ODErrorUserInfoKeyType = string
// ODMatchType - An Open Directory match type.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMatchType
// ODMatchType has base type: uint32_t
type ODMatchType uintptr
// ODNodeRef - An Open Directory node type.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeRef
// ODNodeRef has base type: struct __ODNode *
type ODNodeRef uintptr
// ODNodeType - An Open Directory node type.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODNodeType
// ODNodeType has base type: uint32_t
type ODNodeType uintptr
// ODOptionKeyType type alias
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODOptionKeyType
// ODOptionKeyType is a string typedef
type ODOptionKeyType = string
// ODPolicyAttributeType type alias
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODPolicyAttributeType
// ODPolicyAttributeType is a string typedef
type ODPolicyAttributeType = string
// ODPolicyCategoryType type alias
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODPolicyCategoryType
// ODPolicyCategoryType is a string typedef
type ODPolicyCategoryType = string
// ODPolicyKeyType type alias
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODPolicyKeyType
// ODPolicyKeyType is a string typedef
type ODPolicyKeyType = string
// ODPolicyType type alias
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODPolicyType
// ODPolicyType is a string typedef
type ODPolicyType = string
// ODQueryCallback - A callback function called as results from a scheduled query are returned.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQueryCallback
// ODQueryCallback is a callback function
// C type: void (*)(struct __ODQuery *, const struct __CFArray *, struct __CFError *, void *)
type ODQueryCallback = func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
// ODQueryRef - An Open Directory query type.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODQueryRef
// ODQueryRef has base type: struct __ODQuery *
type ODQueryRef uintptr
// ODRecordRef - An Open Directory record type.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordRef
// ODRecordRef has base type: struct __ODRecord *
type ODRecordRef uintptr
// ODSessionRef - An Open Directory session type.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODSessionRef
// ODSessionRef has base type: struct __ODSession *
type ODSessionRef uintptr
// ODAttributeType - An Open Directory attribute type.
//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeType
// ODAttributeType is a string typedef
type ODAttributeType = string

