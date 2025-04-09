// src/infrastructure/repositories/PatientApiRepository.ts
import axios from 'axios';
import { IPatientRepository } from '../../domain/repositories/IPatientRepository';
import { Patient } from '../../domain/models/Patient';
import * as dotenv from 'dotenv';

dotenv.config();

export class PatientApiRepository implements IPatientRepository {
  async fetchAll(): Promise<Patient[]> {
    const url = process.env.FHIR_API_URL! + '/Patient';
    const response = await axios.get(url);
    const rawData = response.data;

    return rawData.entry.map((entry: any): Patient => ({
      id: entry.resource.id,
      identifier: entry.resource.identifier?.map((i: any) => i.value).join('; ') ?? 'No identifier',
      active: entry.resource.active?.toString() ?? 'false',
      name: this.formatName(entry.resource.name),
      gender: entry.resource.gender ?? 'unknown',
      birthDate: entry.resource.birthDate ?? 'N/A',
      deceased: entry.resource.deceasedBoolean?.toString() ?? 'false',
      address: this.formatAddress(entry.resource.address),
      phone: entry.resource.telecom?.find((t: any) => t.system === 'phone')?.value ?? 'No phone',
      email: entry.resource.telecom?.find((t: any) => t.system === 'email')?.value ?? 'No email',
      maritalStatus: this.formatMaritalStatus(entry.resource.maritalStatus),
      multipleBirth: entry.resource.multipleBirthBoolean?.toString() ?? 'false',
      languages: this.formatLanguages(entry.resource.communication),
      contactName: this.formatContactName(entry.resource.contact),
      contactPhone: entry.resource.contact?.[0]?.telecom?.find((t: any) => t.system === 'phone')?.value ?? 'No contact phone',
      managingOrg: entry.resource.managingOrganization?.reference ?? 'No managing organization',
      generalPractitioner: entry.resource.generalPractitioner?.map((gp: any) => gp.reference).join('; ') ?? 'No general practitioner',
    }));
  }

  private formatName(name: any): string {
    if (name && name.length > 0) {
      const family = name[0]?.family ?? '';
      const given = name[0]?.given?.join(' ') ?? '';
      return `${family} ${given}`.trim() || 'No name';
    }
    return 'No name';
  }

  private formatAddress(address: any): string {
    if (address && address.length > 0) {
      return address.map((a: any) => {
        return `${a.line?.join(' ')} ${a.city}, ${a.state}, ${a.postalCode}, ${a.country}`.trim();
      }).join('; ') || 'No address';
    }
    return 'No address';
  }

  private formatMaritalStatus(maritalStatus: any): string {
    return maritalStatus?.coding?.[0]?.display ?? 'No marital status';
  }

  private formatLanguages(communication: any): string {
    if (communication && communication.length > 0) {
      return communication
        .map((c: any) => {
          const language = c.language?.coding?.[0]?.display;
          return language ?? 'No language'; // 'display' がない場合は 'No language'
        })
        .join('; ') || 'No languages';
    }
    return 'No languages';
  }
  
  private formatContactName(contact: any): string {
    if (contact && contact.length > 0) {
      const name = contact[0]?.name;
      if (name) {
        const family = name?.family ?? '';
        const given = name?.given?.join(' ') ?? '';
        return `${family} ${given}`.trim() || 'No contact name';
      }
    }
    return 'No contact name';
  }
}