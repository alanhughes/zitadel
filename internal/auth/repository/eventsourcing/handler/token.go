package handler

import (
	"context"
	"encoding/json"

	"github.com/zitadel/logging"

	auth_view "github.com/zitadel/zitadel/internal/auth/repository/eventsourcing/view"
	caos_errs "github.com/zitadel/zitadel/internal/errors"
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/eventstore/handler/v2"
	handler2 "github.com/zitadel/zitadel/internal/eventstore/handler/v2"
	es_models "github.com/zitadel/zitadel/internal/eventstore/v1/models"
	proj_model "github.com/zitadel/zitadel/internal/project/model"
	project_es_model "github.com/zitadel/zitadel/internal/project/repository/eventsourcing/model"
	proj_view "github.com/zitadel/zitadel/internal/project/repository/view"
	"github.com/zitadel/zitadel/internal/repository/instance"
	"github.com/zitadel/zitadel/internal/repository/org"
	"github.com/zitadel/zitadel/internal/repository/project"
	"github.com/zitadel/zitadel/internal/repository/user"
	view_model "github.com/zitadel/zitadel/internal/user/repository/view/model"
)

const (
	tokenTable = "auth.tokens"
)

var _ handler2.Projection = (*Token)(nil)

type Token struct {
	view *auth_view.View
	es   handler.EventStore
}

func newToken(
	ctx context.Context,
	config handler2.Config,
	view *auth_view.View,
) *handler2.Handler {
	return handler2.NewHandler(
		ctx,
		&config,
		&Token{
			view: view,
			es:   config.Eventstore,
		},
	)
}

// Name implements [handler.Projection]
func (*Token) Name() string {
	return tokenTable
}

// Reducers implements [handler.Projection]
func (t *Token) Reducers() []handler2.AggregateReducer {
	return []handler2.AggregateReducer{
		{
			Aggregate: user.AggregateType,
			EventRedusers: []handler2.EventReducer{
				{
					Event:  user.UserTokenAddedType,
					Reduce: t.Reduce,
				},
				{
					Event:  user.UserV1ProfileChangedType,
					Reduce: t.Reduce,
				},
				{
					Event:  user.HumanProfileChangedType,
					Reduce: t.Reduce,
				},
				{
					Event:  user.UserV1SignedOutType,
					Reduce: t.Reduce,
				},
				{
					Event:  user.HumanSignedOutType,
					Reduce: t.Reduce,
				},
				{
					Event:  user.UserLockedType,
					Reduce: t.Reduce,
				},
				{
					Event:  user.UserDeactivatedType,
					Reduce: t.Reduce,
				},
				{
					Event:  user.UserRemovedType,
					Reduce: t.Reduce,
				},
				{
					Event:  user.UserTokenRemovedType,
					Reduce: t.Reduce,
				},
				{
					Event:  user.PersonalAccessTokenRemovedType,
					Reduce: t.Reduce,
				},
				{
					Event:  user.HumanRefreshTokenRemovedType,
					Reduce: t.Reduce,
				},
			},
		},
		{
			Aggregate: project.AggregateType,
			EventRedusers: []handler2.EventReducer{
				{
					Event:  project.ApplicationDeactivatedType,
					Reduce: t.Reduce,
				},
				{
					Event:  project.ApplicationRemovedType,
					Reduce: t.Reduce,
				},
				{
					Event:  project.ProjectDeactivatedType,
					Reduce: t.Reduce,
				},
				{
					Event:  project.ProjectRemovedType,
					Reduce: t.Reduce,
				},
			},
		},
		{
			Aggregate: org.AggregateType,
			EventRedusers: []handler2.EventReducer{
				{
					Event:  org.OrgRemovedEventType,
					Reduce: t.Reduce,
				},
			},
		},
		{
			Aggregate: instance.AggregateType,
			EventRedusers: []handler2.EventReducer{
				{
					Event:  instance.InstanceRemovedEventType,
					Reduce: t.Reduce,
				},
			},
		},
	}
}

func (t *Token) Reduce(event eventstore.Event) (_ *handler2.Statement, err error) {
	switch event.Type() {
	case user.UserTokenAddedType,
		user.PersonalAccessTokenAddedType:
		token := new(view_model.TokenView)
		err := token.AppendEvent(event)
		if err != nil {
			return nil, err
		}
		return handler2.NewStatement(event,
			func(ex handler2.Executer, projectionName string) error {
				return t.view.PutToken(token)
			}), nil
	case user.UserV1ProfileChangedType,
		user.HumanProfileChangedType:
		user := new(view_model.UserView)
		err := user.AppendEvent(event)
		if err != nil {
			return nil, err
		}
		tokens, err := t.view.TokensByUserID(event.Aggregate().ID, event.Aggregate().InstanceID)
		if err != nil {
			return nil, err
		}
		for _, token := range tokens {
			token.PreferredLanguage = user.PreferredLanguage
		}
		return handler2.NewStatement(event,
			func(ex handler2.Executer, projectionName string) error {
				return t.view.PutTokens(tokens)
			}), nil
	case user.UserV1SignedOutType,
		user.HumanSignedOutType:
		id, err := agentIDFromSession(event)
		if err != nil {
			return nil, err
		}
		return handler2.NewStatement(event,
			func(ex handler2.Executer, projectionName string) error {
				return t.view.DeleteSessionTokens(id, event)
			}), nil
	case user.UserLockedType,
		user.UserDeactivatedType,
		user.UserRemovedType:
		return handler2.NewStatement(event,
			func(ex handler2.Executer, projectionName string) error {
				return t.view.DeleteUserTokens(event)
			}), nil
	case user.UserTokenRemovedType,
		user.PersonalAccessTokenRemovedType:
		id, err := tokenIDFromRemovedEvent(event)
		if err != nil {
			return nil, err
		}
		return handler2.NewStatement(event,
			func(ex handler2.Executer, projectionName string) error {
				return t.view.DeleteToken(id, event.Aggregate().InstanceID)
			}), nil

	case user.HumanRefreshTokenRemovedType:
		id, err := refreshTokenIDFromRemovedEvent(event)
		if err != nil {
			return nil, err
		}
		return handler2.NewStatement(event,
			func(ex handler2.Executer, projectionName string) error {
				return t.view.DeleteTokensFromRefreshToken(id, event.Aggregate().InstanceID)
			}), nil
	case project.ApplicationDeactivatedType,
		project.ApplicationRemovedType:
		application, err := applicationFromSession(event)
		if err != nil {
			return nil, err
		}
		return handler2.NewStatement(event,
			func(ex handler2.Executer, projectionName string) error {
				return t.view.DeleteApplicationTokens(event, application.AppID)
			}), nil
	case project.ProjectDeactivatedType,
		project.ProjectRemovedType:
		project, err := t.getProjectByID(context.Background(), event.Aggregate().ID, event.Aggregate().InstanceID)
		if err != nil {
			return nil, err
		}
		applicationIDs := make([]string, 0, len(project.Applications))
		for _, app := range project.Applications {
			if app.OIDCConfig != nil {
				applicationIDs = append(applicationIDs, app.OIDCConfig.ClientID)
			}
		}
		return handler2.NewStatement(event,
			func(ex handler2.Executer, projectionName string) error {
				return t.view.DeleteApplicationTokens(event, applicationIDs...)
			}), nil
	case instance.InstanceRemovedEventType:
		return handler2.NewStatement(event,
			func(ex handler2.Executer, projectionName string) error {
				return t.view.DeleteInstanceTokens(event)
			}), nil
	case org.OrgRemovedEventType:
		// deletes all tokens including PATs, which is expected for now
		// if there is an undo of the org deletion in the future,
		// we will need to have a look on how to handle the deleted PATs
		return handler2.NewStatement(event,
			func(ex handler2.Executer, projectionName string) error {
				return t.view.DeleteOrgTokens(event)
			}), nil
	default:
		return handler2.NewNoOpStatement(event), nil
	}
}

func agentIDFromSession(event eventstore.Event) (string, error) {
	session := make(map[string]interface{})
	if err := json.Unmarshal(event.DataAsBytes(), &session); err != nil {
		logging.WithError(err).Error("could not unmarshal event data")
		return "", caos_errs.ThrowInternal(nil, "MODEL-sd325", "could not unmarshal data")
	}
	return session["userAgentID"].(string), nil
}

func applicationFromSession(event eventstore.Event) (*project_es_model.Application, error) {
	application := new(project_es_model.Application)
	if err := json.Unmarshal(event.DataAsBytes(), &application); err != nil {
		logging.WithError(err).Error("could not unmarshal event data")
		return nil, caos_errs.ThrowInternal(nil, "MODEL-Hrw1q", "could not unmarshal data")
	}
	return application, nil
}

func tokenIDFromRemovedEvent(event eventstore.Event) (string, error) {
	removed := make(map[string]interface{})
	if err := json.Unmarshal(event.DataAsBytes(), &removed); err != nil {
		logging.WithError(err).Error("could not unmarshal event data")
		return "", caos_errs.ThrowInternal(nil, "MODEL-Sff32", "could not unmarshal data")
	}
	return removed["tokenId"].(string), nil
}

func refreshTokenIDFromRemovedEvent(event eventstore.Event) (string, error) {
	removed := make(map[string]interface{})
	if err := json.Unmarshal(event.DataAsBytes(), &removed); err != nil {
		logging.WithError(err).Error("could not unmarshal event data")
		return "", caos_errs.ThrowInternal(nil, "MODEL-Dfb3w", "could not unmarshal data")
	}
	return removed["tokenId"].(string), nil
}

func (t *Token) getProjectByID(ctx context.Context, projID, instanceID string) (*proj_model.Project, error) {
	query, err := proj_view.ProjectByIDQuery(projID, instanceID, 0)
	if err != nil {
		return nil, err
	}
	esProject := &project_es_model.Project{
		ObjectRoot: es_models.ObjectRoot{
			AggregateID: projID,
		},
	}
	events, err := t.es.Filter(ctx, query)
	if err != nil {
		return nil, err
	}
	if err = esProject.AppendEvents(events...); err != nil {
		return nil, err
	}

	if esProject.Sequence == 0 {
		return nil, caos_errs.ThrowNotFound(nil, "EVENT-Dsdw2", "Errors.Project.NotFound")
	}
	return project_es_model.ProjectToModel(esProject), nil
}
