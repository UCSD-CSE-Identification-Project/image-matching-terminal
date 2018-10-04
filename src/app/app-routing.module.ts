import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';


const routes: Routes = [
  {
    path: 'login',
    // component: LoginPageComponent
    loadChildren: './login/login.module#LoginModule',
  },
  {
    path: 'upload-data',
    loadChildren: './upload-data/upload-data.module#UploadDataModule'
  },
  {
    path: 'match-terminal',
    loadChildren: './match-terminal/match-terminal.module#MatchTerminalModule'
  },
  { path: '',   redirectTo: '/login', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
