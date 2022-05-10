# qweqwe.dev

**前言**<br>
 　　某天刷Bilibili時看到[這玩意](https://www.bilibili.com/video/BV1hv411G7iH)，感覺還挺有趣，於是用golang寫了一個半復刻版，當作是對gin框架的練習。

**已完成功能**<br>

1. 依據給定長度取得隨機**lorem**單字陣列
2. 依據給定長度取得隨機**lorem**句子陣列
3. 依據給定長度取得隨機**lorem**段落陣列
4. 依據所給的width和height與副檔名(現支援jpg與png)生成圖片
5. 給出一些以逗點分隔的元素，取得打散後的陣列

**啟動專案**

```
cd qweqwe

go run main.go
```

**如何使用**

1. 返回lorem單字陣列<br>
    ``http://localhost:1234/words/n``<br>
    n為需要多少組，最低一組<br>
    例: ``http://localhost:1234/words/10``<br>
2. 返回lorem句子陣列<br>
    ``http://localhost:1234/sentence/n``<br>
    n為需要多少組，最低一組<br>
    例: ``http://localhost:1234/sentence/10``<br>
3. 返回lorem段落陣列<br>
    ``http://localhost:1234/paragraph/n``<br>
    n為需要多少組，最低一組<br>
    例: ``http://localhost:1234/paragraph/10``<br>
4. 返回佔位圖<br>
    ``http://localhost:1234/image/widthxheight.filepostfix``<br>
    返回檔名為``widthxheight.filepostfix``的圖片<br>
    width: 寬，height: 高，filepostfix: jpg或png二選一<br>
    例1: ``http://localhost:1234/image/800x600.jpg``<br>
    例2: ``http://localhost:1234/image/600x800.png``<br>
5. 陣列洗牌<br>
    ``http://localhost:1234/random/arrayitem``<br>
    arrayitem以逗號分隔<br>
    例1: ``http://localhost:1234/random/1,2,3,4,5,6,7,8``<br>
    例2: ``http://localhost:1234/random/qw,ew,cx,ar,we,yg,zi,ks``<br>