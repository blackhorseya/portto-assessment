# Meme Coin API

## 專案簡介

本專案是一個 Meme Coin API 後端應用程式，使用 Golang、Gin 框架及 GORM 作為 ORM，並支援 Docker 部署。

## 專案結構

## 執行方式

在專案根目錄下建立 `.env` 檔案，並填入以下內容：

```env
DB_HOST=localhost
DB_PORT=5432
DB_PASSWORD=password
DB_USER=postgres
DB_NAME=portto
```

### 本地環境

1. 安裝 Go 1.20+。
2. 安裝依賴：
   ```bash
   go mod tidy
   ```
3. 啟動 PostgreSQL 資料庫：
   ```bash
   docker compose up -d db
   ```
4. 執行應用程式：
   ```bash
   go run . run --verbose
   ```

### Docker 部署

1. 確保已安裝 Docker。
2. 執行以下命令來建構 Docker 映像檔：
   ```bash
   docker compose --build --remove-orphans up
   ```

## 設定與配置

## API 說明

[swagger](http://localhost:8080/api/docs/index.html)

- **POST /api/v1/coins**  
  建立一個新的 Meme Coin。  
  **輸入**：`name`（必填，唯一）、`description`（選填）

- **GET /api/v1/coins/:id**  
  根據 ID 取得 Meme Coin 詳細資訊。  
  **輸出**：`name`、`description`、`created_at`、`popularity_score`

- **PATCH /api/v1/coins/:id**  
  更新 Meme Coin 的描述。  
  **輸入**：`description`（必填）

- **DELETE /api/v1/coins/:id**  
  根據 ID 刪除 Meme Coin。

- **POST /api/v1/coins/:id/poke**  
  Poke Meme Coin，使其人氣分數增加（預設 +1）。
