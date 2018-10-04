import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { MatchAreaComponent } from './match-area/match-area.component';
const routes: Routes = [
  { path: 'match-terminal', component: MatchAreaComponent }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class MatchTerminalRoutingModule { }
