// Copyright (c) Illumio, Inc.
// SPDX-License-Identifier: MPL-2.0
// Code generated by fakeservergen. DO NOT EDIT.

package main

import (
	"context"
	"sync"

	"github.com/google/uuid"
	configv1 "github.com/illumio/terraform-provider-illumio-cloudsecure/api/illumio/cloud/config/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// FakeConfigServer is a fake server implementation of ConfigService that can be used for testing API clients.
type FakeConfigServer struct {
	configv1.UnimplementedConfigServiceServer
	Logger                              *zap.Logger
	AwsAccountMap                       map[string]*AwsAccount
	AwsAccountMutex                     sync.RWMutex
	AwsFlowLogsS3BucketMap              map[string]*AwsFlowLogsS3Bucket
	AwsFlowLogsS3BucketMutex            sync.RWMutex
	K8SClusterOnboardingCredentialMap   map[string]*K8SClusterOnboardingCredential
	K8SClusterOnboardingCredentialMutex sync.RWMutex
}

var _ configv1.ConfigServiceServer = &FakeConfigServer{}

// NewFakeConfigServer creates a fake server implementation of ConfigService that can be used for testing API clients.
func NewFakeConfigServer(logger *zap.Logger) configv1.ConfigServiceServer {
	return &FakeConfigServer{
		Logger:                            logger,
		AwsAccountMap:                     make(map[string]*AwsAccount),
		AwsFlowLogsS3BucketMap:            make(map[string]*AwsFlowLogsS3Bucket),
		K8SClusterOnboardingCredentialMap: make(map[string]*K8SClusterOnboardingCredential),
	}
}

type AwsAccount struct {
	Id             string
	AccountId      string
	Mode           string
	Name           string
	OrganizationId *string
	RoleArn        string
	RoleExternalId string
}

type AwsFlowLogsS3Bucket struct {
	Id          string
	AccountId   string
	S3BucketArn string
}

type K8SClusterOnboardingCredential struct {
	Id            string
	ClientId      string
	ClientSecret  string
	CreatedAt     string
	Description   *string
	IllumioRegion string
	Name          string
}

func (s *FakeConfigServer) CreateAwsAccount(ctx context.Context, req *configv1.CreateAwsAccountRequest) (*configv1.CreateAwsAccountResponse, error) {
	id := uuid.New().String()
	model := &AwsAccount{
		Id:             id,
		AccountId:      req.AccountId,
		Mode:           req.Mode,
		Name:           req.Name,
		OrganizationId: req.OrganizationId,
		RoleArn:        req.RoleArn,
		RoleExternalId: req.RoleExternalId,
	}
	resp := &configv1.CreateAwsAccountResponse{
		Id:             id,
		AccountId:      model.AccountId,
		Mode:           model.Mode,
		Name:           model.Name,
		OrganizationId: model.OrganizationId,
		RoleArn:        model.RoleArn,
	}
	s.AwsAccountMutex.Lock()
	s.AwsAccountMap[id] = model
	s.AwsAccountMutex.Unlock()
	s.Logger.Info("created resource",
		zap.String("type", "aws_account"),
		zap.String("method", "CreateAwsAccount"),
		zap.String("id", id),
	)
	return resp, nil
}

func (s *FakeConfigServer) ReadAwsAccount(ctx context.Context, req *configv1.ReadAwsAccountRequest) (*configv1.ReadAwsAccountResponse, error) {
	id := req.Id
	s.AwsAccountMutex.RLock()
	model, found := s.AwsAccountMap[id]
	if !found {
		s.AwsAccountMutex.RUnlock()
		s.Logger.Error("attempted to read resource with unknown id",
			zap.String("type", "aws_account"),
			zap.String("method", "ReadAwsAccount"),
			zap.String("id", id),
		)
		return nil, status.Errorf(codes.NotFound, "no aws_account found with id %s", id)
	}
	resp := &configv1.ReadAwsAccountResponse{
		Id:             id,
		AccountId:      model.AccountId,
		Mode:           model.Mode,
		Name:           model.Name,
		OrganizationId: model.OrganizationId,
		RoleArn:        model.RoleArn,
	}
	s.AwsAccountMutex.RUnlock()
	s.Logger.Info("read resource",
		zap.String("type", "aws_account"),
		zap.String("method", "ReadAwsAccount"),
		zap.String("id", id),
	)
	return resp, nil
}

func (s *FakeConfigServer) UpdateAwsAccount(ctx context.Context, req *configv1.UpdateAwsAccountRequest) (*configv1.UpdateAwsAccountResponse, error) {
	id := req.Id
	s.AwsAccountMutex.Lock()
	model, found := s.AwsAccountMap[id]
	if !found {
		s.AwsAccountMutex.Unlock()
		s.Logger.Error("attempted to update resource with unknown id",
			zap.String("type", "aws_account"),
			zap.String("method", "UpdateAwsAccount"),
			zap.String("id", id),
		)
		return nil, status.Errorf(codes.NotFound, "no aws_account found with id %s", id)
	}
	updateMask := req.UpdateMask
	var updateMaskPaths []string
	if updateMask != nil {
		updateMaskPaths = updateMask.Paths
	}
	for _, path := range updateMaskPaths {
		switch path {
		case "name":
			model.Name = req.Name
		default:
			s.AwsAccountMutex.Unlock()
			s.Logger.Error("attempted to update resource using invalid update_mask path",
				zap.String("type", "aws_account"),
				zap.String("method", "UpdateAwsAccount"),
				zap.String("id", id),
				zap.Strings("updateMaskPaths", updateMaskPaths),
				zap.String("invalidUpdateMaskPath", path),
			)
			return nil, status.Errorf(codes.InvalidArgument, "invalid path in update_mask for aws_account: %s", path)
		}
	}
	resp := &configv1.UpdateAwsAccountResponse{
		Id:             id,
		AccountId:      model.AccountId,
		Mode:           model.Mode,
		Name:           model.Name,
		OrganizationId: model.OrganizationId,
		RoleArn:        model.RoleArn,
	}
	s.AwsAccountMutex.Unlock()
	s.Logger.Info("updated resource",
		zap.String("type", "aws_account"),
		zap.String("method", "UpdateAwsAccount"),
		zap.String("id", id),
		zap.Strings("updateMaskPaths", updateMaskPaths),
	)
	return resp, nil
}

func (s *FakeConfigServer) DeleteAwsAccount(ctx context.Context, req *configv1.DeleteAwsAccountRequest) (*emptypb.Empty, error) {
	id := req.Id
	s.AwsAccountMutex.Lock()
	_, found := s.AwsAccountMap[id]
	if !found {
		s.AwsAccountMutex.Unlock()
		s.Logger.Error("attempted to delete resource with unknown id",
			zap.String("type", "aws_account"),
			zap.String("method", "DeleteAwsAccount"),
			zap.String("id", id),
		)
		return nil, status.Errorf(codes.NotFound, "no aws_account found with id %s", id)
	}
	delete(s.AwsAccountMap, id)
	s.AwsAccountMutex.Unlock()
	s.Logger.Info("deleted resource",
		zap.String("type", "aws_account"),
		zap.String("method", "DeleteAwsAccount"),
		zap.String("id", id),
	)
	return &emptypb.Empty{}, nil
}
func (s *FakeConfigServer) CreateAwsFlowLogsS3Bucket(ctx context.Context, req *configv1.CreateAwsFlowLogsS3BucketRequest) (*configv1.CreateAwsFlowLogsS3BucketResponse, error) {
	id := uuid.New().String()
	model := &AwsFlowLogsS3Bucket{
		Id:          id,
		AccountId:   req.AccountId,
		S3BucketArn: req.S3BucketArn,
	}
	resp := &configv1.CreateAwsFlowLogsS3BucketResponse{
		Id:          id,
		AccountId:   model.AccountId,
		S3BucketArn: model.S3BucketArn,
	}
	s.AwsFlowLogsS3BucketMutex.Lock()
	s.AwsFlowLogsS3BucketMap[id] = model
	s.AwsFlowLogsS3BucketMutex.Unlock()
	s.Logger.Info("created resource",
		zap.String("type", "aws_flow_logs_s3_bucket"),
		zap.String("method", "CreateAwsFlowLogsS3Bucket"),
		zap.String("id", id),
	)
	return resp, nil
}

func (s *FakeConfigServer) ReadAwsFlowLogsS3Bucket(ctx context.Context, req *configv1.ReadAwsFlowLogsS3BucketRequest) (*configv1.ReadAwsFlowLogsS3BucketResponse, error) {
	id := req.Id
	s.AwsFlowLogsS3BucketMutex.RLock()
	model, found := s.AwsFlowLogsS3BucketMap[id]
	if !found {
		s.AwsFlowLogsS3BucketMutex.RUnlock()
		s.Logger.Error("attempted to read resource with unknown id",
			zap.String("type", "aws_flow_logs_s3_bucket"),
			zap.String("method", "ReadAwsFlowLogsS3Bucket"),
			zap.String("id", id),
		)
		return nil, status.Errorf(codes.NotFound, "no aws_flow_logs_s3_bucket found with id %s", id)
	}
	resp := &configv1.ReadAwsFlowLogsS3BucketResponse{
		Id:          id,
		AccountId:   model.AccountId,
		S3BucketArn: model.S3BucketArn,
	}
	s.AwsFlowLogsS3BucketMutex.RUnlock()
	s.Logger.Info("read resource",
		zap.String("type", "aws_flow_logs_s3_bucket"),
		zap.String("method", "ReadAwsFlowLogsS3Bucket"),
		zap.String("id", id),
	)
	return resp, nil
}

func (s *FakeConfigServer) UpdateAwsFlowLogsS3Bucket(ctx context.Context, req *configv1.UpdateAwsFlowLogsS3BucketRequest) (*configv1.UpdateAwsFlowLogsS3BucketResponse, error) {
	id := req.Id
	s.AwsFlowLogsS3BucketMutex.Lock()
	model, found := s.AwsFlowLogsS3BucketMap[id]
	if !found {
		s.AwsFlowLogsS3BucketMutex.Unlock()
		s.Logger.Error("attempted to update resource with unknown id",
			zap.String("type", "aws_flow_logs_s3_bucket"),
			zap.String("method", "UpdateAwsFlowLogsS3Bucket"),
			zap.String("id", id),
		)
		return nil, status.Errorf(codes.NotFound, "no aws_flow_logs_s3_bucket found with id %s", id)
	}
	updateMask := req.UpdateMask
	var updateMaskPaths []string
	if updateMask != nil {
		updateMaskPaths = updateMask.Paths
	}
	for _, path := range updateMaskPaths {
		switch path {
		default:
			s.AwsAccountMutex.Unlock()
			s.Logger.Error("attempted to update resource using invalid update_mask path",
				zap.String("type", "aws_flow_logs_s3_bucket"),
				zap.String("method", "UpdateAwsFlowLogsS3Bucket"),
				zap.String("id", id),
				zap.Strings("updateMaskPaths", updateMaskPaths),
				zap.String("invalidUpdateMaskPath", path),
			)
			return nil, status.Errorf(codes.InvalidArgument, "invalid path in update_mask for aws_account: %s", path)
		}
	}
	resp := &configv1.UpdateAwsFlowLogsS3BucketResponse{
		Id:          id,
		AccountId:   model.AccountId,
		S3BucketArn: model.S3BucketArn,
	}
	s.AwsFlowLogsS3BucketMutex.Unlock()
	s.Logger.Info("updated resource",
		zap.String("type", "aws_flow_logs_s3_bucket"),
		zap.String("method", "UpdateAwsFlowLogsS3Bucket"),
		zap.String("id", id),
		zap.Strings("updateMaskPaths", updateMaskPaths),
	)
	return resp, nil
}

func (s *FakeConfigServer) DeleteAwsFlowLogsS3Bucket(ctx context.Context, req *configv1.DeleteAwsFlowLogsS3BucketRequest) (*emptypb.Empty, error) {
	id := req.Id
	s.AwsFlowLogsS3BucketMutex.Lock()
	_, found := s.AwsFlowLogsS3BucketMap[id]
	if !found {
		s.AwsFlowLogsS3BucketMutex.Unlock()
		s.Logger.Error("attempted to delete resource with unknown id",
			zap.String("type", "aws_flow_logs_s3_bucket"),
			zap.String("method", "DeleteAwsFlowLogsS3Bucket"),
			zap.String("id", id),
		)
		return nil, status.Errorf(codes.NotFound, "no aws_flow_logs_s3_bucket found with id %s", id)
	}
	delete(s.AwsFlowLogsS3BucketMap, id)
	s.AwsFlowLogsS3BucketMutex.Unlock()
	s.Logger.Info("deleted resource",
		zap.String("type", "aws_flow_logs_s3_bucket"),
		zap.String("method", "DeleteAwsFlowLogsS3Bucket"),
		zap.String("id", id),
	)
	return &emptypb.Empty{}, nil
}
func (s *FakeConfigServer) CreateK8SClusterOnboardingCredential(ctx context.Context, req *configv1.CreateK8SClusterOnboardingCredentialRequest) (*configv1.CreateK8SClusterOnboardingCredentialResponse, error) {
	id := uuid.New().String()
	model := &K8SClusterOnboardingCredential{
		Id:            id,
		Description:   req.Description,
		IllumioRegion: req.IllumioRegion,
		Name:          req.Name,
	}
	resp := &configv1.CreateK8SClusterOnboardingCredentialResponse{
		Id:            id,
		ClientId:      model.ClientId,
		ClientSecret:  model.ClientSecret,
		CreatedAt:     model.CreatedAt,
		Description:   model.Description,
		IllumioRegion: model.IllumioRegion,
		Name:          model.Name,
	}
	s.K8SClusterOnboardingCredentialMutex.Lock()
	s.K8SClusterOnboardingCredentialMap[id] = model
	s.K8SClusterOnboardingCredentialMutex.Unlock()
	s.Logger.Info("created resource",
		zap.String("type", "k8s_cluster_onboarding_credential"),
		zap.String("method", "CreateK8SClusterOnboardingCredential"),
		zap.String("id", id),
	)
	return resp, nil
}

func (s *FakeConfigServer) ReadK8SClusterOnboardingCredential(ctx context.Context, req *configv1.ReadK8SClusterOnboardingCredentialRequest) (*configv1.ReadK8SClusterOnboardingCredentialResponse, error) {
	id := req.Id
	s.K8SClusterOnboardingCredentialMutex.RLock()
	model, found := s.K8SClusterOnboardingCredentialMap[id]
	if !found {
		s.K8SClusterOnboardingCredentialMutex.RUnlock()
		s.Logger.Error("attempted to read resource with unknown id",
			zap.String("type", "k8s_cluster_onboarding_credential"),
			zap.String("method", "ReadK8SClusterOnboardingCredential"),
			zap.String("id", id),
		)
		return nil, status.Errorf(codes.NotFound, "no k8s_cluster_onboarding_credential found with id %s", id)
	}
	resp := &configv1.ReadK8SClusterOnboardingCredentialResponse{
		Id:            id,
		ClientId:      model.ClientId,
		CreatedAt:     model.CreatedAt,
		Description:   model.Description,
		IllumioRegion: model.IllumioRegion,
		Name:          model.Name,
	}
	s.K8SClusterOnboardingCredentialMutex.RUnlock()
	s.Logger.Info("read resource",
		zap.String("type", "k8s_cluster_onboarding_credential"),
		zap.String("method", "ReadK8SClusterOnboardingCredential"),
		zap.String("id", id),
	)
	return resp, nil
}

func (s *FakeConfigServer) UpdateK8SClusterOnboardingCredential(ctx context.Context, req *configv1.UpdateK8SClusterOnboardingCredentialRequest) (*configv1.UpdateK8SClusterOnboardingCredentialResponse, error) {
	id := req.Id
	s.K8SClusterOnboardingCredentialMutex.Lock()
	model, found := s.K8SClusterOnboardingCredentialMap[id]
	if !found {
		s.K8SClusterOnboardingCredentialMutex.Unlock()
		s.Logger.Error("attempted to update resource with unknown id",
			zap.String("type", "k8s_cluster_onboarding_credential"),
			zap.String("method", "UpdateK8SClusterOnboardingCredential"),
			zap.String("id", id),
		)
		return nil, status.Errorf(codes.NotFound, "no k8s_cluster_onboarding_credential found with id %s", id)
	}
	updateMask := req.UpdateMask
	var updateMaskPaths []string
	if updateMask != nil {
		updateMaskPaths = updateMask.Paths
	}
	for _, path := range updateMaskPaths {
		switch path {
		case "description":
			model.Description = req.Description
		case "name":
			model.Name = req.Name
		default:
			s.AwsAccountMutex.Unlock()
			s.Logger.Error("attempted to update resource using invalid update_mask path",
				zap.String("type", "k8s_cluster_onboarding_credential"),
				zap.String("method", "UpdateK8SClusterOnboardingCredential"),
				zap.String("id", id),
				zap.Strings("updateMaskPaths", updateMaskPaths),
				zap.String("invalidUpdateMaskPath", path),
			)
			return nil, status.Errorf(codes.InvalidArgument, "invalid path in update_mask for aws_account: %s", path)
		}
	}
	resp := &configv1.UpdateK8SClusterOnboardingCredentialResponse{
		Id:            id,
		ClientId:      model.ClientId,
		CreatedAt:     model.CreatedAt,
		Description:   model.Description,
		IllumioRegion: model.IllumioRegion,
		Name:          model.Name,
	}
	s.K8SClusterOnboardingCredentialMutex.Unlock()
	s.Logger.Info("updated resource",
		zap.String("type", "k8s_cluster_onboarding_credential"),
		zap.String("method", "UpdateK8SClusterOnboardingCredential"),
		zap.String("id", id),
		zap.Strings("updateMaskPaths", updateMaskPaths),
	)
	return resp, nil
}

func (s *FakeConfigServer) DeleteK8SClusterOnboardingCredential(ctx context.Context, req *configv1.DeleteK8SClusterOnboardingCredentialRequest) (*emptypb.Empty, error) {
	id := req.Id
	s.K8SClusterOnboardingCredentialMutex.Lock()
	_, found := s.K8SClusterOnboardingCredentialMap[id]
	if !found {
		s.K8SClusterOnboardingCredentialMutex.Unlock()
		s.Logger.Error("attempted to delete resource with unknown id",
			zap.String("type", "k8s_cluster_onboarding_credential"),
			zap.String("method", "DeleteK8SClusterOnboardingCredential"),
			zap.String("id", id),
		)
		return nil, status.Errorf(codes.NotFound, "no k8s_cluster_onboarding_credential found with id %s", id)
	}
	delete(s.K8SClusterOnboardingCredentialMap, id)
	s.K8SClusterOnboardingCredentialMutex.Unlock()
	s.Logger.Info("deleted resource",
		zap.String("type", "k8s_cluster_onboarding_credential"),
		zap.String("method", "DeleteK8SClusterOnboardingCredential"),
		zap.String("id", id),
	)
	return &emptypb.Empty{}, nil
}
