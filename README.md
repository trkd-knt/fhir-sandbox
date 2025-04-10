# FHIR CSV <-> JSON 変換 CLI

このツールは、FHIRのPatient / ObservationリソースをCSVとFHIR Bundle形式のJSON間で相互変換できるCLIです。FHIR非対応のツールなどとのデータ連携を目的としています。

## 🔧 実行方法

```bash
go run main.go [COMMAND] [...flags]
```

## 🧪 サンプルデータでの実行例

### JSON → CSV 変換
```bash
go run main.go patient --to-csv --input testdata/fhir_bundle_patient.json  --output outputs/values_patient.csv
go run main.go observation --to-csv --input testdata/fhir_bundle_observation.json --output  outputs/values_observation.csv
```

### CSV → JSON 変換
```bash
go run main.go patient --to-json --input testdata/flatten_data_patient.csv --output outputs/bundle_patient.json
go run main.go observation --to-json --input testdata/flatten_data_observation.csv --output outputs/bundle_observation.json
```

## 🩺 FHIR-mockと連携して動作確認

### FHIR Mock サーバの起動
```bash
docker-compose up -d
```

### データ登録
```bash
curl -X POST http://localhost:8080/fhir \
  -H "Content-Type: application/fhir+json" \
  -d @outputs/bundle_patient.json

curl -X POST http://localhost:8080/fhir \
  -H "Content-Type: application/fhir+json" \
  -d @outputs/bundle_observation.json
```

### 登録内容確認
```bash
curl -X GET http://localhost:8080/fhir/Patient
curl -X GET http://localhost:8080/fhir/Observation
```

### FHIR-mockから再エクスポート
```bash
curl -X GET http://localhost:8080/fhir/Patient \
  -H "Accept: application/fhir+json" \
  -o inputs/patients.json

curl -X GET http://localhost:8080/fhir/Observation \
  -H "Accept: application/fhir+json" \
  -o inputs/observation.json
```

### 再取得データをCSVに変換
```bash
go run main.go patient --to-csv --input inputs/patients.json  --output outputs/values_patient_2.csv
go run main.go observation --to-csv --input inputs/observation.json --output  outputs/values_observation_2.csv
```

## 💡 対応リソース
- [x] Patient
- [x] Observation
- [ ] その他リソース

## 🗂 ディレクトリ構成（簡略）

```bash
.
├── cmd/                # Cobra CLI 各コマンド
├── domain/             # FHIRモデル構造体定義
├── infrastructure/     # ファイル入出力(JSON/CSV)
├── usecase/            # 変換ロジック（flatten/expand）
├── testdata/           # 入出力サンプル
└── main.go             # エントリーポイント
```
