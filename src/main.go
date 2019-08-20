package src

import (
	"context"
	"encoding/json"
	"google.golang.org/api/compute/v1"
	"log"
	"os"
)

var projectID = os.Getenv("GCP_PROJECT")

// Message is struct
// 受け取ったjsonを格納するための構造体
type Message struct {
	Data []byte `json:"data"`
}

// PayLoad is struct
// PayLoad用に対応するjsonのマッピング用構造体
type PayLoad struct {
	Switch string `json:"switch"`
	Target string `json:"target"`
	Zone   string `json:"zone"`
}

// InstanceSwitcher return error
func InstanceSwitcher(ctx context.Context, msg Message) error {

	var payLoad PayLoad

	err := json.Unmarshal(msg.Data, &payLoad)

	if err != nil {
		log.Printf("[ERROR][%T][MSG]: %v", err, err)
		return nil
	}

	service, err := compute.NewService(ctx)
	is := compute.NewInstancesService(service)
	insList, err := is.List(projectID, payLoad.Zone).Do()
	log.Printf("instance:list: %v", insList)

	log.Printf("[ProjectId:%s][Switch:%s][Target:%s][Zone:%s]", projectID, payLoad.Switch, payLoad.Target, payLoad.Zone)

	switch payLoad.Switch {
	case "start":
		log.Println("instance start")
		_, err = is.Start(projectID, payLoad.Zone, payLoad.Target).Do()
	case "stop":
		log.Println("instance stop")
		_, err = is.Stop(projectID, payLoad.Zone, payLoad.Target).Do()
	}

	if err != nil {
		log.Printf("[ERROR][%T][MSG]: %v", err, err)
		return nil
	}

	return nil
}
