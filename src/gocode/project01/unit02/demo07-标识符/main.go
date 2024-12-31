/*
标识符：只要是起名的地方都叫标识符；包名、变量名、常量名、函数名、类型名等
1. 标诗符的命名规范
	1.1 由26个英文字母大小写，0-9，_组成
	1.2 数字不可以开头
	1.3 Golang中严格区分大小写
	1.4 标识符不能包含空格
	1.5 下划线"_"本身在Go中是一个特殊的标识符，称为空白标识符。可以代表任何其他标识符，但是它对应的值会被忽略(黑洞)。
	1.6 不能以系统保留关键字作为标识符
		关键字：break default func interface select case defer go map struct chan else goto package switch const fallthrough if range type continue for import return var
2. 标识符的命名注意事项
	2.1 区分大小写
	2.2 标识符不能包含空格
	2.3 不能以数字开头


3、驼峰命名法 （利用首字母的大小写完成权限控制
	3.1 首字母大写：在Go中，如果一个标识符的首字母是大写的，那么它就是对外可见的(暴露的，公有的)
	3.2 首字母小写：如果一个标识符的首字母是小写的，那么它就是对外不可见的(私有的)

	
4、GOPATH：一套代码的公用路径。从项目绝对路径的src开始，到项目的根目录结束
	例如绝对路径：D:\Code\goProject\src\gocode\project01\unit02\demo06
	需要在本机的环境变量处配置：GOPATH=D:\Code\goProject

5、导包
	5.2 导入的包名必须是全路径，但是因为配置了GOPATH，可以从从src开始写
	5.3 如果导入的包没有使用，代码编译不会报错
	5.4 如果导入的包使用了，但是没有引用，代码编译不会报错
	5.5 如果导入的包使用了，但是引用了包中的函数，但是没有使用，代码编译会报错
	5.6 如果导入的包使用了，但是引用了包中的函数，但是没有使用，可以使用_下划线忽略导入包中的函数
	5.7 如果导入的包使用了，但是引用了包中的函数，但是没有使用，可以使用_下划线忽略导入包中的函数


6、Go语言中的关键字 25个
	break	default	func	interface	select
	case	defer	go	map	struct
	chan	else	goto	package	switch
	const	fallthrough	if	range	type
	continue	for	import	return	var
7、Go语言中的预定义标识符 36个
	append	bool	byte	cap	close	complex	complex64	complex128	uint16
	copy	false	float32	float64	imag	int	int8	int16	uint32
	int32	int64	iota	len	make	new	nil	panic	uint64
	print	println	real	recover	string	true	uint	uint8	uintptr
8、Go语言中的内置函数 25个
	append	bool	byte	cap	close	complex	complex64	complex128	uint16
	copy	false	float32	float64	imag	int	int8	int16	uint32
	int32	int64	iota	len	make	new	nil	panic	uint64
	print	println	real	recover	string	true	uint	uint8	uintptr
9、Go语言中的常量 6个
	true	false	iota	nil
10、Go语言中的类型 29个
	int	int8	int16	int32	int64
	uint	uint8	uint16	uint32	uint64
	float32	float64	complex64	complex128
	bool	byte	rune	string	error
	uintptr
11、Go语言中的默认值
	数值类型：0
	布尔类型：false
	字符串类型：""
	指针类型：nil

*/