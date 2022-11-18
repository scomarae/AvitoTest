package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const userdb = "root"
const passworddb = "1111"
const conn = "@tcp(localhost:3306)"

func IsUserExists(userId string, db *sql.DB) bool {

	results, err := db.Query("SELECT * FROM balance where user_id=?", userId)

	if err != nil {
		log.Fatal(err)
	}

	if results != nil {
		return true
	}
	return false
}

func GetBalance(userId string) *UserBalance {
	db, err := sql.Open("mysql", userdb+":"+passworddb+conn+"/balance_schema")
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

func AccrualMoneyToBalance(accrual AccrualMoney) {

	db, err := sql.Open("mysql", userdb+":"+passworddb+conn+"/balance_schema")

	if err != nil {
		panic(err.Error())
	}

	if IsUserExists(accrual.UserId, db) {
		//если человека с отправленным id еще нет в базе, добавляем
		insert, err := db.Query(
			"INSERT INTO balance (user_id,user_balance) VALUES (?,?)",
			accrual.UserId, accrual.Amount)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
	} else {
		//тут к старому значению баланса прибавляем полученное значение
		update, err := db.Query(
			"UPDATE balance SET user_balance = user_balance +? WHERE user_id=?",
			accrual.Amount, accrual.UserId)
		if err != nil {
			panic(err.Error())
		}
		defer update.Close()
	}

}

func ReserveBalance(reserve Reserve) {
	db, err := sql.Open("mysql", userdb+":"+passworddb+conn+"/balance_schema")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	if IsUserExists(reserve.UserId, db) {
		var user_balance float64
		balance_res := db.QueryRow("SELECT user_balance FROM balance where user_id=?", reserve.UserId)
		err = balance_res.Scan(&user_balance)
		if user_balance >= reserve.Price {
			update, err := db.Query(
				"UPDATE balance SET user_balance = user_balance -? WHERE user_id=?",
				reserve.Price, reserve.UserId)
			defer update.Close()
			insert, err := db.Query(
				"INSERT INTO reserve(user_id,service_id,purchase_id,price) VALUES (?,?,?,?)",
				reserve.UserId, reserve.ServiceId, reserve.PurchaseId, reserve.Price)
			defer insert.Close()
			if err != nil {
				panic(err.Error())
			}
		}

	}
}

func ConfirmTransaction(confirm Reserve) {
	db, err := sql.Open("mysql", userdb+":"+passworddb+conn+"/balance_schema")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	del, err := db.Query(
		"DELETE FROM reserve WHERE (purchase_id = ?)",
		confirm.PurchaseId)
	rep_ins, err := db.Query(
		"INSERT INTO report(user_id,service_id,purchase_id,price) VALUES (?,?,?,?)",
		confirm.UserId, confirm.ServiceId, confirm.PurchaseId, confirm.Price)
	if err != nil {
		panic(err.Error())
	}
	defer del.Close()
	defer rep_ins.Close()

	if err != nil {
		panic(err.Error())
	}

}
