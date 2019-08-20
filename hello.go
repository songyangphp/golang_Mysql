package main

import (
	"strconv"
);

func main() {
	age := (1+2)*(3+9);
	b := "hello";

	//int类型转换为string
	c := strconv.Itoa(age);

	/*if(false){
		println(111);
	}else{
		println(222);
	}*/

	switch age {
		case 1 : println(1);break;
		case 2 : println(2);break;
		case 15 : println(15);break;
		default: println("def");
	}

	max := getMax(1,2);

	println(max);
	println(b+" "+c);
}

func getMax(num1 int,num2 int) int {
	var max int;
	if(num1 > num2){
		max = num1;
	}else{
		max = num2;
	}
	return max;
}