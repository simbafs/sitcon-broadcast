# SITCON 製播組工具

📍 **專案網址**：[https://sitcon.1li.tw](https://sitcon.1li.tw)  
📦 **Docker Image**：[sitcon-broadcast container](https://github.com/simbafs/sitcon-broadcast/pkgs/container/sitcon-broadcast)  
🧱 **Docker Compose**：[docker-compose.yml](https://github.com/simbafs/sitcon-broadcast/blob/main/docker-compose.yml)

---

## 📌 專案架構

此專案為 **前後端分離架構**：

-   **前端**：使用 Next.js
-   **後端**：使用 [gin](https://github.com/gin-gonic/gin/) + [huma](https://huma.rocks/)
-   **開發時**：Gin 會轉發未知路由到 Next.js（:3001）
-   **正式部署時**：將網頁、JS、CSS 靜態資源打包進 Go 程式（使用 `go embed`）

---

## ⚙️ 環境變數（Config）

| 變數     | 預設值               | 說明                    |
| -------- | -------------------- | ----------------------- |
| `Addr`   | `:3000`              | Server 監聽的位址       |
| `Token`  | `token`              | 存取部分頁面所需的權杖  |
| `Domain` | `localhsot`          | Cookie 的 `domain` 欄位 |
| `DB`     | `./data/data.sqlite` | SQLite 資料庫的路徑     |

---

## 🧑‍💻 開發指令（使用 `make`）

| 分類           | 指令                 | 功能描述                                             |
| -------------- | -------------------- | ---------------------------------------------------- |
| 🧪 Development | `make dev`           | 啟動前後端開發伺服器（使用 tmux 分窗）               |
|                | `make devBackend`    | 啟動後端開發伺服器（使用 `nodemon` 自動重啟）        |
|                | `make devFrontend`   | 啟動前端開發伺服器（Next.js）                        |
| 📦 Dependency  | `make dep`           | 安裝前後端相依套件                                   |
|                | `make depBackend`    | 安裝後端相依套件（go mod）                           |
|                | `make depFrontend`   | 安裝前端相依套件（使用 `pnpm`）                      |
| 🛠 Build       | `make build`         | 同時建置前後端                                       |
|                | `make buildFrontend` | 建置前端（執行 `frontend/build.sh`）                 |
|                | `make buildBackend`  | 建置後端並嵌入前端靜態檔案至 Go 程式                 |
|                | `make buildDist`     | 使用 Docker 建置最終可部屬版本                       |
| 🧹 Maintenance | `make clean`         | 清除建置產物與暫存資料夾                             |
|                | `make format`        | 使用 `prettier` 與 `gofmt` 格式化前後端程式碼        |
| 🔍 Others      | `make doctor`        | 檢查必要工具是否已安裝（如 `go`、`pnpm`、`tmux` 等） |
|                | `make help`          | 顯示所有指令說明                                     |
|                | `make session`       | 產生 `sessions.json` 檔案（用於卡片資料）            |
|                | `make checkFrontend` | 使用 knip 檢查 frontend 未使用的程式碼               |
|                | `make ent`           | 重新產生 ent model                                   |
|                | `make staticcheck`   | 使用 `staticcheck` 和 `errcheck` 分析後端程式碼      |

> [!TIP]
> 📌 **註**：如需格式化程式碼請先安裝 `prettier` 與 `gofmt`。

---

## 🌐 前端注意事項

-   **不使用 SSR**（部署時為純靜態檔案）
-   **Hooks / Components 放置原則**：
    -   僅在單一頁面使用的 hook/component → 放在該頁附近
    -   跨頁共用的 → 分別放在：
        -   `frontend/src/hooks/`
        -   `frontend/src/components/`
-   除了 `function Page` 外，**請勿使用 `export default`**
-   撰寫與 API 的互動邏輯時，請依下列順序：
    1. 在 `frontend/src/sdk/` 建立對應的 API 呼叫函式
    2. 如果需要，可在 `frontend/src/hooks/` 建立對應 hook
    3. 最後在 component 中使用該 hook

---

## 🛠️ 後端結構規範

資料流與資料夾結構應遵循以下規則：

1. **資料表定義**：`backend/ent/scheme`
2. **資料庫操作邏輯**：`backend/models`（透過 ent 存取資料庫）
3. **API 處理邏輯**：`backend/api`
> [!CAUTION]
>  **`backend/api` 只能使用 `backend/models`，不可直接操作 ent**
4. 原則上 SSE 的 send 只能在 `model/` 下使用
