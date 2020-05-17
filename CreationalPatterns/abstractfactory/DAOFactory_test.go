package abstractfactory

import (
    "testing"
)

func TestDAOFactory(t *testing.T) {

    var factory DAOFactory
    factory = &MySQLFactory{}
    factory.CreateOrderMainDAO().SaveOrderMain()
    factory.CreateOrderDetailDAO().SaveOrderDetail()

    factory = &OracleFactory{}
    factory.CreateOrderMainDAO().SaveOrderMain()
    factory.CreateOrderDetailDAO().SaveOrderDetail()
}
