import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { UploadDataRoutingModule } from './upload-data-routing.module';
import { UploadComponent } from './upload/upload.component';
import { UploadService } from './upload/upload.service';

import { MatFormFieldModule} from '@angular/material';
import { MatInputModule } from '@angular/material';
import { MatButtonModule } from '@angular/material';


@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    UploadDataRoutingModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
  ],
  declarations: [UploadComponent],
  providers: [UploadService]
})
export class UploadDataModule { }
