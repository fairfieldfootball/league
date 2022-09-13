import { ApiUserData } from './auth';
import { Version } from './server';

export interface ClientApi {
  health: () => Promise<void>;
  version: () => Promise<Version>;
  errors: () => {
    json: () => {
      basic: () => Promise<void>;
      complete: () => Promise<void>;
    };
    payload: () => Promise<void>;
    text: () => Promise<void>;
  };
  updateUser: (data: ApiUserData) => Promise<void>
  yahoo: () => {
    authenticate: (code: string) => Promise<void>;
    logout: () => Promise<void>;
    game: (game: 'nfl') => Promise<void>;
  };
}
