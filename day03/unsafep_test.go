package day3

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafeMain(t *testing.T) {
	a := [3]int64{1, 2, 3}
	fmt.Printf("%p\n", &a)

	// 64位 8字节
	s1 := unsafe.Sizeof(a[0])
	fmt.Printf("%d\n", s1) // a + 8B  a[1]

	// 1. *[3]int64 --> unsafe.Pointer ---> uintptr(uint64)
	p1 := &a // *p1
	fmt.Printf("%d\n", *p1)
	// 这个就是相当于c语言中指针移动，移动sizeof个字节
	fmt.Println(*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(p1)) + unsafe.Sizeof(a[0]))))
	fmt.Println(*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(p1)) + unsafe.Sizeof(a[0])*2)))
	fmt.Println(*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(p1)) + unsafe.Sizeof(a[0]) + 8)))
	// fmt.Println() 标准输入
	// int int32 float32  *int *int32 *float64    -->   T -->
	var x int64 = 10
	y := int8(x) // int64 ---> int8
	x = int64(y) // int8  ---> int64
	x1 := &x     // x1 *int64
	// 不安全的指针，就是可以在指针上随意移动 + -
	unsafeP1 := unsafe.Pointer(x1) // *int64 ---> Pointer
	fmt.Println(unsafeP1)
	// Pointer ---> *int64 就是从不安全的变为安全的，就是不能随便移动(+ -)了
	intPtr1 := (*int64)(unsafeP1) // Pointer ---> *int64
	fmt.Println(intPtr1)
	fmt.Println(*intPtr1)
}

func TestSafePointer1(t *testing.T) {
	str := "pointer_test"
	a := &str
	fmt.Println(a)
	fmt.Println(*a)
}

func TestPonterT(t *testing.T) {
	var x int64 = 20
	a := &x // *int64
	fmt.Printf("%T", a)
	// 没有办法指针a的值进行如下转换: *int64 --> *uint64
}

func TestPonterT1(t *testing.T) {
	var x int64 = 20
	a := &x
	fmt.Println(a)

	var y uint64 = 20
	b := &y
	fmt.Println(b)

	// 我们不能进行 a = b 应该是为了保证安全，如果Pointer类型是不是就可以了？
}

type Man struct {
	Name string
	Age  int64
}

func TestUnsafePonter1(t *testing.T) {
	m := Man{Name: "John", Age: 20}
	fmt.Printf("%p\n", &m)
	// 编码unicode一个4个字节，4个16字节？
	fmt.Println(unsafe.Sizeof(m.Name), unsafe.Sizeof(m.Age), unsafe.Sizeof(m)) // 16 8 24
	fmt.Println(unsafe.Offsetof(m.Name))                                       // 0
	// unsafe.Offsetof 函数的参数必须是一个字段 x.f, 然后返回 f 字段相对于 x 起始地址的偏移量, 包括可能的空洞.
	// 内存对齐有关。因为name占用了16字节，因此age相当于差了16字节
	fmt.Println(unsafe.Offsetof(m.Age))
}

func TestUnsafePointer2(t *testing.T) {
	a := [3]int64{1, 2, 3}
	fmt.Printf("%p\n", &a)

	s1 := unsafe.Sizeof(a[0])
	fmt.Printf("%d\n", s1)

	// 刚开始地址执行数组的起始位置即0索引。但是+s1即+8字节后，相当于内存地址偏移，获取索引为1的元素
	p1 := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + s1))
	fmt.Println(*p1)
}

func TestUnsafePointer3(t *testing.T) {
	type T struct{ a int }
	var t1 T
	fmt.Printf("%p\n", &t1)                          // 0xc0000a0200
	println(&t1)                                     // 0xc0000a0200
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&t1))) // c0000a0200
}

type T struct {
	x bool
	y [3]int16 // 1 index
}

const (
	N = unsafe.Offsetof(T{}.y)
	M = unsafe.Sizeof(T{}.y[0])
)

func TestUnsafePointer4(t *testing.T) {
	t1 := T{y: [3]int16{123, 456, 789}}
	p := unsafe.Pointer(&t1)
	// "uintptr(p) + N + M + M"为t.y[2]的内存地址。 offset
	ty2 := (*int16)(unsafe.Pointer(uintptr(p) + N + M))
	fmt.Println(*ty2) // 789
}
