create table if not exists application (
    id varchar(20) not null,
    project_id varchar(20) not null references project(id),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    name varchar(50) not null,
    description text not null,
    primary key (id, project_id)
);

create table if not exists endpoint (
    id varchar(20) not null,
    application_id varchar(20) not null,
    project_id varchar(20) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    method varchar(10) not null,
    path varchar(255) not null,
    name varchar(255),
    description text,
    settings jsonb not null default '{}',
    foreign key (application_id, project_id) references application(id, project_id),
    primary key (id, application_id, project_id)
);

create table if not exists workflow (
    version int not null,
    endpoint_id varchar(20) not null,
    application_id varchar(20) not null,
    project_id varchar(20) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    data jsonb not null default '{}',
    foreign key (endpoint_id, application_id, project_id) references endpoint(id, application_id, project_id),
    primary key (version, endpoint_id, application_id, project_id)
);