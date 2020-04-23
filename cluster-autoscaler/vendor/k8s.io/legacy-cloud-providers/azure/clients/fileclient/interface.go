// +build !providerless

/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fileclient

// Interface is the client interface for creating file shares, interface for test injection.
// mockgen -source=$GOPATH/src/k8s.io/kubernetes/staging/src/k8s.io/legacy-cloud-providers/azure/clients/fileclient/interface.go -package=mockfileclient Interface > $GOPATH/src/k8s.io/kubernetes/staging/src/k8s.io/legacy-cloud-providers/azure/clients/fileclient/mockfileclient/interface.go
type Interface interface {
	CreateFileShare(accountName, accountKey, name string, sizeGiB int) error
	DeleteFileShare(accountName, accountKey, name string) error
	ResizeFileShare(accountName, accountKey, name string, sizeGiB int) error
}
