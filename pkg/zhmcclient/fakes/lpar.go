// Copyright 2021-2023 IBM Corp. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fakes

import (
	"sync"

	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

type LparAPI struct {
	AttachStorageGroupToPartitionStub        func(string, *zhmcclient.StorageGroupPayload) (int, *zhmcclient.HmcError)
	attachStorageGroupToPartitionMutex       sync.RWMutex
	attachStorageGroupToPartitionArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupPayload
	}
	attachStorageGroupToPartitionReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	attachStorageGroupToPartitionReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	CreateLPARStub        func(string, *zhmcclient.LparProperties) (string, int, *zhmcclient.HmcError)
	createLPARMutex       sync.RWMutex
	createLPARArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.LparProperties
	}
	createLPARReturns struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	createLPARReturnsOnCall map[int]struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	DeleteLPARStub        func(string) (int, *zhmcclient.HmcError)
	deleteLPARMutex       sync.RWMutex
	deleteLPARArgsForCall []struct {
		arg1 string
	}
	deleteLPARReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	deleteLPARReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	DetachStorageGroupToPartitionStub        func(string, *zhmcclient.StorageGroupPayload) (int, *zhmcclient.HmcError)
	detachStorageGroupToPartitionMutex       sync.RWMutex
	detachStorageGroupToPartitionArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupPayload
	}
	detachStorageGroupToPartitionReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	detachStorageGroupToPartitionReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	FetchAsciiConsoleURIStub        func(string, *zhmcclient.AsciiConsoleURIPayload) (*zhmcclient.AsciiConsoleURIResponse, int, *zhmcclient.HmcError)
	fetchAsciiConsoleURIMutex       sync.RWMutex
	fetchAsciiConsoleURIArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.AsciiConsoleURIPayload
	}
	fetchAsciiConsoleURIReturns struct {
		result1 *zhmcclient.AsciiConsoleURIResponse
		result2 int
		result3 *zhmcclient.HmcError
	}
	fetchAsciiConsoleURIReturnsOnCall map[int]struct {
		result1 *zhmcclient.AsciiConsoleURIResponse
		result2 int
		result3 *zhmcclient.HmcError
	}
	GetLparPropertiesStub        func(string) (*zhmcclient.LparProperties, int, *zhmcclient.HmcError)
	getLparPropertiesMutex       sync.RWMutex
	getLparPropertiesArgsForCall []struct {
		arg1 string
	}
	getLparPropertiesReturns struct {
		result1 *zhmcclient.LparProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	getLparPropertiesReturnsOnCall map[int]struct {
		result1 *zhmcclient.LparProperties
		result2 int
		result3 *zhmcclient.HmcError
	}
	ListLPARsStub        func(string, map[string]string) ([]zhmcclient.LPAR, int, *zhmcclient.HmcError)
	listLPARsMutex       sync.RWMutex
	listLPARsArgsForCall []struct {
		arg1 string
		arg2 map[string]string
	}
	listLPARsReturns struct {
		result1 []zhmcclient.LPAR
		result2 int
		result3 *zhmcclient.HmcError
	}
	listLPARsReturnsOnCall map[int]struct {
		result1 []zhmcclient.LPAR
		result2 int
		result3 *zhmcclient.HmcError
	}
	ListNicsStub        func(string) ([]string, int, *zhmcclient.HmcError)
	listNicsMutex       sync.RWMutex
	listNicsArgsForCall []struct {
		arg1 string
	}
	listNicsReturns struct {
		result1 []string
		result2 int
		result3 *zhmcclient.HmcError
	}
	listNicsReturnsOnCall map[int]struct {
		result1 []string
		result2 int
		result3 *zhmcclient.HmcError
	}
	MountIsoImageStub        func(string, string, string) (int, *zhmcclient.HmcError)
	mountIsoImageMutex       sync.RWMutex
	mountIsoImageArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	mountIsoImageReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	mountIsoImageReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	StartLPARStub        func(string) (string, int, *zhmcclient.HmcError)
	startLPARMutex       sync.RWMutex
	startLPARArgsForCall []struct {
		arg1 string
	}
	startLPARReturns struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	startLPARReturnsOnCall map[int]struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	StopLPARStub        func(string) (string, int, *zhmcclient.HmcError)
	stopLPARMutex       sync.RWMutex
	stopLPARArgsForCall []struct {
		arg1 string
	}
	stopLPARReturns struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	stopLPARReturnsOnCall map[int]struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}
	UnmountIsoImageStub        func(string) (int, *zhmcclient.HmcError)
	unmountIsoImageMutex       sync.RWMutex
	unmountIsoImageArgsForCall []struct {
		arg1 string
	}
	unmountIsoImageReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	unmountIsoImageReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	UpdateLparPropertiesStub        func(string, *zhmcclient.LparProperties) (int, *zhmcclient.HmcError)
	updateLparPropertiesMutex       sync.RWMutex
	updateLparPropertiesArgsForCall []struct {
		arg1 string
		arg2 *zhmcclient.LparProperties
	}
	updateLparPropertiesReturns struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	updateLparPropertiesReturnsOnCall map[int]struct {
		result1 int
		result2 *zhmcclient.HmcError
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LparAPI) AttachStorageGroupToPartition(arg1 string, arg2 *zhmcclient.StorageGroupPayload) (int, *zhmcclient.HmcError) {
	fake.attachStorageGroupToPartitionMutex.Lock()
	ret, specificReturn := fake.attachStorageGroupToPartitionReturnsOnCall[len(fake.attachStorageGroupToPartitionArgsForCall)]
	fake.attachStorageGroupToPartitionArgsForCall = append(fake.attachStorageGroupToPartitionArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupPayload
	}{arg1, arg2})
	stub := fake.AttachStorageGroupToPartitionStub
	fakeReturns := fake.attachStorageGroupToPartitionReturns
	fake.recordInvocation("AttachStorageGroupToPartition", []interface{}{arg1, arg2})
	fake.attachStorageGroupToPartitionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LparAPI) AttachStorageGroupToPartitionCallCount() int {
	fake.attachStorageGroupToPartitionMutex.RLock()
	defer fake.attachStorageGroupToPartitionMutex.RUnlock()
	return len(fake.attachStorageGroupToPartitionArgsForCall)
}

func (fake *LparAPI) AttachStorageGroupToPartitionCalls(stub func(string, *zhmcclient.StorageGroupPayload) (int, *zhmcclient.HmcError)) {
	fake.attachStorageGroupToPartitionMutex.Lock()
	defer fake.attachStorageGroupToPartitionMutex.Unlock()
	fake.AttachStorageGroupToPartitionStub = stub
}

func (fake *LparAPI) AttachStorageGroupToPartitionArgsForCall(i int) (string, *zhmcclient.StorageGroupPayload) {
	fake.attachStorageGroupToPartitionMutex.RLock()
	defer fake.attachStorageGroupToPartitionMutex.RUnlock()
	argsForCall := fake.attachStorageGroupToPartitionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *LparAPI) AttachStorageGroupToPartitionReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.attachStorageGroupToPartitionMutex.Lock()
	defer fake.attachStorageGroupToPartitionMutex.Unlock()
	fake.AttachStorageGroupToPartitionStub = nil
	fake.attachStorageGroupToPartitionReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *LparAPI) AttachStorageGroupToPartitionReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.attachStorageGroupToPartitionMutex.Lock()
	defer fake.attachStorageGroupToPartitionMutex.Unlock()
	fake.AttachStorageGroupToPartitionStub = nil
	if fake.attachStorageGroupToPartitionReturnsOnCall == nil {
		fake.attachStorageGroupToPartitionReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.attachStorageGroupToPartitionReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *LparAPI) CreateLPAR(arg1 string, arg2 *zhmcclient.LparProperties) (string, int, *zhmcclient.HmcError) {
	fake.createLPARMutex.Lock()
	ret, specificReturn := fake.createLPARReturnsOnCall[len(fake.createLPARArgsForCall)]
	fake.createLPARArgsForCall = append(fake.createLPARArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.LparProperties
	}{arg1, arg2})
	stub := fake.CreateLPARStub
	fakeReturns := fake.createLPARReturns
	fake.recordInvocation("CreateLPAR", []interface{}{arg1, arg2})
	fake.createLPARMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *LparAPI) CreateLPARCallCount() int {
	fake.createLPARMutex.RLock()
	defer fake.createLPARMutex.RUnlock()
	return len(fake.createLPARArgsForCall)
}

func (fake *LparAPI) CreateLPARCalls(stub func(string, *zhmcclient.LparProperties) (string, int, *zhmcclient.HmcError)) {
	fake.createLPARMutex.Lock()
	defer fake.createLPARMutex.Unlock()
	fake.CreateLPARStub = stub
}

func (fake *LparAPI) CreateLPARArgsForCall(i int) (string, *zhmcclient.LparProperties) {
	fake.createLPARMutex.RLock()
	defer fake.createLPARMutex.RUnlock()
	argsForCall := fake.createLPARArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *LparAPI) CreateLPARReturns(result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.createLPARMutex.Lock()
	defer fake.createLPARMutex.Unlock()
	fake.CreateLPARStub = nil
	fake.createLPARReturns = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) CreateLPARReturnsOnCall(i int, result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.createLPARMutex.Lock()
	defer fake.createLPARMutex.Unlock()
	fake.CreateLPARStub = nil
	if fake.createLPARReturnsOnCall == nil {
		fake.createLPARReturnsOnCall = make(map[int]struct {
			result1 string
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.createLPARReturnsOnCall[i] = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) DeleteLPAR(arg1 string) (int, *zhmcclient.HmcError) {
	fake.deleteLPARMutex.Lock()
	ret, specificReturn := fake.deleteLPARReturnsOnCall[len(fake.deleteLPARArgsForCall)]
	fake.deleteLPARArgsForCall = append(fake.deleteLPARArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteLPARStub
	fakeReturns := fake.deleteLPARReturns
	fake.recordInvocation("DeleteLPAR", []interface{}{arg1})
	fake.deleteLPARMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LparAPI) DeleteLPARCallCount() int {
	fake.deleteLPARMutex.RLock()
	defer fake.deleteLPARMutex.RUnlock()
	return len(fake.deleteLPARArgsForCall)
}

func (fake *LparAPI) DeleteLPARCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.deleteLPARMutex.Lock()
	defer fake.deleteLPARMutex.Unlock()
	fake.DeleteLPARStub = stub
}

func (fake *LparAPI) DeleteLPARArgsForCall(i int) string {
	fake.deleteLPARMutex.RLock()
	defer fake.deleteLPARMutex.RUnlock()
	argsForCall := fake.deleteLPARArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LparAPI) DeleteLPARReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteLPARMutex.Lock()
	defer fake.deleteLPARMutex.Unlock()
	fake.DeleteLPARStub = nil
	fake.deleteLPARReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *LparAPI) DeleteLPARReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.deleteLPARMutex.Lock()
	defer fake.deleteLPARMutex.Unlock()
	fake.DeleteLPARStub = nil
	if fake.deleteLPARReturnsOnCall == nil {
		fake.deleteLPARReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.deleteLPARReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *LparAPI) DetachStorageGroupToPartition(arg1 string, arg2 *zhmcclient.StorageGroupPayload) (int, *zhmcclient.HmcError) {
	fake.detachStorageGroupToPartitionMutex.Lock()
	ret, specificReturn := fake.detachStorageGroupToPartitionReturnsOnCall[len(fake.detachStorageGroupToPartitionArgsForCall)]
	fake.detachStorageGroupToPartitionArgsForCall = append(fake.detachStorageGroupToPartitionArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.StorageGroupPayload
	}{arg1, arg2})
	stub := fake.DetachStorageGroupToPartitionStub
	fakeReturns := fake.detachStorageGroupToPartitionReturns
	fake.recordInvocation("DetachStorageGroupToPartition", []interface{}{arg1, arg2})
	fake.detachStorageGroupToPartitionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LparAPI) DetachStorageGroupToPartitionCallCount() int {
	fake.detachStorageGroupToPartitionMutex.RLock()
	defer fake.detachStorageGroupToPartitionMutex.RUnlock()
	return len(fake.detachStorageGroupToPartitionArgsForCall)
}

func (fake *LparAPI) DetachStorageGroupToPartitionCalls(stub func(string, *zhmcclient.StorageGroupPayload) (int, *zhmcclient.HmcError)) {
	fake.detachStorageGroupToPartitionMutex.Lock()
	defer fake.detachStorageGroupToPartitionMutex.Unlock()
	fake.DetachStorageGroupToPartitionStub = stub
}

func (fake *LparAPI) DetachStorageGroupToPartitionArgsForCall(i int) (string, *zhmcclient.StorageGroupPayload) {
	fake.detachStorageGroupToPartitionMutex.RLock()
	defer fake.detachStorageGroupToPartitionMutex.RUnlock()
	argsForCall := fake.detachStorageGroupToPartitionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *LparAPI) DetachStorageGroupToPartitionReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.detachStorageGroupToPartitionMutex.Lock()
	defer fake.detachStorageGroupToPartitionMutex.Unlock()
	fake.DetachStorageGroupToPartitionStub = nil
	fake.detachStorageGroupToPartitionReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *LparAPI) DetachStorageGroupToPartitionReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.detachStorageGroupToPartitionMutex.Lock()
	defer fake.detachStorageGroupToPartitionMutex.Unlock()
	fake.DetachStorageGroupToPartitionStub = nil
	if fake.detachStorageGroupToPartitionReturnsOnCall == nil {
		fake.detachStorageGroupToPartitionReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.detachStorageGroupToPartitionReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *LparAPI) FetchAsciiConsoleURI(arg1 string, arg2 *zhmcclient.AsciiConsoleURIPayload) (*zhmcclient.AsciiConsoleURIResponse, int, *zhmcclient.HmcError) {
	fake.fetchAsciiConsoleURIMutex.Lock()
	ret, specificReturn := fake.fetchAsciiConsoleURIReturnsOnCall[len(fake.fetchAsciiConsoleURIArgsForCall)]
	fake.fetchAsciiConsoleURIArgsForCall = append(fake.fetchAsciiConsoleURIArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.AsciiConsoleURIPayload
	}{arg1, arg2})
	stub := fake.FetchAsciiConsoleURIStub
	fakeReturns := fake.fetchAsciiConsoleURIReturns
	fake.recordInvocation("FetchAsciiConsoleURI", []interface{}{arg1, arg2})
	fake.fetchAsciiConsoleURIMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *LparAPI) FetchAsciiConsoleURICallCount() int {
	fake.fetchAsciiConsoleURIMutex.RLock()
	defer fake.fetchAsciiConsoleURIMutex.RUnlock()
	return len(fake.fetchAsciiConsoleURIArgsForCall)
}

func (fake *LparAPI) FetchAsciiConsoleURICalls(stub func(string, *zhmcclient.AsciiConsoleURIPayload) (*zhmcclient.AsciiConsoleURIResponse, int, *zhmcclient.HmcError)) {
	fake.fetchAsciiConsoleURIMutex.Lock()
	defer fake.fetchAsciiConsoleURIMutex.Unlock()
	fake.FetchAsciiConsoleURIStub = stub
}

func (fake *LparAPI) FetchAsciiConsoleURIArgsForCall(i int) (string, *zhmcclient.AsciiConsoleURIPayload) {
	fake.fetchAsciiConsoleURIMutex.RLock()
	defer fake.fetchAsciiConsoleURIMutex.RUnlock()
	argsForCall := fake.fetchAsciiConsoleURIArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *LparAPI) FetchAsciiConsoleURIReturns(result1 *zhmcclient.AsciiConsoleURIResponse, result2 int, result3 *zhmcclient.HmcError) {
	fake.fetchAsciiConsoleURIMutex.Lock()
	defer fake.fetchAsciiConsoleURIMutex.Unlock()
	fake.FetchAsciiConsoleURIStub = nil
	fake.fetchAsciiConsoleURIReturns = struct {
		result1 *zhmcclient.AsciiConsoleURIResponse
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) FetchAsciiConsoleURIReturnsOnCall(i int, result1 *zhmcclient.AsciiConsoleURIResponse, result2 int, result3 *zhmcclient.HmcError) {
	fake.fetchAsciiConsoleURIMutex.Lock()
	defer fake.fetchAsciiConsoleURIMutex.Unlock()
	fake.FetchAsciiConsoleURIStub = nil
	if fake.fetchAsciiConsoleURIReturnsOnCall == nil {
		fake.fetchAsciiConsoleURIReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.AsciiConsoleURIResponse
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.fetchAsciiConsoleURIReturnsOnCall[i] = struct {
		result1 *zhmcclient.AsciiConsoleURIResponse
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) GetLparProperties(arg1 string) (*zhmcclient.LparProperties, int, *zhmcclient.HmcError) {
	fake.getLparPropertiesMutex.Lock()
	ret, specificReturn := fake.getLparPropertiesReturnsOnCall[len(fake.getLparPropertiesArgsForCall)]
	fake.getLparPropertiesArgsForCall = append(fake.getLparPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetLparPropertiesStub
	fakeReturns := fake.getLparPropertiesReturns
	fake.recordInvocation("GetLparProperties", []interface{}{arg1})
	fake.getLparPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *LparAPI) GetLparPropertiesCallCount() int {
	fake.getLparPropertiesMutex.RLock()
	defer fake.getLparPropertiesMutex.RUnlock()
	return len(fake.getLparPropertiesArgsForCall)
}

func (fake *LparAPI) GetLparPropertiesCalls(stub func(string) (*zhmcclient.LparProperties, int, *zhmcclient.HmcError)) {
	fake.getLparPropertiesMutex.Lock()
	defer fake.getLparPropertiesMutex.Unlock()
	fake.GetLparPropertiesStub = stub
}

func (fake *LparAPI) GetLparPropertiesArgsForCall(i int) string {
	fake.getLparPropertiesMutex.RLock()
	defer fake.getLparPropertiesMutex.RUnlock()
	argsForCall := fake.getLparPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LparAPI) GetLparPropertiesReturns(result1 *zhmcclient.LparProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getLparPropertiesMutex.Lock()
	defer fake.getLparPropertiesMutex.Unlock()
	fake.GetLparPropertiesStub = nil
	fake.getLparPropertiesReturns = struct {
		result1 *zhmcclient.LparProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) GetLparPropertiesReturnsOnCall(i int, result1 *zhmcclient.LparProperties, result2 int, result3 *zhmcclient.HmcError) {
	fake.getLparPropertiesMutex.Lock()
	defer fake.getLparPropertiesMutex.Unlock()
	fake.GetLparPropertiesStub = nil
	if fake.getLparPropertiesReturnsOnCall == nil {
		fake.getLparPropertiesReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.LparProperties
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.getLparPropertiesReturnsOnCall[i] = struct {
		result1 *zhmcclient.LparProperties
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) ListLPARs(arg1 string, arg2 map[string]string) ([]zhmcclient.LPAR, int, *zhmcclient.HmcError) {
	fake.listLPARsMutex.Lock()
	ret, specificReturn := fake.listLPARsReturnsOnCall[len(fake.listLPARsArgsForCall)]
	fake.listLPARsArgsForCall = append(fake.listLPARsArgsForCall, struct {
		arg1 string
		arg2 map[string]string
	}{arg1, arg2})
	stub := fake.ListLPARsStub
	fakeReturns := fake.listLPARsReturns
	fake.recordInvocation("ListLPARs", []interface{}{arg1, arg2})
	fake.listLPARsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *LparAPI) ListLPARsCallCount() int {
	fake.listLPARsMutex.RLock()
	defer fake.listLPARsMutex.RUnlock()
	return len(fake.listLPARsArgsForCall)
}

func (fake *LparAPI) ListLPARsCalls(stub func(string, map[string]string) ([]zhmcclient.LPAR, int, *zhmcclient.HmcError)) {
	fake.listLPARsMutex.Lock()
	defer fake.listLPARsMutex.Unlock()
	fake.ListLPARsStub = stub
}

func (fake *LparAPI) ListLPARsArgsForCall(i int) (string, map[string]string) {
	fake.listLPARsMutex.RLock()
	defer fake.listLPARsMutex.RUnlock()
	argsForCall := fake.listLPARsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *LparAPI) ListLPARsReturns(result1 []zhmcclient.LPAR, result2 int, result3 *zhmcclient.HmcError) {
	fake.listLPARsMutex.Lock()
	defer fake.listLPARsMutex.Unlock()
	fake.ListLPARsStub = nil
	fake.listLPARsReturns = struct {
		result1 []zhmcclient.LPAR
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) ListLPARsReturnsOnCall(i int, result1 []zhmcclient.LPAR, result2 int, result3 *zhmcclient.HmcError) {
	fake.listLPARsMutex.Lock()
	defer fake.listLPARsMutex.Unlock()
	fake.ListLPARsStub = nil
	if fake.listLPARsReturnsOnCall == nil {
		fake.listLPARsReturnsOnCall = make(map[int]struct {
			result1 []zhmcclient.LPAR
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.listLPARsReturnsOnCall[i] = struct {
		result1 []zhmcclient.LPAR
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) ListNics(arg1 string) ([]string, int, *zhmcclient.HmcError) {
	fake.listNicsMutex.Lock()
	ret, specificReturn := fake.listNicsReturnsOnCall[len(fake.listNicsArgsForCall)]
	fake.listNicsArgsForCall = append(fake.listNicsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListNicsStub
	fakeReturns := fake.listNicsReturns
	fake.recordInvocation("ListNics", []interface{}{arg1})
	fake.listNicsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *LparAPI) ListNicsCallCount() int {
	fake.listNicsMutex.RLock()
	defer fake.listNicsMutex.RUnlock()
	return len(fake.listNicsArgsForCall)
}

func (fake *LparAPI) ListNicsCalls(stub func(string) ([]string, int, *zhmcclient.HmcError)) {
	fake.listNicsMutex.Lock()
	defer fake.listNicsMutex.Unlock()
	fake.ListNicsStub = stub
}

func (fake *LparAPI) ListNicsArgsForCall(i int) string {
	fake.listNicsMutex.RLock()
	defer fake.listNicsMutex.RUnlock()
	argsForCall := fake.listNicsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LparAPI) ListNicsReturns(result1 []string, result2 int, result3 *zhmcclient.HmcError) {
	fake.listNicsMutex.Lock()
	defer fake.listNicsMutex.Unlock()
	fake.ListNicsStub = nil
	fake.listNicsReturns = struct {
		result1 []string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) ListNicsReturnsOnCall(i int, result1 []string, result2 int, result3 *zhmcclient.HmcError) {
	fake.listNicsMutex.Lock()
	defer fake.listNicsMutex.Unlock()
	fake.ListNicsStub = nil
	if fake.listNicsReturnsOnCall == nil {
		fake.listNicsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.listNicsReturnsOnCall[i] = struct {
		result1 []string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) MountIsoImage(arg1 string, arg2 string, arg3 string) (int, *zhmcclient.HmcError) {
	fake.mountIsoImageMutex.Lock()
	ret, specificReturn := fake.mountIsoImageReturnsOnCall[len(fake.mountIsoImageArgsForCall)]
	fake.mountIsoImageArgsForCall = append(fake.mountIsoImageArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.MountIsoImageStub
	fakeReturns := fake.mountIsoImageReturns
	fake.recordInvocation("MountIsoImage", []interface{}{arg1, arg2, arg3})
	fake.mountIsoImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LparAPI) MountIsoImageCallCount() int {
	fake.mountIsoImageMutex.RLock()
	defer fake.mountIsoImageMutex.RUnlock()
	return len(fake.mountIsoImageArgsForCall)
}

func (fake *LparAPI) MountIsoImageCalls(stub func(string, string, string) (int, *zhmcclient.HmcError)) {
	fake.mountIsoImageMutex.Lock()
	defer fake.mountIsoImageMutex.Unlock()
	fake.MountIsoImageStub = stub
}

func (fake *LparAPI) MountIsoImageArgsForCall(i int) (string, string, string) {
	fake.mountIsoImageMutex.RLock()
	defer fake.mountIsoImageMutex.RUnlock()
	argsForCall := fake.mountIsoImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *LparAPI) MountIsoImageReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.mountIsoImageMutex.Lock()
	defer fake.mountIsoImageMutex.Unlock()
	fake.MountIsoImageStub = nil
	fake.mountIsoImageReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *LparAPI) MountIsoImageReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.mountIsoImageMutex.Lock()
	defer fake.mountIsoImageMutex.Unlock()
	fake.MountIsoImageStub = nil
	if fake.mountIsoImageReturnsOnCall == nil {
		fake.mountIsoImageReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.mountIsoImageReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *LparAPI) StartLPAR(arg1 string) (string, int, *zhmcclient.HmcError) {
	fake.startLPARMutex.Lock()
	ret, specificReturn := fake.startLPARReturnsOnCall[len(fake.startLPARArgsForCall)]
	fake.startLPARArgsForCall = append(fake.startLPARArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.StartLPARStub
	fakeReturns := fake.startLPARReturns
	fake.recordInvocation("StartLPAR", []interface{}{arg1})
	fake.startLPARMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *LparAPI) StartLPARCallCount() int {
	fake.startLPARMutex.RLock()
	defer fake.startLPARMutex.RUnlock()
	return len(fake.startLPARArgsForCall)
}

func (fake *LparAPI) StartLPARCalls(stub func(string) (string, int, *zhmcclient.HmcError)) {
	fake.startLPARMutex.Lock()
	defer fake.startLPARMutex.Unlock()
	fake.StartLPARStub = stub
}

func (fake *LparAPI) StartLPARArgsForCall(i int) string {
	fake.startLPARMutex.RLock()
	defer fake.startLPARMutex.RUnlock()
	argsForCall := fake.startLPARArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LparAPI) StartLPARReturns(result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.startLPARMutex.Lock()
	defer fake.startLPARMutex.Unlock()
	fake.StartLPARStub = nil
	fake.startLPARReturns = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) StartLPARReturnsOnCall(i int, result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.startLPARMutex.Lock()
	defer fake.startLPARMutex.Unlock()
	fake.StartLPARStub = nil
	if fake.startLPARReturnsOnCall == nil {
		fake.startLPARReturnsOnCall = make(map[int]struct {
			result1 string
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.startLPARReturnsOnCall[i] = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) StopLPAR(arg1 string) (string, int, *zhmcclient.HmcError) {
	fake.stopLPARMutex.Lock()
	ret, specificReturn := fake.stopLPARReturnsOnCall[len(fake.stopLPARArgsForCall)]
	fake.stopLPARArgsForCall = append(fake.stopLPARArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.StopLPARStub
	fakeReturns := fake.stopLPARReturns
	fake.recordInvocation("StopLPAR", []interface{}{arg1})
	fake.stopLPARMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *LparAPI) StopLPARCallCount() int {
	fake.stopLPARMutex.RLock()
	defer fake.stopLPARMutex.RUnlock()
	return len(fake.stopLPARArgsForCall)
}

func (fake *LparAPI) StopLPARCalls(stub func(string) (string, int, *zhmcclient.HmcError)) {
	fake.stopLPARMutex.Lock()
	defer fake.stopLPARMutex.Unlock()
	fake.StopLPARStub = stub
}

func (fake *LparAPI) StopLPARArgsForCall(i int) string {
	fake.stopLPARMutex.RLock()
	defer fake.stopLPARMutex.RUnlock()
	argsForCall := fake.stopLPARArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LparAPI) StopLPARReturns(result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.stopLPARMutex.Lock()
	defer fake.stopLPARMutex.Unlock()
	fake.StopLPARStub = nil
	fake.stopLPARReturns = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) StopLPARReturnsOnCall(i int, result1 string, result2 int, result3 *zhmcclient.HmcError) {
	fake.stopLPARMutex.Lock()
	defer fake.stopLPARMutex.Unlock()
	fake.StopLPARStub = nil
	if fake.stopLPARReturnsOnCall == nil {
		fake.stopLPARReturnsOnCall = make(map[int]struct {
			result1 string
			result2 int
			result3 *zhmcclient.HmcError
		})
	}
	fake.stopLPARReturnsOnCall[i] = struct {
		result1 string
		result2 int
		result3 *zhmcclient.HmcError
	}{result1, result2, result3}
}

func (fake *LparAPI) UnmountIsoImage(arg1 string) (int, *zhmcclient.HmcError) {
	fake.unmountIsoImageMutex.Lock()
	ret, specificReturn := fake.unmountIsoImageReturnsOnCall[len(fake.unmountIsoImageArgsForCall)]
	fake.unmountIsoImageArgsForCall = append(fake.unmountIsoImageArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.UnmountIsoImageStub
	fakeReturns := fake.unmountIsoImageReturns
	fake.recordInvocation("UnmountIsoImage", []interface{}{arg1})
	fake.unmountIsoImageMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LparAPI) UnmountIsoImageCallCount() int {
	fake.unmountIsoImageMutex.RLock()
	defer fake.unmountIsoImageMutex.RUnlock()
	return len(fake.unmountIsoImageArgsForCall)
}

func (fake *LparAPI) UnmountIsoImageCalls(stub func(string) (int, *zhmcclient.HmcError)) {
	fake.unmountIsoImageMutex.Lock()
	defer fake.unmountIsoImageMutex.Unlock()
	fake.UnmountIsoImageStub = stub
}

func (fake *LparAPI) UnmountIsoImageArgsForCall(i int) string {
	fake.unmountIsoImageMutex.RLock()
	defer fake.unmountIsoImageMutex.RUnlock()
	argsForCall := fake.unmountIsoImageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LparAPI) UnmountIsoImageReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.unmountIsoImageMutex.Lock()
	defer fake.unmountIsoImageMutex.Unlock()
	fake.UnmountIsoImageStub = nil
	fake.unmountIsoImageReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *LparAPI) UnmountIsoImageReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.unmountIsoImageMutex.Lock()
	defer fake.unmountIsoImageMutex.Unlock()
	fake.UnmountIsoImageStub = nil
	if fake.unmountIsoImageReturnsOnCall == nil {
		fake.unmountIsoImageReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.unmountIsoImageReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *LparAPI) UpdateLparProperties(arg1 string, arg2 *zhmcclient.LparProperties) (int, *zhmcclient.HmcError) {
	fake.updateLparPropertiesMutex.Lock()
	ret, specificReturn := fake.updateLparPropertiesReturnsOnCall[len(fake.updateLparPropertiesArgsForCall)]
	fake.updateLparPropertiesArgsForCall = append(fake.updateLparPropertiesArgsForCall, struct {
		arg1 string
		arg2 *zhmcclient.LparProperties
	}{arg1, arg2})
	stub := fake.UpdateLparPropertiesStub
	fakeReturns := fake.updateLparPropertiesReturns
	fake.recordInvocation("UpdateLparProperties", []interface{}{arg1, arg2})
	fake.updateLparPropertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LparAPI) UpdateLparPropertiesCallCount() int {
	fake.updateLparPropertiesMutex.RLock()
	defer fake.updateLparPropertiesMutex.RUnlock()
	return len(fake.updateLparPropertiesArgsForCall)
}

func (fake *LparAPI) UpdateLparPropertiesCalls(stub func(string, *zhmcclient.LparProperties) (int, *zhmcclient.HmcError)) {
	fake.updateLparPropertiesMutex.Lock()
	defer fake.updateLparPropertiesMutex.Unlock()
	fake.UpdateLparPropertiesStub = stub
}

func (fake *LparAPI) UpdateLparPropertiesArgsForCall(i int) (string, *zhmcclient.LparProperties) {
	fake.updateLparPropertiesMutex.RLock()
	defer fake.updateLparPropertiesMutex.RUnlock()
	argsForCall := fake.updateLparPropertiesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *LparAPI) UpdateLparPropertiesReturns(result1 int, result2 *zhmcclient.HmcError) {
	fake.updateLparPropertiesMutex.Lock()
	defer fake.updateLparPropertiesMutex.Unlock()
	fake.UpdateLparPropertiesStub = nil
	fake.updateLparPropertiesReturns = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *LparAPI) UpdateLparPropertiesReturnsOnCall(i int, result1 int, result2 *zhmcclient.HmcError) {
	fake.updateLparPropertiesMutex.Lock()
	defer fake.updateLparPropertiesMutex.Unlock()
	fake.UpdateLparPropertiesStub = nil
	if fake.updateLparPropertiesReturnsOnCall == nil {
		fake.updateLparPropertiesReturnsOnCall = make(map[int]struct {
			result1 int
			result2 *zhmcclient.HmcError
		})
	}
	fake.updateLparPropertiesReturnsOnCall[i] = struct {
		result1 int
		result2 *zhmcclient.HmcError
	}{result1, result2}
}

func (fake *LparAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.attachStorageGroupToPartitionMutex.RLock()
	defer fake.attachStorageGroupToPartitionMutex.RUnlock()
	fake.createLPARMutex.RLock()
	defer fake.createLPARMutex.RUnlock()
	fake.deleteLPARMutex.RLock()
	defer fake.deleteLPARMutex.RUnlock()
	fake.detachStorageGroupToPartitionMutex.RLock()
	defer fake.detachStorageGroupToPartitionMutex.RUnlock()
	fake.fetchAsciiConsoleURIMutex.RLock()
	defer fake.fetchAsciiConsoleURIMutex.RUnlock()
	fake.getLparPropertiesMutex.RLock()
	defer fake.getLparPropertiesMutex.RUnlock()
	fake.listLPARsMutex.RLock()
	defer fake.listLPARsMutex.RUnlock()
	fake.listNicsMutex.RLock()
	defer fake.listNicsMutex.RUnlock()
	fake.mountIsoImageMutex.RLock()
	defer fake.mountIsoImageMutex.RUnlock()
	fake.startLPARMutex.RLock()
	defer fake.startLPARMutex.RUnlock()
	fake.stopLPARMutex.RLock()
	defer fake.stopLPARMutex.RUnlock()
	fake.unmountIsoImageMutex.RLock()
	defer fake.unmountIsoImageMutex.RUnlock()
	fake.updateLparPropertiesMutex.RLock()
	defer fake.updateLparPropertiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *LparAPI) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ zhmcclient.LparAPI = new(LparAPI)
