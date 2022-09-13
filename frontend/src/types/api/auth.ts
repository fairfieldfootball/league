interface ApiUser extends ApiUserData {
  _id: string;
  name: string;
  email: string;
  type: string;
  status: string;
}

export interface ApiUserData {
  first_name: string;
  last_name: string;
}

export const toUser = ({ _id, name, email, first_name, last_name, type, status }: ApiUser) => ({
  id: _id,
  name,
  email,
  firstName: first_name,
  lastName: last_name,
  type,
  status,
});

export type User = ReturnType<typeof toUser>;
