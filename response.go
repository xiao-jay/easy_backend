package main

type SuccessReturn struct {
    Status bool `json:"status"`
    Msg   string      `json:"msg"`
    TokenData  interface{}   `json:"data"`
}
type ErrorReturn struct {
    Status bool    `json:"status"`
    Msg   string      `json:"msg"`
    TokenData  interface{}   `json:"data"`
}

func MakeSuccessReturn(data interface{})(int ,interface{})  {
    return 200,SuccessReturn{
        Status: true,
        Msg: "success",
        TokenData: data,
    }
}
func MakeErrorReturn(msg string)(int ,interface{})  {
    return 400,ErrorReturn{
        Status: false,
        Msg: msg,
        TokenData: "",
    }
}
