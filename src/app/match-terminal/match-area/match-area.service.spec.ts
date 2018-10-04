import { TestBed } from '@angular/core/testing';

import { MatchAreaService } from './match-area.service';

describe('MatchAreaService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: MatchAreaService = TestBed.get(MatchAreaService);
    expect(service).toBeTruthy();
  });
});
