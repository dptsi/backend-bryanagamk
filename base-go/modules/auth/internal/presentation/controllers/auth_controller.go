package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"its.ac.id/base-go/modules/auth/internal/app/config"
	"its.ac.id/base-go/pkg/auth/contracts"
	"its.ac.id/base-go/pkg/auth/services"
	"its.ac.id/base-go/pkg/oidc"
	"its.ac.id/base-go/pkg/session"
)

type AuthController struct {
	i   *do.Injector
	cfg config.AuthConfig
}

func NewAuthController() *AuthController {
	i := do.DefaultInjector
	cfg := do.MustInvoke[config.AuthConfig](i)

	return &AuthController{i, cfg}
}

func (c *AuthController) Logout(ctx *gin.Context) {
	op, err := c.getOidcClient(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "unable_to_get_oidc_client",
			"data":    nil,
		})
		return
	}
	cfg := c.cfg.Oidc()
	endSessionEndpoint, err := op.RPInitiatedLogout(cfg.EndSessionEndpoint, cfg.PostLogoutRedirectURI)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "unable_to_get_end_session_endpoint",
			"data":    nil,
		})
		return
	}

	err = services.Logout(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "logout_failed",
			"data":    nil,
		})
		return
	}
	sess := session.Default(ctx)
	sess.Invalidate()
	sess.RegenerateCSRFToken()
	if err := sess.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "unable_to_save_session",
			"data":    nil,
		})
		return
	}

	session.AddCookieToResponse(ctx, sess.Id())

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "logout_success",
		"data":    endSessionEndpoint,
	})
}

func (c *AuthController) User(ctx *gin.Context) {
	u := services.User(ctx)
	roles := make([]gin.H, 0)
	for _, r := range u.Roles() {
		roles = append(roles, gin.H{
			"name":        r.Name,
			"permissions": r.Permissions,
			"is_default":  r.IsDefault,
		})
	}
	var activeRole any
	activeRole = nil
	if u.ActiveRole() != "" {
		activeRole = u.ActiveRole()
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "user",
		"data": gin.H{
			"id":          u.Id(),
			"active_role": activeRole,
			"roles":       roles,
		},
	})
}

func (c *AuthController) getOidcClient(ctx *gin.Context) (*oidc.Client, error) {
	cfg := c.cfg.Oidc()
	sess := session.Default(ctx)
	op, err := oidc.NewClient(ctx, cfg.Provider, sess, ctx)
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (c *AuthController) Login(ctx *gin.Context) {
	op, err := c.getOidcClient(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "login_failed",
			"data":    nil,
		})
	}
	cfg := c.cfg.Oidc()
	url := op.RedirectURL(cfg.ClientID, cfg.ClientSecret, cfg.RedirectURL, cfg.Scopes)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "login_url",
		"data":    url,
	})
}

func (c *AuthController) Callback(ctx *gin.Context) {
	op, err := c.getOidcClient(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "login_failed",
			"data":    nil,
		})
		return
	}

	userInfo, err := op.UserInfo(c.cfg.Oidc().ClientID, c.cfg.Oidc().ClientSecret, c.cfg.Oidc().RedirectURL, c.cfg.Oidc().Scopes)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == oidc.ErrorRetrieveUserInfo {
			status = http.StatusInternalServerError
		}
		ctx.JSON(status, gin.H{
			"code":    status,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	user := contracts.NewUser(userInfo.Subject)

	err = services.Login(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "login_failed",
			"data":    nil,
		})
		return
	}
	sess := session.Default(ctx)
	sess.Regenerate()
	if err := sess.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "unable_to_save_session",
			"data":    nil,
		})
		return
	}

	session.AddCookieToResponse(ctx, sess.Id())
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "login_success",
		"data":    nil,
	})
}
