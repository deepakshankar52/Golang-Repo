package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Account struct {
	ACC_NO         int    `json:"acc_no"`
	Account_Holder string `json:"account_holder"`
	Amount         int    `json:"amount"`
}

func accountInfo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAccount(w, r)
	case "POST":
		createAccount(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
}

func accountDeposit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		depositUpdate(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func accountWithdrawal(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		withdrawalUpdate(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "User,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,AAthorization")

	acc_no_str := r.URL.Query().Get("acc_no")
	
	lUser := r.Header.Get("User")

	if acc_no_str != "" {
		acc_no, err := strconv.Atoi(acc_no_str)
		if err != nil {
			http.Error(w, "Invalid account number parameter", http.StatusBadRequest)
			return
		}

		for _, account := range accounts {
			if account.ACC_NO == acc_no {
				json.NewEncoder(w).Encode(account)
				return
			}
		}

		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}
	fmt.Println(lUser)

	// cookie := &http.Cookie {
	// 	Name: "id",
	// 	Value: "Cookie-message",
	// 	MaxAge: 300,
	// }
	// http.SetCookie(w, cookie)
	// w.WriteHeader(200)
	// w.Write([]byte("Account Fetched Successfully"))

	cookie := http.Cookie{
		Name:    "token",
		Value:   "0rPKW5dUd914eS4qBYzkPBXlVHCzgPl1JHxYZJrQZUfo9uwWKtMiDtFQQ7rWoEuP",
		Expires: time.Now().Add(24 * time.Hour), // Cookie expires in 24 hours
		HttpOnly: true,
		Path: "/",
		Secure: false,
		SameSite: http.SameSiteStrictMode,
	}

	// Set the cookie in the response
	http.SetCookie(w, &cookie)

	json.NewEncoder(w).Encode(accounts)

	// return



}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var new_account Account
	if err := json.NewDecoder(r.Body).Decode(&new_account); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	accounts = append(accounts, new_account)
	json.NewEncoder(w).Encode(new_account)
}

func depositUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var depositData struct {
		ACC_NO int `json:"acc_no"`
		Amount int `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&depositData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	for i, account := range accounts {
		if account.ACC_NO == depositData.ACC_NO {
			accounts[i].Amount += depositData.Amount
			json.NewEncoder(w).Encode(accounts[i])
			return
		}
	}

	http.Error(w, "Account not found", http.StatusNotFound)
}

func withdrawalUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var withdrawData struct {
		ACC_NO int `json:"acc_no"`
		Amount int `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&withdrawData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	for i, account := range accounts {
		if account.ACC_NO == withdrawData.ACC_NO {
			if accounts[i].Amount < withdrawData.Amount {
				http.Error(w, "Insufficient funds", http.StatusBadRequest)
				return
			}
			accounts[i].Amount -= withdrawData.Amount
			json.NewEncoder(w).Encode(accounts[i])
			return
		}
	}

	http.Error(w, "Account not found", http.StatusNotFound)
}

// func cookieHandler(w http.ResponseWriter, r *http.Request) {
// 	cookie := &http.Cookie {
// 		Name: "id",
// 		Value: "Cookie-message",
// 		MaxAge: 300,
// 	}
// 	http.SetCookie(w, cookie)
// 	w.WriteHeader(200)
// 	w.Write([]byte("Account Fetched Successfully"))
// 	return
// }


var accounts []Account

func main() {
	accounts = append(accounts, Account{ACC_NO: 1, Account_Holder: "ajith", Amount: 100})
	accounts = append(accounts, Account{ACC_NO: 2, Account_Holder: "vijay", Amount: 100})
	accounts = append(accounts, Account{ACC_NO: 3, Account_Holder: "kamal", Amount: 100})
	accounts = append(accounts, Account{ACC_NO: 4, Account_Holder: "rajni", Amount: 100})
	accounts = append(accounts, Account{ACC_NO: 5, Account_Holder: "surya", Amount: 100})

	http.HandleFunc("/accounts", accountInfo)
	http.HandleFunc("/accounts/deposit", accountDeposit)
	http.HandleFunc("/accounts/withdraw", accountWithdrawal)
	// http.HandleFunc("/accounts/cookie", cookieHandler)

	fmt.Println("Starting on port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
