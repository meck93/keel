package auditor

import (
	"github.com/google/uuid"

	"github.com/meck93/keel/extension/notification"
	"github.com/meck93/keel/pkg/store"
	"github.com/meck93/keel/types"

	log "github.com/sirupsen/logrus"
)

type auditor struct {
	store store.Store
}

func New(store store.Store) *auditor {
	return &auditor{
		store: store,
	}
}

func (a *auditor) Configure(config *notification.Config) (bool, error) {

	log.WithFields(log.Fields{
		"name": "auditor",
	}).Info("extension.notification.auditor: audit logger configured")

	return true, nil
}

func (a *auditor) Send(event types.EventNotification) error {
	al := &types.AuditLog{
		ID:           uuid.New().String(),
		AccountID:    "system",
		Username:     "system",
		Action:       event.Type.String(),
		ResourceKind: event.ResourceKind,
		Identifier:   event.Identifier,
		Message:      event.Message,
	}
	al.SetMetadata(event.Metadata)
	_, err := a.store.CreateAuditLog(al)

	return err
}
