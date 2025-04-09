// src/infrastructure/writers/PatientCsvWriter.ts
import { createObjectCsvWriter } from 'csv-writer';
import { Patient } from '../../domain/models/Patient';
import * as dotenv from 'dotenv';
import * as path from 'path';
import * as fs from 'fs';

dotenv.config();

export class PatientCsvWriter {
  async write(patients: Patient[], baseName: string): Promise<void> {
    const dir = process.env.OUTPUT_DIR_PATH!;
    const date = new Date().toISOString().slice(0, 10).replace(/-/g, '');
    const filename = `${baseName}-${date}.csv`;
    const fullPath = path.join(dir, filename);

    // ディレクトリが存在しない場合は作成
    if (!fs.existsSync(dir)) {
      fs.mkdirSync(dir, { recursive: true });
    }

    const writer = createObjectCsvWriter({
      path: fullPath,
      header: [
        { id: 'id', title: 'ID' },
        { id: 'identifier', title: 'Identifier' },
        { id: 'active', title: 'Active' },
        { id: 'name', title: 'Name' },
        { id: 'gender', title: 'Gender' },
        { id: 'birthDate', title: 'Birth Date' },
        { id: 'deceased', title: 'Deceased' },
        { id: 'address', title: 'Address' },
        { id: 'phone', title: 'Phone' },
        { id: 'email', title: 'Email' },
        { id: 'maritalStatus', title: 'Marital Status' },
        { id: 'multipleBirth', title: 'Multiple Birth' },
        { id: 'languages', title: 'Languages' },
        { id: 'contactName', title: 'Contact Name' },
        { id: 'contactPhone', title: 'Contact Phone' },
        { id: 'managingOrg', title: 'Managing Organization' },
        { id: 'generalPractitioner', title: 'General Practitioner' },
      ],
    });

    await writer.writeRecords(patients);
    console.log(`✅ CSVを書き出しました: ${fullPath}`);
  }
}
