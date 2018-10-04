import { Injectable } from '@angular/core';
import {Http} from '@angular/http';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class MatchAreaService {

  constructor(private http: Http) { }

  getStrictMatches(key: string) {
    return this.http.post('http://localhost:12345/strictMatches/getMatches',
      JSON.stringify({'key': key}))
      .pipe(map(response => response.json()));

  }

  getPotentialTierOne(key: string) {
    return this.http.post('http://localhost:12345/potentialTierOne/getMatches',
      JSON.stringify({'key': key}))
       .pipe(map(response => response.json()));
  }

  getPotentialTierTwo(key: string) {
    return this.http.post('http://localhost:12345/potentialTierTwo/getMatches', JSON.stringify({'key': key}))
       .pipe(map(response => response.json()));
  }

}
