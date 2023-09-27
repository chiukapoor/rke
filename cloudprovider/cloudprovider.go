package cloudprovider

import (
	"github.com/chiukapoor/rke/cloudprovider/aws"
	"github.com/chiukapoor/rke/cloudprovider/azure"
	"github.com/chiukapoor/rke/cloudprovider/custom"
	"github.com/chiukapoor/rke/cloudprovider/harvester"
	"github.com/chiukapoor/rke/cloudprovider/openstack"
	"github.com/chiukapoor/rke/cloudprovider/vsphere"
	v3 "github.com/chiukapoor/rke/types"
)

type CloudProvider interface {
	Init(cloudProviderConfig v3.CloudProvider) error
	GenerateCloudConfigFile() (string, error)
	GetName() string
}

func InitCloudProvider(cloudProviderConfig v3.CloudProvider) (CloudProvider, error) {
	var p CloudProvider
	if cloudProviderConfig.AWSCloudProvider != nil || cloudProviderConfig.Name == aws.AWSCloudProviderName {
		p = aws.GetInstance()
	}
	if cloudProviderConfig.AzureCloudProvider != nil || cloudProviderConfig.Name == azure.AzureCloudProviderName {
		p = azure.GetInstance()
	}
	if cloudProviderConfig.OpenstackCloudProvider != nil || cloudProviderConfig.Name == openstack.OpenstackCloudProviderName {
		p = openstack.GetInstance()
	}
	if cloudProviderConfig.VsphereCloudProvider != nil || cloudProviderConfig.Name == vsphere.VsphereCloudProviderName {
		p = vsphere.GetInstance()
	}
	if cloudProviderConfig.HarvesterCloudProvider != nil || cloudProviderConfig.Name == harvester.HarvesterCloudProviderName {
		p = harvester.GetInstance()
	}
	if cloudProviderConfig.CustomCloudProvider != "" {
		p = custom.GetInstance()
	}

	if p != nil {
		if err := p.Init(cloudProviderConfig); err != nil {
			return nil, err
		}
	}
	return p, nil
}
