package postgres

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/go-pg/pg/v9"
	_ "github.com/lib/pq"

	"github.com/konstantinfarrell/go-example-lambda"
	"github.com/konstantinfarrell/go-example-lambda/pkg/util/config"
)

type Database struct {
	Conn	*pg.DB
}

func New(conf *config.Config) (*Database, error) {
	options := connOptionsFromConfig(conf)
	db := pg.Connect(options)
	return &Database{ Conn: db }, nil
}

func connOptionsFromConfig(conf *config.Config) *pg.Options {
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

func (d *Database) Call(hasReturn bool, files *[]golx.File, sp string, args ...interface{}) (*[]golx.File, error){
	log.Printf("Call sp %s called", sp)
	query := formatCall(hasReturn, sp, args...)
	log.Printf("Query: %s", query)
	_, err := d.Conn.Query(files, query)
	if err != nil {
		log.Printf("Error calling sp: %s", err)
		return nil, err
	}
	log.Printf("Files: %s", files)
	return files, nil
}

func formatCall(hasReturn bool, sp string, args ...interface{}) (string) {
	var query, q_args string
	for _, arg := range args {

		fmtStr := "'%v'"
		if arg == "" {
			arg = "null"
			fmtStr = "%v"
		}

		if q_args == "" {
			q_args = fmt.Sprintf(fmtStr, arg)
		} else {
			fmtStr = fmt.Sprintf("%v, %v", "%v", fmtStr)
			q_args = fmt.Sprintf(fmtStr, q_args, arg)
		}
		log.Printf("%s, %s", arg, reflect.TypeOf(arg))
	}
	if hasReturn {
		query = fmt.Sprintf("select * from %s(%s)", sp, q_args)
	} else {
		query = fmt.Sprintf("call %s(%s);", sp, q_args)
	}
	return query
}