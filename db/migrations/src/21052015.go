package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var schema = `
DROP TABLE IF EXISTS book_usage_sessions;
DROP TABLE IF EXISTS book_usage_statistic_items;

CREATE TABLE book_usage_statistic_items (
  id           serial NOT NULL,
  profile_id   integer NOT NULL,
  book_item_id integer NOT NULL,
  created_at   timestamp DEFAULT Now(),
  updated_at   timestamp DEFAULT Now(),
  CONSTRAINT busi_pkey PRIMARY KEY(id)
);


CREATE TABLE book_usage_sessions (
  id         integer CONSTRAINT bus_pk PRIMARY KEY,
  book_usage_statistic_item_id integer NOT NULL REFERENCES book_usage_statistic_items (id) ON DELETE CASCADE,
  duration   integer NOT NULL,
  begin_time timestamp,
  created_at timestamp DEFAULT Now(),
  updated_at timestamp DEFAULT Now()
)`

func main() {
	db, err := sqlx.Connect("postgres", "user=konjoot dbname=reeky sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)
}
