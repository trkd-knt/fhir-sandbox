// src/interfaces/cli/export-patients.ts
import { PatientApiRepository } from '../../infrastructure/repositories/PatientApiRepository';
import { PatientCsvWriter } from '../../infrastructure/writers/PatientCsvWriter';
import { ExportPatientsService } from '../../application/services/ExportPatientsService';

(async () => {
  const patientRepository = new PatientApiRepository();
  const csvWriter = new PatientCsvWriter();
  const exportPatientsService = new ExportPatientsService(patientRepository, csvWriter);

  await exportPatientsService.execute();
})();
