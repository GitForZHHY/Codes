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
 * Func_name :DealWithLogInfo
 * Func      :
 * Author    :zhhy
 * Date      :20190317
 *
 */
func DealWithLogInfo(loginfo string,args []interface{}) (outinfo string) {

	var args_buf strings.Builder
	var loginfo_buf strings.Builder
	//var loginfo_tmp_str string
	var ArgsNum int

	Max_Len := len(loginfo)


	fmt.Printf("loginfolen[%d]\n",Max_Len)
	for i := 0; i < Max_Len; {
		lasti := i
		for i < Max_Len && loginfo[i] != '%' {
			i++
		}     

		if i > lasti {
			//loginfo_buf.Write([]byte(loginfo[lasti:i]))
			loginfo_buf.WriteString(loginfo[lasti:i])
			loginfo_tmp_str := args_buf.String()
			fmt.Printf("lasg[%d]i[%d]loginfonow- %s \n",lasti,i,loginfo_tmp_str)
			fmt.Printf("lasg[%d]i[%d]loginfonow- %s \n",lasti,i,loginfo[lasti:i])
		}     
		if i >= Max_Len {
			// done processing format string
			//args_buf.Write([]byte(loginfo[lasti:i]))
			args_buf.WriteString(loginfo[lasti:i])
			outinfo = args_buf.String()

			fmt.Printf("DEAL DONE!!!\n")
			break
		}     

		// Process one verb
		args_buf.Write([]byte(loginfo[lasti:i]))

		i++
		//for _,arg := range args {
		arg := args[ArgsNum]
		switch verb := arg.(type) {
		case string:
			args_buf.WriteString(verb)
			outinfo = args_buf.String()
		case int:
			int_str_tmp := strconv.Itoa(verb)
			args_buf.WriteString(int_str_tmp)
			outinfo = args_buf.String()

			case bool:      
			case []byte:    
			case float32:   
			case int8:      
			case float64:   
			case int16:     
			case int32:     
			case int64:     
			log.Fatal("u need wait...")
		default:
			log.Fatal("unknown type")
		}
		//	}
		ArgsNum++
		i++   
	}

	 return outinfo
}
/*
 * Func_name :swVdebug
 * Func      :
 * Author    :zhhy
 * Date      :20190317
 *
 */

func swVdebug(msg string,args ...interface{}){
 
	msg_str := DealWithLogInfo(msg,args)
	fmt.Printf("LOGMSG: %s\n",msg_str)
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
