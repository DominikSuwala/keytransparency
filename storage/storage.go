// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package storage provides an API to persistant storage, implemented with spanner.
package storage

import (
	corepb "github.com/google/e2e-key-server/proto/core"
	context "golang.org/x/net/context"
)

const (
	// ChannelSize is the buffer size of the channel used to send an
	// EntryStorage to the tree builder.
	ChannelSize = 100
)

var (
	// EntryStorageDB contains the path to the leveldb database that stores
	// EntryStorage.
	EntryStorageDBPath string
)

type ConsistentStorage interface {
	Reader
	Writer
	Watchable
}

type StaticStorage interface {
	Reader
	Writer
	Closer
}

type Reader interface {
	// Read reads a EntryStroage from the storage.
	Read(ctx context.Context, primaryKey uint64) (*corepb.EntryStorage, error)
}

type Writer interface {
	// Write inserts a new EntryStorage in the storage. Fails if the row
	// already exists.
	Write(ctx context.Context, entry *corepb.EntryStorage) error
}

type Closer interface {
	// Close closes the storage instance and release all resources.
	Close()
}

type Watchable interface {
	// NewEntries  returns a channel containing EntryStorage entries, which
	// are pushed into the channel whenever an EntryStorage is written in
	// the stirage.
	NewEntries() chan *corepb.EntryStorage
}

// TODO(cesarghali): bring back ConkisStorage and make it compatible with the
// new proto.
