package main

import "fmt"

type Bank struct {
	// in JAVA
	// private String bankName;
	// private String accountType;
	// private String accountName;

	// in GoLang
	bankName    string
	accountType string
	accountName string
}

func (b *Bank) String() string {
	return fmt.Sprintf("Bank:{\nbankName    :'%s',\naccountType :'%s',\naccountName :'%s'\n}",
		b.bankName, b.accountType, b.accountName)
}

func BankConstructor(bankName, accountType, accountName string) *Bank {

	newBankDetsils := &Bank{
		bankName:    bankName,
		accountType: accountType,
		accountName: accountName,
	}
	return newBankDetsils
}
func main() {
	passBook := BankConstructor("HDFC Bank", "Savings", "Arjun Singh")
	fmt.Println("passBook Details :", passBook)
	//output
	//passBook Details : &{HDFC Bank Savings Arjun Singh}

	fmt.Printf("\npassBook Details :%+v\n", passBook)
	//passBook Details :&{bankName:HDFC Bank accountType:Savings accountName:Arjun Singh}

	passBook1 := BankConstructor("HDFC Bank", "Savings", "kabbu jii")
	fmt.Println("\npassBook Details :", passBook1)
	//after toString overrideing in (java developer words)output will come like this
	// passBook Details : Bank:{
	// 	bankName    :'HDFC Bank',
	// 	accountType :'Savings',
	// 	accountName :'kabbu jii'
	// 	}
}
