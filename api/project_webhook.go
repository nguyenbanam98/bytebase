package api

import (
	"context"
	"encoding/json"
)

type ProjectWebhook struct {
	ID int `jsonapi:"primary,projectWebhookMember"`

	// Standard fields
	CreatorId int
	Creator   *Principal `jsonapi:"attr,creator"`
	CreatedTs int64      `jsonapi:"attr,createdTs"`
	UpdaterId int
	Updater   *Principal `jsonapi:"attr,updater"`
	UpdatedTs int64      `jsonapi:"attr,updatedTs"`

	// Related fields
	// Just returns ProjectId since it always operates within the project context
	ProjectId int `jsonapi:"attr,projecId"`

	// Domain specific fields
	Type         string   `jsonapi:"attr,type"`
	Name         string   `jsonapi:"attr,name"`
	URL          string   `jsonapi:"attr,url"`
	ActivityList []string `jsonapi:"attr,activityList"`
}

type ProjectWebhookCreate struct {
	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	CreatorId int

	// Related fields
	ProjectId int

	// Domain specific fields
	Type         string   `jsonapi:"attr,type"`
	Name         string   `jsonapi:"attr,name"`
	URL          string   `jsonapi:"attr,url"`
	ActivityList []string `jsonapi:"attr,activityList"`
}

type ProjectWebhookFind struct {
	ID *int

	// Related fields
	ProjectId    *int
	ActivityType *ActivityType
}

func (find *ProjectWebhookFind) String() string {
	str, err := json.Marshal(*find)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

type ProjectWebhookPatch struct {
	ID int

	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	UpdaterId int

	// Domain specific fields
	Name         *string `jsonapi:"attr,name"`
	URL          *string `jsonapi:"attr,url"`
	ActivityList *string `jsonapi:"attr,activityList"`
}

type ProjectWebhookDelete struct {
	ID int

	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	DeleterId int
}

type ProjectWebhookTestResult struct {
	Error string `jsonapi:"attr,error"`
}

type ProjectWebhookService interface {
	CreateProjectWebhook(ctx context.Context, create *ProjectWebhookCreate) (*ProjectWebhook, error)
	FindProjectWebhookList(ctx context.Context, find *ProjectWebhookFind) ([]*ProjectWebhook, error)
	FindProjectWebhook(ctx context.Context, find *ProjectWebhookFind) (*ProjectWebhook, error)
	PatchProjectWebhook(ctx context.Context, patch *ProjectWebhookPatch) (*ProjectWebhook, error)
	DeleteProjectWebhook(ctx context.Context, delete *ProjectWebhookDelete) error
}
