SITCON 2024 製播組要用的東西

-   Lighten talk 的倒數計時
-   直播右上角的字卡

# TODO

-   [ ] 動態引入 session.json，有個參數可以調整 session.json 的路徑
-   [x] 更新字卡頁面樣式
-   [x] 需要一個 debug 頁面
-   [ ] 新增一個議程表頁面
-   [x] 移除 session.json 中的「休息」，動態補齊空洞
-   [x] 持久化儲存 session.json
-   [ ] 統一的設定方式（env、args...）
-   [x] 寫死時區
-   [x] time.Time 型別在 JS 接收、發送
-   [ ] Countdown 的時間對齊問題
-   [x] 統一 Go 到底怎麼把 time.Time 轉成字串
-   [x] linked list 是錯誤的，因為依照議程廳不同，前後關係也會不同
-   [ ] SSE countdown-room 頻道不要一直發更新
-   [ ] /card/admin 修改時間更新太慢
-   [x] 字卡好像不會動
-   [x] 有 broadcast 的議程要同步更新時間
-   [ ] 有些字卡字太多，會影響大小
