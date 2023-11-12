package generator

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/amirrezawh/ocserv-manager/config"
	db "github.com/amirrezawh/ocserv-manager/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func dataGenerator(cfg *config.GeneralConfig, resetDay bool) {

	jsonFile, err := os.Open(cfg.JsonPath)

	if err != nil {
		fmt.Println("Error occured during openning general.json")
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)
	var juser Data

	json.Unmarshal(byteValue, &juser)

	// postgres connection

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db.ConnectToPostgres(cfg),
	}), &gorm.Config{})

	sqldb, err := gormDB.DB()
	if err != nil {
		panic(err)
	}
	sqldb.SetConnMaxLifetime(10 * time.Second)

	finalData := make(map[string]uint64)
	var users []db.Users

	for u := 0; u < len(juser); u++ {
	    if juser[u].Username != "(none)" && juser[u].RX != "" && juser[u].TX != "" {
		rx_int, err := strconv.ParseUint(juser[u].RX, 10, 64)
		if err != nil {
			panic(err)
		}

		tx_int, err := strconv.ParseUint(juser[u].TX, 10, 64)
		if err != nil {
			panic(err)
		}

		// Check reset Day
		currentDate := time.Now()
		if currentDate.Day() == 1 && !resetDay {
			finalData[juser[u].Username] = 0
			resetDay = true
		}
		if currentDate.Day() != 1 {
			finalData[juser[u].Username] += rx_int
			finalData[juser[u].Username] += tx_int
			resetDay = false
		}
	}

	}

	for username, usage := range finalData {
		if query := gormDB.Where("username = ?", username).Find(&users); query.Error != nil {
			fmt.Println("Error on running query", query.Error)
			return
		}
	 	fmt.Println(finalData)
		if len(users) > 0 {
			for _, record := range users {
				if usage > record.RX_TX_BYTE {
                	extraUsage := usage - record.RX_TX_BYTE
					gormDB.Model(&db.Users{}).Where("username = ?", username).Update(
						"rx_tx_byte", gorm.Expr("rx_tx_byte + ?", extraUsage))
				}else {
					gormDB.Model(&db.Users{}).Where("username = ?", username).Update(
						"rx_tx_byte", gorm.Expr("rx_tx_byte + ?", usage))
				}

				gormDB.Model(&db.Users{}).Where("username = ?", username).Update(
				"rx_tx", prettyByteSize(record.RX_TX_BYTE))
			}
			//gormDB.Model(&db.Users{}).Where("username = ?", username).Update(
			//	"rx_tx", prettyByteSize(usage))

		} else {
			gormDB.Create(&db.Users{
				Username:   username,
				RX_TX_BYTE: usage,
				RX_TX:      prettyByteSize(usage),
				LIMIT:      21474836480,
				Active:     true,
			})
		}
		// Check limits
		if len(users) > 0 && users[0].RX_TX_BYTE >= users[0].LIMIT && users[0].Active == true {
			fmt.Printf("User %s reached the limit\n", users[0].Username)
			fmt.Printf("Locking user %s...\n", users[0].Username)
			gormDB.Model(&db.Users{}).Where("username = ?", users[0].Username).Update(
				"active", false)
			// Lock Function
			lockUser(cfg.PasswordFilePath, users[0].Username)

		}

		if len(users) > 0 && users[0].RX_TX_BYTE < users[0].LIMIT && users[0].Active == false {
			fmt.Printf("Unlocking user %s...\n", users[0].Username)
			gormDB.Model(&db.Users{}).Where("username = ?", users[0].Username).Update(
				"active", true)
			// Unlock Function
			unlockUser(cfg.PasswordFilePath, users[0].Username)
		}

	}

	jsonFile.Close()

}

func prettyByteSize(b uint64) string {
	bf := float64(b)
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"} {
		if math.Abs(bf) < 1024.0 {
			return fmt.Sprintf("%3.1f%sB", bf, unit)
		}
		bf /= 1024.0
	}
	return fmt.Sprintf("%.1fYiB", bf)
}

func Interval(cfg *config.GeneralConfig) {
	resetDay := false

	for {
		dataGenerator(cfg, resetDay)
		time.Sleep(2 * time.Second)
		jsonGenerator()

	}
}
