# 定时任务
# 安装gf工具
  wget https://goframe.org/cli/linux_amd64/gf && chmod +x gf && ./gf install
# 更新swagger文档
  更新时记得删除 boot/data-swagger.go文件,没有就忽略
  gf swagger --pack
# 访问接口文档 127.0.0.1：8081/swagger    

# gooxml 使用
## 教程地址
   https://github.com/carmel/gooxml
   https://github.com/amanbolat/gooxml-textbox-test  word里面列表解析demo
## 安装gooxml 解析word文档
  先修改代理，不然下载报错
  export GOPROXY=https://goproxy.cn 
  go get baliance.com/gooxml/
  go build  baliance.com/gooxml/…

## 安装 github.com/ledongthuc/pdf   解析pdf文档
    go get -u github.com/ledongthuc/pdf 
## 注意事项
  相对路径会报错，用绝对路径
  windows 路径要用 \\  
## 生成代码 修改config/config.toml 里面的gfcli.gen.dao
  gf gen dao

