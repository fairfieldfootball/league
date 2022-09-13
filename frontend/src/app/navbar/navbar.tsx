import React from 'react';
import { useNavigate } from 'react-router-dom';

import styled from '@emotion/styled';
import { Button, Link } from '@mmdb/buttons';
import { Dropdown } from '@mmdb/dropdowns';
import { ChevronDown } from '@mmdb/svgs';

import { useClient } from '../../client';
import { LongLogo } from '../../components';

const StyledNav = styled.nav`
  flex: 1;
  display: flex;
  align-items: center;
`;

const StyledNavLink = styled(Link)`
  margin-left: 0.5rem;
  text-decoration: none;
  color: blue;
`;

const StyledUserInfo = styled.div`
  margin-right: 0.5rem;
  flex-basis: 200px;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  justify-content: center;
`;

const HorizontalRuler = styled.hr`
  margin: 0;
`;

function Navbar() {
  const { client, user } = useClient();
  const navigate = useNavigate();
  return (
    <>
      <StyledNavLink margin="auto" padding="0 0.5rem" flexBasis="200px" variant="text" to="/">
        <LongLogo />
      </StyledNavLink>
      <StyledNav>
        <StyledNavLink variant="ghost" to="/league_portal">
          League Portal
        </StyledNavLink>
        <StyledNavLink variant="ghost" to="/record_book">
          Record Book
        </StyledNavLink>
      </StyledNav>
      <StyledUserInfo>
        {!user && (
          <StyledNavLink variant="text" to="/login">Log in</StyledNavLink>
        )}
        {user && (
          <Dropdown
            menuAlign="right"
            trigger={
              <Button variant="text" after={<ChevronDown />}>
                Welcome {user.name}!
              </Button>
            }
          >
            <Link textDecoration="none" colour="text" variant="text" to="user_profile">
              User profile
            </Link>
            <HorizontalRuler />
            <Button variant="text" onClick={() => client.auth.logout().then(() => navigate('/'))}>
              Sign out
            </Button>
          </Dropdown>
        )}
      </StyledUserInfo>
    </>
  );
}

export default Navbar;
