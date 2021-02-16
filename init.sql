CREATE TABLE study_group (
id serial PRIMARY KEY,
name TEXT NOT NULL
);

CREATE TABLE subject (
id serial PRIMARY KEY,
name TEXT NOT NULL,
short_name TEXT NOT NULL
);

CREATE TABLE exam_type (
id serial PRIMARY KEY,
type TEXT NOT NULL
);

CREATE TABLE mark (
id serial PRIMARY KEY,
name TEXT NOT NULL,
value TEXT NOT NULL
);

CREATE TABLE student (
id serial PRIMARY KEY ,
surname TEXT NOT NULL,
name TEXT NOT NULL,
second_name TEXT NOT NULL,
study_group_id INT NOT NULL references study_group(id)
);

CREATE TABLE study_plan (
id serial PRIMARY KEY,
subject_id INT references subject(id),
exam_type_id INT references exam_type(id)
);


CREATE TABLE journal (
id serial PRIMARY KEY,
student_id INT references student(id),
study_plan_id INT references study_plan(id),
in_time INT,
count INT,
mark_id INT references mark(id)
);
