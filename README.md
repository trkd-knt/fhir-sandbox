# fhir-sandbox

## スキーマの参照元

```
https://hl7.org/fhir/R4/patient.html
```

## 起動
```
cp .env.sample .env
docker-compose up --build

curl -X POST http://localhost:8080/fhir \
-H "Content-Type: application/fhir+json" \
-d @migrations/transaction-bundle.json
```

## 動作確認
### fhir-mock
```
curl -X GET http://localhost:8080/fhir/Patient
```

### データ出力
```
* Patients
npx ts-node src/interfaces/cli/export-patients.ts

% cat outputs/patients-20250409.csv 
ID,Identifier,Active,Name,Gender,Birth Date,Deceased,Address,Phone,Email,Marital Status,Multiple Birth,Languages,Contact Name,Contact Phone,Managing Organization,General Practitioner
patient-001,001,true,Yamada Taro,male,1990-01-01,false,"1-1 Shibuya Building 1 Tokyo, Tokyo, 150-0001, Japan",090-1234-5678,taro.yamada@example.com,Married,false,English,Yamada Ichiro,090-8765-4321,Organization/8,Practitioner/9
patient-002,002,true,Suzuki Hanako,female,1985-05-15,false,"3-3 Shibuya Building 2 Tokyo, Tokyo, 150-0003, Japan",090-2345-6789,hanako.suzuki@example.com,Single,false,Japanese,Suzuki Yukiko,090-1234-9876,Organization/8,Practitioner/9
% 
```
