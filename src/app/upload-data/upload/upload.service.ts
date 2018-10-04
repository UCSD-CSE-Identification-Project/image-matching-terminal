import { Injectable } from '@angular/core';
import {Http} from '@angular/http';

@Injectable({
  providedIn: 'root'
})
export class UploadService {

  constructor(private http: Http) { }

  populateDatabase(link: string) {
    this.http.post('http://localhost:12345/uploadToDatabase', JSON.stringify({'link': link})).subscribe(r => {});
  }
}
