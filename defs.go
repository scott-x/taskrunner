/*
* @Author: apple
* @Date:   2019-07-10 05:49:02
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-24 17:50:12
 */
package taskrunner

//预定义变量
const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSE             = "c"
)

type controlChan chan string

//interface{} means any type here
type DataChan chan interface{}

type fn func(dc DataChan) error
