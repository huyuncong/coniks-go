package directory

import (
	"testing"

	"github.com/huyuncong/coniks-go/crypto"
	"github.com/huyuncong/coniks-go/merkletree"
)

// NewTestDirectory creates a ConiksDirectory used for testing server-side
// CONIKS operations.
func NewTestDirectory(t *testing.T) *ConiksDirectory {
	vrfKey := crypto.NewStaticTestVRFKey()
	signKey := crypto.NewStaticTestSigningKey()
	d := New(1, vrfKey, signKey, 10, true)
	d.pad = merkletree.StaticPAD(t, d.policies)
	return d
}
