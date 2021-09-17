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

// Package constants contains shared constants between the client and the server.
package constants

// AttestationPrefix is the protocol-defined prefix for finalizing attestations.
const AttestationPrefix = "TLSAttestationV1"

// EndSessionString gets session-encrypted and sent in an EndSession request.
const EndSessionString = "TLS Tunneled EndSessionRequest V1"

// ExportLabel is the unique label for exporting key material from the TLS session.
const ExportLabel = "EXPERIMENTAL Google Confidential Computing Client Attestation 1.0"

// GrpcPort is the default gRPC server port.
const GrpcPort = 9754

// HTTPPort is the default listening port for the HTTP to gRPC proxy.
const HTTPPort = 9755

// SrvTestCrt test server cert
const SrvTestCrt = "-----BEGIN CERTIFICATE-----\nMIIF7TCCA9WgAwIBAgIUfO5QRVRGKZV5leHWzf80ax+I3NMwDQYJKoZIhvcNAQEL\nBQAweTELMAkGA1UEBhMCR0IxEDAOBgNVBAgMB0VuZ2xhbmQxEjAQBgNVBAcMCUNh\nbWJyaWRnZTEeMBwGA1UECgwVVGhhbGVzIENsb3VkIFNlY3VyaXR5MQwwCgYDVQQL\nDANDUEwxFjAUBgNVBAMMDVRlc3RFMkVTZXJ2ZXIwHhcNMjEwNDI5MTIzODM1WhcN\nMjIwNDI5MTIzODM1WjB5MQswCQYDVQQGEwJHQjEQMA4GA1UECAwHRW5nbGFuZDES\nMBAGA1UEBwwJQ2FtYnJpZGdlMR4wHAYDVQQKDBVUaGFsZXMgQ2xvdWQgU2VjdXJp\ndHkxDDAKBgNVBAsMA0NQTDEWMBQGA1UEAwwNVGVzdEUyRVNlcnZlcjCCAiIwDQYJ\nKoZIhvcNAQEBBQADggIPADCCAgoCggIBALFVWMo7vk9+s+NTEJv7+x5PFVDW4F2v\nruSkhIjul3hR13Wk8Wu3fGskPj8mlHhlEkhvML+acjB8Zc5eslTee1z4G6E6gCzk\nH04csIGUyk9Fig6OMFWNnLuCKPeVEkj/cBNTgoTMiYfHYfPWgO5lIu4aquqrlfJP\nUN+yWkXp0iiULuvoaMrI4C0ClLEdLwzgt++8R/s3fDjIzfUz5O/dGl3TwvoscoUD\nHUpL14uWkASO9JJMUkQHzz4Un1PVrIpvQwFEF8ictCaLCxEUS9HwkA+lfv0RHZP7\nxI9Clvto7pWBjs3hFcx+Pe1J6bi23sFkAKq4oIoLe9xZYMLmocU5PYYG65RsOW1t\nFP6boyO4Op+khTc8xTixUQnl9ohW4CG+o8DtkQZybjJwDI8sml5U/l8CjNQqiwy6\n7g9dyd/FEVX3G2IyyXNjU5vCznah/QVXn8ShlXPPsrpyS3LDzNXiycw1DQoHbGk8\nbabssarBsDlq7Qe7JwZfI+Dki1LzThY/yMTX5KMmjNY6BDYTLFVn/L5S5My/8N/w\nnDfJV5Q0MoajSJ5OyL5+SmUyBbn1QlmK1Fk3UXIWG3XoErvGMln6OaJ0UOKZchpI\n2QSXK2W+fcAKqTscUbjtkGc8RzHOVJ+zb3/3CGJ1Xyj9nSN+dipZXUD/kwG90QKa\nBAlP3I/e4qeBAgMBAAGjbTBrMB0GA1UdDgQWBBTmjyXtiQQ2Umk44iUh1AKSfdkQ\noDAfBgNVHSMEGDAWgBTmjyXtiQQ2Umk44iUh1AKSfdkQoDAPBgNVHRMBAf8EBTAD\nAQH/MBgGA1UdEQQRMA+CDVRlc3RFMkVTZXJ2ZXIwDQYJKoZIhvcNAQELBQADggIB\nAHfpxny8Wpof0bRnW0Ws9C6tn6o1aoIeIogGizn+ywH6mwFkEzR/AnTAlzY8shnS\nf+2WcbYrsySvPZqv+ZhVXsNoZbhUhIxSry4ta/jrfFKMJYBF2WNYWBJJZ5+S+yuH\ntg8FxXRaO8H9Vllm21DFyZznKJBLpVgoAE8YUBttxTQvxjXspj+anHR2efmHOr4u\nsq4tcfcbrTXZz2S/pMq6PzAdXAG7yDaEiKxPNvmTYt/XxrBAye4EHO6wpm6yLXgp\nBEbMVb4hyoHa3fdT5wWNmee3yr6gzDmBGe6/tP9FHolwyF2GIxVCO32RfbtsxYNF\n/X7r1Y8mo6TXUwpaSDdVAm5HU21BCXl8D8SDB1M3W8XotEfwT41Ch7eYkUlozlRc\n/I/83N9PeupWRah042U0HZQHSSiAKuJp3wamefx/QwMkMuH7eM1LPPSk/d4MjJzV\nY2YnNJ3PQNlVLUTXPB0toUXfgJNSNYvthnB24CrNQB8ccbPqskdE6L8aO3CER6h6\nQmGF1EuKtEqO9dcZP/u3eNqGxeIdrA/xZGvxtImeYVpfyDB5EgM24zr8zT1m2QY7\ns/jnPF6cNdFCWUXq38NjLsYAmIcY7jXgpeo4w+hh+/LVpvFc6S1A6BcNZ7UlPLHe\nRiAy+L6HyXCAo72QfYHfieevyQe7++UvzPrKgkGxSTOr\n-----END CERTIFICATE-----\n"

// SrvTestKey test private key for the cert
const SrvTestKey = "-----BEGIN PRIVATE KEY-----\n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\n46OjKS20kZR39lh7R68FCkb3hn3w0A==\n-----END PRIVATE KEY-----\n"
