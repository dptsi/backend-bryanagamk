package oidc

import (
	"errors"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"its.ac.id/base-go/pkg/session"
)

var (
	ErrNoEndSessionEndpoint = errors.New("no end session endpoint configured. Please add OIDC_END_SESSION_ENDPOINT to your .env file")
)

const (
	StateKey    = "oidc.state"
	IdTokenKey  = "oidc.id_token"
	StateMaxAge = 60 * 5 // 5 minutes

	AuthorizationCodeNotFound = "authorization_code_not_found"
	InvalidState              = "invalid_state"
	InvalidIdToken            = "invalid_id_token"
	ErrorRetrieveUserInfo     = "error_retrieve_user_info"
)

type QueryParamsProvider interface {
	GetQuery(key string) (string, bool)
}

type Client struct {
	p           *oidc.Provider
	sess        *session.Data
	qp          QueryParamsProvider
	ctx         *gin.Context
	verifyState bool
}

func NewClient(ctx *gin.Context, pUrl string, sess *session.Data, qp QueryParamsProvider) (*Client, error) {
	provider, err := oidc.NewProvider(ctx, pUrl)
	if err != nil {
		return nil, err
	}

	return &Client{provider, sess, qp, ctx, true}, nil
}

func (c *Client) SetVerifyState(verifyState bool) {
	c.verifyState = verifyState
}

func (c *Client) RedirectURL(clientID string, clientSecret string, redirectURL string, scopes []string) string {
	cfg := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,

		// Discovery returns the OAuth2 endpoints.
		Endpoint: c.p.Endpoint(),

		Scopes: scopes,
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	state := uuid.NewString()
	c.sess.Set(StateKey, state)
	c.sess.Save()
	return cfg.AuthCodeURL(state)
}

func (c *Client) UserInfo(clientID string, clientSecret string, redirectURL string, scopes []string) (*oidc.UserInfo, error) {
	code, exist := c.qp.GetQuery("code")
	if !exist {
		return nil, errors.New(AuthorizationCodeNotFound)
	}
	if c.verifyState {
		state, exist := c.qp.GetQuery("state")
		cookieState, ok := c.sess.Get(StateKey)
		if !ok {
			cookieState = ""
		}
		c.sess.Delete(StateKey)
		c.sess.Save()

		if state == "" || !exist || state != cookieState {
			return nil, errors.New(InvalidState)
		}
	}

	cfg := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,

		// Discovery returns the OAuth2 endpoints.
		Endpoint: c.p.Endpoint(),

		Scopes: scopes,
	}

	token, err := cfg.Exchange(c.ctx, code)
	if err != nil {
		return nil, err
	}
	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New(InvalidIdToken)
	}

	// Parse and verify ID Token payload.
	var verifier = c.p.Verifier(&oidc.Config{ClientID: clientID})
	_, err = verifier.Verify(c.ctx, rawIDToken)
	if err != nil {
		return nil, errors.New(InvalidIdToken)
	}
	c.sess.Set(IdTokenKey, rawIDToken)
	c.sess.Save()
	userInfo, err := c.p.UserInfo(c.ctx, oauth2.StaticTokenSource(token))
	if err != nil {
		return nil, errors.New(ErrorRetrieveUserInfo)
	}

	return userInfo, nil
}

func (c *Client) RPInitiatedLogout(endSessionEndpoint string, postLogoutRedirectURI string) (string, error) {
	if endSessionEndpoint == "" {
		return "", ErrNoEndSessionEndpoint
	}
	req, err := http.NewRequest("GET", endSessionEndpoint, nil)
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	idTokenHintItf, exists := c.sess.Get(IdTokenKey)
	if idTokenHint, ok := idTokenHintItf.(string); exists && ok && idTokenHint != "" {
		q.Add("id_token_hint", idTokenHint)
	}
	if postLogoutRedirectURI != "" {
		q.Add("post_logout_redirect_uri", postLogoutRedirectURI)
	}

	req.URL.RawQuery = q.Encode()
	return req.URL.String(), nil
}
