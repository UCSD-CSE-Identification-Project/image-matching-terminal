import { Component, OnInit } from '@angular/core';
import { UploadService} from './upload.service';
@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.css']
})
export class UploadComponent implements OnInit {
  private link;

  constructor( private l: UploadService) { }

  ngOnInit() {

  }
  submit(): void {
    // alert(this.link)

    this.l.populateDatabase(this.link);
  }


}
