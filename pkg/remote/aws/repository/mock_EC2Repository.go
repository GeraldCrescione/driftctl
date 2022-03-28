// Code generated by mockery v2.10.0. DO NOT EDIT.

package repository

import (
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	mock "github.com/stretchr/testify/mock"
)

// MockEC2Repository is an autogenerated mock type for the EC2Repository type
type MockEC2Repository struct {
	mock.Mock
}

// DescribeLaunchTemplates provides a mock function with given fields:
func (_m *MockEC2Repository) DescribeLaunchTemplates() ([]*ec2.LaunchTemplate, error) {
	ret := _m.Called()

	var r0 []*ec2.LaunchTemplate
	if rf, ok := ret.Get(0).(func() []*ec2.LaunchTemplate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.LaunchTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsEbsEncryptionEnabledByDefault provides a mock function with given fields:
func (_m *MockEC2Repository) IsEbsEncryptionEnabledByDefault() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllAddresses provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllAddresses() ([]*ec2.Address, error) {
	ret := _m.Called()

	var r0 []*ec2.Address
	if rf, ok := ret.Get(0).(func() []*ec2.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllAddressesAssociation provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllAddressesAssociation() ([]*ec2.Address, error) {
	ret := _m.Called()

	var r0 []*ec2.Address
	if rf, ok := ret.Get(0).(func() []*ec2.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllImages provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllImages() ([]*ec2.Image, error) {
	ret := _m.Called()

	var r0 []*ec2.Image
	if rf, ok := ret.Get(0).(func() []*ec2.Image); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Image)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllInstances provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllInstances() ([]*ec2.Instance, error) {
	ret := _m.Called()

	var r0 []*ec2.Instance
	if rf, ok := ret.Get(0).(func() []*ec2.Instance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Instance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllInternetGateways provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllInternetGateways() ([]*ec2.InternetGateway, error) {
	ret := _m.Called()

	var r0 []*ec2.InternetGateway
	if rf, ok := ret.Get(0).(func() []*ec2.InternetGateway); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.InternetGateway)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllKeyPairs provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllKeyPairs() ([]*ec2.KeyPairInfo, error) {
	ret := _m.Called()

	var r0 []*ec2.KeyPairInfo
	if rf, ok := ret.Get(0).(func() []*ec2.KeyPairInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.KeyPairInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllNatGateways provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllNatGateways() ([]*ec2.NatGateway, error) {
	ret := _m.Called()

	var r0 []*ec2.NatGateway
	if rf, ok := ret.Get(0).(func() []*ec2.NatGateway); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.NatGateway)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllNetworkACLs provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllNetworkACLs() ([]*ec2.NetworkAcl, error) {
	ret := _m.Called()

	var r0 []*ec2.NetworkAcl
	if rf, ok := ret.Get(0).(func() []*ec2.NetworkAcl); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.NetworkAcl)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllRouteTables provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllRouteTables() ([]*ec2.RouteTable, error) {
	ret := _m.Called()

	var r0 []*ec2.RouteTable
	if rf, ok := ret.Get(0).(func() []*ec2.RouteTable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.RouteTable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllSecurityGroups provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllSecurityGroups() ([]*ec2.SecurityGroup, []*ec2.SecurityGroup, error) {
	ret := _m.Called()

	var r0 []*ec2.SecurityGroup
	if rf, ok := ret.Get(0).(func() []*ec2.SecurityGroup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.SecurityGroup)
		}
	}

	var r1 []*ec2.SecurityGroup
	if rf, ok := ret.Get(1).(func() []*ec2.SecurityGroup); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*ec2.SecurityGroup)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListAllSnapshots provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllSnapshots() ([]*ec2.Snapshot, error) {
	ret := _m.Called()

	var r0 []*ec2.Snapshot
	if rf, ok := ret.Get(0).(func() []*ec2.Snapshot); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Snapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllSubnets provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllSubnets() ([]*ec2.Subnet, []*ec2.Subnet, error) {
	ret := _m.Called()

	var r0 []*ec2.Subnet
	if rf, ok := ret.Get(0).(func() []*ec2.Subnet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Subnet)
		}
	}

	var r1 []*ec2.Subnet
	if rf, ok := ret.Get(1).(func() []*ec2.Subnet); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*ec2.Subnet)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListAllVPCs provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllVPCs() ([]*ec2.Vpc, []*ec2.Vpc, error) {
	ret := _m.Called()

	var r0 []*ec2.Vpc
	if rf, ok := ret.Get(0).(func() []*ec2.Vpc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Vpc)
		}
	}

	var r1 []*ec2.Vpc
	if rf, ok := ret.Get(1).(func() []*ec2.Vpc); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*ec2.Vpc)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListAllVolumes provides a mock function with given fields:
func (_m *MockEC2Repository) ListAllVolumes() ([]*ec2.Volume, error) {
	ret := _m.Called()

	var r0 []*ec2.Volume
	if rf, ok := ret.Get(0).(func() []*ec2.Volume); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ec2.Volume)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
