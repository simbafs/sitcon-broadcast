SITCON 2024 製播組要用的東西

# Features

-   Lighten talk 的倒數計時
-   直播右上角的字卡

# Tech

## Frontend

-   nextjs
-   tailwindcss

## Backend

-   Go
-   entgo

# TODOs

-   [ ] 超級後台（任意修改 DB）
-   [ ] 圖形化展示現在狀況
-   [ ] 倒數計時

# 更新資料庫流程：
1. user 發出 Get /api/event/{name}/session，server 從 event.URL 抓 sessions.json 回傳給 user
2. user 根據 event.Script 處理，並將處理好的 sessions 用 PUT /api/session 寫入伺服器
3. user 可以修敢 event.Script，並用 PUT /api/event/{name} 更新 event.Script
