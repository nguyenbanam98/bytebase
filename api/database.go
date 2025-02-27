package api

import (
	"context"
	"database/sql"
	"encoding/json"
)

const (
	ALL_DATABASE_NAME          = "*"
	DEFAULT_CHARACTER_SET_NAME = "utf8mb4"
	// Use utf8mb4_general_ci instead of the new MySQL 8.0.1 default utf8mb4_0900_ai_ci
	// because the former is compatible with more other MySQL flavors (e.g. MariaDB)
	DEFAULT_COLLATION_NAME = "utf8mb4_general_ci"
)

type SyncStatus string

const (
	OK       SyncStatus = "OK"
	NotFound SyncStatus = "NOT_FOUND"
)

func (e SyncStatus) String() string {
	switch e {
	case OK:
		return "OK"
	case NotFound:
		return "NOT_FOUND"
	}
	return ""
}

type Database struct {
	ID int `jsonapi:"primary,database"`

	// Standard fields
	CreatorId int
	Creator   *Principal `jsonapi:"attr,creator"`
	CreatedTs int64      `jsonapi:"attr,createdTs"`
	UpdaterId int
	Updater   *Principal `jsonapi:"attr,updater"`
	UpdatedTs int64      `jsonapi:"attr,updatedTs"`

	// Related fields
	ProjectId      int
	Project        *Project `jsonapi:"relation,project"`
	InstanceId     int
	Instance       *Instance     `jsonapi:"relation,instance"`
	DataSourceList []*DataSource `jsonapi:"relation,dataSource"`

	// Domain specific fields
	Name                 string     `jsonapi:"attr,name"`
	CharacterSet         string     `jsonapi:"attr,characterSet"`
	Collation            string     `jsonapi:"attr,collation"`
	SyncStatus           SyncStatus `jsonapi:"attr,syncStatus"`
	LastSuccessfulSyncTs int64      `jsonapi:"attr,lastSuccessfulSyncTs"`
}

type DatabaseCreate struct {
	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	CreatorId int

	// Related fields
	ProjectId  int `jsonapi:"attr,projectId"`
	InstanceId int `jsonapi:"attr,instanceId"`

	// Domain specific fields
	Name           string `jsonapi:"attr,name"`
	CharacterSet   string `jsonapi:"attr,characterSet"`
	Collation      string `jsonapi:"attr,collation"`
	IssueId        int    `jsonapi:"attr,issueId"`
	TimezoneName   string
	TimezoneOffset int
}

type DatabaseFind struct {
	ID *int

	// Related fields
	InstanceId *int
	ProjectId  *int

	// Domain specific fields
	Name               *string
	IncludeAllDatabase bool
}

func (find *DatabaseFind) String() string {
	str, err := json.Marshal(*find)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

type DatabasePatch struct {
	ID int

	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	UpdaterId int

	// Related fields
	ProjectId *int `jsonapi:"attr,projectId"`

	// Domain specific fields
	SyncStatus           *SyncStatus
	LastSuccessfulSyncTs *int64
}

type DatabaseService interface {
	CreateDatabase(ctx context.Context, create *DatabaseCreate) (*Database, error)
	// This is specifically used to create the * database when creating the instance.
	CreateDatabaseTx(ctx context.Context, tx *sql.Tx, create *DatabaseCreate) (*Database, error)
	FindDatabaseList(ctx context.Context, find *DatabaseFind) ([]*Database, error)
	FindDatabase(ctx context.Context, find *DatabaseFind) (*Database, error)
	PatchDatabase(ctx context.Context, patch *DatabasePatch) (*Database, error)
}
