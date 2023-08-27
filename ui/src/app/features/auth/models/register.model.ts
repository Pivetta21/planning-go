import { FormControl } from "@angular/forms";

export type RegisterForm = {
    username: FormControl<string>
    password: FormControl<string>
}

export class RegisterRequest {
    username: string
    password: string

    constructor(formValues: Partial<RegisterRequest>) {
        this.username = formValues.username ?? ""
        this.password = formValues.password ?? ""
    }
}

export type RegisterResponse = {
    message: string
}
