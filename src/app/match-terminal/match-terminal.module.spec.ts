import { MatchTerminalModule } from './match-terminal.module';

describe('MatchTerminalModule', () => {
  let matchTerminalModule: MatchTerminalModule;

  beforeEach(() => {
    matchTerminalModule = new MatchTerminalModule();
  });

  it('should create an instance', () => {
    expect(matchTerminalModule).toBeTruthy();
  });
});
