package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	Id int64
	Username string
	Password string
}

type Category_Articles struct {
	Id int64
	Title string
	Created time.Time `orm:"index"`
	Views int64 `orm:"index"`
	Topictime time.Time `orm:"index"`
	Topiccount int64
	TopicLastUserId int64
	Articles *Articles `orm:"reverse(one)"`
}

type Articles struct {
	Id int64
	Uid int64
	Title string
	Content string `orm:"size(5000)"`
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	LastView int64 `orm:"index"`
	Views int64 `orm:"index"`
	Author string
	CategoryId int64
	Category_Articles *Category_Articles `orm:"rel(one)"`
}

type Category_Photos struct {
	Id int64
	Title string
	Created time.Time `orm:"index"`
	Views int64 `orm:"index"`
	Topictime time.Time `orm:"index"`
	Topiccount int64
	TopicLastUserId int64
	Photos *Photos `orm:"reverse(one)"`
}

type Photos struct {
	Id int64
	Uid int64
	Title string
	Created time.Time `orm:"index"`
	Views int64 `orm:"index"`
	LastView time.Time `orm:"index"`
	Url string
	Author string
	Category_Photos *Category_Photos `orm:"rel(one)"`
}

func RegisterDB() {
	orm.RegisterModel(new(User),new(Category_Articles),new(Articles),new(Photos),new(Category_Photos))
	orm.RegisterDataBase("default", "mysql", "leig:leig2013@/leig?charset=utf8", 10)
}

func SelectArticleAll() {
	
}

func Estimate(name string) string{
	o := orm.NewOrm()
	user := User{Username: name}
	err := o.Read(&user,"username")
	if err == orm.ErrNoRows {
	    fmt.Println("查询不到")
	    return "nil"
	} else if err == orm.ErrMissPK {
	    fmt.Println("找不到主键")
	    return "nil"
	} else {
	    fmt.Println(user.Id, user.Username)
	    return user.Password
	}
}
