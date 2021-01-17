CREATE TABLE users
(
    id       uuid         not null unique,

    name     varchar(250) not null,
    surname  varchar(250) not null,
    birthday timestamp default null,

    email    varchar(250) not null,
    password_hash varchar (250) not null,

    role varchar (250) not null
);

CREATE TABLE resumes
(
    id       uuid         not null unique,
    user_id uuid references users (id) on delete cascade      not null,

    programmer_type varchar(250) not null,
    programmer_level varchar(250) not null,
    programmer_language varchar(250) not null,
    extra_skills text
);

CREATE TABLE work_experience
(
    id uuid not null unique ,
    resume_id uuid references resumes (id) on delete cascade      not null,
    company_name varchar(100),
    from timestamp,
    to timestamp,
    content varchar(150)
);

CREATE TABLE vacancies
(
    id       uuid         not null unique,
    user_id uuid references users (id) on delete cascade      not null,

    programmer_type varchar(250) not null,
    programmer_level varchar(250) not null,
    programmer_language varchar(250) not null,
    company_name varchar(250),
    team text,
    salary_from varchar(15) not null,
    salary_to varchar(15) not null,
    extra_benefits varchar(100)
);

