package main  
  
import (  
    "github.com/go-xorm/xorm"  
    _ "github.com/mattn/go-sqlite3"  
)  
  
type User struct {  
    Id   int64  
    Name string `xorm:"varchar(25) notnull unique 'usr_name'"`  
}  
  
var engine *xorm.Engine  
  
func main() {  
    engine, _ = xorm.NewEngine("sqlite3", "./test.db")  
    engine.Sync2(new(User))  
}