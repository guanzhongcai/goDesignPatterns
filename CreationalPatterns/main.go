package main

import "design.patterns/CreationalPatterns/AbstractFactory"

func main() {

    var factory AbstractFactory.DAOFactory
    //factory = &AbstractFactory.MySQLFactory{}
    factory = &AbstractFactory.OracleFactory{}
    factory.CreateOrderMainDAO().SaveOrderMain()
    factory.CreateOrderDetailDAO().SaveOrderDetail()
}
