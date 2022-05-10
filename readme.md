# qweqwe.dev

**前言**<br>
 　　某天刷Bilibili時看到[這玩意](https://www.bilibili.com/video/BV1hv411G7iH)，感覺還挺有趣，於是用golang寫了一個半復刻版，當作是對gin框架的練習。已經完成隨機單字陣列、隨機句子陣列，隨機段落陣列、依據圖片大小與檔案後綴(現支援jpg和png)生成圖片。

**啟動專案**

```
cd qweqwe

go run main.go
```

**如何使用**

1. 返回lorem單字陣列
    ``http://localhost:1234/words/n``
    n為需要多少組，最低一組
    例: ``http://localhost:1234/words/10``
2. 返回lorem句子陣列
    ``http://localhost:1234/sentence/n``
    n為需要多少組，最低一組
    例: ``http://localhost:1234/sentence/10``
3. 返回lorem段落陣列
    ``http://localhost:1234/paragraph/n``
    n為需要多少組，最低一組
    例: ``http://localhost:1234/paragraph/10``
4. 返回佔位圖
    ``http://localhost:1234/image/widthxheight.filepostfix``
    返回檔名為``widthxheight.filepostfix``的圖片
    width: 寬，height: 高，filepostfix: jpg或png二選一
    例1: ``http://localhost:1234/image/800x600.jpg``
    例2: ``http://localhost:1234/image/600x800.png``
5. 陣列洗牌
    ``http://localhost:1234/random/arrayitem``
    arrayitem以逗號分隔
    例1: ``http://localhost:1234/random/1,2,3,4,5,6,7,8``
    例2: ``http://localhost:1234/random/qw,ew,cx,ar,we,yg,zi,ks``