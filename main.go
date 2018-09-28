package main
import (
    "bufio"
    "fmt"
    // "io"
    "time"
    "os"
    "strings"
    "github.com/kataras/iris"
    "github.com/kataras/iris/middleware/logger"
    "github.com/kataras/iris/middleware/recover"
)

func main() {

    app := iris.New()
    app.Logger().SetLevel("debug")
    app.Use(recover.New())
    app.Use(logger.New())

    app.Get("/ping", func(ctx iris.Context) {
        // 检查并创建数据清洗文件是否存在？
        MakeFile()
        // 读取日志文件并开始清洗转储
        DealWith()
        // fmt.Println(ReadLine(4))
        file, _ := os.Open("./test.txt")
        fileScanner := bufio.NewScanner(file)
        for fileScanner.Scan(){
            ctx.WriteString(fileScanner.Text())
        }
        defer file.Close()

    })
    app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}




// 创建文件

func MakeFile() {
    _dir := "test.txt"
    exist, err := PathExists(_dir)
    if err != nil {
        fmt.Printf("文件失败![%v]\n", err)
        return
    }

    if exist {
        fmt.Printf("文件已经存在![%v]\n", _dir)
    } else {
        fmt.Printf("没有文件![%v]\n", _dir)
        // 创建文件夹
        file,err:=os.Create(_dir)
        file.Close()
        if err != nil {
            fmt.Printf("创建文件失败![%v]\n", err)
        } else {
            fmt.Printf("创建文件成功!\n")
        }
    }

}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}

// 清洗数据并增量的读取清洗数据
func DealWith() {
    file, _ := os.Open("./test.txt")
    fileScanner := bufio.NewScanner(file)
    lineCount := 0
    for fileScanner.Scan(){
        lineCount++
    }
    defer file.Close()
    fmt.Println(lineCount)
    ReadLine1(lineCount)

}

func ReadLine1(lineNumber int){
    time.Sleep(time.Duration(3)*time.Second)
    file, _ := os.Open("./access.log")
    fileScanner := bufio.NewScanner(file)
    lineCount := 0
    for fileScanner.Scan(){
        if lineCount >= lineNumber  {
            var h[] string
            // fmt.Println(fileScanner.Text())
            fileScanner.Text()
            h=strings.Split(string(fileScanner.Text()), "*")
            // fmt.Println(h)

            InputFile(h)
        }
        lineCount++
    }
    defer file.Close()
}


func InputFile(str_content[] string){
    fd,_:=os.OpenFile("./test.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
    // var B string
    InputString:=str_content[0]+"\t"+str_content[1]+"\t"+str_content[2]+"\t"+str_content[3]+"\t"+str_content[4]+"\t"+str_content[5]+"\t"+str_content[6]+"\t"+str_content[7]+"\t"+str_content[8]+"\t"+str_content[9]+"\t"+str_content[10]+"\t"
    fd_content:=strings.Join([]string{InputString,"\n"},"")
    buf:=[]byte(fd_content)
    fd.Write(buf)
    fd.Close()
    fmt.Println("祝好！文件写入成功！")
}

