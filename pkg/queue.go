package life
type Queue []interface{}
func (queue *Queue) Push(value interface{}){
	*queue= append(*queue,value)
}
func (queue *Queue) Pop() interface{} {
	value:=(*queue)[0]
	*queue=(*queue)[1:]
	return value
}
func (queue Queue) IsEmput() bool {
	if len(queue)==0{
		return true
	}
	return false
}