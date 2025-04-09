// src/infrastructure/transformers/PatientFlattener.ts
import { Patient } from '../../domain/models/Patient';

export function flattenPatient(entry: any): Patient {
  const resource = entry.resource;
  return {
    id: resource.id ?? '',
    identifier: (resource.identifier ?? []).map((i: any) => i.value).join('; '),
    active: resource.active?.toString() ?? '',
    name: resource.name?.[0]?.text ?? [resource.name?.[0]?.given?.join(' '), resource.name?.[0]?.family].filter(Boolean).join(' ') ?? '',
    gender: resource.gender ?? '',
    birthDate: resource.birthDate ?? '',
    deceased: resource.deceasedBoolean?.toString() ?? resource.deceasedDateTime ?? '',
    address: resource.address?.map((a: any) => a.text ?? '').join('; ') ?? '',
    phone: resource.telecom?.find((t: any) => t.system === 'phone')?.value ?? '',
    email: resource.telecom?.find((t: any) => t.system === 'email')?.value ?? '',
    maritalStatus: resource.maritalStatus?.text ?? '',
    multipleBirth: resource.multipleBirthBoolean?.toString() ?? resource.multipleBirthInteger?.toString() ?? '',
    languages: (resource.communication ?? []).map((c: any) => c.language?.text ?? '').join('; ') ?? '',
    contactName: resource.contact?.[0]?.name?.text ?? '',
    contactPhone: resource.contact?.[0]?.telecom?.find((t: any) => t.system === 'phone')?.value ?? '',
    managingOrg: resource.managingOrganization?.reference ?? '',
    generalPractitioner: (resource.generalPractitioner ?? []).map((gp: any) => gp.reference).join('; ') ?? '',
  };
}
