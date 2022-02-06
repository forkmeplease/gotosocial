// Code generated by astool. DO NOT EDIT.

package streams

import (
	typeemoji "github.com/superseriousbusiness/activity/streams/impl/toot/type_emoji"
	typeidentityproof "github.com/superseriousbusiness/activity/streams/impl/toot/type_identityproof"
	vocab "github.com/superseriousbusiness/activity/streams/vocab"
)

// IsOrExtendsTootEmoji returns true if the other provided type is the Emoji type
// or extends from the Emoji type.
func IsOrExtendsTootEmoji(other vocab.Type) bool {
	return typeemoji.IsOrExtendsEmoji(other)
}

// IsOrExtendsTootIdentityProof returns true if the other provided type is the
// IdentityProof type or extends from the IdentityProof type.
func IsOrExtendsTootIdentityProof(other vocab.Type) bool {
	return typeidentityproof.IsOrExtendsIdentityProof(other)
}