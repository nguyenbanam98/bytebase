package api

import (
	"context"
	"encoding/json"
)

// BackupStatus is the status of a backup.
type BackupStatus string

const (
	// BackupStatusPendingCreate is the status for PENDING_CREATE.
	BackupStatusPendingCreate BackupStatus = "PENDING_CREATE"
	// BackupStatusDone is the status for DONE.
	BackupStatusDone BackupStatus = "DONE"
	// BackupStatusFailed is the status for FAILED.
	BackupStatusFailed BackupStatus = "FAILED"
)

func (e BackupStatus) String() string {
	switch e {
	case BackupStatusPendingCreate:
		return "PENDING_CREATE"
	case BackupStatusDone:
		return "DONE"
	case BackupStatusFailed:
		return "FAILED"
	}
	return "UNKNOWN"
}

// BackupType is the type of a backup.
type BackupType string

const (
	// BackupTypeAutomatic is the type for automatic backup.
	BackupTypeAutomatic BackupType = "AUTOMATIC"
	// BackupTypeManual is the type for manual backup.
	BackupTypeManual BackupType = "MANUAL"
)

func (e BackupType) String() string {
	switch e {
	case BackupTypeAutomatic:
		return "AUTOMATIC"
	case BackupTypeManual:
		return "MANUAL"
	}
	return "UNKNOWN"
}

// BackupStorageBackend is the storage backend of a backup.
type BackupStorageBackend string

const (
	// BackupStorageBackendLocal is the local storage backend for a backup.
	BackupStorageBackendLocal BackupStorageBackend = "LOCAL"
)

func (e BackupStorageBackend) String() string {
	switch e {
	case BackupStorageBackendLocal:
		return "LOCAL"
	}
	return "UNKNOWN"
}

type Backup struct {
	ID int `jsonapi:"primary,backup"`

	// Standard fields
	CreatorId int
	Creator   *Principal `jsonapi:"attr,creator"`
	CreatedTs int64      `jsonapi:"attr,createdTs"`
	UpdaterId int
	Updater   *Principal `jsonapi:"attr,updater"`
	UpdatedTs int64      `jsonapi:"attr,updatedTs"`

	// Related fields
	DatabaseId int `jsonapi:"attr,databaseId"`
	// Do not return this to the client since the client always has the database context and fetching the
	// database object and all its own related objects is a bit expensive.
	Database *Database

	// Domain specific fields
	Name                    string               `jsonapi:"attr,name"`
	Status                  BackupStatus         `jsonapi:"attr,status"`
	Type                    BackupType           `jsonapi:"attr,type"`
	StorageBackend          BackupStorageBackend `jsonapi:"attr,storageBackend"`
	MigrationHistoryVersion string               `jsonapi:"attr,migrationHistoryVersion"`
	Path                    string               `jsonapi:"attr,path"`
	Comment                 string               `jsonapi:"attr,comment"`
}

type BackupCreate struct {
	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	CreatorId int

	// Related fields
	DatabaseId int `jsonapi:"attr,databaseId"`

	// Domain specific fields
	Name                    string               `jsonapi:"attr,name"`
	Status                  BackupStatus         `jsonapi:"attr,status"`
	Type                    BackupType           `jsonapi:"attr,type"`
	StorageBackend          BackupStorageBackend `jsonapi:"attr,storageBackend"`
	MigrationHistoryVersion string               `jsonapi:"attr,migrationHistoryVersion"`
	Path                    string               `jsonapi:"attr,path"`
	Comment                 string               `jsonapi:"attr,comment"`
}

type BackupFind struct {
	ID *int

	// Related fields
	DatabaseId *int

	// Domain specific fields
	Name *string
}

func (find *BackupFind) String() string {
	str, err := json.Marshal(*find)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

type BackupPatch struct {
	ID int

	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	UpdaterId int

	// Domain specific fields
	Status string
}

// RestoreBackup is the message to restore from a backup.
type RestoreBackup struct {
	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.

	// Related fields

	// Domain specific fields
	BackupId int `jsonapi:"attr,backupId"`
}

// BackupSetting is the backup setting for a database.
type BackupSetting struct {
	ID int `jsonapi:"primary,backupSetting"`

	// Standard fields
	CreatorId int
	Creator   *Principal `jsonapi:"attr,creator"`
	CreatedTs int64      `jsonapi:"attr,createdTs"`
	UpdaterId int
	Updater   *Principal `jsonapi:"attr,updater"`
	UpdatedTs int64      `jsonapi:"attr,updatedTs"`

	// Related fields
	DatabaseId int `jsonapi:"attr,databaseId"`
	// Do not return this to the client since the client always has the database context and fetching the
	// database object and all its own related objects is a bit expensive.
	Database *Database

	// Domain specific fields
	Enabled   bool `jsonapi:"attr,enabled"`
	Hour      int  `jsonapi:"attr,hour"`
	DayOfWeek int  `jsonapi:"attr,dayOfWeek"`
}

// BackupSettingFind is the message to get a backup settings.
type BackupSettingFind struct {
	ID *int

	// Related fields
	DatabaseId *int

	// Domain specific fields
}

// BackupSettingUpsert is the message to upsert a backup settings.
// NOTE: We use PATCH for Upsert, this is inspired by https://google.aip.dev/134#patch-and-put
type BackupSettingUpsert struct {
	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	// CreatorId is the ID of the creator.
	UpdaterId int

	// Related fields
	DatabaseId int `jsonapi:"attr,databaseId"`

	// Domain specific fields
	Enabled   bool `jsonapi:"attr,enabled"`
	Hour      int  `jsonapi:"attr,hour"`
	DayOfWeek int  `jsonapi:"attr,dayOfWeek"`
}

// BackupSettingsMatch is the message to find backup settings matching the conditions.
type BackupSettingsMatch struct {
	Hour      int
	DayOfWeek int
}

// BackupService is the backend for backups.
type BackupService interface {
	CreateBackup(ctx context.Context, create *BackupCreate) (*Backup, error)
	FindBackup(ctx context.Context, find *BackupFind) (*Backup, error)
	FindBackupList(ctx context.Context, find *BackupFind) ([]*Backup, error)
	PatchBackup(ctx context.Context, patch *BackupPatch) (*Backup, error)
	FindBackupSetting(ctx context.Context, find *BackupSettingFind) (*BackupSetting, error)
	UpsertBackupSetting(ctx context.Context, upsert *BackupSettingUpsert) (*BackupSetting, error)
	FindBackupSettingsMatch(ctx context.Context, match *BackupSettingsMatch) ([]*BackupSetting, error)
}
