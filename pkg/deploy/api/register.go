package api

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
)

func init() {
	api.Scheme.AddKnownTypes("",
		&Deployment{},
		&DeploymentList{},
		&DeploymentConfig{},
		&DeploymentConfigList{},
		&DeploymentConfigRollback{},
	)
}

func (*Deployment) IsAnAPIObject()               {}
func (*DeploymentList) IsAnAPIObject()           {}
func (*DeploymentConfig) IsAnAPIObject()         {}
func (*DeploymentConfigList) IsAnAPIObject()     {}
func (*DeploymentConfigRollback) IsAnAPIObject() {}
