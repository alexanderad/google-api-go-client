// Copyright 2020 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package sts provides access to the Security Token Service API.
//
// For product documentation, see: http://cloud.google.com/iam/docs/workload-identity-federation
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/sts/v1"
//   ...
//   ctx := context.Background()
//   stsService, err := sts.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   stsService, err := sts.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   stsService, err := sts.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package sts // import "google.golang.org/api/sts/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "sts:v1"
const apiName = "sts"
const apiVersion = "v1"
const basePath = "https://sts.googleapis.com/"
const mtlsBasePath = "https://sts.mtls.googleapis.com/"

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.V1 = NewV1Service(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	V1 *V1Service
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewV1Service(s *Service) *V1Service {
	rs := &V1Service{s: s}
	return rs
}

type V1Service struct {
	s *Service
}

// GoogleIdentityStsV1ExchangeTokenRequest: Request message for
// ExchangeToken.
type GoogleIdentityStsV1ExchangeTokenRequest struct {
	// GrantType: Required. The grant type. Must be
	// `urn:ietf:params:oauth:grant-type:token-exchange`, which indicates a
	// token exchange.
	GrantType string `json:"grantType,omitempty"`

	// Options: A set of features that Security Token Service supports, in
	// addition to the standard OAuth 2.0 token exchange, formatted as a
	// serialized JSON object of Options.
	Options string `json:"options,omitempty"`

	// RequestedTokenType: Required. An identifier for the type of requested
	// security token. Must be
	// `urn:ietf:params:oauth:token-type:access_token`.
	RequestedTokenType string `json:"requestedTokenType,omitempty"`

	// SubjectToken: Required. The input token. You can use a Google-issued
	// OAuth 2.0 access token with this field to obtain an access token with
	// new security attributes applied, such as a Credential Access
	// Boundary. If an access token already contains security attributes,
	// you cannot apply additional security attributes.
	SubjectToken string `json:"subjectToken,omitempty"`

	// SubjectTokenType: Required. An identifier that indicates the type of
	// the security token in the `subject_token` parameter. Must be
	// `urn:ietf:params:oauth:token-type:access_token`.
	SubjectTokenType string `json:"subjectTokenType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GrantType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GrantType") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIdentityStsV1ExchangeTokenRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIdentityStsV1ExchangeTokenRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIdentityStsV1ExchangeTokenResponse: Response message for
// ExchangeToken.
type GoogleIdentityStsV1ExchangeTokenResponse struct {
	// AccessToken: An OAuth 2.0 security token, issued by Google, in
	// response to the token exchange request.
	AccessToken string `json:"access_token,omitempty"`

	// ExpiresIn: The amount of time, in seconds, between the time when the
	// `access_token` was issued and the time when the `access_token` will
	// expire. This field is absent when the `subject_token` in the request
	// is a Google-issued, short-lived access token. In this case, the
	// `access_token` has the same expiration time as the `subject_token`.
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// IssuedTokenType: The token type. Always matches the value of
	// `requested_token_type` from the request.
	IssuedTokenType string `json:"issued_token_type,omitempty"`

	// TokenType: The type of `access_token`. Always has the value `Bearer`.
	TokenType string `json:"token_type,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AccessToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AccessToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIdentityStsV1ExchangeTokenResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIdentityStsV1ExchangeTokenResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "sts.token":

type V1TokenCall struct {
	s                                       *Service
	googleidentitystsv1exchangetokenrequest *GoogleIdentityStsV1ExchangeTokenRequest
	urlParams_                              gensupport.URLParams
	ctx_                                    context.Context
	header_                                 http.Header
}

// Token: Exchanges a credential for a Google OAuth 2.0 access token.
func (r *V1Service) Token(googleidentitystsv1exchangetokenrequest *GoogleIdentityStsV1ExchangeTokenRequest) *V1TokenCall {
	c := &V1TokenCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.googleidentitystsv1exchangetokenrequest = googleidentitystsv1exchangetokenrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1TokenCall) Fields(s ...googleapi.Field) *V1TokenCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V1TokenCall) Context(ctx context.Context) *V1TokenCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V1TokenCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V1TokenCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20201031")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleidentitystsv1exchangetokenrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/token")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "sts.token" call.
// Exactly one of *GoogleIdentityStsV1ExchangeTokenResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleIdentityStsV1ExchangeTokenResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *V1TokenCall) Do(opts ...googleapi.CallOption) (*GoogleIdentityStsV1ExchangeTokenResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleIdentityStsV1ExchangeTokenResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Exchanges a credential for a Google OAuth 2.0 access token.",
	//   "flatPath": "v1/token",
	//   "httpMethod": "POST",
	//   "id": "sts.token",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/token",
	//   "request": {
	//     "$ref": "GoogleIdentityStsV1ExchangeTokenRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleIdentityStsV1ExchangeTokenResponse"
	//   }
	// }

}
