package machine

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

	machinev1 "github.com/openshift/api/machine/v1"
	machinev1beta1 "github.com/openshift/api/machine/v1beta1"
	machinecontroller "github.com/openshift/machine-api-operator/pkg/controller/machine"

	alibabacloudproviderv1 "github.com/openshift/cluster-api-provider-alibaba/pkg/apis/alibabacloudprovider/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	defaultNamespace                     = "default"
	stubZoneID                           = "cn-beijing"
	stubRegionID                         = "cn-beijing-f"
	alibabaCloudCredentialsSecretName    = "alibabacloud-credentials-secret"
	alibabaCloudMasterUserDataSecretName = "master-user-data-secret"
	alibabaCloudWorkerUserDataSecretName = "worker-user-data-secret"

	stubMasterMachineName = "alibabacloud-actuator-testing-master-machine"
	stubWorkerMachineName = "alibabacloud-actuator-testing-worker-machine"

	stubKeyName                 = "alibabacloud-actuator-key-name"
	stubClusterID               = "alibabacloud-actuator-cluster"
	stubImageID                 = "centos_7_9_x64_20G_alibase_20210318.vhd"
	stubVpcID                   = "vpc-3ze4u29pd4lniym7i1xnp"
	stubVSwitchID               = "vsw-7ze567qrl5das7q8s4rei"
	stubInstanceID              = "i-2ze3hj0qh9d290rpax7w"
	stubSecurityGroupID         = "sg-2zeebk9qd965vc2xqq4w"
	stubSystemDiskCategory      = "cloud_essd"
	stubSystemDiskSize          = 120
	stubInternetMaxBandwidthOut = 100
	stubPassword                = "Hello$1234"
	stubInstanceType            = "ecs.c6.2xlarge"
)

func stubAlibabaCloudCredentialsSecret() *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      alibabaCloudCredentialsSecretName,
			Namespace: defaultNamespace,
		},
		Data: map[string][]byte{
			"accessKeyID":     []byte(os.Getenv("ALIBABACLOUD_ACCESS_KEY_ID")),
			"accessKeySecret": []byte(os.Getenv("ALIBABACLOUD_SECRET_ACCESS_KEY")),
		},
	}
}

func stubProviderConfig() *machinev1.AlibabaCloudMachineProviderConfig {
	vSwitchID := stubVSwitchID
	return &machinev1.AlibabaCloudMachineProviderConfig{
		InstanceType: stubInstanceType,
		ImageID:      stubImageID,
		RegionID:     stubRegionID,
		ZoneID:       stubZoneID,
		SecurityGroups: []machinev1.AlibabaResourceReference{
			{
				Type: machinev1.AlibabaResourceReferenceTypeID,
				ID:   &vSwitchID,
			},
		},
		VpcID: stubVpcID,
		VSwitch: machinev1.AlibabaResourceReference{
			Type: machinev1.AlibabaResourceReferenceTypeID,
			ID:   &vSwitchID,
		},
		SystemDisk: machinev1.SystemDiskProperties{
			Category: stubSystemDiskCategory,
			Size:     stubSystemDiskSize,
		},
		Bandwidth: machinev1.BandwidthProperties{
			InternetMaxBandwidthOut: stubInternetMaxBandwidthOut,
		},
		CredentialsSecret: &corev1.LocalObjectReference{
			Name: alibabaCloudCredentialsSecretName,
		},
		Tags: []machinev1.Tag{
			{Key: "openshift-node-group-config", Value: "node-config-master"},
			{Key: "host-type", Value: "master"},
			{Key: "sub-host-type", Value: "default"},
		},
	}
}

func stubMasterMachine() (*machinev1beta1.Machine, error) {
	masterMachine, err := stubMachine(stubMasterMachineName, map[string]string{
		"node-role.kubernetes.io/master": "",
		"node-role.kubernetes.io/infra":  "",
	})

	if err != nil {
		return nil, err
	}

	return masterMachine, nil
}

func stubWorkerMachine() (*machinev1beta1.Machine, error) {
	workerMachine, err := stubMachine(stubWorkerMachineName, map[string]string{
		"node-role.kubernetes.io/infra": "",
	})

	if err != nil {
		return nil, err
	}

	return workerMachine, nil
}

func stubMachine(machineName string, machineLabels map[string]string) (*machinev1beta1.Machine, error) {
	machineSpec := stubProviderConfig()

	providerSpec, err := alibabacloudproviderv1.RawExtensionFromProviderSpec(machineSpec)
	if err != nil {
		return nil, fmt.Errorf("codec.EncodeProviderSpec failed: %v", err)
	}

	machine := &machinev1beta1.Machine{
		ObjectMeta: metav1.ObjectMeta{
			Name:      machineName,
			Namespace: defaultNamespace,
			Labels: map[string]string{
				machinev1beta1.MachineClusterIDLabel: stubClusterID,
			},
			Annotations: map[string]string{
				// skip node draining since it's not mocked
				machinecontroller.ExcludeNodeDrainingAnnotation: "",
			},
		},
		Spec: machinev1beta1.MachineSpec{
			ObjectMeta: machinev1beta1.ObjectMeta{
				Labels: machineLabels,
			},
			ProviderSpec: machinev1beta1.ProviderSpec{
				Value: providerSpec,
			},
		},
	}
	return machine, nil
}

func stubRunInstancesRequest() *ecs.RunInstancesRequest {
	request := ecs.CreateRunInstancesRequest()
	request.Scheme = "https"
	request.RegionId = stubRegionID
	request.InstanceType = stubInstanceType
	request.ImageId = stubImageID
	request.VSwitchId = stubVSwitchID
	request.SecurityGroupId = stubSecurityGroupID
	request.Password = stubPassword
	request.MinAmount = requests.NewInteger(1)
	request.Amount = requests.NewInteger(1)

	request.SystemDiskCategory = stubSystemDiskCategory
	request.SystemDiskSize = strconv.Itoa(stubSystemDiskSize)

	return request
}

func stubRunInstancesResponse() *ecs.RunInstancesResponse {
	response := ecs.CreateRunInstancesResponse()
	response.InstanceIdSets = ecs.InstanceIdSets{
		InstanceIdSet: []string{stubInstanceID},
	}

	return response
}
