package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
Complete the method/function so that it converts dash/underscore
delimited words into camel casing. The first word within the output
should be capitalized only if the original word was capitalized
(known as Upper Camel Case, also often referred to as Pascal case).

Examples
"the-stealth-warrior" gets converted to "theStealthWarrior"
"The_Stealth_Warrior" gets converted to "TheStealthWarrior"
 */
//func ToCamelCase(s string) []string {
//	result := ""
//	var ab []string
//	//var newStringSlice []string
//	if strings.Contains(s,"-") {
//		fmt.Println("a",s)
//		ab := strings.Split(s,"-")
//		fmt.Println("ab", ab)
//		fmt.Println("ab len",len(ab))
//		fmt.Println("b",s)
//		for i := 0; i < len(s); i++ {
//			result=	strings.ToUpper(s)
//			var res2 []string
//			ab =append(res2, result)
//		}
//		fmt.Println("all",s)
//	}
//
//	return ab
//}
func ToCamelCase(s string) string {
	result:=""
	if strings.Contains(s,"-") {
		s=strings.Title(s)
		s1 :=strings.Split(s, "-")
		result=strings.Join(s1, "")
		return result
	} else {
		s=strings.Title(s)
		s1 :=strings.Split(s, "_")
		result=strings.Join(s1, "")
		return result
	}
}
func toWeirdCaseyy(str string) string {
	//weirde word gang
	 newStr := ""
	 count := 0
	 for _, j := range str {
		if string(j) == " " {
			count = 0
			newStr += " "
		} else {
			if count%2 == 0 {
				newStr += strings.ToUpper(string(j))
			} else {
				newStr += strings.ToLower(string(j))
			}
			count++
		}
	}
	return newStr
}
func toWeirdCase(str string) string {
	//weirde word gang
	newStr := ""
	count := 0
	for i:=0; i<len(str); i++ {
		if string(str[i]) == " " {
			count = 0
			newStr += " "
		} else {
			if count%2 == 0 {
				newStr += strings.ToUpper(string(str[i]))
			} else {
				newStr += strings.ToLower(string(str[i]))
			}
			count++
		}
	}
	return newStr
}

func HighAndLow(in string) string {
	fmt.Println("unsplited",in)
	strArray := []int{}
	strings.Split(in," ")
fmt.Println("splited",in)
	for _,j := range in{
		ab, _ := strconv.Atoi(string(j))
		fmt.Println(ab)
		strArray = append(strArray, ab)
	}
	fmt.Println(strArray)
	sort.Ints(strArray)
fmt.Println(strArray)
	 final := []string{strconv.FormatInt(int64(strArray[len(strArray)-1]),32),
	 	strconv.FormatInt(int64(strArray[0]),32)}
	 return strings.Join(final," ")

}



func toWeirdCase66(str string) string {
	s := ""
	i := 0
	for _, char := range(str) {
		if string(char)==" " {
			i=0
			s += " "
		} else {
			if i%2==0 {
				s += strings.ToUpper(string(char))
			} else {
				s += strings.ToLower(string(char))
			}
			i++
		}
	}
	return s
}

//for(var i = 1 ; i < newArr.length ; i++){
//newArr[i] = newArr[i].charAt(0).toUpperCase();
//}
//
//func abc (s string)string{
//	 var newStr  string
//	 var newArr [] string
//	if strings.Index(s,"-") != -1 {
//		strings.Split(s,"-")
//		for i:=1; i< len(s); i++ {
//			newStr=strings.ToUpper(string(s[0]))
//		}
//		return strings.Join(newArr,"")
//	} else if strings.Index(s,"_") != -1 {
//		strings.Split(s,"_")
//		for i:=1; i< len(s); i++ {
//			newStr=strings.ToUpper(string(s[0]))
//		}
//		return strings.Join(newArr,"")
//	}
//
//	return newStr
//}


func ToCamelCase1 (s string)string{

	var sliceString []string
	if strings.Contains(s, "-") {
		sliceString = strings.Split(s, "-")
	} else {
		sliceString = strings.Split(s, "_")
	}
fmt.Println(".....", sliceString[0])
	for i,_ := range sliceString {
		if i != 0{
			fmt.Println(sliceString[i])
			sliceString[i] = strings.Title(sliceString[i])
			//fmt.Println(sliceString[i])
		}
	}
	result := strings.Join(sliceString, "")
	return result
}
func lengthOfLastWord(s string) int {
	//space := regexp.MustCompile(`\s+`)
	//s = space.ReplaceAllString(s, "")
	slice := []string{}
	for i := 0; i < len(s); i++ {
		strings.Join(slice,"")
	}
	fmt.Println(slice)
	a := slice[len(slice)-1]
	return len(a)
}

type newstruct struct {

}

type Int interface {
	Abcd()
	Abc()
}

func (a newstruct) Abcd()  {

}
func (a newstruct) Abc()  {

}
func (a newstruct) Abce()  {

}

func main() {
	//s := "the-stealth-warrior"
	////abcde := []int{1,2,3,4,5}
	//s=strings.Title(s)
	//s1 :=strings.Split(s, "_")
	//s=strings.Join(s1, "")
	//fmt.Println(ToCamelCase1(s))
	////for(var i = 1 ; i < newArr.length ; i++){
	////	newArr[i] = newArr[i].charAt(0).toUpperCase();
	////}
	//fmt.Println(strings.Title("tolu thomas"))
	//var a bool
	//fmt.Println(a)
	str := "1 2 -3 4 5"
	fmt.Println(HighAndLow(str))
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
	//eat := "eattttt"
	//fmt.Println(AddTodo(eat))
	//fmt.Println(AddTodo(eat))
	//fmt.Println(List)
	//fmt.Println("abc",toWeirdCase(eat))

}
var List []string
func AddTodo (item string) string {
	List = append(List, item)
	return item
}