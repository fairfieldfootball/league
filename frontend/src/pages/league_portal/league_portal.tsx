import React from 'react';

import { List, ListItem } from '@mmdb/lists';
import { Heading, Text } from '@mmdb/words';

function LeaguePortalPage() {
  return (
    <article>
      <Heading>League Portal</Heading>
      <Text>Use this page to stay up to date on the league for the on-going season.</Text>
      <section>
        <Heading as="h2">2022-23 NFL Season: Key Dates</Heading>
        <List>
          <ListItem>
            <Text>Sunday, Sep 4: Beer Mile</Text>
          </ListItem>
          <ListItem>
            <Text>Monday, Sep 5: Draft Day, picks start at 8:30p EST</Text>
          </ListItem>
          <ListItem>
            <Text>Thursday, Sep 8: Week 1 begins (first week of the NFL season)</Text>
          </ListItem>
          <ListItem>
            <Text>Saturday, Nov 19: Trde deadline</Text>
          </ListItem>
          <ListItem>
            <Text>Thursday, Dec 8: Week 14 begins (last week of the fantasy regular season)</Text>
          </ListItem>
          <ListItem>
            <Text>Thursday, Dec 15: Week 15 begins (first week the fantasy playoffs)</Text>
          </ListItem>
          <ListItem>
            <Text>Thursday, Dec 29: Championship begins</Text>
          </ListItem>
        </List>
      </section>
    </article>
  );
}

export default LeaguePortalPage;
