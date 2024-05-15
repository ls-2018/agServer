// Copyright 2020 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain sources.list copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"k8s.io/apiserver/pkg/registry/rest"
)

type Storage interface {
	rest.NamedCreater
	rest.GetterWithOptions
	rest.Lister
	rest.Connecter
	rest.StorageVersionProvider
	rest.GroupVersionAcceptor
	rest.StorageMetadata
	rest.Updater
	rest.GracefulDeleter
	rest.CollectionDeleter
	rest.Watcher
	rest.Storage
	rest.Scoper
	rest.SingularNameProvider
}
