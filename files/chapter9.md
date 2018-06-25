### 模式九 外观模式
```go
type Facade struct {
    M Music
    V Video
    C Count
}

func (this *Facade) GetRecommandVideos() error {
    this.V.GetVideos()
    this.C.GetCountByID(111)

    return nil
}

type Music struct {
}

func (this *Music) GetMusic() error {
    fmt.Println("get music material")
    // logic code here
    return nil
}

type Video struct {
    vid int64
}

func (this *Video) GetVideos() error {
    fmt.Println("get videos1")
    return nil
}

type Count struct {
    PraiseCnt  int64 //点赞数
    CommentCnt int64 //评论数
    CollectCnt int64 //收藏数
}

func (this *Count) GetCountByID(id int64) (*Count, error) {
    fmt.Println("get video counts")
    return this, nil
}

func main() {
    f := &Facade{}
    f.GetRecommandVideos()
}
```

### 总结

>它为一套复杂的调度子系统提供一个统一接入接口，外部所有的对子系统的调用都通过这个外观接口接入，统一管理

>缺点：外观接口的会不断膨胀，代码量会不断增加
