package store

import (
	"context"
	"fmt"
	"strings"

	"github.com/bytebase/bytebase"
	"github.com/bytebase/bytebase/api"
	"go.uber.org/zap"
)

var (
	_ api.BackupService = (*BackupService)(nil)
)

// BackupService represents a service for managing backup.
type BackupService struct {
	l  *zap.Logger
	db *DB
}

// NewBackupService returns a new instance of BackupService.
func NewBackupService(logger *zap.Logger, db *DB) *BackupService {
	return &BackupService{l: logger, db: db}
}

// CreateBackup creates a new backup.
func (s *BackupService) CreateBackup(ctx context.Context, create *api.BackupCreate) (*api.Backup, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	backup, err := s.createBackup(ctx, tx, create)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, FormatError(err)
	}

	return backup, nil
}

// FindBackup retrieves a single backup based on find.
// Returns ENOTFOUND if no matching record.
// Returns ECONFLICT if finding more than 1 matching records.
func (s *BackupService) FindBackup(ctx context.Context, find *api.BackupFind) (*api.Backup, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	list, err := s.findBackupList(ctx, tx, find)
	if err != nil {
		return nil, err
	} else if len(list) == 0 {
		return nil, &bytebase.Error{Code: bytebase.ENOTFOUND, Message: fmt.Sprintf("backup not found: %+v", find)}
	} else if len(list) > 1 {
		return nil, &bytebase.Error{Code: bytebase.ECONFLICT, Message: fmt.Sprintf("found %d backups with filter %+v, expect 1. ", len(list), find)}
	}

	return list[0], nil
}

// FindBackupList retrieves a list of backups based on find.
func (s *BackupService) FindBackupList(ctx context.Context, find *api.BackupFind) ([]*api.Backup, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	list, err := s.findBackupList(ctx, tx, find)
	if err != nil {
		return []*api.Backup{}, err
	}

	return list, nil
}

// PatchBackup updates an existing backup by ID.
// Returns ENOTFOUND if backup does not exist.
func (s *BackupService) PatchBackup(ctx context.Context, patch *api.BackupPatch) (*api.Backup, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	backup, err := s.patchBackup(ctx, tx, patch)
	if err != nil {
		return nil, FormatError(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, FormatError(err)
	}

	return backup, nil
}

// createBackup creates a new backup.
func (s *BackupService) createBackup(ctx context.Context, tx *Tx, create *api.BackupCreate) (*api.Backup, error) {
	// Insert row into backup.
	row, err := tx.QueryContext(ctx, `
		INSERT INTO backup (
			creator_id,
			updater_id,
			database_id,
			name,
			`+"`status`,"+`
			`+"`type`,"+`
			storage_backend,
			migration_history_version,
			path,
			comment
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id, creator_id, created_ts, updater_id, updated_ts, database_id, name, `+"`status`,"+` `+"`type`, storage_backend, migration_history_version, path, comment"+`
	`,
		create.CreatorId,
		create.CreatorId,
		create.DatabaseId,
		create.Name,
		create.Status,
		create.Type,
		create.StorageBackend,
		create.MigrationHistoryVersion,
		create.Path,
		create.Comment,
	)

	if err != nil {
		return nil, FormatError(err)
	}
	defer row.Close()

	row.Next()
	var backup api.Backup
	if err := row.Scan(
		&backup.ID,
		&backup.CreatorId,
		&backup.CreatedTs,
		&backup.UpdaterId,
		&backup.UpdatedTs,
		&backup.DatabaseId,
		&backup.Name,
		&backup.Status,
		&backup.Type,
		&backup.StorageBackend,
		&backup.MigrationHistoryVersion,
		&backup.Path,
		&backup.Comment,
	); err != nil {
		return nil, FormatError(err)
	}

	return &backup, nil
}

func (s *BackupService) findBackupList(ctx context.Context, tx *Tx, find *api.BackupFind) (_ []*api.Backup, err error) {
	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := find.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}
	if v := find.DatabaseId; v != nil {
		where, args = append(where, "database_id = ?"), append(args, *v)
	}
	if v := find.Name; v != nil {
		where, args = append(where, "name = ?"), append(args, *v)
	}

	rows, err := tx.QueryContext(ctx, `
		SELECT
		    id,
		    creator_id,
		    created_ts,
		    updater_id,
		    updated_ts,
			database_id,
		    name,
			`+"`status`,"+`
			`+"`type`,"+`
			storage_backend,
			migration_history_version,
			path,
			comment
		FROM backup
		WHERE `+strings.Join(where, " AND "),
		args...,
	)
	if err != nil {
		return nil, FormatError(err)
	}
	defer rows.Close()

	// Iterate over result set and deserialize rows into list.
	list := make([]*api.Backup, 0)
	for rows.Next() {
		var backup api.Backup
		if err := rows.Scan(
			&backup.ID,
			&backup.CreatorId,
			&backup.CreatedTs,
			&backup.UpdaterId,
			&backup.UpdatedTs,
			&backup.DatabaseId,
			&backup.Name,
			&backup.Status,
			&backup.Type,
			&backup.StorageBackend,
			&backup.MigrationHistoryVersion,
			&backup.Path,
			&backup.Comment,
		); err != nil {
			return nil, FormatError(err)
		}

		list = append(list, &backup)
	}
	if err := rows.Err(); err != nil {
		return nil, FormatError(err)
	}

	return list, nil
}

// patchBackup updates a backup by ID. Returns the new state of the backup after update.
func (s *BackupService) patchBackup(ctx context.Context, tx *Tx, patch *api.BackupPatch) (*api.Backup, error) {
	// Build UPDATE clause.
	set, args := []string{"updater_id = ?"}, []interface{}{patch.UpdaterId}
	set, args = append(set, "status = ?"), append(args, patch.Status)

	args = append(args, patch.ID)

	// Execute update query with RETURNING.
	row, err := tx.QueryContext(ctx, `
		UPDATE backup
		SET `+strings.Join(set, ", ")+`
		WHERE id = ?
		RETURNING id, creator_id, created_ts, updater_id, updated_ts, database_id, name, `+"`status`,"+` `+"`type`, storage_backend, migration_history_version, path, comment"+`
	`,
		args...,
	)
	if err != nil {
		return nil, FormatError(err)
	}
	defer row.Close()

	if row.Next() {
		var backup api.Backup
		if err := row.Scan(
			&backup.ID,
			&backup.CreatorId,
			&backup.CreatedTs,
			&backup.UpdaterId,
			&backup.UpdatedTs,
			&backup.DatabaseId,
			&backup.Name,
			&backup.Status,
			&backup.Type,
			&backup.StorageBackend,
			&backup.MigrationHistoryVersion,
			&backup.Path,
			&backup.Comment,
		); err != nil {
			return nil, FormatError(err)
		}
		return &backup, nil
	}

	return nil, &bytebase.Error{Code: bytebase.ENOTFOUND, Message: fmt.Sprintf("backup ID not found: %d", patch.ID)}
}

// FindBackupSetting finds the backup setting for a database.
// Returns ENOTFOUND if no matching record.
// Returns ECONFLICT if finding more than 1 matching records.
func (s *BackupService) FindBackupSetting(ctx context.Context, find *api.BackupSettingFind) (*api.BackupSetting, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	list, err := s.findBackupSetting(ctx, tx, find)
	if err != nil {
		return nil, err
	} else if len(list) == 0 {
		return nil, &bytebase.Error{Code: bytebase.ENOTFOUND, Message: fmt.Sprintf("backup setting not found: %+v", find)}
	} else if len(list) > 1 {
		return nil, &bytebase.Error{Code: bytebase.ECONFLICT, Message: fmt.Sprintf("found %d backup settings with filter %+v, expect 1. ", len(list), find)}
	}

	return list[0], nil
}

func (s *BackupService) findBackupSetting(ctx context.Context, tx *Tx, find *api.BackupSettingFind) (_ []*api.BackupSetting, err error) {
	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := find.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}
	if v := find.DatabaseId; v != nil {
		where, args = append(where, "database_id = ?"), append(args, *v)
	}

	rows, err := tx.QueryContext(ctx, `
		SELECT
			id,
			creator_id,
			created_ts,
			updater_id,
			updated_ts,
			database_id,
			enabled,
			hour,
			day_of_week
		FROM backup_setting
		WHERE `+strings.Join(where, " AND "),
		args...,
	)
	if err != nil {
		return nil, FormatError(err)
	}
	defer rows.Close()

	// Iterate over result set and deserialize rows into list.
	list := make([]*api.BackupSetting, 0)
	for rows.Next() {
		var backupSetting api.BackupSetting
		if err := rows.Scan(
			&backupSetting.ID,
			&backupSetting.CreatorId,
			&backupSetting.CreatedTs,
			&backupSetting.UpdaterId,
			&backupSetting.UpdatedTs,
			&backupSetting.DatabaseId,
			&backupSetting.Enabled,
			&backupSetting.Hour,
			&backupSetting.DayOfWeek,
		); err != nil {
			return nil, FormatError(err)
		}

		list = append(list, &backupSetting)
	}
	if err := rows.Err(); err != nil {
		return nil, FormatError(err)
	}

	return list, nil
}

// UpsertBackupSetting sets the backup settings for a database.
func (s *BackupService) UpsertBackupSetting(ctx context.Context, upsert *api.BackupSettingUpsert) (*api.BackupSetting, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	backup, err := s.upsertBackupSetting(ctx, tx, upsert)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, FormatError(err)
	}

	return backup, nil
}

// upsertBackupSetting updates an existing backup setting.
func (s *BackupService) upsertBackupSetting(ctx context.Context, tx *Tx, upsert *api.BackupSettingUpsert) (*api.BackupSetting, error) {
	// Upsert row into backup_setting.
	row, err := tx.QueryContext(ctx, `
		INSERT INTO backup_setting (
			creator_id,
			updater_id,
			database_id,
			`+"`enabled`,"+`
			hour,
			day_of_week
		)
		VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT(database_id) DO UPDATE SET
				enabled = excluded.enabled,
				hour = excluded.hour,
				day_of_week = excluded.day_of_week
		RETURNING id, creator_id, created_ts, updater_id, updated_ts, database_id, `+"`enabled`,"+` `+"hour, day_of_week"+`
		`,
		upsert.UpdaterId,
		upsert.UpdaterId,
		upsert.DatabaseId,
		upsert.Enabled,
		upsert.Hour,
		upsert.DayOfWeek,
	)

	if err != nil {
		return nil, FormatError(err)
	}
	defer row.Close()

	row.Next()
	var backupSetting api.BackupSetting
	if err := row.Scan(
		&backupSetting.ID,
		&backupSetting.CreatorId,
		&backupSetting.CreatedTs,
		&backupSetting.UpdaterId,
		&backupSetting.UpdatedTs,
		&backupSetting.DatabaseId,
		&backupSetting.Enabled,
		&backupSetting.Hour,
		&backupSetting.DayOfWeek,
	); err != nil {
		return nil, FormatError(err)
	}

	return &backupSetting, nil
}

// FindBackupSettingsMatch retrieves a list of backup settings based on match condition.
func (s *BackupService) FindBackupSettingsMatch(ctx context.Context, match *api.BackupSettingsMatch) ([]*api.BackupSetting, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FormatError(err)
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
		SELECT
			id,
			creator_id,
			created_ts,
			updater_id,
			updated_ts,
			database_id,
			enabled,
			hour,
			day_of_week
		FROM backup_setting
		WHERE
			enabled = 1
			AND (
				(hour = ? AND day_of_week = ?)
				OR
				(hour = ? AND day_of_week = -1)
				OR
				(hour = -1 AND day_of_week = ?)
			)
		`,
		match.Hour, match.DayOfWeek, match.Hour, match.DayOfWeek,
	)
	if err != nil {
		return nil, FormatError(err)
	}
	defer rows.Close()

	// Iterate over result set and deserialize rows into list.
	list := make([]*api.BackupSetting, 0)
	for rows.Next() {
		var backupSetting api.BackupSetting
		if err := rows.Scan(
			&backupSetting.ID,
			&backupSetting.CreatorId,
			&backupSetting.CreatedTs,
			&backupSetting.UpdaterId,
			&backupSetting.UpdatedTs,
			&backupSetting.DatabaseId,
			&backupSetting.Enabled,
			&backupSetting.Hour,
			&backupSetting.DayOfWeek,
		); err != nil {
			return nil, FormatError(err)
		}

		list = append(list, &backupSetting)
	}
	if err := rows.Err(); err != nil {
		return nil, FormatError(err)
	}

	return list, nil
}
