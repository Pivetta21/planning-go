import { FormControl } from "@angular/forms";

export type LoginForm = {
    username: FormControl<string>
    password: FormControl<string>
}

export class LoginRequest {
    username: string
    password: string
    origin: number

    constructor(formValues: Partial<LoginRequest>) {
        this.username = formValues.username ?? ""
        this.password = formValues.password ?? ""
        this.origin = 1 
    }
}
