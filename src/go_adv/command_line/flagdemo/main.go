package main
import (
	"fmt"
	"flag"
)

type Parameter struct {
	user string
	pwd string
	port int 
	host string 
}
func main(){
	par :=Parameter{}
	flag.StringVar(&par.user,"u","","用户名，默认为空")
	flag.StringVar(&par.host,"h","localhost","主机名,默认为localhost")
	flag.IntVar(&par.port,"port",3306,"端口名,默认为3306")
	flag.StringVar(&par.pwd,"pwd","","密码")

	flag.Parse()
	fmt.Printf("user:%s,pwd:%s,port:%d,host:%s",par.user,par.pwd,par.port,par.host)
}