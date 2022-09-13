import React from 'react';

import { Heading, Text } from '@mmdb/words';

import { useClient } from '../../client';

import UserDataForm from './user_data';

function UserProfilePage() {
  const { client, user } = useClient();

  if (!user) {
    return null;
  }

  return (
    <article>
      <Heading>User Profile</Heading>
      <Text>The user data we have stored on you.</Text>
      <UserDataForm
        user={user}
        saveUser={data =>
          client.api.updateUser(data).then(() => {
            client.auth.whoami();
          })
        }
      />
    </article>
  );
}

export default UserProfilePage;
