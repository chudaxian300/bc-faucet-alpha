package service

import (
	"api/entity"
	"fmt"
)

type TransferLog entity.TransferLog_XieLinXuan

func (transferLog *TransferLog) AddLog() error {
	_, err := dbControl.Exec(`INSERT INTO transfer_log VALUES (CURDATE(),?,?)`,
		transferLog.OutNum,
		transferLog.OutCount)
	if err != nil {
		return err
	}
	return nil
}

func (transferLog *TransferLog) UpdateLog() error {
	_, err := dbControl.Exec(`UPDATE transfer_log SET out_num = out_num + ?, out_count = out_count + 1 WHERE time = (SELECT CURDATE());`,
		transferLog.OutNum)
	if err != nil {
		return err
	}
	return nil
}

func (transferLog *TransferLog) IsExistByTime() (bool, error) {
	result, err := dbControl.Query(`SELECT * FROM transfer_log WHERE time = (SELECT CURDATE());`)
	if err != nil {
		return false, err
	}
	return result.Next(), nil
}

func (transferLog *TransferLog) GetOutNumInNow() (float64, error) {
	var num float64
	err := dbControl.QueryRow(`SELECT out_num FROM transfer_log WHERE time = (SELECT CURDATE());`).Scan(&num)
	if err != nil {
		fmt.Println(err)
		return  0, err
	}
	return num, nil
}

func (transferLog *TransferLog) GetOutCountInNow() (float64, error) {
	var num float64
	err := dbControl.QueryRow(`SELECT out_count FROM transfer_log WHERE time = (SELECT CURDATE());`).Scan(&num)
	if err != nil {
		return  -1, err
	}
	return num, nil
}