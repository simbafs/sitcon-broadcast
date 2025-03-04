SITCON 2024 製播組要用的東西

* Lighten talk 的倒數計時
* 直播右上角的字卡

# TODO

- [ ] 動態引入 session.json，有個參數可以調整 session.json 的路徑
- [ ] 更新字卡頁面樣式
- [ ] 需要一個 debug 頁面
- [ ] 新增一個議程表頁面
- [ ] 移除 session.json 中的「休息」，動態補齊空洞
- [ ] 持久化儲存 session.json
- [ ] 統一的設定方式（env、args...）
- [ ] 寫死時區
- [ ] time.Time 型別在 JS 接收、發送
- [ ] Countdown 的時間對齊問題
- [ ] 統一 Go 到底怎麼把 time.Time 轉成字串
- [ ] linked list 是錯誤的，因為依照議程廳不同，前後關係也會不同

# Known Bugs
* 字卡好像不會動
* 有 broadcast 的議程要同步更新時間
 
