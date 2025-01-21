package main

import (
	"notify-withdrawal/repository"

	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "your-db-host"
	port     = 5432
	user     = "your-db-user"
	password = "your-db-password"
	dbname   = "your-db-name"
)

func receiveEmail() string {
	// メール受信処理（省略）
	return "payment email"
}

func parsePayment(email string) int {
	// メールから決済額を解析する処理（省略）
	return 5000
}

func sendEmail(message string) {
	// メール送信処理（省略）
	fmt.Println("メール送信:", message)
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("データベース接続に失敗しました:", err)
		return
	}
	defer db.Close()

	balance := repository.ChangeBalance(db)

	for {
		email := receiveEmail()
		payment := parsePayment(email)

		if payment > 0 {
			newBalance := balance - payment
			if newBalance < 0 {
				sendEmail("残高が不足しているため更新できません。")
				// 残高は変更しない
			} else {
				if newBalance < 10000 {
					sendEmail("残高が1万円未満です。")
				}
				balance = newBalance
			}
		}
	}
}
