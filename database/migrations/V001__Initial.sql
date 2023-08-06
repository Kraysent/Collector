CREATE SCHEMA IF NOT EXISTS collector;
CREATE TABLE IF NOT EXISTS collector.t_event (
    ts timestamp not null,
    source text not null,
    data jsonb,
    primary key (ts)
);
