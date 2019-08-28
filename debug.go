package main

import (
	"os"
	"log"
	"fmt"
	"time"
	"runtime"
	"strings"
	"strconv"
	"path/filepath"

)

/*
 * Func_name :swVdebug
 * Func      :
 * Author    :zhhy
 * Date      :20190317
 *
 */

func swVdebug(msg string,args ...interface{}){
 
	var  line int
	//org_file_name,line := "???",0
        
	_,org_file_name,_,ok := runtime.Caller(0)
	if ok {
		//funcname = runtime.FuncForPC(pc).Name()
		//funcname = filepath.Ext(funcname)
		//funcname = strings.TrimPrefix(funcname,".")
		org_file_name = filepath.Base(org_file_name)
	}
	file_name := fmt.Sprintf("%s/log/%s_%d.debug",os.Getenv("HOME"),os.Args[0],os.Getpid())

	_,_,line,ok = runtime.Caller(1)
	file_fp,err := os.OpenFile(file_name,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		log.Printf("open %s fail...\n",file_name)
	}
	msg_str := fmt.Sprintf(msg,args...)
	
	//time|pid|file_name|line|msg
	//file_line := fmt.Sprintf("%s|%s|%s|%d|%s\n",time.Now().Format("2006/01/02-15:04:05"),org_file_name,funcname,line,msg)
	file_line := fmt.Sprintf(" %s|%6d|%s|%5d|%s   \n",time.Now().Format("2006/01/02-15:04:05"),os.Getpid(),org_file_name,line,msg_str)
	fmt.Println(file_line)

	if _,err := file_fp.Write([]byte(file_line)); err != nil {
		log.Printf("write %s-%s fail[%s]...\n",msg,file_name,err)
	}

	if err := file_fp.Close(); err != nil {
		log.Printf("close-%s fail...\n",file_name)
	}
 }

/*    
 * Func_name :main
 * Func      :
 * Author    :zhhy
 * Date      :20190317
 *
 */

func main(){
	var i int
	i = 10
	swVdebug("this is a msglog:string[%s]int[%d]string[%s]","XXX",i,"1234")
	swVdebug("this is a msglog:string[%s]","1234")
	swVdebug("this is a msglog:string")
	
}
