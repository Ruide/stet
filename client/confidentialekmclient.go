// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Defines an HTTP client for making requests to Confidential EKM services.

package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	cwpb "github.com/GoogleCloudPlatform/stet/proto/confidential_wrap_go_proto"
	sspb "github.com/GoogleCloudPlatform/stet/proto/secure_session_go_proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const (
	beginSessionEndpoint         = "/session/beginsession"
	handshakeEndpoint            = "/session/handshake"
	negotiateAttestationEndpoint = "/session/negotiateattestation"
	finalizeEndpoint             = "/session/finalize"
	endSessionEndpoint           = "/session/endsession"
	confidentialWrapEndpoint     = ":confidentialwrap"
	confidentialUnwrapEndpoint   = ":confidentialunwrap"
)

type confidentialEkmClient struct {
	uri       string
	authToken string
	certPool  *x509.CertPool
}

// Removes the last two path component from the key URI.
func removeEndpointPathComponent(url string) string {
	for i := 0; i < 2; i++ {
		substrIndex := strings.LastIndex(url, "/")
		if substrIndex == -1 {
			break
		}
		url = url[0:substrIndex]
	}

	return url
}

func (c confidentialEkmClient) post(ctx context.Context, url string, protoReq, protoResp proto.Message) error {
	marshaled, err := protojson.Marshal(protoReq)
	if err != nil {
		return fmt.Errorf("Error marshaling request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(marshaled))
	if err != nil {
		return fmt.Errorf("Error creating HTTP request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")

	if c.authToken != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.authToken)
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: c.certPool,
			},
		},
	}

	httpResp, err := client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("HTTP call returned with error: %w", err)
	}
	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("Non-OK status returned: %s", httpResp.Status)
	}

	defer httpResp.Body.Close()
	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return fmt.Errorf("Error reading HTTP response body: %w", err)
	}

	if err = protojson.Unmarshal(respBody, protoResp); err != nil {
		return fmt.Errorf("Error unmarshaling response: %w", err)
	}

	return nil
}

func (c confidentialEkmClient) BeginSession(ctx context.Context, req *sspb.BeginSessionRequest) (*sspb.BeginSessionResponse, error) {
	resp := &sspb.BeginSessionResponse{}
	url := removeEndpointPathComponent(c.uri) + beginSessionEndpoint
	if err := c.post(ctx, url, req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c confidentialEkmClient) Handshake(ctx context.Context, req *sspb.HandshakeRequest) (*sspb.HandshakeResponse, error) {
	resp := &sspb.HandshakeResponse{}
	url := removeEndpointPathComponent(c.uri) + handshakeEndpoint
	if err := c.post(ctx, url, req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c confidentialEkmClient) NegotiateAttestation(ctx context.Context, req *sspb.NegotiateAttestationRequest) (*sspb.NegotiateAttestationResponse, error) {
	resp := &sspb.NegotiateAttestationResponse{}
	url := removeEndpointPathComponent(c.uri) + negotiateAttestationEndpoint
	if err := c.post(ctx, url, req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c confidentialEkmClient) Finalize(ctx context.Context, req *sspb.FinalizeRequest) (*sspb.FinalizeResponse, error) {
	resp := &sspb.FinalizeResponse{}
	url := removeEndpointPathComponent(c.uri) + finalizeEndpoint
	if err := c.post(ctx, url, req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c confidentialEkmClient) EndSession(ctx context.Context, req *sspb.EndSessionRequest) (*sspb.EndSessionResponse, error) {
	resp := &sspb.EndSessionResponse{}
	url := removeEndpointPathComponent(c.uri) + endSessionEndpoint
	if err := c.post(ctx, url, req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c confidentialEkmClient) ConfidentialWrap(ctx context.Context, req *cwpb.ConfidentialWrapRequest) (*cwpb.ConfidentialWrapResponse, error) {
	resp := &cwpb.ConfidentialWrapResponse{}
	url := c.uri + confidentialWrapEndpoint
	if err := c.post(ctx, url, req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c confidentialEkmClient) ConfidentialUnwrap(ctx context.Context, req *cwpb.ConfidentialUnwrapRequest) (*cwpb.ConfidentialUnwrapResponse, error) {
	resp := &cwpb.ConfidentialUnwrapResponse{}
	url := c.uri + confidentialUnwrapEndpoint
	if err := c.post(ctx, url, req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
