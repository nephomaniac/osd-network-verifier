// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/clients/aws/aws.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	gomock "github.com/golang/mock/gomock"
)

// MockEC2Client is a mock of EC2Client interface.
type MockEC2Client struct {
	ctrl     *gomock.Controller
	recorder *MockEC2ClientMockRecorder
}

// MockEC2ClientMockRecorder is the mock recorder for MockEC2Client.
type MockEC2ClientMockRecorder struct {
	mock *MockEC2Client
}

// NewMockEC2Client creates a new mock instance.
func NewMockEC2Client(ctrl *gomock.Controller) *MockEC2Client {
	mock := &MockEC2Client{ctrl: ctrl}
	mock.recorder = &MockEC2ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEC2Client) EXPECT() *MockEC2ClientMockRecorder {
	return m.recorder
}

// AuthorizeSecurityGroupEgress mocks base method.
func (m *MockEC2Client) AuthorizeSecurityGroupEgress(ctx context.Context, params *ec2.AuthorizeSecurityGroupEgressInput, optFns ...func(*ec2.Options)) (*ec2.AuthorizeSecurityGroupEgressOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthorizeSecurityGroupEgress", varargs...)
	ret0, _ := ret[0].(*ec2.AuthorizeSecurityGroupEgressOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthorizeSecurityGroupEgress indicates an expected call of AuthorizeSecurityGroupEgress.
func (mr *MockEC2ClientMockRecorder) AuthorizeSecurityGroupEgress(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizeSecurityGroupEgress", reflect.TypeOf((*MockEC2Client)(nil).AuthorizeSecurityGroupEgress), varargs...)
}

// CreateSecurityGroup mocks base method.
func (m *MockEC2Client) CreateSecurityGroup(ctx context.Context, params *ec2.CreateSecurityGroupInput, optFns ...func(*ec2.Options)) (*ec2.CreateSecurityGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSecurityGroup", varargs...)
	ret0, _ := ret[0].(*ec2.CreateSecurityGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecurityGroup indicates an expected call of CreateSecurityGroup.
func (mr *MockEC2ClientMockRecorder) CreateSecurityGroup(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecurityGroup", reflect.TypeOf((*MockEC2Client)(nil).CreateSecurityGroup), varargs...)
}

// CreateTags mocks base method.
func (m *MockEC2Client) CreateTags(ctx context.Context, params *ec2.CreateTagsInput, optFns ...func(*ec2.Options)) (*ec2.CreateTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTags", varargs...)
	ret0, _ := ret[0].(*ec2.CreateTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTags indicates an expected call of CreateTags.
func (mr *MockEC2ClientMockRecorder) CreateTags(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTags", reflect.TypeOf((*MockEC2Client)(nil).CreateTags), varargs...)
}

// DeleteSecurityGroup mocks base method.
func (m *MockEC2Client) DeleteSecurityGroup(ctx context.Context, params *ec2.DeleteSecurityGroupInput, optFns ...func(*ec2.Options)) (*ec2.DeleteSecurityGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSecurityGroup", varargs...)
	ret0, _ := ret[0].(*ec2.DeleteSecurityGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSecurityGroup indicates an expected call of DeleteSecurityGroup.
func (mr *MockEC2ClientMockRecorder) DeleteSecurityGroup(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecurityGroup", reflect.TypeOf((*MockEC2Client)(nil).DeleteSecurityGroup), varargs...)
}

// DescribeInstanceTypes mocks base method.
func (m *MockEC2Client) DescribeInstanceTypes(ctx context.Context, input *ec2.DescribeInstanceTypesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstanceTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstanceTypes", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstanceTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstanceTypes indicates an expected call of DescribeInstanceTypes.
func (mr *MockEC2ClientMockRecorder) DescribeInstanceTypes(ctx, input interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstanceTypes", reflect.TypeOf((*MockEC2Client)(nil).DescribeInstanceTypes), varargs...)
}

// DescribeInstances mocks base method.
func (m *MockEC2Client) DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstances", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstances indicates an expected call of DescribeInstances.
func (mr *MockEC2ClientMockRecorder) DescribeInstances(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstances", reflect.TypeOf((*MockEC2Client)(nil).DescribeInstances), varargs...)
}

// DescribeSecurityGroups mocks base method.
func (m *MockEC2Client) DescribeSecurityGroups(ctx context.Context, params *ec2.DescribeSecurityGroupsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSecurityGroups", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSecurityGroups indicates an expected call of DescribeSecurityGroups.
func (mr *MockEC2ClientMockRecorder) DescribeSecurityGroups(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecurityGroups", reflect.TypeOf((*MockEC2Client)(nil).DescribeSecurityGroups), varargs...)
}

// DescribeSubnets mocks base method.
func (m *MockEC2Client) DescribeSubnets(ctx context.Context, params *ec2.DescribeSubnetsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSubnets", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSubnetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSubnets indicates an expected call of DescribeSubnets.
func (mr *MockEC2ClientMockRecorder) DescribeSubnets(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSubnets", reflect.TypeOf((*MockEC2Client)(nil).DescribeSubnets), varargs...)
}

// DescribeVpcAttribute mocks base method.
func (m *MockEC2Client) DescribeVpcAttribute(ctx context.Context, input *ec2.DescribeVpcAttributeInput, optFns ...func(*ec2.Options)) (*ec2.DescribeVpcAttributeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVpcAttribute", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeVpcAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcAttribute indicates an expected call of DescribeVpcAttribute.
func (mr *MockEC2ClientMockRecorder) DescribeVpcAttribute(ctx, input interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcAttribute", reflect.TypeOf((*MockEC2Client)(nil).DescribeVpcAttribute), varargs...)
}

// GetConsoleOutput mocks base method.
func (m *MockEC2Client) GetConsoleOutput(ctx context.Context, input *ec2.GetConsoleOutputInput, optFns ...func(*ec2.Options)) (*ec2.GetConsoleOutputOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConsoleOutput", varargs...)
	ret0, _ := ret[0].(*ec2.GetConsoleOutputOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsoleOutput indicates an expected call of GetConsoleOutput.
func (mr *MockEC2ClientMockRecorder) GetConsoleOutput(ctx, input interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsoleOutput", reflect.TypeOf((*MockEC2Client)(nil).GetConsoleOutput), varargs...)
}

// RevokeSecurityGroupEgress mocks base method.
func (m *MockEC2Client) RevokeSecurityGroupEgress(ctx context.Context, params *ec2.RevokeSecurityGroupEgressInput, optFns ...func(*ec2.Options)) (*ec2.RevokeSecurityGroupEgressOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RevokeSecurityGroupEgress", varargs...)
	ret0, _ := ret[0].(*ec2.RevokeSecurityGroupEgressOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeSecurityGroupEgress indicates an expected call of RevokeSecurityGroupEgress.
func (mr *MockEC2ClientMockRecorder) RevokeSecurityGroupEgress(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeSecurityGroupEgress", reflect.TypeOf((*MockEC2Client)(nil).RevokeSecurityGroupEgress), varargs...)
}

// RunInstances mocks base method.
func (m *MockEC2Client) RunInstances(ctx context.Context, params *ec2.RunInstancesInput, optFns ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunInstances", varargs...)
	ret0, _ := ret[0].(*ec2.RunInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunInstances indicates an expected call of RunInstances.
func (mr *MockEC2ClientMockRecorder) RunInstances(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInstances", reflect.TypeOf((*MockEC2Client)(nil).RunInstances), varargs...)
}

// TerminateInstances mocks base method.
func (m *MockEC2Client) TerminateInstances(ctx context.Context, input *ec2.TerminateInstancesInput, optFns ...func(*ec2.Options)) (*ec2.TerminateInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TerminateInstances", varargs...)
	ret0, _ := ret[0].(*ec2.TerminateInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TerminateInstances indicates an expected call of TerminateInstances.
func (mr *MockEC2ClientMockRecorder) TerminateInstances(ctx, input interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateInstances", reflect.TypeOf((*MockEC2Client)(nil).TerminateInstances), varargs...)
}
