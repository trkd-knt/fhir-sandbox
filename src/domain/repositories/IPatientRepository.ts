// src/domain/repositories/IPatientRepository.ts
import { Patient } from '../models/Patient';

export interface IPatientRepository {
  fetchAll(): Promise<Patient[]>;
}
