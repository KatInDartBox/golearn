create table if not exists student(
  id serial unique primary key,
  first_name varchar(100),
  last_name varchar(100),
  email varchar(200),
  gender varchar(70),
  phone varchar(50),
  balance decimal(14,4)
);

create table if not exists teacher (
  id serial unique primary key,
  first_name varchar(100),
  last_name varchar(100),
  email varchar(200),
  gender varchar(70),
  phone varchar(50),
  balance decimal(14,4)
);

create table if not exists subject (
	id serial unique primary key,
	teacher_id int references teacher(id),
	title VARCHAR(50)
);

create table if not exists subject_student (
	student_id int references student(id),
	subject_id int references subject(id)
);

create table if not exists tbempty(
  id serial unique primary key,
  name varchar(70) 
);

 