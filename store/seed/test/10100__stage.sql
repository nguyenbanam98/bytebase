-- Stage for Pipeline 9001 "Hello world"
INSERT INTO
    stage (
        id,
        creator_id,
        updater_id,
        pipeline_id,
        environment_id,
        name
    )
VALUES
    (
        10001,
        101,
        101,
        9001,
        5004,
        'Prod'
    );

-- Stage for Pipeline 9002 simulating webhook push event
INSERT INTO
    stage (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        pipeline_id,
        environment_id,
        name
    )
VALUES
    (
        10002,
        1,
        1624873710,
        1,
        1624873710,
        9002,
        5001,
        'Dev'
    );

INSERT INTO
    stage (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        pipeline_id,
        environment_id,
        name
    )
VALUES
    (
        10003,
        1,
        1624873710,
        1,
        1624873710,
        9002,
        5002,
        'Integration'
    );

INSERT INTO
    stage (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        pipeline_id,
        environment_id,
        name
    )
VALUES
    (
        10004,
        1,
        1624873710,
        1,
        1624873710,
        9002,
        5003,
        'Staging'
    );

INSERT INTO
    stage (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        pipeline_id,
        environment_id,
        name
    )
VALUES
    (
        10005,
        1,
        1624873710,
        1,
        1624873710,
        9002,
        5004,
        'Prod'
    );

-- Stage for Pipeline 9003 create table
INSERT INTO
    stage (
        id,
        creator_id,
        updater_id,
        pipeline_id,
        environment_id,
        name
    )
VALUES
    (
        10006,
        103,
        103,
        9003,
        5001,
        'Dev'
    );

INSERT INTO
    stage (
        id,
        creator_id,
        updater_id,
        pipeline_id,
        environment_id,
        name
    )
VALUES
    (
        10007,
        103,
        103,
        9003,
        5002,
        'Integration'
    );

-- Stage for Pipeline 9004 simulating webhook push event
INSERT INTO
    stage (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        pipeline_id,
        environment_id,
        name
    )
VALUES
    (
        10008,
        1,
        1624865387,
        1,
        1624865387,
        9004,
        5001,
        'Dev'
    );

-- Stage for Pipeline 9005 simulating webhook push event
INSERT INTO
    stage (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        pipeline_id,
        environment_id,
        name
    )
VALUES
    (
        10009,
        1,
        1624866790,
        1,
        1624866790,
        9005,
        5002,
        'Integration'
    );

-- Stage for Pipeline 9006 simulating webhook push event
INSERT INTO
    stage (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        pipeline_id,
        environment_id,
        name
    )
VALUES
    (
        10010,
        1,
        1624868407,
        1,
        1624868407,
        9006,
        5003,
        'Staging'
    );

-- Stage for Pipeline 9007 simulating webhook push event
INSERT INTO
    stage (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        pipeline_id,
        environment_id,
        name
    )
VALUES
    (
        10011,
        1,
        1624868680,
        1,
        1624868680,
        9007,
        5004,
        'Prod'
    );

-- Stage for Pipeline 9008 simulating webhook push event
INSERT INTO
    stage (
        id,
        creator_id,
        created_ts,
        updater_id,
        updated_ts,
        pipeline_id,
        environment_id,
        name
    )
VALUES
    (
        10012,
        1,
        1624869944,
        1,
        1624869944,
        9008,
        5001,
        'Dev'
    );