package model

type Account struct{
	account string
	balance float64
	possward string
}

func (a *Account)SetAccount(account1 string){
	if(len(account1)>10 ||len(account1)<6){
		fmt.Println("账号长度必须在6-10之间")
		return
	}else{
		a.account:account1
	}
}

func (a Account)SetBalance(balance1 string){
	if(balance1<20){
		fmt.Println("余额必须大于20")
	}else{
		a.balance：balance1
	}
}

func (a Account)SetAccount(account1 string){
	if(len(account1)>10 ||len(account1)<6){
		fmt.Println("账号长度必须在6-10之间")

	}else{
		a.account=account1
	}
}
