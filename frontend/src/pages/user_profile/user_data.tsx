import React from 'react';

import { List, ListItem } from '@mmdb/lists';
import { makeForm, Form, FormField, SubmitButton, TextInput } from '@mmdb/forms';
import { SpinnerIcon } from '@mmdb/icons';
import { Subheading, Text } from '@mmdb/words';

import { ApiUserData, User } from '../../types';

interface Props {
  user: User;
  saveUser: (data: ApiUserData) => Promise<void>;
}

function UserDataForm({ user, saveUser }: Props) {
  const { doSubmit, errors, submitting, useField } = makeForm(
    { username: user.name, email: user.email, firstName: user.firstName, lastName: user.lastName },
    {
      onSubmit: data => saveUser({ first_name: data.firstName, last_name: data.lastName }).catch(),
      options: {
        firstName: {
          type: 'text',
          checkValue: v => {
            if (!v) {
              return 'must provide first name';
            }
          },
        },
        lastName: {
          type: 'text',
          checkValue: v => {
            if (!v) {
              return 'must provide last name';
            }
          },
        },
      },
    }
  );
  const hasError = !!errors.submit || Object.values(errors.fields).some(err => !!err);
  return (
    <>
      <section>
        <Subheading>Personal Info</Subheading>
        <Form width={{ _: '100%', md: '61.8%' }} display="grid" gridTemplateColumns="auto auto" onSubmit={doSubmit}>
          <FormField label="Username" labelFor="username">
            <TextInput {...useField('username')} readOnly />
          </FormField>
          <FormField label="Email" labelFor="email">
            <TextInput {...useField('email')} readOnly />
          </FormField>
          <FormField label="First Name" labelFor="firstName">
            <TextInput {...useField('firstName')} />
          </FormField>
          <FormField label="Last Name" labelFor="lastName">
            <TextInput {...useField('lastName')} />
          </FormField>
          <SubmitButton disabled={hasError} submitting={submitting} value="Save" />
        </Form>
      </section>
      {hasError && (
        <div>
          <Text>Please resolve the following errors before saving user data:</Text>
          <List>
            {errors.submit && (
              <ListItem>
                <Text colour="red">{`Failed to submit: ${errors.submit}`}</Text>
              </ListItem>
            )}
            {Object.entries(errors.fields)
              .filter(([_, err]) => !!err)
              .map(([field, err]) => (
                <ListItem>
                  <Text key={field} colour="red">
                    {`Invalid ${field}: ${err}`}
                  </Text>
                </ListItem>
              ))}
          </List>
        </div>
      )}
    </>
  );
}

export default UserDataForm;
