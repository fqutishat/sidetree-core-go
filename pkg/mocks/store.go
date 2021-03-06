/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mocks

import (
	"errors"
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/api/batch"
)

// MockOperationStore mocks store for testing purposes.
type MockOperationStore struct {
	sync.RWMutex
	operations map[string][]batch.Operation
	Err        error
}

// NewMockOperationStore creates mock operations store
func NewMockOperationStore(err error) *MockOperationStore {
	return &MockOperationStore{operations: make(map[string][]batch.Operation), Err: err}
}

//Put mocks storing operation
func (m *MockOperationStore) Put(op batch.Operation) error {
	if m.Err != nil {
		return m.Err
	}

	m.Lock()
	defer m.Unlock()

	m.operations[op.UniqueSuffix] = append(m.operations[op.UniqueSuffix], op)

	return nil
}

//Get mocks retrieving operations from the store
func (m *MockOperationStore) Get(uniqueSuffix string) ([]batch.Operation, error) {
	if m.Err != nil {
		return nil, m.Err
	}

	m.RLock()
	defer m.RUnlock()

	if ops, ok := m.operations[uniqueSuffix]; ok {
		return ops, nil
	}

	return nil, errors.New("uniqueSuffix not found in the store")
}
