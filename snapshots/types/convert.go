package types

import (
	fmt "fmt"

	proto "github.com/gogo/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"
)

// Converts an ABCI snapshot to a snapshot
func SnapshotFromABCI(in *abci.Snapshot) (Snapshot, error) {
	snapshot := Snapshot{
		Height: in.Height,
		Format: in.Format,
		Chunks: in.Chunks,
		Hash:   in.Hash,
	}
	err := proto.Unmarshal(in.Metadata, &snapshot.Metadata)
	if err != nil {
		return Snapshot{}, fmt.Errorf("failed to unmarshal snapshot metadata: %w", err)
	}
	return snapshot, nil
}

// Converts a Snapshot to its ABCI representation.
func (s Snapshot) ToABCI() (abci.Snapshot, error) {
	out := abci.Snapshot{
		Height: s.Height,
		Format: s.Format,
		Chunks: s.Chunks,
		Hash:   s.Hash,
	}
	var err error
	out.Metadata, err = proto.Marshal(&s.Metadata)
	if err != nil {
		return abci.Snapshot{}, fmt.Errorf("failed to marshal snapshot metadata: %w", err)
	}
	return out, nil
}