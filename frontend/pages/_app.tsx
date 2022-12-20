import React, { useState } from 'react';
import type { AppProps } from 'next/app'
import { QueryClientProvider, QueryClient } from '@tanstack/react-query';

import 'the-new-css-reset/css/reset.css'
import '../styles/globals.css'
import { AuthContext } from '../contexts/auth';

const queryClient = new QueryClient();

export default function App({ Component, pageProps }: AppProps) {
  const [token, setToken] = useState('');
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const login = (token: string) => {
    setToken(token);
    setIsLoggedIn(true);
  };
  const authValue = { token, isLoggedIn, login };

  return (
    <AuthContext.Provider value={authValue}>
      <QueryClientProvider client={queryClient}>
        <Component {...pageProps} />
      </QueryClientProvider>
    </AuthContext.Provider>
  );
}
