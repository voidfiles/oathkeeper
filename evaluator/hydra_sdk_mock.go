// Automatically generated by MockGen. DO NOT EDIT!
// Source: ./vendor/github.com/ory/hydra/sdk/go/hydra/sdk_api.go

package evaluator

import (
	gomock "github.com/golang/mock/gomock"
	swagger "github.com/ory/hydra/sdk/go/hydra/swagger"
	oauth2 "golang.org/x/oauth2"
	clientcredentials "golang.org/x/oauth2/clientcredentials"
)

// Mock of SDK interface
type MockSDK struct {
	ctrl     *gomock.Controller
	recorder *_MockSDKRecorder
}

// Recorder for MockSDK (not exported)
type _MockSDKRecorder struct {
	mock *MockSDK
}

func NewMockSDK(ctrl *gomock.Controller) *MockSDK {
	mock := &MockSDK{ctrl: ctrl}
	mock.recorder = &_MockSDKRecorder{mock}
	return mock
}

func (_m *MockSDK) EXPECT() *_MockSDKRecorder {
	return _m.recorder
}

func (_m *MockSDK) GetOAuth2ClientConfig() *clientcredentials.Config {
	ret := _m.ctrl.Call(_m, "GetOAuth2ClientConfig")
	ret0, _ := ret[0].(*clientcredentials.Config)
	return ret0
}

func (_mr *_MockSDKRecorder) GetOAuth2ClientConfig() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOAuth2ClientConfig")
}

func (_m *MockSDK) GetOAuth2Config() *oauth2.Config {
	ret := _m.ctrl.Call(_m, "GetOAuth2Config")
	ret0, _ := ret[0].(*oauth2.Config)
	return ret0
}

func (_mr *_MockSDKRecorder) GetOAuth2Config() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOAuth2Config")
}

func (_m *MockSDK) CreatePolicy(body swagger.Policy) (*swagger.Policy, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "CreatePolicy", body)
	ret0, _ := ret[0].(*swagger.Policy)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) CreatePolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreatePolicy", arg0)
}

func (_m *MockSDK) DeletePolicy(id string) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DeletePolicy", id)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSDKRecorder) DeletePolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeletePolicy", arg0)
}

func (_m *MockSDK) GetPolicy(id string) (*swagger.Policy, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetPolicy", id)
	ret0, _ := ret[0].(*swagger.Policy)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) GetPolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPolicy", arg0)
}

func (_m *MockSDK) ListPolicies(offset int64, limit int64) ([]swagger.Policy, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "ListPolicies", offset, limit)
	ret0, _ := ret[0].([]swagger.Policy)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) ListPolicies(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListPolicies", arg0, arg1)
}

func (_m *MockSDK) UpdatePolicy(id string, body swagger.Policy) (*swagger.Policy, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "UpdatePolicy", id, body)
	ret0, _ := ret[0].(*swagger.Policy)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) UpdatePolicy(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdatePolicy", arg0, arg1)
}

func (_m *MockSDK) AddMembersToGroup(id string, body swagger.GroupMembers) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "AddMembersToGroup", id, body)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSDKRecorder) AddMembersToGroup(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddMembersToGroup", arg0, arg1)
}

func (_m *MockSDK) CreateGroup(body swagger.Group) (*swagger.Group, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateGroup", body)
	ret0, _ := ret[0].(*swagger.Group)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) CreateGroup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateGroup", arg0)
}

func (_m *MockSDK) DeleteGroup(id string) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteGroup", id)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSDKRecorder) DeleteGroup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteGroup", arg0)
}

func (_m *MockSDK) DoesWardenAllowAccessRequest(body swagger.WardenAccessRequest) (*swagger.WardenAccessRequestResponse, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DoesWardenAllowAccessRequest", body)
	ret0, _ := ret[0].(*swagger.WardenAccessRequestResponse)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) DoesWardenAllowAccessRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DoesWardenAllowAccessRequest", arg0)
}

func (_m *MockSDK) DoesWardenAllowTokenAccessRequest(body swagger.WardenTokenAccessRequest) (*swagger.WardenTokenAccessRequestResponse, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DoesWardenAllowTokenAccessRequest", body)
	ret0, _ := ret[0].(*swagger.WardenTokenAccessRequestResponse)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) DoesWardenAllowTokenAccessRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DoesWardenAllowTokenAccessRequest", arg0)
}

func (_m *MockSDK) ListGroups(member string, limit int64, offset int64) ([]swagger.Group, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "ListGroups", member, limit, offset)
	ret0, _ := ret[0].([]swagger.Group)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) ListGroups(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListGroups", arg0, arg1, arg2)
}

func (_m *MockSDK) GetGroup(id string) (*swagger.Group, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetGroup", id)
	ret0, _ := ret[0].(*swagger.Group)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) GetGroup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetGroup", arg0)
}

func (_m *MockSDK) RemoveMembersFromGroup(id string, body swagger.GroupMembers) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "RemoveMembersFromGroup", id, body)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSDKRecorder) RemoveMembersFromGroup(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveMembersFromGroup", arg0, arg1)
}

func (_m *MockSDK) CreateJsonWebKeySet(set string, body swagger.JsonWebKeySetGeneratorRequest) (*swagger.JsonWebKeySet, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateJsonWebKeySet", set, body)
	ret0, _ := ret[0].(*swagger.JsonWebKeySet)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) CreateJsonWebKeySet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateJsonWebKeySet", arg0, arg1)
}

func (_m *MockSDK) DeleteJsonWebKey(kid string, set string) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteJsonWebKey", kid, set)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSDKRecorder) DeleteJsonWebKey(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteJsonWebKey", arg0, arg1)
}

func (_m *MockSDK) DeleteJsonWebKeySet(set string) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteJsonWebKeySet", set)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSDKRecorder) DeleteJsonWebKeySet(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteJsonWebKeySet", arg0)
}

func (_m *MockSDK) GetJsonWebKey(kid string, set string) (*swagger.JsonWebKeySet, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetJsonWebKey", kid, set)
	ret0, _ := ret[0].(*swagger.JsonWebKeySet)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) GetJsonWebKey(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJsonWebKey", arg0, arg1)
}

func (_m *MockSDK) GetJsonWebKeySet(set string) (*swagger.JsonWebKeySet, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetJsonWebKeySet", set)
	ret0, _ := ret[0].(*swagger.JsonWebKeySet)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) GetJsonWebKeySet(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJsonWebKeySet", arg0)
}

func (_m *MockSDK) UpdateJsonWebKey(kid string, set string, body swagger.JsonWebKey) (*swagger.JsonWebKey, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "UpdateJsonWebKey", kid, set, body)
	ret0, _ := ret[0].(*swagger.JsonWebKey)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) UpdateJsonWebKey(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateJsonWebKey", arg0, arg1, arg2)
}

func (_m *MockSDK) UpdateJsonWebKeySet(set string, body swagger.JsonWebKeySet) (*swagger.JsonWebKeySet, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "UpdateJsonWebKeySet", set, body)
	ret0, _ := ret[0].(*swagger.JsonWebKeySet)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) UpdateJsonWebKeySet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateJsonWebKeySet", arg0, arg1)
}

func (_m *MockSDK) AcceptOAuth2ConsentRequest(id string, body swagger.ConsentRequestAcceptance) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "AcceptOAuth2ConsentRequest", id, body)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSDKRecorder) AcceptOAuth2ConsentRequest(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AcceptOAuth2ConsentRequest", arg0, arg1)
}

func (_m *MockSDK) CreateOAuth2Client(body swagger.OAuth2Client) (*swagger.OAuth2Client, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateOAuth2Client", body)
	ret0, _ := ret[0].(*swagger.OAuth2Client)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) CreateOAuth2Client(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateOAuth2Client", arg0)
}

func (_m *MockSDK) DeleteOAuth2Client(id string) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteOAuth2Client", id)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSDKRecorder) DeleteOAuth2Client(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteOAuth2Client", arg0)
}

func (_m *MockSDK) GetOAuth2Client(id string) (*swagger.OAuth2Client, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetOAuth2Client", id)
	ret0, _ := ret[0].(*swagger.OAuth2Client)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) GetOAuth2Client(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOAuth2Client", arg0)
}

func (_m *MockSDK) GetOAuth2ConsentRequest(id string) (*swagger.OAuth2ConsentRequest, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetOAuth2ConsentRequest", id)
	ret0, _ := ret[0].(*swagger.OAuth2ConsentRequest)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) GetOAuth2ConsentRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOAuth2ConsentRequest", arg0)
}

func (_m *MockSDK) GetWellKnown() (*swagger.WellKnown, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetWellKnown")
	ret0, _ := ret[0].(*swagger.WellKnown)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) GetWellKnown() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWellKnown")
}

func (_m *MockSDK) IntrospectOAuth2Token(token string, scope string) (*swagger.OAuth2TokenIntrospection, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "IntrospectOAuth2Token", token, scope)
	ret0, _ := ret[0].(*swagger.OAuth2TokenIntrospection)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) IntrospectOAuth2Token(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IntrospectOAuth2Token", arg0, arg1)
}

func (_m *MockSDK) ListOAuth2Clients() ([]swagger.OAuth2Client, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "ListOAuth2Clients")
	ret0, _ := ret[0].([]swagger.OAuth2Client)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) ListOAuth2Clients() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListOAuth2Clients")
}

func (_m *MockSDK) RejectOAuth2ConsentRequest(id string, body swagger.ConsentRequestRejection) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "RejectOAuth2ConsentRequest", id, body)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSDKRecorder) RejectOAuth2ConsentRequest(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RejectOAuth2ConsentRequest", arg0, arg1)
}

func (_m *MockSDK) RevokeOAuth2Token(token string) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "RevokeOAuth2Token", token)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSDKRecorder) RevokeOAuth2Token(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RevokeOAuth2Token", arg0)
}

func (_m *MockSDK) UpdateOAuth2Client(id string, body swagger.OAuth2Client) (*swagger.OAuth2Client, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "UpdateOAuth2Client", id, body)
	ret0, _ := ret[0].(*swagger.OAuth2Client)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockSDKRecorder) UpdateOAuth2Client(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateOAuth2Client", arg0, arg1)
}

// Mock of PolicyAPI interface
type MockPolicyAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockPolicyAPIRecorder
}

// Recorder for MockPolicyAPI (not exported)
type _MockPolicyAPIRecorder struct {
	mock *MockPolicyAPI
}

func NewMockPolicyAPI(ctrl *gomock.Controller) *MockPolicyAPI {
	mock := &MockPolicyAPI{ctrl: ctrl}
	mock.recorder = &_MockPolicyAPIRecorder{mock}
	return mock
}

func (_m *MockPolicyAPI) EXPECT() *_MockPolicyAPIRecorder {
	return _m.recorder
}

func (_m *MockPolicyAPI) CreatePolicy(body swagger.Policy) (*swagger.Policy, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "CreatePolicy", body)
	ret0, _ := ret[0].(*swagger.Policy)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockPolicyAPIRecorder) CreatePolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreatePolicy", arg0)
}

func (_m *MockPolicyAPI) DeletePolicy(id string) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DeletePolicy", id)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockPolicyAPIRecorder) DeletePolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeletePolicy", arg0)
}

func (_m *MockPolicyAPI) GetPolicy(id string) (*swagger.Policy, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetPolicy", id)
	ret0, _ := ret[0].(*swagger.Policy)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockPolicyAPIRecorder) GetPolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPolicy", arg0)
}

func (_m *MockPolicyAPI) ListPolicies(offset int64, limit int64) ([]swagger.Policy, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "ListPolicies", offset, limit)
	ret0, _ := ret[0].([]swagger.Policy)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockPolicyAPIRecorder) ListPolicies(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListPolicies", arg0, arg1)
}

func (_m *MockPolicyAPI) UpdatePolicy(id string, body swagger.Policy) (*swagger.Policy, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "UpdatePolicy", id, body)
	ret0, _ := ret[0].(*swagger.Policy)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockPolicyAPIRecorder) UpdatePolicy(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdatePolicy", arg0, arg1)
}

// Mock of WardenAPI interface
type MockWardenAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockWardenAPIRecorder
}

// Recorder for MockWardenAPI (not exported)
type _MockWardenAPIRecorder struct {
	mock *MockWardenAPI
}

func NewMockWardenAPI(ctrl *gomock.Controller) *MockWardenAPI {
	mock := &MockWardenAPI{ctrl: ctrl}
	mock.recorder = &_MockWardenAPIRecorder{mock}
	return mock
}

func (_m *MockWardenAPI) EXPECT() *_MockWardenAPIRecorder {
	return _m.recorder
}

func (_m *MockWardenAPI) AddMembersToGroup(id string, body swagger.GroupMembers) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "AddMembersToGroup", id, body)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockWardenAPIRecorder) AddMembersToGroup(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddMembersToGroup", arg0, arg1)
}

func (_m *MockWardenAPI) CreateGroup(body swagger.Group) (*swagger.Group, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateGroup", body)
	ret0, _ := ret[0].(*swagger.Group)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockWardenAPIRecorder) CreateGroup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateGroup", arg0)
}

func (_m *MockWardenAPI) DeleteGroup(id string) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteGroup", id)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockWardenAPIRecorder) DeleteGroup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteGroup", arg0)
}

func (_m *MockWardenAPI) DoesWardenAllowAccessRequest(body swagger.WardenAccessRequest) (*swagger.WardenAccessRequestResponse, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DoesWardenAllowAccessRequest", body)
	ret0, _ := ret[0].(*swagger.WardenAccessRequestResponse)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockWardenAPIRecorder) DoesWardenAllowAccessRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DoesWardenAllowAccessRequest", arg0)
}

func (_m *MockWardenAPI) DoesWardenAllowTokenAccessRequest(body swagger.WardenTokenAccessRequest) (*swagger.WardenTokenAccessRequestResponse, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DoesWardenAllowTokenAccessRequest", body)
	ret0, _ := ret[0].(*swagger.WardenTokenAccessRequestResponse)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockWardenAPIRecorder) DoesWardenAllowTokenAccessRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DoesWardenAllowTokenAccessRequest", arg0)
}

func (_m *MockWardenAPI) ListGroups(member string, limit int64, offset int64) ([]swagger.Group, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "ListGroups", member, limit, offset)
	ret0, _ := ret[0].([]swagger.Group)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockWardenAPIRecorder) ListGroups(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListGroups", arg0, arg1, arg2)
}

func (_m *MockWardenAPI) GetGroup(id string) (*swagger.Group, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetGroup", id)
	ret0, _ := ret[0].(*swagger.Group)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockWardenAPIRecorder) GetGroup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetGroup", arg0)
}

func (_m *MockWardenAPI) RemoveMembersFromGroup(id string, body swagger.GroupMembers) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "RemoveMembersFromGroup", id, body)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockWardenAPIRecorder) RemoveMembersFromGroup(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveMembersFromGroup", arg0, arg1)
}

// Mock of JWKApi interface
type MockJWKApi struct {
	ctrl     *gomock.Controller
	recorder *_MockJWKApiRecorder
}

// Recorder for MockJWKApi (not exported)
type _MockJWKApiRecorder struct {
	mock *MockJWKApi
}

func NewMockJWKApi(ctrl *gomock.Controller) *MockJWKApi {
	mock := &MockJWKApi{ctrl: ctrl}
	mock.recorder = &_MockJWKApiRecorder{mock}
	return mock
}

func (_m *MockJWKApi) EXPECT() *_MockJWKApiRecorder {
	return _m.recorder
}

func (_m *MockJWKApi) CreateJsonWebKeySet(set string, body swagger.JsonWebKeySetGeneratorRequest) (*swagger.JsonWebKeySet, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateJsonWebKeySet", set, body)
	ret0, _ := ret[0].(*swagger.JsonWebKeySet)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockJWKApiRecorder) CreateJsonWebKeySet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateJsonWebKeySet", arg0, arg1)
}

func (_m *MockJWKApi) DeleteJsonWebKey(kid string, set string) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteJsonWebKey", kid, set)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJWKApiRecorder) DeleteJsonWebKey(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteJsonWebKey", arg0, arg1)
}

func (_m *MockJWKApi) DeleteJsonWebKeySet(set string) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteJsonWebKeySet", set)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockJWKApiRecorder) DeleteJsonWebKeySet(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteJsonWebKeySet", arg0)
}

func (_m *MockJWKApi) GetJsonWebKey(kid string, set string) (*swagger.JsonWebKeySet, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetJsonWebKey", kid, set)
	ret0, _ := ret[0].(*swagger.JsonWebKeySet)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockJWKApiRecorder) GetJsonWebKey(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJsonWebKey", arg0, arg1)
}

func (_m *MockJWKApi) GetJsonWebKeySet(set string) (*swagger.JsonWebKeySet, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetJsonWebKeySet", set)
	ret0, _ := ret[0].(*swagger.JsonWebKeySet)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockJWKApiRecorder) GetJsonWebKeySet(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJsonWebKeySet", arg0)
}

func (_m *MockJWKApi) UpdateJsonWebKey(kid string, set string, body swagger.JsonWebKey) (*swagger.JsonWebKey, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "UpdateJsonWebKey", kid, set, body)
	ret0, _ := ret[0].(*swagger.JsonWebKey)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockJWKApiRecorder) UpdateJsonWebKey(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateJsonWebKey", arg0, arg1, arg2)
}

func (_m *MockJWKApi) UpdateJsonWebKeySet(set string, body swagger.JsonWebKeySet) (*swagger.JsonWebKeySet, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "UpdateJsonWebKeySet", set, body)
	ret0, _ := ret[0].(*swagger.JsonWebKeySet)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockJWKApiRecorder) UpdateJsonWebKeySet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateJsonWebKeySet", arg0, arg1)
}

// Mock of OAuth2API interface
type MockOAuth2API struct {
	ctrl     *gomock.Controller
	recorder *_MockOAuth2APIRecorder
}

// Recorder for MockOAuth2API (not exported)
type _MockOAuth2APIRecorder struct {
	mock *MockOAuth2API
}

func NewMockOAuth2API(ctrl *gomock.Controller) *MockOAuth2API {
	mock := &MockOAuth2API{ctrl: ctrl}
	mock.recorder = &_MockOAuth2APIRecorder{mock}
	return mock
}

func (_m *MockOAuth2API) EXPECT() *_MockOAuth2APIRecorder {
	return _m.recorder
}

func (_m *MockOAuth2API) AcceptOAuth2ConsentRequest(id string, body swagger.ConsentRequestAcceptance) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "AcceptOAuth2ConsentRequest", id, body)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOAuth2APIRecorder) AcceptOAuth2ConsentRequest(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AcceptOAuth2ConsentRequest", arg0, arg1)
}

func (_m *MockOAuth2API) CreateOAuth2Client(body swagger.OAuth2Client) (*swagger.OAuth2Client, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateOAuth2Client", body)
	ret0, _ := ret[0].(*swagger.OAuth2Client)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockOAuth2APIRecorder) CreateOAuth2Client(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateOAuth2Client", arg0)
}

func (_m *MockOAuth2API) DeleteOAuth2Client(id string) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteOAuth2Client", id)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOAuth2APIRecorder) DeleteOAuth2Client(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteOAuth2Client", arg0)
}

func (_m *MockOAuth2API) GetOAuth2Client(id string) (*swagger.OAuth2Client, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetOAuth2Client", id)
	ret0, _ := ret[0].(*swagger.OAuth2Client)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockOAuth2APIRecorder) GetOAuth2Client(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOAuth2Client", arg0)
}

func (_m *MockOAuth2API) GetOAuth2ConsentRequest(id string) (*swagger.OAuth2ConsentRequest, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetOAuth2ConsentRequest", id)
	ret0, _ := ret[0].(*swagger.OAuth2ConsentRequest)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockOAuth2APIRecorder) GetOAuth2ConsentRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOAuth2ConsentRequest", arg0)
}

func (_m *MockOAuth2API) GetWellKnown() (*swagger.WellKnown, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "GetWellKnown")
	ret0, _ := ret[0].(*swagger.WellKnown)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockOAuth2APIRecorder) GetWellKnown() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWellKnown")
}

func (_m *MockOAuth2API) IntrospectOAuth2Token(token string, scope string) (*swagger.OAuth2TokenIntrospection, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "IntrospectOAuth2Token", token, scope)
	ret0, _ := ret[0].(*swagger.OAuth2TokenIntrospection)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockOAuth2APIRecorder) IntrospectOAuth2Token(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IntrospectOAuth2Token", arg0, arg1)
}

func (_m *MockOAuth2API) ListOAuth2Clients() ([]swagger.OAuth2Client, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "ListOAuth2Clients")
	ret0, _ := ret[0].([]swagger.OAuth2Client)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockOAuth2APIRecorder) ListOAuth2Clients() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListOAuth2Clients")
}

func (_m *MockOAuth2API) RejectOAuth2ConsentRequest(id string, body swagger.ConsentRequestRejection) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "RejectOAuth2ConsentRequest", id, body)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOAuth2APIRecorder) RejectOAuth2ConsentRequest(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RejectOAuth2ConsentRequest", arg0, arg1)
}

func (_m *MockOAuth2API) RevokeOAuth2Token(token string) (*swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "RevokeOAuth2Token", token)
	ret0, _ := ret[0].(*swagger.APIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOAuth2APIRecorder) RevokeOAuth2Token(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RevokeOAuth2Token", arg0)
}

func (_m *MockOAuth2API) UpdateOAuth2Client(id string, body swagger.OAuth2Client) (*swagger.OAuth2Client, *swagger.APIResponse, error) {
	ret := _m.ctrl.Call(_m, "UpdateOAuth2Client", id, body)
	ret0, _ := ret[0].(*swagger.OAuth2Client)
	ret1, _ := ret[1].(*swagger.APIResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockOAuth2APIRecorder) UpdateOAuth2Client(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateOAuth2Client", arg0, arg1)
}
