package command

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"golang.org/x/text/language"

	"github.com/zitadel/zitadel/internal/crypto"
	"github.com/zitadel/zitadel/internal/domain"
	"github.com/zitadel/zitadel/internal/errors"
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/eventstore/repository"
	"github.com/zitadel/zitadel/internal/eventstore/repository/mock"
	action_repo "github.com/zitadel/zitadel/internal/repository/action"
	"github.com/zitadel/zitadel/internal/repository/idpintent"
	iam_repo "github.com/zitadel/zitadel/internal/repository/instance"
	key_repo "github.com/zitadel/zitadel/internal/repository/keypair"
	"github.com/zitadel/zitadel/internal/repository/org"
	proj_repo "github.com/zitadel/zitadel/internal/repository/project"
	"github.com/zitadel/zitadel/internal/repository/session"
	usr_repo "github.com/zitadel/zitadel/internal/repository/user"
	"github.com/zitadel/zitadel/internal/repository/usergrant"
)

type expect func(mockRepository *mock.MockRepository)

func eventstoreExpect(t *testing.T, expects ...expect) *eventstore.Eventstore {
	m := mock.NewRepo(t)
	for _, e := range expects {
		e(m)
	}
	es := eventstore.NewEventstore(
		&eventstore.Config{
			Querier: m.MockQuerier,
			Pusher:  m.MockPusher,
		},
	)
	iam_repo.RegisterEventMappers(es)
	org.RegisterEventMappers(es)
	usr_repo.RegisterEventMappers(es)
	proj_repo.RegisterEventMappers(es)
	usergrant.RegisterEventMappers(es)
	key_repo.RegisterEventMappers(es)
	action_repo.RegisterEventMappers(es)
	session.RegisterEventMappers(es)
	idpintent.RegisterEventMappers(es)
	return es
}

func eventPusherToEvents(eventsPushes ...eventstore.Command) []*repository.Event {
	events := make([]*repository.Event, len(eventsPushes))
	for i, event := range eventsPushes {
		data, err := eventstore.EventData(event)
		if err != nil {
			return nil
		}
		events[i] = &repository.Event{
			AggregateID:   event.Aggregate().ID,
			AggregateType: event.Aggregate().Type,
			ResourceOwner: sql.NullString{String: event.Aggregate().ResourceOwner, Valid: event.Aggregate().ResourceOwner != ""},
			EditorService: "zitadel",
			EditorUser:    event.Creator(),
			Typ:           event.Type(),
			Version:       event.Aggregate().Version,
			Data:          data,
		}
	}
	return events
}

func expectPush(commands ...eventstore.Command) expect {
	return func(m *mock.MockRepository) {
		m.ExpectPush(commands)
	}
}

func expectPushFailed(err error, commands ...eventstore.Command) expect {
	return func(m *mock.MockRepository) {
		m.ExpectPushFailed(err, commands)
	}
}

func expectFilter(events ...eventstore.Event) expect {
	return func(m *mock.MockRepository) {
		m.ExpectFilterEvents(events...)
	}
}
func expectFilterError(err error) expect {
	return func(m *mock.MockRepository) {
		m.ExpectFilterEventsError(err)
	}
}

func expectFilterOrgDomainNotFound() expect {
	return func(m *mock.MockRepository) {
		m.ExpectFilterNoEventsNoError()
	}
}

func expectFilterOrgMemberNotFound() expect {
	return func(m *mock.MockRepository) {
		m.ExpectFilterNoEventsNoError()
	}
}

func eventFromCommand(event eventstore.Command) eventstore.Event {
	data, _ := eventstore.EventData(event)
	return &repository.Event{
		InstanceID:                    event.Aggregate().InstanceID,
		ID:                            "",
		Seq:                           0,
		PreviousAggregateSequence:     0,
		PreviousAggregateTypeSequence: 0,
		CreationDate:                  time.Time{},
		Typ:                           event.Type(),
		Data:                          data,
		EditorService:                 "zitadel",
		EditorUser:                    event.Creator(),
		Version:                       event.Aggregate().Version,
		AggregateID:                   event.Aggregate().ID,
		AggregateType:                 event.Aggregate().Type,
		ResourceOwner:                 sql.NullString{String: event.Aggregate().ResourceOwner, Valid: event.Aggregate().ResourceOwner != ""},
	}
}

func eventFromEventPusherWithInstanceID(instanceID string, event eventstore.Command) *repository.Event {
	data, _ := eventstore.EventData(event)
	return &repository.Event{
		ID:                            "",
		Seq:                           0,
		PreviousAggregateSequence:     0,
		PreviousAggregateTypeSequence: 0,
		CreationDate:                  time.Time{},
		Typ:                           event.Type(),
		Data:                          data,
		EditorService:                 "zitadel",
		EditorUser:                    event.Creator(),
		Version:                       event.Aggregate().Version,
		AggregateID:                   event.Aggregate().ID,
		AggregateType:                 event.Aggregate().Type,
		ResourceOwner:                 sql.NullString{String: event.Aggregate().ResourceOwner, Valid: event.Aggregate().ResourceOwner != ""},
		InstanceID:                    instanceID,
	}
}

func eventFromCommandWithCreationDateNow(event eventstore.Command) eventstore.Event {
	e := eventFromCommand(event)
	e.(*repository.Event).CreationDate = time.Now()
	return e
}

func GetMockSecretGenerator(t *testing.T) crypto.Generator {
	ctrl := gomock.NewController(t)
	alg := crypto.CreateMockEncryptionAlg(ctrl)
	generator := crypto.NewMockGenerator(ctrl)
	generator.EXPECT().Length().Return(uint(1)).AnyTimes()
	generator.EXPECT().Runes().Return([]rune("aa")).AnyTimes()
	generator.EXPECT().Alg().Return(alg).AnyTimes()
	generator.EXPECT().Expiry().Return(time.Hour * 1).AnyTimes()

	return generator
}

type mockInstance struct{}

func (m *mockInstance) InstanceID() string {
	return "INSTANCE"
}

func (m *mockInstance) ProjectID() string {
	return "projectID"
}

func (m *mockInstance) ConsoleClientID() string {
	return "consoleID"
}

func (m *mockInstance) ConsoleApplicationID() string {
	return "consoleApplicationID"
}

func (m *mockInstance) DefaultLanguage() language.Tag {
	return language.English
}

func (m *mockInstance) DefaultOrganisationID() string {
	return "defaultOrgID"
}

func (m *mockInstance) RequestedDomain() string {
	return "zitadel.cloud"
}

func (m *mockInstance) RequestedHost() string {
	return "zitadel.cloud:443"
}

func (m *mockInstance) SecurityPolicyAllowedOrigins() []string {
	return nil
}

func newMockPermissionCheckAllowed() domain.PermissionCheck {
	return func(ctx context.Context, permission, orgID, resourceID string) (err error) {
		return nil
	}
}

func newMockPermissionCheckNotAllowed() domain.PermissionCheck {
	return func(ctx context.Context, permission, orgID, resourceID string) (err error) {
		return errors.ThrowPermissionDenied(nil, "AUTHZ-HKJD33", "Errors.PermissionDenied")
	}
}
