package main

import (
    "design.patterns/CreationalPatterns/AbstractFactory/abstractfactory"
)

func main() {

    var factory abstractfactory.DAOFactory
    //factory = &AbstractFactory.MySQLFactory{}
    factory = &abstractfactory.OracleFactory{}
    factory.CreateOrderMainDAO().SaveOrderMain()
    factory.CreateOrderDetailDAO().SaveOrderDetail()
}
