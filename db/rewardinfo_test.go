package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func initTable(t *testing.T) error {
	err := Init()
	if err != nil {
		t.Fatal("open database failed, err ", err)
		return err
	}
	CreateRewardInfoTable()
	return nil
}

func TestNodeName_Create(t *testing.T) {
	initTable(t)

	datas := make([]*RewardInfo, 0)
	list, err := ioutil.ReadFile("nodelist.json")
	if err != nil {
		fmt.Println("open nodelist failed.")
		return
	}
	err = json.Unmarshal(list, &datas)
	if err != nil {
		fmt.Printf("unmarshal nodelist failed, err:%s.\n", err)
		return
	}

	for _, d := range datas {
		if err := d.Create(); err != nil {
			t.Fatal("create failed, err ", err)
		} else {
			fmt.Println("create new passed ")
		}
	}
	//gorm.DropTable(&RewardInfo{})
}

func TestNodeName_GetAllInfo(t *testing.T) {
	var d = &RewardInfo{}
	all, err := d.GetAllInfo()
	if err != nil {
		fmt.Printf("get all failed.\n")
	}
	for _, info := range all {
		d, _ := json.Marshal(info)
		fmt.Printf("get info ", string(d))
	}
}

func TestNodeName_CloseDB(t *testing.T) {
	//GetORM().
}
