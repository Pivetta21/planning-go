import { HTTP_INTERCEPTORS } from '@angular/common/http';

import { DefaultInterceptor } from './default-interceptor';
import { AuthInterceptor } from './auth-interceptor';

export const httpInterceptorProviders = [
    { provide: HTTP_INTERCEPTORS, useClass: DefaultInterceptor, multi: true },
    { provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true },
];