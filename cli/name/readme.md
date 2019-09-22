# Name 转换并格式化名称模块

将给定的字符串转换为名称对象，用于根据使用场景生成各种标准化的名称字符串

## 使用方式

    import  "github.com/herb-go/util/name"

    //通过New函数生成对象。第一个参数代表传入的字符串是否带有前缀的路径
    n,err:=name.New(true,"testname")

## 格式限制与转换规则

传入的字符串必须为字母,数字,点,空格,下划线,减号,必须以字母开头
包含路径的字符串还可以带 "/" 符号

空格,下划线,减号会被用作字符串切割符号

转换为驼峰格式时,会自动根据 go lint 将特殊名词转为全大写格式

## 转换后的结构

    type Name struct {
        Raw                         string // 原始字符串
        Parents                     string //父目录(前缀)
        ParentsList                 []string //分割后的父目录列表
        Title                       string //首字母大写格式
        Lower                       string //全小写格式
        Camel                       string  //首字母小写的驼峰格式
        Pascal                      string //首字母大写的驼峰格式
        LowerWithParentDotSeparated string //所有父目录以 . 连接的小写格式
        LowerWithParentPath         string //带路径的小写格式
        PascalWithParents           string //带父目录的首字母大写驼峰格式
    }

## 快速方法

    //将父目录,小写格式,传入的字符串列表 拼接为文件目录格式
    n.LowerPath("abc","def")
