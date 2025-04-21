# Slot, Signal 和 Property

Dotherside 封装了 Qt 的 Slot 和 Signal，并支持 Slot、Singal、Property 三种接口。

其中 Slot 和 Signal 和 Qt 中的概念一致，声明一个 struct 时，可以标记一个方法为 Slot，用于响应信号。

也可以标记一个方法为 Singal，其作用是发送信号给 Dothreside / Qt，触发相同名称的 Slot。

Property 可以看作是 Slot 和 Signal 的结合体，用于读写属性：
- reader: 一个 slot，用于读取属性值
- writer: 一个 slot，用于设置属性值
- emitter: 一个 signal，用于通知属性值变化

本项目 goqml 实现了 Dotherside 的 Slot、Signal 和 Property。

## 语法

1. slot: 可以通过 `// @goqml.slot` 或者 `// @goqml.slot("my-slot")` 注释方法来标记一个方法为 Slot。
后者可以指定 Slot 的名称，默认为方法名。

2. signal: 可以通过 `// @goqml.signal` 或者 `// @goqml.signal("my-signal")` 注释方法来标记一个函数类型的 field 为 Signal。

3. property: 可以通过 `// @goqml.property` 或者 `// @goqml.property("my-property")` 注释一个 field 为 Property。也可以通过 `// @goqml.property("name").getter` 来指定 Property 的 getter 对应的方法，setter / emitter 同理，当然，也可以指定对应的名称。

## 生成规则

### 约束检查
- signal 方法不能有返回值
- signal 方法不能有函数体
- property 指定 getter / setter / emitter 时，同 property 的核心类型保持一致，检查对应的方法类型

### 解析
结构体需要解析的值：
- Struct 的名称 structName
- Struct 的父 Struct 类型 parentStructName

每个 slot 需要解析的值：
- slot 的方法名称
- slot 的名称，默认是 slot 方法名，也可以是注释中指定的名称
- slot 的参数列表
- slot 的返回值类型

每个 signal 需要解析的值：
- signal 方法名称
- signal 名称，默认是 signal 方法名，也可以是注释中指定的名称
- signal 的参数列表

每个 field 型的 property：
- property 名称
- property 核心类型

每个方法型 property 需要解析的值：
- property 名称
- property 核心类型
- getter / setter / emitter 的方法名称和对应的名称

### 生成
1. 每个 signal 生成一个 go 方法
    - 函数名是 goqml+方法名称，注意原函数名的首字母改为大写
    - 函数体调用 struct 的 Emit 方法，第一个实参是 signal 名称，然后将函数头中传入的参数逐个调用 qoaml.NewQVariant 传入
2. 每个 signal 的函数体通过 plan 9 生成到 s 文件中
    - 函数实现将所有参数转发给 goqml+方法名称 这个方法
3. 每个 property 和 signal 一样生成一个 goqml 前缀的方法和用于转发的汇编
4. 构造一个 static+Struct名称+QMetaObject 变量作为这个类型的 QMetaObject 变量和 一个 StaticQMetaObject 方法返回这个变量
5. 构造一个 OnSlotCalled 方法，参数是 slot 的名称和 Property 的 setter 和 getter，根据参数类型传入参数，根据返回值情况设置返回值

## 例子

以下结构体 MyStruct，用来演示 Slot、Signal 和 Property 的用法。

```go
type MyStruct struct {
    goqml.QObject

    // @goqml.property
    name string

    // @goqml.property("p2")
    prop2 int
}

// @goqml.slot
func (s *MyStruct) MySlot(a int, b string) {}

// @goqml.slot("my-slot")
func (s *MyStruct) MySlot2(c float64) int {
    return 0;
}

// @goqml.signal
func (s *MyStruct) MySignal()

// @goqml.signal("my-signal")
func (s *MyStruct) MySignal2(int)

// @goqml.property("p1").getter
func (s *MyStruct) MyProperty() int {
    return 1
}

// @goqml.property("p1").setter("qiekenao")
func (s *MyStruct) SetMyProperty(int) {}

// @goqml.property("p1").emitter
func (s *MyStruct) MyPropertyChanged(int)
```

将会被扩展

```go
var staticMyStructQMetaObject = goqml.NewQMetaObject(
    (goqml.QObject*)(nil).staticMetaObject(),
    "MyStruct",
    []*goqml.SignalDefinition{
        goqml.NewSignalDefinition("MySignal"),
        goqml.NewSignalDefinition("my-signal", goqml.QMetaTypeInt),
    },
    []*goqml.SlotDefinition{
        goqml.NewSlotDefinition("MySlot", goqml.QMetaTypeVoid, goqml.QMetaTypeInt, goqml.QMetaTypeString),
        goqml.NewSlotDefinition("my-slot", goqml.QMetaTypeInt, goqml.QMetaTypeFloat64),
    },
    []*goqml.PrpertyDefinition{
        goqml.NewPropertyDefinition("p1", goqml.QMetaTypeInt, "getP1", "qiekenao", "p1Changed"),
        goqml.NewPropertyDefinition("p2", goqml.QMetaTypeInt, "getP2", "setP2", "p2Changed"),
        goqml.NewPropertyDefinition("name", goqml.QMetaTypeString, "getName", "setName", "nameChanged"),
    },
)

func (s *MyStruct) StaticQMetaObject() *QMetaObject {
    return staticMyStructQMetaObject
}

func (s *MyStruct) goqmlMySignal() {
    goqml.QObject.Emit("MySignal")
}

func (s *MyStruct) goqmlMySignal2(value int) {
    goqml.QObject.Emit("my-signal", value)
}

func (s *MyStruct) goqmlMyPropertyChanged(value int) {
    goqml.QObject.Emit("MyPropertyChanged", value)
}

func (s *MyStruct) goqmlNameChanged(value string) {
    goqml.QObject.Emit("nameChanged", value)
}

func (s *MyStruct) goqmlP2Changed(value int) {
    goqml.QObject.Emit("p2Changed", value)
}

func (s *MyStruct) OnSlotCalled(slotName string, arguments []*QVariant) {
    switch slotName {
    case "MySlot":
        s.MySlot(arguments[1].ToInt(), arguments[2].ToString())
    case "my-slot":
        arguments[0].SetInt(s.MySlot2(arguments[1].ToFloat64()))
    case "getP1":
        arguments[0].SetInt(s.MyProperty())
    case "qiekenao":
        s.SetMyProperty(arguments[1].ToInt())
        s.MyPropertyChanged(arguments[1].ToInt())
    case "getName":
        arguments[0].SetString(s.name)
    case "setName":
        s.name = arguments[1].ToString()
        s.goqmlNameChanged(s.name)
    case "getP2":
        arguments[0].SetInt(s.prop2)
    case "setP2":
        s.prop2 = arguments[1].ToInt()
        s.goqmlP2Changed(s.prop2)
    default:
        fmt.Println("unknown slot:", slotName)
    }
}
```

有一个 plan9 汇编文件，实现 signal 的结构体，转发给 goqml 前缀的函数

```plan9
#include "textflag.h"

TEXT ·MyStruct·MySignal(SB), NOSPLIT, $0-16
    CALL ·MyStruct·goqmlMySignal(SB)
    RET 
```
