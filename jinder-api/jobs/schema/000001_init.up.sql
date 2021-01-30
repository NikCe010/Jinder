CREATE TABLE users
(
    id       uuid         not null unique,

    name     varchar(250) not null,
    surname  varchar(250) not null,
    birthday timestamp with time zone default null,

    email    varchar(250) not null,
    password_hash varchar (250) not null,

    role int not null
);

CREATE TABLE resumes
(
    id       uuid         not null unique,
    user_id uuid references users (id) on delete cascade      not null,

    programmer_type int not null,
    programmer_level int not null,
    programmer_language int not null
);

CREATE TABLE work_experience
(
    id uuid not null unique ,
    resume_id uuid references resumes (id) on delete cascade      not null,
    company_name varchar(100),
    experience_from timestamp with time zone,
    experience_to timestamp with time zone,
    extra_content varchar(150)
);

CREATE TABLE vacancies
(
    id       uuid         not null unique,
    user_id uuid references users (id) on delete cascade      not null,

    programmer_type int not null,
    programmer_level int not null,
    programmer_language int not null,
    company_name varchar(250),
    salary_from varchar(15) not null,
    salary_to varchar(15) not null,
    extra_benefits varchar(100)
);

