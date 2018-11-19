package client

import (
	"github.com/huyuncong/coniks-go/application"
	"github.com/huyuncong/coniks-go/protocol"
)

// CreateRegistrationMsg returns a JSON encoding of
// a protocol.RegistrationRequest for the given (name, key) pair.
func CreateRegistrationMsg(name string, key []byte) ([]byte, error) {
	return application.MarshalRequest(protocol.RegistrationType,
		&protocol.RegistrationRequest{
			Username: name,
			Key:      key,
		})
}

// CreateKeyLookupMsg returns a JSON encoding of
// a protocol.KeyLookupRequest for the given name.
func CreateKeyLookupMsg(name string) ([]byte, error) {
	return application.MarshalRequest(protocol.KeyLookupType,
		&protocol.KeyLookupRequest{
			Username: name,
		})
}

// CreateKeyLookupInEpochMsg returns a JSON encoding of
// a protocol.KeyLookupInEpochRequest for the given name.
func CreateKeyLookupInEpochMsg(name string, epoch uint64) ([]byte, error) {
	return application.MarshalRequest(protocol.KeyLookupInEpochType,
		&protocol.KeyLookupInEpochRequest{
			Username: name,
			Epoch: epoch,
		})
}

// CreateEpochIncreaseMsg returns a JSON encoding of 
// a protocol.EpochIncreaseRequest for the given name.
func CreateEpochIncreaseMsg(epoch uint64) ([]byte, error) {
	return application.MarshalRequest(protocol.EpochIncreaseType,
		&protocol.EpochIncreaseRequest{
			Epoch: epoch,
		})
}

func CreateWorkloadInitMsg(perEpochNewRecord uint64, epochs uint64) ([]byte, error) {
	return application.MarshalRequest(protocol.WorkloadInitType,
		&protocol.WorkloadInitRequest{
			PerEpochNewRecord: perEpochNewRecord,
			Epochs: epochs,
		})
}
