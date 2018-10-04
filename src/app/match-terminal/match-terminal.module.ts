import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatchAreaComponent } from './match-area/match-area.component';
import {MatchTerminalRoutingModule} from './match-terminal-routing.module';
import {FormsModule} from '@angular/forms';
import {HttpModule} from '@angular/http';
import { MatFormFieldModule} from '@angular/material';
import { MatInputModule } from '@angular/material';
import { MatButtonModule } from '@angular/material';
import { MatMenuModule } from '@angular/material';
import { MatSelectModule } from '@angular/material';

@NgModule({
  imports: [
    CommonModule,
    MatchTerminalRoutingModule,
    HttpModule,
    FormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatMenuModule,
    MatSelectModule
  ],
  declarations: [MatchAreaComponent]
})
export class MatchTerminalModule { }
