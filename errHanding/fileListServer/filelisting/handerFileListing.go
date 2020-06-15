package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)
const Prefix string="/list/"
/* 实现userError接口*/
type userError string
func (userErr userError)Error() string{
	return userErr.Message()
}
func (userErr userError) Message() string{
	return string(userErr)
}
/*  请求处理函数 */
func HandlerFileListing(writer http.ResponseWriter, request *http.Request) error { //url 请求处理器
	    if strings.Index(request.URL.Path,Prefix)!=0{
	    	//返回自定义错务类型userError
	    return 	userError("path must be start with "+Prefix)
		}
		path := request.URL.Path[len(Prefix):]
		fmt.Println(path)
		file, err := os.Open(path) //打开文件
		if err != nil {
			return err
		}
		defer file.Close()
		all, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		writer.Write(all)
	return nil
	}

