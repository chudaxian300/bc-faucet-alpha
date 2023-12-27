package service

import (
	"api/db"
	"api/entity"
	"fmt"
)

type TransferInfo entity.TransferInfo_XieLinXuan

var dbControl = db.Db

func (transferInfo *TransferInfo) AddTransferInfo() error {
	_, err := dbControl.Exec(`INSERT INTO transfer_info (to_address,award,get_time) VALUES (?,?,?)`,
		transferInfo.ToAddress,
		transferInfo.Award,
		transferInfo.GetTime)
	if err != nil {
		return err
	}
	return nil
}

func (transferInfo *TransferInfo) IsExist() (int, error) {
	var id int
	result, err := dbControl.Query(`SELECT id FROM transfer_info where to_address = ?;`, transferInfo.ToAddress)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	for result.Next() {
		err = result.Scan(&id)
	}
	if id == 0 {
		return -1, nil
	}
	return id, nil
}

func (transferInfo *TransferInfo) UpdateInfo() error {
	_, err := dbControl.Exec(`UPDATE transfer_info SET award = ?,get_time=? WHERE id=?`,
		transferInfo.Award,
		transferInfo.GetTime,
		transferInfo.ID)
	if err != nil {
		return err
	}
	return nil
}
