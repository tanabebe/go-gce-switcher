package src

import (
	"context"
	"encoding/json"
	"google.golang.org/api/compute/v1"
	"log"
	"os"
)

var projectId = os.Getenv("GCP_PROJECT")

type Message struct {
	Data []byte `json:"data"`
}

type PayLoad struct {
	Switch string `json:"switch"`
	Target string `json:"target"`
	Zone   string `json:"zone"`
}

func InstanceSwitcher(ctx context.Context, msg Message) error {

	var payLoad PayLoad

	err := json.Unmarshal(msg.Data, &payLoad)

	if err != nil {
		log.Printf("[ERROR][%T][MSG]: %v", err, err)
		return nil
	}

	service, err := compute.NewService(ctx)
	is := compute.NewInstancesService(service)
	insList, err := is.List(projectId, payLoad.Zone).Do()
	log.Printf("instance:list: %v", insList)

	log.Printf("[ProjectId:%s][Switch:%s][Target:%s][Zone:%s]", projectId, payLoad.Switch, payLoad.Target, payLoad.Zone)

	switch payLoad.Switch {
	case "start":
		log.Println("instance start")
		_, err = is.Start(projectId, payLoad.Zone, payLoad.Target).Do()
	case "stop":
		log.Println("instance stop")
		_, err = is.Stop(projectId, payLoad.Zone, payLoad.Target).Do()
	}

	if err != nil {
		log.Printf("[ERROR][%T][MSG]: %v", err, err)
		return nil
	}

	return nil
}
