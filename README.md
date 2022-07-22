# octopus-meta
<a href="https://goreportcard.com/report/github.com/JJApplication/octopus_meta"><img src="./copyright/goreport.svg" /></a>
<a><img src="./copyright/version.svg"/></a>
<a href="http://service.renj.io"><img src="./copyright/renj.io.svg"/></a>
<a href="https://github.com/JJApplication"><img src="./copyright/copyright-JJService.svg"/></a>

octopus-meta是Apollo运行时的App 模型定义

## 拥有完整的测试用例
```bash
$ bash ./test.sh
```
## 模型
`App | Meta` 为octopus-meta定义的app模型

## 对外接口
### AutoLoad
自动加载meta文件， 当环境变量`APP_ROOT`存在并且`$APP_ROOT/.octopus`或`$APP_ROOT/meta`存在时
会自动从此目录加载模型文件

### Load(p string)
传入指定的模型文件目录并加载模型文件

### SetOctopusMetaDir(p string)
配置全局生效的模型文件路径，会在Load接口中生效，不影响Autoload逻辑

### Octopus{}
Octopus结构体暴露了内部的json模型加载方法
```bash
var OctopusIterator = Octopus{Type: "default", AutoEnv: true}
```
在octopus-meta内部定义一个全局的OctopusIterator默认使用，通过`AutoEnv`可以控制meta的值是否支持从环境变量加载

