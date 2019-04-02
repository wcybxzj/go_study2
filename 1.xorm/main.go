package main

import (
	"fmt"
	"go_study2/1.xorm/models"
	)

const prompt =`plz enter a number
1.create new account
2.show detail of account
3.Deposit(存款)
4.Withdraw(取款)
5.转账
6.List account by Id
7.List account by balance
8.Delete account
9.清空表-生成10条新测试数据
10.获取总数
11.迭代方式1
12.迭代方式2
99.exit
`

func main() {
	models.InitDatabase()
	if err := models.NewAccount("abc", 123); err != nil{
		fmt.Println(err)
	}

Exit:
	for{
		fmt.Println(prompt)
		var num int
		fmt.Scanf("%d\n", &num)

		switch num{
		case 1:
			fmt.Println("<name> <balance>")
			var name string
			var balance float64
			fmt.Scanf("%s %f\n", &name, &balance)
			if err := models.NewAccount(name, balance); err != nil{
				fmt.Println(err)
			}

		case 2:
			fmt.Println("<id>")
			var id int64
			fmt.Scanf("%d\n", &id)
			if account, err := models.GetAccount(id); err != nil{
				fmt.Println(err)
			}else{
				fmt.Printf("%#v\n",account)
			}

		case 3:
			fmt.Println("<id> <存款数>")
			var id int64
			var deposit float64
			fmt.Scanf("%d %f\n", &id, &deposit)
			account, err := models.MakeDeposit(id, deposit)
			if err != nil{
				fmt.Println(err)
			}else{
				fmt.Printf("%#v\n",account)
			}
		case 4:
			fmt.Println("<id> <取款数>")
			var id int64
			var deposit float64
			fmt.Scanf("%d %f\n", &id, &deposit)
			account, err := models.MakeDeposit(id, deposit)
			if err != nil{
				fmt.Println(err)
			}else{
				fmt.Printf("%#v\n",account)
			}

		case 5:
			fmt.Println("<转出者id> <款数> <转入者id>")
			var id1 ,id2 int64
			var balance float64
			fmt.Scanf("%d %d %f\n", &id1, &id2, &balance)
			if err := models.MakeTransfer(id1, id2, balance); err != nil{
				fmt.Println(err)
			}

		case 6:
			as, err := models.GetAccountsAscId()
			if err != nil {
				fmt.Println(err)
			}else{
				for i,a := range as {
					fmt.Println("%d: %#v\n", i, a)
				}
			}
		case 7:

		case 8:
			fmt.Println("<id>")
			var id int64
			fmt.Scanf("%d \n", &id )
			if err := models.DeleteAccount(id); err != nil{
				fmt.Println(err)
			}

		//9.创建10条测试数据
		case 9:
			models.CreateTestData()

		//10.获取总数
		case 10:
			fmt.Println(models.GetAccountCount())

		//11.迭代方式1
		case 11:
			models.Iterate1()

		//12.迭代方式2
		case 12:
			models.Iterate2()




		case 99:
			break Exit
		}
	}
}
