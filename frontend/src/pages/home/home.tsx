import React, { useState } from 'react';
import { useErrorHandler } from 'react-error-boundary';
import styled from '@emotion/styled';
import { Anchor, Button } from '@mmdb/buttons';
import { makeForm, Form, FormField, SubmitButton, TextInput } from '@mmdb/forms';
import { Text } from '@mmdb/words';

import { useClient } from '../../client';

const CLIENT_ID = 'dj0yJmk9amQ4ZGJIRUVmZ0FCJmQ9WVdrOU9EWnBNVTVVTXpBbWNHbzlNQS0tJnM9Y29uc3VtZXJzZWNyZXQmc3Y9MCZ4PTc3';

const StyledApp = styled.article`
  text-align: center;
`;

function HomePage() {
  const { client } = useClient();
  const [showForm, setShowForm] = useState(false);
  const { doSubmit, errors, useField } = makeForm(
    { code: '' },
    {
      onSubmit: data => client.api.yahoo().authenticate(data.code),
      validate: data => {
        if (!data.code) {
          throw new Error('must enter code first');
        }
      },
    }
  );

  const fieldCode = useField('code');
  const handleError = useErrorHandler()
  return (
    <StyledApp>
      <Text>You are not logged in with Yahoo! yet</Text>
      <Anchor
        href={`https://api.login.yahoo.com/oauth2/request_auth?client_id=${CLIENT_ID}&redirect_uri=oob&response_type=code`}
        target="_blank"
        rel="noopener noreferrer"
        onClick={() => setShowForm(true)}
      >
        Authenticate with Yahoo!
      </Anchor>
      {showForm && (
        <Form onSubmit={doSubmit}>
          <FormField label="Code:" labelFor="title">
            <TextInput {...fieldCode} />
          </FormField>
          <SubmitButton error={errors.submit} />
        </Form>
      )}
      <Button
        onClick={() =>
          client.api
            .yahoo()
            .game('nfl')
            .catch(handleError)
        }
      >
        Get Games
      </Button>
    </StyledApp>
  );
}

export default HomePage;
