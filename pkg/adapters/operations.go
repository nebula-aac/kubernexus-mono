package adapters

import (
	"context"
	"net/url"

	"github.com/nebula-aac/kubernexus-mono/pkg/utils"
)

var (
	NoneVersion  = []Version{"none"}
	NoneTemplate = []Template{"none"}
)

type Version string

type Template string

type Service string

func (t Template) String() string {
	_, err := url.ParseRequestURI(string(t))
	if err != nil {
		return string(t)
	}

	st, err := utils.ReadFileSource(string(t))
	if err != nil {
		return ""
	}

	return st
}

type Operation struct {
	Type                 int32             `json:"type,string,omitempty"`
	Description          string            `json:"description,omitempty"`
	Versions             []Version         `json:"versions,omitempty"`
	Templates            []Template        `json:"templates,omitempty"`
	Services             []Service         `json:"services,omitempty"`
	AdditionalProperties map[string]string `json:"additional_properties,omitempty"`
}

type Operations map[string]*Operation

type OperationRequest struct {
	OperationName     string
	Namespace         string
	Username          string
	CustomBody        string
	IsDeleteOperation bool
	OperationID       string
	K8sConfigs        []string
}

func (h *Adapter) ApplyOperation(ctx context.Context, op OperationRequest) error {
	return nil
}
