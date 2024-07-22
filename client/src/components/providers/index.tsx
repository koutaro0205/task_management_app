'use client';

import React from 'react';

import { UrqlProvider } from './UrqlProvider';

type Props = {
  children: React.ReactNode;
};

export const AppProviders: React.FC<Props> = React.memo(({ children }) => {
  return <UrqlProvider>{children}</UrqlProvider>;
});
