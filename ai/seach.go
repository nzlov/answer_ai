package ai

import(
	"strings"
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

/*
  将关键字放在百度搜索结果
  返回搜索到的相关记录数
*/
func Seach(key string) (count int64){
	count = 0
	url := fmt.Sprintf("http://www.baidu.com/s?tn=ichuner&lm=-1&word=%s&rn=1", url.QueryEscape(key))
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println(err.Error())
	}else{
		rd := bufio.NewReader(resp.Body)
		for {
			linebuf,_, err := rd.ReadLine()
			if err != nil || io.EOF == err {
				break
			}
			line := string(linebuf)
			if strings.Contains(line, "百度为您找到相关结果约") {
				staIndex := strings.Index(line, "百度为您找到相关结果约") +33
				line = line[staIndex:]
				endIndex := strings.Index(line, "个")
				line = line[0:endIndex]
				line = strings.Replace(line, ",", "", -1)
				count, err = strconv.ParseInt(line, 10, 64)
				if err != nil{
					count = 0
				}
				break
			}
		}
	}
	return count
}