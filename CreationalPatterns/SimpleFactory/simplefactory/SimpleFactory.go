package simplefactory

//中国人，你好
//英国人，hello

type API interface {
    Say(name string) string
}

func NewAPI(str string) API {
    if str == "en" {
        return &English{}
    } else if str == "cn" {
        return &Chinese{}
    } else {
        return nil
    }
}

type Chinese struct {
}

func (*Chinese) Say(name string) string {

    return "你好 " + name
}

type English struct {
}

func (*English) Say(name string) string {
    return "hello " + name
}
