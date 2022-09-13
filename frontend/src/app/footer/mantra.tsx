import React from 'react';
import { Anchor } from '@mmdb/buttons';
import { Text } from '@mmdb/words';

const Mantra = () => (
  <Text size="deci" mb="0">
    a{' '}
    <Anchor
      colour="blue"
      variant="text"
      baseSize=""
      href="https://football.fantasysports.yahoo.com/league/myfg"
      target="_blank"
      rel="noreferrer"
    >
      fantasy football league
    </Anchor>{' '}
    since 2013.
  </Text>
);

export default Mantra;
