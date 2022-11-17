package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetBalance(userId string) *UserBalance {
	db, err := sql.Open("mysql", "root:1111@tcp(localhost:3306)/balance_schema")
	balance := &UserBalance{}
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM balance where user_id=?", userId)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&balance.UserId, &balance.Balance)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return balance
}

func AddBalance(balance UserBalance) {

	db, err := sql.Open("mysql", "root:1111@tcp(localhost:3306)/balance_schema")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM balance where id=?", balance.UserId)

	if results != nil {
		//тут меняем баланс существующего пользователя
		update, err := db.Query(
			"UPDATE balance SET user_balance +=? WHERE user_id=?",
			balance.Balance, balance.UserId)
		if err != nil {
			panic(err.Error())
		}
		defer update.Close()
	} else {
		//если человека с отправленным id еще нет в базе, добавляем
		insert, err := db.Query(
			"INSERT INTO balance (user_id,user_balance) VALUES (?,?)",
			balance.UserId, balance.Balance)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
	}

}

func ReserveBalance(rbalance Reserve) {
	db, err := sql.Open("mysql", "root:1111@tcp(localhost:3306)/balance_schema")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO reserve(user_id,service_id,purchase_id,price) VALUES (?,?,?,?)",
		rbalance.UserId, rbalance.ServiceId, rbalance.PurchaseId, rbalance.Price)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

//func ConfirmBalance(confirm Confirm) {
//	db, err := sql.Open("mysql", "root:1111@tcp(localhost:3306)/balance_schema")
//
//	if err != nil {
//		panic(err.Error())
//	}
//
//	defer db.Close()
//
//	balance_res, err := db.Query("SELECT balance FROM balance where id=?", confirm.UserId)
//
//	if balance_res >= confirm.Price { //как поменять получаемый тип из строки?
//		insert, err := db.Query(
//			"INSERT INTO confirm(user_id,service_id,purchase_id,price) VALUES (?,?,?,?)",
//			confirm.UserId, confirm.ServiceId, confirm.PurchaseId, confirm.Price)
//		del, err := db.Query(
//			"DELETE FROM reserve WHERE (purchase_id = ?)",
//			confirm.PurchaseId)
//		rep_ins, err := db.Query(
//			"INSERT INTO report(user_id,service_id,purchase_id,price) VALUES (?,?,?,?)",
//			confirm.UserId, confirm.ServiceId, confirm.PurchaseId, confirm.Price)
//
//		defer insert.Close()
//		defer del.Close()
//		defer rep_ins.Close()
//	}
//
//	if err != nil {
//		panic(err.Error())
//	}
//
//}
