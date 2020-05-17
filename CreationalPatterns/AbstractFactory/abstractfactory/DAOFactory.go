package abstractfactory

//mysql
//sql server
//oracle

//sql sentence

//订单
//订单报表

//把操作抽象出来
type OrderMainDAO interface { //订单记录
    SaveOrderMain() //保存
    //DeleteOrderMain()
    //SearchOrderMain()
}

type OrderDetailDAO interface { //订单报表详情
    SaveOrderDetail() //保存
}

type DAOFactory interface { //抽象工厂接口
    CreateOrderMainDAO() OrderMainDAO
    CreateOrderDetailDAO() OrderDetailDAO
}

//详情请见多文件
