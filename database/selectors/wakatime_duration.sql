SELECT
    ts as time,
    1 as active
FROM
    collector.t_event
WHERE
    source = 'wakatime'
UNION ALL
SELECT
    ts + (data ->> 'duration')::float * interval '1 second' as time,
    0 as active
FROM
    collector.t_event
WHERE
    source = 'wakatime';
