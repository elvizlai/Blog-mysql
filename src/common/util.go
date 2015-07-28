/**
@author ElvizLai
*/

package common

import (
	"regexp"
	"strings"
	"errors"
)

//todo buglist 1、只传照片img时，摘要错误的bug   2、复制黏贴内容时，摘要错误的bug

//整个程序分为3大部分
//1、正则过滤
//2、筛选固定长度
//3、标签检查及补全

func SubHtml(content string, limit int) string {
	if len(content)<limit {
		return content
	}

	//-->正则过滤获取content中的所有标签
	all := regexp.MustCompile(`<[\S\s]+?>`)
	allResult := all.FindAllString(content, -1)

	needRemove := regexp.MustCompile(`<[\S\s]+?/>`)//筛选项目，匹配诸如<br/>之类的标签

	//筛选固定长度-->limit
	var i, totalLength int
	cpOfContent := content
	//计算所需要截取的长度
	for strLength := 0; i<len(allResult); i++ {
		index := strings.Index(content, allResult[i])//标签索引
		lenOfIndex := len(allResult[i])//标签长度

		subStr := []byte(content)[:index]//去开始到标签索引位置所对应的字符串

		content=string([]byte(content)[index+lenOfIndex:])//去除所取的字符串

		strLength+=len([]rune(string(subStr)))//将新取出的字符串长度加到已截取的字符串长度中

		totalLength=len([]rune(string(subStr)))+lenOfIndex+totalLength//需要截取的总长度

		if strLength>limit {
			break
		}
	}

	//截取后的字符串
	if totalLength==0 {
		cpOfContent=string([]rune(cpOfContent)[:limit])
	}else {
		cpOfContent=string([]rune(cpOfContent)[:totalLength])
	}

	//-->修复标签，如果i==len(allResult)，则无需修复标签
	if i!=len(allResult) {
		allResult=allResult[:i+1]

		//移除不需要补全的标签
		for i := 0; i<len(allResult); i++ {
			if len(needRemove.FindStringIndex(allResult[i]))==2 {
				allResult=append(allResult[:i], allResult[i+1:]...)
				i--
			}
		}

		//使用栈来维护标签是否配对
		stack := NewStack(len(allResult))

		for i := 0; i<len(allResult); i++ {
			if strings.HasPrefix(allResult[i], "</") {
				stack.Pop()
			}else {
				stack.Push(allResult[i])
			}
		}

		for {
			if stack.Len()!=0 {
				str, _ := stack.Pop()
				index := strings.Index(str, " ")
				cpOfContent += "</"+string([]byte(str)[1:index])+">"
			}else {
				break
			}
		}
	}

	return cpOfContent
}


type Stack struct {
	st []string
	len int
	cap int
}

func NewStack(cap int) *Stack {
	st := make([]string, 0, cap)
	return &Stack{st, 0, cap}
}

func (this *Stack)Push(p string) {
	this.st = append(this.st, p)
	this.len = len(this.st)
	this.cap = cap(this.st)
}

func (this *Stack)Pop() (string, error) {
	if this.len == 0 {
		return "", errors.New("Can't pop an empty stack")
	}
	this.len -= 1
	out := this.st[this.len]
	this.st = this.st[:this.len]
	return out, nil
}

func (this *Stack)Len() int {
	return this.len
}

func (this *Stack)Cap() int {
	return this.cap
}

