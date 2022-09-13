import React from 'react';

import { List, ListItem } from '@mmdb/lists';
import { Heading, Subheading, Text } from '@mmdb/words';

function RecordBookPage() {
  return (
    <article>
      <Heading>Record Book</Heading>
      <Text>A look back at seasons passed.</Text>
      <section>
        <Heading as="h2">2021-22 Results</Heading>
        <Subheading>Podium</Subheading>
        <List ordered horizontal>
          <ListItem>
            <Text>Team A</Text>
          </ListItem>
          <ListItem>
            <Text>Team B</Text>
          </ListItem>
          <ListItem>
            <Text>Team C</Text>
          </ListItem>
        </List>
        <Subheading>Gigantic Loser</Subheading>
        <Text>Team D</Text>
      </section>
    </article>
  );
}

export default RecordBookPage;
