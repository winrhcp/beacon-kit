// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package primitives

import (
	"github.com/berachain/beacon-kit/mod/primitives/math"
	"github.com/berachain/beacon-kit/mod/primitives/version"
)

// BeaconBlockDeneb represents a block in the beacon chain during
// the Deneb fork.
//
//go:generate go run github.com/ferranbt/fastssz/sszgen --path block.go -objs BeaconBlockDeneb -include ./primitives.go,./payload.go,./withdrawal.go,./kzg.go,./bytes.go,./eth1data.go,./math,./execution.go,./deposit.go,./withdrawal_credentials.go,./body.go,$GETH_PKG_INCLUDE/common,$GETH_PKG_INCLUDE/common/hexutil -output block.ssz.go
type BeaconBlockDeneb struct {
	// Slot represents the position of the block in the chain.
	Slot math.Slot

	// ProposerIndex is the index of the validator who proposed the block.
	ProposerIndex math.ValidatorIndex

	// ParentBlockRoot is the hash of the parent block
	ParentBlockRoot Root

	// StateRoot is the hash of the state at the block.
	StateRoot Root

	// Body is the body of the BeaconBlockDeneb, containing the block's
	// operations.
	Body *BeaconBlockBodyDeneb
}

// Version identifies the version of the BeaconBlockDeneb.
func (b *BeaconBlockDeneb) Version() uint32 {
	return version.Deneb
}

// IsNil checks if the BeaconBlockDeneb instance is nil.
func (b *BeaconBlockDeneb) IsNil() bool {
	return b == nil
}

// GetSlot retrieves the slot of the BeaconBlockDeneb.
func (b *BeaconBlockDeneb) GetSlot() math.Slot {
	return b.Slot
}

// GetSlot retrieves the slot of the BeaconBlockDeneb.
func (b *BeaconBlockDeneb) GetProposerIndex() math.ValidatorIndex {
	return b.ProposerIndex
}

// GetBody retrieves the body of the BeaconBlockDeneb.
func (b *BeaconBlockDeneb) GetBody() BeaconBlockBody {
	return b.Body
}

// GetParentBlockRoot retrieves the parent block root of the BeaconBlockDeneb.
func (b *BeaconBlockDeneb) GetParentBlockRoot() Root {
	return b.ParentBlockRoot
}

// GetStateRoot retrieves the state root of the BeaconBlockDeneb.
func (b *BeaconBlockDeneb) GetStateRoot() Root {
	return b.StateRoot
}

// GetHeader builds a BeaconBlockHeader from the BeaconBlockDeneb.
func (b BeaconBlockDeneb) GetHeader() *BeaconBlockHeader {
	bodyRoot, err := b.GetBody().HashTreeRoot()
	if err != nil {
		return nil
	}

	return &BeaconBlockHeader{
		Slot:          b.Slot,
		ProposerIndex: b.ProposerIndex,
		ParentRoot:    b.ParentBlockRoot,
		StateRoot:     b.StateRoot,
		BodyRoot:      bodyRoot,
	}
}