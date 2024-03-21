package lib

import (
	"errors"
	"strings"

	"github.com/abhishekghoshhh/gms/pkg/client"
	"github.com/abhishekghoshhh/gms/pkg/model"
)

const (
	BEARER_TOKEN_PREFIX = "Bearer "
)

type TokenFlow struct {
	iamClient *client.IamClient
}

func NewTokenFlow(iamClient *client.IamClient) *TokenFlow {
	return &TokenFlow{
		iamClient,
	}
}

func (flow *TokenFlow) GetGroups(gmsModel *model.GmsModel) (string, error) {
	token := gmsModel.Token()
	if !strings.HasPrefix(token, BEARER_TOKEN_PREFIX) {
		return "", errors.New("Invalid token " + token)
	}
	iamProfile, err := flow.iamClient.FetchUser(token)
	if err != nil {
		return "", err
	}
	var groups strings.Builder
	for _, group := range iamProfile.Groups {
		groups.WriteString(group.Display)
		groups.WriteString("\n")
	}
	return groups.String(), nil
}

type ClientCredentialFlow struct {
}

func NewClientCredentialFlow() *ClientCredentialFlow {
	return &ClientCredentialFlow{}
}

func (flow *ClientCredentialFlow) GetGroups(gmsModel *model.GmsModel) (string, error) {
	return "group1\ngroup2", nil
}

type PasswordGrantFlow struct {
	cfg *model.PasswordGrantFlowConfig
}

func NewPasswordGrantFlow(cfg *model.PasswordGrantFlowConfig) *PasswordGrantFlow {
	return &PasswordGrantFlow{
		cfg,
	}
}

func (flow *PasswordGrantFlow) GetGroups(gmsModel *model.GmsModel) (string, error) {
	return "group1\ngroup2", nil
}
