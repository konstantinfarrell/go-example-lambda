package file

import (
	"encoding/json"
	"log"

	"github.com/konstantinfarrell/go-example-lambda"
	"github.com/konstantinfarrell/go-example-lambda/pkg/util/config"
	pg "github.com/konstantinfarrell/go-example-lambda/pkg/util/postgres"
	pgf "github.com/konstantinfarrell/go-example-lambda/pkg/file/platform/postgres"

)

type FileHandler struct {
	Config 	*config.Config
	DB		*pg.Database
}

type KinesisData struct {
	Command		string
	Data		string
} 

func New() (*FileHandler, error) {
	var h FileHandler
	conf, err := config.LoadFromEnvVar()
	h.Config = conf
	if err != nil {
		return nil, err
	}

	log.Printf("Initializing DB connection from config")
	db, err := pg.New(h.Config)
	h.DB = db
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (h *FileHandler) Handle(data string) (error){
	var kinesisData KinesisData
	err := json.Unmarshal([]byte(data), &kinesisData)
	if err != nil {
		log.Printf("Unable to decode json payload: %s", err)
		return err
	}

	file := new(golx.File)
	file, err = file.FromJson(kinesisData.Data)
	if err != nil {
		log.Printf("Unable to bind json to file object: %s", err)
		return err
	}

	switch kinesisData.Command {
	case "create":
		_, err = pgf.CreateFile(h.DB, file)
	case "delete":
		err = pgf.DeleteFile(h.DB, file)
	}
	if err != nil {
		return err
	}
	
	return nil
}