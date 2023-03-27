create table if not exists account(
  id bigserial unique primary key,
  holder varchar(100) not null,
  balance bigint default(0) not null
);

create table if not exists entry (
  id bigserial unique primary key,
  account_id bigint references account(id) not null,
  amount bigint not null,
  note varchar(200) not null
);

create table if not exists tbempty(
  id bigserial unique primary key,
  name varchar(70)  not null
);

 