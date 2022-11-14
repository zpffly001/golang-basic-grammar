package day2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	payload := [4]int{1}
	fmt.Printf("%p\n", &payload)
	change(payload)
}

func change(payload [4]int) {
	fmt.Printf("%p\n", &payload)
	fmt.Printf("%d\n", payload[0])
	payload[0] = 10
	fmt.Printf("%d\n", payload[0])
}

func TestMaxint(t *testing.T) {
	var a []int
	a = append(a, 1)
	a = append(a, 22)
	// 打印数组中的所有元素
	fmt.Println(a)
	// 数组长度
	fmt.Println(len(a))
}

func TestArrayType(t *testing.T) {
	var (
		a1 [4]int
		a2 [5]int
	)
	// [4]int
	fmt.Println(reflect.TypeOf(a1))
	// [5]int
	fmt.Println(reflect.TypeOf(a2))
}

func TestArrayD(t *testing.T) {
	// [0 0 0 0 0]
	var a1 [5]int
	fmt.Println(a1)

	// [   ]
	var a2 [4]string
	fmt.Println(a2)
}

func TestArrayD1(t *testing.T) {
	a1 := [3]int{1, 2, 3}
	fmt.Println(a1)
	// 如果将元素个数指定为特殊符号...，则表示通过初始化时的给定的值个数来推断数组长度
	a2 := [...]int{1, 2, 3, 4}
	fmt.Println(a2)
	a3 := [...]int{1, 1, 1}
	fmt.Println(a3)
	// 如果声明数组时，只想给其中某几个元素初始化赋值，则使用索引号
	// 3: 5表明在数组3索引上填充数字5
	a4 := [4]int{0: 1, 3: 5}
	fmt.Println(a4)
}

func TestArrayV1(t *testing.T) {
	a := [4]int{0: 1, 3: 5}
	fmt.Println(a[0])
	fmt.Println(a[3])

	a[0] = 10
	a[3] = 20
	fmt.Println(a[0])
	fmt.Println(a[3])
}

func TestArrayPtr(t *testing.T) {
	// 指针数组，数组中的元素都是指针
	a := [4]*int{0: new(int), 3: new(int)}
	fmt.Println(a)

	// 如果指针地址为空, 是会报空指针错误的, 比如
	// *a[1] = 3

	*a[0] = 10
	*a[3] = 20
	// 打印的都是地址，因为数组中存放的就是地址
	fmt.Println(a)
	// 打印的是指针解引用对应的值
	fmt.Println(*a[0], *a[3])

	// 为1赋值
	a[1] = new(int)
	*a[1] = 30
	fmt.Println(a, *a[1])
}

func TestArrayCopy(t *testing.T) {
	a1 := [4]string{"a", "b", "c", "m"}
	a2 := [4]string{"x", "y", "z", "n"}
	// 浅拷贝？拷贝的是值？
	a1 = a2
	fmt.Println(a1, a2)
	// 数组中元素的值相同
	fmt.Println(a1[0], a2[0])
	// 数组中元素的值的地址不一致；浅拷贝？拷贝的是值？
	fmt.Println(&a1[0], &a2[0])
}

func TestArrayPtrCopy(t *testing.T) {
	a1 := [4]*string{new(string), new(string), new(string), new(string)}
	a2 := a1
	// 经过复制后，
	fmt.Println(a1, a2)
	// 数组中元素的值相同
	fmt.Println(a1[0], a2[0])
	// 数组中元素的值的地址不一致；浅拷贝？拷贝的是值？
	fmt.Println(&a1[0], &a2[0])

	*a1[0] = "A"
	*a1[1] = "B"
	*a1[2] = "C"
	*a1[3] = "C"
	// 因为a1[0]和a2[0]是同一个指针，因此修改一个另外一个也会跟着变化
	fmt.Println(*a1[0], *a2[0])
}

func TestArrayIter(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	for i, v := range a {
		fmt.Println(i, v)
	}
}

func TestArray2x(t *testing.T) {
	// 二维数组
	pos := [4][2]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
	fmt.Println(pos)

	pos[0] = [2]int{10, 10}
	fmt.Println(pos)
}

func TestArray(t *testing.T) {
	a := [2]int{}
	fmt.Println(a, a[1])
}
