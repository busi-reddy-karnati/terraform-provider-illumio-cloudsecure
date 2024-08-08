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
	Logger          *zap.Logger
	AwsAccountMap   map[string]*AwsAccount
	AwsAccountMutex sync.RWMutex
}

var _ configv1.ConfigServiceServer = &FakeConfigServer{}

// NewFakeConfigServer creates a fake server implementation of ConfigService that can be used for testing API clients.
func NewFakeConfigServer(logger *zap.Logger) configv1.ConfigServiceServer {
	return &FakeConfigServer{
		Logger:        logger,
		AwsAccountMap: make(map[string]*AwsAccount),
	}
}

type AwsAccount struct {
	Id                  string
	AccountId           string
	AccountType         string
	ManagementAccountId *string
	Mode                string
	Name                string
	OrganizationId      *string
	RoleArn             string
	RoleExternalId      string
}

func (s *FakeConfigServer) CreateAwsAccount(ctx context.Context, req *configv1.CreateAwsAccountRequest) (*configv1.CreateAwsAccountResponse, error) {
	id := uuid.New().String()
	model := &AwsAccount{
		Id:                  id,
		AccountId:           req.AccountId,
		AccountType:         req.AccountType,
		ManagementAccountId: req.ManagementAccountId,
		Mode:                req.Mode,
		Name:                req.Name,
		OrganizationId:      req.OrganizationId,
		RoleArn:             req.RoleArn,
		RoleExternalId:      req.RoleExternalId,
	}
	resp := &configv1.CreateAwsAccountResponse{
		Id:          id,
		AccountId:   model.AccountId,
		AccountType: model.AccountType,
		Mode:        model.Mode,
		Name:        model.Name,
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
		Id:          id,
		AccountId:   model.AccountId,
		AccountType: model.AccountType,
		Mode:        model.Mode,
		Name:        model.Name,
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
		Id:          id,
		AccountId:   model.AccountId,
		AccountType: model.AccountType,
		Mode:        model.Mode,
		Name:        model.Name,
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
