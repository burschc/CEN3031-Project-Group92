import { TestBed } from '@angular/core/testing';

import { DecalService } from './decal.service';

describe('DecalService', () => {
  let service: DecalService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DecalService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
