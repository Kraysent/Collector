CREATE SCHEMA IF NOT EXISTS collector;
CREATE TABLE IF NOT EXISTS collector.t_event (
    ts timestamp not null,
    source text not null,
    data json,
    primary key (ts)
);
