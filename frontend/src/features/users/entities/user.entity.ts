export interface User {
  id: number;
  username: string;
  email: string;
  refresh_token: {
    String: string;
    Valid: boolean;
  };
  created_at: string;
  updated_at: string;
}
