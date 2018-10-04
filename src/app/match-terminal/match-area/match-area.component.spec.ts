import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MatchAreaComponent } from './match-area.component';

describe('MatchAreaComponent', () => {
  let component: MatchAreaComponent;
  let fixture: ComponentFixture<MatchAreaComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MatchAreaComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MatchAreaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
