package mocks

import ecr "github.com/aws/aws-sdk-go/service/ecr"

import mock "github.com/stretchr/testify/mock"
import request "github.com/aws/aws-sdk-go/aws/request"

// ECRAPI is an autogenerated mock type for the ECRAPI type
type ECRAPI struct {
	mock.Mock
}

// BatchCheckLayerAvailability provides a mock function with given fields: _a0
func (_m *ECRAPI) BatchCheckLayerAvailability(_a0 *ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.BatchCheckLayerAvailabilityOutput
	if rf, ok := ret.Get(0).(func(*ecr.BatchCheckLayerAvailabilityInput) *ecr.BatchCheckLayerAvailabilityOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.BatchCheckLayerAvailabilityOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.BatchCheckLayerAvailabilityInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchCheckLayerAvailabilityRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) BatchCheckLayerAvailabilityRequest(_a0 *ecr.BatchCheckLayerAvailabilityInput) (*request.Request, *ecr.BatchCheckLayerAvailabilityOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.BatchCheckLayerAvailabilityInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.BatchCheckLayerAvailabilityOutput
	if rf, ok := ret.Get(1).(func(*ecr.BatchCheckLayerAvailabilityInput) *ecr.BatchCheckLayerAvailabilityOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.BatchCheckLayerAvailabilityOutput)
		}
	}

	return r0, r1
}

// BatchDeleteImage provides a mock function with given fields: _a0
func (_m *ECRAPI) BatchDeleteImage(_a0 *ecr.BatchDeleteImageInput) (*ecr.BatchDeleteImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.BatchDeleteImageOutput
	if rf, ok := ret.Get(0).(func(*ecr.BatchDeleteImageInput) *ecr.BatchDeleteImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.BatchDeleteImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.BatchDeleteImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchDeleteImageRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) BatchDeleteImageRequest(_a0 *ecr.BatchDeleteImageInput) (*request.Request, *ecr.BatchDeleteImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.BatchDeleteImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.BatchDeleteImageOutput
	if rf, ok := ret.Get(1).(func(*ecr.BatchDeleteImageInput) *ecr.BatchDeleteImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.BatchDeleteImageOutput)
		}
	}

	return r0, r1
}

// BatchGetImage provides a mock function with given fields: _a0
func (_m *ECRAPI) BatchGetImage(_a0 *ecr.BatchGetImageInput) (*ecr.BatchGetImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.BatchGetImageOutput
	if rf, ok := ret.Get(0).(func(*ecr.BatchGetImageInput) *ecr.BatchGetImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.BatchGetImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.BatchGetImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchGetImageRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) BatchGetImageRequest(_a0 *ecr.BatchGetImageInput) (*request.Request, *ecr.BatchGetImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.BatchGetImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.BatchGetImageOutput
	if rf, ok := ret.Get(1).(func(*ecr.BatchGetImageInput) *ecr.BatchGetImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.BatchGetImageOutput)
		}
	}

	return r0, r1
}

// CompleteLayerUpload provides a mock function with given fields: _a0
func (_m *ECRAPI) CompleteLayerUpload(_a0 *ecr.CompleteLayerUploadInput) (*ecr.CompleteLayerUploadOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.CompleteLayerUploadOutput
	if rf, ok := ret.Get(0).(func(*ecr.CompleteLayerUploadInput) *ecr.CompleteLayerUploadOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.CompleteLayerUploadOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.CompleteLayerUploadInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompleteLayerUploadRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) CompleteLayerUploadRequest(_a0 *ecr.CompleteLayerUploadInput) (*request.Request, *ecr.CompleteLayerUploadOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.CompleteLayerUploadInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.CompleteLayerUploadOutput
	if rf, ok := ret.Get(1).(func(*ecr.CompleteLayerUploadInput) *ecr.CompleteLayerUploadOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.CompleteLayerUploadOutput)
		}
	}

	return r0, r1
}

// CreateRepository provides a mock function with given fields: _a0
func (_m *ECRAPI) CreateRepository(_a0 *ecr.CreateRepositoryInput) (*ecr.CreateRepositoryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.CreateRepositoryOutput
	if rf, ok := ret.Get(0).(func(*ecr.CreateRepositoryInput) *ecr.CreateRepositoryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.CreateRepositoryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.CreateRepositoryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRepositoryRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) CreateRepositoryRequest(_a0 *ecr.CreateRepositoryInput) (*request.Request, *ecr.CreateRepositoryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.CreateRepositoryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.CreateRepositoryOutput
	if rf, ok := ret.Get(1).(func(*ecr.CreateRepositoryInput) *ecr.CreateRepositoryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.CreateRepositoryOutput)
		}
	}

	return r0, r1
}

// DeleteRepository provides a mock function with given fields: _a0
func (_m *ECRAPI) DeleteRepository(_a0 *ecr.DeleteRepositoryInput) (*ecr.DeleteRepositoryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.DeleteRepositoryOutput
	if rf, ok := ret.Get(0).(func(*ecr.DeleteRepositoryInput) *ecr.DeleteRepositoryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.DeleteRepositoryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.DeleteRepositoryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRepositoryPolicy provides a mock function with given fields: _a0
func (_m *ECRAPI) DeleteRepositoryPolicy(_a0 *ecr.DeleteRepositoryPolicyInput) (*ecr.DeleteRepositoryPolicyOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.DeleteRepositoryPolicyOutput
	if rf, ok := ret.Get(0).(func(*ecr.DeleteRepositoryPolicyInput) *ecr.DeleteRepositoryPolicyOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.DeleteRepositoryPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.DeleteRepositoryPolicyInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRepositoryPolicyRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) DeleteRepositoryPolicyRequest(_a0 *ecr.DeleteRepositoryPolicyInput) (*request.Request, *ecr.DeleteRepositoryPolicyOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.DeleteRepositoryPolicyInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.DeleteRepositoryPolicyOutput
	if rf, ok := ret.Get(1).(func(*ecr.DeleteRepositoryPolicyInput) *ecr.DeleteRepositoryPolicyOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.DeleteRepositoryPolicyOutput)
		}
	}

	return r0, r1
}

// DeleteRepositoryRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) DeleteRepositoryRequest(_a0 *ecr.DeleteRepositoryInput) (*request.Request, *ecr.DeleteRepositoryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.DeleteRepositoryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.DeleteRepositoryOutput
	if rf, ok := ret.Get(1).(func(*ecr.DeleteRepositoryInput) *ecr.DeleteRepositoryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.DeleteRepositoryOutput)
		}
	}

	return r0, r1
}

// DescribeRepositories provides a mock function with given fields: _a0
func (_m *ECRAPI) DescribeRepositories(_a0 *ecr.DescribeRepositoriesInput) (*ecr.DescribeRepositoriesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.DescribeRepositoriesOutput
	if rf, ok := ret.Get(0).(func(*ecr.DescribeRepositoriesInput) *ecr.DescribeRepositoriesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.DescribeRepositoriesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.DescribeRepositoriesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeRepositoriesRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) DescribeRepositoriesRequest(_a0 *ecr.DescribeRepositoriesInput) (*request.Request, *ecr.DescribeRepositoriesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.DescribeRepositoriesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.DescribeRepositoriesOutput
	if rf, ok := ret.Get(1).(func(*ecr.DescribeRepositoriesInput) *ecr.DescribeRepositoriesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.DescribeRepositoriesOutput)
		}
	}

	return r0, r1
}

// GetAuthorizationToken provides a mock function with given fields: _a0
func (_m *ECRAPI) GetAuthorizationToken(_a0 *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.GetAuthorizationTokenOutput
	if rf, ok := ret.Get(0).(func(*ecr.GetAuthorizationTokenInput) *ecr.GetAuthorizationTokenOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.GetAuthorizationTokenOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.GetAuthorizationTokenInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAuthorizationTokenRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) GetAuthorizationTokenRequest(_a0 *ecr.GetAuthorizationTokenInput) (*request.Request, *ecr.GetAuthorizationTokenOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.GetAuthorizationTokenInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.GetAuthorizationTokenOutput
	if rf, ok := ret.Get(1).(func(*ecr.GetAuthorizationTokenInput) *ecr.GetAuthorizationTokenOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.GetAuthorizationTokenOutput)
		}
	}

	return r0, r1
}

// GetDownloadUrlForLayer provides a mock function with given fields: _a0
func (_m *ECRAPI) GetDownloadUrlForLayer(_a0 *ecr.GetDownloadUrlForLayerInput) (*ecr.GetDownloadUrlForLayerOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.GetDownloadUrlForLayerOutput
	if rf, ok := ret.Get(0).(func(*ecr.GetDownloadUrlForLayerInput) *ecr.GetDownloadUrlForLayerOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.GetDownloadUrlForLayerOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.GetDownloadUrlForLayerInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDownloadUrlForLayerRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) GetDownloadUrlForLayerRequest(_a0 *ecr.GetDownloadUrlForLayerInput) (*request.Request, *ecr.GetDownloadUrlForLayerOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.GetDownloadUrlForLayerInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.GetDownloadUrlForLayerOutput
	if rf, ok := ret.Get(1).(func(*ecr.GetDownloadUrlForLayerInput) *ecr.GetDownloadUrlForLayerOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.GetDownloadUrlForLayerOutput)
		}
	}

	return r0, r1
}

// GetRepositoryPolicy provides a mock function with given fields: _a0
func (_m *ECRAPI) GetRepositoryPolicy(_a0 *ecr.GetRepositoryPolicyInput) (*ecr.GetRepositoryPolicyOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.GetRepositoryPolicyOutput
	if rf, ok := ret.Get(0).(func(*ecr.GetRepositoryPolicyInput) *ecr.GetRepositoryPolicyOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.GetRepositoryPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.GetRepositoryPolicyInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepositoryPolicyRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) GetRepositoryPolicyRequest(_a0 *ecr.GetRepositoryPolicyInput) (*request.Request, *ecr.GetRepositoryPolicyOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.GetRepositoryPolicyInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.GetRepositoryPolicyOutput
	if rf, ok := ret.Get(1).(func(*ecr.GetRepositoryPolicyInput) *ecr.GetRepositoryPolicyOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.GetRepositoryPolicyOutput)
		}
	}

	return r0, r1
}

// InitiateLayerUpload provides a mock function with given fields: _a0
func (_m *ECRAPI) InitiateLayerUpload(_a0 *ecr.InitiateLayerUploadInput) (*ecr.InitiateLayerUploadOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.InitiateLayerUploadOutput
	if rf, ok := ret.Get(0).(func(*ecr.InitiateLayerUploadInput) *ecr.InitiateLayerUploadOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.InitiateLayerUploadOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.InitiateLayerUploadInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitiateLayerUploadRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) InitiateLayerUploadRequest(_a0 *ecr.InitiateLayerUploadInput) (*request.Request, *ecr.InitiateLayerUploadOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.InitiateLayerUploadInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.InitiateLayerUploadOutput
	if rf, ok := ret.Get(1).(func(*ecr.InitiateLayerUploadInput) *ecr.InitiateLayerUploadOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.InitiateLayerUploadOutput)
		}
	}

	return r0, r1
}

// ListImages provides a mock function with given fields: _a0
func (_m *ECRAPI) ListImages(_a0 *ecr.ListImagesInput) (*ecr.ListImagesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.ListImagesOutput
	if rf, ok := ret.Get(0).(func(*ecr.ListImagesInput) *ecr.ListImagesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.ListImagesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.ListImagesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListImagesRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) ListImagesRequest(_a0 *ecr.ListImagesInput) (*request.Request, *ecr.ListImagesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.ListImagesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.ListImagesOutput
	if rf, ok := ret.Get(1).(func(*ecr.ListImagesInput) *ecr.ListImagesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.ListImagesOutput)
		}
	}

	return r0, r1
}

// PutImage provides a mock function with given fields: _a0
func (_m *ECRAPI) PutImage(_a0 *ecr.PutImageInput) (*ecr.PutImageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.PutImageOutput
	if rf, ok := ret.Get(0).(func(*ecr.PutImageInput) *ecr.PutImageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.PutImageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.PutImageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutImageRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) PutImageRequest(_a0 *ecr.PutImageInput) (*request.Request, *ecr.PutImageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.PutImageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.PutImageOutput
	if rf, ok := ret.Get(1).(func(*ecr.PutImageInput) *ecr.PutImageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.PutImageOutput)
		}
	}

	return r0, r1
}

// SetRepositoryPolicy provides a mock function with given fields: _a0
func (_m *ECRAPI) SetRepositoryPolicy(_a0 *ecr.SetRepositoryPolicyInput) (*ecr.SetRepositoryPolicyOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.SetRepositoryPolicyOutput
	if rf, ok := ret.Get(0).(func(*ecr.SetRepositoryPolicyInput) *ecr.SetRepositoryPolicyOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.SetRepositoryPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.SetRepositoryPolicyInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetRepositoryPolicyRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) SetRepositoryPolicyRequest(_a0 *ecr.SetRepositoryPolicyInput) (*request.Request, *ecr.SetRepositoryPolicyOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.SetRepositoryPolicyInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.SetRepositoryPolicyOutput
	if rf, ok := ret.Get(1).(func(*ecr.SetRepositoryPolicyInput) *ecr.SetRepositoryPolicyOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.SetRepositoryPolicyOutput)
		}
	}

	return r0, r1
}

// UploadLayerPart provides a mock function with given fields: _a0
func (_m *ECRAPI) UploadLayerPart(_a0 *ecr.UploadLayerPartInput) (*ecr.UploadLayerPartOutput, error) {
	ret := _m.Called(_a0)

	var r0 *ecr.UploadLayerPartOutput
	if rf, ok := ret.Get(0).(func(*ecr.UploadLayerPartInput) *ecr.UploadLayerPartOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecr.UploadLayerPartOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecr.UploadLayerPartInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadLayerPartRequest provides a mock function with given fields: _a0
func (_m *ECRAPI) UploadLayerPartRequest(_a0 *ecr.UploadLayerPartInput) (*request.Request, *ecr.UploadLayerPartOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*ecr.UploadLayerPartInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *ecr.UploadLayerPartOutput
	if rf, ok := ret.Get(1).(func(*ecr.UploadLayerPartInput) *ecr.UploadLayerPartOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ecr.UploadLayerPartOutput)
		}
	}

	return r0, r1
}
