import { FormControl } from "@angular/forms";

export interface IProfileModel {
    username: string;
    createdAt: string;
    sessionLimit: number;
    activeSessions: number;
}

export type ProfileUpdateForm = {
    username: FormControl<string>
}

export class ProfileUpdateRequest {
    username: string | undefined

    constructor(formValues: Partial<ProfileUpdateRequest>) {
        this.username = formValues.username
    }
}