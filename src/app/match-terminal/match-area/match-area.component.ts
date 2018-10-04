import { Component, OnInit } from '@angular/core';
import {MatchAreaService} from './match-area.service';
import { ViewEncapsulation} from '@angular/core';

@Component({
  selector: 'app-match-area',
  templateUrl: './match-area.component.html',
  styleUrls: ['./match-area.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class MatchAreaComponent implements OnInit {

  private imageName;
  private imageFolderPath = '/assets/Images/';
  private strictmatches$: Object;
  private tier1matches$: Object;
  private tier2matches$: Object;
  private imageSource;
  displayPictures: boolean;


  constructor(private s: MatchAreaService) {
    this.displayPictures = false;
  }

  ngOnInit() {
  }

  matchFound( matchName: string): void {
    // update in database that imagename has a match in matchName
    alert(matchName);
  }
  submit(): void {
    this.imageSource = this.imageFolderPath + this.imageName;
    this.populateMatches();
  }

  populateMatches(): void {
    this.s.getStrictMatches(this.imageName).subscribe((data) => {
      this.strictmatches$ = data;
      alert(this.strictmatches$);
    });
    this.s.getPotentialTierOne(this.imageName).subscribe( (data) => {
      this.tier1matches$ = data;
      // alert(this.tier1matches$);
    });
    this.s.getPotentialTierTwo(this.imageName).subscribe((data) => {
      this.tier2matches$ = data;
      // alert(this.tier2matches$);
    });
  }

}
