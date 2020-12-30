// fetch 输出从url获取的内容
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const httpPrefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		// 第13页 练习1.8 如果输入的url不包含http前缀，则添加
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }
		// // 必须这样打印，负责打印的是数字
		// fmt.Printf("%s", b)

		// 第13页 练习1.7 函数io.Copy(dst, src)从src读取，并写入dst
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			os.Exit(1)
		}
		// 第13页 练习1.9 输出HTTP状态码
		fmt.Println()
		fmt.Println("HTTP State:", resp.StatusCode)
	}
}
