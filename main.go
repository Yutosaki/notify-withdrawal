package main

import (
	"fmt"
)

func getInitialBalance() int {
	var balance int
	fmt.Print("最初の貯金残高を入力してください: ")
	fmt.Scan(&balance)
	return balance
}

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
	balance := getInitialBalance()

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
