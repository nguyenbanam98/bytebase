-- For testing, we reset data on each run
-- Do not reset bb.auth.secret so that we don't need to re-login after restart
DELETE FROM
    setting
WHERE
    name != 'bb.auth.secret';

DELETE FROM
    repo;

DELETE FROM
    vcs;

DELETE FROM
    bookmark;

DELETE FROM
    inbox;

DELETE FROM
    activity;

DELETE FROM
    issue_subscriber;

DELETE FROM
    issue;

DELETE FROM
    task_run;

DELETE FROM
    task;

DELETE FROM
    stage;

DELETE FROM
    pipeline;

DELETE FROM
    data_source;

DELETE FROM
    idx;

DELETE FROM
    col;

DELETE FROM
    tbl;

DELETE FROM
    backup;

DELETE FROM
    backup_setting;

DELETE FROM
    db;

DELETE FROM
    instance_user;

DELETE FROM
    instance;

DELETE FROM
    environment;

DELETE FROM
    project_webhook;

DELETE FROM
    project_member;

-- Project 1 refers to DEFAULT project which is considered as part of schema
DELETE FROM
    project
WHERE
    id != 1;

DELETE FROM
    member;

-- Principal 1 refers to bytebase system account which is considered as part of schema
DELETE FROM
    principal
WHERE
    id != 1;

-- By default, for every table, we reserve id <=100 for internal user
UPDATE
    sqlite_sequence
SET
    seq = 100;