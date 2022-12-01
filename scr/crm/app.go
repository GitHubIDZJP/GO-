package main //声明main包，表明当前是一个可执行程序
import (
	//导入外部包-->模块函数名必须大写开头
	"crm/function1"
	"crm/function2"
	"crm/ginModule"
)

//"GO学习/scr/crm/test/common/message"

// main函数 程序执行的入口
func main() {

	//s1 := "字符串"
	//s2 := "拼接"
	//s3 := s1 + s2
	//fmt.Println(s3) //s3 = "打印字符串"

	//CEMTERComMain_sd()
	//调用外部的函数
	function1.Function1000()
	function2.Function2()
	//gin官网:

	//ginModule.Redirection() //重定向
	//ginModule.PostSendHttpRedirection() //从 POST 发出 HTTP 重定向
	//	ginModule.SenderRouterRedirection() //发出路由器重定向，使用HandleContext

	//ginModule.CEMTERComMain()//gin框架中的 basicauth 认证模块
	//ginModule.A_file_upload() //单文件上传
	//ginModule.Multiple_file_upload() //多个文件上传
	//ginModule.HttpEncrypt() //加密,搞不懂
	//ginModule.SetOrFetchCookie() //设置并获取cookie
	//ginModule.ProvideStaticFile() //加载模版文件夹里的HTML
	//ginModule.FeedDataFromTheReader() //从阅读器提供数据
	//ginModule.SecureJSON() //安全json---->
	//ginModule.SecureJSONExample() //secure(安全)json---->写了个Api
	//ginModule.RunMultipteServer() //运行多个服务

	//ginModule.SearchStrParameter() //查询字符串参数
	//ginModule.ginNew() //中间件
	//ginModule.InitCoroutine() //协程学习
	//ginModule.CoroutinesWithinMiddleware() //在中间件中使用协程 gorouetine

	//ginModule.MultipleTemplatePlate() //模版
	//ginModule.MultipleTemplateAdvanced() //多模版进阶
	//ginModule.BindOnlyInQueryString() //只绑定查询字符串
	//ginModule.BindOnlyInQueryStringTest() //只绑定查询字符串
	//ginModule.CodedForm() //多部分/Urlen编码表单
	//
	//ginModule.URLCodingBind() //多部分/URL 编码绑定-->未授权

	//ginModule.Model_binding_and_validation() //模型绑定和验证-->只有第三个方法是可以行得通的，

	//ginModule.MapStrPostFormParam() //映射为查询字符串或 postform 参数（回传数组参数）
	//ginModule.MapStrPostFormParamTest() // 映射为查询字符串(测试-map)
	//ginModule.MapStrPostFormParamPostFormArray() //映射为查询字符串或 postform 参数（array）

	//ginModule.GoLangTemplateDemo()
	//ginModule.HTMLRendering() //html渲染
	//ginModule.CustomTemplate() //HTML自定义模版
	//ginModule.WritingLog() //写log文件
	//ginModule.CustomLogFile() //自定义日志文件 ---->调用router.Use(gin.LoggerWithFormatter(logFormat))
	//ginModule.FroupRouter() //分组路由
	//ginModule.RrstartOrStop() //优雅重启或停止
	//ginModule.EndlessShutDownService() //使用endless 来优雅重启

	//ginModule.DefineFormatofTheRouteLog() //定义路由日志的格式
	//ginModule.MainCustomValidators() //自定义验证器(预定)
	//ginModule.CustomHttpConfig()

	//ginModule.MultipleApi() //一个函数里多个Api，直接改请求参数就行
	//ginModule.XML_JSON_YAML_ProtoBuf() //XML/JSON/YAML/ProtoBuf 渲染
	//ginModule.RenderingMain() //XML/JSON/YAML/ProtoBuf 渲染
	//ginModule.PureJSON() //纯JSON
	//ginModule.ParamtersPath() //路径中的参数//--没搞定有啥用

	//ginModule.QueryAndPublishForms() //查询和发布表单--对应gin官网的Query 和 post form
	//ginModule.DefaultQuery()//c.DefaultQuery上传
	//ginModule.PureJSON() // PureJSON可以使特殊字符正常显示
	//ginModule.CodedForm() //Multipart/Urlencoded 表单---->c.DefaultPostForm--
	//ginModule.MultipartUrl_Not_Bind() //Multipart/Urlencoded 未绑定
	//ginModule.MultipartUrlBind() //Multipart/Urlencoded 绑定
	//ginModule.JSONP() //使用 JSONP跨域

	ginModule.ServerPush() //HTTP2 服务器推送
	//语法学习
	//grammarLearning.Init_variable() //变量
	//grammarLearning.OpenLocalFile()//打开本地文件，并且在终端打印文本内容
	//grammarLearning.Init_constant() //常量
	//grammarLearning.MainApi() //结构体

	//gin框架demo
	//ginDemo.MainStudentApi() //学生接口
	//ginDemo.Api01()
	//ginDemo.RootMain()
	//ginDemo.BashRequestFramework()// 基础请求框架
	//ginDemo.MainLogin()

	/**
	注意事项1：执行顺序只从上到下，同端口时 下面的会覆盖上面的
	前提是:每新建一个类来写一个接口
	网络Api,但是一旦只要端口号都是8080，回调方法按顺序来，哪个在下面哪个就执行，会直接覆盖上面的
	*/
	/*
	 注意事项2：执行顺序只从下到上，同端口时 上面的会覆盖下面的(run时只运行最上面的接口)
	 前提是：接口写一个类里才行
	*/
	//ginModule.AsciiJSON_Api()

	ginModule.QueryParams()

	//bashRequestFramework() //基础网络框架
	//post 请求Api

	//postRequest()
	//gorm1() //gorm快速入门--增删改查
	//gorm2() //gorm声明model
	//gorm3()//gorm模型使用
	//gorm4() //gorm模型加强
	//gorm5() //gorm模型加强1
	//gorm_batch_insert()
	//gorm_Creating_Record() //gorm样板
	//gorm_association_relationship_between_tables() //表之间关联关系
	//gorm_association_gl() //表之间关联关系1

	//官网:

	//mainKoyeb() //Koyeb
	//rootMain()

	//t1 := time.Time{}
	//t2 := time.Now()

	//create_model_and_generate_table() //创建模型和生成表
	//PartialFieldPropertiesAreInserted() //部分字段属性插入
	//t1Second := t1.Unix()
	//t2Second := t2.Unix()
	//fmt.Printf("时间1:%v, 时间1Second:%v \n", t1, t1Second) // UTC时间1970-01-01 08:00:50的秒数为0,所以在此时间之前的为负值,且在int32范围之外
	//fmt.Printf("时间2:%v, 时间2Second:%v \n", t2, t2Second)
	//fmt.Println("t2Second-t1Second=", t2Second-t1Second)

}
