package postgresql

import (
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/limanmys/go/helpers"
	"github.com/limanmys/go/models"
)

func AddOrUpdateGoEngine(token string, machineID string, ipAddress string, port int) error {
	newData := &models.EngineModel{
		Token:     token,
		MachineID: machineID,
		IPAddress: ipAddress,
		Port:      port,
		Enabled:   true,
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	engine := GetGoEngine(machineID)
	if engine.ID != "" {
		newData.CreatedAt = engine.CreatedAt
		_, err := db.Model(newData).Where("id = ?", engine.ID).Update()
		if err != nil {
			return err
		}
		return nil
	}
	newID, _ := uuid.NewUUID()
	newData.ID = newID.String()
	newData.CreatedAt = time.Now().Format(time.RFC3339)
	_, err := db.Model(newData).Insert()
	if err != nil {
		return err
	}
	return nil
}

func StoreEngineData() {
	key, _ := uuid.NewUUID()
	machineID, err := ioutil.ReadFile("/sys/class/dmi/id/product_uuid")
	if err != nil {
		panic(err.Error())
	}
	localIP := helpers.GetLocalIP()
	if localIP == "" {
		panic("Cannot find local IP Address, please add CURRENT_IP to configuration.")
	} else {
		log.Printf("Current IP Address %v\n", localIP)
	}
	machineIDSTR := strings.TrimSpace(strings.ToUpper(string(machineID)))
	err = AddOrUpdateGoEngine(key.String(), machineIDSTR, localIP, 5454)
	if err != nil {
		log.Panic(err.Error())
	}
}