package postgres

import (
	"log"
	"github.com/konstantinfarrell/go-example-lambda"
	"github.com/konstantinfarrell/go-example-lambda/pkg/util/postgres"
)

func CreateFile(d *postgres.Database, f *golx.File) (*golx.File, error) {
	spname := "create_file"

	_, err := d.Call(spname, f.Filename, f.Path, f.Permissions, f.Created, f.Modified, f.FileId, f.Data, f.Received)
	if err != nil {
		log.Printf("Error while calling create_file: %s", err)
		return nil, err
	}
	return f, nil
}

func DeleteFile(d *postgres.Database, f *golx.File) (error) {
	spname := "delete_file"

	_, err := d.Call(spname, f.FileId)
	if err != nil {
		log.Printf("Error while calling delete_file: %s", err)
		return err
	}
	return nil
}