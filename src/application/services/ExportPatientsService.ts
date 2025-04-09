// src/application/services/ExportPatientsService.ts
import { IPatientRepository } from '../../domain/repositories/IPatientRepository';
import { PatientCsvWriter } from '../../infrastructure/writers/PatientCsvWriter';

export class ExportPatientsService {
  constructor(
    private repository: IPatientRepository,
    private writer: PatientCsvWriter
  ) {}

  async execute(): Promise<void> {
    const patients = await this.repository.fetchAll();
    await this.writer.write(patients, 'patients');
  }
}
