import { Injectable } from '@angular/core';
import {Http} from '@angular/http';
import { map } from 'rxjs/operators';
import {environment} from '../../environments/environment';


@Injectable({
  providedIn: 'root'
})
export class LoginPageService {


  constructor(private http: Http) { }

  getUsernamePassword(userName, password: string) {
    return this.http.post('http://localhost:12345/login/check', // `${environment.serverUrl}/login-check`,
      JSON.stringify({'username': userName, 'password': password}))
      .pipe(map(response => response.json()));
  }

  addUsernamePassword(userName, password: string) {
    // this.http.post(`${environment.serverUrl}/login-add`, JSON.stringify({'username': userName, 'password': password}));
    this.http.post('http://localhost:12345/login/add', JSON.stringify({'username': userName, 'password': password})).subscribe(r => {});
  }
}
