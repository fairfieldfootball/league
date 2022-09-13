import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';

import MakesUI from '@mmdb/ui';

import App from './app';
import ClientProvider from './client';
import reportWebVitals from './reportWebVitals';
import theme from './theme';

const rootId = 'root';

const root = ReactDOM.createRoot(document.getElementById(rootId) as HTMLElement);
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <ClientProvider>
        <MakesUI
          prefix="ffc"
          themes={{ default: theme }}
          mode="default"
          globalStyles={() => [{ body: { margin: 0 } }, { ['#' + rootId]: { minHeight: '100vh' } }]}
        >
          <App />
        </MakesUI>
      </ClientProvider>
    </BrowserRouter>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
