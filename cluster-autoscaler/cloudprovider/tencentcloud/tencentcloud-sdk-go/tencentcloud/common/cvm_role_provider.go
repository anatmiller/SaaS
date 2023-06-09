/*
Copyright 2016 The Kubernetes Authors.

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

package common

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	tcerr "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

const (
	metaUrl = "http://metadata.tencentyun.com/latest/meta-data/"
	roleUrl = metaUrl + "cam/security-credentials/"
)

var roleNotBound = errors.New("get cvm role name failed, Please confirm whether the role is bound")

type CvmRoleProvider struct {
	roleName string
}

type roleRsp struct {
	TmpSecretId  string    `json:"TmpSecretId"`
	TmpSecretKey string    `json:"TmpSecretKey"`
	ExpiredTime  int64     `json:"ExpiredTime"`
	Expiration   time.Time `json:"Expiration"`
	Token        string    `json:"Token"`
	Code         string    `json:"Code"`
}

// NewCvmRoleProvider need you to specify the roleName of the cvm currently in use
func NewCvmRoleProvider(roleName string) *CvmRoleProvider {
	return &CvmRoleProvider{roleName: roleName}
}

// DefaultCvmRoleProvider will auto get the cvm role name by accessing the metadata api
// more info please lookup: https://cloud.tencent.com/document/product/213/4934
func DefaultCvmRoleProvider() *CvmRoleProvider {
	return NewCvmRoleProvider("")
}

func get(url string) ([]byte, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode == http.StatusNotFound {
		return nil, roleNotBound
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func (r *CvmRoleProvider) getRoleName() (string, error) {
	if r.roleName != "" {
		return r.roleName, nil
	}
	rn, err := get(roleUrl)
	return string(rn), err
}

func (r *CvmRoleProvider) GetCredential() (CredentialIface, error) {
	roleName, err := r.getRoleName()
	if err != nil {
		return nil, noCvmRole
	}
	// get the cvm role name by accessing the metadata api
	// https://cloud.tencent.com/document/product/213/4934
	body, err := get(roleUrl + roleName)

	if err != nil {
		return nil, err
	}
	rspSt := new(roleRsp)
	if err = json.Unmarshal(body, rspSt); err != nil {
		return nil, tcerr.NewTencentCloudSDKError(creErr, err.Error(), "")
	}
	if rspSt.Code != "Success" {
		return nil, tcerr.NewTencentCloudSDKError(creErr, "Get credential from metadata server by role name "+roleName+" failed, code="+rspSt.Code, "")
	}
	cre := &CvmRoleCredential{
		tmpSecretId:  rspSt.TmpSecretId,
		tmpSecretKey: rspSt.TmpSecretKey,
		token:        rspSt.Token,
		roleName:     roleName,
		expiredTime:  rspSt.ExpiredTime,
		source:       r,
	}
	return cre, nil
}
