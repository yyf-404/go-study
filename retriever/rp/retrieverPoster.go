package rp

const url = "http://www.baidu.com"
/* 定义post接口 */
type Poster interface {
	Post(url string,
		form map[string]string) string
}
/* 定义Retriever接口 */
type Retriever interface {
	Get(url string) string
}
/* 组合Retriever和Post接口 */
type RetrieverPoster interface {
	Retriever
	Poster
}
/* 通过Session方法调用两个接口的方法 */
func Session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked  website",
	})
	return s.Get(url)
}
/* RetrieverPoster接口的实现类 */
type RPCometrue struct {
	Contents string
}
/* 要用指针函数才能修改结构体属性值 */
func (r *RPCometrue) Post(url string,
	form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
func (r *RPCometrue) Get(url string) string {
	return r.Contents
}

