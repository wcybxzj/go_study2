package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //只执行里面的init,不使用里面的API
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"log"
)

//这里xorm这个tag是让xorm来识别的
//所有字段必须首字母大写才能让xorm反射出去
type Account struct {
	Id int64 //Id int64默认会做成主键(Id生成到mysql中字段名是id,ID生成到mysql中的字段名是i_d)
	Name string `xorm:"unique"` //唯一key
	Balance float64
	Version int `xorm:"version"` //version 说明这个是乐观锁
}

var x *xorm.Engine

func InitDatabase()  {
	var err error
	x, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatalf("Fail to create engine:%v",err)
	}

	//mysql要自己创建database
	//sqlite连database都不用创建
	//同步Account表结构到database
	err = x.Sync2(Account{})
	if err != nil {
		log.Fatalf("Fail to sync database:%v",err)
	}
}

func NewAccount(name string, balance float64) error {
	a := new(Account)
	a.Name = name
	a.Balance = balance
	_, err := x.Insert(a)
	return err
}

func GetAccount(id int64)(*Account, error){
	a := &Account{}
	has, err := x.Id(id).Get(a)
	if err != nil{
		return nil, err
	}else if !has{
		return nil, errors.New("Account not found")
	}
	return a, nil
}

func MakeDeposit(id int64, deposit float64)(*Account ,error)  {
	a, err :=GetAccount(id)
	if err!=nil{
		return nil, err
	}
	a.Balance += deposit
	_, err= x.Update(a)
	return a, err
}

func MakeWithdraw(id int64, withdraw float64)(*Account, error)  {
	a, err :=GetAccount(id)
	if err!=nil{
		return nil, err
	}

	if a.Balance <= withdraw{
		return nil, errors.New("Not enouugh balance")
	}

	a.Balance -= withdraw
	_, err= x.Update(a)
	return a, err
}

func MakeTransfer(id1, id2 int64, balance float64) error {
	a1, err :=GetAccount(id1)
	if err != nil {
		return err
	}
	
	a2, err :=GetAccount(id2)
	if err != nil {
		return err
	}

	if a1.Balance <= balance {
		return errors.New("NOt enough balance")
	}
	a1.Balance -=balance
	a2.Balance +=balance

	if _,err = x.Update(a1); err!=nil {
		return err
	}else if _,err = x.Update(a2); err!=nil{
		return err
	}
	return nil
}

func GetAccountsAscId()(as []*Account, err error)  {
	err = x.Asc("id").Find(&as)//这里id必须是mysql字段列名
	return as, err
}

func DeleteAccount(id int64) error {
	_,err := x.Delete(&Account{Id:id})
	return err
}

func GetAccountCount() (int64, error) {
	return x.Count(new(Account))
}

func CreateTestData() (error) {
	sql := "truncate account"
	_, err := x.Query(sql)
	if err != nil{
		return err
	}

	for i:=0; i<10 ;i++  {
		err := NewAccount(fmt.Sprintf("ybx%d", i), float64(i)*100)
		if err!=nil{
			log.Fatalf("Fail to create account")
		}
	}
	return err
}

var printFn = func(idx int, bean interface{}) error{
	fmt.Printf("%d:%#v\n", idx, bean.(*Account))
	return nil
}

func Iterate1()  {
	x.Iterate(new(Account), printFn)
}

func Iterate2()  {
	a := new(Account)
	rows, err := x.Rows(new (Account))
	if err != nil{
		log.Fatalf("Fail to get rows")
	}
	defer rows.Close()

	for rows.Next()  {
		if err = rows.Scan(a); err != nil{
			log.Fatalf("Fail to get row:%v\n",err)
		}
		fmt.Printf("%#v\n",a)
	}
}

func TestIn(ids[]int64) ([]*Account){
	var err error
	accounts := make([]*Account, 0)
	err = x.Cols("id", "name").In("id", ids).Limit(1,0).Find(&accounts)
	if err != nil {
		panic("err:"+err.Error())
	}
	return accounts
}