/*
Copyright The Kubernetes Authors.

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

// Code generated by counterfeiter. DO NOT EDIT.
package spdxfakes

import (
	"sync"

	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"k8s.io/release/pkg/license"
	"k8s.io/release/pkg/spdx"
)

type FakeSpdxImplementation struct {
	AnalyzeImageLayerStub        func(string, *spdx.Package) error
	analyzeImageLayerMutex       sync.RWMutex
	analyzeImageLayerArgsForCall []struct {
		arg1 string
		arg2 *spdx.Package
	}
	analyzeImageLayerReturns struct {
		result1 error
	}
	analyzeImageLayerReturnsOnCall map[int]struct {
		result1 error
	}
	ApplyIgnorePatternsStub        func([]string, []gitignore.Pattern) []string
	applyIgnorePatternsMutex       sync.RWMutex
	applyIgnorePatternsArgsForCall []struct {
		arg1 []string
		arg2 []gitignore.Pattern
	}
	applyIgnorePatternsReturns struct {
		result1 []string
	}
	applyIgnorePatternsReturnsOnCall map[int]struct {
		result1 []string
	}
	ExtractTarballTmpStub        func(string) (string, error)
	extractTarballTmpMutex       sync.RWMutex
	extractTarballTmpArgsForCall []struct {
		arg1 string
	}
	extractTarballTmpReturns struct {
		result1 string
		result2 error
	}
	extractTarballTmpReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetDirectoryLicenseStub        func(*license.Reader, string, *spdx.Options) (*license.License, error)
	getDirectoryLicenseMutex       sync.RWMutex
	getDirectoryLicenseArgsForCall []struct {
		arg1 *license.Reader
		arg2 string
		arg3 *spdx.Options
	}
	getDirectoryLicenseReturns struct {
		result1 *license.License
		result2 error
	}
	getDirectoryLicenseReturnsOnCall map[int]struct {
		result1 *license.License
		result2 error
	}
	GetDirectoryTreeStub        func(string) ([]string, error)
	getDirectoryTreeMutex       sync.RWMutex
	getDirectoryTreeArgsForCall []struct {
		arg1 string
	}
	getDirectoryTreeReturns struct {
		result1 []string
		result2 error
	}
	getDirectoryTreeReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetGoDependenciesStub        func(string, *spdx.Options) ([]*spdx.Package, error)
	getGoDependenciesMutex       sync.RWMutex
	getGoDependenciesArgsForCall []struct {
		arg1 string
		arg2 *spdx.Options
	}
	getGoDependenciesReturns struct {
		result1 []*spdx.Package
		result2 error
	}
	getGoDependenciesReturnsOnCall map[int]struct {
		result1 []*spdx.Package
		result2 error
	}
	IgnorePatternsStub        func(string, []string, bool) ([]gitignore.Pattern, error)
	ignorePatternsMutex       sync.RWMutex
	ignorePatternsArgsForCall []struct {
		arg1 string
		arg2 []string
		arg3 bool
	}
	ignorePatternsReturns struct {
		result1 []gitignore.Pattern
		result2 error
	}
	ignorePatternsReturnsOnCall map[int]struct {
		result1 []gitignore.Pattern
		result2 error
	}
	ImageRefToPackageStub        func(string, *spdx.Options) (*spdx.Package, error)
	imageRefToPackageMutex       sync.RWMutex
	imageRefToPackageArgsForCall []struct {
		arg1 string
		arg2 *spdx.Options
	}
	imageRefToPackageReturns struct {
		result1 *spdx.Package
		result2 error
	}
	imageRefToPackageReturnsOnCall map[int]struct {
		result1 *spdx.Package
		result2 error
	}
	LicenseReaderStub        func(*spdx.Options) (*license.Reader, error)
	licenseReaderMutex       sync.RWMutex
	licenseReaderArgsForCall []struct {
		arg1 *spdx.Options
	}
	licenseReaderReturns struct {
		result1 *license.Reader
		result2 error
	}
	licenseReaderReturnsOnCall map[int]struct {
		result1 *license.Reader
		result2 error
	}
	PackageFromImageTarballStub        func(string, *spdx.Options) (*spdx.Package, error)
	packageFromImageTarballMutex       sync.RWMutex
	packageFromImageTarballArgsForCall []struct {
		arg1 string
		arg2 *spdx.Options
	}
	packageFromImageTarballReturns struct {
		result1 *spdx.Package
		result2 error
	}
	packageFromImageTarballReturnsOnCall map[int]struct {
		result1 *spdx.Package
		result2 error
	}
	PackageFromLayerTarballStub        func(string, *spdx.TarballOptions) (*spdx.Package, error)
	packageFromLayerTarballMutex       sync.RWMutex
	packageFromLayerTarballArgsForCall []struct {
		arg1 string
		arg2 *spdx.TarballOptions
	}
	packageFromLayerTarballReturns struct {
		result1 *spdx.Package
		result2 error
	}
	packageFromLayerTarballReturnsOnCall map[int]struct {
		result1 *spdx.Package
		result2 error
	}
	PullImagesToArchiveStub func(string, string) ([]struct {
		Reference string
		Archive   string
		Arch      string
		OS        string
	}, error)
	pullImagesToArchiveMutex       sync.RWMutex
	pullImagesToArchiveArgsForCall []struct {
		arg1 string
		arg2 string
	}
	pullImagesToArchiveReturns struct {
		result1 []struct {
			Reference string
			Archive   string
			Arch      string
			OS        string
		}
		result2 error
	}
	pullImagesToArchiveReturnsOnCall map[int]struct {
		result1 []struct {
			Reference string
			Archive   string
			Arch      string
			OS        string
		}
		result2 error
	}
	ReadArchiveManifestStub        func(string) (*spdx.ArchiveManifest, error)
	readArchiveManifestMutex       sync.RWMutex
	readArchiveManifestArgsForCall []struct {
		arg1 string
	}
	readArchiveManifestReturns struct {
		result1 *spdx.ArchiveManifest
		result2 error
	}
	readArchiveManifestReturnsOnCall map[int]struct {
		result1 *spdx.ArchiveManifest
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpdxImplementation) AnalyzeImageLayer(arg1 string, arg2 *spdx.Package) error {
	fake.analyzeImageLayerMutex.Lock()
	ret, specificReturn := fake.analyzeImageLayerReturnsOnCall[len(fake.analyzeImageLayerArgsForCall)]
	fake.analyzeImageLayerArgsForCall = append(fake.analyzeImageLayerArgsForCall, struct {
		arg1 string
		arg2 *spdx.Package
	}{arg1, arg2})
	stub := fake.AnalyzeImageLayerStub
	fakeReturns := fake.analyzeImageLayerReturns
	fake.recordInvocation("AnalyzeImageLayer", []interface{}{arg1, arg2})
	fake.analyzeImageLayerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSpdxImplementation) AnalyzeImageLayerCallCount() int {
	fake.analyzeImageLayerMutex.RLock()
	defer fake.analyzeImageLayerMutex.RUnlock()
	return len(fake.analyzeImageLayerArgsForCall)
}

func (fake *FakeSpdxImplementation) AnalyzeImageLayerCalls(stub func(string, *spdx.Package) error) {
	fake.analyzeImageLayerMutex.Lock()
	defer fake.analyzeImageLayerMutex.Unlock()
	fake.AnalyzeImageLayerStub = stub
}

func (fake *FakeSpdxImplementation) AnalyzeImageLayerArgsForCall(i int) (string, *spdx.Package) {
	fake.analyzeImageLayerMutex.RLock()
	defer fake.analyzeImageLayerMutex.RUnlock()
	argsForCall := fake.analyzeImageLayerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpdxImplementation) AnalyzeImageLayerReturns(result1 error) {
	fake.analyzeImageLayerMutex.Lock()
	defer fake.analyzeImageLayerMutex.Unlock()
	fake.AnalyzeImageLayerStub = nil
	fake.analyzeImageLayerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpdxImplementation) AnalyzeImageLayerReturnsOnCall(i int, result1 error) {
	fake.analyzeImageLayerMutex.Lock()
	defer fake.analyzeImageLayerMutex.Unlock()
	fake.AnalyzeImageLayerStub = nil
	if fake.analyzeImageLayerReturnsOnCall == nil {
		fake.analyzeImageLayerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.analyzeImageLayerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpdxImplementation) ApplyIgnorePatterns(arg1 []string, arg2 []gitignore.Pattern) []string {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []gitignore.Pattern
	if arg2 != nil {
		arg2Copy = make([]gitignore.Pattern, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.applyIgnorePatternsMutex.Lock()
	ret, specificReturn := fake.applyIgnorePatternsReturnsOnCall[len(fake.applyIgnorePatternsArgsForCall)]
	fake.applyIgnorePatternsArgsForCall = append(fake.applyIgnorePatternsArgsForCall, struct {
		arg1 []string
		arg2 []gitignore.Pattern
	}{arg1Copy, arg2Copy})
	stub := fake.ApplyIgnorePatternsStub
	fakeReturns := fake.applyIgnorePatternsReturns
	fake.recordInvocation("ApplyIgnorePatterns", []interface{}{arg1Copy, arg2Copy})
	fake.applyIgnorePatternsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSpdxImplementation) ApplyIgnorePatternsCallCount() int {
	fake.applyIgnorePatternsMutex.RLock()
	defer fake.applyIgnorePatternsMutex.RUnlock()
	return len(fake.applyIgnorePatternsArgsForCall)
}

func (fake *FakeSpdxImplementation) ApplyIgnorePatternsCalls(stub func([]string, []gitignore.Pattern) []string) {
	fake.applyIgnorePatternsMutex.Lock()
	defer fake.applyIgnorePatternsMutex.Unlock()
	fake.ApplyIgnorePatternsStub = stub
}

func (fake *FakeSpdxImplementation) ApplyIgnorePatternsArgsForCall(i int) ([]string, []gitignore.Pattern) {
	fake.applyIgnorePatternsMutex.RLock()
	defer fake.applyIgnorePatternsMutex.RUnlock()
	argsForCall := fake.applyIgnorePatternsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpdxImplementation) ApplyIgnorePatternsReturns(result1 []string) {
	fake.applyIgnorePatternsMutex.Lock()
	defer fake.applyIgnorePatternsMutex.Unlock()
	fake.ApplyIgnorePatternsStub = nil
	fake.applyIgnorePatternsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeSpdxImplementation) ApplyIgnorePatternsReturnsOnCall(i int, result1 []string) {
	fake.applyIgnorePatternsMutex.Lock()
	defer fake.applyIgnorePatternsMutex.Unlock()
	fake.ApplyIgnorePatternsStub = nil
	if fake.applyIgnorePatternsReturnsOnCall == nil {
		fake.applyIgnorePatternsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.applyIgnorePatternsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeSpdxImplementation) ExtractTarballTmp(arg1 string) (string, error) {
	fake.extractTarballTmpMutex.Lock()
	ret, specificReturn := fake.extractTarballTmpReturnsOnCall[len(fake.extractTarballTmpArgsForCall)]
	fake.extractTarballTmpArgsForCall = append(fake.extractTarballTmpArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ExtractTarballTmpStub
	fakeReturns := fake.extractTarballTmpReturns
	fake.recordInvocation("ExtractTarballTmp", []interface{}{arg1})
	fake.extractTarballTmpMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSpdxImplementation) ExtractTarballTmpCallCount() int {
	fake.extractTarballTmpMutex.RLock()
	defer fake.extractTarballTmpMutex.RUnlock()
	return len(fake.extractTarballTmpArgsForCall)
}

func (fake *FakeSpdxImplementation) ExtractTarballTmpCalls(stub func(string) (string, error)) {
	fake.extractTarballTmpMutex.Lock()
	defer fake.extractTarballTmpMutex.Unlock()
	fake.ExtractTarballTmpStub = stub
}

func (fake *FakeSpdxImplementation) ExtractTarballTmpArgsForCall(i int) string {
	fake.extractTarballTmpMutex.RLock()
	defer fake.extractTarballTmpMutex.RUnlock()
	argsForCall := fake.extractTarballTmpArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSpdxImplementation) ExtractTarballTmpReturns(result1 string, result2 error) {
	fake.extractTarballTmpMutex.Lock()
	defer fake.extractTarballTmpMutex.Unlock()
	fake.ExtractTarballTmpStub = nil
	fake.extractTarballTmpReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) ExtractTarballTmpReturnsOnCall(i int, result1 string, result2 error) {
	fake.extractTarballTmpMutex.Lock()
	defer fake.extractTarballTmpMutex.Unlock()
	fake.ExtractTarballTmpStub = nil
	if fake.extractTarballTmpReturnsOnCall == nil {
		fake.extractTarballTmpReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.extractTarballTmpReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) GetDirectoryLicense(arg1 *license.Reader, arg2 string, arg3 *spdx.Options) (*license.License, error) {
	fake.getDirectoryLicenseMutex.Lock()
	ret, specificReturn := fake.getDirectoryLicenseReturnsOnCall[len(fake.getDirectoryLicenseArgsForCall)]
	fake.getDirectoryLicenseArgsForCall = append(fake.getDirectoryLicenseArgsForCall, struct {
		arg1 *license.Reader
		arg2 string
		arg3 *spdx.Options
	}{arg1, arg2, arg3})
	stub := fake.GetDirectoryLicenseStub
	fakeReturns := fake.getDirectoryLicenseReturns
	fake.recordInvocation("GetDirectoryLicense", []interface{}{arg1, arg2, arg3})
	fake.getDirectoryLicenseMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSpdxImplementation) GetDirectoryLicenseCallCount() int {
	fake.getDirectoryLicenseMutex.RLock()
	defer fake.getDirectoryLicenseMutex.RUnlock()
	return len(fake.getDirectoryLicenseArgsForCall)
}

func (fake *FakeSpdxImplementation) GetDirectoryLicenseCalls(stub func(*license.Reader, string, *spdx.Options) (*license.License, error)) {
	fake.getDirectoryLicenseMutex.Lock()
	defer fake.getDirectoryLicenseMutex.Unlock()
	fake.GetDirectoryLicenseStub = stub
}

func (fake *FakeSpdxImplementation) GetDirectoryLicenseArgsForCall(i int) (*license.Reader, string, *spdx.Options) {
	fake.getDirectoryLicenseMutex.RLock()
	defer fake.getDirectoryLicenseMutex.RUnlock()
	argsForCall := fake.getDirectoryLicenseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSpdxImplementation) GetDirectoryLicenseReturns(result1 *license.License, result2 error) {
	fake.getDirectoryLicenseMutex.Lock()
	defer fake.getDirectoryLicenseMutex.Unlock()
	fake.GetDirectoryLicenseStub = nil
	fake.getDirectoryLicenseReturns = struct {
		result1 *license.License
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) GetDirectoryLicenseReturnsOnCall(i int, result1 *license.License, result2 error) {
	fake.getDirectoryLicenseMutex.Lock()
	defer fake.getDirectoryLicenseMutex.Unlock()
	fake.GetDirectoryLicenseStub = nil
	if fake.getDirectoryLicenseReturnsOnCall == nil {
		fake.getDirectoryLicenseReturnsOnCall = make(map[int]struct {
			result1 *license.License
			result2 error
		})
	}
	fake.getDirectoryLicenseReturnsOnCall[i] = struct {
		result1 *license.License
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) GetDirectoryTree(arg1 string) ([]string, error) {
	fake.getDirectoryTreeMutex.Lock()
	ret, specificReturn := fake.getDirectoryTreeReturnsOnCall[len(fake.getDirectoryTreeArgsForCall)]
	fake.getDirectoryTreeArgsForCall = append(fake.getDirectoryTreeArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetDirectoryTreeStub
	fakeReturns := fake.getDirectoryTreeReturns
	fake.recordInvocation("GetDirectoryTree", []interface{}{arg1})
	fake.getDirectoryTreeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSpdxImplementation) GetDirectoryTreeCallCount() int {
	fake.getDirectoryTreeMutex.RLock()
	defer fake.getDirectoryTreeMutex.RUnlock()
	return len(fake.getDirectoryTreeArgsForCall)
}

func (fake *FakeSpdxImplementation) GetDirectoryTreeCalls(stub func(string) ([]string, error)) {
	fake.getDirectoryTreeMutex.Lock()
	defer fake.getDirectoryTreeMutex.Unlock()
	fake.GetDirectoryTreeStub = stub
}

func (fake *FakeSpdxImplementation) GetDirectoryTreeArgsForCall(i int) string {
	fake.getDirectoryTreeMutex.RLock()
	defer fake.getDirectoryTreeMutex.RUnlock()
	argsForCall := fake.getDirectoryTreeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSpdxImplementation) GetDirectoryTreeReturns(result1 []string, result2 error) {
	fake.getDirectoryTreeMutex.Lock()
	defer fake.getDirectoryTreeMutex.Unlock()
	fake.GetDirectoryTreeStub = nil
	fake.getDirectoryTreeReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) GetDirectoryTreeReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getDirectoryTreeMutex.Lock()
	defer fake.getDirectoryTreeMutex.Unlock()
	fake.GetDirectoryTreeStub = nil
	if fake.getDirectoryTreeReturnsOnCall == nil {
		fake.getDirectoryTreeReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getDirectoryTreeReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) GetGoDependencies(arg1 string, arg2 *spdx.Options) ([]*spdx.Package, error) {
	fake.getGoDependenciesMutex.Lock()
	ret, specificReturn := fake.getGoDependenciesReturnsOnCall[len(fake.getGoDependenciesArgsForCall)]
	fake.getGoDependenciesArgsForCall = append(fake.getGoDependenciesArgsForCall, struct {
		arg1 string
		arg2 *spdx.Options
	}{arg1, arg2})
	stub := fake.GetGoDependenciesStub
	fakeReturns := fake.getGoDependenciesReturns
	fake.recordInvocation("GetGoDependencies", []interface{}{arg1, arg2})
	fake.getGoDependenciesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSpdxImplementation) GetGoDependenciesCallCount() int {
	fake.getGoDependenciesMutex.RLock()
	defer fake.getGoDependenciesMutex.RUnlock()
	return len(fake.getGoDependenciesArgsForCall)
}

func (fake *FakeSpdxImplementation) GetGoDependenciesCalls(stub func(string, *spdx.Options) ([]*spdx.Package, error)) {
	fake.getGoDependenciesMutex.Lock()
	defer fake.getGoDependenciesMutex.Unlock()
	fake.GetGoDependenciesStub = stub
}

func (fake *FakeSpdxImplementation) GetGoDependenciesArgsForCall(i int) (string, *spdx.Options) {
	fake.getGoDependenciesMutex.RLock()
	defer fake.getGoDependenciesMutex.RUnlock()
	argsForCall := fake.getGoDependenciesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpdxImplementation) GetGoDependenciesReturns(result1 []*spdx.Package, result2 error) {
	fake.getGoDependenciesMutex.Lock()
	defer fake.getGoDependenciesMutex.Unlock()
	fake.GetGoDependenciesStub = nil
	fake.getGoDependenciesReturns = struct {
		result1 []*spdx.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) GetGoDependenciesReturnsOnCall(i int, result1 []*spdx.Package, result2 error) {
	fake.getGoDependenciesMutex.Lock()
	defer fake.getGoDependenciesMutex.Unlock()
	fake.GetGoDependenciesStub = nil
	if fake.getGoDependenciesReturnsOnCall == nil {
		fake.getGoDependenciesReturnsOnCall = make(map[int]struct {
			result1 []*spdx.Package
			result2 error
		})
	}
	fake.getGoDependenciesReturnsOnCall[i] = struct {
		result1 []*spdx.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) IgnorePatterns(arg1 string, arg2 []string, arg3 bool) ([]gitignore.Pattern, error) {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.ignorePatternsMutex.Lock()
	ret, specificReturn := fake.ignorePatternsReturnsOnCall[len(fake.ignorePatternsArgsForCall)]
	fake.ignorePatternsArgsForCall = append(fake.ignorePatternsArgsForCall, struct {
		arg1 string
		arg2 []string
		arg3 bool
	}{arg1, arg2Copy, arg3})
	stub := fake.IgnorePatternsStub
	fakeReturns := fake.ignorePatternsReturns
	fake.recordInvocation("IgnorePatterns", []interface{}{arg1, arg2Copy, arg3})
	fake.ignorePatternsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSpdxImplementation) IgnorePatternsCallCount() int {
	fake.ignorePatternsMutex.RLock()
	defer fake.ignorePatternsMutex.RUnlock()
	return len(fake.ignorePatternsArgsForCall)
}

func (fake *FakeSpdxImplementation) IgnorePatternsCalls(stub func(string, []string, bool) ([]gitignore.Pattern, error)) {
	fake.ignorePatternsMutex.Lock()
	defer fake.ignorePatternsMutex.Unlock()
	fake.IgnorePatternsStub = stub
}

func (fake *FakeSpdxImplementation) IgnorePatternsArgsForCall(i int) (string, []string, bool) {
	fake.ignorePatternsMutex.RLock()
	defer fake.ignorePatternsMutex.RUnlock()
	argsForCall := fake.ignorePatternsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSpdxImplementation) IgnorePatternsReturns(result1 []gitignore.Pattern, result2 error) {
	fake.ignorePatternsMutex.Lock()
	defer fake.ignorePatternsMutex.Unlock()
	fake.IgnorePatternsStub = nil
	fake.ignorePatternsReturns = struct {
		result1 []gitignore.Pattern
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) IgnorePatternsReturnsOnCall(i int, result1 []gitignore.Pattern, result2 error) {
	fake.ignorePatternsMutex.Lock()
	defer fake.ignorePatternsMutex.Unlock()
	fake.IgnorePatternsStub = nil
	if fake.ignorePatternsReturnsOnCall == nil {
		fake.ignorePatternsReturnsOnCall = make(map[int]struct {
			result1 []gitignore.Pattern
			result2 error
		})
	}
	fake.ignorePatternsReturnsOnCall[i] = struct {
		result1 []gitignore.Pattern
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) ImageRefToPackage(arg1 string, arg2 *spdx.Options) (*spdx.Package, error) {
	fake.imageRefToPackageMutex.Lock()
	ret, specificReturn := fake.imageRefToPackageReturnsOnCall[len(fake.imageRefToPackageArgsForCall)]
	fake.imageRefToPackageArgsForCall = append(fake.imageRefToPackageArgsForCall, struct {
		arg1 string
		arg2 *spdx.Options
	}{arg1, arg2})
	stub := fake.ImageRefToPackageStub
	fakeReturns := fake.imageRefToPackageReturns
	fake.recordInvocation("ImageRefToPackage", []interface{}{arg1, arg2})
	fake.imageRefToPackageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSpdxImplementation) ImageRefToPackageCallCount() int {
	fake.imageRefToPackageMutex.RLock()
	defer fake.imageRefToPackageMutex.RUnlock()
	return len(fake.imageRefToPackageArgsForCall)
}

func (fake *FakeSpdxImplementation) ImageRefToPackageCalls(stub func(string, *spdx.Options) (*spdx.Package, error)) {
	fake.imageRefToPackageMutex.Lock()
	defer fake.imageRefToPackageMutex.Unlock()
	fake.ImageRefToPackageStub = stub
}

func (fake *FakeSpdxImplementation) ImageRefToPackageArgsForCall(i int) (string, *spdx.Options) {
	fake.imageRefToPackageMutex.RLock()
	defer fake.imageRefToPackageMutex.RUnlock()
	argsForCall := fake.imageRefToPackageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpdxImplementation) ImageRefToPackageReturns(result1 *spdx.Package, result2 error) {
	fake.imageRefToPackageMutex.Lock()
	defer fake.imageRefToPackageMutex.Unlock()
	fake.ImageRefToPackageStub = nil
	fake.imageRefToPackageReturns = struct {
		result1 *spdx.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) ImageRefToPackageReturnsOnCall(i int, result1 *spdx.Package, result2 error) {
	fake.imageRefToPackageMutex.Lock()
	defer fake.imageRefToPackageMutex.Unlock()
	fake.ImageRefToPackageStub = nil
	if fake.imageRefToPackageReturnsOnCall == nil {
		fake.imageRefToPackageReturnsOnCall = make(map[int]struct {
			result1 *spdx.Package
			result2 error
		})
	}
	fake.imageRefToPackageReturnsOnCall[i] = struct {
		result1 *spdx.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) LicenseReader(arg1 *spdx.Options) (*license.Reader, error) {
	fake.licenseReaderMutex.Lock()
	ret, specificReturn := fake.licenseReaderReturnsOnCall[len(fake.licenseReaderArgsForCall)]
	fake.licenseReaderArgsForCall = append(fake.licenseReaderArgsForCall, struct {
		arg1 *spdx.Options
	}{arg1})
	stub := fake.LicenseReaderStub
	fakeReturns := fake.licenseReaderReturns
	fake.recordInvocation("LicenseReader", []interface{}{arg1})
	fake.licenseReaderMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSpdxImplementation) LicenseReaderCallCount() int {
	fake.licenseReaderMutex.RLock()
	defer fake.licenseReaderMutex.RUnlock()
	return len(fake.licenseReaderArgsForCall)
}

func (fake *FakeSpdxImplementation) LicenseReaderCalls(stub func(*spdx.Options) (*license.Reader, error)) {
	fake.licenseReaderMutex.Lock()
	defer fake.licenseReaderMutex.Unlock()
	fake.LicenseReaderStub = stub
}

func (fake *FakeSpdxImplementation) LicenseReaderArgsForCall(i int) *spdx.Options {
	fake.licenseReaderMutex.RLock()
	defer fake.licenseReaderMutex.RUnlock()
	argsForCall := fake.licenseReaderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSpdxImplementation) LicenseReaderReturns(result1 *license.Reader, result2 error) {
	fake.licenseReaderMutex.Lock()
	defer fake.licenseReaderMutex.Unlock()
	fake.LicenseReaderStub = nil
	fake.licenseReaderReturns = struct {
		result1 *license.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) LicenseReaderReturnsOnCall(i int, result1 *license.Reader, result2 error) {
	fake.licenseReaderMutex.Lock()
	defer fake.licenseReaderMutex.Unlock()
	fake.LicenseReaderStub = nil
	if fake.licenseReaderReturnsOnCall == nil {
		fake.licenseReaderReturnsOnCall = make(map[int]struct {
			result1 *license.Reader
			result2 error
		})
	}
	fake.licenseReaderReturnsOnCall[i] = struct {
		result1 *license.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) PackageFromImageTarball(arg1 string, arg2 *spdx.Options) (*spdx.Package, error) {
	fake.packageFromImageTarballMutex.Lock()
	ret, specificReturn := fake.packageFromImageTarballReturnsOnCall[len(fake.packageFromImageTarballArgsForCall)]
	fake.packageFromImageTarballArgsForCall = append(fake.packageFromImageTarballArgsForCall, struct {
		arg1 string
		arg2 *spdx.Options
	}{arg1, arg2})
	stub := fake.PackageFromImageTarballStub
	fakeReturns := fake.packageFromImageTarballReturns
	fake.recordInvocation("PackageFromImageTarball", []interface{}{arg1, arg2})
	fake.packageFromImageTarballMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSpdxImplementation) PackageFromImageTarballCallCount() int {
	fake.packageFromImageTarballMutex.RLock()
	defer fake.packageFromImageTarballMutex.RUnlock()
	return len(fake.packageFromImageTarballArgsForCall)
}

func (fake *FakeSpdxImplementation) PackageFromImageTarballCalls(stub func(string, *spdx.Options) (*spdx.Package, error)) {
	fake.packageFromImageTarballMutex.Lock()
	defer fake.packageFromImageTarballMutex.Unlock()
	fake.PackageFromImageTarballStub = stub
}

func (fake *FakeSpdxImplementation) PackageFromImageTarballArgsForCall(i int) (string, *spdx.Options) {
	fake.packageFromImageTarballMutex.RLock()
	defer fake.packageFromImageTarballMutex.RUnlock()
	argsForCall := fake.packageFromImageTarballArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpdxImplementation) PackageFromImageTarballReturns(result1 *spdx.Package, result2 error) {
	fake.packageFromImageTarballMutex.Lock()
	defer fake.packageFromImageTarballMutex.Unlock()
	fake.PackageFromImageTarballStub = nil
	fake.packageFromImageTarballReturns = struct {
		result1 *spdx.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) PackageFromImageTarballReturnsOnCall(i int, result1 *spdx.Package, result2 error) {
	fake.packageFromImageTarballMutex.Lock()
	defer fake.packageFromImageTarballMutex.Unlock()
	fake.PackageFromImageTarballStub = nil
	if fake.packageFromImageTarballReturnsOnCall == nil {
		fake.packageFromImageTarballReturnsOnCall = make(map[int]struct {
			result1 *spdx.Package
			result2 error
		})
	}
	fake.packageFromImageTarballReturnsOnCall[i] = struct {
		result1 *spdx.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) PackageFromLayerTarball(arg1 string, arg2 *spdx.TarballOptions) (*spdx.Package, error) {
	fake.packageFromLayerTarballMutex.Lock()
	ret, specificReturn := fake.packageFromLayerTarballReturnsOnCall[len(fake.packageFromLayerTarballArgsForCall)]
	fake.packageFromLayerTarballArgsForCall = append(fake.packageFromLayerTarballArgsForCall, struct {
		arg1 string
		arg2 *spdx.TarballOptions
	}{arg1, arg2})
	stub := fake.PackageFromLayerTarballStub
	fakeReturns := fake.packageFromLayerTarballReturns
	fake.recordInvocation("PackageFromLayerTarball", []interface{}{arg1, arg2})
	fake.packageFromLayerTarballMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSpdxImplementation) PackageFromLayerTarballCallCount() int {
	fake.packageFromLayerTarballMutex.RLock()
	defer fake.packageFromLayerTarballMutex.RUnlock()
	return len(fake.packageFromLayerTarballArgsForCall)
}

func (fake *FakeSpdxImplementation) PackageFromLayerTarballCalls(stub func(string, *spdx.TarballOptions) (*spdx.Package, error)) {
	fake.packageFromLayerTarballMutex.Lock()
	defer fake.packageFromLayerTarballMutex.Unlock()
	fake.PackageFromLayerTarballStub = stub
}

func (fake *FakeSpdxImplementation) PackageFromLayerTarballArgsForCall(i int) (string, *spdx.TarballOptions) {
	fake.packageFromLayerTarballMutex.RLock()
	defer fake.packageFromLayerTarballMutex.RUnlock()
	argsForCall := fake.packageFromLayerTarballArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpdxImplementation) PackageFromLayerTarballReturns(result1 *spdx.Package, result2 error) {
	fake.packageFromLayerTarballMutex.Lock()
	defer fake.packageFromLayerTarballMutex.Unlock()
	fake.PackageFromLayerTarballStub = nil
	fake.packageFromLayerTarballReturns = struct {
		result1 *spdx.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) PackageFromLayerTarballReturnsOnCall(i int, result1 *spdx.Package, result2 error) {
	fake.packageFromLayerTarballMutex.Lock()
	defer fake.packageFromLayerTarballMutex.Unlock()
	fake.PackageFromLayerTarballStub = nil
	if fake.packageFromLayerTarballReturnsOnCall == nil {
		fake.packageFromLayerTarballReturnsOnCall = make(map[int]struct {
			result1 *spdx.Package
			result2 error
		})
	}
	fake.packageFromLayerTarballReturnsOnCall[i] = struct {
		result1 *spdx.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) PullImagesToArchive(arg1 string, arg2 string) ([]struct {
	Reference string
	Archive   string
	Arch      string
	OS        string
}, error) {
	fake.pullImagesToArchiveMutex.Lock()
	ret, specificReturn := fake.pullImagesToArchiveReturnsOnCall[len(fake.pullImagesToArchiveArgsForCall)]
	fake.pullImagesToArchiveArgsForCall = append(fake.pullImagesToArchiveArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.PullImagesToArchiveStub
	fakeReturns := fake.pullImagesToArchiveReturns
	fake.recordInvocation("PullImagesToArchive", []interface{}{arg1, arg2})
	fake.pullImagesToArchiveMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSpdxImplementation) PullImagesToArchiveCallCount() int {
	fake.pullImagesToArchiveMutex.RLock()
	defer fake.pullImagesToArchiveMutex.RUnlock()
	return len(fake.pullImagesToArchiveArgsForCall)
}

func (fake *FakeSpdxImplementation) PullImagesToArchiveCalls(stub func(string, string) ([]struct {
	Reference string
	Archive   string
	Arch      string
	OS        string
}, error)) {
	fake.pullImagesToArchiveMutex.Lock()
	defer fake.pullImagesToArchiveMutex.Unlock()
	fake.PullImagesToArchiveStub = stub
}

func (fake *FakeSpdxImplementation) PullImagesToArchiveArgsForCall(i int) (string, string) {
	fake.pullImagesToArchiveMutex.RLock()
	defer fake.pullImagesToArchiveMutex.RUnlock()
	argsForCall := fake.pullImagesToArchiveArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpdxImplementation) PullImagesToArchiveReturns(result1 []struct {
	Reference string
	Archive   string
	Arch      string
	OS        string
}, result2 error) {
	fake.pullImagesToArchiveMutex.Lock()
	defer fake.pullImagesToArchiveMutex.Unlock()
	fake.PullImagesToArchiveStub = nil
	fake.pullImagesToArchiveReturns = struct {
		result1 []struct {
			Reference string
			Archive   string
			Arch      string
			OS        string
		}
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) PullImagesToArchiveReturnsOnCall(i int, result1 []struct {
	Reference string
	Archive   string
	Arch      string
	OS        string
}, result2 error) {
	fake.pullImagesToArchiveMutex.Lock()
	defer fake.pullImagesToArchiveMutex.Unlock()
	fake.PullImagesToArchiveStub = nil
	if fake.pullImagesToArchiveReturnsOnCall == nil {
		fake.pullImagesToArchiveReturnsOnCall = make(map[int]struct {
			result1 []struct {
				Reference string
				Archive   string
				Arch      string
				OS        string
			}
			result2 error
		})
	}
	fake.pullImagesToArchiveReturnsOnCall[i] = struct {
		result1 []struct {
			Reference string
			Archive   string
			Arch      string
			OS        string
		}
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) ReadArchiveManifest(arg1 string) (*spdx.ArchiveManifest, error) {
	fake.readArchiveManifestMutex.Lock()
	ret, specificReturn := fake.readArchiveManifestReturnsOnCall[len(fake.readArchiveManifestArgsForCall)]
	fake.readArchiveManifestArgsForCall = append(fake.readArchiveManifestArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ReadArchiveManifestStub
	fakeReturns := fake.readArchiveManifestReturns
	fake.recordInvocation("ReadArchiveManifest", []interface{}{arg1})
	fake.readArchiveManifestMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSpdxImplementation) ReadArchiveManifestCallCount() int {
	fake.readArchiveManifestMutex.RLock()
	defer fake.readArchiveManifestMutex.RUnlock()
	return len(fake.readArchiveManifestArgsForCall)
}

func (fake *FakeSpdxImplementation) ReadArchiveManifestCalls(stub func(string) (*spdx.ArchiveManifest, error)) {
	fake.readArchiveManifestMutex.Lock()
	defer fake.readArchiveManifestMutex.Unlock()
	fake.ReadArchiveManifestStub = stub
}

func (fake *FakeSpdxImplementation) ReadArchiveManifestArgsForCall(i int) string {
	fake.readArchiveManifestMutex.RLock()
	defer fake.readArchiveManifestMutex.RUnlock()
	argsForCall := fake.readArchiveManifestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSpdxImplementation) ReadArchiveManifestReturns(result1 *spdx.ArchiveManifest, result2 error) {
	fake.readArchiveManifestMutex.Lock()
	defer fake.readArchiveManifestMutex.Unlock()
	fake.ReadArchiveManifestStub = nil
	fake.readArchiveManifestReturns = struct {
		result1 *spdx.ArchiveManifest
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) ReadArchiveManifestReturnsOnCall(i int, result1 *spdx.ArchiveManifest, result2 error) {
	fake.readArchiveManifestMutex.Lock()
	defer fake.readArchiveManifestMutex.Unlock()
	fake.ReadArchiveManifestStub = nil
	if fake.readArchiveManifestReturnsOnCall == nil {
		fake.readArchiveManifestReturnsOnCall = make(map[int]struct {
			result1 *spdx.ArchiveManifest
			result2 error
		})
	}
	fake.readArchiveManifestReturnsOnCall[i] = struct {
		result1 *spdx.ArchiveManifest
		result2 error
	}{result1, result2}
}

func (fake *FakeSpdxImplementation) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.analyzeImageLayerMutex.RLock()
	defer fake.analyzeImageLayerMutex.RUnlock()
	fake.applyIgnorePatternsMutex.RLock()
	defer fake.applyIgnorePatternsMutex.RUnlock()
	fake.extractTarballTmpMutex.RLock()
	defer fake.extractTarballTmpMutex.RUnlock()
	fake.getDirectoryLicenseMutex.RLock()
	defer fake.getDirectoryLicenseMutex.RUnlock()
	fake.getDirectoryTreeMutex.RLock()
	defer fake.getDirectoryTreeMutex.RUnlock()
	fake.getGoDependenciesMutex.RLock()
	defer fake.getGoDependenciesMutex.RUnlock()
	fake.ignorePatternsMutex.RLock()
	defer fake.ignorePatternsMutex.RUnlock()
	fake.imageRefToPackageMutex.RLock()
	defer fake.imageRefToPackageMutex.RUnlock()
	fake.licenseReaderMutex.RLock()
	defer fake.licenseReaderMutex.RUnlock()
	fake.packageFromImageTarballMutex.RLock()
	defer fake.packageFromImageTarballMutex.RUnlock()
	fake.packageFromLayerTarballMutex.RLock()
	defer fake.packageFromLayerTarballMutex.RUnlock()
	fake.pullImagesToArchiveMutex.RLock()
	defer fake.pullImagesToArchiveMutex.RUnlock()
	fake.readArchiveManifestMutex.RLock()
	defer fake.readArchiveManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSpdxImplementation) recordInvocation(key string, args []interface{}) {
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
