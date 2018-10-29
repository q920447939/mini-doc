

## gin项目采坑经历

1. 未解决`jsonBind `无法绑定的问题,自己写了一个json转换(无法获取参数坑了好久)
2. 前端调用fa.ico文件,后台会拦截，暂时不知道怎么解决



​	





## gorm项目采坑经历

1. 想看sql的话  gorm包下,  `scope.go`文件    `callCallbacks`方法,`scope.Sql`

2. 暂时不知道怎么解决insert之后返回id的方式,至少框架是不支持的,但是又不想自己的SQL

3. gorm默认是会在表名后面加个s的,需要关闭的话调用  `db.SingularTable(true)` 

4. 新增的时候一定要使用指针,要不然插入不进入，至少我测试是这样。

5. ![123](./img/gin_img_create.jpg)

   写SQL最好加上Error方法,要不然根本不会报错，根本不知道是什么错误！虽然返回的是MySql的错误，但是总比没有好

6. 报错

`(sql: Scan error on column index 4: unsupported Scan, storing driver.Value type []uint8 into type *time.Time) `

解决办法:在mysql连接后面加上`charset=utf8&parseTime=True&loc=Local`