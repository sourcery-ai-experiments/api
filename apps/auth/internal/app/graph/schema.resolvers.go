package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"errors"

	"github.com/kloudlite/api/apps/auth/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/auth/internal/app/graph/model"
	"github.com/kloudlite/api/common"
	klErrors "github.com/kloudlite/api/pkg/errors"
	httpServer "github.com/kloudlite/api/pkg/http-server"
	"github.com/kloudlite/api/pkg/repos"
)

// AuthSetRemoteAuthHeader is the resolver for the auth_setRemoteAuthHeader field.
func (r *mutationResolver) AuthSetRemoteAuthHeader(ctx context.Context, loginID string, authHeader *string) (bool, error) {
	err := r.d.SetRemoteLoginAuthHeader(ctx, repos.ID(loginID), *authHeader)
	return err == nil, klErrors.NewE(err)
}

// AuthCreateRemoteLogin is the resolver for the auth_createRemoteLogin field.
func (r *mutationResolver) AuthCreateRemoteLogin(ctx context.Context, secret *string) (string, error) {
	login, err := r.d.CreateRemoteLogin(ctx, *secret)
	if err != nil {
		return "", klErrors.NewE(err)
	}
	return string(login), nil
}

// AuthLogin is the resolver for the auth_login field.
func (r *mutationResolver) AuthLogin(ctx context.Context, email string, password string) (*model.Session, error) {
	sessionEntity, err := r.d.Login(ctx, email, password)
	if err != nil {
		return nil, klErrors.NewE(err)
	}

	if !sessionEntity.UserVerified {
		return nil, errors.New("user email not verified")
	}

	httpServer.SetSession(ctx, sessionEntity)
	return sessionModelFromAuthSession(sessionEntity), err
}

// AuthSignup is the resolver for the auth_signup field.
func (r *mutationResolver) AuthSignup(ctx context.Context, name string, email string, password string) (*model.Session, error) {
	sessionEntity, err := r.d.SignUp(ctx, name, email, password)
	if err != nil {
		return nil, klErrors.NewE(err)
	}
	httpServer.SetSession(ctx, sessionEntity)
	session := sessionModelFromAuthSession(sessionEntity)
	return session, err
}

// AuthLogout is the resolver for the auth_logout field.
func (r *mutationResolver) AuthLogout(ctx context.Context) (bool, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return true, nil
	}
	httpServer.DeleteSession(ctx)
	return true, nil
}

// AuthSetMetadata is the resolver for the auth_setMetadata field.
func (r *mutationResolver) AuthSetMetadata(ctx context.Context, values map[string]interface{}) (*model.User, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return nil, errors.New("user not logged in")
	}
	userEntity, err := r.d.SetUserMetadata(ctx, session.UserId, values)
	return userModelFromEntity(userEntity), klErrors.NewE(err)
}

// AuthClearMetadata is the resolver for the auth_clearMetadata field.
func (r *mutationResolver) AuthClearMetadata(ctx context.Context) (*model.User, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return nil, errors.New("user not logged in")
	}
	userEntity, err := r.d.ClearUserMetadata(ctx, session.UserId)
	return userModelFromEntity(userEntity), klErrors.NewE(err)
}

// AuthVerifyEmail is the resolver for the auth_verifyEmail field.
func (r *mutationResolver) AuthVerifyEmail(ctx context.Context, token string) (*model.Session, error) {
	sessionEntity, err := r.d.VerifyEmail(ctx, token)
	httpServer.SetSession(ctx, sessionEntity)
	return sessionModelFromAuthSession(sessionEntity), klErrors.NewE(err)
}

// AuthResetPassword is the resolver for the auth_resetPassword field.
func (r *mutationResolver) AuthResetPassword(ctx context.Context, token string, password string) (bool, error) {
	return r.d.ResetPassword(ctx, token, password)
}

// AuthRequestResetPassword is the resolver for the auth_requestResetPassword field.
func (r *mutationResolver) AuthRequestResetPassword(ctx context.Context, email string) (bool, error) {
	return r.d.RequestResetPassword(ctx, email)
}

// AuthChangeEmail is the resolver for the auth_changeEmail field.
func (r *mutationResolver) AuthChangeEmail(ctx context.Context, email string) (bool, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return false, errors.New("user is not logged in")
	}
	return r.d.ChangeEmail(ctx, session.UserId, email)
}

// AuthResendVerificationEmail is the resolver for the auth_resendVerificationEmail field.
func (r *mutationResolver) AuthResendVerificationEmail(ctx context.Context) (bool, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return false, errors.New("user is not logged in")
	}
	return r.d.ResendVerificationEmail(ctx, session.UserId)
}

// AuthChangePassword is the resolver for the auth_changePassword field.
func (r *mutationResolver) AuthChangePassword(ctx context.Context, currentPassword string, newPassword string) (bool, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return false, errors.New("user is not logged in")
	}
	return r.d.ChangePassword(ctx, session.UserId, currentPassword, newPassword)
}

// OAuthLogin is the resolver for the oAuth_login field.
func (r *mutationResolver) OAuthLogin(ctx context.Context, provider string, code string, state *string) (*model.Session, error) {
	st := ""
	if state != nil {
		st = *state
	}
	sessionEntity, err := r.d.OauthLogin(ctx, provider, st, code)
	if err != nil {
		return nil, klErrors.NewEf(err, "could not create session")
	}
	httpServer.SetSession(ctx, sessionEntity)
	return sessionModelFromAuthSession(sessionEntity), klErrors.NewE(err)
}

// OAuthAddLogin is the resolver for the oAuth_addLogin field.
func (r *mutationResolver) OAuthAddLogin(ctx context.Context, provider string, state string, code string) (bool, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return false, errors.New("user is not logged in")
	}
	return r.d.OauthAddLogin(ctx, session.UserId, provider, state, code)
}

// AuthMe is the resolver for the auth_me field.
func (r *queryResolver) AuthMe(ctx context.Context) (*model.User, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return nil, errors.New("user not logged in")
	}
	u, err := r.d.GetUserById(ctx, session.UserId)
	if err != nil {
		return nil, klErrors.NewE(err)
	}
	if u == nil {
		return nil, klErrors.Newf("user(email=%s) does not exist in system", session.UserEmail)
	}
	return userModelFromEntity(u), klErrors.NewE(err)
}

// AuthFindByEmail is the resolver for the auth_findByEmail field.
func (r *queryResolver) AuthFindByEmail(ctx context.Context, email string) (*model.User, error) {
	userEntity, err := r.d.GetUserByEmail(ctx, email)
	return userModelFromEntity(userEntity), klErrors.NewE(err)
}

// OAuthRequestLogin is the resolver for the oAuth_requestLogin field.
func (r *queryResolver) OAuthRequestLogin(ctx context.Context, provider string, state *string) (string, error) {
	_state := ""
	if state != nil {
		_state = *state
	}
	url, err := r.d.OauthRequestLogin(ctx, provider, _state)
	if err != nil {
		return "", klErrors.NewE(err)
	}
	return url, nil
}

// AuthGetRemoteLogin is the resolver for the auth_getRemoteLogin field.
func (r *queryResolver) AuthGetRemoteLogin(ctx context.Context, loginID string, secret string) (*model.RemoteLogin, error) {
	login, err := r.d.GetRemoteLogin(ctx, repos.ID(loginID), secret)
	if err != nil {
		return nil, klErrors.NewE(err)
	}
	return &model.RemoteLogin{
		Status:     string(login.LoginStatus),
		AuthHeader: &login.AuthHeader,
	}, nil
}

// AuthListOAuthProviders is the resolver for the auth_listOAuthProviders field.
func (r *queryResolver) AuthListOAuthProviders(ctx context.Context) ([]*model.OAuthProviderStatus, error) {
	if !r.ev.OAuth2Enabled {
		return []*model.OAuthProviderStatus{}, nil
	}
	return []*model.OAuthProviderStatus{
		{
			Provider: "github",
			Enabled:  r.ev.OAuth2GithubEnabled,
		},
		{
			Provider: "gitlab",
			Enabled:  r.ev.OAuth2GitlabEnabled,
		},
		{
			Provider: "google",
			Enabled:  r.ev.OAuth2GoogleEnabled,
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
