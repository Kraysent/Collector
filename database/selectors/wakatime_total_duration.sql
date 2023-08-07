SELECT
    current_date as time,
    sum(extract ('epoch' from (data ->> 'duration')::float * interval '1 second')) as number
FROM collector.t_event
WHERE ts + interval '3h' BETWEEN current_date AND current_date + interval '1 day';
