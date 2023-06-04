package main

import "fmt"

// append无论如何都是从slice的尾部开始追加数据; 如果有append操作，很可能会引发扩容，要特别注意

func main() {

	sli := make([]string, 1)

	sli[0] = "宋江"

	fmt.Printf("[main]原始sli为%#v,长度:%d,容量:%d,内存地址:%p\n", sli, len(sli), cap(sli), sli)
	// [main]原始sli为[]string{"宋江"},长度:1,容量:1,内存地址:0x14000010230

	f2(sli)

	fmt.Printf("[main]调用f2()之后sli为%#v,长度:%d,容量:%d,内存地址:%p\n", sli, len(sli), cap(sli), sli)
	// [main]调用f2()之后sli为[]string{"晁盖"},长度:1,容量:1,内存地址:0x14000010230

	// 可见，只可能会影响底层数组的值，**不会影响长度和容量**
}

func f2(sli1 []string) []string {

	fmt.Printf("[f2]f2中append之前sli1的长度%d,容量%d,内存地址%p\n", len(sli1), cap(sli1), sli1)
	// [f2]f2中append之前sli1的长度1,容量1,内存地址0x14000010230

	sli1[0] = "晁盖" // 此时没有扩容，sli1和main中的sli地址一样，修改sli1[0]自然会影响main

	sli1 = append(sli1, "卢俊义", "吴用", "公孙胜", "关胜")

	// 如果将上面的sli1[0] = "晁盖"去掉，而在下方赋值，此时sli1和main中的sli内存地址不同，此时再修改sli1[0]不会影响到main
	//sli1[0] = "晁盖"

	fmt.Printf("[f2]f2中append之后sli1的长度%d,容量%d,内存地址%p\n", len(sli1), cap(sli1), sli1)
	// [f2]f2中append之后sli1的长度5,容量5,内存地址0x14000070050
	return sli1
}

// append一定会改变原始slice的内存地址吗? 不一定，不发生扩容就不会改变~
