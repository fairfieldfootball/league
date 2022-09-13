import React, { useState } from 'react';
import { ErrorBoundary } from 'react-error-boundary';
import { useNavigate, Routes, Route } from 'react-router-dom';

import styled from '@emotion/styled';

import { useClient } from '../client';
import { sizes } from '../theme';

import AuthRoute from './auth_route';
import Footer from './footer';
import Navbar from './navbar';

import HomePage from '../pages/home';

import LeaguePortalPage from '../pages/league_portal';
import RecordBookPage from '../pages/record_book';
import UserProfilePage from '../pages/user_profile';

import LoginPage from '../pages/auth/login';

import FAQPage from '../pages/info/faq';
import PrivacyPolicyPage from '../pages/info/privacy_policy';
import ResponsibleGamingPage from '../pages/info/responsible_gaming';
import TermsOfServicePage from '../pages/info/terms_of_service';

import { ErrorFallback, NotFoundPage } from '../pages/errors';

const StyledHeader = styled.header`
  display: flex;
  justify-content: space-between;
  height: ${sizes.components.navbar.height};
  border-bottom: 1px solid;
`;

const StyledMain = styled.main`
  min-height: calc(100vh - ${sizes.components.navbar.height});
  min-height: -o-calc(100vh - ${sizes.components.navbar.height}); /* opera */
  min-height: -webkit-calc(100vh - ${sizes.components.navbar.height}); /* google, safari */
  min-height: -moz-calc(100vh - ${sizes.components.navbar.height}); /* firefox */
`;

const StyledFooter = styled.footer`
  display: flex;
  justify-content: space-between;
  border-top: 1px solid;
`;

function App() {
  const [errorReset, setErrorReset] = useState(false);
  const { initialized } = useClient();
  const navTo = useNavigate();

  if (!initialized) {
    return <>initializing...</>;
  }

  return (
    <ErrorBoundary
      FallbackComponent={ErrorFallback}
      onError={(err, info) => console.error(`${err.name}: ${err.message}`, info.componentStack.toString())}
      onReset={() => {
        setErrorReset(!errorReset);
        navTo('/');
      }}
      resetKeys={[errorReset]}
    >
      <StyledHeader>
        <Navbar />
      </StyledHeader>
      <StyledMain>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route
            path="league_portal"
            element={
              <AuthRoute>
                <LeaguePortalPage />
              </AuthRoute>
            }
          />
          <Route
            path="record_book"
            element={
              <AuthRoute>
                <RecordBookPage />
              </AuthRoute>
            }
          />
          <Route
            path="user_profile"
            element={
              <AuthRoute>
                <UserProfilePage />
              </AuthRoute>
            }
          />

          {/* auth */}
          <Route path="login" element={<LoginPage />} />

          {/* info */}
          <Route path="faq" element={<FAQPage />} />
          <Route path="privacy_policy" element={<PrivacyPolicyPage />} />
          <Route path="responsible_gaming" element={<ResponsibleGamingPage />} />
          <Route path="terms_of_service" element={<TermsOfServicePage />} />

          {/* errors */}
          <Route path="*" element={<NotFoundPage />} />
        </Routes>
      </StyledMain>
      <StyledFooter>
        <Footer />
      </StyledFooter>
    </ErrorBoundary>
  );
}

export default App;
