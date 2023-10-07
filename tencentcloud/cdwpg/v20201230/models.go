// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
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

package v20201230

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CBSSpec struct {
	// 盘类型
	DiskType *string `json:"DiskType,omitnil" name:"DiskType"`

	// 大小
	DiskSize *int64 `json:"DiskSize,omitnil" name:"DiskSize"`

	// 个数
	DiskCount *int64 `json:"DiskCount,omitnil" name:"DiskCount"`
}

type ChargeProperties struct {
	// 1-需要自动续期
	RenewFlag *int64 `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 订单时间范围
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 时间单位，一般为h和m
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 计费类型0-按量计费，1-包年包月
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// PREPAID、POSTPAID_BY_HOUR
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`
}

// Predefined struct for user
type CreateInstanceByApiRequestParams struct {
	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 私有网络
	UserVPCId *string `json:"UserVPCId,omitnil" name:"UserVPCId"`

	// 子网
	UserSubnetId *string `json:"UserSubnetId,omitnil" name:"UserSubnetId"`

	// 计费方式
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil" name:"ChargeProperties"`

	// 集群密码
	AdminPassword *string `json:"AdminPassword,omitnil" name:"AdminPassword"`

	// 资源信息
	Resources []*ResourceSpecNew `json:"Resources,omitnil" name:"Resources"`

	// 标签列表
	Tags *Tag `json:"Tags,omitnil" name:"Tags"`
}

type CreateInstanceByApiRequest struct {
	*tchttp.BaseRequest
	
	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 私有网络
	UserVPCId *string `json:"UserVPCId,omitnil" name:"UserVPCId"`

	// 子网
	UserSubnetId *string `json:"UserSubnetId,omitnil" name:"UserSubnetId"`

	// 计费方式
	ChargeProperties *ChargeProperties `json:"ChargeProperties,omitnil" name:"ChargeProperties"`

	// 集群密码
	AdminPassword *string `json:"AdminPassword,omitnil" name:"AdminPassword"`

	// 资源信息
	Resources []*ResourceSpecNew `json:"Resources,omitnil" name:"Resources"`

	// 标签列表
	Tags *Tag `json:"Tags,omitnil" name:"Tags"`
}

func (r *CreateInstanceByApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceByApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "Zone")
	delete(f, "UserVPCId")
	delete(f, "UserSubnetId")
	delete(f, "ChargeProperties")
	delete(f, "AdminPassword")
	delete(f, "Resources")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceByApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceByApiResponseParams struct {
	// 流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateInstanceByApiResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceByApiResponseParams `json:"Response"`
}

func (r *CreateInstanceByApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceByApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceByApiRequestParams struct {
	// 实例名称，例如"cdwpg-xxxx"
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DestroyInstanceByApiRequest struct {
	*tchttp.BaseRequest
	
	// 实例名称，例如"cdwpg-xxxx"
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DestroyInstanceByApiRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceByApiRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyInstanceByApiRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyInstanceByApiResponseParams struct {
	// 销毁流程Id
	FlowId *string `json:"FlowId,omitnil" name:"FlowId"`

	// 错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DestroyInstanceByApiResponse struct {
	*tchttp.BaseResponse
	Response *DestroyInstanceByApiResponseParams `json:"Response"`
}

func (r *DestroyInstanceByApiResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyInstanceByApiResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceSpecNew struct {
	// 资源名称
	SpecName *string `json:"SpecName,omitnil" name:"SpecName"`

	// 资源数
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 磁盘信息
	DiskSpec *CBSSpec `json:"DiskSpec,omitnil" name:"DiskSpec"`

	// 资源类型，DATA
	Type *string `json:"Type,omitnil" name:"Type"`
}

type Tag struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}