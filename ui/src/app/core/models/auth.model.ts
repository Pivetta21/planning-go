export interface IAuthRequest {
    username: string,
    password: string,
    origin?: number,
}

export interface IAuthResponse {
    message: string
}
