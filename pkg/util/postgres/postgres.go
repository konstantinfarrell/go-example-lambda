package postgres

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-pg/pg/v9"
	_ "github.com/lib/pq"

	"github.com/konstantinfarrell/go-example-lambda/pkg/util/config"
)

type Database struct {
	Conn	*pg.DB
}

func New(conf *config.Config) (*Database, error) {
	connStr := connStringFromConfig(conf)
	u, err := pg.ParseURL(connStr)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(u)
	return &Database{ Conn: db }, nil
}

func connOptionsFromConfig(conf *config.Configuration) *pg.Options {
	c := conf.DB
	port := strconv.Itoa(c.Port)
	addr := fmt.Sprintf("%s:%s", c.Addr, port)
	return &pg.Options{
		Addr: addr,
		User: c.User,
		Password: c.Pass,
		Database: c.Name,
	}
}

func (d *Database) Call(hasReturn bool, files *[]golx.File, sp string, args ...interface{}) (*[]gox.File, error){
	log.Printf("Call sp %s called", sp)
	query := formatCall(hasReturn, sp, args)
	_, err := d.Conn.Query(files, query)
	if err != nil {
		log.Printf("Error calling sp: %s", err)
		return nil, err
	}
	return files, nil
}

func formatCall(hasReturn bool, sp string, args ...interface{}) (string) {
	var query, q_args string

	for _, arg := range args {
		if q_args == "" {
			q_args = fmt.Sprintf("%v", arg)
		} else {
			q_args = fmt.Sprintf("%v, %v", q_args, arg)
		}
	}
	if q_args == "[]" {
		q_args = ""
	}
	if hasReturn {
		query = fmt.Sprintf("select * from %s(%s)", sp, q_args)
	} else {
		query = fmt.Sprintf("call %s(%s);", sp, q_args)
	}
	return query
}