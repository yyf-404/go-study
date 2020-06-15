package main
import (
	"go-study/errHanding/fileListServer/filelisting"
	"log"
	"net/http"
	"os"
)
//给对应函数起别名
type apphandler func (writer http.ResponseWriter, request *http.Request) error
/*  通过对函数包装和函数式接口改变 函数返回值 集中处理错误*/
func errorWraper(handler apphandler) func (writer http.ResponseWriter, request *http.Request){
	return func (writer http.ResponseWriter, request *http.Request){
		defer func() {  //通过recover处理 panic
			if r := recover();r!=nil{
				log.Printf("Error occurred "+
					"handling request:%v)",r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err:=handler(writer,request) //交给传入的apphandler实现函数处理
		if userErr,ok := err.(userError);ok{  //如果是用户自定义异常 直接返回
			http.Error(writer,userErr.Message(),http.StatusBadRequest)
			return
		}
		if err!=nil{
			log.Printf("Error occurred "+
				"handling request: %s",
				err.Error())
			code:=http.StatusOK
			switch{
			case os.IsNotExist(err):
				code=http.StatusNotFound
			default:
				code=http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)// 返回给客户端的错误
		}
	}

}
//自定义错务类型
type userError interface {
	error
	Message() string
}
func main() {
	http.HandleFunc("/",errorWraper(filelisting.HandlerFileListing) )
	err := http.ListenAndServe(":8888", nil) //启动服务器
	if err != nil {
		panic(err)
	}
}
