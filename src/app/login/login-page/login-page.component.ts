import { Component, OnInit } from '@angular/core';
import { LoginPageService} from '../login-page.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {

  private username;
  private password;

  constructor( private l: LoginPageService) {

  }

  ngOnInit() {
  }

  login(): void {

    let correctCred = 'False';
    this.l.getUsernamePassword(this.username, this.password).subscribe((data) => {
      correctCred = data.correctValue;
      if (correctCred === 'False') {
        alert('wrong username/password combo');
      } else {
        alert('logged in successfully');
      }
    });


  }

  signUp(): void {
    // alert("inside sign up")
    this.l.addUsernamePassword( this.username, this.password);
  }
}
