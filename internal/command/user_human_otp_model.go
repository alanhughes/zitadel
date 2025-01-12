package command

import (
	"github.com/zitadel/zitadel/internal/crypto"
	"github.com/zitadel/zitadel/internal/domain"
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/repository/user"
)

type HumanTOTPWriteModel struct {
	eventstore.WriteModel

	State  domain.MFAState
	Secret *crypto.CryptoValue
}

func NewHumanTOTPWriteModel(userID, resourceOwner string) *HumanTOTPWriteModel {
	return &HumanTOTPWriteModel{
		WriteModel: eventstore.WriteModel{
			AggregateID:   userID,
			ResourceOwner: resourceOwner,
		},
	}
}

func (wm *HumanTOTPWriteModel) Reduce() error {
	for _, event := range wm.Events {
		switch e := event.(type) {
		case *user.HumanOTPAddedEvent:
			wm.Secret = e.Secret
			wm.State = domain.MFAStateNotReady
		case *user.HumanOTPVerifiedEvent:
			wm.State = domain.MFAStateReady
		case *user.HumanOTPRemovedEvent:
			wm.State = domain.MFAStateRemoved
		case *user.UserRemovedEvent:
			wm.State = domain.MFAStateRemoved
		}
	}
	return wm.WriteModel.Reduce()
}

func (wm *HumanTOTPWriteModel) Query() *eventstore.SearchQueryBuilder {
	query := eventstore.NewSearchQueryBuilder(eventstore.ColumnsEvent).
		AddQuery().
		AggregateTypes(user.AggregateType).
		AggregateIDs(wm.AggregateID).
		EventTypes(user.HumanMFAOTPAddedType,
			user.HumanMFAOTPVerifiedType,
			user.HumanMFAOTPRemovedType,
			user.UserRemovedType,
			user.UserV1MFAOTPAddedType,
			user.UserV1MFAOTPVerifiedType,
			user.UserV1MFAOTPRemovedType).
		Builder()

	if wm.ResourceOwner != "" {
		query.ResourceOwner(wm.ResourceOwner)
	}
	return query
}

type HumanOTPSMSWriteModel struct {
	eventstore.WriteModel

	phoneVerified bool
	otpAdded      bool
}

func NewHumanOTPSMSWriteModel(userID, resourceOwner string) *HumanOTPSMSWriteModel {
	return &HumanOTPSMSWriteModel{
		WriteModel: eventstore.WriteModel{
			AggregateID:   userID,
			ResourceOwner: resourceOwner,
		},
	}
}

func (wm *HumanOTPSMSWriteModel) Reduce() error {
	for _, event := range wm.Events {
		switch event.(type) {
		case *user.HumanPhoneVerifiedEvent:
			wm.phoneVerified = true
		case *user.HumanOTPSMSAddedEvent:
			wm.otpAdded = true
		case *user.HumanOTPSMSRemovedEvent:
			wm.otpAdded = false
		case *user.HumanPhoneRemovedEvent,
			*user.UserRemovedEvent:
			wm.phoneVerified = false
			wm.otpAdded = false
		}
	}
	return wm.WriteModel.Reduce()
}

func (wm *HumanOTPSMSWriteModel) Query() *eventstore.SearchQueryBuilder {
	query := eventstore.NewSearchQueryBuilder(eventstore.ColumnsEvent).
		AddQuery().
		AggregateTypes(user.AggregateType).
		AggregateIDs(wm.AggregateID).
		EventTypes(user.HumanPhoneVerifiedType,
			user.HumanOTPSMSAddedType,
			user.HumanOTPSMSRemovedType,
			user.HumanPhoneRemovedType,
			user.UserRemovedType,
		).
		Builder()

	if wm.ResourceOwner != "" {
		query.ResourceOwner(wm.ResourceOwner)
	}
	return query
}

type HumanOTPEmailWriteModel struct {
	eventstore.WriteModel

	emailVerified bool
	otpAdded      bool
}

func NewHumanOTPEmailWriteModel(userID, resourceOwner string) *HumanOTPEmailWriteModel {
	return &HumanOTPEmailWriteModel{
		WriteModel: eventstore.WriteModel{
			AggregateID:   userID,
			ResourceOwner: resourceOwner,
		},
	}
}

func (wm *HumanOTPEmailWriteModel) Reduce() error {
	for _, event := range wm.Events {
		switch event.(type) {
		case *user.HumanEmailVerifiedEvent:
			wm.emailVerified = true
		case *user.HumanOTPEmailAddedEvent:
			wm.otpAdded = true
		case *user.HumanOTPEmailRemovedEvent:
			wm.otpAdded = false
		case *user.UserRemovedEvent:
			wm.emailVerified = false
			wm.otpAdded = false
		}
	}
	return wm.WriteModel.Reduce()
}

func (wm *HumanOTPEmailWriteModel) Query() *eventstore.SearchQueryBuilder {
	query := eventstore.NewSearchQueryBuilder(eventstore.ColumnsEvent).
		AddQuery().
		AggregateTypes(user.AggregateType).
		AggregateIDs(wm.AggregateID).
		EventTypes(user.HumanEmailVerifiedType,
			user.HumanOTPEmailAddedType,
			user.HumanOTPEmailRemovedType,
			user.UserRemovedType,
		).
		Builder()

	if wm.ResourceOwner != "" {
		query.ResourceOwner(wm.ResourceOwner)
	}
	return query
}
