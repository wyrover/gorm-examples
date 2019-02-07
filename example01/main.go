package main

import (
    "fmt"

    "github.com/BurntSushi/toml"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)


type DbConfig struct {
    Driver    string
    Server    string
    User      string
    Password  string
    Database  string
    Charset   string
    ParseTime string
}

type Config struct {
    Database DbConfig
}

func (d DbConfig) DSN() string {
    return fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=%s", d.User, d.Password, d.Server, d.Database, d.Charset, d.ParseTime)
}

func (c Config) Db() (string, string) {
    return c.Database.Driver, c.Database.DSN()
}

func GetConfig() Config {
    var config Config
    _, err := toml.DecodeFile("config.toml", &config)
    if err != nil {
        panic("unloaded config file")
    }

    return config
}

var db *gorm.DB

func init() {
    var err error
    config := GetConfig()

    db, err = gorm.Open(config.Db())
    if err != nil {
        panic(err.Error())
    }
}



//type Account struct {
//    ID int
//    Name string
//    Email string
//    Address string
//}


type Account struct {
    gorm.Model
    Name string
    Email string
    Address string
}

func main() {
    fmt.Println("hello world!")

    // Account -> accounts 复数形式的表
    db.AutoMigrate(&Account{})


    // 创建
    account1 := Account{Name: "Jinzhu", Email: "test@test.com", Address: "tokorozawa"}
    db.Create(&account1)

    // 查询
    account := Account{}
    db.First(&account)

    fmt.Println(account.ID)
    fmt.Println(account.CreatedAt)
    fmt.Println(account.UpdatedAt)
    fmt.Println(account.DeletedAt)

    fmt.Println(account.Name)
    fmt.Println(account.Email)
    fmt.Println(account.Address)

    account2 := Account{}
    db.Debug().Where("id = ?", account.ID).Find(&account2)


    // 更新
    account2.Address = "Beijing"
    db.Debug().Save(&account2)

    account3 := Account{}
    db.First(&account3)
    fmt.Println(account3.Address)

    // 删除
    db.Delete(&account3)

    account4 := Account{}
    db.First(&account4)
    fmt.Println(account4.ID)
    
}

