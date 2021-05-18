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

func connStringFromConfig(conf *config.Config) string {
	// postgresql://user:password@url:port/dbname
	base := "postgresql://%s:%s@%s:%s/%s"
	c := conf.DB
	return fmt.Sprintf(base, c.User, c.Pass, c.Addr, strconv.Itoa(c.Port), c.Name)
}

func (d *Database) Call(sp string, args ...interface{}) (*[]interface{}, error){
	log.Printf("Call sp %s called", sp)
	query := formatCall(sp, args)
	result, err := d.Conn.Exec(query)
	if err != nil {
		log.Printf("Error calling sp: %s", err)
		return nil, err
	}
	var results []interface{}

	return &results, nil
}

func formatCall(sp string, args ...interface{}) (string) {
	var query, q_args string

	for _, arg := range args {
		if q_args == "" {
			q_args = fmt.Sprintf("%v", arg)
		} else {
			q_args = fmt.Sprintf("%v, %v", q_args, arg)
		}
	}
	query = fmt.Sprintf("call %s(%s);", sp, q_args)
	return query
}