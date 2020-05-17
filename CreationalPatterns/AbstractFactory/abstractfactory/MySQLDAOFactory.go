package abstractfactory

import "fmt"

//mysql针对两个接口的实现

type MySQLMainDAO struct {
}

func (*MySQLMainDAO) SaveOrderMain() {
    fmt.Println("MySQL main save")
}

type MySQLDetailDAO struct {
}

func (*MySQLDetailDAO) SaveOrderDetail() {
    fmt.Println("MySQL detail save")
}